package server

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	dockerref "github.com/docker/distribution/reference"
	dockermessage "github.com/docker/docker/pkg/jsonmessage"
	// dockerstdcopy "github.com/docker/docker/pkg/stdcopy"
	// dockerapi "github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/blkiodev"
	"github.com/docker/engine-api/types/container"
	"github.com/docker/engine-api/types/filters"
	// "github.com/docker/engine-api/types/mount"
	"github.com/docker/engine-api/types/network"
	"github.com/docker/engine-api/types/strslice"
	"github.com/docker/go-connections/nat"
	"github.com/docker/go-units"
	"github.com/golang/glog"

	"golang.org/x/net/context"
	"k8s.io/kubernetes/pkg/util/parsers"

	"github.com/tangfeixiong/go-to-docker/pb"
	"github.com/tangfeixiong/go-to-docker/pb/moby"
	"github.com/tangfeixiong/go-to-docker/pkg/dockerctl"
)

func (m *myService) runContainer(req *pb.DockerRunData) (*pb.DockerRunData, error) {
	resp := new(pb.DockerRunData)
	if nil == req {
		return resp, fmt.Errorf("Request required")
	}

	if nil == req.Config || 0 == len(req.Config.Image) {
		return resp, fmt.Errorf("Image required")
	}
	if nil == req.HostConfig {
		return resp, fmt.Errorf("Host config required")
	}
	if nil == req.NetworkConfig {
		return resp, fmt.Errorf("Network config required")
	}

	cc := &container.Config{
		Hostname:     req.Config.Hostname,
		Domainname:   req.Config.Domainname,
		User:         req.Config.User,
		AttachStdin:  req.Config.AttachStdin,
		AttachStdout: req.Config.AttachStdout,
		AttachStderr: req.Config.AttachStderr,
		ExposedPorts: nat.PortSet{},
		Tty:          req.Config.Tty,
		OpenStdin:    req.Config.OpenStdin,
		StdinOnce:    req.Config.StdinOnce,
		Env:          make([]string, 0),
		Cmd:          make(strslice.StrSlice, 0),
		// HealthCheck:     new(container.HealthCheck),
		ArgsEscaped: req.Config.ArgsEscaped,
		Image:       req.Config.Image,
		// Volumes:         make(map[string]struct{}),
		WorkingDir:      req.Config.WorkingDir,
		Entrypoint:      req.Config.Entrypoint,
		NetworkDisabled: req.Config.NetworkDisabled,
		MacAddress:      req.Config.MacAddress,
		OnBuild:         make([]string, 0),
		Labels:          make(map[string]string),
		StopSignal:      req.Config.StopSignal,
		// StopTimeout:     new(int),
		// Shell:           make(strslice.StrSlice, 0),
	}
	for k, _ := range req.Config.ExposedPorts.Value {
		cc.ExposedPorts[nat.Port(k)] = struct{}{}
	}
	cc.Env = append(cc.Env, req.Config.Env...)
	cc.Cmd = append(cc.Cmd, req.Config.Cmd...)
	cc.OnBuild = append(cc.OnBuild, req.Config.OnBuild...)
	for k, v := range req.Config.Labels {
		cc.Labels[k] = v
	}
	// *(cc.StopTimeout) = int(req.Config.StopTimeout)
	// cc.Shell = append(cc.Shell, req.Config.Shell...)

	chc := &container.HostConfig{
		Binds:           make([]string, 0),
		ContainerIDFile: req.HostConfig.ContainerIdFile,
		LogConfig:       container.LogConfig{},
		NetworkMode:     "",
		PortBindings:    nat.PortMap{},
		RestartPolicy: container.RestartPolicy{
			Name:              "none",
			MaximumRetryCount: 3,
		},
		AutoRemove:   req.HostConfig.AutoRemove,
		VolumeDriver: req.HostConfig.VolumeDriver,
		VolumesFrom:  make([]string, 0),

		CapAdd:          make(strslice.StrSlice, 0),
		CapDrop:         make(strslice.StrSlice, 0),
		DNS:             make([]string, 0),
		DNSOptions:      make([]string, 0),
		DNSSearch:       make([]string, 0),
		ExtraHosts:      make([]string, 0),
		GroupAdd:        make([]string, 0),
		IpcMode:         "",
		Cgroup:          "",
		Links:           make([]string, 0),
		OomScoreAdj:     0,
		PidMode:         "",
		Privileged:      req.HostConfig.Privileged,
		PublishAllPorts: req.HostConfig.PublishAllPorts,
		ReadonlyRootfs:  req.HostConfig.ReadonlyRootfs,
		SecurityOpt:     make([]string, 0),
		StorageOpt:      make(map[string]string),
		Tmpfs:           make(map[string]string),
		UTSMode:         "",
		UsernsMode:      "",
		ShmSize:         0,
		Sysctls:         make(map[string]string),
		// Runtime:         req.HostConfig.Runtime,
		// ConsoleSize:     [2]uint{},
		ConsoleSize: [2]int{},
		Isolation:   "",
		Resources: container.Resources{
			// CPUShares:            req.HostConfig.Resources.CpuShares,
			// Memory:               req.HostConfig.Resources.Memory,
			// NanoCPUs:             req.HostConfig.Resources.NanoCpus,
			CgroupParent: "",
			// BlkioWeight:          uint16(req.HostConfig.Resources.BlkioWeight),
			BlkioWeightDevice:    make([]*blkiodev.WeightDevice, 0),
			BlkioDeviceReadBps:   make([]*blkiodev.ThrottleDevice, 0),
			BlkioDeviceWriteBps:  make([]*blkiodev.ThrottleDevice, 0),
			BlkioDeviceReadIOps:  make([]*blkiodev.ThrottleDevice, 0),
			BlkioDeviceWriteIOps: make([]*blkiodev.ThrottleDevice, 0),
			// CPUPeriod:            req.HostConfig.Resources.CpuPeriod,
			// CPUQuota:             req.HostConfig.Resources.CpuQuota,
			// CPURealtimePeriod:    req.HostConfig.Resources.CpuRealtimePeriod,
			// CPURealtimeRuntime:   req.HostConfig.Resources.CpuRealtimeRuntime,
			CpusetCpus: "",
			CpusetMems: "",
			Devices:    make([]container.DeviceMapping, 0),
			// DeviceCgroupRules:  make([]string, 0),
			DiskQuota:          0,
			KernelMemory:       0,
			MemoryReservation:  0,
			MemorySwap:         0,
			MemorySwappiness:   new(int64),
			OomKillDisable:     new(bool),
			PidsLimit:          0,
			Ulimits:            make([]*units.Ulimit, 0),
			CPUCount:           0,
			CPUPercent:         0,
			IOMaximumIOps:      0,
			IOMaximumBandwidth: 0,
		},
		// Mounts: make([]mount.Mount, 0),
		// Init:   new(bool),
	}
	for k, v := range req.HostConfig.PortBindings.Value {
		chc.PortBindings[nat.Port(k)] = make([]nat.PortBinding, 0)
		chc.PortBindings[nat.Port(k)] = append(chc.PortBindings[nat.Port(k)], nat.PortBinding{
			HostIP:   v.HostIp,
			HostPort: v.HostPort,
		})
	}

	cnc := &network.NetworkingConfig{
		EndpointsConfig: make(map[string]*network.EndpointSettings),
	}
	for k, v := range req.NetworkConfig.EndpointsConfig {
		cnc.EndpointsConfig[k] = &network.EndpointSettings{
			IPAMConfig: &network.EndpointIPAMConfig{
				IPv4Address: v.IpamConfig.Ipv4Address,
				IPv6Address: v.IpamConfig.Ipv6Address,
				// LinkLocalIPs: make([]string, 0),
			},
			Links:               make([]string, 0),
			Aliases:             make([]string, 0),
			NetworkID:           v.NetworkId,
			EndpointID:          v.EndpointId,
			Gateway:             v.Gateway,
			IPAddress:           v.IpAddress,
			IPPrefixLen:         int(v.IpPrefixLen),
			IPv6Gateway:         v.Ipv6Gateway,
			GlobalIPv6Address:   v.GlobalIpv6Address,
			GlobalIPv6PrefixLen: int(v.GlobalIpv6PrefixLen),
			MacAddress:          v.MacAddress,
			// DriverOpts:          make(map[string]string),
		}
	}

	ctl := dockerctl.NewEngine1_12Client()
	if result, err := ctl.CreateContainer(cc, chc, cnc, req.ContainerName); nil != err {
		resp.StateCode = 100
		resp.StateMessage = err.Error()
		return resp, err
	} else if err := ctl.StartContainer(result.ID); nil != err {
		resp.StateCode = 101
		resp.StateMessage = err.Error()
		return resp, err
	} else {
		resp.ContainerId = result.ID
	}
	resp.Config = req.Config
	resp.HostConfig = req.HostConfig
	resp.NetworkConfig = req.NetworkConfig
	return resp, nil
}

