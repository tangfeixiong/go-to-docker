// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/moby/network/network.proto

/*
Package network is a generated protocol buffer package.

It is generated from these files:
	pb/moby/network/network.proto

It has these top-level messages:
	Address
	IPAM
	IPAMConfig
	EndpointIPAMConfig
	PeerInfo
	EndpointSettings
	Task
	ServiceInfo
	NetworkingConfig
	ConfigReference
*/
package network

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

// Address represents an IP address
// type Address struct
type Address struct {
	// Addr string
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	// PrefixLen int
	PrefixLen int32 `protobuf:"varint,2,opt,name=prefix_len,json=prefixLen,proto3" json:"prefix_len,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{0} }

func (m *Address) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Address) GetPrefixLen() int32 {
	if m != nil {
		return m.PrefixLen
	}
	return 0
}

// IPAM represents IP Address Management
// type IPAM struct
type IPAM struct {
	// Driver string
	Driver string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	// Options map[string]string //Per network IPAM driver options
	Options map[string]string `protobuf:"bytes,2,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Config []IPAMConfig
	Config []*IPAMConfig `protobuf:"bytes,3,rep,name=config" json:"config,omitempty"`
}

func (m *IPAM) Reset()                    { *m = IPAM{} }
func (m *IPAM) String() string            { return proto.CompactTextString(m) }
func (*IPAM) ProtoMessage()               {}
func (*IPAM) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{1} }

func (m *IPAM) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *IPAM) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *IPAM) GetConfig() []*IPAMConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

