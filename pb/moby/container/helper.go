package container

import (
	"errors"
	"time"

	blkiodevtypes "github.com/docker/docker/api/types/blkiodev"
	containertypes "github.com/docker/docker/api/types/container"
	mounttypes "github.com/docker/docker/api/types/mount"
	strslicetypes "github.com/docker/docker/api/types/strslice"
	nattypes "github.com/docker/go-connections/nat"
	unitstypes "github.com/docker/go-units"
	gogoprotobuftypes "github.com/gogo/protobuf/types"
	"github.com/golang/glog"

	blkiodevpb "github.com/tangfeixiong/go-to-docker/pb/moby/blkiodev"
	mountpb "github.com/tangfeixiong/go-to-docker/pb/moby/mount"
	natpb "github.com/tangfeixiong/go-to-docker/pb/moby/nat"
	unitspb "github.com/tangfeixiong/go-to-docker/pb/moby/units"
)

var (
	ErrNilPointer        = errors.New("found nil pointer")
	ErrConfigIsNil       = errors.New("Config is nil")
	ErrImageNotSpecified = errors.New("Config's filed (Image) not specified")
	ErrHostConfigIsNil   = errors.New("HostConfig is nil")
)

func (m *Config) DeepCopyChecked() (*Config, []error) {
	if m == nil {
		return nil, []error{ErrConfigIsNil}
	}
	errorList := []error{}
	obj := &Config{
		Hostname:        m.Hostname,
		Domainname:      m.Domainname,
		User:            m.User,
		AttachStdin:     m.AttachStdin,
		AttachStdout:    m.AttachStdout,
		AttachStderr:    m.AttachStderr,
		ExposedPorts:    new(natpb.PortSet),
		Tty:             m.Tty,
		OpenStdin:       m.OpenStdin,
		StdinOnce:       m.StdinOnce,
		Env:             make([]string, 0),
		Cmd:             make([]string, 0),
		Healthcheck:     (*HealthConfig)(nil),
		ArgsEscaped:     m.ArgsEscaped,
		Image:           m.Image,
		Volumes:         make(map[string]*Config_VoidStruct),
		WorkingDir:      m.WorkingDir,
		Entrypoint:      make([]string, 0),
		NetworkDisabled: m.NetworkDisabled,
		MacAddress:      m.MacAddress,
		OnBuild:         make([]string, 0),
		Labels:          make(map[string]string),
		StopSignal:      m.StopSignal,
		StopTimeout:     (*Config_Int32Struct)(nil),
		Shell:           make([]string, 0),
	}
	if m.ExposedPorts != nil {
		obj.ExposedPorts.InternalMap = make(map[string]*natpb.PortSet_VoidStruct)
		for k, v := range m.ExposedPorts.InternalMap {
			obj.ExposedPorts.InternalMap[k] = v
		}
	}
	if len(m.Env) != 0 {
		obj.Env = append([]string{}, m.Env...)
	}
	if len(m.Cmd) != 0 {
		obj.Cmd = append([]string{}, m.Cmd...)
	}
	obj.Healthcheck, _ = m.Healthcheck.deepCopyChecked()
	if len(m.Image) == 0 {
		errorList = append(errorList, ErrImageNotSpecified)
	}
	for k, v := range m.Volumes {
		obj.Volumes[k] = v
	}
	if len(m.Entrypoint) != 0 {
		obj.Entrypoint = append([]string{}, m.Entrypoint...)
	}
	if len(m.OnBuild) != 0 {
		obj.OnBuild = append([]string{}, m.OnBuild...)
	}
	for k, v := range m.Labels {
		obj.Labels[k] = v
	}
	if v := m.StopTimeout; v != nil {
		obj.StopTimeout = &Config_Int32Struct{
			Value: v.Value,
		}
	}
	if len(m.Shell) != 0 {
		obj.Shell = append([]string{}, m.Shell...)
	}
	return obj, errorList
}

func (m *Config) ExportIntoDockerApiType() (tgt *containertypes.Config) {
	if m == nil {
		return
	}
	tgt = &containertypes.Config{
		Hostname:        m.Hostname,
		Domainname:      m.Domainname,
		User:            m.User,
		AttachStdin:     m.AttachStdin,
		AttachStdout:    m.AttachStdout,
		AttachStderr:    m.AttachStderr,
		ExposedPorts:    nattypes.PortSet{},
		Tty:             m.Tty,
		OpenStdin:       m.OpenStdin,
		StdinOnce:       m.StdinOnce,
		Env:             make([]string, 0),
		Cmd:             strslicetypes.StrSlice{},
		Healthcheck:     (*containertypes.HealthConfig)(nil),
		ArgsEscaped:     m.ArgsEscaped,
		Image:           m.Image,
		Volumes:         make(map[string]struct{}),
		WorkingDir:      m.WorkingDir,
		Entrypoint:      strslicetypes.StrSlice{},
		NetworkDisabled: m.NetworkDisabled,
		MacAddress:      m.MacAddress,
		OnBuild:         make([]string, 0),
		Labels:          make(map[string]string),
		StopSignal:      m.StopSignal,
		StopTimeout:     (*int)(nil),
		Shell:           strslicetypes.StrSlice{},
	}
	tgt.ExposedPorts = m.ExposedPorts.ExportIntoDockerApiType()
	for _, v := range m.Env {
		tgt.Env = append(tgt.Env, v)
	}
	for _, v := range m.Cmd {
		tgt.Cmd = append(tgt.Cmd, v)
	}
	tgt.Healthcheck = m.Healthcheck.exportIntoDockerApiType()
	for k, _ := range m.Volumes {
		tgt.Volumes[k] = struct{}{}
	}
	for _, v := range m.Entrypoint {
		tgt.Entrypoint = append(tgt.Entrypoint, v)
	}
	for _, v := range m.OnBuild {
		tgt.OnBuild = append(tgt.OnBuild, v)
	}
	for k, v := range m.Labels {
		tgt.Labels[k] = v
	}
	if v := m.StopTimeout; v != nil {
		num := int(v.Value)
		tgt.StopTimeout = &num
	}
	for _, v := range m.Shell {
		tgt.Shell = append(tgt.Shell, v)
	}
	return
}

