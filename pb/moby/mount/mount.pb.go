// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/moby/mount/mount.proto

/*
Package mount is a generated protocol buffer package.

It is generated from these files:
	pb/moby/mount/mount.proto

It has these top-level messages:
	Mount
	BindOptions
	VolumeOptions
	Driver
	TmpfsOptions
*/
package mount

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Mount represents a mount (volume).
// type Mount struct
type Mount struct {
	// Type Type `json:",omitempty"`
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Source specifies the name of the mount. Depending on mount type, this may be a volume name or a host path, or even ignored.
	// Source is not supported for tmpfs (must be an empty value)
	// Source string `json:",omitempty"`
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	// Target string `json:",omitempty"`
	Target string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	// ReadOnly bool `json:",omitempty"`
	ReadOnly bool `protobuf:"varint,4,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
	// Consistency consistency `json:",omitempty"`
	// Consistency represents the consistency requirements of a mount. // type Consistency string
	Consistency string `protobuf:"bytes,5,opt,name=consistency,proto3" json:"consistency,omitempty"`
	// BindOptions *BindOptions `json:",omitempty"`
	BindOptions *BindOptions `protobuf:"bytes,6,opt,name=bind_options,json=bindOptions" json:"bind_options,omitempty"`
	// VolumeOptions *VolumeOptions `json:",omitempty"`
	VolumeOptions *VolumeOptions `protobuf:"bytes,7,opt,name=volume_options,json=volumeOptions" json:"volume_options,omitempty"`
	// TmpfsOptions *TmpfsOptions `json:",omitempty"`
	TmpfsOptions *TmpfsOptions `protobuf:"bytes,8,opt,name=tmpfs_options,json=tmpfsOptions" json:"tmpfs_options,omitempty"`
}

func (m *Mount) Reset()                    { *m = Mount{} }
func (m *Mount) String() string            { return proto.CompactTextString(m) }
func (*Mount) ProtoMessage()               {}
func (*Mount) Descriptor() ([]byte, []int) { return fileDescriptorMount, []int{0} }

func (m *Mount) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Mount) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Mount) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Mount) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

func (m *Mount) GetConsistency() string {
	if m != nil {
		return m.Consistency
	}
	return ""
}

func (m *Mount) GetBindOptions() *BindOptions {
	if m != nil {
		return m.BindOptions
	}
	return nil
}

func (m *Mount) GetVolumeOptions() *VolumeOptions {
	if m != nil {
		return m.VolumeOptions
	}
	return nil
}

func (m *Mount) GetTmpfsOptions() *TmpfsOptions {
	if m != nil {
		return m.TmpfsOptions
	}
	return nil
}

// BindOptions defines options specific to mounts of type "bind".
// type BindOptions struct
type BindOptions struct {
	// Propagation Propagation `json",omitempty"`
	// Propagation represents the propagation of a mount. // type Propagation string
	Propagation string `protobuf:"bytes,1,opt,name=propagation,proto3" json:"propagation,omitempty"`
}

func (m *BindOptions) Reset()                    { *m = BindOptions{} }
func (m *BindOptions) String() string            { return proto.CompactTextString(m) }
func (*BindOptions) ProtoMessage()               {}
func (*BindOptions) Descriptor() ([]byte, []int) { return fileDescriptorMount, []int{1} }

func (m *BindOptions) GetPropagation() string {
	if m != nil {
		return m.Propagation
	}
	return ""
}

// VolumeOptions represents the options for a mount of type volume.
// type VolumeOptions struct
type VolumeOptions struct {
	// NoCopy bool `json:",omitempty"`
	NoCopy bool `protobuf:"varint,1,opt,name=no_copy,json=noCopy,proto3" json:"no_copy,omitempty"`
	// Labels map[string]string `json:",omitempty"`
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// DriverConfig *Driver `json:",omitempty"`
	DriverConfig *Driver `protobuf:"bytes,3,opt,name=driver_config,json=driverConfig" json:"driver_config,omitempty"`
}

func (m *VolumeOptions) Reset()                    { *m = VolumeOptions{} }
func (m *VolumeOptions) String() string            { return proto.CompactTextString(m) }
func (*VolumeOptions) ProtoMessage()               {}
func (*VolumeOptions) Descriptor() ([]byte, []int) { return fileDescriptorMount, []int{2} }

func (m *VolumeOptions) GetNoCopy() bool {
	if m != nil {
		return m.NoCopy
	}
	return false
}

func (m *VolumeOptions) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *VolumeOptions) GetDriverConfig() *Driver {
	if m != nil {
		return m.DriverConfig
	}
	return nil
}

// Driver represents a volume driver.
// type Driver struct
type Driver struct {
	// Name string `json:",omitempty"`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Options map[string]string `json:",omitempty"`
	Options map[string]string `protobuf:"bytes,2,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Driver) Reset()                    { *m = Driver{} }
func (m *Driver) String() string            { return proto.CompactTextString(m) }
func (*Driver) ProtoMessage()               {}
func (*Driver) Descriptor() ([]byte, []int) { return fileDescriptorMount, []int{3} }

func (m *Driver) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Driver) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

