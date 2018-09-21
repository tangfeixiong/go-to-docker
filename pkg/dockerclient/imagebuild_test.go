/*
[vagrant@kubedev-172-17-4-59 go-to-docker]$ GOPATH=/Users/fanhongling/Downloads/workspace/:/Users/fanhongling/go go test -test.v -run ImageBuild ./pkg/dockerclient/
*/
package dockerclient

import (
	"archive/tar"
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
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

	"github.com/tangfeixiong/go-to-docker/pb"
	mobypb "github.com/tangfeixiong/go-to-docker/pb/moby"
	"github.com/tangfeixiong/go-to-docker/pkg/api/artifact"
)

func TestImageBuildDockerfileAPI(t *testing.T) {
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
		Tags:           []string{"tangfeixiong/basedongofileserver"},
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
	err = tte.Execute(dockerfile, struct{ SrcStr, DstDir string }{
		SrcStr: "/",
		DstDir: "",
	})
	if err != nil {
		t.Skip(err)
	}

	return dockerfile.Bytes(), src
}

var (
	dockerfile []byte = []byte(`
FROM alpine
RUN apk add --update netcat-openbsd && rm -rf /var/cache/apk/*
RUN echo -e "#!/bin/sh\n\ 
set -e\n\
while true; do echo -e \"HTTP/1.1 200 OK\n\n \$(date) Hello world\" | nc -l 80; done" > /entrypoint.sh \
    && chmod +x /entrypoint.sh
# RUN touch /entrypoint.sh && chmod +x /entrypoint.sh && echo -e "#!/bin/sh\nset -e\nwhile true; do nc -l 80 < index.html; done" > /entrypoint.sh
RUN echo -e "\n\
<html>\
        <head>\
                <title>Hello Page</title>\
        </head>\
        <body>\
                <h1>Hello</h1>\
                <h2>Container</h2>\
                <p>Powered by nc</p>\
        </body>\
</html>\
" > /index.html

ENTRYPOINT ["/entrypoint.sh"]
EXPOSE 80
`)
	dockerfileB64 string = "CkZST00gYWxwaW5lClJVTiBhcGsgYWRkIC0tdXBkYXRlIG5ldGNhdC1vcGVuYnNkICYmIHJtIC1yZiAvdmFyL2NhY2hlL2Fway8qClJVTiBlY2hvIC1lICIjIS9iaW4vc2hcblwgCnNldCAtZVxuXAp3aGlsZSB0cnVlOyBkbyBuYyAtbCA4MCA8IGluZGV4Lmh0bWw7IGRvbmUiID4gL2VudHJ5cG9pbnQuc2ggXAogICAgJiYgY2htb2QgK3ggL2VudHJ5cG9pbnQuc2gKUlVOIGVjaG8gLWUgIlxuXAo8aHRtbD5cCiAgICAgICAgPGhlYWQ+XAogICAgICAgICAgICAgICAgPHRpdGxlPkhlbGxvIFBhZ2U8L3RpdGxlPlwKICAgICAgICA8L2hlYWQ+XAogICAgICAgIDxib2R5PlwKICAgICAgICAgICAgICAgIDxoMT5IZWxsbzwvaDE+XAogICAgICAgICAgICAgICAgPGgyPkNvbnRhaW5lcjwvaDI+XAogICAgICAgICAgICAgICAgPHA+UG93ZXJlZCBieSBuYzwvcD5cCiAgICAgICAgPC9ib2R5PlwKPC9odG1sPlwKIiA+IC9pbmRleC5odG1sCgpFTlRSWVBPSU5UIFsiL2VudHJ5cG9pbnQuc2giXQpFWFBPU0UgODAK"
)

func TestImageBuildClientDockerfile(t *testing.T) {
	client := NewOrDie()

	dockerfileB64 = base64.StdEncoding.EncodeToString(dockerfile)
	fmt.Println(dockerfileB64)

	req := &pb.DockerImageBuildReqResp{
		BuildContext: dockerfile,
		ImageBuildOptions: &mobypb.ImageBuildOptions{
			Tags:           []string{"tangfeixiong/hello-world:netcat-http"},
			NoCache:        true,
			SuppressOutput: true,
			Remove:         true,
			ForceRemove:    true,
			PullParent:     true,
		},
	}
	resp, err := client.BuildImage(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(resp.ImageBuildResponse.Body))
}

var (
	gitMetadata []byte = []byte(`# ** metadata:git **
#

https://github.com/tangfeixiong/nta#:docs`)
	gitMetadataB64 string = "IyAqKiBtZXRhZGF0YTpnaXQgKioKIwoKaHR0cHM6Ly9naXRodWIuY29tL3RhbmdmZWl4aW9uZy9udGEjOmRvY3M="
)

func TestImageBuildClientGit(t *testing.T) {
	client := NewOrDie()

	gitMetadataB64 = base64.StdEncoding.EncodeToString(gitMetadata)
	fmt.Println(gitMetadataB64)

	req := &pb.DockerImageBuildReqResp{
		BuildContext: gitMetadata,
		ImageBuildOptions: &mobypb.ImageBuildOptions{
			Tags:           []string{"tangfeixiong/basedongofileserver"},
			NoCache:        true,
			SuppressOutput: true,
			Remove:         true,
			ForceRemove:    true,
			PullParent:     true,
		},
	}
	resp, err := client.BuildImage(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(resp.ImageBuildResponse.Body))
}
