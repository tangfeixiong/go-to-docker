package manifest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"text/template"

	"github.com/golang/glog"

	"github.com/tangfeixiong/go-to-docker/pkg/api"
	"github.com/tangfeixiong/go-to-docker/pkg/api/artifact"
)

type DockerfileManifest struct {
	Filename  string
	SavingDir string
	// CreatingImgDir string
}

// Create object
// parameters:
// - expecting Dockerfile name, default is Dockerfile
// - expecting directory to save Dockerfile
func NewDockerfileManifest(dockerfileName, savingDir string) *DockerfileManifest {
	return &DockerfileManifest{
		Filename:  dockerfileName,
		SavingDir: savingDir,
	}
}

// Using go template to generate default Dockerfile that is building file server
// parameters:
// - dir must write into Dockerfile e.g. directive COPY ... .../<imgDirCreate>
// - Directory of Docker build context, in local file system
// - Rules filting file and dirs
// return:
// - target saved in current file system
// - occurs in reading template, executing template, reading directory, writing Dockerfile etc.
func (m *DockerfileManifest) WriteFileUsingGofileserverbasedTemplate(imgDirCreate, dir string, rules map[string][]string) (string, error) {
	if len(m.Filename) == 0 {
		m.Filename = "Dockerfile.gofileserver-based"
	}
	if len(m.SavingDir) == 0 {
		m.SavingDir = filepath.Join(os.TempDir(), api.DefaultProfileDir[1:], "dockerbuildsource")
	}

	asset, err := artifact.Asset("template/Dockerfile.gofileserver-based.go-tpl")
	if err != nil {
		return "", err
	}

	buf := bytes.Buffer{}
	if len(dir) != 0 {
		// re := regexp.MustCompile(`\.git|\.cache|\.m2|\.gradle|\.docker|\.kube|\.ssh|\.vagrant(\.d)?|\.eclipse|\.npm|\.glide`)
		re := regexp.MustCompile(`\..*`)
		//	fs, err := ioutil.ReadDir(".")
		fs, err := ioutil.ReadDir(dir)
		if err != nil {
			glog.Warningf("Unable to find specified dir: %v", err)
		}
		for _, fi := range fs {
			if fi.Mode().IsRegular() {
				if re.MatchString(fi.Name()) == false {
					buf.WriteString(fi.Name())
					buf.WriteString(" ")
				}
			}
			for k, v := range rules {
				switch k {
				case "inclusions":
					break
				case "regexpinc":
					break
				case "exclusions":
					break
				case "regexpexc":
					break
				default:
					glog.Warningf("Unknown rule %s %v", k, v)
				}
			}
		}
		if buf.Len() == 0 {
			glog.Warningln("None file or directory find")
		}
	}

	buf.Reset()
	tte := template.Must(template.New("basedgofs").Parse(string(asset)))
	err = tte.Execute(&buf, struct{ SrcStr, DstDir string }{
		SrcStr: "/",
		DstDir: imgDirCreate,
	})
	if err != nil {
		return "", err
	}

	fi, err := os.Stat(m.SavingDir)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(m.SavingDir, os.ModePerm); err != nil {
				return "", err
			}
		} else if !os.IsExist(err) {
			return "", err
		}
	} else if fi.IsDir() == false {
		return "", fmt.Errorf("Unable to create as a same name file is already existed")
	}

	if len(m.Filename) == 0 {
		m.Filename = "Dockerfile"
	}
	dockerfileDir := filepath.Join(m.SavingDir, m.Filename)

	err = ioutil.WriteFile(dockerfileDir, buf.Bytes(), 0644)
	if err != nil {
		return "", err
	}
	return dockerfileDir, nil
}
