// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: resirvation.proto

package resirvation

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

type Restuarant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address     string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Phone       string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Restuarant) Reset() {
	*x = Restuarant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resirvation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restuarant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restuarant) ProtoMessage() {}

func (x *Restuarant) ProtoReflect() protoreflect.Message {
	mi := &file_resirvation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restuarant.ProtoReflect.Descriptor instead.
func (*Restuarant) Descriptor() ([]byte, []int) {
	return file_resirvation_proto_rawDescGZIP(), []int{0}
}

func (x *Restuarant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Restuarant) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Restuarant) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Restuarant) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resirvation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_resirvation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_resirvation_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type AllRestuarant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AllRestuarant) Reset() {
	*x = AllRestuarant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resirvation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllRestuarant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllRestuarant) ProtoMessage() {}

func (x *AllRestuarant) ProtoReflect() protoreflect.Message {
	mi := &file_resirvation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllRestuarant.ProtoReflect.Descriptor instead.
func (*AllRestuarant) Descriptor() ([]byte, []int) {
	return file_resirvation_proto_rawDescGZIP(), []int{2}
}

type Restuanants struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Restuanants []*GetRes `protobuf:"bytes,1,rep,name=restuanants,proto3" json:"restuanants,omitempty"`
}

func (x *Restuanants) Reset() {
	*x = Restuanants{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resirvation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restuanants) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restuanants) ProtoMessage() {}

func (x *Restuanants) ProtoReflect() protoreflect.Message {
	mi := &file_resirvation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restuanants.ProtoReflect.Descriptor instead.
func (*Restuanants) Descriptor() ([]byte, []int) {
	return file_resirvation_proto_rawDescGZIP(), []int{3}
}

func (x *Restuanants) GetRestuanants() []*GetRes {
	if x != nil {
		return x.Restuanants
	}
	return nil
}

type RestuanantId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestuarantId string `protobuf:"bytes,1,opt,name=restuarantId,proto3" json:"restuarantId,omitempty"`
}

func (x *RestuanantId) Reset() {
	*x = RestuanantId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resirvation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestuanantId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestuanantId) ProtoMessage() {}

func (x *RestuanantId) ProtoReflect() protoreflect.Message {
	mi := &file_resirvation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestuanantId.ProtoReflect.Descriptor instead.
func (*RestuanantId) Descriptor() ([]byte, []int) {
	return file_resirvation_proto_rawDescGZIP(), []int{4}
}

func (x *RestuanantId) GetRestuarantId() string {
	if x != nil {
		return x.RestuarantId
	}
	return ""
}

type GetRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address     string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Phone       string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   string `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   string `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *GetRes) Reset() {
	*x = GetRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resirvation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRes) ProtoMessage() {}

func (x *GetRes) ProtoReflect() protoreflect.Message {
	mi := &file_resirvation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRes.ProtoReflect.Descriptor instead.
func (*GetRes) Descriptor() ([]byte, []int) {
	return file_resirvation_proto_rawDescGZIP(), []int{5}
}

func (x *GetRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetRes) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetRes) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GetRes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetRes) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GetRes) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_resirvation_proto protoreflect.FileDescriptor

var file_resirvation_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x73, 0x69, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x65, 0x73, 0x69, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x72, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x74, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x74, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x22, 0x44, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x74, 0x75,
	0x61, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x35, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x74, 0x75, 0x61,
	0x6e, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65,
	0x73, 0x69, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x52, 0x0b, 0x72, 0x65, 0x73, 0x74, 0x75, 0x61, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x32, 0x0a,
	0x0c, 0x52, 0x65, 0x73, 0x74, 0x75, 0x61, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x72, 0x65, 0x73, 0x74, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0xba, 0x01, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xe7,
	0x02, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x69, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x74, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x1a, 0x13, 0x2e, 0x72, 0x65,
	0x73, 0x69, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x00, 0x12, 0x4b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x74, 0x75, 0x61, 0x72,
	0x61, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x75, 0x61, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x00, 0x12,
	0x41, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74,
	0x12, 0x19, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x74, 0x75, 0x61, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x72, 0x65,
	0x73, 0x69, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74,
	0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x1a, 0x13, 0x2e, 0x72, 0x65,
	0x73, 0x69, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x00, 0x12, 0x44, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74,
	0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x75, 0x61, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x1a, 0x13, 0x2e, 0x72, 0x65, 0x73, 0x69, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resirvation_proto_rawDescOnce sync.Once
	file_resirvation_proto_rawDescData = file_resirvation_proto_rawDesc
)

func file_resirvation_proto_rawDescGZIP() []byte {
	file_resirvation_proto_rawDescOnce.Do(func() {
		file_resirvation_proto_rawDescData = protoimpl.X.CompressGZIP(file_resirvation_proto_rawDescData)
	})
	return file_resirvation_proto_rawDescData
}

var file_resirvation_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_resirvation_proto_goTypes = []any{
	(*Restuarant)(nil),    // 0: resirvation.Restuarant
	(*Status)(nil),        // 1: resirvation.Status
	(*AllRestuarant)(nil), // 2: resirvation.AllRestuarant
	(*Restuanants)(nil),   // 3: resirvation.Restuanants
	(*RestuanantId)(nil),  // 4: resirvation.RestuanantId
	(*GetRes)(nil),        // 5: resirvation.GetRes
}
var file_resirvation_proto_depIdxs = []int32{
	5, // 0: resirvation.Restuanants.restuanants:type_name -> resirvation.GetRes
	0, // 1: resirvation.Resirvation.CreateRestaurant:input_type -> resirvation.Restuarant
	2, // 2: resirvation.Resirvation.GetAllRestaurants:input_type -> resirvation.AllRestuarant
	4, // 3: resirvation.Resirvation.GetRestuarant:input_type -> resirvation.RestuanantId
	5, // 4: resirvation.Resirvation.UpdateRestuarant:input_type -> resirvation.GetRes
	4, // 5: resirvation.Resirvation.DeleteRestuarant:input_type -> resirvation.RestuanantId
	1, // 6: resirvation.Resirvation.CreateRestaurant:output_type -> resirvation.Status
	3, // 7: resirvation.Resirvation.GetAllRestaurants:output_type -> resirvation.Restuanants
	5, // 8: resirvation.Resirvation.GetRestuarant:output_type -> resirvation.GetRes
	1, // 9: resirvation.Resirvation.UpdateRestuarant:output_type -> resirvation.Status
	1, // 10: resirvation.Resirvation.DeleteRestuarant:output_type -> resirvation.Status
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_resirvation_proto_init() }
func file_resirvation_proto_init() {
	if File_resirvation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resirvation_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Restuarant); i {
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
		file_resirvation_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Status); i {
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
		file_resirvation_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AllRestuarant); i {
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
		file_resirvation_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Restuanants); i {
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
		file_resirvation_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RestuanantId); i {
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
		file_resirvation_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetRes); i {
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
			RawDescriptor: file_resirvation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_resirvation_proto_goTypes,
		DependencyIndexes: file_resirvation_proto_depIdxs,
		MessageInfos:      file_resirvation_proto_msgTypes,
	}.Build()
	File_resirvation_proto = out.File
	file_resirvation_proto_rawDesc = nil
	file_resirvation_proto_goTypes = nil
	file_resirvation_proto_depIdxs = nil
}