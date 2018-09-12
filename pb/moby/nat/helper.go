package nat

import (
	"errors"

	nattypes "github.com/docker/go-connections/nat"
)

var (
	ErrNilPointer = errors.New("found nil pointer")
)

func (m *PortMap) DeepCopyChecked() (*PortMap, []error) {
	var obj *PortMap = nil
	var errorList []error = []error{}
	if m == nil {
		errorList = append(errorList, ErrNilPointer)
	} else {
		obj = &PortMap{
			InternalMap: make(map[string]*PortMap_PortBindingSlice),
		}
		for k, v := range m.InternalMap {
			ele := new(PortMap_PortBindingSlice)
			if v != nil {
				ele.InternalList = make([]*PortBinding, len(v.InternalList))
				for _, v1 := range v.InternalList {
					ele.InternalList = append(ele.InternalList, &PortBinding{
						HostIp:   v1.HostIp,
						HostPort: v1.HostPort,
					})
				}
			}
			obj.InternalMap[k] = ele
		}
	}
	return obj, []error{}
}

func (m *PortMap) ExportIntoDockerApiType() (tgt nattypes.PortMap) {
	if m == nil || m.InternalMap == nil {
		return
	}
	tgt = nattypes.PortMap(make(map[nattypes.Port][]nattypes.PortBinding))
	for k, v := range m.InternalMap {
		var ele []nattypes.PortBinding
		if v != nil {
			ele = make([]nattypes.PortBinding, len(v.InternalList))
			for _, v1 := range v.InternalList {
				if v1 != nil {
					ele = append(ele, nattypes.PortBinding{
						HostIP:   v1.HostIp,
						HostPort: v1.HostPort,
					})
				}
			}
		}
		tgt[nattypes.Port(k)] = ele
	}
	return
}

func ConvertFromDockerApiTypePortMap(s nattypes.PortMap) (dst *PortMap) {
	dst = &PortMap{
		InternalMap: make(map[string]*PortMap_PortBindingSlice),
	}
	for k, v := range s {
		ele := &PortMap_PortBindingSlice{
			InternalList: make([]*PortBinding, len(v)),
		}
		for _, v1 := range v {
			ele.InternalList = append(ele.InternalList, &PortBinding{
				HostIp:   v1.HostIP,
				HostPort: v1.HostPort,
			})
		}
		dst.InternalMap[string(k)] = ele
	}
	return
}

func (m *PortSet) DeepCopyChecked() (*PortSet, []error) {
	if m == nil {
		return nil, []error{ErrNilPointer}
	}
	obj := &PortSet{
		InternalMap: make(map[string]*PortSet_VoidStruct),
	}
	for k, v := range m.InternalMap {
		obj.InternalMap[k] = v
	}
	return obj, []error{}
}

func (m *PortSet) ExportIntoDockerApiType() (tgt nattypes.PortSet) {
	if m == nil || m.InternalMap == nil {
		return
	}
	tgt = nattypes.PortSet(make(map[nattypes.Port]struct{}))
	for k, _ := range m.InternalMap {
		tgt[nattypes.Port(k)] = struct{}{}
	}
	return
}

func ConvertFromDockerApiTypePortSet(s *nattypes.PortSet) (dst *PortSet) {
	if s == nil {
		return
	}
	dst = &PortSet{
		InternalMap: (map[string]*PortSet_VoidStruct)(nil),
	}
	for k, _ := range *s {
		dst.InternalMap[string(k)] = &PortSet_VoidStruct{}
	}
	return
}
