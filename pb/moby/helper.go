package moby

import (
	"bytes"
	"errors"

	dockertypes "github.com/docker/docker/api/types"
	dockercontainertypes "github.com/docker/docker/api/types/container"
	dockernetworktypes "github.com/docker/docker/api/types/network"
	dockerunitstypes "github.com/docker/go-units"
	gogoprotobuftypes "github.com/gogo/protobuf/types"
	"github.com/golang/glog"

	containerpb "github.com/tangfeixiong/go-to-docker/pb/moby/container"
	filterspb "github.com/tangfeixiong/go-to-docker/pb/moby/filters"
	networkpb "github.com/tangfeixiong/go-to-docker/pb/moby/network"
	unitspb "github.com/tangfeixiong/go-to-docker/pb/moby/units"
)

const (
	BuilderV1       string = "1"
	BuilderBuildKit string = "2"
)

var (
	ErrNilPointer                      = errors.New("found nil pointer")
	ErrImagePullOptionsIsNil           = errors.New("ImagePullOptions not specified")
	ErrImagePushOptionsIsNil           = errors.New("ImagePushOptions not specified")
	ErrImageBuildOptionsIsNil          = errors.New("ImageBuildOptions is nil")
	ErrImageBuildDockefileNotSpecified = errors.New("Dockefile not specified")
	ErrImageBuildOptCtxNotSpecified    = errors.New("Option context not specified")
	ErrNetworkCreateIsNil              = errors.New("NetworkCreate not specified")
)

func (m *ContainerListOptions) DeepCopyChecked() (*ContainerListOptions, []error) {
	if m == nil {
		return nil, []error{ErrNilPointer}
	}
	errorList := []error{}
	dst := &ContainerListOptions{
		Quiet:   m.Quiet,
		Size_:   m.Size_,
		All:     m.All,
		Latest:  m.Latest,
		Since:   m.Since,
		Before:  m.Before,
		Limit:   m.Limit,
		Filters: (*filterspb.Args)(nil),
	}
	if m.Filters != nil {
		dst.Filters = m.Filters.DeepCopyChecked()
	}
	return dst, errorList
}

func (m *ContainerListOptions) ExportIntoDockerApiType() (tgt dockertypes.ContainerListOptions) {
	if m != nil {
		tgt.Quiet = m.Quiet
		tgt.Size = m.Size_
		tgt.All = m.All
		tgt.Latest = m.Latest
		tgt.Since = m.Since
		tgt.Before = m.Before
		tgt.Limit = int(m.Limit)
		if m.Filters != nil {
			tgt.Filters = m.Filters.ExportIntoDockerApiType()
		}
	}
	return
}

func (m *ContainerRemoveOptions) ExportIntoDockerApiType() (tgt dockertypes.ContainerRemoveOptions) {
	if m != nil {
		tgt.RemoveVolumes = m.RemoveVolumes
		tgt.RemoveLinks = m.RemoveLinks
		tgt.Force = m.Force
	}
	return
}

