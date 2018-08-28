package moby

import (
	"errors"

	"github.com/moby/moby/api/types"
	// networktypes "github.com/moby/moby/api/types/network"
)

const (
	BuilderV1       string = "1"
	BuilderBuildKit string = "2"
)

var (
	ErrNilPointer         = errors.New("Pointer is nil")
	ErrNetworkCreateIsNil = errors.New("NetworkCreate not specified")
)

func (m *ImagePullOptions) CopyWithValidation() (*ImagePullOptions, []error) {
	tgt := new(ImagePullOptions)
	if m == nil {
		return tgt, []error{ErrNilPointer}
	}
	tgt.All = m.All
	tgt.RegistryAuth = m.RegistryAuth
	return tgt, []error{}
}

func (m *ImagePullOptions) ConvertIntoDockerApiTypes() types.ImagePullOptions {
	return types.ImagePullOptions{
		All:          m.All,
		RegistryAuth: m.RegistryAuth,
		Platform:     m.Platform,
	}
}

func (m *NetworkCreate) CopyWithValidation() (*NetworkCreate, []error) {
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
	tgt.IPAM, errlst = m.IPAM.CopyWithValidation()
	if len(errlst) != 0 {
		errs = append(errs, errlst...)
	}
	tgt.Internal = m.Internal
	tgt.Attachable = m.Attachable
	tgt.Ingress = m.Ingress
	tgt.ConfigOnly = m.ConfigOnly
	tgt.ConfigFrom, _ = m.ConfigFrom.CopyWithValidation()
	tgt.Options = make(map[sstring]string)
	for k, v := range m.Options {
		tgt.Options[k] = v
	}
	tgt.Lables = make(map[string]string)
	for k, v := range m.Labels {
		tgt.Labels[k] = v
	}
	return tgt, errs
}

func (m *NetworkCreate) ConvertIntoDockerApiTypes() types.NetworkCreate {
	tgt := types.NetworkCreate{
		CheckDuplicate: m.CheckDuplicate,
		Driver:         m.Driver,
		Scope:          m.Scope,
		EnableIPv6:     m.EnableIpv6,
		IPAM:           m.Ipam.ConvertIntoDockerApiTypes(),
		Internal:       m.Internal,
		Attachable:     m.Attachable,
		Ingress:        m.Ingress,
		ConfigOnly:     m.ConfigOnly,
		ConfigFrom:     m.ConfigFrom.ConvertIntoDockerApiTypes(),
		Options:        make(map[string]string),
		Lables:         make(map[string]string),
	}

	for k, v := range m.Options {
		tgt.Options[k] = v
	}

	for k, v := range m.Labels {
		tgt.Lables[k] = v
	}

	return tgt
}

func (m *NetworkCreateResponse) ConvertFromDockerApiTypes(resp types.NetworkCreateResponse) *NetworkCreateResponse {
	return &NetworkCreateResponse{
		Id:      resp.ID,
		Warning: resp.Warning,
	}
}
