
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

type DockerPort struct {
    PrivatePort int64  `json:"PrivatePort,omitempty" yaml:"PrivatePort,omitempty"`
    PublicPort  int64  `json:"PublicPort,omitempty" yaml:"PublicPort,omitempty"`
    Type        string `json:"Type,omitempty" yaml:"Type,omitempty"`
    IP          string `json:"IP,omitempty" yaml:"IP,omitempty"`
}    

type DockerContainer struct {
    Id          string      `json:"Id" yaml:"Id"`
    Image       string      `json:"Image,omitempty" yaml:"Image,omitempty"`
    Command     string      `json:"Command,omitempty" yaml:"Command,omitempty"`
    Created     int64       `json:"Created,omitempty" yaml:"Created,omitempty"`
    Status      string      `json:"Status,omitempty" yaml:"Status,omitempty"`
    Ports       []DockerPort `json:"Ports,omitempty" yaml:"Ports,omitempty"`
    SizeRw      int64       `json:"SizeRw,omitempty" yaml:"SizeRw,omitempty"`
    SizeRootFs  int64 `json:"SizeRootFs,omitempty" yaml:"SizeRootFs,omitempty"`
    Names       []string     `json:"Names,omitempty" yaml:"Names,omitempty"`
    Labels map[string]string `json:"Labels,omitempty" yaml:"Labels,omitempty"`
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
    req, err1 := http.NewRequest("GET", "/containers/json?all=0", data)
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

    var containers []DockerContainer
    err = json.Unmarshal(body, &containers)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    for _, container := range containers {
        fmt.Println("Id:          ", container.Id)
        fmt.Println("Image:       ", container.Image)
        fmt.Println("Command:     ", container.Command)
        fmt.Println("Created:     ", container.Created)
        fmt.Println("Status:      ", container.Status)
        fmt.Println("Ports:       ", container.Ports)
        fmt.Println("SizeRw:      ", container.SizeRw)
        fmt.Println("SizeRootFs   ", container.SizeRootFs)
        fmt.Println("Names:       ", container.Names)
        fmt.Println("Labels:      ", container.Labels)
    }
}