func ConvertFromDockerApiTypeContainer(s dockertypes.Container) (dst *Container) {
	dst = &Container{
		Id:         s.ID,
		Names:      make([]string, len(s.Names)),
		Image:      s.Image,
		ImageId:    s.ImageID,
		Created:    s.Created,
		Ports:      make([]*Port, len(s.Ports)),
		SizeRw:     s.SizeRw,
		SizeRootFs: s.SizeRootFs,
		Labels:     make(map[string]string),
		State:      s.State,
		Status:     s.Status,
		HostConfig: &Container_HostConfig{
			NetworkMode: s.HostConfig.NetworkMode,
		},
		NetworkSettings: (*SummaryNetworkSettings)(nil),
		Mounts:          make([]*MountPoint, len(s.Mounts)),
	}
	for _, v := range s.Names {
		dst.Names = append(dst.Names, v)
	}
	for _, v := range s.Ports {
		dst.Ports = append(dst.Ports, &Port{
			Ip:          v.IP,
			PrivatePort: int32(v.PrivatePort),
			PublicPort:  int32(v.PublicPort),
			Type:        v.Type,
		})
	}
	for k, v := range s.Labels {
		dst.Labels[k] = v
	}
	if s.NetworkSettings != nil {
		dst.NetworkSettings = &SummaryNetworkSettings{
			Networks: make(map[string]*networkpb.EndpointSettings),
		}
		for k, v := range s.NetworkSettings.Networks {
			var ele *networkpb.EndpointSettings = nil
			if v != nil {
				ele = &networkpb.EndpointSettings{
					IpamConfig:          (*networkpb.EndpointIPAMConfig)(nil),
					Links:               make([]string, len(v.Links)),
					Aliases:             make([]string, len(v.Aliases)),
					NetworkId:           v.NetworkID,
					EndpointId:          v.EndpointID,
					Gateway:             v.Gateway,
					IpAddress:           v.IPAddress,
					IpPrefixLen:         int32(v.IPPrefixLen),
					Ipv6Gateway:         v.IPv6Gateway,
					GlobalIpv6Address:   v.GlobalIPv6Address,
					GlobalIpv6PrefixLen: int32(v.GlobalIPv6PrefixLen),
					MacAddress:          v.MacAddress,
					DriverOpts:          make(map[string]string),
				}
				if v.IPAMConfig != nil {
					ele.IpamConfig = &networkpb.EndpointIPAMConfig{
						Ipv4Address:  v.IPAMConfig.IPv4Address,
						Ipv6Address:  v.IPAMConfig.IPv6Address,
						LinkLocalIps: make([]string, len(v.IPAMConfig.LinkLocalIPs)),
					}
					for _, v1 := range v.IPAMConfig.LinkLocalIPs {
						ele.IpamConfig.LinkLocalIps = append(ele.IpamConfig.LinkLocalIps, v1)
					}
				}
				for _, v1 := range v.Links {
					ele.Links = append(ele.Links, v1)
				}
				for _, v1 := range v.Aliases {
					ele.Aliases = append(ele.Aliases, v1)
				}
				for k1, v1 := range v.DriverOpts {
					ele.DriverOpts[k1] = v1
				}
			}
			dst.NetworkSettings.Networks[k] = ele
		}
	}
	for _, v := range s.Mounts {
		dst.Mounts = append(dst.Mounts, convertFromDockerApiTypeMountPoint(v))
	}
	return dst
}

func ConvertFromDockerApiTypeContainerJSON(s *dockertypes.ContainerJSON) (dst *ContainerJSON) {
	if s != nil {
		dst = &ContainerJSON{
			ContainerJsonBase: (*ContainerJSONBase)(nil),
			Mounts:            make([]*MountPoint, len(s.Mounts)),
			Config:            (*containerpb.Config)(nil),
			NetworkSettings:   (*NetworkSettings)(nil),
		}
		dst.ContainerJsonBase = convertFromDockerApiTypeContainerJSONBase(s.ContainerJSONBase)
		for _, v := range s.Mounts {
			dst.Mounts = append(dst.Mounts, convertFromDockerApiTypeMountPoint(v))
		}
	}
	return
}

