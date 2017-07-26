package server

import (
	"fmt"
	"io"
	"mime"
	"net"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/philips/go-bindata-assetfs"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/tangfeixiong/go-to-docker/pb"
	"github.com/tangfeixiong/go-to-docker/pkg/ui/data/swagger"
)

type myService struct {
}

func newServer() *myService {
	return &myService{}
}

func (m *myService) RunContainer(ctx context.Context, in *pb.DockerRunData) (*pb.DockerRunData, error) {
	if glog.V(2) {
		glog.V(2).Infoln("[gRPC] run container!")
	} else {
		glog.Infoln("[gRPC] run container.")
	}
	return m.runContainer(in)
}

func (m *myService) RemoveContainer(ctx context.Context, in *pb.DockerRunData) (*pb.DockerRunData, error) {
	return m.removeContainer(in)
}

func (m *myService) GetImageFromRegistry(ctx context.Context, in *pb.ImageRegistryData) (*pb.ImageRegistryData, error) {
	return nil, fmt.Errorf("not ready")
}

func (m *myService) PullImage(ctx context.Context, in *pb.ImageRegistryData) (*pb.ImageRegistryData, error) {
	return nil, fmt.Errorf("not ready")
}

func (m *myService) rmiArchived(ctx context.Context, in *pb.ImageArchiveData) (*pb.ImageArchiveData, error) {
	return nil, fmt.Errorf("not ready")
}

func (m *myService) InspectImage(ctx context.Context, in *pb.ImageArchiveData) (*pb.ImageArchiveData, error) {
	return nil, fmt.Errorf("not ready")
}

func (m *myService) DefineImageCatalog(ctx context.Context, in *pb.ImageCatalogData) (*pb.ImageCatalogData, error) {
	return nil, fmt.Errorf("not ready")
}

func (m *myService) UndefineImageCatalog(ctx context.Context, in *pb.ImageCatalogData) (*pb.ImageCatalogData, error) {
	return nil, fmt.Errorf("not ready")
}

func (m *myService) Echo(c context.Context, s *pb.EchoMessage) (*pb.EchoMessage, error) {
	fmt.Printf("rpc request Echo(%q)\n", s.Value)
	return s, nil
}

func StartGRPC(ch chan string) {
	host := ":10051"
	if v, ok := os.LookupEnv("GRPC_PORT_GTD"); ok && 0 != len(v) {
		if strings.Contains(v, ":") {
			host = v
		} else {
			host = "localhost:" + v
		}
	}
	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, newServer())

	l, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}

	ch <- host
	fmt.Println("Start gRPC on host", l.Addr())
	if err := s.Serve(l); nil != err {
		panic(err)
	}
}

func StartGateway(gRPCHost string) {
	host := ":10052"
	if v, ok := os.LookupEnv("HTTP_PORT_GTD"); ok && 0 != len(v) {
		if strings.Contains(v, ":") {
			host = v
		} else {
			host = ":" + v
		}
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	// mux.HandleFunc("/swagger/", serveSwagger2)
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, strings.NewReader(pb.Swagger))
	})

	dopts := []grpc.DialOption{grpc.WithInsecure()}

	fmt.Println("Start gRPC Gateway into host", gRPCHost)
	gwmux := runtime.NewServeMux()
	if err := pb.RegisterEchoServiceHandlerFromEndpoint(ctx, gwmux, gRPCHost, dopts); err != nil {
		fmt.Println("Failed to run HTTP server. ", err)
		return
	}

	mux.Handle("/", gwmux)
	serveSwagger(mux)
	//	fmt.Printf("Start HTTP")
	//	if err := http.ListenAndServe(host, allowCORS(mux)); nil != err {
	//		fmt.Fprintf(os.Stderr, "Server died: %s\n", err)
	//	}

	lstn, err := net.Listen("tcp", host)
	if nil != err {
		panic(err)
	}

	fmt.Printf("http on host: %s\n", lstn.Addr())
	srv := &http.Server{
		Handler: allowCORS(mux),
	}

	if err := srv.Serve(lstn); nil != err {
		fmt.Fprintln(os.Stderr, "Server died.", err.Error())
	}
}

func serveSwagger(mux *http.ServeMux) {
	mime.AddExtensionType(".svg", "image/svg+xml")

	// Expose files in third_party/swagger-ui/ on <host>/swagger-ui
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

var swaggerDir string = "examples/examplepb"

func serveSwagger2(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
		glog.Errorf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	glog.Infof("Serving %s", r.URL.Path)
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(swaggerDir, p)
	http.ServeFile(w, r, p)
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	// glog.Infof("preflight request for %s", r.URL.Path)
	return
}
