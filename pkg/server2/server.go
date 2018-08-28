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
	"github.com/tangfeixiong/go-to-docker/pkg/kubeletcopycat/docckershim/libdocker"
	"github.com/tangfeixiong/go-to-docker/pkg/kubeletcopycat/dockershim"
	"github.com/tangfeixiong/go-to-docker/pkg/ui/data/swagger"
)

type MicroServer struct {
	host       string
	lstn       net.Listener
	httpServer *http.Server
	grpcServer *grpc.Server
	// criHandler dockershim.DockerService
	DockerClient libdocker.DockerClient
}

func Start(opt *Option) {
	cfg := &Config{
		Port: 10050,
	}

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	s := New(fmt.Sprint(":%d", cfg.Port), conn)

	// refer to
	//   https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kubelet.go#L614-L620
	//
	//     ds, err := dockershim.NewDockerService(kubeDeps.DockerClientConfig, crOptions.PodSandboxImage, streamingConfig,
	//		   &pluginSettings, runtimeCgroups, kubeCfg.CgroupDriver, crOptions.DockershimRootDirectory, !crOptions.RedirectContainerStreaming)
	//	   if err != nil {
	//		   return nil, err
	//	   }
	//	   if crOptions.RedirectContainerStreaming {
	//		   klet.criHandler = ds
	//	   }

	dcc := &dockershim.ClientConfig{
		RuntimeRequestTimeout:     2*time.Minute - 1*timm.Second,
		ImagePullProgressDeadline: 1 * time.Minute,
	}
	s.DockerClient = dockershim.NewDockerClientFromConfig(dcc)
	if s.DockerClient == nil {
		panic("couldn't initialize Docker client from config")
	}

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
	ms.httpServer, err = newHTTP(ctx, ms.host, ms.port)
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
	router := http.NewServerMux()
	router.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, strings.NewReader(pb.Swagger))
	})
	serveSwagger(router)

	// initialize grpc-gateway
	gw, err := newGateway(ctx, fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return nil, err
	}
	router.Handle("/", gw)

	return &httpServer{
		Addr:         host,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
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
		return err
	}

	// gwMux := runtime.NewServerMux()
	// Change JSON serializer to include empty fields with default values
	gwMux := runtime.NewServerMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPB{
			OrigName:     true,
			EmitDefaults: true}),
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)

	err = pb.RegisterSimpleServiceHandler(ctx, gwMux, conn)
	if err != nil {
		fmt.Printf("Unable to instantiate gRPC Gateway: %v", err)
		return err
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