func convertFromDockerApiTypeContainerJSONBase(s *dockertypes.ContainerJSONBase) (dst *ContainerJSONBase) {
	if s == nil {
		return
	}
	dst = &ContainerJSONBase{
		Id:              s.ID,
		Created:         s.Created,
		Path:            s.Path,
		Args:            make([]string, len(s.Args)),
		State:           (*ContainerState)(nil),
		Image:           s.Image,
		ResolvConfPath:  s.ResolvConfPath,
		HostnamePath:    s.HostnamePath,
		HostsPath:       s.HostsPath,
		LogPath:         s.LogPath,
		Node:            (*ContainerNode)(nil),
		Name:            s.Name,
		RestartCount:    int32(s.RestartCount),
		Driver:          s.Driver,
		Platform:        s.Platform,
		MountLabel:      s.MountLabel,
		ProcessLabel:    s.ProcessLabel,
		AppArmorProfile: s.AppArmorProfile,
		ExecIds:         make([]string, len(s.ExecIDs)),
		HostConfig:      (*containerpb.HostConfig)(nil),
		GraphDriver:     (*GraphDriverData)(nil),
		SizeRw:          (*ContainerJSONBase_Int64Struct)(nil),
		SizeRootFs:      (*ContainerJSONBase_Int64Struct)(nil),
	}
	for _, v := range s.Args {
		dst.Args = append(dst.Args, v)
	}
	if v := s.State; v != nil {
		dst.State = &ContainerState{
			Status:     v.Status,
			Running:    v.Running,
			Paused:     v.Paused,
			Restarting: v.Restarting,
			OomKilled:  v.OOMKilled,
			Dead:       v.Dead,
			Pid:        int32(v.Pid),
			ExitCode:   int32(v.ExitCode),
			Error:      v.Error,
			StartedAt:  v.StartedAt,
			FinishedAt: v.FinishedAt,
			Health:     (*Health)(nil),
		}
		if v1 := v.Health; v1 != nil {
			dst.State.Health = &Health{
				Status:        v1.Status,
				FailingStreak: int32(v1.FailingStreak),
				Log:           make([]*HealthcheckResult, len(v1.Log)),
			}
			for _, v2 := range v1.Log {
				var ele *HealthcheckResult = nil
				if v2 != nil {
					ele = &HealthcheckResult{
						Start:    (*gogoprotobuftypes.Timestamp)(nil),
						End:      (*gogoprotobuftypes.Timestamp)(nil),
						ExitCode: int32(v2.ExitCode),
						Output:   v2.Output,
					}
					var err error
					ele.Start, err = gogoprotobuftypes.TimestampProto(v2.Start)
					if err != nil {
						glog.Warningf("Could not parse time into protobuf timestamp: %v", err)
						err = nil
					}
					ele.End, err = gogoprotobuftypes.TimestampProto(v2.End)
					if err != nil {
						glog.Warningf("Couldn't parse time into protobuf timestamep: %v", err)
					}
				}
				dst.State.Health.Log = append(dst.State.Health.Log, ele)
			}
		}
	}
	if v := s.Node; v != nil {
		dst.Node = &ContainerNode{
			Id:        v.ID,
			IpAddress: v.IPAddress,
			Addr:      v.Addr,
			Name:      v.Name,
			Cpus:      int32(v.Cpus),
			Memory:    v.Memory,
			Labels:    make(map[string]string),
		}
		for k1, v1 := range v.Labels {
			dst.Node.Labels[k1] = v1
		}
	}
	for _, v := range s.ExecIDs {
		dst.ExecIds = append(dst.ExecIds, v)
	}
	dst.HostConfig = containerpb.ConvertFromDockerApiTypeHostConfig(s.HostConfig)
	return
}

func convertFromDockerApiTypeMountPoint(s dockertypes.MountPoint) *MountPoint {
	return &MountPoint{
		Type:        string(s.Type),
		Name:        s.Name,
		Source:      s.Source,
		Destination: s.Destination,
		Driver:      s.Driver,
		Mode:        s.Mode,
		Rw:          s.RW,
		Propagation: string(s.Propagation),
	}
}

func ConvertFromDockerApiTypeContainersPruneReport(s dockertypes.ContainersPruneReport) (dst *ContainersPruneReport) {
	dst = &ContainersPruneReport{
		ContainersDeleted: make([]string, len(s.ContainersDeleted)),
		SpaceReclaimed:    s.SpaceReclaimed,
	}
	for _, v := range s.ContainersDeleted {
		dst.ContainersDeleted = append(dst.ContainersDeleted, v)
	}
	return
}

func (m *ImagePullOptions) DeepCopyChecked() (*ImagePullOptions, []error) {
	if m == nil {
		return nil, []error{ErrNilPointer}
	}
	errorList := []error{}
	tgt := &ImagePullOptions{
		All:          m.All,
		RegistryAuth: m.RegistryAuth,
		Platform:     m.Platform,
	}
	return tgt, errorList
}

func (m *ImagePullOptions) ExportIntoDockerApiType() (tgt dockertypes.ImagePullOptions) {
	if m != nil {
		tgt = dockertypes.ImagePullOptions{
			All:          m.All,
			RegistryAuth: m.RegistryAuth,
			// PrivilegeFunc: m.PrivilegeFunc,
			// Platform:      m.Platform,
		}
	}
	return
}

func (m *ImageListOptions) DeepCopyChecked() (*ImageListOptions, []error) {
	if m != nil {
		return nil, []error{ErrNilPointer}
	}
	errorList := []error{}
	tgt := &ImageListOptions{
		All:     m.All,
		Filters: new(filterspb.Args),
	}
	tgt.Filters = m.Filters.DeepCopyChecked()
	return tgt, errorList
}