// IPAMConfig represents IPAM configurations
// type IPAMConfig struct
type IPAMConfig struct {
	// Subnet string `json:",omitempty"`
	Subnet string `protobuf:"bytes,1,opt,name=subnet,proto3" json:"subnet,omitempty"`
	// IPRange string `json:",omitempty"`
	IpRange string `protobuf:"bytes,2,opt,name=ip_range,json=ipRange,proto3" json:"ip_range,omitempty"`
	// Gateway string `json:",omitempty"`
	Gateway string `protobuf:"bytes,3,opt,name=gateway,proto3" json:"gateway,omitempty"`
	// AuxAddress map[string]string `json:"AuxiliaryAddresses,omitempty"
	AuxAddress map[string]string `protobuf:"bytes,4,rep,name=aux_address,json=auxAddress" json:"aux_address,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *IPAMConfig) Reset()                    { *m = IPAMConfig{} }
func (m *IPAMConfig) String() string            { return proto.CompactTextString(m) }
func (*IPAMConfig) ProtoMessage()               {}
func (*IPAMConfig) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{2} }

func (m *IPAMConfig) GetSubnet() string {
	if m != nil {
		return m.Subnet
	}
	return ""
}

func (m *IPAMConfig) GetIpRange() string {
	if m != nil {
		return m.IpRange
	}
	return ""
}

func (m *IPAMConfig) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *IPAMConfig) GetAuxAddress() map[string]string {
	if m != nil {
		return m.AuxAddress
	}
	return nil
}

// EndpointIPAMConfig represents IPAM configurations for the endpoint
// type EndpointIPAMConfig struct
type EndpointIPAMConfig struct {
	// IPv4Address string `json:",omitempty"`
	Ipv4Address string `protobuf:"bytes,1,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	// IPv6Address string `json:",omitempty"`
	Ipv6Address string `protobuf:"bytes,2,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	// LinkLocalIPs []string `json:",omitempty"`
	LinkLocalIps []string `protobuf:"bytes,3,rep,name=link_local_ips,json=linkLocalIps" json:"link_local_ips,omitempty"`
}

func (m *EndpointIPAMConfig) Reset()                    { *m = EndpointIPAMConfig{} }
func (m *EndpointIPAMConfig) String() string            { return proto.CompactTextString(m) }
func (*EndpointIPAMConfig) ProtoMessage()               {}
func (*EndpointIPAMConfig) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{3} }

func (m *EndpointIPAMConfig) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *EndpointIPAMConfig) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

func (m *EndpointIPAMConfig) GetLinkLocalIps() []string {
	if m != nil {
		return m.LinkLocalIps
	}
	return nil
}

// PeerInfo represents one peer of an overlay network
// type PeerInfo struct
type PeerInfo struct {
	// Name string
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// IP string
	Ip string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (m *PeerInfo) Reset()                    { *m = PeerInfo{} }
func (m *PeerInfo) String() string            { return proto.CompactTextString(m) }
func (*PeerInfo) ProtoMessage()               {}
func (*PeerInfo) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{4} }

func (m *PeerInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PeerInfo) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

// EndpointSettings stores the network endpoint details
// type EndpointSettings struct
type EndpointSettings struct {
	// IPAMConfig *EndpointIPAMConfig	// Configurations
	IpamConfig *EndpointIPAMConfig `protobuf:"bytes,1,opt,name=ipam_config,json=ipamConfig" json:"ipam_config,omitempty"`
	// Links []string
	Links []string `protobuf:"bytes,2,rep,name=links" json:"links,omitempty"`
	// Aliases []string
	Aliases []string `protobuf:"bytes,3,rep,name=aliases" json:"aliases,omitempty"`
	// NetworkID string // Operational data
	NetworkId string `protobuf:"bytes,4,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// EndpointID string
	EndpointId string `protobuf:"bytes,5,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	// Gateway string
	Gateway string `protobuf:"bytes,6,opt,name=gateway,proto3" json:"gateway,omitempty"`
	// IPAddress string
	IpAddress string `protobuf:"bytes,7,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	// IPPrefixLen int
	IpPrefixLen int32 `protobuf:"varint,8,opt,name=ip_prefix_len,json=ipPrefixLen,proto3" json:"ip_prefix_len,omitempty"`
	// IPv6Gateway string
	Ipv6Gateway string `protobuf:"bytes,9,opt,name=ipv6_gateway,json=ipv6Gateway,proto3" json:"ipv6_gateway,omitempty"`
	// GlobalIPv6Address string
	GlobalIpv6Address string `protobuf:"bytes,10,opt,name=global_ipv6_address,json=globalIpv6Address,proto3" json:"global_ipv6_address,omitempty"`
	// GlobalIPv6PrefixLen int
	GlobalIpv6PrefixLen int32 `protobuf:"varint,11,opt,name=global_ipv6_prefix_len,json=globalIpv6PrefixLen,proto3" json:"global_ipv6_prefix_len,omitempty"`
	// MacAddress string
	MacAddress string `protobuf:"bytes,12,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	// DriverOpts map[string]string
	DriverOpts map[string]string `protobuf:"bytes,13,rep,name=driver_opts,json=driverOpts" json:"driver_opts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *EndpointSettings) Reset()                    { *m = EndpointSettings{} }
func (m *EndpointSettings) String() string            { return proto.CompactTextString(m) }
func (*EndpointSettings) ProtoMessage()               {}
func (*EndpointSettings) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{5} }

func (m *EndpointSettings) GetIpamConfig() *EndpointIPAMConfig {
	if m != nil {
		return m.IpamConfig
	}
	return nil
}

func (m *EndpointSettings) GetLinks() []string {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *EndpointSettings) GetAliases() []string {
	if m != nil {
		return m.Aliases
	}
	return nil
}

func (m *EndpointSettings) GetNetworkId() string {
	if m != nil {
		return m.NetworkId
	}
	return ""
}

func (m *EndpointSettings) GetEndpointId() string {
	if m != nil {
		return m.EndpointId
	}
	return ""
}

func (m *EndpointSettings) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *EndpointSettings) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *EndpointSettings) GetIpPrefixLen() int32 {
	if m != nil {
		return m.IpPrefixLen
	}
	return 0
}

func (m *EndpointSettings) GetIpv6Gateway() string {
	if m != nil {
		return m.Ipv6Gateway
	}
	return ""
}

