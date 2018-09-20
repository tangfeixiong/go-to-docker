package pb

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

type ImageBuildMetadata struct {
	content []byte
}

func NewImageBuildMetadata(buildcontext []byte) *ImageBuildMetadata {
	return &ImageBuildMetadata{
		content: buildcontext,
	}
}

// Treat as Dockerfile as find directive FROM and return instruction:
// - image reference (e.g. include :tag or @digest)
// - alias name
func (m *ImageBuildMetadata) AsDockerfile() (string, string, bool) {
	return m.dockerfileDirectiveFROM()
}

func (m *ImageBuildMetadata) IsDockerfile() bool {
	_, _, ok := m.dockerfileDirectiveFROM()
	return ok
}

// Dockerfile directive FROM
func (m *ImageBuildMetadata) dockerfileDirectiveFROM() (string, string, bool) {
	var ndx int
	var start, stop int = 0, -1
	var imgref, alias string = "", ""
	var err error = nil
	var ok *bool = nil
	for ndx = 0; ndx < len(m.content); ndx++ {
		if m.content[ndx] == byte('\n') {
			stop = ndx
			if start == stop {
				start = ndx + 1
				continue // blank line
			}
			line := m.content[start:stop]
			if line[0] == byte('#') {
				start = ndx + 1
				continue // comment line
			}
			line = m.trimPrefixBlank(line)
			fmt.Println(string(line))
			if ok == nil {
				if len(line) > 4 &&
					(line[0] == byte('F') || line[0] == byte('f')) &&
					(line[1] == byte('R') || line[1] == byte('r')) &&
					(line[2] == byte('O') || line[2] == byte('o')) &&
					(line[3] == byte('M') || line[3] == byte('m')) {
					instruction := m.trimBlank(line[4:])
					imgref, alias, err = m.dockefileInstructionFROM(instruction)
					ok = new(bool)
					if err != nil {
						// return imgref, alias, false
						*ok = false
					}
					// return imgref, alias, true
					*ok = true
				}
			}
			start = ndx + 1
		}
	}
	if ok == nil {
		fmt.Printf("\x1b[36;1m%q\x1b[0m\n", m.content)
		return "", "", false
	}
	if *ok == false {
		fmt.Fprintf(os.Stderr, "\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %v", err))
	} else {
		fmt.Fprintf(os.Stdout, "\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf("imgref=%s alias=%s", imgref, alias))
	}
	return imgref, alias, *ok
}

func (m *ImageBuildMetadata) dockefileInstructionFROM(bs []byte) (string, string, error) {
	imgref, remainder := m.contentBeforeBlank(bs)
	if len(remainder) == 0 {
		return string(imgref), "", nil
	}
	guess, remainder := m.contentBeforeBlank(remainder)
	if len(guess) == 2 &&
		(guess[0] == byte('A') || guess[0] == byte('a')) &&
		(guess[1] == byte('S') || guess[1] == byte('s')) {
		alias, remainder := m.contentBeforeBlank(remainder)
		if len(alias) == 0 || len(remainder) > 0 {
			return string(imgref), string(alias), errors.New("unexpected alias format")
		}
		return string(imgref), string(alias), nil
	}
	return string(imgref), "", errors.New("unexpected format")
}

func (m *ImageBuildMetadata) trimPrefixBlank(bs []byte) []byte {
	for index := 0; index < len(bs); index++ {
		if bs[index] != byte(' ') {
			return bs[index:]
		}
	}
	return bs
}

func (m *ImageBuildMetadata) trimSuffixBlank(bs []byte) []byte {
	for index := len(bs) - 1; index >= 0; index-- {
		if bs[index] != byte(' ') {
			return bs[:index+1]
		}
	}
	return bs
}

func (m *ImageBuildMetadata) trimBlank(bs []byte) []byte {
	tr := m.trimSuffixBlank(bs)
	return m.trimPrefixBlank(tr)
}

func (m *ImageBuildMetadata) contentBeforeBlank(bs []byte) ([]byte, []byte) {
	bs = m.trimPrefixBlank(bs)
	for index := 0; index < len(bs); index++ {
		if bs[index] == byte(' ') {
			return bs[:index], bs[index:]
		}
	}
	return bs, []byte(nil)
}