func ConvertFromDockerApiTypeConfig(s *containertypes.Config) (dst *Config) {
	if s == nil {
		return
	}
	dst = &Config{
		Hostname:        s.Hostname,
		Domainname:      s.Domainname,
		User:            s.User,
		AttachStdin:     s.AttachStdin,
		AttachStdout:    s.AttachStdout,
		AttachStderr:    s.AttachStderr,
		ExposedPorts:    (*natpb.PortSet)(nil),
		Tty:             s.Tty,
		OpenStdin:       s.OpenStdin,
		StdinOnce:       s.StdinOnce,
		Env:             ([]string)(nil),
		Cmd:             ([]string)(nil),
		Healthcheck:     (*HealthConfig)(nil),
		ArgsEscaped:     s.ArgsEscaped,
		Image:           s.Image,
		Volumes:         (map[string]*Config_VoidStruct)(nil),
		WorkingDir:      s.WorkingDir,
		Entrypoint:      ([]string)(nil),
		NetworkDisabled: s.NetworkDisabled,
		MacAddress:      s.MacAddress,
		OnBuild:         ([]string)(nil),
		Labels:          (map[string]string)(nil),
		StopSignal:      s.StopSignal,
		StopTimeout:     (*Config_Int32Struct)(nil),
		Shell:           ([]string)(nil),
	}
	dst.ExposedPorts = natpb.ConvertFromDockerApiTypePortSet(&s.ExposedPorts)
	if s.Env != nil {
		dst.Env = make([]string, len(s.Env))
		for _, v := range s.Env {
			dst.Env = append(dst.Env, v)
		}
	}
	if s.Cmd != nil {
		dst.Cmd = make([]string, len(s.Cmd))
		for _, v := range s.Cmd {
			dst.Cmd = append(dst.Cmd, v)
		}
	}
	dst.Healthcheck = convertFromDockerApiTypeHealthConfig(s.Healthcheck)
	if s.Volumes != nil {
		dst.Volumes = make(map[string]*Config_VoidStruct)
		for k, _ := range s.Volumes {
			dst.Volumes[k] = &Config_VoidStruct{}
		}
	}
	if s.Entrypoint != nil {
		dst.Entrypoint = make([]string, len(s.Entrypoint))
		for _, v := range s.Entrypoint {
			dst.Entrypoint = append(dst.Entrypoint, v)
		}
	}
	if s.OnBuild != nil {
		dst.OnBuild = make([]string, len(s.OnBuild))
		for _, v := range s.OnBuild {
			dst.OnBuild = append(dst.OnBuild, v)
		}
	}
	if s.Labels != nil {
		s.Labels = make(map[string]string)
		for k, v := range s.Labels {
			dst.Labels[k] = v
		}
	}
	if v := s.StopTimeout; v != nil {
		dst.StopTimeout = &Config_Int32Struct{
			Value: int32(*v),
		}
	}
	if s.Shell != nil {
		dst.Shell = make([]string, len(s.Shell))
		for _, v := range s.Shell {
			dst.Shell = append(dst.Shell, v)
		}
	}
	return
}

func (m *HealthConfig) deepCopyChecked() (*HealthConfig, []error) {
	if m == nil {
		return nil, []error{ErrNilPointer}
	}
	obj := &HealthConfig{
		Test:        make([]string, 0),
		Interval:    (*gogoprotobuftypes.Duration)(nil),
		Timeout:     (*gogoprotobuftypes.Duration)(nil),
		StartPeriod: (*gogoprotobuftypes.Duration)(nil),
		Retries:     m.Retries,
	}
	if len(m.Test) != 0 {
		obj.Test = append([]string{}, m.Test...)
	}
	if v := m.Interval; v != nil {
		bytes, err := v.Marshal()
		if err != nil {
			glog.Warningln(err.Error())
			obj.Interval = &gogoprotobuftypes.Duration{
				Seconds: v.Seconds,
				Nanos:   v.Nanos,
			}
		} else {
			p := &gogoprotobuftypes.Duration{}
			err = p.Unmarshal(bytes)
			if err != nil {
				glog.Warningln(err.Error())
			} else {
				obj.Interval = p
			}
		}
	}
	if v := m.Timeout; v != nil {
		bytes, err := v.Marshal()
		if err != nil {
			glog.Warningln(err.Error())
			obj.Timeout = &gogoprotobuftypes.Duration{
				Seconds: v.Seconds,
				Nanos:   v.Nanos,
			}
		} else {
			p := &gogoprotobuftypes.Duration{}
			err = p.Unmarshal(bytes)
			if err != nil {
				glog.Warningln(err.Error())
			} else {
				obj.Timeout = p
			}
		}
	}
	if v := m.StartPeriod; v != nil {
		bytes, err := v.Marshal()
		if err != nil {
			glog.Warningln(err.Error())
			obj.StartPeriod = &gogoprotobuftypes.Duration{
				Seconds: v.Seconds,
				Nanos:   v.Nanos,
			}
		} else {
			p := &gogoprotobuftypes.Duration{}
			err = p.Unmarshal(bytes)
			if err != nil {
				glog.Warningln(err.Error())
			} else {
				obj.StartPeriod = p
			}
		}
	}
	return obj, []error{}
}

