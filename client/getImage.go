
package main

import(
    "fmt"
    "net"
    "net/url"
    "net/http"
    "bufio"
    "io"
    "io/ioutil"
    "encoding/json"    
)

type DockerImage struct {
    Id          string      `json:"Id" yaml:"Id"`
    RepoTags    []string    `json:"RepoTags,omitempty" yaml:"RepoTags,omitempty"`
    Created     int64       `json:"Created,omitempty" yaml:"Created,omitempty"`
    Size        int64       `json:"Size,omitempty" yaml:"Size,omitempty"`
    VirtualSize int64       `json:"VirtualSize,omitempty" yaml:"VirtualSize,omitempty"`
    ParentId    string  `json:"ParentId,omitempty" yaml:"ParentId,omitempty"`
    RepoDigests []string     `json:"RepoDigests,omitempty" yaml:"RepoDigests,omitempty"`
    Labels  map[string]string `json:"Labels,omitempty" yaml:"Labels,omitempty"`
}

func main() {
    url, err0 := url.Parse("unix:///var/run/docker.sock")
    if err0 != nil {
        fmt.Println(err0)
        return
    }
    fmt.Println(url.String(), url.Path)

    conn, err := net.Dial("unix", url.Path)
    if err != nil {
        fmt.Println(err)
        return
    }

    defer conn.Close()

    buffReader := bufio.NewReader(conn)
       
    var data io.Reader
    req, err1 := http.NewRequest("GET", "/images/json?all=0", data)
    if err1 != nil {
        fmt.Println(err1)
        return
    }
    
    req.Header.Set("User-Agent", "docker api client")
    err1 = req.Write(conn)
    if err1 != nil {
        fmt.Println(err1)
        return
    }

    resp, err2 := http.ReadResponse(buffReader, req)
    if err2 != nil {
        fmt.Println(err2)
        return
    }

    defer resp.Body.Close()
    body, err3 := ioutil.ReadAll(resp.Body)
    if err3 != nil {
        fmt.Println(err3)
        return
    }

    fmt.Println(resp.Status)
    if resp.StatusCode >= 400 {
        return
    }

    var images []DockerImage
    err = json.Unmarshal(body, &images)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    for _, image := range images {
        fmt.Println("Id:          ", image.Id)
        fmt.Println("RepoTags:    ", image.RepoTags)
        fmt.Println("Created:     ", image.Created)
        fmt.Println("Size:        ", image.Size)
        fmt.Println("VirtualSize: ", image.VirtualSize)
        fmt.Println("ParentId:    ", image.ParentId)
    }
}