// TmpfsOptions defines options specific to mounts of type "tmpfs".
// type TmpfsOptions struct
type TmpfsOptions struct {
	// Size sets the size of the tmpfs, in bytes.
	//
	// This will be converted to an operating system specific value
	// depending on the host. For example, on linux, it will be converted to
	// use a 'k', 'm' or 'g' syntax. BSD, though not widely supported with
	// docker, uses a straight byte value.
	//
	// Percentages are not supported.
	// SizeBytes int84 `json:",omitempty"`
	SizeBytes int64 `protobuf:"varint,1,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	// Mode of the tmpfs upon creation
	// Mode os.FileMode `json:",omitempty"`
	Mode uint32 `protobuf:"varint,2,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (m *TmpfsOptions) Reset()                    { *m = TmpfsOptions{} }
func (m *TmpfsOptions) String() string            { return proto.CompactTextString(m) }
func (*TmpfsOptions) ProtoMessage()               {}
func (*TmpfsOptions) Descriptor() ([]byte, []int) { return fileDescriptorMount, []int{4} }

func (m *TmpfsOptions) GetSizeBytes() int64 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

func (m *TmpfsOptions) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func init() {
	proto.RegisterType((*Mount)(nil), "mount.Mount")
	proto.RegisterType((*BindOptions)(nil), "mount.BindOptions")
	proto.RegisterType((*VolumeOptions)(nil), "mount.VolumeOptions")
	proto.RegisterType((*Driver)(nil), "mount.Driver")
	proto.RegisterType((*TmpfsOptions)(nil), "mount.TmpfsOptions")
}

func init() { proto.RegisterFile("pb/moby/mount/mount.proto", fileDescriptorMount) }

var fileDescriptorMount = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x25, 0xed, 0x36, 0x6d, 0x6f, 0x12, 0x91, 0xeb, 0xa2, 0xb1, 0x22, 0x84, 0x3e, 0xf5, 0xa9,
	0x85, 0xaa, 0x50, 0xd7, 0x27, 0x77, 0xf5, 0x4d, 0x59, 0x08, 0xe2, 0x6b, 0xc8, 0xc7, 0x6c, 0x09,
	0x26, 0x73, 0x87, 0x64, 0x1a, 0x18, 0x7f, 0x84, 0xbf, 0xc7, 0xbf, 0xe1, 0x3f, 0x92, 0x99, 0x49,
	0x6a, 0x0a, 0xbe, 0xec, 0x4b, 0xb8, 0xe7, 0xdc, 0x7b, 0x2e, 0xe7, 0xcc, 0x4c, 0xe0, 0xa5, 0xc8,
	0x76, 0x35, 0x65, 0x6a, 0x57, 0xd3, 0x89, 0x4b, 0xfb, 0xdd, 0x8a, 0x86, 0x24, 0xe1, 0xcc, 0x80,
	0xf5, 0xef, 0x09, 0xcc, 0xbe, 0xea, 0x0a, 0x11, 0xae, 0xa4, 0x12, 0x2c, 0x74, 0x22, 0x67, 0xb3,
	0x8c, 0x4d, 0x8d, 0xcf, 0xc1, 0x6d, 0xe9, 0xd4, 0xe4, 0x2c, 0x9c, 0x18, 0xb6, 0x47, 0x9a, 0x97,
	0x69, 0x73, 0x64, 0x32, 0x9c, 0x5a, 0xde, 0x22, 0x7c, 0x05, 0xcb, 0x86, 0xa5, 0x45, 0x42, 0xbc,
	0x52, 0xe1, 0x55, 0xe4, 0x6c, 0x16, 0xf1, 0x42, 0x13, 0xf7, 0xbc, 0x52, 0x18, 0x81, 0x97, 0x13,
	0x6f, 0xcb, 0x56, 0x32, 0x9e, 0xab, 0x70, 0x66, 0x94, 0x63, 0x0a, 0xdf, 0x81, 0x9f, 0x95, 0xbc,
	0x48, 0x48, 0xc8, 0x92, 0x78, 0x1b, 0xba, 0x91, 0xb3, 0xf1, 0xf6, 0xb8, 0xb5, 0xbe, 0x6f, 0x4b,
	0x5e, 0xdc, 0xdb, 0x4e, 0xec, 0x65, 0xff, 0x00, 0x7e, 0x80, 0x27, 0x1d, 0x55, 0xa7, 0x9a, 0x9d,
	0x85, 0x73, 0x23, 0xbc, 0xee, 0x85, 0xdf, 0x4d, 0x73, 0x90, 0x06, 0xdd, 0x18, 0xe2, 0x01, 0x02,
	0x59, 0x8b, 0x87, 0xf6, 0xac, 0x5d, 0x18, 0xed, 0xb3, 0x5e, 0xfb, 0x4d, 0xf7, 0x06, 0xa9, 0x2f,
	0x47, 0x68, 0xbd, 0x03, 0x6f, 0x64, 0x49, 0xc7, 0x13, 0x0d, 0x89, 0xf4, 0x98, 0x6a, 0xdc, 0x1f,
	0xe3, 0x98, 0x5a, 0xff, 0x71, 0x20, 0xb8, 0xf0, 0x82, 0x2f, 0x60, 0xce, 0x29, 0xc9, 0x49, 0x28,
	0x33, 0xbf, 0x88, 0x5d, 0x4e, 0x77, 0x24, 0x14, 0x1e, 0xc0, 0xad, 0xd2, 0x8c, 0x55, 0x6d, 0x38,
	0x89, 0xa6, 0x1b, 0x6f, 0x1f, 0xfd, 0x2f, 0xca, 0xf6, 0x8b, 0x19, 0xf9, 0xcc, 0x65, 0xa3, 0xe2,
	0x7e, 0x1e, 0xf7, 0x10, 0x14, 0x4d, 0xd9, 0xb1, 0x26, 0xc9, 0x89, 0x3f, 0x94, 0x47, 0x73, 0x43,
	0xde, 0x3e, 0xe8, 0x17, 0x7c, 0x32, 0xbd, 0xd8, 0xb7, 0x33, 0x77, 0x66, 0x64, 0xf5, 0x1e, 0xbc,
	0xd1, 0x2a, 0x7c, 0x0a, 0xd3, 0x1f, 0x4c, 0xf5, 0x09, 0x74, 0x89, 0xd7, 0x30, 0xeb, 0xd2, 0xea,
	0x34, 0x3c, 0x03, 0x0b, 0x6e, 0x26, 0x07, 0x67, 0xfd, 0xcb, 0x01, 0xd7, 0xee, 0xd4, 0x0f, 0x88,
	0xa7, 0xf5, 0xf9, 0x01, 0xe9, 0x1a, 0xdf, 0xc2, 0x7c, 0x38, 0x57, 0x1b, 0x64, 0x75, 0xe1, 0x63,
	0xdb, 0x27, 0xb1, 0x11, 0x86, 0xd1, 0xd5, 0x0d, 0xf8, 0xe3, 0xc6, 0xa3, 0x0c, 0x7d, 0x04, 0x7f,
	0x7c, 0x67, 0xf8, 0x1a, 0xa0, 0x2d, 0x7f, 0xb2, 0x24, 0x53, 0x92, 0xb5, 0x66, 0xc5, 0x34, 0x5e,
	0x6a, 0xe6, 0x56, 0x13, 0xda, 0x74, 0x4d, 0x85, 0xdd, 0x13, 0xc4, 0xa6, 0xce, 0x5c, 0xf3, 0x87,
	0xbc, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x48, 0x38, 0x48, 0xf2, 0x3e, 0x03, 0x00, 0x00,
}