func (m *myService) containersProvisioning(req *pb.ProvisioningsData) (*pb.ProvisioningsData, error) {
	resp := new(pb.ProvisioningsData)
	if nil == req || 0 == len(req.Provisionings) || 0 == len(req.Name) {
		return resp, fmt.Errorf("Request required")
	}

	if 0 == len(req.Metadata.CategoryName) {
		req.Metadata.CategoryName = "default"
	}
	if 0 == len(req.Metadata.ClassName) {
		req.Metadata.ClassName = "default"
	}
	if 0 == len(req.Metadata.FieldName) {
		req.Metadata.FieldName = "default"
	}
	if 0 == len(req.Namespace) {
		req.Namespace = "default"
	}

	resp.Name = req.Name
	resp.Namespace = req.Namespace
	resp.Metadata = req.Metadata
	resp.Provisionings = make([]*pb.DockerRunData, 0)
	for _, item := range req.Provisionings {
		if nil == item.Config.Labels {
			item.Config.Labels = make(map[string]string)
		}
		item.Config.Labels["created-by"] = fmt.Sprintf(`{
  "category_name": "%s",
  "class_name": "%s",
  "field_name": "%s"
}`, req.Metadata.CategoryName, req.Metadata.ClassName, req.Metadata.FieldName)
		item.Config.Labels["stackdocker.io"] = fmt.Sprintf("%s/%s", req.Namespace, req.Name)

		result, err := m.runContainer(item)
		if nil != err {
			return resp, err
		}
		resp.Provisionings = append(resp.Provisionings, result)
	}
	return resp, nil
}

