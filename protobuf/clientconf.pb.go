// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: clientconf.proto

package tdproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegistrarType int32

const (
	// Prefix type names with `REGISTRAR_TYPE` to avoid namespace pollution
	RegistrarType_REGISTRAR_TYPE_NULL  RegistrarType = 0
	RegistrarType_REGISTRAR_TYPE_API   RegistrarType = 1
	RegistrarType_REGISTRAR_TYPE_DECOY RegistrarType = 2
	RegistrarType_REGISTRAR_TYPE_OTHER RegistrarType = 3
)

// Enum value maps for RegistrarType.
var (
	RegistrarType_name = map[int32]string{
		0: "REGISTRAR_TYPE_NULL",
		1: "REGISTRAR_TYPE_API",
		2: "REGISTRAR_TYPE_DECOY",
		3: "REGISTRAR_TYPE_OTHER",
	}
	RegistrarType_value = map[string]int32{
		"REGISTRAR_TYPE_NULL":  0,
		"REGISTRAR_TYPE_API":   1,
		"REGISTRAR_TYPE_DECOY": 2,
		"REGISTRAR_TYPE_OTHER": 3,
	}
)

func (x RegistrarType) Enum() *RegistrarType {
	p := new(RegistrarType)
	*p = x
	return p
}

func (x RegistrarType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegistrarType) Descriptor() protoreflect.EnumDescriptor {
	return file_clientconf_proto_enumTypes[0].Descriptor()
}

func (RegistrarType) Type() protoreflect.EnumType {
	return &file_clientconf_proto_enumTypes[0]
}

func (x RegistrarType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *RegistrarType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = RegistrarType(num)
	return nil
}

// Deprecated: Use RegistrarType.Descriptor instead.
func (RegistrarType) EnumDescriptor() ([]byte, []int) {
	return file_clientconf_proto_rawDescGZIP(), []int{0}
}

// Represents an individual config
type DeploymentConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DecoyList          *DecoyList          `protobuf:"bytes,1,opt,name=decoy_list,json=decoyList" json:"decoy_list,omitempty"`
	Generation         *uint32             `protobuf:"varint,2,opt,name=generation" json:"generation,omitempty"`
	DefaultPubkey      *PubKey             `protobuf:"bytes,3,opt,name=default_pubkey,json=defaultPubkey" json:"default_pubkey,omitempty"`
	PhantomSubnetsList *PhantomSubnetsList `protobuf:"bytes,4,opt,name=phantom_subnets_list,json=phantomSubnetsList" json:"phantom_subnets_list,omitempty"`
	ConjurePubkey      *PubKey             `protobuf:"bytes,5,opt,name=conjure_pubkey,json=conjurePubkey" json:"conjure_pubkey,omitempty"`
	VersionNumber      *uint32             `protobuf:"varint,6,opt,name=version_number,json=versionNumber" json:"version_number,omitempty"`
	SuportedTransports []TransportType     `protobuf:"varint,7,rep,name=suported_transports,json=suportedTransports,enum=tapdance.TransportType" json:"suported_transports,omitempty"`
	Registrars         []*Registrar        `protobuf:"bytes,8,rep,name=registrars" json:"registrars,omitempty"`
}

func (x *DeploymentConfig) Reset() {
	*x = DeploymentConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientconf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentConfig) ProtoMessage() {}

