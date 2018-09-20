/*
  Inspired by
  - https://github.com/Stoakes/grpc-gateway-example
*/

package server2

import (
	"context"
	"fmt"
	"io"
	"mime"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	assetfs "github.com/philips/go-bindata-assetfs"
	"github.com/soheilhy/cmux"

	"google.golang.org/grpc"

	"github.com/tangfeixiong/go-to-docker/pb"
	"github.com/tangfeixiong/go-to-docker/pkg/dockerclient"
	"github.com/tangfeixiong/go-to-docker/pkg/ui/data/swagger"
)

type MicroServer struct {
	host       string
	lstn       net.Listener
	httpServer *http.Server
	grpcServer *grpc.Server

	config       *Config
	DockerClient *dockerclient.DockerClient
}

func Start(opt *Option) {
	cfg := NewConfig()

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	s := New(fmt.Sprint(":%d", cfg.Port), conn)
	s.config = cfg

	s.DockerClient = dockerclient.NewOrDie()

	if err = s.Start(); err != nil {
		panic(err)
	}
}

func New(addr string, l net.Listener) *MicroServer {
	return &MicroServer{
		host: addr,
		lstn: l,
	}
}

func (ms *MicroServer) Start() error {
	var err error

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tcpMux := cmux.New(ms.lstn)
	grpcL := tcpMux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
	httpL := tcpMux.Match(cmux.HTTP1Fast())

	ms.grpcServer, err = newGRPC(ctx, ms)
	if err != nil {
		panic(err)
	}
	ms.httpServer, err = newHTTP(ctx, ms.host, ms.config.Port)
	if err != nil {
		panic(err)
	}

	go func() {
		if err := ms.grpcServer.Serve(grpcL); err != nil {
			panic(err)
		}
	}()
	go func() {
		if err := ms.httpServer.Serve(httpL); err != nil {
			panic(err)
		}
	}()

	return tcpMux.Serve()
}

func newGRPC(ctx context.Context, ms *MicroServer) (*grpc.Server, error) {
	grpcServer := grpc.NewServer()
	pb.RegisterSimpleServiceServer(grpcServer, ms)
	return grpcServer, nil
}

func newHTTP(ctx context.Context, host string, port int) (*http.Server, error) {
	router := http.NewServeMux()
	router.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, strings.NewReader(pb.Swagger))
	})
	serveSwagger(router)
	serveReSTful(router)

	// initialize grpc-gateway
	gw, err := newGateway(ctx, fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return nil, err
	}
	router.Handle("/", gw)

	return &http.Server{
		Addr:    host,
		Handler: router,
		// ReadTimeout:  5 * time.Second,
		// WriteTimeout: 10 * time.Second,
		// IdleTimeout:  120 * time.Second,
	}, nil
}

func newGateway(ctx context.Context, host string) (http.Handler, error) {
	opts := []grpc.DialOption{
		grpc.WithTimeout(10 * time.Second),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		fmt.Printf("Failed to dial: %v\n", err)
		return nil, err
	}

	// gwMux := runtime.NewServerMux()
	// Change JSON serializer to include empty fields with default values
	gwMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			OrigName:     true,
			EmitDefaults: true}),
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)

	err = pb.RegisterSimpleServiceHandler(ctx, gwMux, conn)
	if err != nil {
		fmt.Printf("Unable to instantiate gRPC Gateway: %v", err)
		return nil, err
	}
	return gwMux, nil
}

func serveSwagger(mux *http.ServeMux) {
	mime.AddExtensionType(".svg", "image/svg+xml")

	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
