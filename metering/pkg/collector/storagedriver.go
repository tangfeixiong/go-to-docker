package collector

import (
	// "flag"
	// "fmt"
	// "strings"
	"time"

	"github.com/golang/glog"
	// "github.com/google/cadvisor/cache/memory"
	// "github.com/google/cadvisor/storage"
	//	_ "github.com/google/cadvisor/storage/bigquery"
	//	_ "github.com/google/cadvisor/storage/elasticsearch"
	//	_ "github.com/google/cadvisor/storage/influxdb"
	//	_ "github.com/google/cadvisor/storage/kafka"
	//	_ "github.com/google/cadvisor/storage/redis"
	//	_ "github.com/google/cadvisor/storage/statsd"
	//	_ "github.com/google/cadvisor/storage/stdout"

	"github.com/tangfeixiong/go-to-docker/metering/pkg/cache/memory"
	"github.com/tangfeixiong/go-to-docker/metering/pkg/storage"
)

//var (
//	storageDriver   = flag.String("storage_driver", "", fmt.Sprintf("Storage `driver` to use. Data is always cached shortly in memory, this controls where data is pushed besides the local cache. Empty means none. Options are: <empty>, %s", strings.Join(storage.ListDrivers(), ", ")))
//	storageDuration = flag.Duration("storage_duration", 2*time.Minute, "How long to keep data stored (Default: 2min).")
//)

// NewMemoryStorage creates a memory storage with an optional backend storage option.
func NewMemoryStorage(storageDriver *string, storageDuration *time.Duration) (*memory.InMemoryCache, error) {
	backendStorage, err := storage.New(*storageDriver)
	if err != nil {
		return nil, err
	}
	if *storageDriver != "" {
		glog.Infof("Using backend storage type %q", *storageDriver)
	}
	glog.Infof("Caching stats in memory for %v", *storageDuration)
	return memory.New(*storageDuration, backendStorage), nil
}
