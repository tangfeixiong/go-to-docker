package manifest

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/golang/glog"

	"github.com/tangfeixiong/go-to-docker/pkg/api/artifact"
)

type Manifest struct {
	DockerfileName string
	DockerfileDir  string
	IncludeDirSelf bool
}

func (m *Manifest) FromGofileserverTemplate() (string, error) {
	if len(m.DockerfileName) == 0 {
		m.DockerfileName = "Dockerfile.gofileserver-based"
	}
	asset, err := artifact.Asset("template/Dockerfile.gofileserver-based.go-tpl")
	if err != nil {
		glog.Errorf("Unable to get template: %v", err)
		return "", err
	}

	buf := bytes.Buffer{}
	dstDir := ""
	if m.IncludeDirSelf {
		buf.WriteString("/")
	} else {
		//	fs, err := ioutil.ReadDir(".")
		fs, err := ioutil.ReadDir(dir)
		if err != nil {
			glog.Errorf("Unable to find specified dir: %v", err)
			return "", err
		}
		dir = filepath.Clean(dir)
		dstDir = filepath.Base(dir)
		for len(dstDir) > 0 {
			if dstDir[0] == byte('/') || dstDir[0] == byte('.') {
				dstDir = dstDir[1:]
			}
		}
		// re := regexp.MustCompile(`\.git|\.cache|\.m2|\.gradle|\.docker|\.kube|\.ssh|\.vagrant(\.d)?|\.eclipse|\.npm|\.glide`)
		re := regexp.MustCompile(`\..*`)
		for _, fi := range fs {
			if f.Mode().IsRegular() {
				if re.MatchString(fi.Name()) == false {
					buf.WriteString(f.Name())
					buf.WriteString(" ")
				}
			}
		}
	}

	if buf.Len() == 0 {
		glog.Warningln("None file or directory find")
	}

	dockerfileBuf := &bytes.Buffer{}
	tte := template.Must(template.New("basedgofs").Parse(string(asset)))
	err = tte.Execute(dockerfileBuf, struct{ SrcStr string }{
		SrcStr: "/",
		dstDir: dstDir,
	})
	if err != nil {
		glog.Errorf("Unexpected to execute gofileserver dockerfile template: %v", err)
		return "", err
	}

	if len(m.DockerfileDir) == 0 {
		m.DockerfileDir, err = ioutil.TempDir(os.TempDir(), "")
		if err != nil {
			glog.Errorf("Unable to create temp dir to save dockerfile", err)
			return "", err
		}
	}
	dockerfilePath := filepath.Join(m.DockerfileDir, m.DockerfileName)
	err = ioutil.WriteFile(dockerfilePath, dockerfileBuf.Bytes(), 0644)
	if err != nil {
		glog.Errorf("Unable to write out dockerfile: ", err)
		return "", err
	}
	return dockerfilePath, nil
}
