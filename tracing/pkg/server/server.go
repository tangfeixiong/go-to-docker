package server

type myServer struct {
	grpcHost               string
	httpHost               string
	packgesHome            string
	redisSentinelAddresses string
	redisAddresses         []string
	redisDB                int
	etcdAddresses          string
	mysqlAddress           string
	gnatsdAddresses        string
	kafkaAddresses         string
	zookeeperAddresses     string
	rabbitAddress          string
	priorityCMDB           []string
	priorityMQ             []string
	subSubject             string
	subQueue               string
	pubSubject             string
	cmCache                string
	unsubCh                chan string
	dispatchersignal       chan bool
}
