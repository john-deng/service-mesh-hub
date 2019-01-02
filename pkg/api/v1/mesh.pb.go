// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mesh.proto

package v1 // import "github.com/solo-io/supergloo/pkg/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
// @solo-kit:resource.short_name=mesh
// @solo-kit:resource.plural_name=meshes
// @solo-kit:resource.resource_groups=translator.supergloo.solo.io
type Mesh struct {
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata"`
	// mesh-specific configuration
	//
	// Types that are valid to be assigned to MeshType:
	//	*Mesh_Istio
	//	*Mesh_Linkerd2
	//	*Mesh_Consul
	//	*Mesh_AppMesh
	MeshType isMesh_MeshType `protobuf_oneof:"mesh_type"`
	// policy applied to the mesh
	// TODO: rick-ducott, yuval-k: consider splitting these out as in routing.proto
	Encryption           *Encryption    `protobuf:"bytes,98,opt,name=encryption" json:"encryption,omitempty"`
	Observability        *Observability `protobuf:"bytes,99,opt,name=observability" json:"observability,omitempty"`
	Policy               *Policy        `protobuf:"bytes,100,opt,name=policy" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Mesh) Reset()         { *m = Mesh{} }
func (m *Mesh) String() string { return proto.CompactTextString(m) }
func (*Mesh) ProtoMessage()    {}
func (*Mesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_3ccd824e12b34443, []int{0}
}
func (m *Mesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mesh.Unmarshal(m, b)
}
func (m *Mesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mesh.Marshal(b, m, deterministic)
}
func (dst *Mesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mesh.Merge(dst, src)
}
func (m *Mesh) XXX_Size() int {
	return xxx_messageInfo_Mesh.Size(m)
}
func (m *Mesh) XXX_DiscardUnknown() {
	xxx_messageInfo_Mesh.DiscardUnknown(m)
}

var xxx_messageInfo_Mesh proto.InternalMessageInfo

type isMesh_MeshType interface {
	isMesh_MeshType()
	Equal(interface{}) bool
}

type Mesh_Istio struct {
	Istio *Istio `protobuf:"bytes,10,opt,name=istio,oneof"`
}
type Mesh_Linkerd2 struct {
	Linkerd2 *Linkerd2 `protobuf:"bytes,20,opt,name=linkerd2,oneof"`
}
type Mesh_Consul struct {
	Consul *Consul `protobuf:"bytes,30,opt,name=consul,oneof"`
}
type Mesh_AppMesh struct {
	AppMesh *AppMesh `protobuf:"bytes,50,opt,name=app_mesh,json=appMesh,oneof"`
}

func (*Mesh_Istio) isMesh_MeshType()    {}
func (*Mesh_Linkerd2) isMesh_MeshType() {}
func (*Mesh_Consul) isMesh_MeshType()   {}
func (*Mesh_AppMesh) isMesh_MeshType()  {}

func (m *Mesh) GetMeshType() isMesh_MeshType {
	if m != nil {
		return m.MeshType
	}
	return nil
}

func (m *Mesh) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Mesh) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Mesh) GetIstio() *Istio {
	if x, ok := m.GetMeshType().(*Mesh_Istio); ok {
		return x.Istio
	}
	return nil
}

func (m *Mesh) GetLinkerd2() *Linkerd2 {
	if x, ok := m.GetMeshType().(*Mesh_Linkerd2); ok {
		return x.Linkerd2
	}
	return nil
}

func (m *Mesh) GetConsul() *Consul {
	if x, ok := m.GetMeshType().(*Mesh_Consul); ok {
		return x.Consul
	}
	return nil
}

func (m *Mesh) GetAppMesh() *AppMesh {
	if x, ok := m.GetMeshType().(*Mesh_AppMesh); ok {
		return x.AppMesh
	}
	return nil
}

func (m *Mesh) GetEncryption() *Encryption {
	if m != nil {
		return m.Encryption
	}
	return nil
}

func (m *Mesh) GetObservability() *Observability {
	if m != nil {
		return m.Observability
	}
	return nil
}

func (m *Mesh) GetPolicy() *Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Mesh) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Mesh_OneofMarshaler, _Mesh_OneofUnmarshaler, _Mesh_OneofSizer, []interface{}{
		(*Mesh_Istio)(nil),
		(*Mesh_Linkerd2)(nil),
		(*Mesh_Consul)(nil),
		(*Mesh_AppMesh)(nil),
	}
}

func _Mesh_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Mesh)
	// mesh_type
	switch x := m.MeshType.(type) {
	case *Mesh_Istio:
		_ = b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Istio); err != nil {
			return err
		}
	case *Mesh_Linkerd2:
		_ = b.EncodeVarint(20<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Linkerd2); err != nil {
			return err
		}
	case *Mesh_Consul:
		_ = b.EncodeVarint(30<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Consul); err != nil {
			return err
		}
	case *Mesh_AppMesh:
		_ = b.EncodeVarint(50<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AppMesh); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Mesh.MeshType has unexpected type %T", x)
	}
	return nil
}

func _Mesh_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Mesh)
	switch tag {
	case 10: // mesh_type.istio
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Istio)
		err := b.DecodeMessage(msg)
		m.MeshType = &Mesh_Istio{msg}
		return true, err
	case 20: // mesh_type.linkerd2
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Linkerd2)
		err := b.DecodeMessage(msg)
		m.MeshType = &Mesh_Linkerd2{msg}
		return true, err
	case 30: // mesh_type.consul
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Consul)
		err := b.DecodeMessage(msg)
		m.MeshType = &Mesh_Consul{msg}
		return true, err
	case 50: // mesh_type.app_mesh
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AppMesh)
		err := b.DecodeMessage(msg)
		m.MeshType = &Mesh_AppMesh{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Mesh_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Mesh)
	// mesh_type
	switch x := m.MeshType.(type) {
	case *Mesh_Istio:
		s := proto.Size(x.Istio)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Mesh_Linkerd2:
		s := proto.Size(x.Linkerd2)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Mesh_Consul:
		s := proto.Size(x.Consul)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Mesh_AppMesh:
		s := proto.Size(x.AppMesh)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// configuration for an istio mesh. this will be autogenerated if Supergloo installs Istio for you.
type Istio struct {
	// which namespace is istio installed to?
	InstallationNamespace string `protobuf:"bytes,1,opt,name=installation_namespace,json=installationNamespace,proto3" json:"installation_namespace,omitempty"`
	// the namespaces istio is watching for its crd-based configuration. leave empty if istio install is cluster-wide
	WatchNamespaces []string `protobuf:"bytes,2,rep,name=watch_namespaces,json=watchNamespaces" json:"watch_namespaces,omitempty"`
	// if provided, this will give Supergloo a reference to the prometheus configuration associated with this istio install
	// if empty, Supergloo will look for the configmap `istio-system.prometheus`
	PrometheusConfigmap  *core.ResourceRef `protobuf:"bytes,3,opt,name=prometheus_configmap,json=prometheusConfigmap" json:"prometheus_configmap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Istio) Reset()         { *m = Istio{} }
func (m *Istio) String() string { return proto.CompactTextString(m) }
func (*Istio) ProtoMessage()    {}
func (*Istio) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_3ccd824e12b34443, []int{1}
}
func (m *Istio) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Istio.Unmarshal(m, b)
}
func (m *Istio) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Istio.Marshal(b, m, deterministic)
}
func (dst *Istio) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Istio.Merge(dst, src)
}
func (m *Istio) XXX_Size() int {
	return xxx_messageInfo_Istio.Size(m)
}
func (m *Istio) XXX_DiscardUnknown() {
	xxx_messageInfo_Istio.DiscardUnknown(m)
}

