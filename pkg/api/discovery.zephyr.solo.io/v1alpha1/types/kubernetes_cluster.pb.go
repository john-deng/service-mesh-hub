// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/discovery/v1alpha1/kubernetes_cluster.proto

package types

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//Representation of a Kubernetes cluster that has been registered in Service Mesh Hub.
type KubernetesClusterSpec struct {
	// pointer to secret which contains the kubeconfig with information to connect to the remote cluster.
	SecretRef *types.ResourceRef `protobuf:"bytes,1,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref,omitempty"`
	// context to use within the kubeconfig pointed to by the above reference
	Context string `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	// version of kubernetes
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// cloud provider, empty if unknown
	Cloud string `protobuf:"bytes,4,opt,name=cloud,proto3" json:"cloud,omitempty"`
	// namespace to use when writing Service Mesh Hub resources to this cluster
	WriteNamespace       string   `protobuf:"bytes,5,opt,name=write_namespace,json=writeNamespace,proto3" json:"write_namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KubernetesClusterSpec) Reset()         { *m = KubernetesClusterSpec{} }
func (m *KubernetesClusterSpec) String() string { return proto.CompactTextString(m) }
func (*KubernetesClusterSpec) ProtoMessage()    {}
func (*KubernetesClusterSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_cebb6fc6a7744031, []int{0}
}
func (m *KubernetesClusterSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesClusterSpec.Unmarshal(m, b)
}
func (m *KubernetesClusterSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesClusterSpec.Marshal(b, m, deterministic)
}
func (m *KubernetesClusterSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesClusterSpec.Merge(m, src)
}
func (m *KubernetesClusterSpec) XXX_Size() int {
	return xxx_messageInfo_KubernetesClusterSpec.Size(m)
}
func (m *KubernetesClusterSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesClusterSpec.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesClusterSpec proto.InternalMessageInfo

func (m *KubernetesClusterSpec) GetSecretRef() *types.ResourceRef {
	if m != nil {
		return m.SecretRef
	}
	return nil
}

func (m *KubernetesClusterSpec) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

func (m *KubernetesClusterSpec) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *KubernetesClusterSpec) GetCloud() string {
	if m != nil {
		return m.Cloud
	}
	return ""
}

func (m *KubernetesClusterSpec) GetWriteNamespace() string {
	if m != nil {
		return m.WriteNamespace
	}
	return ""
}

func init() {
	proto.RegisterType((*KubernetesClusterSpec)(nil), "discovery.zephyr.solo.io.KubernetesClusterSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/discovery/v1alpha1/kubernetes_cluster.proto", fileDescriptor_cebb6fc6a7744031)
}

var fileDescriptor_cebb6fc6a7744031 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x18, 0x84, 0x89, 0x5a, 0xa5, 0x2b, 0x28, 0x04, 0x85, 0xc5, 0x53, 0xf1, 0x62, 0x0f, 0x76, 0x43,
	0xf5, 0x01, 0x04, 0x3d, 0x0a, 0x1e, 0x52, 0xbc, 0x78, 0x29, 0xc9, 0x76, 0xd2, 0x2c, 0x4d, 0xf3,
	0x2f, 0xff, 0xee, 0x46, 0xeb, 0x2b, 0xfa, 0x52, 0xd2, 0xc4, 0xb6, 0x58, 0x3c, 0x78, 0x9c, 0x99,
	0xdd, 0x6f, 0xf8, 0x47, 0x4c, 0xe6, 0xc6, 0x97, 0x21, 0x57, 0x9a, 0x96, 0x89, 0xa3, 0x8a, 0x46,
	0x86, 0x12, 0x07, 0x6e, 0x8c, 0xc6, 0x68, 0x09, 0x57, 0x8e, 0xca, 0x90, 0x27, 0x99, 0x35, 0xc9,
	0xcc, 0x38, 0x4d, 0x0d, 0x78, 0x95, 0x34, 0xe3, 0xac, 0xb2, 0x65, 0x36, 0x4e, 0x16, 0x21, 0x07,
	0xd7, 0xf0, 0x70, 0x53, 0x5d, 0x05, 0xe7, 0xc1, 0xca, 0x32, 0x79, 0x8a, 0xe5, 0xf6, 0xb1, 0xfa,
	0x84, 0x2d, 0x57, 0xac, 0xd6, 0x68, 0x65, 0xe8, 0xea, 0xf6, 0x4f, 0xb6, 0x26, 0xc6, 0x0e, 0xcb,
	0x28, 0x3a, 0xce, 0xf5, 0x57, 0x24, 0x2e, 0x9f, 0xb7, 0x25, 0x4f, 0x5d, 0xc7, 0xc4, 0x42, 0xc7,
	0x0f, 0x42, 0x38, 0x68, 0x86, 0x9f, 0x32, 0x0a, 0x19, 0x0d, 0xa2, 0xe1, 0xe9, 0xdd, 0x40, 0xad,
	0x39, 0x7b, 0x8d, 0x2a, 0x85, 0xa3, 0xc0, 0x1a, 0x29, 0x8a, 0xb4, 0xdf, 0xfd, 0x49, 0x51, 0xc4,
	0x52, 0x9c, 0x68, 0xaa, 0x3d, 0x3e, 0xbc, 0x3c, 0x18, 0x44, 0xc3, 0x7e, 0xba, 0x91, 0xeb, 0xa4,
	0x01, 0x3b, 0x43, 0xb5, 0x3c, 0xec, 0x92, 0x1f, 0x19, 0x5f, 0x88, 0x9e, 0xae, 0x28, 0xcc, 0xe4,
	0x51, 0xeb, 0x77, 0x22, 0xbe, 0x11, 0xe7, 0xef, 0x6c, 0x3c, 0xa6, 0x75, 0xb6, 0x84, 0xb3, 0x99,
	0x86, 0xec, 0xb5, 0xf9, 0x59, 0x6b, 0xbf, 0x6c, 0xdc, 0xc7, 0xd7, 0xb7, 0x7f, 0x8d, 0x6d, 0x17,
	0xf3, 0xdf, 0x83, 0xef, 0x5d, 0xb4, 0x1b, 0xca, 0xaf, 0x2c, 0x5c, 0x7e, 0xdc, 0x6e, 0x75, 0xff,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x41, 0xa6, 0x60, 0xdf, 0xca, 0x01, 0x00, 0x00,
}