func (m *ImageListOptions) ExportIntoDockerApiType() (obj dockertypes.ImageListOptions) {
	if m == nil {
		return
	}
	obj.All = m.All
	obj.Filters = m.Filters.ExportIntoDockerApiType()
	return
}

func (m *ImageRemoveOptions) DeepCopyChecked() (*ImageRemoveOptions, []error) {
	if m != nil {
		return nil, []error{ErrNilPointer}
	}
	errorList := []error{}
	tgt := &ImageRemoveOptions{
		Force:         m.Force,
		PruneChildren: m.PruneChildren,
	}
	return tgt, errorList
}

func (m *ImageRemoveOptions) ExportIntoDockerApiType() (obj dockertypes.ImageRemoveOptions) {
	if m == nil {
		return
	}
	obj.Force = m.Force
	obj.PruneChildren = m.PruneChildren
	return
}

func (m *ImagePushOptions) DeepCopyChecked() (*ImagePushOptions, []error) {
	if m == nil {
		return nil, []error{ErrNilPointer}
	}
	errorList := []error{}
	tgt := &ImagePushOptions{
		All:          m.All,
		RegistryAuth: m.RegistryAuth,
		Platform:     m.Platform,
	}
	return tgt, errorList
}

func (m *ImagePushOptions) ExportIntoDockerApiType() (obj dockertypes.ImagePushOptions) {
	if m == nil {
		return
	}
	obj.All = m.All
	obj.RegistryAuth = m.RegistryAuth
	// obj.Platform = m.Platform
	return
}

func (m *ImageBuildOptions) DeepCopyChecked() (*ImageBuildOptions, []error) {
	if m == nil {
		return nil, []error{ErrImageBuildOptionsIsNil}
	}
	errorList := []error{}
	tgt := &ImageBuildOptions{
		Tags:           make([]string, 0),
		SuppressOutput: m.SuppressOutput,
		RemoteContext:  m.RemoteContext,
		NoCache:        m.NoCache,
		Remove:         m.Remove,
		ForceRemove:    m.ForceRemove,
		PullParent:     m.PullParent,
		Isolation:      m.Isolation,
		CpuSetCpus:     m.CpuSetCpus,
		CpuSetMems:     m.CpuSetMems,
		CpuShares:      m.CpuShares,
		CpuQuota:       m.CpuQuota,
		CpuPeriod:      m.CpuPeriod,
		Memory:         m.Memory,
		MemorySwap:     m.MemorySwap,
		CgroupParent:   m.CgroupParent,
		NetworkMode:    m.NetworkMode,
		ShmSize:        m.ShmSize,
		Dockerfile:     m.Dockerfile,
		Ulimits:        make([]*unitspb.Ulimit, 0),
		BuildArgs:      make(map[string]*ImageBuildOptions_StringStruct),
		AuthConfigs:    make(map[string]*AuthConfig),
		Context:        make([]byte, 0),
		Labels:         make(map[string]string),
		Squash:         m.Squash,
		CacheFrom:      make([]string, 0),
		SecurityOpt:    make([]string, 0),
		ExtraHosts:     make([]string, 0),
		Target:         m.Target,
		SessionId:      m.SessionId,
		Platform:       m.Platform,
		Version:        m.Version,
		BuildId:        m.BuildId,
	}
	for _, v := range m.Tags {
		tgt.Tags = append(tgt.Tags, v)
	}
	if len(m.Dockerfile) == 0 {
		errorList = append(errorList, ErrImageBuildDockefileNotSpecified)
	}
	for _, v := range m.Ulimits {
		if v != nil {
			tgt.Ulimits = append(tgt.Ulimits, &unitspb.Ulimit{
				Name: v.Name,
				Hard: v.Hard,
				Soft: v.Soft,
			})
		}
	}
	for k, v := range m.BuildArgs {
		if v != nil {
			ele := &ImageBuildOptions_StringStruct{
				Value: v.Value,
			}
			tgt.BuildArgs[k] = ele
		}
	}
	for k, v := range m.AuthConfigs {
		if v != nil {
			ele := &AuthConfig{
				Username:      v.Username,
				Password:      v.Password,
				Auth:          v.Auth,
				Email:         v.Email,
				ServerAddress: v.ServerAddress,
				IdentityToken: v.IdentityToken,
				RegistryToken: v.RegistryToken,
			}
			tgt.AuthConfigs[k] = ele
		}
	}
	if len(m.Context) == 0 {
		errorList = append(errorList, ErrImageBuildOptCtxNotSpecified)
	}
	for _, v := range m.Context {
		tgt.Context = append(tgt.Context, v)
	}
	for k, v := range m.Labels {
		tgt.Labels[k] = v
	}
	for _, v := range m.CacheFrom {
		tgt.CacheFrom = append(tgt.CacheFrom, v)
	}
	for _, v := range m.SecurityOpt {
		tgt.SecurityOpt = append(tgt.SecurityOpt, v)
	}
	for _, v := range m.ExtraHosts {
		tgt.ExtraHosts = append(tgt.ExtraHosts, v)
	}

	return tgt, errorList
}