func (m *myService) containersTerminating(req *pb.ProvisioningsData) (*pb.ProvisioningsData, error) {
	resp := new(pb.ProvisioningsData)
	if nil == req || 0 == len(req.Name) {
		return resp, fmt.Errorf("Request required")
	}
	if 0 == len(req.Namespace) {
		req.Namespace = "default"
	}

	ctl := dockerctl.NewEngine1_12Client()
	filter := filters.NewArgs()
	filter.Add("label", fmt.Sprintf("stackdocker.io=%s/%s", req.Namespace, req.Name))
	resultcontainers, err := ctl.ProcessStatusContainers(types.ContainerListOptions{
		Filter: filter,
	})
	if nil != err {
		return resp, err
	}

	resp.Name = req.Name
	resp.Namespace = req.Namespace
	resp.Metadata = req.Metadata
	resp.Provisionings = make([]*pb.DockerRunData, 0)
	for _, item := range resultcontainers {
		if strings.Title(item.Status) != "Exited" {
			if err := ctl.StopContainer(item.ID, time.Second*5); nil != err {
				return resp, fmt.Errorf("Could not stop container: %s; %s", item.ID, err.Error())
			}
		}
		if err := ctl.RemoveContainer(item.ID); nil != err {
			resp.Provisionings = append(resp.Provisionings, &pb.DockerRunData{
				StateCode:    102,
				StateMessage: err.Error(),
			})
			return resp, err
		}
		resp.Provisionings = append(resp.Provisionings, &pb.DockerRunData{
			ContainerId: item.ID,
		})
	}
	return resp, nil
}

