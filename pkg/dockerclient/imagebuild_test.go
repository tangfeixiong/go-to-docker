/*
[vagrant@kubedev-172-17-4-59 go-to-docker]$ GOPATH=/Users/fanhongling/Downloads/workspace/:/Users/fanhongling/go go test -test.v -run ImageBuild_api ./pkg/dockerclient/
*/
package dockerclient

import (
	"archive/tar"
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"text/template"

	dockertypes "github.com/docker/docker/api/types"

	"github.com/tangfeixiong/go-to-docker/pkg/api/artifact"
)

func TestImageBuild_api_dockerfile(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasSuffix(wd, "go-to-docker") {
		ndx := strings.LastIndex(wd, "go-to-docker")
		if ndx < 0 {
			t.Skip("Unable to change working dir to go-to-docker")
		}
		wd = wd[:ndx+len("go-to-docker")]
		//		if err := os.Chdir(wd); err != nil {
		//			t.Skipf("Failed to change working dir: %v", err)
		//		}
	}

	apiClient, err := getDockerClient("")
	if err != nil {
		t.Skip(err)
	}

	dockerfile := "Dockerfile.gofileserver-based"
	opts := dockertypes.ImageBuildOptions{
		Tags:           []string{"tangfeixiong/basegofs"},
		NoCache:        true,
		SuppressOutput: false,
		Remove:         true,
		ForceRemove:    true,
		PullParent:     true,
	}
	opts.Dockerfile = dockerfile
	// opts.Memory =
	// opts.MemorySwap =
	// opts.CgroupParent =

	dockerfileData, srcList := dockerfileFromTemplate(wd, t)
	dir, err := ioutil.TempDir(os.TempDir(), "dockerbuildupload")
	if err != nil {
		t.Skip(err)
	}
	dockerfilePath := filepath.Join(dir, dockerfile)
	ioutil.WriteFile(dockerfilePath, dockerfileData, 0644)
	tarFile, err := ioutil.TempFile(dir, "tar")
	if err != nil {
		t.Skip(err)
	}
	defer tarFile.Close()
	tarWriter := tar.NewWriter(tarFile)
	defer tarWriter.Close()

	var sepfn shouldExcludePathFn = func(path string) bool {
		for _, ele := range srcList {
			if strings.HasPrefix(path, filepath.Join(wd, ele)) {
				fmt.Println(path)
				return false
			}
		}
		return true
	}
	if err := CreateTarStreamToTarWriter(wd, false, tarWriter, nil, sepfn); err != nil {
		t.Fatal(err)
	}
	if err := tarWriter.Flush(); err != nil {
		t.Fatal(err)
	}
	if err := CreateTarStreamToTarWriter(dockerfilePath, true, tarWriter, nil, func(path string) bool {
		return false
	}); err != nil {
		t.Fatal(err)
	}
	if err := tarWriter.Flush(); err != nil {
		t.Fatal(err)
	}
	tarPath := tarFile.Name()
	if err := tarFile.Close(); err != nil {
		t.Fatal(err)
	}
	tarFile, err = os.Open(tarPath)
	if err != nil {
		t.Fatal(err)
	}
	defer tarFile.Close()
	fread := bufio.NewReader(tarFile)

	resp, err := apiClient.ImageBuild(context.Background(), fread, opts)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	_, err = io.Copy(os.Stdout, resp.Body)
}

func dockerfileFromTemplate(dir string, t *testing.T) ([]byte, []string) {
	asset, err := artifact.Asset("template/Dockerfile.gofileserver-based.go-tpl")
	if err != nil {
		t.Skip(err)
	}

	//	fs, err := ioutil.ReadDir(".")
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Skipf("Unable to find files to test in working dir: %v", err)
	}
	src := []string{}
	buf := bytes.Buffer{}
	re := regexp.MustCompile(`\.md|\.txt|\.ya?ml|\.json|\.toml|Dockerfile.*`)
	for _, f := range fs {
		if f.Mode().IsRegular() && re.Match([]byte(f.Name())) {
			src = append(src, f.Name())
			buf.WriteString(f.Name())
			buf.WriteString(" ")
			continue
		}
		if f.Mode().IsDir() && (f.Name() == "cmd" || f.Name() == "script" || f.Name() == "template") {
			src = append(src, f.Name())
			buf.WriteString(f.Name())
			buf.WriteString("/ ")
		}
	}
	if buf.Len() == 0 {
		t.Skip("please change working dir to repo home, e.g. go-to-docker")
	}

	dockerfile := &bytes.Buffer{}
	tte := template.Must(template.New("basedgofs").Parse(string(asset)))
	err = tte.Execute(dockerfile, struct{ SrcStr string }{
		SrcStr: "/",
	})
	if err != nil {
		t.Skip(err)
	}

	return dockerfile.Bytes(), src
}