func (m *HealthConfig) exportIntoDockerApiType() (tgt *containertypes.HealthConfig) {
	if m == nil {
		return
	}
	tgt = &containertypes.HealthConfig{
		Test:        make([]string, 0),
		Interval:    time.Duration(0),
		Timeout:     time.Duration(0),
		StartPeriod: time.Duration(0),
		Retries:     int(m.Retries),
	}
	for _, v := range m.Test {
		tgt.Test = append(tgt.Test, v)
	}
	if v := m.Interval; v != nil {
		var err error
		tgt.Interval, err = gogoprotobuftypes.DurationFromProto(v)
		if err != nil {
			glog.Warningln(err.Error())
		}
	}
	if v := m.Timeout; v != nil {
		var err error
		tgt.Timeout, err = gogoprotobuftypes.DurationFromProto(v)
		if err != nil {
			glog.Warningln(err.Error())
		}
	}
	if v := m.StartPeriod; v != nil {
		var err error
		tgt.StartPeriod, err = gogoprotobuftypes.DurationFromProto(v)
		if err != nil {
			glog.Warningln(err.Error())
		}
	}
	return
}

func convertFromDockerApiTypeHealthConfig(s *containertypes.HealthConfig) (d *HealthConfig) {
	if s == nil {
		return
	}
	d = &HealthConfig{
		Test:        make([]string, 0),
		Interval:    gogoprotobuftypes.DurationProto(s.Interval),
		Timeout:     gogoprotobuftypes.DurationProto(s.Timeout),
		StartPeriod: gogoprotobuftypes.DurationProto(s.StartPeriod),
		Retries:     int32(s.Retries),
	}
	for _, v := range s.Test {
		d.Test = append(d.Test, v)
	}
	return
}

func (m *HostConfig) DeepCopyChecked() (*HostConfig, []error) {
	if m == nil {
		return nil, []error{ErrHostConfigIsNil}
	}
	errorList := []error{}
	obj := &HostConfig{
		Binds:           make([]string, 0),
		ContainerIdFile: m.ContainerIdFile,
		LogConfig:       new(LogConfig),
		NetworkMode:     m.NetworkMode,
		PortBindings:    new(natpb.PortMap),
		RestartPolicy:   new(RestartPolicy),
		AutoRemove:      m.AutoRemove,
		VolumeDriver:    m.VolumeDriver,
		VolumesFrom:     make([]string, 0),

		CapAdd:          make([]string, 0),
		CapDrop:         make([]string, 0),
		Dns:             make([]string, 0),
		DnsOptions:      make([]string, 0),
		DnsSearch:       make([]string, 0),
		ExtraHosts:      make([]string, 0),
		GroupAdd:        make([]string, 0),
		IpcMode:         m.IpcMode,
		Cgroup:          m.Cgroup,
		Links:           make([]string, 0),
		OomScoreAdj:     m.OomScoreAdj,
		PidMode:         m.PidMode,
		Privileged:      m.Privileged,
		PublishAllPorts: m.PublishAllPorts,
		ReadonlyRootfs:  m.ReadonlyRootfs,
		SecurityOpt:     make([]string, 0),
		StorageOpt:      make(map[string]string),
		Tmpfs:           make(map[string]string),
		UtsMode:         m.UtsMode,
		UsernsMode:      m.UsernsMode,
		ShmSize:         m.ShmSize,
		Sysctls:         make(map[string]string),
		Runtime:         m.Runtime,

		ConsoleSizeHeight: m.ConsoleSizeHeight,
		ConsoleSizeWidth:  m.ConsoleSizeWidth,
		Isolation:         m.Isolation,

		Resources: new(Resources),

		Mounts: make([]*mountpb.Mount, 0),

		MaskedPaths: make([]string, 0),

		ReadonlyPaths: make([]string, 0),

		Init: (*HostConfig_BoolStruct)(nil),
	}

	for _, v := range m.Binds {
		obj.Binds = append(obj.Binds, v)
	}
	if v := m.LogConfig; v != nil {
		obj.LogConfig, _ = v.deepCopyChecked()
	}
	if v := m.PortBindings; v != nil {
		obj.PortBindings, _ = v.DeepCopyChecked()
	}
	if v := m.RestartPolicy; v != nil {
		obj.RestartPolicy = &RestartPolicy{
			Name:              v.Name,
			MaximumRetryCount: v.MaximumRetryCount,
		}
	}
	for _, v := range m.VolumesFrom {
		obj.VolumesFrom = append(obj.VolumesFrom, v)
	}
	for _, v := range m.CapAdd {
		obj.CapAdd = append(obj.CapAdd, v)
	}
	for _, v := range m.CapDrop {
		obj.CapDrop = append(obj.CapDrop, v)
	}
	for _, v := range m.Dns {
		obj.Dns = append(obj.Dns, v)
	}
	for _, v := range m.DnsOptions {
		obj.DnsOptions = append(obj.DnsOptions, v)
	}
	for _, v := range m.DnsSearch {
		obj.DnsSearch = append(obj.DnsSearch, v)
	}
	for _, v := range m.ExtraHosts {
		obj.ExtraHosts = append(obj.ExtraHosts, v)
	}
	for _, v := range m.GroupAdd {
		obj.GroupAdd = append(obj.GroupAdd, v)
	}
	for _, v := range m.Links {
		obj.Links = append(obj.Links, v)
	}
	for _, v := range m.SecurityOpt {
		obj.SecurityOpt = append(obj.SecurityOpt, v)
	}
	for k, v := range m.StorageOpt {
		obj.StorageOpt[k] = v
	}
	for k, v := range m.Tmpfs {
		obj.Tmpfs[k] = v
	}
	for k, v := range m.Sysctls {
		obj.Sysctls[k] = v
	}
	if m.Resources != nil {
		var errs []error
		obj.Resources, errs = m.Resources.deepCopyChecked()
		if len(errs) != 0 {
			errorList = append(errorList, errs...)
		}
	}
	for _, v := range m.Mounts {
		ele, _ := v.DeepCopyChecked()
		obj.Mounts = append(obj.Mounts, ele)
	}
	for _, v := range m.MaskedPaths {
		obj.MaskedPaths = append(obj.MaskedPaths, v)
	}
	for _, v := range m.ReadonlyPaths {
		obj.ReadonlyPaths = append(obj.ReadonlyPaths, v)
	}
	if v := m.Init; v != nil {
		obj.Init = &HostConfig_BoolStruct{
			Value: v.Value,
		}
	}
	return m, errorList
}

