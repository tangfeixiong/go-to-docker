
package main

import(
    "fmt"
    "net"
    "net/url"
    "net/http"
    "bytes"
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
/*
type KeyValuePair struct {
    Key         string  `json:"Key,omitempty" yaml:"Key,omitempty"`
    Value       string  `json:"Value,omitempty" yaml:"Value,omitempty"`
}
*/
type PortBinding struct {
    HostIP      string  `json:"HostIP,omitempty" yaml:"HostIP,omitempty"`
    HostPort    string  `json:"HostPort,omitempty" yaml:"HostPort,omitempty"`
}

type RestartPolicy struct {
    Name        string  `json:"Name,omitempty" yaml:"Name,omitempty"`
    MaximumRetryCount   int `json:"MaximumRetryCount,omitempty" yaml:"MaximumRetryCount,omitempty"`
}

type Device struct {
    PathOnHost string `json:"PathOnHost,omitempty" yaml:"PathOnHost,omitempty"`
    PathInContainer string  `json:"PathInContainer,omitempty" yaml:"PathInContainer,omitempty"`
    CgroupPermissions   string  `json:"CgroupPermissions,omitempty" yaml:"CgroupPermissions,omitempty"`
}

type Ulimit struct {
    Name    string  `json:"Name,omitempty" yaml:"Name,omitempty"`
    Soft    int64   `json:"Soft,omitempty" yaml:"Soft,omitempty"`
    Hard    int64   `json:"Hard,omitempty" yaml:"Hard,omitempty"`
}

type LogConfig struct {
    Type    string  `json:"Type,omitempty" yaml:"Type,omitempty"`
    Config map[string]string `json:"Config,omitempty" yaml:"Config,omitempty"`
}

type DockerHostConfig struct {
    Binds       []string    `json:"Binds,omitempty" yaml:"Binds,omitempty"`
    Links       []string    `json:"Links,omitempty" yaml:"Links,omitempty"`
    LxcConf     map[string]string  `json:"LxcConf,omitempty" yaml:"LxcConf,omitempty"`
    Memory      int64       `json:"Memory,omitempty" yaml:"Memory,omitempty"`
    MemorySwap  int64 `json:"MemorySwap,omitempty" yaml:"MemorySwap,omitempty"`
    CpuShares   int64 `json:"CpuShares,omitempty" yaml:"CpuShares,omitempty"`
    CpuPeriod   int64 `json:"CpuPeriod,omitempty" yaml:"CpuPeriod,omitempty"`
    CpusetCpus string `json:"CpusetCpus,omitempty" yaml:"CpusetCpus,omitempty"`
    CpusetMems string `json:"CpusetMems,omitempty" yaml:"CpusetMems,omitempty"`
    BlkioWeight int64 `json:"BlkioWeight,omitempty" yaml:"BlkioWeight,omitempty"`
    OomKillDisable bool     `json:"OomKillDisable,omitempty" yaml:"OomKillDisable,omitempty"`
    PortBindings map[string][]PortBinding `json:"PortBindings,omitempty" yaml:"PortBindings,omitempty"`
    PublishAllPorts bool    `json:"PublishAllPorts,omitempty" yaml:"PublishAllPorts,omitempty"`
    Privileged  bool  `json:"Privileged,omitempty" yaml:"Privileged,omitempty"`
    ReadonlyRootfs  bool    `json:"ReadonlyRootfs,omitempty" yaml:"ReadonlyRootfs,omitempty"`
    Dns         []string    `json:"Dns,omitempty" yaml:"Dns,omitempty"` // For Docker API v1.10 and above only
    DnsSearch  []string `json:"DnsSearch,omitempty" yaml:"DnsSearch,omitempty"`
    ExtraHosts      []string    `json:"ExtraHosts,omitempty" yaml:"ExtraHosts,omitempty"`
    VolumesFrom     []string    `json:"VolumesFrom,omitempty" yaml:"VolumesFrom,omitempty"`
    CapAdd      []string    `json:"CapAdd,omitempty" yaml:"CapAdd,omitempty"`
    CapDrop     []string    `json:"CapDrop,omitempty" yaml:"CapDrop,omitempty"`
    RestartPolicy   RestartPolicy `json:"RestartPolicy,omitempty" yaml:"RestartPolicy,omitempty"`
    NetworkMode     string  `json:"NetworkMode,omitempty" yaml:"NetworkMode,omitempty"`
    Devices     []Device    `json:"Devices,omitempty" yaml:"Devices,omitempty"`
    Ulimits     []Ulimit    `json:"Ulimits,omitempty" yaml:"Ulimits,omitempty"`
    LogConfig LogConfig `json:"LogConfig,omitempty" yaml:"LogConfig,omitempty"`
    SecurityOpt     []string    `json:"SecurityOpt,omitempty" yaml:"SecurityOpt,omitempty"`
    CgroupParent    string  `json:"CgroupParent,omitempty" yaml:"CgroupParent,omitempty"`
}

type DockerContainerConfig struct {
    Hostname    string  `json:"Hostname,omitempty" yaml:"Hostname,omitempty"`
    Domainname string `json:"Domainname,omitempty" yaml:"Domainname,omitempty"`
    User       string   `json:"User,omitempty" yaml:"User,omitempty"`
    AttachStdin bool `json:"AttachStdin,omitempty" yaml:"AttachStdin,omitempty"`
    AttachStdout    bool    `json:"AttachStdout,omitempty" yaml:"AttachStdout,omitempty"`
    AttachStderr    bool    `json:"AttachStderr,omitempty" yaml:"AttachStderr,omitempty"`
    Tty         bool    `json:"Tty,omitempty" yaml:"Tty,omitempty"`
    OpenStdin   bool    `json:"OpenStdin,omitempty" yaml:"OpenStdin,omitempty"`
    StdinOnce   bool    `json:"StdinOnce,omitempty" yaml:"StdinOnce,omitempty"`
    Env         []string    `json:"Env,omitempty" yaml:"Env,omitempty"`
    Cmd         []string    `json:"Cmd" yaml:"Cmd"`
    Entrypoint  []string    `json:"Entrypoint" yaml:"Entrypoint"`
    Image       string      `json:"Image,omitempty" yaml:"Image,omitempty"`
    Labels  map[string]string `json:"Labels,omitempty" yaml:"Labels,omitempty"`
    Volumes map[string]struct{} `json:"Volumes,omitempty" yaml:"Volumes,omitempty"`
    WorkingDir string `json:"WorkingDir,omitempty" yaml:"WorkingDir,omitempty"`
    NetworkDisabled bool    `json:"NetworkDisabled,omitempty" yaml:"NetworkDisabled,omitempty"`
    MacAddress string `json:"MacAddress,omitempty" yaml:"MacAddress,omitempty"`
    ExposedPorts map[string]struct{} `json:"ExposedPorts,omitempty" yaml:"ExposedPorts,omitempty"`
    HostConfig DockerHostConfig `json:"HostConfig,omitempty" yaml:"HostConfig,omitempty"`
}

type DockerContainerCreation struct {
    Id          string      `json:"Id", yaml:"Id"`
    Warnings   []string   `json:"Warnings,omitempty" yaml:"Warnings,omitempty"`
}

func main() {
    var cc DockerContainerConfig
    cc.AttachStdin = false
    cc.AttachStdout = true
    cc.AttachStderr = true
    cc.Tty = false
    cc.OpenStdin = false
    cc.StdinOnce = false
    cc.Cmd = append(cc.Cmd, "date")
    cc.Image = "ubuntu" //"tangfeixiong/exercise-golang-1-4-onbuild"
    cc.NetworkDisabled = false
    var host *DockerHostConfig = &cc.HostConfig
    //host.LxcConf = make(map[string]string)
    //host.LxcConf["lxc.utsname"] = "docker"
    host.Memory = 256000000
    host.MemorySwap = -1
    //host.CpuShares = 512
    //host.CpuPeriod = 100000
    //host.CpusetCpus = "0,1"
    //host.CpusetMems = "0,1"
    //host.BlkioWeight = 300
    //host.OomKillDisable = false
    host.PublishAllPorts = false
    host.Privileged = false
    host.ReadonlyRootfs = false
    //host.Dns = append(host.Dns, "8.8.4.4")
    //host.VolumesFrom = append(host.Dns, "parent", "other:ro")
    //host.CapAdd = append(host.CapAdd, "NET_ADMIN")
    //host.CapDrop = append(host.CapDrop, "MKNOD")
    //host.RestartPolicy = make(map[string]string)
    host.RestartPolicy.Name = ""
    host.RestartPolicy.MaximumRetryCount = 0
    host.NetworkMode = "bridge"
    host.LogConfig.Type = "json-file"
    
    buf, err_1 := json.Marshal(cc)
    if err_1 != nil {
        fmt.Println(err_1)
        return
    }

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
       
    var data io.Reader = bytes.NewBuffer(buf)
    req, err1 := http.NewRequest("POST", "/containers/create", data)
    if err1 != nil {
        fmt.Println(err1)
        return
    }
    
    req.Header.Set("User-Agent", "docker api client")
    req.Header.Set("Content-Type", "application/json")
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

    var container DockerContainerCreation
    err = json.Unmarshal(body, &container)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    fmt.Println("Id:       ", container.Id)
    fmt.Println("Warnings: ", container.Warnings)
}