func (m *EndpointSettings) GetGlobalIpv6Address() string {
	if m != nil {
		return m.GlobalIpv6Address
	}
	return ""
}

func (m *EndpointSettings) GetGlobalIpv6PrefixLen() int32 {
	if m != nil {
		return m.GlobalIpv6PrefixLen
	}
	return 0
}

func (m *EndpointSettings) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *EndpointSettings) GetDriverOpts() map[string]string {
	if m != nil {
		return m.DriverOpts
	}
	return nil
}

// Task carries the information about one backend task
// type Task struct
type Task struct {
	// Name string
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// EndpointID string
	EndpointId string `protobuf:"bytes,2,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	// EndpointIP string
	EndpointIp string `protobuf:"bytes,3,opt,name=endpoint_ip,json=endpointIp,proto3" json:"endpoint_ip,omitempty"`
	// Info map[string]string
	Info map[string]string `protobuf:"bytes,4,rep,name=info" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{6} }

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetEndpointId() string {
	if m != nil {
		return m.EndpointId
	}
	return ""
}

func (m *Task) GetEndpointIp() string {
	if m != nil {
		return m.EndpointIp
	}
	return ""
}

func (m *Task) GetInfo() map[string]string {
	if m != nil {
		return m.Info
	}
	return nil
}

// ServiceInfo represents service parameters with the list of service's tasks
// type ServiceInfo struct
type ServiceInfo struct {
	// VIP string
	Vip string `protobuf:"bytes,1,opt,name=vip,proto3" json:"vip,omitempty"`
	// Ports []string
	Ports []string `protobuf:"bytes,2,rep,name=ports" json:"ports,omitempty"`
	// LocalLBIndex int
	LocalLbIndex int32 `protobuf:"varint,3,opt,name=local_lb_index,json=localLbIndex,proto3" json:"local_lb_index,omitempty"`
	// Tasks []Task
	Tasks []*Task `protobuf:"bytes,4,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *ServiceInfo) Reset()                    { *m = ServiceInfo{} }
func (m *ServiceInfo) String() string            { return proto.CompactTextString(m) }
func (*ServiceInfo) ProtoMessage()               {}
func (*ServiceInfo) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{7} }

func (m *ServiceInfo) GetVip() string {
	if m != nil {
		return m.Vip
	}
	return ""
}

func (m *ServiceInfo) GetPorts() []string {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *ServiceInfo) GetLocalLbIndex() int32 {
	if m != nil {
		return m.LocalLbIndex
	}
	return 0
}

func (m *ServiceInfo) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

// NetworkingConfig represents the container's networking configuration for each of its interfaces Carries the networking configs specified in the 'docker run' and 'docker network connect' commands
// type NetworkingConfig struct
type NetworkingConfig struct {
	// EndpointsConfig map[string]*EndpointSettings // Endpoint configs for each connecting network
	EndpointsConfig map[string]*EndpointSettings `protobuf:"bytes,1,rep,name=endpoints_config,json=endpointsConfig" json:"endpoints_config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *NetworkingConfig) Reset()                    { *m = NetworkingConfig{} }
func (m *NetworkingConfig) String() string            { return proto.CompactTextString(m) }
func (*NetworkingConfig) ProtoMessage()               {}
func (*NetworkingConfig) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{8} }

func (m *NetworkingConfig) GetEndpointsConfig() map[string]*EndpointSettings {
	if m != nil {
		return m.EndpointsConfig
	}
	return nil
}

// ConfigReference specifies the source which provides a network's cconfiguration
// type ConfigReference struct
type ConfigReference struct {
	// Network string
	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
}

func (m *ConfigReference) Reset()                    { *m = ConfigReference{} }
func (m *ConfigReference) String() string            { return proto.CompactTextString(m) }
func (*ConfigReference) ProtoMessage()               {}
func (*ConfigReference) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{9} }