func (m *HostConfig) ExportIntoDockerApiType() (tgt *containertypes.HostConfig) {
	if m == nil {
		return
	}
	tgt = &containertypes.HostConfig{
		Binds:           make([]string, 0),
		ContainerIDFile: m.ContainerIdFile,
		LogConfig:       containertypes.LogConfig{},
		NetworkMode:     containertypes.NetworkMode(m.NetworkMode),
		PortBindings:    make(nattypes.PortMap),
		RestartPolicy:   containertypes.RestartPolicy{},
		AutoRemove:      m.AutoRemove,
		VolumeDriver:    m.VolumeDriver,
		VolumesFrom:     make([]string, 0),

		CapAdd:          make(strslicetypes.StrSlice, 0),
		CapDrop:         make(strslicetypes.StrSlice, 0),
		DNS:             make([]string, 0),
		DNSOptions:      make([]string, 0),
		DNSSearch:       make([]string, 0),
		ExtraHosts:      make([]string, 0),
		GroupAdd:        make([]string, 0),
		IpcMode:         containertypes.IpcMode(m.IpcMode),
		Cgroup:          containertypes.CgroupSpec(m.Cgroup),
		Links:           make([]string, 0),
		OomScoreAdj:     int(m.OomScoreAdj),
		PidMode:         containertypes.PidMode(m.PidMode),
		Privileged:      m.Privileged,
		PublishAllPorts: m.PublishAllPorts,
		ReadonlyRootfs:  m.ReadonlyRootfs,
		SecurityOpt:     make([]string, 0),
		StorageOpt:      make(map[string]string),
		Tmpfs:           make(map[string]string),
		UTSMode:         containertypes.UTSMode(m.UtsMode),
		UsernsMode:      containertypes.UsernsMode(m.UsernsMode),
		ShmSize:         m.ShmSize,
		Sysctls:         make(map[string]string),
		Runtime:         m.Runtime,

		ConsoleSize: [2]uint{uint(m.ConsoleSizeHeight), uint(m.ConsoleSizeWidth)},
		Isolation:   containertypes.Isolation(m.Isolation),

		Resources: containertypes.Resources{},

		Mounts: make([]mounttypes.Mount, 0),

		// MaskedPaths: make([]string, 0),

		// ReadonlyPaths: make([]string, 0),

		Init: (*bool)(nil),
	}
	for _, v := range m.Binds {
		tgt.Binds = append(tgt.Binds, v)
	}
	if m.LogConfig != nil {
		tgt.LogConfig = m.LogConfig.exportIntoDockerApiType()
	}
	if m.PortBindings != nil {
		tgt.PortBindings = m.PortBindings.ExportIntoDockerApiType()
	}
	if v := m.RestartPolicy; v != nil {
		tgt.RestartPolicy = containertypes.RestartPolicy{
			Name:              v.Name,
			MaximumRetryCount: int(v.MaximumRetryCount),
		}
	}
	for _, v := range m.VolumesFrom {
		tgt.VolumesFrom = append(tgt.VolumesFrom, v)
	}
	for _, v := range m.CapAdd {
		tgt.CapAdd = append(tgt.CapAdd, v)
	}
	for _, v := range m.CapDrop {
		tgt.CapDrop = append(tgt.CapDrop, v)
	}
	for _, v := range m.Dns {
		tgt.DNS = append(tgt.DNS, v)
	}
	for _, v := range m.DnsOptions {
		tgt.DNSOptions = append(tgt.DNSOptions, v)
	}
	for _, v := range m.DnsSearch {
		tgt.DNSSearch = append(tgt.DNSSearch, v)
	}
	for _, v := range m.ExtraHosts {
		tgt.ExtraHosts = append(tgt.ExtraHosts, v)
	}
	for _, v := range m.GroupAdd {
		tgt.GroupAdd = append(tgt.GroupAdd, v)
	}
	for _, v := range m.Links {
		tgt.Links = append(tgt.Links, v)
	}
	for _, v := range m.SecurityOpt {
		tgt.SecurityOpt = append(tgt.SecurityOpt, v)
	}
	for k, v := range m.StorageOpt {
		tgt.StorageOpt[k] = v
	}
	for k, v := range m.Tmpfs {
		tgt.Tmpfs[k] = v
	}
	for k, v := range m.Sysctls {
		tgt.Sysctls[k] = v
	}
	tgt.Resources = m.Resources.exportIntoDockerApiType()
	for _, v := range m.Mounts {
		tgt.Mounts = append(tgt.Mounts, v.ExportIntoDockerApiType())
	}
	// for _, v := range m.MaskedPaths {
	//	  tgt.MaskedPaths = append(tgt.MaskedPaths, v)
	// }
	// for _, v := range m.ReadonlyPaths {
	// 	  tgt.ReadonlyPaths = append(tgt.ReadonlyPaths, v)
	// }
	if v := m.Init; v != nil {
		f := v.Value
		tgt.Init = &f
	}
	return
}

