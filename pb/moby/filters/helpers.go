package filters

import (
	"github.com/docker/docker/api/types/filters"
)

func (s *Args) DeepCopyChecked() *Args {
	if s == nil {
		return nil
	}
	dst := &Args{
		Fields: make(map[string]*Args_Value),
	}
	for k, v := range s.Fields {
		ele := &Args_Value{
			Value: make(map[string]bool),
		}
		for k1, v1 := range v.Value {
			ele.Value[k1] = v1
		}
		dst.Fields[k] = ele
	}
	return dst
}

func (s *Args) ExportIntoDockerApiType() (d filters.Args) {
	if s != nil {
		for k, v := range s.Fields {
			if v != nil {
				for k1, _ := range v.Value {
					d.Add(k, k1)
				}
			}
		}
	}
	return
}