func (m *ImageBuildOptions) ExportIntoDockerApiType() (obj dockertypes.ImageBuildOptions) {
	if m == nil {
		return
	}
	if m.Tags != nil {
		obj.Tags = make([]string, len(m.Tags))
		for _, v := range m.Tags {
			obj.Tags = append(obj.Tags, v)
		}
	}
	obj.SuppressOutput = m.SuppressOutput
	obj.RemoteContext = m.RemoteContext
	obj.NoCache = m.NoCache
	obj.Remove = m.Remove
	obj.ForceRemove = m.ForceRemove
	obj.PullParent = m.PullParent
	obj.Isolation = dockercontainertypes.Isolation(m.Isolation)
	obj.CPUSetCPUs = m.CpuSetCpus
	obj.CPUSetMems = m.CpuSetMems
	obj.CPUShares = m.CpuShares
	obj.CPUQuota = m.CpuQuota
	obj.CPUPeriod = m.CpuPeriod
	obj.Memory = m.Memory
	obj.MemorySwap = m.MemorySwap
	obj.CgroupParent = m.CgroupParent
	obj.NetworkMode = m.NetworkMode
	obj.ShmSize = m.ShmSize
	obj.Dockerfile = m.Dockerfile
	if m.Ulimits != nil {
		obj.Ulimits = make([]*dockerunitstypes.Ulimit, len(m.Ulimits))
		for _, v := range m.Ulimits {
			if v != nil {
				obj.Ulimits = append(obj.Ulimits, &dockerunitstypes.Ulimit{
					Name: v.Name,
					Hard: v.Hard,
					Soft: v.Soft,
				})
			}
		}
	}
	if m.BuildArgs != nil {
		obj.BuildArgs = make(map[string]*string)
		for k, v := range m.BuildArgs {
			if v != nil {
				s := v.Value
				obj.BuildArgs[k] = &s
			} else {
				obj.BuildArgs[k] = (*string)(nil)
			}
		}
	}
	if m.AuthConfigs != nil {
		obj.AuthConfigs = make(map[string]dockertypes.AuthConfig)
		for k, v := range m.AuthConfigs {
			if v != nil {
				obj.AuthConfigs[k] = dockertypes.AuthConfig{
					Username:      v.Username,
					Password:      v.Password,
					Auth:          v.Auth,
					Email:         v.Email,
					ServerAddress: v.ServerAddress,
					IdentityToken: v.IdentityToken,
					RegistryToken: v.RegistryToken,
				}
			}
		}
	}
	if m.Context != nil {
		buf := bytes.NewBuffer(m.Context)
		obj.Context = buf
	}
	if m.Labels != nil {
		obj.Labels = make(map[string]string)
		for k, v := range m.Labels {
			obj.Labels[k] = v
		}
	}
	obj.Squash = m.Squash
	if m.CacheFrom != nil {
		obj.CacheFrom = make([]string, len(m.CacheFrom))
		for _, v := range m.CacheFrom {
			obj.CacheFrom = append(obj.CacheFrom, v)
		}
	}
	if m.SecurityOpt != nil {
		obj.SecurityOpt = make([]string, len(m.SecurityOpt))
		for _, v := range m.SecurityOpt {
			obj.SecurityOpt = append(obj.SecurityOpt, v)
		}
	}
	if m.ExtraHosts != nil {
		obj.ExtraHosts = make([]string, len(m.ExtraHosts))
		for _, v := range m.ExtraHosts {
			obj.ExtraHosts = append(obj.ExtraHosts, v)
		}
	}
	obj.Target = m.Target
	obj.SessionID = m.SessionId
	// obj.Platform = m.Platform
	// obj.Version = dockertypes.BuildVersion(m.Version)
	// obj.BuildID = m.BuildId
	return
}

