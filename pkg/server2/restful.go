package server2

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"

	// "github.com/tangfeixiong/go-to-docker/pkg/dockerclient"
	"github.com/tangfeixiong/go-to-docker/pkg/ui/data/openapi"
)

func serveReSTful(mux *http.ServeMux) {
	wsc := restful.NewContainer()
	wsc.ServeMux = mux
	//	wsc.DoNotRecover(true)
	wsc.RecoverHandler(logStackOnRecover)
	wsc.ServiceErrorHandler(writeServiceError)
	//	wsc.Router(restful.CurlyRouter{})
	//	wsc.EnableContentEncoding(false)
	u := UserResource{map[string]User{}}
	//	restful.DefaultContainer.Add(u.WebService())
	rs := u.WebService()
	// dockerclient.DockerBuildRestfulSerive(rs)
	wsc.Add(rs)
	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(), // you control what services are visible
		APIPath:     "/apidocs.json",
		PostBuildSwaggerObjectHandler: /* enrichSwaggerObject */ func(swo *spec.Swagger) {
			swo.Info = &spec.Info{
				InfoProps: spec.InfoProps{
					Title:       "A docker build/push service",
					Description: "Resource for managing Docker Build/Push",
					Contact: &spec.ContactInfo{
						Name:  "Feixiong Tang",
						Email: "tangfx128@gmail.com",
						URL:   "https://github.com/tangfeixiong",
					},
					License: &spec.License{
						Name: "APACHE",
						URL:  "http://www.apache.org/licenses/",
					},
					Version: "0.2.0",
				},
			}
			swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
				Name:        "users",
				Description: "Managing users"}}}
		},
	}
	//restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))
	wsc.Add(restfulspec.NewOpenAPIService(config))

	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json
	//	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("/Users/emicklei/Projects/swagger-ui/dist"))))
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    openapi.Asset,
		AssetDir: openapi.AssetDir,
		Prefix:   "third_party/OpenAPI",
	})
	mux.Handle("/apidocs/", http.StripPrefix("/apidocs/", fileServer))

	// Optionally, you may need to enable CORS for the UI to work.
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		CookiesAllowed: false,
		Container:/* restful.DefaultContainer */ wsc,
	}
	//	restful.DefaultContainer.Filter(cors.Filter)
	wsc.Filter(cors.Filter)

	//	log.Printf("Get the API using http://localhost:8080/apidocs.json")
	//	log.Printf("Open Swagger UI using http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json")
	logger := log.New(os.Stderr, "[server2 serveReSTful] ", log.LstdFlags|log.Lshortfile)
	logger.Printf("Get the API using http://<host>/apidocs.json")
	logger.Printf("Open OpenAPI UI using http://<host>/apidocs/?url=http://<host>/apidocs.json")
}

// logStackOnRecover is the default RecoverHandleFunction and is called
// when DoNotRecover is false and the recoverHandleFunc is not set for the container.
// Default implementation logs the stacktrace and writes the stacktrace on the response.
// This may be a security issue as it exposes sourcecode information.
func logStackOnRecover(panicReason interface{}, httpWriter http.ResponseWriter) {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("recover from panic situation: - %v\r\n", panicReason))
	for i := 2; ; i += 1 {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		buffer.WriteString(fmt.Sprintf("    %s:%d\r\n", file, line))
	}
	//	log.Print(buffer.String())
	logger := log.New(os.Stderr, "[server2 logStackOnRecover] ", log.LstdFlags|log.Lshortfile)
	logger.Print(buffer.String())
	httpWriter.WriteHeader(http.StatusInternalServerError)
	httpWriter.Write(buffer.Bytes())
}

// writeServiceError is the default ServiceErrorHandleFunction and is called
// when a ServiceError is returned during route selection. Default implementation
// calls resp.WriteErrorString(err.Code, err.Message)
func writeServiceError(err restful.ServiceError, req *restful.Request, resp *restful.Response) {
	resp.WriteErrorString(err.Code, err.Message)
}

//package main

//import (
//	"log"
//	"net/http"

//	"github.com/emicklei/go-restful"
//	restfulspec "github.com/emicklei/go-restful-openapi"
//	"github.com/go-openapi/spec"
//)

// UserResource is the REST layer to the User domain
type UserResource struct {
	// normally one would use DAO (data access object)
	users map[string]User
}

// WebService creates a new service that can handle REST requests for User resources.
func (u UserResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/api/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	tags := []string{"users"}

	ws.Route(ws.GET("/").To(u.findAllUsers).
		// docs
		Doc("get all users").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]User{}).
		Returns(200, "OK", []User{}))

	ws.Route(ws.GET("/{user-id}").To(u.findUser).
		// docs
		Doc("get a user").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(User{}). // on the response
		Returns(200, "OK", User{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{user-id}").To(u.updateUser).
		// docs
		Doc("update a user").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(User{})) // from the request

	ws.Route(ws.PUT("").To(u.createUser).
		// docs
		Doc("create a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(User{})) // from the request

	ws.Route(ws.DELETE("/{user-id}").To(u.removeUser).
		// docs
		Doc("delete a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")))

	return ws
}

// GET http://localhost:8080/users
//
func (u UserResource) findAllUsers(request *restful.Request, response *restful.Response) {
	list := []User{}
	for _, each := range u.users {
		list = append(list, each)
	}
	response.WriteEntity(list)
}

// GET http://localhost:8080/users/1
//
func (u UserResource) findUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	usr := u.users[id]
	if len(usr.ID) == 0 {
		response.WriteErrorString(http.StatusNotFound, "User could not be found.")
	} else {
		response.WriteEntity(usr)
	}
}

// PUT http://localhost:8080/users/1
// <User><Id>1</Id><Name>Melissa Raspberry</Name></User>
//
func (u *UserResource) updateUser(request *restful.Request, response *restful.Response) {
	usr := new(User)
	err := request.ReadEntity(&usr)
	if err == nil {
		u.users[usr.ID] = *usr
		response.WriteEntity(usr)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

// PUT http://localhost:8080/users/1
// <User><Id>1</Id><Name>Melissa</Name></User>
//
func (u *UserResource) createUser(request *restful.Request, response *restful.Response) {
	usr := User{ID: request.PathParameter("user-id")}
	err := request.ReadEntity(&usr)
	if err == nil {
		u.users[usr.ID] = usr
		response.WriteHeaderAndEntity(http.StatusCreated, usr)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

// DELETE http://localhost:8080/users/1
//
func (u *UserResource) removeUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	delete(u.users, id)
}

func main() {
	u := UserResource{map[string]User{}}
	restful.DefaultContainer.Add(u.WebService())

	config := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(), // you control what services are visible
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))

	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json
	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("/Users/emicklei/Projects/swagger-ui/dist"))))

	log.Printf("start listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "UserService",
			Description: "Resource for managing Users",
			Contact: &spec.ContactInfo{
				Name:  "john",
				Email: "john@doe.rp",
				URL:   "http://johndoe.org",
			},
			License: &spec.License{
				Name: "MIT",
				URL:  "http://mit.org",
			},
			Version: "1.0.0",
		},
	}
	swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
		Name:        "users",
		Description: "Managing users"}}}
}

// User is just a sample type
type User struct {
	ID   string `json:"id" description:"identifier of the user"`
	Name string `json:"name" description:"name of the user" default:"john"`
	Age  int    `json:"age" description:"age of the user" default:"21"`
}
