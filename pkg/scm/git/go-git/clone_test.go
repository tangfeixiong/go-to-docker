/*
  Refer to:
  - https://github.com/src-d/go-git/blob/master/_examples/clone/main.go"
  - [vagrant@kubedev-172-17-4-59 go-to-docker]$ GOPATH=/Users/fanhongling/Downloads/workspace/:/Users/fanhongling/go go test -test.v -run Clone ./pkg/scm/git/gogit
*/
package /* main */ gogit

import (
	"fmt"
	// "os"
	"testing"

	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
)

// Basic example of how to clone a repository using clone options.
func /* main() */ TestClone(t *testing.T) {
	//	CheckArgs("<url>", "<directory>")
	//	url := os.Args[1]
	//	directory := os.Args[2]
	url := "https://github.com/tangfeixiong/nta"
	directory := "/tmp/test/nta"

	// Clone the given repository to the given directory
	Info("git clone %s %s --recursive", url, directory)

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	CheckIfError(err)

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	CheckIfError(err)
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(commit)
}
