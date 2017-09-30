package exporter

import (
	//	"bufio"
	//	"bytes"
	//	b64 "encoding/base64"
	//	"errors"
	//	"fmt"
	//	"io/ioutil"
	//	"math"
	//	"os"
	//	"path/filepath"
	//	"regexp"
	//	"strings"
	"sync"
	"time"

	//	"github.com/golang/glog"

	//	"github.com/tangfeixiong/go-to-docker/metering/pb"
)

type ExporterManager struct {
	MeteringNameURLs map[string]string
	Dispatcher       map[string][]string
	name             string
	command          []string
	args             []string
	env              []string
	conf             map[string]string
	workdir          string
	periodic         int32
	duration         int32
	destinationPath  string
	RootPath         string
	ticker           *time.Ticker
	timestamp        time.Time
	result           []byte
	mutex            sync.Mutex
}

func (em *ExporterManager) Start(ch <-chan bool) {

}