func (m *myService) reapInstantiation(req *pb.InstantiationData) (*pb.InstantiationData, error) {
	resp := new(pb.InstantiationData)
	if nil == req || 0 == len(req.Name) {
		resp.StateCode = 10
		resp.StateMessage = "Request required"
		return resp, fmt.Errorf("Request required")
	}
	if 0 == len(req.Namespace) {
		req.Namespace = "default"
	}

	ctl := dockerctl.NewEngine1_12Client()
	filter := filters.NewArgs()
	filter.Add("label", fmt.Sprintf("stackdocker.io=%s/%s", req.Namespace, req.Name))
	resultcontainers, err := ctl.ProcessStatusContainers(types.ContainerListOptions{
		Filter: filter,
	})
	if nil != err {
		resp.StateCode = 100
		resp.StateMessage = err.Error()
		return resp, fmt.Errorf("Failed to get containers status: %s", err.Error())
	}

	resp.Name = req.Name
	resp.Namespace = req.Namespace
	resp.Metadata = req.Metadata
	resp.Instantiation = make([]*moby.Container, 0)
	for _, item := range resultcontainers {
		resp.Instantiation = append(resp.Instantiation, &moby.Container{
			Id:         item.ID,
			Names:      item.Names,
			Image:      item.Image,
			ImageId:    item.ImageID,
			Command:    item.Command,
			Created:    item.Created,
			Ports:      make([]*moby.Port, len(item.Ports)),
			SizeRw:     item.SizeRw,
			SizeRootFs: item.SizeRootFs,
			Labels:     item.Labels,
			State:      item.State,
			Status:     item.Status,
			HostConfig: &moby.Container_HostConfig{},
			NetworkSettings: &moby.SummaryNetworkSettings{
				Networks: make(map[string]*moby.EndpointSettings),
			},
			Mounts: make([]*moby.MountPoint, len(item.Mounts)),
		})
		for _, v := range item.Ports {
			resp.Instantiation[len(resp.Instantiation)-1].Ports = append(resp.Instantiation[len(resp.Instantiation)-1].Ports, &moby.Port{
				Ip:          v.IP,
				PrivatePort: int32(v.PrivatePort),
				PublicPort:  int32(v.PublicPort),
				Type:        v.Type,
			})
		}
		if len(item.HostConfig.NetworkMode) > 0 {
			resp.Instantiation[len(resp.Instantiation)-1].HostConfig.NetworkMode = item.HostConfig.NetworkMode
		}
		if len(item.NetworkSettings.Networks) > 0 {
			for k, v := range item.NetworkSettings.Networks {
				resp.Instantiation[len(resp.Instantiation)-1].NetworkSettings.Networks[k] = &moby.EndpointSettings{
					IpamConfig: &moby.EndpointIPAMConfig{
						Ipv4Address: v.IPAMConfig.IPv4Address,
						Ipv6Address: v.IPAMConfig.IPv6Address,
						// LinkLocalIps: v.IPAMConfig.LinkLoclIPs,
					},
					Links:               v.Links,
					Aliases:             v.Aliases,
					NetworkId:           v.NetworkID,
					EndpointId:          v.EndpointID,
					Gateway:             v.Gateway,
					IpAddress:           v.IPAddress,
					IpPrefixLen:         int32(v.IPPrefixLen),
					Ipv6Gateway:         v.IPv6Gateway,
					GlobalIpv6Address:   v.GlobalIPv6Address,
					GlobalIpv6PrefixLen: int32(v.GlobalIPv6PrefixLen),
					MacAddress:          v.MacAddress,
					// DriverOpts:          v.DriverOpts,
				}
			}
		}
		if len(item.Mounts) > 0 {
			for _, v := range item.Mounts {
				resp.Instantiation[len(resp.Instantiation)-1].Mounts = append(resp.Instantiation[len(resp.Instantiation)-1].Mounts, &moby.MountPoint{
					Name:        v.Name,
					Source:      v.Source,
					Destination: v.Destination,
					Driver:      v.Driver,
					Mode:        v.Mode,
					Rw:          v.RW,
					Propagation: v.Propagation,
				})
			}
		}
	}
	return resp, nil
}