func ConvertFromDockerApiTypeImageInspect(s *dockertypes.ImageInspect) (dst *ImageInspect) {
	if s != nil {
		dst = &ImageInspect{
			Id:              s.ID,
			RepoTags:        ([]string)(nil),
			RepoDigests:     ([]string)(nil),
			Parent:          s.Parent,
			Comment:         s.Comment,
			Created:         s.Created,
			Container:       s.Container,
			ContainerConfig: (*containerpb.Config)(nil),
			DockerVersion:   s.DockerVersion,
			Author:          s.Author,
			Config:          (*containerpb.Config)(nil),
			Architecture:    s.Architecture,
			Os:              s.Os,
			OsVersion:       s.OsVersion,
			Size_:           s.Size,
			VirtualSize:     s.VirtualSize,
			GraphDriver:     (*GraphDriverData)(nil),
			RootFs:          (*RootFS)(nil),
			Metadata:        (*ImageMetadata)(nil),
		}
		if s.RepoTags != nil {
			dst.RepoTags = make([]string, len(s.RepoTags))
			for _, v := range s.RepoTags {
				dst.RepoTags = append(dst.RepoTags, v)
			}
		}
		if s.RepoDigests != nil {
			dst.RepoDigests = make([]string, len(s.RepoDigests))
			for _, v := range s.RepoDigests {
				dst.RepoDigests = append(dst.RepoDigests, v)
			}
		}
		dst.ContainerConfig = containerpb.ConvertFromDockerApiTypeConfig(s.ContainerConfig)
		dst.Config = containerpb.ConvertFromDockerApiTypeConfig(s.Config)
		dst.GraphDriver = convertFromDockerApiTypeGraphDriverData(s.GraphDriver)
		dst.RootFs = convertFromDockerApiTypeRootFS(s.RootFS)
		dst.Metadata = convertFromDockerApiTypeImageMetaData(s.Metadata)
	}
	return
}

func convertFromDockerApiTypeGraphDriverData(s dockertypes.GraphDriverData) (d *GraphDriverData) {
	d = &GraphDriverData{
		Data: (map[string]string)(nil),
		Name: s.Name,
	}
	if s.Data != nil {
		d.Data = make(map[string]string)
		for k, v := range s.Data {
			d.Data[k] = v
		}
	}
	return
}

func convertFromDockerApiTypeRootFS(s dockertypes.RootFS) (d *RootFS) {
	d = &RootFS{
		Type:      s.Type,
		Layers:    ([]string)(nil),
		BaseLayer: s.BaseLayer,
	}
	if s.Layers != nil {
		for _, v := range s.Layers {
			d.Layers = append(d.Layers, v)
		}
	}
	return
}

func convertFromDockerApiTypeImageMetaData(s dockertypes.ImageMetadata) (d *ImageMetadata) {
	d = &ImageMetadata{
		LastTagTime: (*gogoprotobuftypes.Timestamp)(nil),
	}
	var err error
	d.LastTagTime, err = gogoprotobuftypes.TimestampProto(s.LastTagTime)
	if err != nil {
		glog.Warningf("Could not parse time: %v", err)
	}
	return
}

func ConvertFromDockerApiTypeImageSummary(s dockertypes.ImageSummary) (dst *ImageSummary) {
	dst = &ImageSummary{
		Containers:  s.Containers,
		Created:     s.Created,
		Id:          s.ID,
		Labels:      make(map[string]string),
		ParentId:    s.ParentID,
		RepoDigests: make([]string, 0),
		RepoTags:    make([]string, 0),
		SharedSize:  s.SharedSize,
		Size_:       s.Size,
		VirtualSize: s.VirtualSize,
	}
	for k, v := range s.Labels {
		dst.Labels[k] = v
	}
	for _, v := range s.RepoDigests {
		dst.RepoDigests = append(dst.RepoDigests, v)
	}
	for _, v := range s.RepoTags {
		dst.RepoTags = append(dst.RepoTags, v)
	}
	return
}