/*
According https://docs.docker.com/engine/reference/commandline/build/#git-repositories

$ docker build https://github.com/docker/rootfs.git#container:docker
The following table represents all the valid suffixes with their build contexts:

Build Syntax Suffix	            Commit Used	            Build Context Used
myrepo.git	                    refs/heads/master	    /
myrepo.git#mytag	            refs/tags/mytag	        /
myrepo.git#mybranch	            refs/heads/mybranch	/
myrepo.git#pull/42/head	        refs/pull/42/head	    /
myrepo.git#:myfolder	        refs/heads/master	    /myfolder
myrepo.git#master:myfolder	    refs/heads/master	    /myfolder
myrepo.git#mytag:myfolder	    refs/tags/mytag	        /myfolder
myrepo.git#mybranch:myfolder	refs/heads/mybranch	    /myfolder
*/

/*
According https://www.git-scm.com/docs/git-clone#_git_urls_a_id_urls_a

The following syntaxes may be used with them:

ssh://[user@]host.xz[:port]/path/to/repo.git/
git://host.xz[:port]/path/to/repo.git/
http[s]://host.xz[:port]/path/to/repo.git/
ftp[s]://host.xz[:port]/path/to/repo.git/

An alternative scp-like syntax may also be used with the ssh protocol:

[user@]host.xz:path/to/repo.git/

This syntax is only recognized if there are no slashes before the first colon. This helps differentiate a local path that contains a colon. For example the local path foo:bar could be specified as an absolute path or ./foo:bar to avoid being misinterpreted as an ssh url.

The ssh and git protocols additionally support ~username expansion:

ssh://[user@]host.xz[:port]/~[user]/path/to/repo.git/
git://host.xz[:port]/~[user]/path/to/repo.git/
[user@]host.xz:/~[user]/path/to/repo.git/

For local repositories, also supported by Git natively, the following syntaxes may be used:

/path/to/repo.git/
file:///path/to/repo.git/
*/

// Must match first line:
//    # ** metadata:git **
// then other comment lines, include git access control directive:
//    ...
// And last, must only single content line:
//    https://github.com/tangfeixiong/nta#:docs
func (m *ImageBuildMetadata) AsGitReopsitory() (string, string, string, bool) {
	return m.examineGitDescription()
}

func (m *ImageBuildMetadata) IsGitReopsitory() bool {
	_, _, _, ok := m.examineGitDescription()
	return ok
}

func (m *ImageBuildMetadata) examineGitDescription() (string, string, string, bool) {
	var ndx int
	var start, stop int = 0, -1
	var url, ref, dir string = "", "", ""
	var err error = nil
	var ok *bool = nil
	for ndx = 0; ndx < len(m.content); ndx++ {
		if m.content[ndx] == byte('\n') {
			stop = ndx
			if start == stop {
				start = ndx + 1
				continue // empty line
			}
			line := m.content[start:stop]
			if line[0] == byte('#') {
				if start == 0 && !bytes.Equal(line, []byte("# ** metadata:git **")) {
					err = errors.New("unexpected first specifal line")
					ok = new(bool)
					// return "", "", "", false
					*ok = false
				} else {
					// TODO: ensure metadata
					// comment line
				}
				start = ndx + 1
				continue
			}

			line = bytes.TrimSpace(line)
			fmt.Println(string(line))
			switch {
			case bytes.HasPrefix(line, []byte("HTTPS://")) ||
				bytes.HasPrefix(line, []byte("https://")) ||
				bytes.HasPrefix(line, []byte("HTTP://")) ||
				bytes.HasPrefix(line, []byte("http://")):
				if len(url) > 0 {
					err = errors.New("multi repo url not permitted")
					ok = new(bool)
					// return url, ref, dir, false
					*ok = false
					break
				}
				if i := bytes.IndexByte(line, byte('#')); i > 0 {
					url = string(line[:i])
					remainder := line[i+1:]
					if len(remainder) > 0 {
						if j := bytes.IndexByte(remainder, byte(':')); j > 0 {
							ref = string(remainder[:j])
							dir = string(remainder[j+1:])
						} else if remainder[0] != byte(':') {
							ref = string(remainder[1:])
						} else {
							dir = string(remainder[1:])
						}
					}
				} else {
					url = string(line)
				}
				ok = new(bool)
				// return url, ref, dir, true
				*ok = true
			}

			start = ndx + 1
		}
	}
	if ok == nil {
		fmt.Printf("\x1b[36;1m%q\x1b[0m\n", m.content)
		return "", "", "", false
	}
	if *ok == false {
		fmt.Fprintf(os.Stderr, "\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %v", err))
	} else {
		fmt.Fprintf(os.Stdout, "\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf("url=%s ref=%s dir=%s", url, ref, dir))
	}
	return url, ref, dir, *ok
}
