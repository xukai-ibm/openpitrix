// Code generated by protoc-gen-go. DO NOT EDIT.
// source: openpitrix/types/etcd.proto

package pbtypes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EtcdConfig struct {
	User     string          `protobuf:"bytes,1,opt,name=user" json:"user"`
	Password string          `protobuf:"bytes,2,opt,name=password" json:"password"`
	NodeList []*EtcdEndpoint `protobuf:"bytes,3,rep,name=node_list,json=nodeList" json:"node_list"`
}

func (m *EtcdConfig) Reset()                    { *m = EtcdConfig{} }
func (m *EtcdConfig) String() string            { return proto.CompactTextString(m) }
func (*EtcdConfig) ProtoMessage()               {}
func (*EtcdConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *EtcdConfig) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *EtcdConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *EtcdConfig) GetNodeList() []*EtcdEndpoint {
	if m != nil {
		return m.NodeList
	}
	return nil
}

type EtcdEndpoint struct {
	Host string `protobuf:"bytes,1,opt,name=host" json:"host"`
	Port int32  `protobuf:"varint,2,opt,name=port" json:"port"`
}

func (m *EtcdEndpoint) Reset()                    { *m = EtcdEndpoint{} }
func (m *EtcdEndpoint) String() string            { return proto.CompactTextString(m) }
func (*EtcdEndpoint) ProtoMessage()               {}
func (*EtcdEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *EtcdEndpoint) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *EtcdEndpoint) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*EtcdConfig)(nil), "openpitrix.types.EtcdConfig")
	proto.RegisterType((*EtcdEndpoint)(nil), "openpitrix.types.EtcdEndpoint")
}

func init() { proto.RegisterFile("openpitrix/types/etcd.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x31, 0x6b, 0xc3, 0x30,
	0x10, 0x85, 0x71, 0xdd, 0x16, 0x5b, 0xed, 0x50, 0x34, 0x99, 0x16, 0x8a, 0xf1, 0xe4, 0xa5, 0x12,
	0xb4, 0xd0, 0xc5, 0x5b, 0x8b, 0xb7, 0x4e, 0x1e, 0xbb, 0x84, 0xd8, 0x52, 0x1c, 0x91, 0xa0, 0x3b,
	0xa4, 0x0b, 0x89, 0xff, 0x7d, 0x90, 0x42, 0x12, 0x93, 0x49, 0x4f, 0xef, 0x71, 0x7c, 0x7c, 0xec,
	0x0d, 0x50, 0x5b, 0x34, 0xe4, 0xcc, 0x41, 0xd2, 0x84, 0xda, 0x4b, 0x4d, 0x83, 0x12, 0xe8, 0x80,
	0x80, 0xbf, 0x5c, 0x47, 0x11, 0xc7, 0x6a, 0x62, 0xac, 0xa5, 0x41, 0xfd, 0x82, 0x5d, 0x99, 0x91,
	0x73, 0x76, 0xbf, 0xf3, 0xda, 0x15, 0x49, 0x99, 0xd4, 0x79, 0x17, 0x33, 0x7f, 0x65, 0x19, 0x2e,
	0xbd, 0xdf, 0x83, 0x53, 0xc5, 0x5d, 0xec, 0x2f, 0x7f, 0xde, 0xb0, 0xdc, 0x82, 0xd2, 0x8b, 0xad,
	0xf1, 0x54, 0xa4, 0x65, 0x5a, 0x3f, 0x7d, 0xbe, 0x8b, 0x5b, 0x86, 0x08, 0x80, 0xd6, 0x2a, 0x04,
	0x63, 0xa9, 0xcb, 0xc2, 0xc1, 0x9f, 0xf1, 0x54, 0x7d, 0xb3, 0xe7, 0xf9, 0x12, 0xe0, 0x6b, 0xf0,
	0x74, 0x86, 0x87, 0x1c, 0x3a, 0x04, 0x47, 0x11, 0xfc, 0xd0, 0xc5, 0xfc, 0x23, 0xff, 0x3f, 0x66,
	0x08, 0x03, 0x72, 0x66, 0x8c, 0x9b, 0x51, 0x62, 0x7f, 0x12, 0x6f, 0xb0, 0x8f, 0x6f, 0xff, 0x18,
	0xe5, 0xbf, 0x8e, 0x01, 0x00, 0x00, 0xff, 0xff, 0x43, 0xe7, 0x5b, 0x3e, 0x1b, 0x01, 0x00, 0x00,
}