func (m *ConfigReference) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func init() {
	proto.RegisterType((*Address)(nil), "network.Address")
	proto.RegisterType((*IPAM)(nil), "network.IPAM")
	proto.RegisterType((*IPAMConfig)(nil), "network.IPAMConfig")
	proto.RegisterType((*EndpointIPAMConfig)(nil), "network.EndpointIPAMConfig")
	proto.RegisterType((*PeerInfo)(nil), "network.PeerInfo")
	proto.RegisterType((*EndpointSettings)(nil), "network.EndpointSettings")
	proto.RegisterType((*Task)(nil), "network.Task")
	proto.RegisterType((*ServiceInfo)(nil), "network.ServiceInfo")
	proto.RegisterType((*NetworkingConfig)(nil), "network.NetworkingConfig")
	proto.RegisterType((*ConfigReference)(nil), "network.ConfigReference")
}

func init() { proto.RegisterFile("pb/moby/network/network.proto", fileDescriptorNetwork) }

var fileDescriptorNetwork = []byte{
	// 799 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0x96, 0xf3, 0x8b, 0x8f, 0x03, 0xe4, 0x0e, 0x88, 0x6b, 0x72, 0x85, 0xe0, 0x86, 0xbb, 0xe0,
	0x0a, 0xc9, 0x48, 0x80, 0x68, 0x85, 0xe8, 0x02, 0x15, 0x54, 0xa5, 0xa2, 0x05, 0x99, 0x6e, 0xba,
	0xa8, 0xac, 0x49, 0x3c, 0x89, 0x46, 0x71, 0xc6, 0x23, 0xdb, 0x09, 0xc9, 0xae, 0x2f, 0xd6, 0x4d,
	0xd7, 0xdd, 0xf4, 0x05, 0xfa, 0x2c, 0xd5, 0xfc, 0x39, 0x4e, 0x48, 0x17, 0xac, 0xe2, 0x73, 0xce,
	0x77, 0xfe, 0xcf, 0x37, 0x81, 0x3d, 0xde, 0x3d, 0x19, 0xc5, 0xdd, 0xd9, 0x09, 0x23, 0xd9, 0x53,
	0x9c, 0x0c, 0xcd, 0xaf, 0xc7, 0x93, 0x38, 0x8b, 0x51, 0x5d, 0x8b, 0xed, 0x2b, 0xa8, 0x5f, 0x87,
	0x61, 0x42, 0xd2, 0x14, 0x21, 0xa8, 0xe0, 0x30, 0x4c, 0x5c, 0xeb, 0xc0, 0x3a, 0xb2, 0x7d, 0xf9,
	0x8d, 0xf6, 0x00, 0x78, 0x42, 0xfa, 0x74, 0x1a, 0x44, 0x84, 0xb9, 0xa5, 0x03, 0xeb, 0xa8, 0xea,
	0xdb, 0x4a, 0x73, 0x47, 0x58, 0xfb, 0x9b, 0x05, 0x95, 0xce, 0xc3, 0xf5, 0x07, 0xb4, 0x03, 0xb5,
	0x30, 0xa1, 0x13, 0x62, 0xbc, 0xb5, 0x84, 0xce, 0xa1, 0x1e, 0xf3, 0x8c, 0xc6, 0x2c, 0x75, 0x4b,
	0x07, 0xe5, 0x23, 0xe7, 0xb4, 0xe5, 0x99, 0x42, 0x84, 0x9f, 0x77, 0xaf, 0x8c, 0xb7, 0x2c, 0x4b,
	0x66, 0xbe, 0x81, 0xa2, 0x63, 0xa8, 0xf5, 0x62, 0xd6, 0xa7, 0x03, 0xb7, 0x2c, 0x9d, 0xb6, 0x16,
	0x9c, 0xde, 0x4a, 0x93, 0xaf, 0x21, 0xad, 0x4b, 0x68, 0x14, 0xa3, 0xa0, 0x26, 0x94, 0x87, 0x64,
	0xa6, 0xeb, 0x10, 0x9f, 0x68, 0x1b, 0xaa, 0x13, 0x1c, 0x8d, 0x89, 0xac, 0xdf, 0xf6, 0x95, 0x70,
	0x59, 0x7a, 0x6d, 0xb5, 0x7f, 0x59, 0x00, 0xf3, 0x90, 0xa2, 0x8b, 0x74, 0xdc, 0x65, 0x24, 0x33,
	0x5d, 0x28, 0x09, 0xed, 0xc2, 0x1a, 0xe5, 0x41, 0x82, 0xd9, 0xc0, 0xc4, 0xa8, 0x53, 0xee, 0x0b,
	0x11, 0xb9, 0x50, 0x1f, 0xe0, 0x8c, 0x3c, 0xe1, 0x99, 0x5b, 0x56, 0x16, 0x2d, 0xa2, 0x1b, 0x70,
	0xf0, 0x78, 0x1a, 0x60, 0x35, 0x5d, 0xb7, 0x22, 0x3b, 0x39, 0x5c, 0xd1, 0x89, 0x77, 0x3d, 0x9e,
	0xea, 0x1d, 0xa8, 0x39, 0x00, 0xce, 0x15, 0xad, 0x37, 0xb0, 0xb9, 0x64, 0x7e, 0x51, 0x83, 0x5f,
	0x2d, 0x40, 0xb7, 0x2c, 0xe4, 0x31, 0x65, 0x59, 0xa1, 0xd1, 0x7f, 0xa1, 0x41, 0xf9, 0xe4, 0x3c,
	0x2f, 0x4e, 0xc5, 0x72, 0x84, 0xce, 0x5c, 0x83, 0x82, 0x5c, 0xe4, 0x90, 0x52, 0x0e, 0xb9, 0x30,
	0x90, 0xff, 0x60, 0x23, 0xa2, 0x6c, 0x18, 0x44, 0x71, 0x0f, 0x47, 0x01, 0xe5, 0xa9, 0x5c, 0x97,
	0xed, 0x37, 0x84, 0xf6, 0x4e, 0x28, 0x3b, 0x3c, 0x6d, 0x7b, 0xb0, 0xf6, 0x40, 0x48, 0xd2, 0x61,
	0xfd, 0x58, 0x9c, 0x18, 0xc3, 0x23, 0x62, 0x4e, 0x4c, 0x7c, 0xa3, 0x0d, 0x28, 0x51, 0xae, 0xc3,
	0x97, 0x28, 0x6f, 0xff, 0xac, 0x40, 0xd3, 0x94, 0xfc, 0x48, 0xb2, 0x8c, 0xb2, 0x41, 0x8a, 0xae,
	0xc0, 0xa1, 0x1c, 0x8f, 0x02, 0x7d, 0x16, 0xc2, 0xdf, 0x39, 0xfd, 0x27, 0x1f, 0xe6, 0xf3, 0x16,
	0x7d, 0x10, 0x78, 0xdd, 0xee, 0x36, 0x54, 0x45, 0x49, 0xea, 0x06, 0x6d, 0x5f, 0x09, 0x62, 0x75,
	0x38, 0xa2, 0x38, 0x25, 0xa6, 0x6e, 0x23, 0x8a, 0xab, 0xd7, 0x91, 0x03, 0x1a, 0xba, 0x15, 0x59,
	0x9a, 0xad, 0x35, 0x9d, 0x10, 0xed, 0x83, 0x43, 0x74, 0x42, 0x61, 0xaf, 0x4a, 0x3b, 0x18, 0x55,
	0x27, 0x2c, 0x1e, 0x45, 0x6d, 0xf1, 0x28, 0xf6, 0x00, 0x28, 0xcf, 0x67, 0x5a, 0x57, 0x91, 0x29,
	0x37, 0x13, 0x6d, 0xc3, 0x3a, 0xe5, 0x41, 0x81, 0x71, 0x6b, 0x92, 0x71, 0x0e, 0xe5, 0x0f, 0x86,
	0x73, 0xf9, 0x62, 0x4c, 0x06, 0x7b, 0xbe, 0x98, 0x77, 0x3a, 0x8b, 0x07, 0x5b, 0x83, 0x28, 0xee,
	0xca, 0xa5, 0x14, 0x56, 0x08, 0x12, 0xf9, 0x97, 0x32, 0x75, 0x0a, 0x8b, 0x3c, 0x83, 0x9d, 0x22,
	0xbe, 0x90, 0xdf, 0x91, 0xf9, 0xb7, 0xe6, 0x2e, 0xf3, 0x3a, 0xf6, 0xc1, 0x19, 0xe1, 0x5e, 0x1e,
	0xbc, 0xa1, 0xa6, 0x30, 0xc2, 0x3d, 0x13, 0xf5, 0x3d, 0x38, 0xea, 0x15, 0x08, 0x62, 0x9e, 0xa5,
	0xee, 0xba, 0x24, 0xc0, 0xff, 0xcf, 0x76, 0x66, 0x76, 0xec, 0xdd, 0x48, 0xf0, 0x3d, 0xcf, 0x0c,
	0x0d, 0xc2, 0x5c, 0x21, 0x68, 0xb0, 0x64, 0x7e, 0x11, 0x0d, 0xbe, 0x5b, 0x50, 0xf9, 0x84, 0xd3,
	0xe1, 0xca, 0x03, 0x5c, 0x5a, 0x67, 0xe9, 0xd9, 0x3a, 0x17, 0x00, 0x5c, 0xf3, 0x7c, 0x0e, 0xe0,
	0xe8, 0x18, 0x2a, 0x94, 0xf5, 0x63, 0xcd, 0xf1, 0xbf, 0xf3, 0x16, 0x45, 0x4a, 0x4f, 0x1c, 0xbe,
	0x6a, 0x48, 0x82, 0x5a, 0xaf, 0xc0, 0xce, 0x55, 0x2f, 0xe5, 0xb2, 0xf3, 0x48, 0x92, 0x09, 0xed,
	0x11, 0x49, 0xa6, 0x26, 0x94, 0x27, 0x94, 0x1b, 0xdf, 0x09, 0xe5, 0xc2, 0x97, 0xc7, 0x49, 0x96,
	0xdf, 0xb9, 0x14, 0x24, 0x4d, 0x25, 0x43, 0xa3, 0x6e, 0x40, 0x59, 0x48, 0xa6, 0xb2, 0x83, 0xaa,
	0xdf, 0x90, 0xda, 0xbb, 0x6e, 0x47, 0xe8, 0xd0, 0x21, 0x54, 0x33, 0x9c, 0x0e, 0xcd, 0x43, 0xb5,
	0xbe, 0xd0, 0x84, 0xaf, 0x6c, 0xed, 0x1f, 0x16, 0x34, 0x3f, 0x2a, 0x3d, 0x65, 0x03, 0xcd, 0xae,
	0xcf, 0xd0, 0x34, 0xb3, 0x48, 0xe7, 0x04, 0x15, 0x41, 0xbc, 0x3c, 0xc8, 0xb2, 0x53, 0xbe, 0xfd,
	0x54, 0xc9, 0x6a, 0x40, 0x9b, 0x64, 0x51, 0xdb, 0xfa, 0x02, 0xdb, 0xab, 0x80, 0x2b, 0xc6, 0x76,
	0x52, 0x1c, 0x9b, 0x73, 0xba, 0xfb, 0xc7, 0x33, 0x2b, 0x4e, 0xf4, 0x18, 0x36, 0xf5, 0x6b, 0x41,
	0xfa, 0x24, 0x21, 0xac, 0x27, 0xdf, 0x73, 0xed, 0xa9, 0xa3, 0x1b, 0xb1, 0x5b, 0x93, 0xff, 0x9c,
	0x67, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x01, 0xe9, 0x84, 0x5a, 0x07, 0x00, 0x00,
}