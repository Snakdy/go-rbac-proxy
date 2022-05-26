// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: pkg/rbac/authority.proto

package rbac

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Verb int32

const (
	Verb_CREATE Verb = 0
	Verb_READ   Verb = 1
	Verb_UPDATE Verb = 2
	Verb_DELETE Verb = 3
	Verb_SUDO   Verb = 4
)

// Enum value maps for Verb.
var (
	Verb_name = map[int32]string{
		0: "CREATE",
		1: "READ",
		2: "UPDATE",
		3: "DELETE",
		4: "SUDO",
	}
	Verb_value = map[string]int32{
		"CREATE": 0,
		"READ":   1,
		"UPDATE": 2,
		"DELETE": 3,
		"SUDO":   4,
	}
)

func (x Verb) Enum() *Verb {
	p := new(Verb)
	*p = x
	return p
}

func (x Verb) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Verb) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_rbac_authority_proto_enumTypes[0].Descriptor()
}

func (Verb) Type() protoreflect.EnumType {
	return &file_pkg_rbac_authority_proto_enumTypes[0]
}

func (x Verb) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Verb.Descriptor instead.
func (Verb) EnumDescriptor() ([]byte, []int) {
	return file_pkg_rbac_authority_proto_rawDescGZIP(), []int{0}
}

type RoleBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject  string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject"`
	Resource string `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource"`
	Action   Verb   `protobuf:"varint,3,opt,name=action,proto3,enum=api.Verb" json:"action"`
}

func (x *RoleBinding) Reset() {
	*x = RoleBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rbac_authority_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleBinding) ProtoMessage() {}

func (x *RoleBinding) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rbac_authority_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleBinding.ProtoReflect.Descriptor instead.
func (*RoleBinding) Descriptor() ([]byte, []int) {
	return file_pkg_rbac_authority_proto_rawDescGZIP(), []int{0}
}

func (x *RoleBinding) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *RoleBinding) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *RoleBinding) GetAction() Verb {
	if x != nil {
		return x.Action
	}
	return Verb_CREATE
}

type AccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject  string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject"`
	Resource string `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource"`
	Action   Verb   `protobuf:"varint,3,opt,name=action,proto3,enum=api.Verb" json:"action"`
}

func (x *AccessRequest) Reset() {
	*x = AccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rbac_authority_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessRequest) ProtoMessage() {}

func (x *AccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rbac_authority_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessRequest.ProtoReflect.Descriptor instead.
func (*AccessRequest) Descriptor() ([]byte, []int) {
	return file_pkg_rbac_authority_proto_rawDescGZIP(), []int{1}
}

func (x *AccessRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *AccessRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *AccessRequest) GetAction() Verb {
	if x != nil {
		return x.Action
	}
	return Verb_CREATE
}

type GenericResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Ok      bool   `protobuf:"varint,2,opt,name=ok,proto3" json:"ok"`
}

func (x *GenericResponse) Reset() {
	*x = GenericResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rbac_authority_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericResponse) ProtoMessage() {}