func ConvertFromDockerApiTypeImageDeleteResponseItem(s dockertypes.ImageDeleteResponseItem) (dst *ImageDeleteResponseItem) {
	dst = &ImageDeleteResponseItem{
		Deleted:  s.Deleted,
		Untagged: s.Untagged,
	}
	return
}

func ConvertFromDockerApiTypeImagesPruneReport(s dockertypes.ImagesPruneReport) (dst *ImagesPruneReport) {
	dst = &ImagesPruneReport{
		ImagesDeleted:  make([]*ImageDeleteResponseItem, len(s.ImagesDeleted)),
		SpaceReclaimed: s.SpaceReclaimed,
	}
	for _, v := range s.ImagesDeleted {
		dst.ImagesDeleted = append(dst.ImagesDeleted, &ImageDeleteResponseItem{
			Deleted:  v.Deleted,
			Untagged: v.Untagged,
		})
	}
	return
}

func (m *NetworkCreate) DeepCopyChecked() (*NetworkCreate, []error) {
	tgt := new(NetworkCreate)
	if m == nil {
		return tgt, []error{ErrNetworkCreateIsNil}
	}

	var errlst []error
	errs := make([]error, 0)

	tgt.CheckDuplicate = m.CheckDuplicate
	tgt.Driver = m.Driver
	tgt.Scope = m.Scope
	tgt.EnableIpv6 = m.EnableIpv6
	errlst = nil
	tgt.Ipam, errlst = m.Ipam.DeepCopyChecked()
	if len(errlst) != 0 {
		errs = append(errs, errlst...)
	}
	tgt.Internal = m.Internal
	tgt.Attachable = m.Attachable
	tgt.Ingress = m.Ingress
	tgt.ConfigOnly = m.ConfigOnly
	tgt.ConfigFrom, _ = m.ConfigFrom.DeepCopyChecked()
	tgt.Options = make(map[string]string)
	for k, v := range m.Options {
		tgt.Options[k] = v
	}
	tgt.Labels = make(map[string]string)
	for k, v := range m.Labels {
		tgt.Labels[k] = v
	}
	return tgt, errs
}

func (m *NetworkCreate) ExportIntoDockerApiType() (tgt dockertypes.NetworkCreate) {
	if m != nil {
		tgt = dockertypes.NetworkCreate{
			CheckDuplicate: m.CheckDuplicate,
			Driver:         m.Driver,
			Scope:          m.Scope,
			EnableIPv6:     m.EnableIpv6,
			IPAM:           new(dockernetworktypes.IPAM),
			Internal:       m.Internal,
			Attachable:     m.Attachable,
			Ingress:        m.Ingress,
			ConfigOnly:     m.ConfigOnly,
			ConfigFrom:     new(dockernetworktypes.ConfigReference),
			Options:        make(map[string]string),
			// Lables:         make(map[string]string),
		}

		ipam := m.Ipam.ExportIntoDockerApiTypes()
		tgt.IPAM = &ipam

		configfrom := m.ConfigFrom.ExportIntoDockerApiType()
		tgt.ConfigFrom = &configfrom

		for k, v := range m.Options {
			tgt.Options[k] = v
		}

		// for k, v := range m.Labels {
		// 	tgt.Lables[k] = v
		// }
	}
	return
}

func ConvertFromDockerApiTypeNetworkCreateResponse(resp dockertypes.NetworkCreateResponse) *NetworkCreateResponse {
	return &NetworkCreateResponse{
		Id:      resp.ID,
		Warning: resp.Warning,
	}
}

func (m *NetworkInspectOptions) DeepCopyChecked() *NetworkInspectOptions {
	tgt := new(NetworkInspectOptions)
	if m != nil {
		tgt.Scope = m.Scope
		tgt.Verbose = m.Verbose
	}
	return tgt
}

func (m *NetworkInspectOptions) ExportIntoDockerApiType() (tgt dockertypes.NetworkInspectOptions) {
	if m != nil {
		tgt.Scope = m.Scope
		tgt.Verbose = m.Verbose
	}
	return
}

