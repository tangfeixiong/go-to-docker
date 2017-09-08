package counselor

import (
	"runtime"
	"testing"
	"time"
)

func Test_checkout(t *testing.T) {
	cm := WebXCheckerConfigMgmt{
		RootPath:        "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/python/checkalive",
		destinationPath: "web1check.py",
	}
	cm.checkout()
}

func Test_createticker(t *testing.T) {
	cm := WebXCheckerConfigMgmt{
		RootPath:        "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/python/checkalive",
		destinationPath: "web1check.py",
		periodic:        3,
	}
	cm.CreateTicker()
	time.Sleep(time.Second * 10)
	runtime.Goexit()
}

func Test_updateticker(t *testing.T) {
	cm := WebXCheckerConfigMgmt{
		RootPath:        "/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-docker/examples/python/checkalive",
		destinationPath: "web1check.py",
		periodic:        3,
	}
	cm.CreateTicker()
	time.Sleep(time.Second * 5)

	cm.periodic = 4
	cm.UpdateTicker()
	time.Sleep(time.Second * 5)
}