func (x *DeploymentConfig) ProtoReflect() protoreflect.Message {
	mi := &file_clientconf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentConfig.ProtoReflect.Descriptor instead.
func (*DeploymentConfig) Descriptor() ([]byte, []int) {
	return file_clientconf_proto_rawDescGZIP(), []int{0}
}

func (x *DeploymentConfig) GetDecoyList() *DecoyList {
	if x != nil {
		return x.DecoyList
	}
	return nil
}

func (x *DeploymentConfig) GetGeneration() uint32 {
	if x != nil && x.Generation != nil {
		return *x.Generation
	}
	return 0
}

func (x *DeploymentConfig) GetDefaultPubkey() *PubKey {
	if x != nil {
		return x.DefaultPubkey
	}
	return nil
}

func (x *DeploymentConfig) GetPhantomSubnetsList() *PhantomSubnetsList {
	if x != nil {
		return x.PhantomSubnetsList
	}
	return nil
}

func (x *DeploymentConfig) GetConjurePubkey() *PubKey {
	if x != nil {
		return x.ConjurePubkey
	}
	return nil
}

func (x *DeploymentConfig) GetVersionNumber() uint32 {
	if x != nil && x.VersionNumber != nil {
		return *x.VersionNumber
	}
	return 0
}

func (x *DeploymentConfig) GetSuportedTransports() []TransportType {
	if x != nil {
		return x.SuportedTransports
	}
	return nil
}

func (x *DeploymentConfig) GetRegistrars() []*Registrar {
	if x != nil {
		return x.Registrars
	}
	return nil
}

// Houses multiple deployment configs
type ClientConfig2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VersionNumber     *uint32             `protobuf:"varint,1,opt,name=version_number,json=versionNumber" json:"version_number,omitempty"`
	DeploymentConfigs []*DeploymentConfig `protobuf:"bytes,2,rep,name=deployment_configs,json=deploymentConfigs" json:"deployment_configs,omitempty"`
}

func (x *ClientConfig2) Reset() {
	*x = ClientConfig2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientconf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConfig2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConfig2) ProtoMessage() {}

func (x *ClientConfig2) ProtoReflect() protoreflect.Message {
	mi := &file_clientconf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConfig2.ProtoReflect.Descriptor instead.
func (*ClientConfig2) Descriptor() ([]byte, []int) {
	return file_clientconf_proto_rawDescGZIP(), []int{1}
}

func (x *ClientConfig2) GetVersionNumber() uint32 {
	if x != nil && x.VersionNumber != nil {
		return *x.VersionNumber
	}
	return 0
}

func (x *ClientConfig2) GetDeploymentConfigs() []*DeploymentConfig {
	if x != nil {
		return x.DeploymentConfigs
	}
	return nil
}

type Registrar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegistrarType *RegistrarType `protobuf:"varint,1,opt,name=registrar_type,json=registrarType,enum=tapdance.RegistrarType" json:"registrar_type,omitempty"`
	Registrar     *anypb.Any     `protobuf:"bytes,2,opt,name=registrar" json:"registrar,omitempty"` // Any: arbitrary serialized
}

func (x *Registrar) Reset() {
	*x = Registrar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientconf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registrar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registrar) ProtoMessage() {}

func (x *Registrar) ProtoReflect() protoreflect.Message {
	mi := &file_clientconf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registrar.ProtoReflect.Descriptor instead.
func (*Registrar) Descriptor() ([]byte, []int) {
	return file_clientconf_proto_rawDescGZIP(), []int{2}
}

func (x *Registrar) GetRegistrarType() RegistrarType {
	if x != nil && x.RegistrarType != nil {
		return *x.RegistrarType
	}
	return RegistrarType_REGISTRAR_TYPE_NULL
}

func (x *Registrar) GetRegistrar() *anypb.Any {
	if x != nil {
		return x.Registrar
	}
	return nil
}

var File_clientconf_proto protoreflect.FileDescriptor

var file_clientconf_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x74, 0x61, 0x70, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x03, 0x0a, 0x10, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32,
	0x0a, 0x0a, 0x64, 0x65, 0x63, 0x6f, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x61, 0x70, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x44, 0x65,
	0x63, 0x6f, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x09, 0x64, 0x65, 0x63, 0x6f, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x70, 0x75,
	0x62, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x61, 0x70,
	0x64, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x0d, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x4e, 0x0a, 0x14, 0x70,
	0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x5f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x61, 0x70, 0x64,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x50, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d, 0x53, 0x75, 0x62, 0x6e,
	0x65, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x12, 0x70, 0x68, 0x61, 0x6e, 0x74, 0x6f, 0x6d,
	0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0e, 0x63,
	0x6f, 0x6e, 0x6a, 0x75, 0x72, 0x65, 0x5f, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x61, 0x70, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x50,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x6a, 0x75, 0x72, 0x65, 0x50, 0x75,
	0x62, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x13, 0x73,
	0x75, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x74, 0x61, 0x70, 0x64, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x12, 0x73, 0x75, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x61, 0x70, 0x64,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x52, 0x0a,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x0d, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0x12, 0x25, 0x0a, 0x0e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x12, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x74, 0x61, 0x70, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x11, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0x7f,
	0x0a, 0x09, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x12, 0x3e, 0x0a, 0x0e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x74, 0x61, 0x70, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x2a,
	0x74, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x41, 0x52, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x47,
	0x49, 0x53, 0x54, 0x52, 0x41, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x50, 0x49, 0x10,
	0x01, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x41, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x43, 0x4f, 0x59, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x52,
	0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x41, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x54,
	0x48, 0x45, 0x52, 0x10, 0x03, 0x42, 0x09, 0x5a, 0x07, 0x74, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_clientconf_proto_rawDescOnce sync.Once
	file_clientconf_proto_rawDescData = file_clientconf_proto_rawDesc
)

