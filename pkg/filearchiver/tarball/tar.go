package tarball

import (
	"archive/tar"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/golang/glog"

	utiltimeout "github.com/tangfeixiong/go-to-docker/pkg/util/timeout"
)

// defaultTimeout is the amount of time that the untar will wait for a tar
// stream to extract a single file. A timeout is needed to guard against broken
// connections in which it would wait for a long time to untar and nothing would happen
const defaultTimeout = 30 * time.Second

// DefaultExclusionPattern is the pattern of files that will not be included in a tar
// file when creating one. By default it is any file inside a .git metadata directory
var DefaultExclusionPattern = regexp.MustCompile(`(^|/)\.git(/|$)`)

// ChmodAdapter changes the mode of files and directories inline as a tarfile is
// being written
type ChmodAdapter struct {
	*tar.Writer
	NewFileMode     int64
	NewExecFileMode int64
	NewDirMode      int64
}

// WriteHeader changes the mode of files and directories inline as a tarfile is
// being written
func (a ChmodAdapter) WriteHeader(hdr *tar.Header) error {
	if hdr.FileInfo().Mode()&os.ModeSymlink == 0 {
		newMode := hdr.Mode &^ 0777
		if hdr.FileInfo().IsDir() {
			newMode |= a.NewDirMode
		} else if hdr.FileInfo().Mode()&0010 != 0 { // S_IXUSR
			newMode |= a.NewExecFileMode
		} else {
			newMode |= a.NewFileMode
		}
		hdr.Mode = newMode
	}
	return a.Writer.WriteHeader(hdr)
}

// RenameAdapter renames files and directories inline as a tarfile is being
// written
type RenameAdapter struct {
	*tar.Writer
	Old string
	New string
}

// WriteHeader renames files and directories inline as a tarfile is being
// written
func (a RenameAdapter) WriteHeader(hdr *tar.Header) error {
	if hdr.Name == a.Old {
		hdr.Name = a.New
	} else if strings.HasPrefix(hdr.Name, a.Old+"/") {
		hdr.Name = a.New + hdr.Name[len(a.Old):]
	}

	return a.Writer.WriteHeader(hdr)
}

// New creates a new Tar
func New() *TapArchiver {
	return &TapArchiver{
		fs:      NewFileSystem(),
		exclude: DefaultExclusionPattern,
		timeout: defaultTimeout,
	}
}

// NewParanoid creates a new Tar that has restrictions
// on what it can do while extracting files.
func NewParanoid() *TapArchiver {
	return &TapArchiver{
		fs:                   NewFileSystem(),
		exclude:              DefaultExclusionPattern,
		timeout:              defaultTimeout,
		disallowOverwrite:    true,
		disallowOutsidePaths: true,
		disallowSpecialFiles: true,
	}
}

// stiTar is an implementation of the Tar interface
type TapArchiver struct {
	*fs
	timeout              time.Duration
	exclude              *regexp.Regexp
	includeDirInPath     bool
	disallowOverwrite    bool
	disallowOutsidePaths bool
	disallowSpecialFiles bool
}

// SetExclusionPattern sets the exclusion pattern for tar creation.  The
// exclusion pattern always uses UNIX-style (/) path separators, even on
// Windows.
func (t *TapArchiver) SetExclusionPattern(p *regexp.Regexp) {
	t.exclude = p
}

// CreateTarFile creates a tar file from the given directory
// while excluding files that match the given exclusion pattern
// It returns the name of the created file
func (t *TapArchiver) CreateTarFile(base, dir string) (string, error) {
	tarFile, err := ioutil.TempFile(base, "tar")
	defer tarFile.Close()
	if err != nil {
		return "", err
	}
	if err = t.CreateTarStream(dir, false, tarFile); err != nil {
		return "", err
	}
	return tarFile.Name(), nil
}

func (t *TapArchiver) shouldExclude(path string) bool {
	return t.exclude != nil && t.exclude.String() != "" && t.exclude.MatchString(filepath.ToSlash(path))
}

// CreateTarStream calls CreateTarStreamToTarWriter with a nil logger
func (t *TapArchiver) CreateTarStream(dir string, includeDirInPath bool, writer io.Writer) error {
	tarWriter := tar.NewWriter(writer)
	defer tarWriter.Close()

	return t.CreateTarStreamToTarWriter(dir, includeDirInPath, tarWriter, nil)
}

// CreateTarStreamReader returns an io.ReadCloser from which a tar stream can be
// read.  The tar stream is created using CreateTarStream.
func (t *TapArchiver) CreateTarStreamReader(dir string, includeDirInPath bool) io.ReadCloser {
	r, w := io.Pipe()
	go func() {
		w.CloseWithError(t.CreateTarStream(dir, includeDirInPath, w))
	}()
	return r
}

