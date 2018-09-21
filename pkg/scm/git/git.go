package gitscm

// import "github.com/tangfeixiong/go-to-docker/pkg/scm/git

import (
	"errors"
	"io"

	"github.com/tangfeixiong/go-to-docker/pkg/util/logging"
)

type Interface interface {
	Clone(url, directory string) error
}

type githubClient struct {
}

type gitlabClient struct {
}

type gogsClient struct {
}

type Client struct {
	githubclient *githubClient
	gitlabclient *gitlabClient
	gogsclient   *gogsClient

	gogitclient      *gogitClient
	cloneStateReader io.Reader
}

type GitClientOptFn func(*Client) error

func New(optFns ...GitClientOptFn) (*Client, error) {
	client := &Client{
		gogitclient: &gogitClient{},
	}
	for _, optFn := range optFns {
		if err := optFn(client); err != nil {
			return nil, err
		}
	}
	return client, nil
}

var (
	ErrNilPointer error = errors.New("client is nil")

	ErrRepoURLNotSpecified       error = errors.New("URL not specified")
	ErrRepoDirectoryNotSpecified error = errors.New("Local directory not specified")
)

// Branch, Tag, commit Hash to clone, default is master/HEAD
func CloneOptsReferenceName(refName string) GitClientOptFn {
	return func(client *Client) (err error) {
		if client == nil {
			err = ErrNilPointer
		} else {
			err = client.gogitclient.cloneOptRef(refName)
		}
		return
	}
}

// If true, only clone single branch
func CloneOptsSingleBranchOnly(singleBranch bool) GitClientOptFn {
	return func(client *Client) (err error) {
		if client == nil {
			err = ErrNilPointer
		} else {
			err = client.gogitclient.cloneOptSingle(singleBranch)
		}
		if err != nil {
			logging.New("SingleBranchCloneOption").Errorf("error: %s", err.Error())
		}
		return
	}
}

// Limit commits number
func CloneOptsCommitsLimit(limitDepth int) GitClientOptFn {
	return func(client *Client) (err error) {
		if client == nil {
			err = ErrNilPointer
		} else {
			err = client.gogitclient.cloneOptDepth(limitDepth)
		}
		return
	}
}

func CloneOptRecursionSubmodulesDepth(resursiveDepth int) GitClientOptFn {
	return func(client *Client) (err error) {
		if client == nil {
			err = ErrNilPointer
		} else {
			err = client.gogitclient.CloneOptRecursiveDepth(resursiveDepth)
		}
		return
	}
}

func CloneOptStateReader(reader io.Reader) GitClientOptFn {
	return func(client *Client) (err error) {
		if client == nil {
			err = ErrNilPointer
		} else {
			if reader != nil {
				client.cloneStateReader = reader
			}
		}
		return
	}
}

func (client *Client) Clone(url, directory string) error {
	//	var writer io.Writer = nil
	//	if client.cloneStateReader != nil {
	//		writer = &bytes.NewBuffer(make([]byte, 0, 1024))
	//		io.Copy(writer, client.cloneStateReader)
	//	}
	return client.gogitclient.Clone(url, directory)
}