func file_clientconf_proto_rawDescGZIP() []byte {
	file_clientconf_proto_rawDescOnce.Do(func() {
		file_clientconf_proto_rawDescData = protoimpl.X.CompressGZIP(file_clientconf_proto_rawDescData)
	})
	return file_clientconf_proto_rawDescData
}

var file_clientconf_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_clientconf_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_clientconf_proto_goTypes = []interface{}{
	(RegistrarType)(0),         // 0: tapdance.RegistrarType
	(*DeploymentConfig)(nil),   // 1: tapdance.DeploymentConfig
	(*ClientConfig2)(nil),      // 2: tapdance.ClientConfig2
	(*Registrar)(nil),          // 3: tapdance.Registrar
	(*DecoyList)(nil),          // 4: tapdance.DecoyList
	(*PubKey)(nil),             // 5: tapdance.PubKey
	(*PhantomSubnetsList)(nil), // 6: tapdance.PhantomSubnetsList
	(TransportType)(0),         // 7: tapdance.TransportType
	(*anypb.Any)(nil),          // 8: google.protobuf.Any
}
var file_clientconf_proto_depIdxs = []int32{
	4, // 0: tapdance.DeploymentConfig.decoy_list:type_name -> tapdance.DecoyList
	5, // 1: tapdance.DeploymentConfig.default_pubkey:type_name -> tapdance.PubKey
	6, // 2: tapdance.DeploymentConfig.phantom_subnets_list:type_name -> tapdance.PhantomSubnetsList
	5, // 3: tapdance.DeploymentConfig.conjure_pubkey:type_name -> tapdance.PubKey
	7, // 4: tapdance.DeploymentConfig.suported_transports:type_name -> tapdance.TransportType
	3, // 5: tapdance.DeploymentConfig.registrars:type_name -> tapdance.Registrar
	1, // 6: tapdance.ClientConfig2.deployment_configs:type_name -> tapdance.DeploymentConfig
	0, // 7: tapdance.Registrar.registrar_type:type_name -> tapdance.RegistrarType
	8, // 8: tapdance.Registrar.registrar:type_name -> google.protobuf.Any
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_clientconf_proto_init() }
func file_clientconf_proto_init() {
	if File_clientconf_proto != nil {
		return
	}
	file_signalling_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_clientconf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_clientconf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConfig2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_clientconf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registrar); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_clientconf_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_clientconf_proto_goTypes,
		DependencyIndexes: file_clientconf_proto_depIdxs,
		EnumInfos:         file_clientconf_proto_enumTypes,
		MessageInfos:      file_clientconf_proto_msgTypes,
	}.Build()
	File_clientconf_proto = out.File
	file_clientconf_proto_rawDesc = nil
	file_clientconf_proto_goTypes = nil
	file_clientconf_proto_depIdxs = nil
}