func ConvertFromDockerApiTypeHostConfig(s *containertypes.HostConfig) (d *HostConfig) {
	if s == nil {
		return
	}
	d = &HostConfig{
		Binds:           make([]string, len(s.Binds)),
		ContainerIdFile: s.ContainerIDFile,
		LogConfig:       new(LogConfig),
		NetworkMode:     string(s.NetworkMode),
		PortBindings:    (*natpb.PortMap)(nil),
		RestartPolicy:   new(RestartPolicy),
		AutoRemove:      s.AutoRemove,
		VolumeDriver:    s.VolumeDriver,
		VolumesFrom:     make([]string, len(s.VolumesFrom)),

		CapAdd:          make([]string, len(s.CapAdd)),
		CapDrop:         make([]string, len(s.CapDrop)),
		Dns:             make([]string, len(s.DNS)),
		DnsOptions:      make([]string, len(s.DNSOptions)),
		DnsSearch:       make([]string, len(s.DNSSearch)),
		ExtraHosts:      make([]string, len(s.ExtraHosts)),
		GroupAdd:        make([]string, len(s.GroupAdd)),
		IpcMode:         string(s.IpcMode),
		Cgroup:          string(s.Cgroup),
		Links:           make([]string, len(s.Links)),
		OomScoreAdj:     int32(s.OomScoreAdj),
		PidMode:         string(s.PidMode),
		Privileged:      s.Privileged,
		PublishAllPorts: s.PublishAllPorts,
		ReadonlyRootfs:  s.ReadonlyRootfs,
		SecurityOpt:     make([]string, len(s.SecurityOpt)),
		StorageOpt:      make(map[string]string),
		Tmpfs:           make(map[string]string),
		UtsMode:         string(s.UTSMode),
		UsernsMode:      string(s.UsernsMode),
		ShmSize:         s.ShmSize,
		Sysctls:         make(map[string]string),
		Runtime:         s.Runtime,

		ConsoleSizeHeight: uint32(s.ConsoleSize[0]),
		ConsoleSizeWidth:  uint32(s.ConsoleSize[1]),
		Isolation:         string(s.Isolation),

		Resources: (*Resources)(nil),

		Mounts: make([]*mountpb.Mount, len(s.Mounts)),

		MaskedPaths: make([]string, 0),

		ReadonlyPaths: make([]string, 0),

		Init: (*HostConfig_BoolStruct)(nil),
	}
	for _, v := range s.Binds {
		d.Binds = append(d.Binds, v)
	}
	d.LogConfig = convertFromDockerApiTypeLogConfig(s.LogConfig)
	d.PortBindings = natpb.ConvertFromDockerApiTypePortMap(s.PortBindings)
	d.RestartPolicy = &RestartPolicy{
		Name:              s.RestartPolicy.Name,
		MaximumRetryCount: int32(s.RestartPolicy.MaximumRetryCount),
	}
	for _, v := range s.VolumesFrom {
		d.VolumesFrom = append(d.VolumesFrom, v)
	}
	for _, v := range s.CapAdd {
		d.CapAdd = append(d.CapAdd, v)
	}
	for _, v := range s.CapDrop {
		d.CapDrop = append(d.CapDrop, v)
	}
	for _, v := range s.DNS {
		d.Dns = append(d.Dns, v)
	}
	for _, v := range s.DNSOptions {
		d.DnsOptions = append(d.DnsOptions, v)
	}
	for _, v := range s.DNSSearch {
		d.DnsSearch = append(d.DnsSearch, v)
	}
	for _, v := range s.ExtraHosts {
		d.ExtraHosts = append(d.ExtraHosts, v)
	}
	for _, v := range s.GroupAdd {
		d.GroupAdd = append(d.GroupAdd, v)
	}
	for _, v := range s.Links {
		d.Links = append(d.Links, v)
	}
	for _, v := range s.SecurityOpt {
		d.SecurityOpt = append(d.SecurityOpt, v)
	}
	for k, v := range s.StorageOpt {
		d.StorageOpt[k] = v
	}
	for k, v := range s.Tmpfs {
		d.Tmpfs[k] = v
	}
	for k, v := range s.Sysctls {
		d.Sysctls[k] = v
	}
	d.Resources = convertFromDockerApiTypeResources(s.Resources)
	for _, v := range s.Mounts {
		d.Mounts = append(d.Mounts, mountpb.ConvertFromDockerApiTypeMount(v))
	}
	// for _, v := range s.MaskedPaths {
	// 	d.MaskedPaths = append(d.MaskedPaths, v)
	// }
	// for _, v := range s.ReadonlyPaths {
	// 	d.ReadonlyPaths = append(d.ReadonlyPaths, v)
	// }
	if v := s.Init; v != nil {
		d.Init = &HostConfig_BoolStruct{
			Value: *v,
		}
	}
	return
}

