package server

import (
	"fmt"

	"github.com/docker/docker/api/types/blkiodev"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/go-connections/nat"
	"github.com/docker/go-units"

	"github.com/tangfeixiong/go-to-docker/pb"
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
		StopTimeout:     new(int),
		Shell:           make(strslice.StrSlice, 0),
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
	*(cc.StopTimeout) = int(req.Config.StopTimeout)
	cc.Shell = append(cc.Shell, req.Config.Shell...)

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
		Runtime:         req.HostConfig.Runtime,
		ConsoleSize:     [2]uint{},
		Isolation:       "",
		Resources: container.Resources{
			CPUShares:            req.HostConfig.Resources.CpuShares,
			Memory:               req.HostConfig.Resources.Memory,
			NanoCPUs:             req.HostConfig.Resources.NanoCpus,
			CgroupParent:         "",
			BlkioWeight:          uint16(req.HostConfig.Resources.BlkioWeight),
			BlkioWeightDevice:    make([]*blkiodev.WeightDevice, 0),
			BlkioDeviceReadBps:   make([]*blkiodev.ThrottleDevice, 0),
			BlkioDeviceWriteBps:  make([]*blkiodev.ThrottleDevice, 0),
			BlkioDeviceReadIOps:  make([]*blkiodev.ThrottleDevice, 0),
			BlkioDeviceWriteIOps: make([]*blkiodev.ThrottleDevice, 0),
			CPUPeriod:            req.HostConfig.Resources.CpuPeriod,
			CPUQuota:             req.HostConfig.Resources.CpuQuota,
			CPURealtimePeriod:    req.HostConfig.Resources.CpuRealtimePeriod,
			CPURealtimeRuntime:   req.HostConfig.Resources.CpuRealtimeRuntime,
			CpusetCpus:           "",
			CpusetMems:           "",
			Devices:              make([]container.DeviceMapping, 0),
			DeviceCgroupRules:    make([]string, 0),
			DiskQuota:            0,
			KernelMemory:         0,
			MemoryReservation:    0,
			MemorySwap:           0,
			MemorySwappiness:     new(int64),
			OomKillDisable:       new(bool),
			PidsLimit:            0,
			Ulimits:              make([]*units.Ulimit, 0),
			CPUCount:             0,
			CPUPercent:           0,
			IOMaximumIOps:        0,
			IOMaximumBandwidth:   0,
		},
		Mounts: make([]mount.Mount, 0),
		Init:   new(bool),
	}

	cnc := &network.NetworkingConfig{
		EndpointsConfig: make(map[string]*network.EndpointSettings),
	}
	for k, v := range req.NetworkConfig.EndpointsConfig {
		cnc.EndpointsConfig[k] = &network.EndpointSettings{
			IPAMConfig: &network.EndpointIPAMConfig{
				IPv4Address:  v.IpamConfig.Ipv4Address,
				IPv6Address:  v.IpamConfig.Ipv6Address,
				LinkLocalIPs: make([]string, 0),
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
			DriverOpts:          make(map[string]string),
		}
	}

	if result, err := dockerctl.NewMobyClient("1.12").CreateContainer(cc, chc, cnc, req.ContainerName); nil != err {
		resp.StateCode = 100
		resp.StateMessage = "Failed to create docker container. " + err.Error()
		return resp, fmt.Errorf("Failed to create docker container. %v", err)
	} else {
		resp.ContainerId = result.ID
		resp.Config = req.Config
		resp.HostConfig = req.HostConfig
		resp.NetworkConfig = req.NetworkConfig
		return resp, nil
	}
}

func (m *myService) provisionTarget() {

}
