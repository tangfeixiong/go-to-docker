package dockerclient

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/golang/glog"
)

type shouldExcludePathFn func(string) bool

// CreateTarStreamToTarWriter creates a tar stream on the given writer from
// the given directory while excluding files that match the given
// exclusion pattern.
func CreateTarStreamToTarWriter(dir string, includeDirInPath bool, tarWriter *tar.Writer, logger io.Writer, shouldExclude shouldExcludePathFn) error {
	dir = filepath.Clean(dir) // remove relative paths and extraneous slashes
	glog.Infof("Adding %q to tar ...", dir)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// on Windows, directory symlinks report as a directory and as a symlink.
		// They should be treated as symlinks.
		if !shouldExclude(path) {
			// if file is a link just writing header info is enough
			if info.Mode()&os.ModeSymlink != 0 {
				if dir == path {
					return nil
				}
				if err = writeTarHeader(tarWriter, dir, path, info, includeDirInPath, logger); err != nil {
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
				if err = writeTarHeader(tarWriter, dir, path, info, includeDirInPath, logger); err != nil {
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
			if err = writeTarHeader(tarWriter, dir, path, info, includeDirInPath, logger); err != nil {
				glog.Errorf("Error writing header for %q: %v", info.Name(), err)
				return err
			}
			if _, err = io.Copy(tarWriter, file); err != nil {
				glog.Errorf("Error copying file %q to tar: %v", path, err)
				return err
			}
		}
		if info.IsDir() && path != dir {
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		glog.Errorf("Error writing tar: %v", err)
		return err
	}

	return nil
}

func writeTarHeader(tarWriter *tar.Writer, dir string, path string, info os.FileInfo, includeDirInPath bool, logger io.Writer) error {
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
	glog.Infof("Adding to tar: %s as %s", path, header.Name)
	return tarWriter.WriteHeader(header)
}

func logFile(logger io.Writer, name string) {
	if logger == nil {
		return
	}
	fmt.Fprintf(logger, "%s\n", name)
}