func (m *LogConfig) deepCopyChecked() (*LogConfig, []error) {
	var obj *LogConfig = nil
	var errorList []error = []error{}
	if m == nil {
		errorList = append(errorList, ErrNilPointer)
	} else {
		obj = &LogConfig{
			Type:   m.Type,
			Config: make(map[string]string),
		}
		for k, v := range m.Config {
			obj.Config[k] = v
		}
	}
	return obj, errorList
}

func (m *LogConfig) exportIntoDockerApiType() (tgt containertypes.LogConfig) {
	if m != nil {
		tgt = containertypes.LogConfig{
			Type:   m.Type,
			Config: make(map[string]string),
		}
		for k, v := range m.Config {
			tgt.Config[k] = v
		}
	}
	return
}

func convertFromDockerApiTypeLogConfig(s containertypes.LogConfig) (d *LogConfig) {
	d = &LogConfig{
		Type:   s.Type,
		Config: make(map[string]string),
	}
	for k, v := range s.Config {
		d.Config[k] = v
	}
	return
}

func (m *Resources) deepCopyChecked() (*Resources, []error) {
	if m == nil {
		return nil, []error{}
	}
	errorList := []error{}
	obj := &Resources{
		CpuShares: m.CpuShares,
		Memory:    m.Memory,
		NanoCpus:  m.NanoCpus,

		CgroupParent:         m.CgroupParent,
		BlkioWeight:          m.BlkioWeight,
		BlkioWeightDevice:    make([]*blkiodevpb.WeightDevice, 0),
		BlkioDeviceReadBps:   make([]*blkiodevpb.ThrottleDevice, 0),
		BlkioDeviceWriteBps:  make([]*blkiodevpb.ThrottleDevice, 0),
		BlkioDeviceReadIops:  make([]*blkiodevpb.ThrottleDevice, 0),
		BlkioDeviceWriteIops: make([]*blkiodevpb.ThrottleDevice, 0),
		CpuPeriod:            m.CpuPeriod,
		CpuQuota:             m.CpuQuota,
		CpuRealtimePeriod:    m.CpuRealtimePeriod,
		CpuRealtimeRuntime:   m.CpuRealtimeRuntime,
		CpusetCpus:           m.CpusetCpus,
		CpusetMems:           m.CpusetMems,
		Devices:              make([]*DeviceMapping, 0),
		DeviceCgroupRules:    make([]string, 0),
		DiskQuota:            m.DiskQuota,
		KernelMemory:         m.KernelMemory,
		MemoryReservation:    m.MemoryReservation,
		MemorySwap:           m.MemorySwap,
		MemorySwappiness:     (*Resources_Int64Struct)(nil),
		OomKillDisable:       (*Resources_BoolStruct)(nil),
		PidsLimit:            m.PidsLimit,
		Ulimits:              make([]*unitspb.Ulimit, 0),

		CpuCount:           m.CpuCount,
		CpuPercent:         m.CpuPercent,
		IoMaximumIops:      m.IoMaximumIops,
		IoMaximumBandwidth: m.IoMaximumBandwidth,
	}
	for _, v := range m.BlkioWeightDevice {
		var ele *blkiodevpb.WeightDevice = nil
		if v != nil {
			ele = &blkiodevpb.WeightDevice{
				Path:   v.Path,
				Weight: v.Weight,
			}
		}
		obj.BlkioWeightDevice = append(obj.BlkioWeightDevice, ele)
	}
	for _, v := range m.BlkioDeviceReadBps {
		var ele *blkiodevpb.ThrottleDevice = nil
		if v != nil {
			ele = &blkiodevpb.ThrottleDevice{
				Path: v.Path,
				Rate: v.Rate,
			}
		}
		obj.BlkioDeviceReadBps = append(obj.BlkioDeviceReadBps, ele)
	}
	for _, v := range m.BlkioDeviceWriteBps {
		var ele *blkiodevpb.ThrottleDevice = nil
		if v != nil {
			ele = &blkiodevpb.ThrottleDevice{
				Path: v.Path,
				Rate: v.Rate,
			}
		}
		obj.BlkioDeviceWriteBps = append(obj.BlkioDeviceWriteBps, ele)
	}
	for _, v := range m.BlkioDeviceReadIops {
		var ele *blkiodevpb.ThrottleDevice = nil
		if v != nil {
			ele = &blkiodevpb.ThrottleDevice{
				Path: v.Path,
				Rate: v.Rate,
			}
		}
		obj.BlkioDeviceReadIops = append(obj.BlkioDeviceReadIops, ele)
	}
	for _, v := range m.BlkioDeviceWriteIops {
		var ele *blkiodevpb.ThrottleDevice = nil
		if v != nil {
			ele = &blkiodevpb.ThrottleDevice{
				Path: v.Path,
				Rate: v.Rate,
			}
		}
		obj.BlkioDeviceWriteIops = append(obj.BlkioDeviceWriteIops, ele)
	}
	for _, v := range m.Devices {
		if v != nil {
			obj.Devices = append(obj.Devices, &DeviceMapping{
				PathOnHost:        v.PathOnHost,
				PathInContainer:   v.PathInContainer,
				CgroupPermissions: v.CgroupPermissions,
			})
		}
	}
	for _, v := range m.DeviceCgroupRules {
		obj.DeviceCgroupRules = append(obj.DeviceCgroupRules, v)
	}
	if v := m.MemorySwappiness; v != nil {
		obj.MemorySwappiness = &Resources_Int64Struct{
			Value: v.Value,
		}
	}
	if v := m.OomKillDisable; v != nil {
		obj.OomKillDisable = &Resources_BoolStruct{
			Value: v.Value,
		}
	}
	for _, v := range m.Ulimits {
		if v != nil {
			obj.Ulimits = append(obj.Ulimits, &unitspb.Ulimit{
				Name: v.Name,
				Hard: v.Hard,
				Soft: v.Soft,
			})
		}
	}
	return obj, errorList
}