var xxx_messageInfo_Istio proto.InternalMessageInfo

func (m *Istio) GetInstallationNamespace() string {
	if m != nil {
		return m.InstallationNamespace
	}
	return ""
}

func (m *Istio) GetWatchNamespaces() []string {
	if m != nil {
		return m.WatchNamespaces
	}
	return nil
}

func (m *Istio) GetPrometheusConfigmap() *core.ResourceRef {
	if m != nil {
		return m.PrometheusConfigmap
	}
	return nil
}

// configuration for an linkerd2 mesh. this will be autogenerated if Supergloo installs Linkerd2 for you.
type Linkerd2 struct {
	// which namespace is linkerd2 installed to?
	InstallationNamespace string `protobuf:"bytes,1,opt,name=installation_namespace,json=installationNamespace,proto3" json:"installation_namespace,omitempty"`
	// the namespaces linkerd2 is watching for its crd-based configuration. leave empty if linkerd2 install is cluster-wide
	WatchNamespaces []string `protobuf:"bytes,2,rep,name=watch_namespaces,json=watchNamespaces" json:"watch_namespaces,omitempty"`
	// if provided, this will give Supergloo a reference to the prometheus configuration associated with this linkerd2 install
	// if empty, Supergloo will look for the configmap `linkerd.prometheus`
	PrometheusConfigmap  *core.ResourceRef `protobuf:"bytes,3,opt,name=prometheus_configmap,json=prometheusConfigmap" json:"prometheus_configmap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Linkerd2) Reset()         { *m = Linkerd2{} }
func (m *Linkerd2) String() string { return proto.CompactTextString(m) }
func (*Linkerd2) ProtoMessage()    {}
func (*Linkerd2) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_3ccd824e12b34443, []int{2}
}
func (m *Linkerd2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Linkerd2.Unmarshal(m, b)
}
func (m *Linkerd2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Linkerd2.Marshal(b, m, deterministic)
}
func (dst *Linkerd2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Linkerd2.Merge(dst, src)
}
func (m *Linkerd2) XXX_Size() int {
	return xxx_messageInfo_Linkerd2.Size(m)
}
func (m *Linkerd2) XXX_DiscardUnknown() {
	xxx_messageInfo_Linkerd2.DiscardUnknown(m)
}

var xxx_messageInfo_Linkerd2 proto.InternalMessageInfo

func (m *Linkerd2) GetInstallationNamespace() string {
	if m != nil {
		return m.InstallationNamespace
	}
	return ""
}

func (m *Linkerd2) GetWatchNamespaces() []string {
	if m != nil {
		return m.WatchNamespaces
	}
	return nil
}

func (m *Linkerd2) GetPrometheusConfigmap() *core.ResourceRef {
	if m != nil {
		return m.PrometheusConfigmap
	}
	return nil
}

// configuration for an consul mesh. this will be autogenerated if Supergloo installs Consul for you.
type Consul struct {
	// which namespace is consul instatlled to?
	InstallationNamespace string `protobuf:"bytes,1,opt,name=installation_namespace,json=installationNamespace,proto3" json:"installation_namespace,omitempty"`
	// address of the consul api server
	ServerAddress string `protobuf:"bytes,2,opt,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"`
	// if provided, this will give Supergloo a reference to the prometheus configuration associated with this consul install
	// if empty, Supergloo will look for the configmap `linkerd.prometheus`
	PrometheusConfigmap  *core.ResourceRef `protobuf:"bytes,3,opt,name=prometheus_configmap,json=prometheusConfigmap" json:"prometheus_configmap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Consul) Reset()         { *m = Consul{} }
func (m *Consul) String() string { return proto.CompactTextString(m) }
func (*Consul) ProtoMessage()    {}
func (*Consul) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_3ccd824e12b34443, []int{3}
}
func (m *Consul) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consul.Unmarshal(m, b)
}
func (m *Consul) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consul.Marshal(b, m, deterministic)
}
func (dst *Consul) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consul.Merge(dst, src)
}
func (m *Consul) XXX_Size() int {
	return xxx_messageInfo_Consul.Size(m)
}
func (m *Consul) XXX_DiscardUnknown() {
	xxx_messageInfo_Consul.DiscardUnknown(m)
}