func ConvertFromDockerApiTypeNetworkResource(resp dockertypes.NetworkResource) *NetworkResource {
	tgt := &NetworkResource{
		Name:       resp.Name,
		Id:         resp.ID,
		Created:    new(gogoprotobuftypes.Timestamp),
		Scope:      resp.Scope,
		Driver:     resp.Driver,
		EnableIpv6: resp.EnableIPv6,
		Ipam: &networkpb.IPAM{
			Driver:  resp.IPAM.Driver,
			Options: make(map[string]string),
			Config:  make([]*networkpb.IPAMConfig, 0),
		},
		Internal:   resp.Internal,
		Attachable: resp.Attachable,
		Ingress:    resp.Ingress,
		ConfigFrom: &networkpb.ConfigReference{
			Network: resp.ConfigFrom.Network,
		},
		ConfigOnly: resp.ConfigOnly,
		Containers: make(map[string]*EndpointResource),
		Options:    make(map[string]string),
		Labels:     make(map[string]string),
		Peers:      make([]*networkpb.PeerInfo, 0),
		Services:   make(map[string]*networkpb.ServiceInfo),
	}
	var err error
	tgt.Created, err = gogoprotobuftypes.TimestampProto(resp.Created)
	if err != nil {
		glog.Warningf("failed to convert Time to protobuf Timestamp: %v", err)
	}
	for k, v := range resp.IPAM.Options {
		tgt.Ipam.Options[k] = v
	}
	for _, v := range resp.IPAM.Config {
		ele := &networkpb.IPAMConfig{
			Subnet:     v.Subnet,
			IpRange:    v.IPRange,
			Gateway:    v.Gateway,
			AuxAddress: make(map[string]string),
		}
		for k1, v1 := range v.AuxAddress {
			ele.AuxAddress[k1] = v1
		}
		tgt.Ipam.Config = append(tgt.Ipam.Config, ele)
	}
	for k, v := range resp.Containers {
		ele := &EndpointResource{
			Name:        v.Name,
			EndpointId:  v.EndpointID,
			MacAddress:  v.MacAddress,
			Ipv4Address: v.IPv4Address,
			Ipv6Address: v.IPv6Address,
		}
		tgt.Containers[k] = ele
	}

	for k, v := range resp.Options {
		tgt.Options[k] = v
	}
	for k, v := range resp.Labels {
		tgt.Labels[k] = v
	}

	for _, v := range resp.Peers {
		ele := &networkpb.PeerInfo{
			Name: v.Name,
			Ip:   v.IP,
		}
		tgt.Peers = append(tgt.Peers, ele)
	}
	for k, v := range resp.Services {
		ele := &networkpb.ServiceInfo{
			Vip:          v.VIP,
			Ports:        make([]string, 0),
			LocalLbIndex: int32(v.LocalLBIndex),
			Tasks:        make([]*networkpb.Task, 0),
		}
		for _, v1 := range v.Ports {
			ele.Ports = append(ele.Ports, v1)
		}
		for _, v1 := range v.Tasks {
			ele1 := &networkpb.Task{
				Name:       v1.Name,
				EndpointId: v1.EndpointID,
				EndpointIp: v1.EndpointIP,
				Info:       make(map[string]string),
			}
			for k2, v2 := range v1.Info {
				ele1.Info[k2] = v2
			}
		}
		tgt.Services[k] = ele
	}
	return tgt
}

func (m *NetworkListOptions) DeepCopyChecked() *NetworkListOptions {
	obj := new(NetworkListOptions)
	if m != nil {
		obj.Filters = m.Filters.DeepCopyChecked()
	}
	return obj
}

func (m *NetworkListOptions) ExportIntoDockerApiType() (tgt dockertypes.NetworkListOptions) {
	if m != nil {
		tgt.Filters = m.Filters.ExportIntoDockerApiType()
	}
	return
}

func ConvertFromDockerApiTypeNetworksPruneReport(s dockertypes.NetworksPruneReport) (dst *NetworksPruneReport) {
	dst = &NetworksPruneReport{
		NetworksDeleted: make([]string, len(s.NetworksDeleted)),
	}
	for _, v := range s.NetworksDeleted {
		dst.NetworksDeleted = append(dst.NetworksDeleted, v)
	}
	return dst
}