// CreateTarStreamToTarWriter creates a tar stream on the given writer from
// the given directory while excluding files that match the given
// exclusion pattern.
func (t *TapArchiver) CreateTarStreamToTarWriter(dir string, includeDirInPath bool, tarWriter *tar.Writer, logger io.Writer) error {
	dir = filepath.Clean(dir) // remove relative paths and extraneous slashes
	glog.V(5).Infof("Adding %q to tar ...", dir)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// on Windows, directory symlinks report as a directory and as a symlink.
		// They should be treated as symlinks.
		if !t.shouldExclude(path) {
			// if file is a link just writing header info is enough
			if info.Mode()&os.ModeSymlink != 0 {
				if dir == path {
					return nil
				}
				if err = t.writeTarHeader(tarWriter, dir, path, info, includeDirInPath, logger); err != nil {
					glog.Errorf("Error writing header for %q: %v", info.Name(), err)
				}
				// on Windows, filepath.Walk recurses into directory symlinks when it
				// shouldn't.  https://github.com/golang/go/issues/17540
				if err == nil && info.Mode()&os.ModeDir != 0 {
					return filepath.SkipDir
				}
				return err
			}
			if info.IsDir() {
				if dir == path {
					return nil
				}
				if err = t.writeTarHeader(tarWriter, dir, path, info, includeDirInPath, logger); err != nil {
					glog.Errorf("Error writing header for %q: %v", info.Name(), err)
				}
				return err
			}

			// regular files are copied into tar, if accessible
			file, err := os.Open(path)
			if err != nil {
				glog.Errorf("Ignoring file %s: %v", path, err)
				return nil
			}
			defer file.Close()
			if err = t.writeTarHeader(tarWriter, dir, path, info, includeDirInPath, logger); err != nil {
				glog.Errorf("Error writing header for %q: %v", info.Name(), err)
				return err
			}
			if _, err = io.Copy(tarWriter, file); err != nil {
				glog.Errorf("Error copying file %q to tar: %v", path, err)
				return err
			}
		}
		return nil
	})

	if err != nil {
		glog.Errorf("Error writing tar: %v", err)
		return err
	}

	return nil
}

// writeTarHeader writes tar header for given file, returns error if operation fails
func (t *TapArchiver) writeTarHeader(tarWriter *tar.Writer, dir string, path string, info os.FileInfo, includeDirInPath bool, logger io.Writer) error {
	var (
		link string
		err  error
	)
	if info.Mode()&os.ModeSymlink != 0 {
		link, err = os.Readlink(path)
		if err != nil {
			return err
		}
	}
	header, err := tar.FileInfoHeader(info, link)
	if err != nil {
		return err
	}
	// on Windows, tar.FileInfoHeader incorrectly interprets directory symlinks
	// as directories.  https://github.com/golang/go/issues/17541
	if info.Mode()&os.ModeSymlink != 0 && info.Mode()&os.ModeDir != 0 {
		header.Typeflag = tar.TypeSymlink
		header.Mode &^= 040000 // c_ISDIR
		header.Mode |= 0120000 // c_ISLNK
		header.Linkname = link
	}
	prefix := dir
	if includeDirInPath {
		prefix = filepath.Dir(prefix)
	}
	fileName := path
	if prefix != "." {
		fileName = path[1+len(prefix):]
	}
	header.Name = filepath.ToSlash(fileName)
	header.Linkname = filepath.ToSlash(header.Linkname)
	logFile(logger, header.Name)
	glog.V(5).Infof("Adding to tar: %s as %s", path, header.Name)
	return tarWriter.WriteHeader(header)
}

// ExtractTarStream calls ExtractTarStreamFromTarReader with a default reader and nil logger
func (t *TapArchiver) ExtractTarStream(dir string, reader io.Reader) error {
	tarReader := tar.NewReader(reader)
	return t.ExtractTarStreamFromTarReader(dir, tarReader, nil)
}

// ExtractTarStreamWithLogging calls ExtractTarStreamFromTarReader with a default reader
func (t *TapArchiver) ExtractTarStreamWithLogging(dir string, reader io.Reader, logger io.Writer) error {
	tarReader := tar.NewReader(reader)
	return t.ExtractTarStreamFromTarReader(dir, tarReader, logger)
}

