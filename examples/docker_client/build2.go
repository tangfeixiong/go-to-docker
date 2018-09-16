package main

import (
    "fmt"
    "io/ioutil"
    "os"

    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
    "github.com/jhoonb/archivex"
    "golang.org/x/net/context"
)

func main() {
    tar := new(archivex.TarFile)
    tar.Create("/path/to/tar/archieve")
    tar.AddAll("/path/to/base/folder", false)
    tar.Close()
    dockerBuildContext, err := os.Open("/path/to/tar/archieve.tar")
    defer dockerBuildContext.Close()
    defaultHeaders := map[string]string{"Content-Type": "application/tar"}
    cli, _ := client.NewClient("unix:///var/run/docker.sock", "v1.24", nil, defaultHeaders)
    options := types.ImageBuildOptions{
        SuppressOutput: true,
        Remove:         true,
        ForceRemove:    true,
        PullParent:     true,
        Tags:           []string{"xxx"}}
    buildResponse, err := cli.ImageBuild(context.Background(), dockerBuildContext, options)
    defer buildResponse.Body.Close()
    if err != nil {
        fmt.Printf("%s", err.Error())
    }
    //time.Sleep(5000 * time.Millisecond)
    fmt.Printf("********* %s **********", buildResponse.OSType)
    response, err := ioutil.ReadAll(buildResponse.Body)
    if err != nil {
        fmt.Printf("%s", err.Error())
    }
    fmt.Println(string(response))
}
Be careful that a lot of