func (m *Resources) exportIntoDockerApiType() (tgt containertypes.Resources) {
	if m != nil {
		tgt := containertypes.Resources{
			CPUShares: m.CpuShares,
			Memory:    m.Memory,
			NanoCPUs:  m.NanoCpus,

			CgroupParent:         m.CgroupParent,
			BlkioWeight:          uint16(m.BlkioWeight),
			BlkioWeightDevice:    make([]*blkiodevtypes.WeightDevice, 0),
			BlkioDeviceReadBps:   make([]*blkiodevtypes.ThrottleDevice, 0),
			BlkioDeviceWriteBps:  make([]*blkiodevtypes.ThrottleDevice, 0),
			BlkioDeviceReadIOps:  make([]*blkiodevtypes.ThrottleDevice, 0),
			BlkioDeviceWriteIOps: make([]*blkiodevtypes.ThrottleDevice, 0),
			CPUPeriod:            m.CpuPeriod,
			CPUQuota:             m.CpuQuota,
			CPURealtimePeriod:    m.CpuRealtimePeriod,
			CPURealtimeRuntime:   m.CpuRealtimeRuntime,
			CpusetCpus:           m.CpusetCpus,
			CpusetMems:           m.CpusetMems,
			Devices:              make([]containertypes.DeviceMapping, 0),
			DeviceCgroupRules:    make([]string, 0),
			DiskQuota:            m.DiskQuota,
			KernelMemory:         m.KernelMemory,
			MemoryReservation:    m.MemoryReservation,
			MemorySwap:           m.MemorySwap,
			MemorySwappiness:     (*int64)(nil),
			OomKillDisable:       (*bool)(nil),
			PidsLimit:            m.PidsLimit,
			Ulimits:              make([]*unitstypes.Ulimit, 0),

			CPUCount:           m.CpuCount,
			CPUPercent:         m.CpuPercent,
			IOMaximumIOps:      m.IoMaximumIops,
			IOMaximumBandwidth: m.IoMaximumBandwidth,
		}
		for _, v := range m.BlkioWeightDevice {
			if v != nil {
				tgt.BlkioWeightDevice = append(tgt.BlkioWeightDevice, &blkiodevtypes.WeightDevice{
					Path:   v.Path,
					Weight: uint16(v.Weight),
				})
			}
		}
		for _, v := range m.BlkioDeviceReadBps {
			if v != nil {
				tgt.BlkioDeviceReadBps = append(tgt.BlkioDeviceReadBps, &blkiodevtypes.ThrottleDevice{
					Path: v.Path,
					Rate: v.Rate,
				})
			}
		}
		for _, v := range m.BlkioDeviceWriteBps {
			if v != nil {
				tgt.BlkioDeviceWriteBps = append(tgt.BlkioDeviceWriteBps, &blkiodevtypes.ThrottleDevice{
					Path: v.Path,
					Rate: v.Rate,
				})
			}
		}
		for _, v := range m.BlkioDeviceReadIops {
			if v != nil {
				tgt.BlkioDeviceReadIOps = append(tgt.BlkioDeviceReadIOps, &blkiodevtypes.ThrottleDevice{
					Path: v.Path,
					Rate: v.Rate,
				})
			}
		}
		for _, v := range m.BlkioDeviceWriteIops {
			if v != nil {
				tgt.BlkioDeviceWriteIOps = append(tgt.BlkioDeviceWriteIOps, &blkiodevtypes.ThrottleDevice{
					Path: v.Path,
					Rate: v.Rate,
				})
			}
		}
		for _, v := range m.Devices {
			if v != nil {
				tgt.Devices = append(tgt.Devices, containertypes.DeviceMapping{
					PathOnHost:        v.PathOnHost,
					PathInContainer:   v.PathInContainer,
					CgroupPermissions: v.CgroupPermissions,
				})
			}
		}
		for _, v := range m.DeviceCgroupRules {
			tgt.DeviceCgroupRules = append(tgt.DeviceCgroupRules, v)
		}
		if v := m.MemorySwappiness; v != nil {
			num := v.Value
			tgt.MemorySwappiness = &num
		}
		if v := m.OomKillDisable; v != nil {
			f := v.Value
			tgt.OomKillDisable = &f
		}
		for _, v := range m.Ulimits {
			if v != nil {
				tgt.Ulimits = append(tgt.Ulimits, &unitstypes.Ulimit{
					Name: v.Name,
					Hard: v.Hard,
					Soft: v.Soft,
				})
			}
		}
	}
	return
}

