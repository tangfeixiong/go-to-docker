/*
  [vagrant@kubedev-172-17-4-59 go-to-docker]$ GOPATH=/Users/fanhongling/Downloads/workspace/:/Users/fanhongling/go go test -test.v -run GitClone ./pkg/scm/git/
*/
package gitscm

import (
	"testing"
)

func TestGitClone(t *testing.T) {
	git, err := New(CloneOptsSingleBranchOnly(true))
	if err != nil {
		t.Fatal(err)
	}
	if err := git.Clone("https://github.com/tangfeixiong/nta", "/tmp/gittestrepos/nta"); err != nil {
		t.Fatal(err)
	}
}