// There are 2 kinds of docker operations categorized by running time:
// * Long running operation: The long running operation could run for arbitrary long time, and the running time
// usually depends on some uncontrollable factors. These operations include: PullImage, Logs, StartExec, AttachToContainer.
// * Non-long running operation: Given the maximum load of the system, the non-long running operation should finish
// in expected and usually short time. These include all other operations.
// kubeDockerClient only applies timeout on non-long running operations.
const (
	// defaultTimeout is the default timeout of short running docker operations.
	defaultTimeout = 2 * time.Minute

	// defaultShmSize is the default ShmSize to use (in bytes) if not specified.
	defaultShmSize = int64(1024 * 1024 * 64)

	// defaultImagePullingProgressReportInterval is the default interval of image pulling progress reporting.
	defaultImagePullingProgressReportInterval = 10 * time.Second

	// defaultImagePullingStuckTimeout is the default timeout for image pulling stuck. If no progress
	// is made for defaultImagePullingStuckTimeout, the image pulling will be cancelled.
	// Docker reports image progress for every 512kB block, so normally there shouldn't be too long interval
	// between progress updates.
	// TODO(random-liu): Make this configurable
	defaultImagePullingStuckTimeout = 1 * time.Minute
)

func base64EncodeAuth(auth types.AuthConfig) (string, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(auth); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(buf.Bytes()), nil
}

// progress is a wrapper of dockermessage.JSONMessage with a lock protecting it.
type progress struct {
	sync.RWMutex
	// message stores the latest docker json message.
	message *dockermessage.JSONMessage
	// timestamp of the latest update.
	timestamp time.Time
}

func newProgress() *progress {
	return &progress{timestamp: time.Now()}
}

func (p *progress) set(msg *dockermessage.JSONMessage) {
	p.Lock()
	defer p.Unlock()
	p.message = msg
	p.timestamp = time.Now()
}

func (p *progress) get() (string, time.Time) {
	p.RLock()
	defer p.RUnlock()
	if p.message == nil {
		return "No progress", p.timestamp
	}
	// The following code is based on JSONMessage.Display
	var prefix string
	if p.message.ID != "" {
		prefix = fmt.Sprintf("%s: ", p.message.ID)
	}
	if p.message.Progress == nil {
		return fmt.Sprintf("%s%s", prefix, p.message.Status), p.timestamp
	}
	return fmt.Sprintf("%s%s %s", prefix, p.message.Status, p.message.Progress.String()), p.timestamp
}

// progressReporter keeps the newest image pulling progress and periodically report the newest progress.
type progressReporter struct {
	*progress
	image  string
	cancel context.CancelFunc
	stopCh chan struct{}
}

// newProgressReporter creates a new progressReporter for specific image with specified reporting interval
func newProgressReporter(image string, cancel context.CancelFunc) *progressReporter {
	return &progressReporter{
		progress: newProgress(),
		image:    image,
		cancel:   cancel,
		stopCh:   make(chan struct{}),
	}
}