func convertFromDockerApiTypeResources(s containertypes.Resources) (d *Resources) {
	d = &Resources{
		CpuShares: s.CPUShares,
		Memory:    s.Memory,
		NanoCpus:  s.NanoCPUs,

		CgroupParent:         s.CgroupParent,
		BlkioWeight:          int32(s.BlkioWeight),
		BlkioWeightDevice:    ([]*blkiodevpb.WeightDevice)(nil),
		BlkioDeviceReadBps:   ([]*blkiodevpb.ThrottleDevice)(nil),
		BlkioDeviceWriteBps:  ([]*blkiodevpb.ThrottleDevice)(nil),
		BlkioDeviceReadIops:  ([]*blkiodevpb.ThrottleDevice)(nil),
		BlkioDeviceWriteIops: ([]*blkiodevpb.ThrottleDevice)(nil),
		CpuPeriod:            s.CPUPeriod,
		CpuQuota:             s.CPUQuota,
		CpuRealtimePeriod:    s.CPURealtimePeriod,
		CpuRealtimeRuntime:   s.CPURealtimeRuntime,
		CpusetCpus:           s.CpusetCpus,
		CpusetMems:           s.CpusetMems,
		Devices:              make([]*DeviceMapping, len(s.Devices)),
		DeviceCgroupRules:    make([]string, len(s.Devices)),
		DiskQuota:            s.DiskQuota,
		KernelMemory:         s.KernelMemory,
		MemoryReservation:    s.MemoryReservation,
		MemorySwap:           s.MemorySwap,
		MemorySwappiness:     (*Resources_Int64Struct)(nil),
		OomKillDisable:       (*Resources_BoolStruct)(nil),
		PidsLimit:            s.PidsLimit,
		Ulimits:              make([]*unitspb.Ulimit, 0),

		CpuCount:           s.CPUCount,
		CpuPercent:         s.CPUPercent,
		IoMaximumIops:      s.IOMaximumIOps,
		IoMaximumBandwidth: s.IOMaximumBandwidth,
	}
	for _, v := range s.BlkioWeightDevice {
		var ele *blkiodevpb.WeightDevice = nil
		if v != nil {
			ele = &blkiodevpb.WeightDevice{
				Path:   v.Path,
				Weight: int32(v.Weight),
			}
		}
		d.BlkioWeightDevice = append(d.BlkioWeightDevice, ele)
	}
	for _, v := range s.BlkioDeviceReadBps {
		var ele *blkiodevpb.ThrottleDevice = nil
		if v != nil {
			ele = &blkiodevpb.ThrottleDevice{
				Path: v.Path,
				Rate: v.Rate,
			}
		}
		d.BlkioDeviceReadBps = append(d.BlkioDeviceReadBps, ele)
	}
	for _, v := range s.BlkioDeviceWriteBps {
		var ele *blkiodevpb.ThrottleDevice = nil
		if v != nil {
			ele = &blkiodevpb.ThrottleDevice{
				Path: v.Path,
				Rate: v.Rate,
			}
		}
		d.BlkioDeviceWriteBps = append(d.BlkioDeviceWriteBps, ele)
	}
	for _, v := range s.BlkioDeviceReadIOps {
		var ele *blkiodevpb.ThrottleDevice = nil
		if v != nil {
			ele = &blkiodevpb.ThrottleDevice{
				Path: v.Path,
				Rate: v.Rate,
			}
		}
		d.BlkioDeviceReadIops = append(d.BlkioDeviceReadIops, ele)
	}
	for _, v := range s.BlkioDeviceWriteIOps {
		var ele *blkiodevpb.ThrottleDevice = nil
		if v != nil {
			ele = &blkiodevpb.ThrottleDevice{
				Path: v.Path,
				Rate: v.Rate,
			}
		}
		d.BlkioDeviceWriteIops = append(d.BlkioDeviceWriteIops, ele)
	}
	for _, v := range s.Devices {
		ele := &DeviceMapping{
			PathOnHost:        v.PathOnHost,
			PathInContainer:   v.PathInContainer,
			CgroupPermissions: v.CgroupPermissions,
		}
		d.Devices = append(d.Devices, ele)
	}
	for _, v := range s.DeviceCgroupRules {
		d.DeviceCgroupRules = append(d.DeviceCgroupRules, v)
	}
	if v := s.MemorySwappiness; v != nil {
		d.MemorySwappiness = &Resources_Int64Struct{
			Value: *v,
		}
	}
	if v := s.OomKillDisable; v != nil {
		d.OomKillDisable = &Resources_BoolStruct{
			Value: *v,
		}
	}
	for _, v := range s.Ulimits {
		var ele *unitspb.Ulimit = nil
		if v != nil {
			ele = &unitspb.Ulimit{
				Name: v.Name,
				Hard: v.Hard,
				Soft: v.Soft,
			}
		}
		d.Ulimits = append(d.Ulimits, ele)
	}
	return
}

func ConvertFromDockerApiTypeContainerCreateCreatedBody(s *containertypes.ContainerCreateCreatedBody) (dst *ContainerCreateCreatedBody) {
	if s != nil {
		dst = &ContainerCreateCreatedBody{
			Id:       s.ID,
			Warnings: make([]string, len(s.Warnings)),
		}
		for _, v := range s.Warnings {
			dst.Warnings = append(dst.Warnings, v)
		}
	}
	return
}
