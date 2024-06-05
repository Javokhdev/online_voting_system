// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: public.proto

package public

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

type CreatePublicReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Birthday  string `protobuf:"bytes,3,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender    string `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Nation    string `protobuf:"bytes,5,opt,name=nation,proto3" json:"nation,omitempty"`
	PartyId   string `protobuf:"bytes,6,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
}

func (x *CreatePublicReq) Reset() {
	*x = CreatePublicReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePublicReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePublicReq) ProtoMessage() {}

func (x *CreatePublicReq) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePublicReq.ProtoReflect.Descriptor instead.
func (*CreatePublicReq) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePublicReq) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreatePublicReq) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreatePublicReq) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *CreatePublicReq) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *CreatePublicReq) GetNation() string {
	if x != nil {
		return x.Nation
	}
	return ""
}

func (x *CreatePublicReq) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

type GetPublicRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Birthday  string `protobuf:"bytes,3,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender    string `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Nation    string `protobuf:"bytes,5,opt,name=nation,proto3" json:"nation,omitempty"`
	PartyId   string `protobuf:"bytes,6,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
}

func (x *GetPublicRes) Reset() {
	*x = GetPublicRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPublicRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPublicRes) ProtoMessage() {}

func (x *GetPublicRes) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPublicRes.ProtoReflect.Descriptor instead.
func (*GetPublicRes) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{1}
}

func (x *GetPublicRes) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetPublicRes) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GetPublicRes) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *GetPublicRes) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *GetPublicRes) GetNation() string {
	if x != nil {
		return x.Nation
	}
	return ""
}

func (x *GetPublicRes) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

type GetAllPublicRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parties []*GetPublicRes `protobuf:"bytes,1,rep,name=Parties,proto3" json:"Parties,omitempty"`
}

func (x *GetAllPublicRes) Reset() {
	*x = GetAllPublicRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPublicRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPublicRes) ProtoMessage() {}

func (x *GetAllPublicRes) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPublicRes.ProtoReflect.Descriptor instead.
func (*GetAllPublicRes) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllPublicRes) GetParties() []*GetPublicRes {
	if x != nil {
		return x.Parties
	}
	return nil
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	Gender    string `protobuf:"bytes,2,opt,name=gender,proto3" json:"gender,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{3}
}

func (x *Filter) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Filter) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

type ById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ById) Reset() {
	*x = ById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ById) ProtoMessage() {}

func (x *ById) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ById.ProtoReflect.Descriptor instead.
func (*ById) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{4}
}

func (x *ById) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_public_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_public_proto_rawDescGZIP(), []int{5}
}

var File_public_proto protoreflect.FileDescriptor

var file_public_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0xb4, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64,
	0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64,
	0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x22, 0xb1, 0x01,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49,
	0x64, 0x22, 0x41, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x52, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x50, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x73, 0x52, 0x07, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x22, 0x3f, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x16, 0x0a, 0x04, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x06, 0x0a,
	0x04, 0x56, 0x6f, 0x69, 0x64, 0x32, 0x80, 0x02, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x17, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x42,
	0x79, 0x49, 0x64, 0x1a, 0x14, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x0e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x17, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x2e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x73,
	0x1a, 0x0c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00,
	0x12, 0x26, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_public_proto_rawDescOnce sync.Once
	file_public_proto_rawDescData = file_public_proto_rawDesc
)

func file_public_proto_rawDescGZIP() []byte {
	file_public_proto_rawDescOnce.Do(func() {
		file_public_proto_rawDescData = protoimpl.X.CompressGZIP(file_public_proto_rawDescData)
	})
	return file_public_proto_rawDescData
}

var file_public_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_public_proto_goTypes = []interface{}{
	(*CreatePublicReq)(nil), // 0: public.CreatePublicReq
	(*GetPublicRes)(nil),    // 1: public.GetPublicRes
	(*GetAllPublicRes)(nil), // 2: public.GetAllPublicRes
	(*Filter)(nil),          // 3: public.Filter
	(*ById)(nil),            // 4: public.ById
	(*Void)(nil),            // 5: public.Void
}
var file_public_proto_depIdxs = []int32{
	1, // 0: public.GetAllPublicRes.Parties:type_name -> public.GetPublicRes
	0, // 1: public.PublicService.Create:input_type -> public.CreatePublicReq
	4, // 2: public.PublicService.GetById:input_type -> public.ById
	3, // 3: public.PublicService.GetAll:input_type -> public.Filter
	1, // 4: public.PublicService.Update:input_type -> public.GetPublicRes
	4, // 5: public.PublicService.Delete:input_type -> public.ById
	5, // 6: public.PublicService.Create:output_type -> public.Void
	1, // 7: public.PublicService.GetById:output_type -> public.GetPublicRes
	2, // 8: public.PublicService.GetAll:output_type -> public.GetAllPublicRes
	5, // 9: public.PublicService.Update:output_type -> public.Void
	5, // 10: public.PublicService.Delete:output_type -> public.Void
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_public_proto_init() }
func file_public_proto_init() {
	if File_public_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_public_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePublicReq); i {
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
		file_public_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPublicRes); i {
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
		file_public_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPublicRes); i {
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
		file_public_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_public_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ById); i {
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
		file_public_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
			RawDescriptor: file_public_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_public_proto_goTypes,
		DependencyIndexes: file_public_proto_depIdxs,
		MessageInfos:      file_public_proto_msgTypes,
	}.Build()
	File_public_proto = out.File
	file_public_proto_rawDesc = nil
	file_public_proto_goTypes = nil
	file_public_proto_depIdxs = nil
}