// ExtractTarStreamFromTarReader extracts files from a given tar stream.
// Times out if reading from the stream for any given file
// exceeds the value of timeout
func (t *TapArchiver) ExtractTarStreamFromTarReader(dir string, tarReader *tar.Reader, logger io.Writer) error {
	err := utiltimeout.TimeoutAfter(t.timeout, "", func(timeoutTimer *time.Timer) error {
		for {
			header, err := tarReader.Next()
			if !timeoutTimer.Stop() {
				return &utiltimeout.TimeoutError{}
			}
			timeoutTimer.Reset(t.timeout)
			if err == io.EOF {
				return nil
			}
			if err != nil {
				glog.Errorf("Error reading next tar header: %v", err)
				return err
			}

			if t.disallowSpecialFiles {
				switch header.Typeflag {
				case tar.TypeReg, tar.TypeRegA, tar.TypeLink, tar.TypeSymlink, tar.TypeDir, tar.TypeGNUSparse:
				default:
					glog.Warningf("Skipping special file %s, type: %v", header.Name, header.Typeflag)
					continue
				}
			}

			p := header.Name
			if t.disallowOutsidePaths {
				p = filepath.Clean(filepath.Join(dir, p))
				if !strings.HasPrefix(p, dir) {
					glog.Warningf("Skipping relative path file in tar: %s", header.Name)
					continue
				}
			}

			if header.FileInfo().IsDir() {
				dirPath := filepath.Join(dir, filepath.Clean(header.Name))
				glog.V(3).Infof("Creating directory %s", dirPath)
				if err = os.MkdirAll(dirPath, 0700); err != nil {
					glog.Errorf("Error creating dir %q: %v", dirPath, err)
					return err
				}
				t.Chmod(dirPath, header.FileInfo().Mode())
			} else {
				fileDir := filepath.Dir(header.Name)
				dirPath := filepath.Join(dir, filepath.Clean(fileDir))
				glog.V(3).Infof("Creating directory %s", dirPath)
				if err = os.MkdirAll(dirPath, 0700); err != nil {
					glog.Errorf("Error creating dir %q: %v", dirPath, err)
					return err
				}
				if header.Typeflag == tar.TypeSymlink {
					if err := t.extractLink(dir, header, tarReader); err != nil {
						glog.Errorf("Error extracting link %q: %v", header.Name, err)
						return err
					}
					continue
				}
				logFile(logger, header.Name)
				if err := t.extractFile(dir, header, tarReader); err != nil {
					glog.Errorf("Error extracting file %q: %v", header.Name, err)
					return err
				}
			}
		}
	})

	if err != nil {
		glog.Error("Error extracting tar stream")
	} else {
		glog.V(2).Info("Done extracting tar stream")
	}

	if utiltimeout.IsTimeoutError(err) {
		err = NewTarTimeoutError()
	}

	return err
}

func (t *TapArchiver) extractLink(dir string, header *tar.Header, tarReader io.Reader) error {
	dest := filepath.Join(dir, header.Name)
	source := header.Linkname

	if t.disallowOutsidePaths {
		target := filepath.Clean(filepath.Join(dest, "..", source))
		if !strings.HasPrefix(target, dir) {
			glog.Warningf("Skipping symlink that points to relative path: %s", header.Linkname)
			return nil
		}
	}

	if t.disallowOverwrite {
		if _, err := os.Stat(dest); !os.IsNotExist(err) {
			glog.Warningf("Refusing to overwrite existing file: %s", dest)
			return nil
		}
	}

	glog.V(3).Infof("Creating symbolic link from %q to %q", dest, source)

	// TODO: set mtime for symlink (unfortunately we can't use os.Chtimes() and probably should use syscall)
	return os.Symlink(source, dest)
}

func (t *TapArchiver) extractFile(dir string, header *tar.Header, tarReader io.Reader) error {
	path := filepath.Join(dir, header.Name)
	if t.disallowOverwrite {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			glog.Warningf("Refusing to overwrite existing file: %s", path)
			return nil
		}
	}

	glog.V(3).Infof("Creating %s", path)

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	// The file times need to be modified after it's been closed thus this function
	// is deferred after the file close (LIFO order for defer)
	defer os.Chtimes(path, time.Now(), header.FileInfo().ModTime())
	defer file.Close()
	glog.V(3).Infof("Extracting/writing %s", path)
	written, err := io.Copy(file, tarReader)
	if err != nil {
		return err
	}
	if written != header.Size {
		return fmt.Errorf("wrote %d bytes, expected to write %d", written, header.Size)
	}
	return t.Chmod(path, header.FileInfo().Mode())
}

// NewFileSystem creates a new instance of the default FileSystem
// implementation
func NewFileSystem() *fs {
	return &fs{
		fileModes:    make(map[string]os.FileMode),
		keepSymlinks: false,
	}
}

type fs struct {
	// on Windows, fileModes is used to track the UNIX file mode of every file we
	// work with; m is used to synchronize access to fileModes.
	fileModes    map[string]os.FileMode
	m            sync.Mutex
	keepSymlinks bool
}

// Chmod sets the file mode
func (h *fs) Chmod(file string, mode os.FileMode) error {
	err := os.Chmod(file, mode)
	if runtime.GOOS == "windows" && err == nil {
		h.m.Lock()
		h.fileModes[file] = mode
		h.m.Unlock()
		return nil
	}
	return err
}

func logFile(logger io.Writer, name string) {
	if logger == nil {
		return
	}
	fmt.Fprintf(logger, "%s\n", name)
}

const (
	WorkdirError int = 1 + iota
	TarTimeoutError
	DownloadError
	URLHandlerError
	SourcePathError
	UserNotAllowedError
	EmptyGitRepositoryError
)

// Error represents an error thrown during S2I execution
type Error struct {
	Message    string
	Details    error
	ErrorCode  int
	Suggestion string
}

// Error returns a string for a given error
func (s Error) Error() string {
	return s.Message
}

// NewTarTimeoutError returns a new error which indicates there was a problem
// when sending or receiving tar stream
func NewTarTimeoutError() error {
	return Error{
		Message:    fmt.Sprintf("timeout waiting for tar stream"),
		Details:    nil,
		ErrorCode:  TarTimeoutError,
		Suggestion: "check if it accepts tar stream for any one",
	}
}