// start starts the progressReporter
func (p *progressReporter) start() {
	go func() {
		ticker := time.NewTicker(defaultImagePullingProgressReportInterval)
		defer ticker.Stop()
		for {
			// TODO(random-liu): Report as events.
			select {
			case <-ticker.C:
				progress, timestamp := p.progress.get()
				// If there is no progress for defaultImagePullingStuckTimeout, cancel the operation.
				if time.Now().Sub(timestamp) > defaultImagePullingStuckTimeout {
					glog.Errorf("Cancel pulling image %q because of no progress for %v, latest progress: %q", p.image, defaultImagePullingStuckTimeout, progress)
					p.cancel()
					return
				}
				glog.V(2).Infof("Pulling image %q: %q", p.image, progress)
			case <-p.stopCh:
				progress, _ := p.progress.get()
				glog.V(2).Infof("Stop pulling image %q: %q", p.image, progress)
				return
			}
		}
	}()
}

// stop stops the progressReporter
func (p *progressReporter) stop() {
	close(p.stopCh)
}

// applyDefaultImageTag parses a docker image string, if it doesn't contain any tag or digest,
// a default tag will be applied.
// https://github.com/kubernetes/kubernetes/pkg/kubelet/dockertools/docker.go
func applyDefaultImageTag(image string) (string, error) {
	named, err := dockerref.ParseNamed(image)
	if err != nil {
		return "", fmt.Errorf("couldn't parse image reference %q: %v", image, err)
	}
	_, isTagged := named.(dockerref.Tagged)
	_, isDigested := named.(dockerref.Digested)
	if !isTagged && !isDigested {
		named, err := dockerref.WithTag(named, parsers.DefaultImageTag)
		if err != nil {
			return "", fmt.Errorf("failed to apply default image tag %q: %v", image, err)
		}
		image = named.String()
	}
	return image, nil
}

/*
  Inspired from https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/dockertools/kube_docker_client.go
*/
func (m *myService) pullImage(req *pb.DockerPullData) (*pb.DockerPullData, error) {
	glog.Infoln("Go to pull image", req.Image)
	resp := new(pb.DockerPullData)
	if nil == req || 0 == len(req.Image) {
		return resp, fmt.Errorf("Request required")
	}
	if img, err := applyDefaultImageTag(req.Image); nil != err {
		return resp, err
	} else {
		resp.Image = img
	}

	ctl := dockerctl.NewEngine1_12Client()
	cli, err := ctl.DockerClient()
	if nil != err {
		resp.StateCode = 100
		resp.StateMessage = err.Error()
		return resp, err
	}

	auth := types.AuthConfig{
		// Username: "",
		// Password: "",
		Auth:          "",
		Email:         "",
		ServerAddress: "127.0.0.1:5000",
		// IdentityToken: "",
		// RegistryToken: "",
	}
	auth = ctl.RegistryAuth(resp.Image)

	// RegistryAuth is the base64 encoded credentials for the registry
	base64Auth, err := base64EncodeAuth(auth)
	if err != nil {
		resp.StateCode = 101
		resp.StateMessage = err.Error()
		return resp, err
	}
	opts := types.ImagePullOptions{
		RegistryAuth: base64Auth,
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	result, err := cli.ImagePull(ctx, resp.Image, opts)
	if err != nil {
		resp.StateCode = 102
		resp.StateMessage = err.Error()
		return resp, err
	}
	defer result.Close()
	reporter := newProgressReporter(resp.Image, cancel)
	reporter.start()
	defer reporter.stop()
	decoder := json.NewDecoder(result)
	for {
		var msg dockermessage.JSONMessage
		err := decoder.Decode(&msg)
		if err == io.EOF {
			break
		}
		if err != nil {
			resp.StateCode = 103
			resp.StateMessage = err.Error()
			return resp, err
		}
		if msg.Error != nil {
			resp.StateCode = 104
			resp.StateMessage = fmt.Sprintf("code: %d, message: %s, %s", msg.Error.Code, msg.Error.Message, msg.ErrorMessage)
			return resp, fmt.Errorf("Failed to pull image %s; %s", resp.Image, resp.StateMessage)
		}
		reporter.set(&msg)
	}
	resp.ProgressReport, _ = reporter.get()
	return resp, nil
}
