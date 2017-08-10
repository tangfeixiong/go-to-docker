package server

import (
	"fmt"
	"strings"
	"time"

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
		for _, item := range v.PortBindings {
			chc.PortBindings[nat.Port(k)] = append(chc.PortBindings[nat.Port(k)], nat.PortBinding{
				HostIP:   item.HostIp,
				HostPort: item.HostPort,
			})
		}
	}

	cnc := &network.NetworkingConfig{
		EndpointsConfig: make(map[string]*network.EndpointSettings),
	}
	for k, v := range req.NetworkConfig.EndpointsConfig {
		cnc.EndpointsConfig[k] = &network.EndpointSettings{
			IPAMConfig: &network.EndpointIPAMConfig{
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
		if v.IpamConfig != nil {
			cnc.EndpointsConfig[k].IPAMConfig.IPv4Address = v.IpamConfig.Ipv4Address
			cnc.EndpointsConfig[k].IPAMConfig.IPv6Address = v.IpamConfig.Ipv6Address
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

func (m *myService) inspectContainer(req *pb.DockerContainerInspection) (*pb.DockerContainerInspection, error) {
	resp := new(pb.DockerContainerInspection)
	if nil == req || nil == req.ContainerInfo || nil == req.ContainerInfo.ContainerJsonBase || "" == req.ContainerInfo.ContainerJsonBase.Id {
		return resp, fmt.Errorf("Request required")
	}

	ctl := dockerctl.NewEngine1_12Client()
	result, err := ctl.InspectContainer(req.ContainerInfo.ContainerJsonBase.Id)
	if nil != err {
		return resp, err
	}

	resp.ContainerInfo = &moby.ContainerJSON{
		ContainerJsonBase: &moby.ContainerJSONBase{
			Id:      result.ContainerJSONBase.ID,
			Created: result.ContainerJSONBase.Created,
			Path:    result.ContainerJSONBase.Path,
			Args:    result.ContainerJSONBase.Args,
			State: &moby.ContainerState{
				Status:     result.ContainerJSONBase.State.Status,
				Running:    result.ContainerJSONBase.State.Running,
				Paused:     result.ContainerJSONBase.State.Paused,
				Restarting: result.ContainerJSONBase.State.Restarting,
				OomKilled:  result.ContainerJSONBase.State.OOMKilled,
				Dead:       result.ContainerJSONBase.State.Dead,
				Pid:        int32(result.ContainerJSONBase.State.Pid),
				ExitCode:   int32(result.ContainerJSONBase.State.ExitCode),
				Error:      result.ContainerJSONBase.State.Error,
				StartedAt:  result.ContainerJSONBase.State.StartedAt,
				FinishedAt: result.ContainerJSONBase.State.FinishedAt,
			},
			Image:           result.ContainerJSONBase.Image,
			ResolvConfPath:  result.ContainerJSONBase.ResolvConfPath,
			HostnamePath:    result.ContainerJSONBase.HostnamePath,
			HostsPath:       result.ContainerJSONBase.HostsPath,
			LogPath:         result.ContainerJSONBase.LogPath,
			Node:            &moby.ContainerNode{},
			Name:            result.ContainerJSONBase.Name,
			RestartCount:    int32(result.ContainerJSONBase.RestartCount),
			Driver:          result.ContainerJSONBase.Driver,
			MountLabel:      result.ContainerJSONBase.MountLabel,
			ProcessLabel:    result.ContainerJSONBase.ProcessLabel,
			AppArmorProfile: result.ContainerJSONBase.AppArmorProfile,
			ExecIds:         result.ContainerJSONBase.ExecIDs,
			HostConfig: &moby.HostConfig{
				Binds:           result.ContainerJSONBase.HostConfig.Binds,
				ContainerIdFile: result.ContainerJSONBase.HostConfig.ContainerIDFile,
				LogConfig: &moby.LogConfig{
					Type:   result.ContainerJSONBase.HostConfig.LogConfig.Type,
					Config: result.ContainerJSONBase.HostConfig.LogConfig.Config,
				},
				NetworkMode: string(result.ContainerJSONBase.HostConfig.NetworkMode),
				PortBindings: &moby.PortMap{
					Value: make(map[string]*moby.PortMap_PortBindings, len(result.ContainerJSONBase.HostConfig.PortBindings)),
				},
				RestartPolicy: &moby.RestartPolicy{
					Name:              result.ContainerJSONBase.HostConfig.RestartPolicy.Name,
					MaximumRetryCount: int32(result.ContainerJSONBase.HostConfig.RestartPolicy.MaximumRetryCount),
				},
				AutoRemove:   result.ContainerJSONBase.HostConfig.AutoRemove,
				VolumeDriver: result.ContainerJSONBase.HostConfig.VolumeDriver,
				VolumesFrom:  result.ContainerJSONBase.HostConfig.VolumesFrom,

				CapAdd:          result.ContainerJSONBase.HostConfig.CapAdd[:],
				CapDrop:         result.ContainerJSONBase.HostConfig.CapDrop[:],
				Dns:             result.ContainerJSONBase.HostConfig.DNS,
				DnsOptions:      result.ContainerJSONBase.HostConfig.DNSOptions,
				DnsSearch:       result.ContainerJSONBase.HostConfig.DNSSearch,
				ExtraHosts:      result.ContainerJSONBase.HostConfig.ExtraHosts,
				GroupAdd:        result.ContainerJSONBase.HostConfig.GroupAdd,
				IpcMode:         string(result.ContainerJSONBase.HostConfig.IpcMode),
				Cgroup:          string(result.ContainerJSONBase.HostConfig.Cgroup),
				Links:           result.ContainerJSONBase.HostConfig.Links,
				OomScoreAdj:     int32(result.ContainerJSONBase.HostConfig.OomScoreAdj),
				PidMode:         string(result.ContainerJSONBase.HostConfig.PidMode),
				Privileged:      result.ContainerJSONBase.HostConfig.Privileged,
				PublishAllPorts: result.ContainerJSONBase.HostConfig.PublishAllPorts,
				ReadonlyRootfs:  result.ContainerJSONBase.HostConfig.ReadonlyRootfs,
				SecurityOpt:     result.ContainerJSONBase.HostConfig.SecurityOpt,
				StorageOpt:      result.ContainerJSONBase.HostConfig.StorageOpt,
				Tmpfs:           result.ContainerJSONBase.HostConfig.Tmpfs,
				UtsMode:         string(result.ContainerJSONBase.HostConfig.UTSMode),
				UsernsMode:      string(result.ContainerJSONBase.HostConfig.UsernsMode),
				ShmSize:         result.ContainerJSONBase.HostConfig.ShmSize,
				Sysctls:         result.ContainerJSONBase.HostConfig.Sysctls,
				// Runtime:         result.ContainerJSONBase.HostConfig.Runtime,
				// ConsoleSize:     [2]uint32{},
				ConsoleSizeHeight: uint32(result.ContainerJSONBase.HostConfig.ConsoleSize[0]),
				ConsoleSizeWidth:  uint32(result.ContainerJSONBase.HostConfig.ConsoleSize[1]),
				Isolation:         string(result.ContainerJSONBase.HostConfig.Isolation),
				Resources: &moby.Resources{
					// CPUShares:            result.ContainerJSONBase.HostConfig.Resources.CpuShares,
					// Memory:               result.ContainerJSONBase.HostConfig.Resources.Memory,
					// NanoCPUs:             result.ContainerJSONBase.HostConfig.Resources.NanoCpus,
					CgroupParent: result.ContainerJSONBase.HostConfig.Resources.CgroupParent,
					// BlkioWeight:          uint16(result.ContainerJSONBase.HostConfig.Resources.BlkioWeight),
					BlkioWeightDevice:    make([]*moby.WeightDevice, len(result.ContainerJSONBase.HostConfig.Resources.BlkioWeightDevice)),
					BlkioDeviceReadBps:   make([]*moby.ThrottleDevice, len(result.ContainerJSONBase.HostConfig.Resources.BlkioDeviceReadBps)),
					BlkioDeviceWriteBps:  make([]*moby.ThrottleDevice, len(result.ContainerJSONBase.HostConfig.Resources.BlkioDeviceWriteBps)),
					BlkioDeviceReadIops:  make([]*moby.ThrottleDevice, len(result.ContainerJSONBase.HostConfig.Resources.BlkioDeviceReadIOps)),
					BlkioDeviceWriteIops: make([]*moby.ThrottleDevice, len(result.ContainerJSONBase.HostConfig.Resources.BlkioDeviceWriteIOps)),
					// CPUPeriod:            result.ContainerJSONBase.HostConfig.Resources.CpuPeriod,
					// CPUQuota:             result.ContainerJSONBase.HostConfig.Resources.CpuQuota,
					// CPURealtimePeriod:    result.ContainerJSONBase.HostConfig.Resources.CpuRealtimePeriod,
					// CPURealtimeRuntime:   result.ContainerJSONBase.HostConfig.Resources.CpuRealtimeRuntime,
					CpusetCpus: result.ContainerJSONBase.HostConfig.Resources.CpusetCpus,
					CpusetMems: result.ContainerJSONBase.HostConfig.Resources.CpusetMems,
					Devices:    make([]*moby.DeviceMapping, len(result.ContainerJSONBase.HostConfig.Resources.Devices)),
					// DeviceCgroupRules:  result.ContainerJSONBase.HostConfig.Resources.DeviceCgroupRules,
					DiskQuota:          result.ContainerJSONBase.HostConfig.Resources.DiskQuota,
					KernelMemory:       result.ContainerJSONBase.HostConfig.Resources.KernelMemory,
					MemoryReservation:  result.ContainerJSONBase.HostConfig.Resources.MemoryReservation,
					MemorySwap:         result.ContainerJSONBase.HostConfig.Resources.MemorySwap,
					MemorySwappiness:   *result.ContainerJSONBase.HostConfig.Resources.MemorySwappiness,
					OomKillDisable:     *result.ContainerJSONBase.HostConfig.Resources.OomKillDisable,
					PidsLimit:          result.ContainerJSONBase.HostConfig.Resources.PidsLimit,
					Ulimits:            make([]*moby.Ulimit, len(result.ContainerJSONBase.HostConfig.Resources.Ulimits)),
					CpuCount:           result.ContainerJSONBase.HostConfig.Resources.CPUCount,
					CpuPercent:         result.ContainerJSONBase.HostConfig.Resources.CPUPercent,
					IoMaximumIops:      result.ContainerJSONBase.HostConfig.Resources.IOMaximumIOps,
					IoMaximumBandwidth: result.ContainerJSONBase.HostConfig.Resources.IOMaximumBandwidth,
				},
				// Mounts: make([]*moby.Mount, len(result.ContainerJSONBase.HostConfig.Mounts)),
				// Init:   &result.ContainerJSONBase.HostConfig.Init,

			},
			GraphDriver: &moby.GraphDriverData{
				Name: result.ContainerJSONBase.GraphDriver.Name,
				Data: result.ContainerJSONBase.GraphDriver.Data,
			},
			SizeRw:     0,
			SizeRootFs: 0,
		},
		Mounts: make([]*moby.MountPoint, len(result.Mounts)),
		Config: &moby.Config{
			Hostname:     result.Config.Hostname,
			Domainname:   result.Config.Domainname,
			User:         result.Config.User,
			AttachStdin:  result.Config.AttachStdin,
			AttachStdout: result.Config.AttachStdout,
			AttachStderr: result.Config.AttachStderr,
			ExposedPorts: &moby.PortSet{
				Value: make(map[string]string, len(result.Config.ExposedPorts)),
			},
			Tty:             result.Config.Tty,
			OpenStdin:       result.Config.OpenStdin,
			StdinOnce:       result.Config.StdinOnce,
			Env:             result.Config.Env,
			Cmd:             result.Config.Cmd[:],
			ArgsEscaped:     result.Config.ArgsEscaped,
			Image:           result.Config.Image,
			Volumes:         make(map[string]string, len(result.Config.Volumes)),
			WorkingDir:      result.Config.WorkingDir,
			Entrypoint:      result.Config.Entrypoint[:],
			NetworkDisabled: result.Config.NetworkDisabled,
			MacAddress:      result.Config.MacAddress,
			OnBuild:         result.Config.OnBuild,
			Labels:          result.Config.Labels,
			StopSignal:      result.Config.StopSignal,
		},
		NetworkSettings: &moby.NetworkSettings{
			NetworkSettingsBase: &moby.NetworkSettingsBase{
				Bridge:                 result.NetworkSettings.NetworkSettingsBase.Bridge,
				SandboxId:              result.NetworkSettings.NetworkSettingsBase.SandboxID,
				HairpinMode:            result.NetworkSettings.NetworkSettingsBase.HairpinMode,
				LinkLocalIpv6Address:   result.NetworkSettings.NetworkSettingsBase.LinkLocalIPv6Address,
				LinkLocalIpv6PrefixLen: int32(result.NetworkSettings.NetworkSettingsBase.LinkLocalIPv6PrefixLen),
				Ports: &moby.PortMap{
					Value: make(map[string]*moby.PortMap_PortBindings, len(result.NetworkSettings.NetworkSettingsBase.Ports)),
				},
				SandboxKey:             result.NetworkSettings.NetworkSettingsBase.SandboxKey,
				SecondaryIpAddresses:   make([]*moby.Address, len(result.NetworkSettings.NetworkSettingsBase.SecondaryIPAddresses)),
				SecondaryIpv6Addresses: make([]*moby.Address, len(result.NetworkSettings.NetworkSettingsBase.SecondaryIPv6Addresses)),
			},
			DefaultNetworkSettings: &moby.DefaultNetworkSettings{
				EndpointId:          result.NetworkSettings.DefaultNetworkSettings.EndpointID,
				Gateway:             result.NetworkSettings.DefaultNetworkSettings.Gateway,
				GlobalIpv6Address:   result.NetworkSettings.DefaultNetworkSettings.GlobalIPv6Address,
				GlobalIpv6PrefixLen: int32(result.NetworkSettings.DefaultNetworkSettings.GlobalIPv6PrefixLen),
				IpAddress:           result.NetworkSettings.DefaultNetworkSettings.IPAddress,
				IpPrefixLen:         int32(result.NetworkSettings.DefaultNetworkSettings.IPPrefixLen),
				Ipv6Gateway:         result.NetworkSettings.DefaultNetworkSettings.IPv6Gateway,
				MacAddress:          result.NetworkSettings.DefaultNetworkSettings.MacAddress,
			},
			Networks: make(map[string]*moby.EndpointSettings, len(result.NetworkSettings.Networks)),
		},
	}

	if result.ContainerJSONBase.Node != nil {
		resp.ContainerInfo.ContainerJsonBase.Node = &moby.ContainerNode{
			Id:        result.ContainerJSONBase.Node.ID,
			IpAddress: result.ContainerJSONBase.Node.IPAddress,
			Addr:      result.ContainerJSONBase.Node.Addr,
			Name:      result.ContainerJSONBase.Node.Name,
			Cpus:      int32(result.ContainerJSONBase.Node.Cpus),
			Memory:    int32(result.ContainerJSONBase.Node.Memory),
			Labels:    result.ContainerJSONBase.Node.Labels,
		}
	}
	if result.ContainerJSONBase.SizeRw != nil {
		resp.ContainerInfo.ContainerJsonBase.SizeRw = *result.ContainerJSONBase.SizeRw
	}
	if result.ContainerJSONBase.SizeRootFs != nil {
		resp.ContainerInfo.ContainerJsonBase.SizeRootFs = *result.ContainerJSONBase.SizeRootFs
	}

	for k, v := range result.ContainerJSONBase.HostConfig.PortBindings {
		bs := make([]*moby.PortBinding, len(v))
		for i := 0; i < len(v); i++ {
			bs[i] = &moby.PortBinding{
				HostIp:   v[i].HostIP,
				HostPort: v[i].HostPort,
			}
		}
		resp.ContainerInfo.ContainerJsonBase.HostConfig.PortBindings.Value[string(k)] = &moby.PortMap_PortBindings{
			PortBindings: bs,
		}
	}
	for i := 0; i < len(result.ContainerJSONBase.HostConfig.Resources.BlkioWeightDevice); i++ {
		resp.ContainerInfo.ContainerJsonBase.HostConfig.Resources.BlkioWeightDevice[i] = &moby.WeightDevice{
			Path:   result.ContainerJSONBase.HostConfig.Resources.BlkioWeightDevice[i].Path,
			Weight: int32(result.ContainerJSONBase.HostConfig.Resources.BlkioWeightDevice[i].Weight),
		}
	}
	for i := 0; i < len(result.ContainerJSONBase.HostConfig.Resources.BlkioDeviceReadBps); i++ {
		resp.ContainerInfo.ContainerJsonBase.HostConfig.Resources.BlkioDeviceReadBps[i] = &moby.ThrottleDevice{
			Path: result.ContainerJSONBase.HostConfig.Resources.BlkioDeviceReadBps[i].Path,
			Rate: result.ContainerJSONBase.HostConfig.Resources.BlkioDeviceReadBps[i].Rate,
		}
	}
	for i := 0; i < len(result.ContainerJSONBase.HostConfig.Resources.BlkioDeviceWriteBps); i++ {
		resp.ContainerInfo.ContainerJsonBase.HostConfig.Resources.BlkioDeviceWriteBps[i] = &moby.ThrottleDevice{
			Path: result.ContainerJSONBase.HostConfig.Resources.BlkioDeviceWriteBps[i].Path,
			Rate: result.ContainerJSONBase.HostConfig.Resources.BlkioDeviceWriteBps[i].Rate,
		}
	}
	for i := 0; i < len(result.ContainerJSONBase.HostConfig.Resources.BlkioDeviceReadIOps); i++ {
		resp.ContainerInfo.ContainerJsonBase.HostConfig.Resources.BlkioDeviceReadIops[i] = &moby.ThrottleDevice{
			Path: result.ContainerJSONBase.HostConfig.Resources.BlkioDeviceReadIOps[i].Path,
			Rate: result.ContainerJSONBase.HostConfig.Resources.BlkioDeviceReadIOps[i].Rate,
		}
	}
	for i := 0; i < len(result.ContainerJSONBase.HostConfig.Resources.BlkioDeviceWriteIOps); i++ {
		resp.ContainerInfo.ContainerJsonBase.HostConfig.Resources.BlkioDeviceWriteIops[i] = &moby.ThrottleDevice{
			Path: result.ContainerJSONBase.HostConfig.Resources.BlkioDeviceWriteIOps[i].Path,
			Rate: result.ContainerJSONBase.HostConfig.Resources.BlkioDeviceWriteIOps[i].Rate,
		}
	}
	for i := 0; i < len(result.ContainerJSONBase.HostConfig.Resources.Devices); i++ {
		resp.ContainerInfo.ContainerJsonBase.HostConfig.Resources.Devices[i] = &moby.DeviceMapping{
			PathOnHost:        result.ContainerJSONBase.HostConfig.Resources.Devices[i].PathOnHost,
			PathInContainer:   result.ContainerJSONBase.HostConfig.Resources.Devices[i].PathInContainer,
			CgroupPermissions: result.ContainerJSONBase.HostConfig.Resources.Devices[i].CgroupPermissions,
		}
	}
	for i := 0; i < len(result.ContainerJSONBase.HostConfig.Resources.Ulimits); i++ {
		resp.ContainerInfo.ContainerJsonBase.HostConfig.Resources.Ulimits[i] = &moby.Ulimit{
			Name: result.ContainerJSONBase.HostConfig.Resources.Ulimits[i].Name,
			Hard: result.ContainerJSONBase.HostConfig.Resources.Ulimits[i].Hard,
			Soft: result.ContainerJSONBase.HostConfig.Resources.Ulimits[i].Soft,
		}
	}
	for i := 0; i < len(result.Mounts); i++ {
		resp.ContainerInfo.Mounts[i] = &moby.MountPoint{
			Name:        result.Mounts[i].Name,
			Source:      result.Mounts[i].Source,
			Destination: result.Mounts[i].Destination,
			Driver:      result.Mounts[i].Driver,
			Mode:        result.Mounts[i].Mode,
			Rw:          result.Mounts[i].RW,
			Propagation: result.Mounts[i].Propagation,
		}
	}
	for k, _ := range result.Config.ExposedPorts {
		resp.ContainerInfo.Config.ExposedPorts.Value[string(k)] = string(k)
	}
	for k, _ := range result.Config.Volumes {
		resp.ContainerInfo.Config.Volumes[string(k)] = string(k)
	}

	for k, v := range result.NetworkSettings.NetworkSettingsBase.Ports {
		bs := make([]*moby.PortBinding, len(v))
		for i := 0; i < len(v); i++ {
			bs[i] = &moby.PortBinding{
				HostIp:   v[i].HostIP,
				HostPort: v[i].HostPort,
			}
		}
		resp.ContainerInfo.NetworkSettings.NetworkSettingsBase.Ports.Value[string(k)] = &moby.PortMap_PortBindings{
			PortBindings: bs,
		}
	}
	for i := 0; i < len(result.NetworkSettings.NetworkSettingsBase.SecondaryIPAddresses); i++ {
		resp.ContainerInfo.NetworkSettings.NetworkSettingsBase.SecondaryIpAddresses[i] = &moby.Address{
			Addr:      result.NetworkSettings.NetworkSettingsBase.SecondaryIPAddresses[i].Addr,
			PrefixLen: int32(result.NetworkSettings.NetworkSettingsBase.SecondaryIPAddresses[i].PrefixLen),
		}
	}
	for i := 0; i < len(result.NetworkSettings.NetworkSettingsBase.SecondaryIPv6Addresses); i++ {
		resp.ContainerInfo.NetworkSettings.NetworkSettingsBase.SecondaryIpv6Addresses[i] = &moby.Address{
			Addr:      result.NetworkSettings.NetworkSettingsBase.SecondaryIPv6Addresses[i].Addr,
			PrefixLen: int32(result.NetworkSettings.NetworkSettingsBase.SecondaryIPv6Addresses[i].PrefixLen),
		}
	}
	for k, v := range result.NetworkSettings.Networks {
		resp.ContainerInfo.NetworkSettings.Networks[k] = &moby.EndpointSettings{
			IpamConfig: &moby.EndpointIPAMConfig{
				Ipv4Address: v.IPAMConfig.IPv4Address,
				Ipv6Address: v.IPAMConfig.IPv6Address,
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
		}
	}

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
		// for _, v := range item.Ports {
		for i := 0; i < len(item.Ports); i++ {
			// resp.Instantiation[len(resp.Instantiation)-1].Ports = append(resp.Instantiation[len(resp.Instantiation)-1].Ports, &moby.Port{
			resp.Instantiation[len(resp.Instantiation)-1].Ports[i] = &moby.Port{
				Ip:          item.Ports[i].IP,
				PrivatePort: int32(item.Ports[i].PrivatePort),
				PublicPort:  int32(item.Ports[i].PublicPort),
				Type:        item.Ports[i].Type,
			}
		}
		if len(item.HostConfig.NetworkMode) > 0 {
			resp.Instantiation[len(resp.Instantiation)-1].HostConfig.NetworkMode = item.HostConfig.NetworkMode
		}
		if len(item.NetworkSettings.Networks) > 0 {
			for k, v := range item.NetworkSettings.Networks {
				resp.Instantiation[len(resp.Instantiation)-1].NetworkSettings.Networks[k] = &moby.EndpointSettings{
					IpamConfig: &moby.EndpointIPAMConfig{
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
				if v.IPAMConfig != nil {
					resp.Instantiation[len(resp.Instantiation)-1].NetworkSettings.Networks[k].IpamConfig.Ipv4Address = v.IPAMConfig.IPv4Address
					resp.Instantiation[len(resp.Instantiation)-1].NetworkSettings.Networks[k].IpamConfig.Ipv6Address = v.IPAMConfig.IPv6Address
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
