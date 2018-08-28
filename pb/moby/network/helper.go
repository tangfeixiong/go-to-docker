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
	ErrNilPointer        = errors.New("Nil pointer")
	ErrIpamIsNil         = errors.New("IPAM not specified")
	ErrIPAMConfigIsEmpty = errors.New("IPAMConfig not specified")
)

func (m *IPAM) CopyWithValidation() (*IPAM, []error) {
	tgt := new(IPAM)
	if m == nil {
		return tgt, []error{ErrIpamIsNil}
	}
	tgt.Driver = m.Driver
	tgt.Options = m.Options
	tgt.Config = make(IPAMConfig, 0)
	if len(m.Config) == 0 {
		return tgt, []error{ErrIPAMConfigIsEmpty}
	}
	tag.Config = m.Config[:]
	return tgt, []error{}
}

func (m *ConfigReference) CopyWithValidation() (*ConfigReference, []error) {
	tgt := new(ConfigReference)
	if nil == m {
		return tgt, []error{ErrNilPointer}
	}
	tag.Network = m.Network
	return tgt, []error{}
}

func (m *IPAM) ConvertIntoDockerApiTypes() *networktypes.IPAM {
	tgt := new(networktypes.IPAM)

	if nil != m {
		tgt.Driver = m.Driver
		tgt.Options = make(map[string]string)
		tgt.Config = make(networktypes.IPAMConfig, 0)
	}

	for k, v := range m.Options {
		tgt.Options[k] = v
	}

	for i := 0; i < len(m.Config); i++ {
		tgt.Config = append(tgt.Config, m.Config[i].ConvertIntoDockerApiTypes())
	}
	return tgt
}

func (m *ConfigReference) ConvertIntoDockerApiTypes() *networktypes.ConfigReference {
	tgt := networktypes.IPAMConfig{
		Subnet:  m.Subnet,
		IPRange: m.IpRange,
		Gateway, m.Gateway,
		AuxAddress, make(map[string]string),
	}
	for k, v := range m.AuxAddress {
		tgt.AuxAddress[k] = v
	}
	return tgt
}