func (x *GenericResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rbac_authority_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericResponse.ProtoReflect.Descriptor instead.
func (*GenericResponse) Descriptor() ([]byte, []int) {
	return file_pkg_rbac_authority_proto_rawDescGZIP(), []int{2}
}

func (x *GenericResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GenericResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type AddRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject  string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject"`
	Resource string `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource"`
	Action   Verb   `protobuf:"varint,3,opt,name=action,proto3,enum=api.Verb" json:"action"`
}

func (x *AddRoleRequest) Reset() {
	*x = AddRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rbac_authority_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoleRequest) ProtoMessage() {}

func (x *AddRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rbac_authority_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoleRequest.ProtoReflect.Descriptor instead.
func (*AddRoleRequest) Descriptor() ([]byte, []int) {
	return file_pkg_rbac_authority_proto_rawDescGZIP(), []int{3}
}

func (x *AddRoleRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *AddRoleRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *AddRoleRequest) GetAction() Verb {
	if x != nil {
		return x.Action
	}
	return Verb_CREATE
}

type AddGlobalRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject"`
	Role    string `protobuf:"bytes,2,opt,name=role,proto3" json:"role"`
}

func (x *AddGlobalRoleRequest) Reset() {
	*x = AddGlobalRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rbac_authority_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddGlobalRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGlobalRoleRequest) ProtoMessage() {}

func (x *AddGlobalRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rbac_authority_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGlobalRoleRequest.ProtoReflect.Descriptor instead.
func (*AddGlobalRoleRequest) Descriptor() ([]byte, []int) {
	return file_pkg_rbac_authority_proto_rawDescGZIP(), []int{4}
}

func (x *AddGlobalRoleRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *AddGlobalRoleRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type ListBySubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject"`
}

func (x *ListBySubRequest) Reset() {
	*x = ListBySubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rbac_authority_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBySubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBySubRequest) ProtoMessage() {}

func (x *ListBySubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rbac_authority_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBySubRequest.ProtoReflect.Descriptor instead.
func (*ListBySubRequest) Descriptor() ([]byte, []int) {
	return file_pkg_rbac_authority_proto_rawDescGZIP(), []int{5}
}

func (x *ListBySubRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

type ListByRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role"`
}

func (x *ListByRoleRequest) Reset() {
	*x = ListByRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rbac_authority_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListByRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListByRoleRequest) ProtoMessage() {}

func (x *ListByRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rbac_authority_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListByRoleRequest.ProtoReflect.Descriptor instead.
func (*ListByRoleRequest) Descriptor() ([]byte, []int) {
	return file_pkg_rbac_authority_proto_rawDescGZIP(), []int{6}
}

func (x *ListByRoleRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*RoleBinding `protobuf:"bytes,1,rep,name=results,proto3" json:"results"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rbac_authority_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rbac_authority_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_rbac_authority_proto_rawDescGZIP(), []int{7}
}

func (x *ListResponse) GetResults() []*RoleBinding {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_pkg_rbac_authority_proto protoreflect.FileDescriptor

var file_pkg_rbac_authority_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22,
	0x66, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x65, 0x72, 0x62, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x68, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x21,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x65, 0x72, 0x62, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x3b, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x69,
	0x0a, 0x0e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x65, 0x72,
	0x62, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x44, 0x0a, 0x14, 0x41, 0x64, 0x64,
	0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22,
	0x2c, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x53, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x27, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x3a, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x2a, 0x3e, 0x0a, 0x04, 0x56, 0x65, 0x72, 0x62, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x55, 0x44, 0x4f,
	0x10, 0x04, 0x32, 0xa4, 0x02, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x2f, 0x0a, 0x03, 0x43, 0x61, 0x6e, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x13, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x47, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x64, 0x64, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x79, 0x53, 0x75, 0x62, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x79, 0x53, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x37, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x16,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x69, 0x73, 0x6d,
	0x2f, 0x67, 0x6f, 0x2d, 0x72, 0x62, 0x61, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x72, 0x62, 0x61, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_rbac_authority_proto_rawDescOnce sync.Once
	file_pkg_rbac_authority_proto_rawDescData = file_pkg_rbac_authority_proto_rawDesc
)

func file_pkg_rbac_authority_proto_rawDescGZIP() []byte {
	file_pkg_rbac_authority_proto_rawDescOnce.Do(func() {
		file_pkg_rbac_authority_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_rbac_authority_proto_rawDescData)
	})
	return file_pkg_rbac_authority_proto_rawDescData
}

var file_pkg_rbac_authority_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_rbac_authority_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_rbac_authority_proto_goTypes = []interface{}{
	(Verb)(0),                    // 0: api.Verb
	(*RoleBinding)(nil),          // 1: api.RoleBinding
	(*AccessRequest)(nil),        // 2: api.AccessRequest
	(*GenericResponse)(nil),      // 3: api.GenericResponse
	(*AddRoleRequest)(nil),       // 4: api.AddRoleRequest
	(*AddGlobalRoleRequest)(nil), // 5: api.AddGlobalRoleRequest
	(*ListBySubRequest)(nil),     // 6: api.ListBySubRequest
	(*ListByRoleRequest)(nil),    // 7: api.ListByRoleRequest
	(*ListResponse)(nil),         // 8: api.ListResponse
}
var file_pkg_rbac_authority_proto_depIdxs = []int32{
	0, // 0: api.RoleBinding.action:type_name -> api.Verb
	0, // 1: api.AccessRequest.action:type_name -> api.Verb
	0, // 2: api.AddRoleRequest.action:type_name -> api.Verb
	1, // 3: api.ListResponse.results:type_name -> api.RoleBinding
	2, // 4: api.Authority.Can:input_type -> api.AccessRequest
	4, // 5: api.Authority.AddRole:input_type -> api.AddRoleRequest
	5, // 6: api.Authority.AddGlobalRole:input_type -> api.AddGlobalRoleRequest
	6, // 7: api.Authority.ListBySub:input_type -> api.ListBySubRequest
	7, // 8: api.Authority.ListByRole:input_type -> api.ListByRoleRequest
	3, // 9: api.Authority.Can:output_type -> api.GenericResponse
	3, // 10: api.Authority.AddRole:output_type -> api.GenericResponse
	3, // 11: api.Authority.AddGlobalRole:output_type -> api.GenericResponse
	8, // 12: api.Authority.ListBySub:output_type -> api.ListResponse
	8, // 13: api.Authority.ListByRole:output_type -> api.ListResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_rbac_authority_proto_init() }
func file_pkg_rbac_authority_proto_init() {
	if File_pkg_rbac_authority_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_rbac_authority_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleBinding); i {
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
		file_pkg_rbac_authority_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessRequest); i {
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
		file_pkg_rbac_authority_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericResponse); i {
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
		file_pkg_rbac_authority_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRoleRequest); i {
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
		file_pkg_rbac_authority_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddGlobalRoleRequest); i {
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
		file_pkg_rbac_authority_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBySubRequest); i {
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
		file_pkg_rbac_authority_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListByRoleRequest); i {
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
		file_pkg_rbac_authority_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
			RawDescriptor: file_pkg_rbac_authority_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_rbac_authority_proto_goTypes,
		DependencyIndexes: file_pkg_rbac_authority_proto_depIdxs,
		EnumInfos:         file_pkg_rbac_authority_proto_enumTypes,
		MessageInfos:      file_pkg_rbac_authority_proto_msgTypes,
	}.Build()
	File_pkg_rbac_authority_proto = out.File
	file_pkg_rbac_authority_proto_rawDesc = nil
	file_pkg_rbac_authority_proto_goTypes = nil
	file_pkg_rbac_authority_proto_depIdxs = nil
}
