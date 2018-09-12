package mount

import (
	"errors"
	"os"

	mounttypes "github.com/docker/docker/api/types/mount"
)

var (
	ErrNilPointer = errors.New("found nil pointer")
)

func (m *Mount) DeepCopyChecked() (*Mount, []error) {
	if m == nil {
		return nil, []error{ErrNilPointer}
	}
	obj := &Mount{
		Type:        m.Type,
		Source:      m.Source,
		Target:      m.Target,
		ReadOnly:    m.ReadOnly,
		Consistency: m.Consistency,
	}
	if v := m.BindOptions; v != nil {
		obj.BindOptions = &BindOptions{
			Propagation: v.Propagation,
		}
	}
	if v := m.VolumeOptions; v != nil {
		obj.VolumeOptions = &VolumeOptions{
			NoCopy: v.NoCopy,
		}
		if len(v.Labels) != 0 {
			obj.VolumeOptions.Labels = make(map[string]string)
			for k1, v1 := range v.Labels {
				obj.VolumeOptions.Labels[k1] = v1
			}
		}
		if v1 := v.DriverConfig; v1 != nil {
			obj.VolumeOptions.DriverConfig = &Driver{
				Name: v1.Name,
			}
			if len(v1.Options) != 0 {
				obj.VolumeOptions.DriverConfig.Options = make(map[string]string)
				for k2, v2 := range v1.Options {
					obj.VolumeOptions.DriverConfig.Options[k2] = v2
				}
			}
		}
	}
	if v := m.TmpfsOptions; v != nil {
		obj.TmpfsOptions = &TmpfsOptions{
			SizeBytes: v.SizeBytes,
			Mode:      v.Mode,
		}
	}
	return obj, []error{}
}

func (m *Mount) ExportIntoDockerApiType() (tgt mounttypes.Mount) {
	if m != nil {
		tgt = mounttypes.Mount{
			Type:        mounttypes.Type(m.Type),
			Source:      m.Source,
			ReadOnly:    m.ReadOnly,
			Consistency: mounttypes.Consistency(m.Consistency),
		}
		if v := m.BindOptions; v != nil {
			tgt.BindOptions = &mounttypes.BindOptions{
				Propagation: mounttypes.Propagation(v.Propagation),
			}
		}
		if v := m.VolumeOptions; v != nil {
			tgt.VolumeOptions = &mounttypes.VolumeOptions{
				NoCopy: v.NoCopy,
				Labels: make(map[string]string),
			}
			for k1, v1 := range v.Labels {
				tgt.VolumeOptions.Labels[k1] = v1
			}
			if v1 := v.DriverConfig; v1 != nil {
				tgt.VolumeOptions.DriverConfig = &mounttypes.Driver{
					Name:    v1.Name,
					Options: make(map[string]string),
				}
				for k2, v2 := range v1.Options {
					tgt.VolumeOptions.DriverConfig.Options[k2] = v2
				}
			}
		}
		if v := m.TmpfsOptions; v != nil {
			tgt.TmpfsOptions = &mounttypes.TmpfsOptions{
				SizeBytes: v.SizeBytes,
				Mode:      os.FileMode(v.Mode),
			}
		}
	}
	return
}

func ConvertFromDockerApiTypeMount(s mounttypes.Mount) (dst *Mount) {
	dst = &Mount{
		Type:        string(s.Type),
		Source:      s.Source,
		Target:      s.Target,
		ReadOnly:    s.ReadOnly,
		Consistency: string(s.Consistency),

		BindOptions:   (*BindOptions)(nil),
		VolumeOptions: (*VolumeOptions)(nil),
		TmpfsOptions:  (*TmpfsOptions)(nil),
	}
	if v := s.BindOptions; v != nil {
		s := &BindOptions{
			Propagation: string(v.Propagation),
		}
		dst.BindOptions = s
	}
	if v := s.VolumeOptions; v != nil {
		s := &VolumeOptions{
			NoCopy:       v.NoCopy,
			Labels:       make(map[string]string),
			DriverConfig: (*Driver)(nil),
		}
		for k1, v1 := range v.Labels {
			s.Labels[k1] = v1
		}
		if v1 := v.DriverConfig; v1 != nil {
			o1 := &Driver{
				Name:    v1.Name,
				Options: make(map[string]string),
			}
			for k2, v2 := range v1.Options {
				o1.Options[k2] = v2
			}
			s.DriverConfig = o1
		}
		dst.VolumeOptions = s
	}
	if v := s.TmpfsOptions; v != nil {
		s := &TmpfsOptions{
			SizeBytes: v.SizeBytes,
			Mode:      uint32(v.Mode),
		}
		dst.TmpfsOptions = s
	}
	return
}
