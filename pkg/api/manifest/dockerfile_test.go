/*
  [vagrant@kubedev-172-17-4-59 go-to-docker]$ GOPATH=/Users/fanhongling/Downloads/workspace/:/Users/fanhongling/go go test -test.v -run GofileserverBased ./pkg/api/manifest
*/

package manifest

import (
	"os"
	"testing"
)

func TestGofileserverBased(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log("wd=", wd)
	mf := NewDockerfileManifest("", "")
	dfp, err := mf.WriteFileUsingGofileserverbasedTemplate("pkg", wd, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dfp)
}