var xxx_messageInfo_Consul proto.InternalMessageInfo

func (m *Consul) GetInstallationNamespace() string {
	if m != nil {
		return m.InstallationNamespace
	}
	return ""
}

func (m *Consul) GetServerAddress() string {
	if m != nil {
		return m.ServerAddress
	}
	return ""
}

func (m *Consul) GetPrometheusConfigmap() *core.ResourceRef {
	if m != nil {
		return m.PrometheusConfigmap
	}
	return nil
}

// configuration for an AWS AppMesh mesh.
type AppMesh struct {
	// the region where this mesh should be deployed
	AwsRegion string `protobuf:"bytes,1,opt,name=aws_region,json=awsRegion,proto3" json:"aws_region,omitempty"`
	// ref to a Secret containing aws credentials
	AwsCredentials       *core.ResourceRef `protobuf:"bytes,3,opt,name=aws_credentials,json=awsCredentials" json:"aws_credentials,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AppMesh) Reset()         { *m = AppMesh{} }
func (m *AppMesh) String() string { return proto.CompactTextString(m) }
func (*AppMesh) ProtoMessage()    {}
func (*AppMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_3ccd824e12b34443, []int{4}
}
func (m *AppMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppMesh.Unmarshal(m, b)
}
func (m *AppMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppMesh.Marshal(b, m, deterministic)
}
func (dst *AppMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppMesh.Merge(dst, src)
}
func (m *AppMesh) XXX_Size() int {
	return xxx_messageInfo_AppMesh.Size(m)
}
func (m *AppMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_AppMesh.DiscardUnknown(m)
}

var xxx_messageInfo_AppMesh proto.InternalMessageInfo

func (m *AppMesh) GetAwsRegion() string {
	if m != nil {
		return m.AwsRegion
	}
	return ""
}

func (m *AppMesh) GetAwsCredentials() *core.ResourceRef {
	if m != nil {
		return m.AwsCredentials
	}
	return nil
}

func init() {
	proto.RegisterType((*Mesh)(nil), "supergloo.solo.io.Mesh")
	proto.RegisterType((*Istio)(nil), "supergloo.solo.io.Istio")
	proto.RegisterType((*Linkerd2)(nil), "supergloo.solo.io.Linkerd2")
	proto.RegisterType((*Consul)(nil), "supergloo.solo.io.Consul")
	proto.RegisterType((*AppMesh)(nil), "supergloo.solo.io.AppMesh")
}
func (this *Mesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Mesh)
	if !ok {
		that2, ok := that.(Mesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if that1.MeshType == nil {
		if this.MeshType != nil {
			return false
		}
	} else if this.MeshType == nil {
		return false
	} else if !this.MeshType.Equal(that1.MeshType) {
		return false
	}
	if !this.Encryption.Equal(that1.Encryption) {
		return false
	}
	if !this.Observability.Equal(that1.Observability) {
		return false
	}
	if !this.Policy.Equal(that1.Policy) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Mesh_Istio) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Mesh_Istio)
	if !ok {
		that2, ok := that.(Mesh_Istio)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Istio.Equal(that1.Istio) {
		return false
	}
	return true
}
func (this *Mesh_Linkerd2) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Mesh_Linkerd2)
	if !ok {
		that2, ok := that.(Mesh_Linkerd2)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Linkerd2.Equal(that1.Linkerd2) {
		return false
	}
	return true
}
func (this *Mesh_Consul) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Mesh_Consul)
	if !ok {
		that2, ok := that.(Mesh_Consul)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Consul.Equal(that1.Consul) {
		return false
	}
	return true
}
func (this *Mesh_AppMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Mesh_AppMesh)
	if !ok {
		that2, ok := that.(Mesh_AppMesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.AppMesh.Equal(that1.AppMesh) {
		return false
	}
	return true
}
func (this *Istio) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Istio)
	if !ok {
		that2, ok := that.(Istio)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.InstallationNamespace != that1.InstallationNamespace {
		return false
	}
	if len(this.WatchNamespaces) != len(that1.WatchNamespaces) {
		return false
	}
	for i := range this.WatchNamespaces {
		if this.WatchNamespaces[i] != that1.WatchNamespaces[i] {
			return false
		}
	}
	if !this.PrometheusConfigmap.Equal(that1.PrometheusConfigmap) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Linkerd2) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Linkerd2)
	if !ok {
		that2, ok := that.(Linkerd2)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.InstallationNamespace != that1.InstallationNamespace {
		return false
	}
	if len(this.WatchNamespaces) != len(that1.WatchNamespaces) {
		return false
	}
	for i := range this.WatchNamespaces {
		if this.WatchNamespaces[i] != that1.WatchNamespaces[i] {
			return false
		}
	}
	if !this.PrometheusConfigmap.Equal(that1.PrometheusConfigmap) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Consul) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Consul)
	if !ok {
		that2, ok := that.(Consul)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.InstallationNamespace != that1.InstallationNamespace {
		return false
	}
	if this.ServerAddress != that1.ServerAddress {
		return false
	}
	if !this.PrometheusConfigmap.Equal(that1.PrometheusConfigmap) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AppMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AppMesh)
	if !ok {
		that2, ok := that.(AppMesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.AwsRegion != that1.AwsRegion {
		return false
	}
	if !this.AwsCredentials.Equal(that1.AwsCredentials) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("mesh.proto", fileDescriptor_mesh_3ccd824e12b34443) }

var fileDescriptor_mesh_3ccd824e12b34443 = []byte{
	// 626 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x94, 0xcd, 0x6a, 0x1b, 0x3b,
	0x14, 0x80, 0x3d, 0x37, 0x89, 0x63, 0x9f, 0xdc, 0xfc, 0x29, 0xbe, 0x41, 0xc9, 0x25, 0x89, 0x31,
	0x5c, 0x6e, 0x0a, 0xcd, 0x4c, 0x93, 0x50, 0xfa, 0x03, 0x5d, 0xc4, 0xa1, 0xad, 0x0b, 0x49, 0x5b,
	0xd4, 0x5d, 0x37, 0x46, 0x9e, 0x91, 0xc7, 0x22, 0xe3, 0x91, 0x90, 0xe4, 0x98, 0xbc, 0x51, 0xa1,
	0xdb, 0x42, 0xa1, 0x4f, 0xd0, 0xa7, 0xc8, 0xa2, 0x8f, 0xd0, 0x27, 0x28, 0xd2, 0xc8, 0x76, 0x4c,
	0x5d, 0x1a, 0x4a, 0x37, 0x5d, 0x79, 0x38, 0xe7, 0xfb, 0xce, 0x1c, 0xcd, 0x39, 0x16, 0x40, 0x9f,
	0xe9, 0x5e, 0x28, 0x95, 0x30, 0x02, 0xad, 0xeb, 0x81, 0x64, 0x2a, 0xcd, 0x84, 0x08, 0xb5, 0xc8,
	0x44, 0xc8, 0xc5, 0x76, 0x2d, 0x15, 0xa9, 0x70, 0xd9, 0xc8, 0x3e, 0x15, 0xe0, 0xf6, 0x61, 0xca,
	0x4d, 0x6f, 0xd0, 0x09, 0x63, 0xd1, 0x8f, 0x2c, 0x79, 0xc0, 0x45, 0xf1, 0x7b, 0xc1, 0x4d, 0x44,
	0x25, 0x8f, 0x2e, 0x0f, 0xa3, 0x3e, 0x33, 0x34, 0xa1, 0x86, 0x7a, 0x25, 0xba, 0x85, 0xa2, 0x0d,
	0x35, 0x03, 0xed, 0x85, 0xbb, 0xb7, 0x10, 0x14, 0xeb, 0x7a, 0x7a, 0x43, 0x74, 0x34, 0x53, 0x97,
	0xb4, 0xc3, 0x33, 0x6e, 0xae, 0x7c, 0x70, 0x8d, 0xe5, 0xb1, 0xba, 0x92, 0x86, 0x8b, 0xdc, 0x47,
	0xfe, 0x96, 0x22, 0xe3, 0xb1, 0xcf, 0x37, 0xde, 0xcf, 0xc3, 0xfc, 0x39, 0xd3, 0x3d, 0xf4, 0x1c,
	0xca, 0xc5, 0xbb, 0x71, 0xb9, 0x1e, 0xec, 0x2f, 0x1d, 0xd5, 0xc2, 0x58, 0x28, 0x36, 0xfa, 0x08,
	0xe1, 0x1b, 0x97, 0x6b, 0x6e, 0x7d, 0xbe, 0xde, 0x2b, 0x7d, 0xbd, 0xde, 0x5b, 0x37, 0x4c, 0x9b,
	0x84, 0x77, 0xbb, 0x8f, 0x1b, 0x3c, 0xcd, 0x85, 0x62, 0x0d, 0xe2, 0x75, 0xf4, 0x10, 0x2a, 0xa3,
	0x73, 0xe3, 0x45, 0x57, 0x6a, 0x73, 0xba, 0xd4, 0xb9, 0xcf, 0x36, 0xe7, 0x6d, 0x31, 0x32, 0xa6,
	0xd1, 0x3d, 0x58, 0xe0, 0xda, 0x70, 0x81, 0xc1, 0x69, 0x38, 0xfc, 0x6e, 0x16, 0xe1, 0x0b, 0x9b,
	0x6f, 0x95, 0x48, 0x01, 0xa2, 0x47, 0x50, 0xc9, 0x78, 0x7e, 0xc1, 0x54, 0x72, 0x84, 0x6b, 0x4e,
	0xfa, 0x77, 0x86, 0x74, 0xe6, 0x91, 0x56, 0x89, 0x8c, 0x71, 0x74, 0x0c, 0xe5, 0x58, 0xe4, 0x7a,
	0x90, 0xe1, 0x5d, 0x27, 0x6e, 0xcd, 0x10, 0x4f, 0x1d, 0xd0, 0x2a, 0x11, 0x8f, 0xa2, 0x07, 0x50,
	0xa1, 0x52, 0xb6, 0xed, 0xbe, 0xe0, 0x23, 0xa7, 0x6d, 0xcf, 0xd0, 0x4e, 0xa4, 0xb4, 0x9f, 0xb4,
	0x55, 0x22, 0x8b, 0xb4, 0x78, 0x44, 0x4f, 0x00, 0x26, 0x83, 0xc0, 0x1d, 0xa7, 0xee, 0xcc, 0x50,
	0x9f, 0x8e, 0x21, 0x72, 0x43, 0x40, 0xcf, 0x60, 0x79, 0x6a, 0xb8, 0x38, 0x76, 0x15, 0xea, 0x33,
	0x2a, 0xbc, 0xba, 0xc9, 0x91, 0x69, 0x0d, 0x1d, 0x42, 0xb9, 0x98, 0x3e, 0x4e, 0x7e, 0x78, 0xe8,
	0xd7, 0x0e, 0x20, 0x1e, 0x6c, 0x2e, 0x41, 0xd5, 0x1e, 0xb7, 0x6d, 0xae, 0x24, 0x6b, 0x7c, 0x0c,
	0x60, 0xc1, 0x8d, 0x00, 0xdd, 0x87, 0x4d, 0x9e, 0x6b, 0x43, 0xb3, 0x8c, 0xda, 0x0e, 0xdb, 0x39,
	0xed, 0x33, 0x2d, 0x69, 0xcc, 0x70, 0x50, 0x0f, 0xf6, 0xab, 0xe4, 0x9f, 0x9b, 0xd9, 0x97, 0xa3,
	0x24, 0xba, 0x03, 0x6b, 0x43, 0x6a, 0xe2, 0xde, 0x84, 0xd7, 0xf8, 0xaf, 0xfa, 0xdc, 0x7e, 0x95,
	0xac, 0xba, 0xf8, 0x98, 0xd4, 0xe8, 0x0c, 0x6a, 0x52, 0x89, 0x3e, 0x33, 0x3d, 0x36, 0xd0, 0xed,
	0x58, 0xe4, 0x5d, 0x9e, 0xf6, 0xa9, 0xc4, 0x73, 0xbe, 0xf3, 0xa9, 0x9d, 0x22, 0x4c, 0x8b, 0x81,
	0x8a, 0x19, 0x61, 0x5d, 0xb2, 0x31, 0xd1, 0x4e, 0x47, 0x56, 0xe3, 0x53, 0x00, 0x95, 0xd1, 0x1e,
	0xfc, 0x71, 0xcd, 0x7f, 0x08, 0xa0, 0x5c, 0xec, 0xe2, 0xaf, 0xb6, 0xfe, 0x1f, 0xac, 0xd8, 0x3d,
	0x60, 0xaa, 0x4d, 0x93, 0x44, 0x31, 0x6d, 0x1b, 0xb7, 0xf8, 0x72, 0x11, 0x3d, 0x29, 0x82, 0xbf,
	0xb9, 0xed, 0x0c, 0x16, 0xfd, 0x5f, 0x01, 0xed, 0x00, 0xd0, 0xa1, 0x6e, 0x2b, 0x96, 0xda, 0xfd,
	0x2f, 0x5a, 0xad, 0xd2, 0xa1, 0x26, 0x2e, 0x80, 0x9a, 0xb0, 0x6a, 0xd3, 0xb1, 0x62, 0x09, 0xcb,
	0x0d, 0xa7, 0x99, 0xfe, 0xf9, 0x2b, 0x57, 0xe8, 0x50, 0x9f, 0x4e, 0x84, 0xe6, 0xc1, 0xbb, 0x2f,
	0xbb, 0xc1, 0xdb, 0xff, 0x67, 0x5d, 0x99, 0xa3, 0x1d, 0x8f, 0xe4, 0x45, 0xea, 0xef, 0xcd, 0x4e,
	0xd9, 0xdd, 0x7f, 0xc7, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x63, 0x86, 0xb9, 0x33, 0xfd, 0x05,
	0x00, 0x00,
}