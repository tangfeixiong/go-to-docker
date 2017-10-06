package elasticsearch

import (
	"context"
	// "flag"
	"fmt"
	"os"
	"sync"
	"time"

	info "github.com/google/cadvisor/info/v1"
	// storage "github.com/google/cadvisor/storage"

	// "gopkg.in/olivere/elastic.v2"
	"gopkg.in/olivere/elastic.v5"

	"github.com/tangfeixiong/go-to-docker/metering/pkg/storage"
)

//func init() {
//	storage.RegisterStorageDriver("elasticsearch", new)
//}

func Register(elastichost, indexname, typename *string, enablesniffer *bool) {
	argElasticHost = elastichost
	argIndexName = indexname
	argTypeName = typename
	argEnableSniffer = enablesniffer
	storage.RegisterStorageDriver("elasticsearch", new)
}

type elasticStorage struct {
	client      *elastic.Client
	machineName string
	indexName   string
	typeName    string
	lock        sync.Mutex
}

type detailSpec struct {
	Timestamp      int64                `json:"timestamp"`
	MachineName    string               `json:"machine_name,omitempty"`
	ContainerName  string               `json:"container_Name,omitempty"`
	ContainerStats *info.ContainerStats `json:"container_stats,omitempty"`
}

//var (
//	argElasticHost   = flag.String("storage_driver_es_host", "http://localhost:9200", "ElasticSearch host:port")
//	argIndexName     = flag.String("storage_driver_es_index", "cadvisor", "ElasticSearch index name")
//	argTypeName      = flag.String("storage_driver_es_type", "stats", "ElasticSearch type name")
//	argEnableSniffer = flag.Bool("storage_driver_es_enable_sniffer", false, "ElasticSearch uses a sniffing process to find all nodes of your cluster by default, automatically")
//)

var (
	argElasticHost   *string
	argIndexName     *string
	argTypeName      *string
	argEnableSniffer *bool
)

func new() (storage.StorageDriver, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	return newStorage(
		hostname,
		*argIndexName,
		*argTypeName,
		*argElasticHost,
		*argEnableSniffer,
	)
}

func (self *elasticStorage) containerStatsAndDefaultValues(
	ref info.ContainerReference, stats *info.ContainerStats) *detailSpec {
	timestamp := stats.Timestamp.UnixNano() / 1E3
	var containerName string
	if len(ref.Aliases) > 0 {
		containerName = ref.Aliases[0]
	} else {
		containerName = ref.Name
	}
	detail := &detailSpec{
		Timestamp:      timestamp,
		MachineName:    self.machineName,
		ContainerName:  containerName,
		ContainerStats: stats,
	}
	return detail
}

func (self *elasticStorage) AddStats(ref info.ContainerReference, stats *info.ContainerStats) error {
	if stats == nil {
		return nil
	}
	func() {
		// AddStats will be invoked simultaneously from multiple threads and only one of them will perform a write.
		self.lock.Lock()
		defer self.lock.Unlock()
		// Add some default params based on ContainerStats
		detail := self.containerStatsAndDefaultValues(ref, stats)
		// Index a cadvisor (using JSON serialization)
		//		_, err := self.client.Index().
		//			Index(self.indexName).
		//			Type(self.typeName).
		//			BodyJson(detail).
		//			Do()
		_, err := self.client.Index().
			Index(self.indexName).
			Type(self.typeName).
			BodyJson(detail).
			Do(context.Background())
		if err != nil {
			// Handle error
			fmt.Printf("failed to write stats to ElasticSearch - %s", err)
			return
		}
	}()
	return nil
}

func (self *elasticStorage) Close() error {
	self.client = nil
	return nil
}

// machineName: A unique identifier to identify the host that current cAdvisor
// instance is running on.
// ElasticHost: The host which runs ElasticSearch.
func newStorage(
	machineName,
	indexName,
	typeName,
	elasticHost string,
	enableSniffer bool,
) (storage.StorageDriver, error) {
	// Obtain a client and connect to the default Elasticsearch installation
	// on 127.0.0.1:9200. Of course you can configure your client to connect
	// to other hosts and configure it in various other ways.
	//	client, err := elastic.NewClient(
	//		elastic.SetHealthcheck(true),
	//		elastic.SetSniff(enableSniffer),
	//		elastic.SetHealthcheckInterval(30*time.Second),
	//		elastic.SetURL(elasticHost),
	//	)
	client, err := elastic.NewSimpleClient(
		elastic.SetURL(elasticHost),
		elastic.SetBasicAuth("elastic", "changeme"),
	)
	if err != nil {
		// Handle error
		return nil, fmt.Errorf("failed to create the elasticsearch client - %s", err)
	}

	for begin := time.Now(); time.Since(begin) < time.Second*60; time.Sleep(time.Second * 15) {
		// Ping the Elasticsearch server to get e.g. the version number
		//info, code, err := client.Ping().URL(elasticHost).Do()
		info, code, err := client.Ping(elasticHost).Timeout("30s").Do(context.Background())
		if err != nil {
			if time.Since(begin) < time.Second*40 {
				continue
			}
			// Handle error
			return nil, fmt.Errorf("failed to ping the elasticsearch - %s", err)

		}
		fmt.Printf("Elasticsearch returned with code %d and version %s", code, info.Version.Number)
		break
	}

	ret := &elasticStorage{
		client:      client,
		machineName: machineName,
		indexName:   indexName,
		typeName:    typeName,
	}
	return ret, nil
}
