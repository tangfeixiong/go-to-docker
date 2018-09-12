package network

import (
	"errors"

	networktypes "github.com/docker/docker/api/types/network"
)

const (
	Driver_Filter string = "driver"
	Type_Filter   string = "type"
	Name_Filter   string = "name"
	Id_Filter     string = "id"
	Label_Filter  string = "label"
	Scope_Filter  string = "scope"
)

var (
	ErrNilPointer            = errors.New("found nil pointer")
	ErrIpamIsNil             = errors.New("IPAM not specified")
	ErrIPAMConfigListIsEmpty = errors.New("IPAMConfig list not specified")
	ErrNetworkingConfigIsNil = errors.New("NetworkingConfig is nil")
)

func (m *NetworkingConfig) DeepCopyChecked() (*NetworkingConfig, []error) {
	if m == nil {
		return nil, []error{ErrNetworkingConfigIsNil}
	}
	errorList := []error{}
	tgt := &NetworkingConfig{
		EndpointsConfig: make(map[string]*EndpointSettings),
	}
	for k, v := range m.EndpointsConfig {
		ele, _ := v.deepCopyChecked()
		tgt.EndpointsConfig[k] = ele
	}
	return tgt, errorList
}

func (m *NetworkingConfig) ExportIntoDockerApiType() (tgt *networktypes.NetworkingConfig) {
	if m != nil {
		tgt = &networktypes.NetworkingConfig{
			EndpointsConfig: make(map[string]*networktypes.EndpointSettings),
		}
		for k, v := range m.EndpointsConfig {
			if v != nil {
				ele := &networktypes.EndpointSettings{
					IPAMConfig:          new(networktypes.EndpointIPAMConfig),
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
					DriverOpts:          make(map[string]string),
				}
				if v.IpamConfig != nil {
					ele.IPAMConfig.IPv4Address = v.IpamConfig.Ipv4Address
					ele.IPAMConfig.IPv6Address = v.IpamConfig.Ipv6Address
					ele.IPAMConfig.LinkLocalIPs = make([]string, 0)
					for _, v := range v.IpamConfig.LinkLocalIps {
						ele.IPAMConfig.LinkLocalIPs = append(ele.IPAMConfig.LinkLocalIPs, v)
					}
				}
				for _, v := range v.Links {
					ele.Links = append(ele.Links, v)
				}
				for _, v := range v.Aliases {
					ele.Aliases = append(ele.Aliases, v)
				}
				for k, v := range v.DriverOpts {
					ele.DriverOpts[k] = v
				}
				tgt.EndpointsConfig[k] = ele
			} else {
				tgt.EndpointsConfig[k] = new(networktypes.EndpointSettings)
			}
		}
	}
	return
}

func (m *EndpointSettings) deepCopyChecked() (*EndpointSettings, []error) {
	if m == nil {
		return new(EndpointSettings), []error{ErrNilPointer}
	}
	tgt := &EndpointSettings{
		Links:               make([]string, 0),
		Aliases:             make([]string, 0),
		NetworkId:           m.NetworkId,
		EndpointId:          m.EndpointId,
		Gateway:             m.Gateway,
		IpAddress:           m.IpAddress,
		IpPrefixLen:         m.IpPrefixLen,
		Ipv6Gateway:         m.Ipv6Gateway,
		GlobalIpv6Address:   m.GlobalIpv6Address,
		GlobalIpv6PrefixLen: m.GlobalIpv6PrefixLen,
		MacAddress:          m.MacAddress,
	}
	if m.IpamConfig != nil {
		tgt.IpamConfig = &EndpointIPAMConfig{
			Ipv4Address: m.IpamConfig.Ipv4Address,
			Ipv6Address: m.IpamConfig.Ipv6Address,
		}
		if len(m.IpamConfig.LinkLocalIps) != 0 {
			tgt.IpamConfig.LinkLocalIps = make([]string, len(m.IpamConfig.LinkLocalIps))
			for _, v := range m.IpamConfig.LinkLocalIps {
				tgt.IpamConfig.LinkLocalIps = append(tgt.IpamConfig.LinkLocalIps, v)
			}
		}
	}
	if len(m.Links) != 0 {
		tgt.Links = make([]string, len(m.Links))
		for _, v := range m.Links {
			tgt.Links = append(tgt.Links, v)
		}
	}
	if len(m.Aliases) != 0 {
		tgt.Aliases = make([]string, len(m.Aliases))
		for _, v := range m.Aliases {
			tgt.Aliases = append(tgt.Aliases, v)
		}
	}
	if len(m.DriverOpts) != 0 {
		tgt.DriverOpts = make(map[string]string)
		for k, v := range m.DriverOpts {
			tgt.DriverOpts[k] = v
		}
	}
	return tgt, []error{}
}

func (m *IPAM) DeepCopyChecked() (*IPAM, []error) {
	tgt := &IPAM{
		Options: make(map[string]string),
		Config:  make([]*IPAMConfig, 0),
	}
	if m == nil {
		return tgt, []error{ErrIpamIsNil}
	}
	tgt.Driver = m.Driver
	tgt.Options = m.Options
	if len(m.Config) == 0 {
		return tgt, []error{ErrIPAMConfigListIsEmpty}
	}
	for _, v := range m.Config {
		ele, _ := v.deepCopyChecked()
		tgt.Config = append(m.Config, ele)
	}
	return tgt, []error{}
}

func (m *IPAM) ExportIntoDockerApiTypes() (tgt networktypes.IPAM) {
	if nil != m {
		tgt.Driver = m.Driver
		tgt.Options = make(map[string]string)
		tgt.Config = make([]networktypes.IPAMConfig, 0)

		for k, v := range m.Options {
			tgt.Options[k] = v
		}
		for i := 0; i < len(m.Config); i++ {
			config := m.Config[i]
			tgt.Config = append(tgt.Config, config.exportIntoDockerApiType())
		}
	}
	return
}

func (m *IPAMConfig) deepCopyChecked() (*IPAMConfig, []error) {
	tgt := &IPAMConfig{
		AuxAddress: make(map[string]string),
	}
	tgt.Subnet = m.Subnet
	tgt.IpRange = m.IpRange
	tgt.Gateway = m.Gateway
	for k, v := range m.AuxAddress {
		tgt.AuxAddress[k] = v
	}
	return tgt, []error{}
}

func (m *IPAMConfig) exportIntoDockerApiType() (tgt networktypes.IPAMConfig) {
	if nil != m {
		tgt = networktypes.IPAMConfig{
			Subnet:     m.Subnet,
			IPRange:    m.IpRange,
			Gateway:    m.Gateway,
			AuxAddress: make(map[string]string),
		}
		for k, v := range m.AuxAddress {
			tgt.AuxAddress[k] = v
		}
	}
	return
}

func (m *ConfigReference) DeepCopyChecked() (*ConfigReference, []error) {
	tgt := new(ConfigReference)
	if nil == m {
		return tgt, []error{ErrNilPointer}
	}
	tgt.Network = m.Network
	return tgt, []error{}
}

func (m *ConfigReference) ExportIntoDockerApiType() (tgt networktypes.ConfigReference) {
	if nil != m {
		tgt.Network = m.Network
	}
	return
}
