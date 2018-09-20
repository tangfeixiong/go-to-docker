/*
  [vagrant@kubedev-172-17-4-59 go-to-docker]$ GOPATH=/Users/fanhongling/Downloads/workspace/:/Users/fanhongling/go go test -test.v -run BuildMetadata ./pkg/api/
*/
package api

import (
	"encoding/json"
	"testing"
)

func TestBuildMetadata(t *testing.T) {
	source := &GitRepoConfig{
		URL:         "https://github.com/tangfeixiong/nta",
		AllTags:     true,
		ForceRemote: true,
	}
	project := &ProjectMetadata{
		Name: "nta",
	}
	dockerImgBldCfg := BasedGofileserverDockerBuildConfig("tangfeixiong/basegofs", "docs", true)

	metadata := DockerImageBuildConfigWithGitRepo(project, dockerImgBldCfg, source)

	bs, err := json.Marshal(metadata)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%q", string(bs))
}
