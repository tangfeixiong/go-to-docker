/*
  Refer to:
  - https://github.com/src-d/go-git
  - https://github.com/src-d/go-git/blob/master/_examples/clone/main.go"
*/
package gitscm

import (
	"fmt"

	"github.com/golang/glog"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type gogitClient struct {
	url      *string
	dir      *string
	cloneOpt *git.CloneOptions
}

// Basic example of how to clone a repository using clone options.
func (client *gogitClient) Clone(url, directory string) error {
	//	CheckArgs("<url>", "<directory>")
	if len(url) == 0 {
		return ErrRepoURLNotSpecified
	}
	if len(directory) == 0 {
		return ErrRepoDirectoryNotSpecified
	}
	//	url := os.Args[1]
	//	directory := os.Args[2]

	// Clone the given repository to the given directory
	glog.Infof("git clone %s %s --recursive", url, directory)

	opt := &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	}
	if client.cloneOpt != nil {
		opt = client.cloneOpt
		opt.URL = url
	}
	r, err := git.PlainClone(directory, false, opt)

	if err != nil {
		glog.Errorf("failed to clone %s", err)
		return err
	}

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	if err != nil {
		glog.Errorf("failed to find ref HEAD")
		return err
	}

	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		glog.Errorf("failed to get commit hash")
		return err
	}

	fmt.Println(commit)
	return nil
}

func (client *gogitClient) cloneOptURL(url string) error {
	if len(url) == 0 {
		return ErrRepoURLNotSpecified
	}
	if client.cloneOpt == nil {
		client.cloneOpt = new(git.CloneOptions)
	}
	client.cloneOpt.URL = url
	return nil
}

func (client *gogitClient) cloneOptRef(ref string) error {
	if client.cloneOpt == nil {
		client.cloneOpt = new(git.CloneOptions)
	}
	client.cloneOpt.ReferenceName = plumbing.ReferenceName(ref)
	return nil
}

func (client *gogitClient) cloneOptSingle(single bool) error {
	if client.cloneOpt == nil {
		client.cloneOpt = new(git.CloneOptions)
	}
	client.cloneOpt.SingleBranch = single
	return nil
}

func (client *gogitClient) cloneOptDepth(depth int) error {
	if client.cloneOpt == nil {
		client.cloneOpt = new(git.CloneOptions)
	}
	if depth > 0 {
		client.cloneOpt.Depth = depth
	}
	return nil
}

func (client *gogitClient) CloneOptRecursiveDepth(depth int) error {
	if client.cloneOpt == nil {
		client.cloneOpt = new(git.CloneOptions)
	}
	client.cloneOpt.RecurseSubmodules = git.DefaultSubmoduleRecursionDepth
	if depth == 0 {
		client.cloneOpt.RecurseSubmodules = git.NoRecurseSubmodules
	} else if 0 < depth && depth < 15 {
		client.cloneOpt.RecurseSubmodules = git.SubmoduleRescursivity(uint(depth))
	}
	return nil
}
