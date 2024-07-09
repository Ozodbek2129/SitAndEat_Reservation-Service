// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: menu.proto

package menu

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

type CreateF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestuarantId string  `protobuf:"bytes,1,opt,name=restuarantId,proto3" json:"restuarantId,omitempty"`
	Name         string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description  string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price        float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Image        []byte  `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *CreateF) Reset() {
	*x = CreateF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateF) ProtoMessage() {}

func (x *CreateF) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateF.ProtoReflect.Descriptor instead.
func (*CreateF) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{0}
}

func (x *CreateF) GetRestuarantId() string {
	if x != nil {
		return x.RestuarantId
	}
	return ""
}

func (x *CreateF) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateF) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateF) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateF) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[1]
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
	return file_menu_proto_rawDescGZIP(), []int{1}
}

type Food struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RestuarantId string  `protobuf:"bytes,2,opt,name=restuarantId,proto3" json:"restuarantId,omitempty"`
	Name         string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description  string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Price        float32 `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	Image        []byte  `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	CreatedAt    string  `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdateAt     string  `protobuf:"bytes,8,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
}

func (x *Food) Reset() {
	*x = Food{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Food) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Food) ProtoMessage() {}

func (x *Food) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Food.ProtoReflect.Descriptor instead.
func (*Food) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{2}
}

func (x *Food) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Food) GetRestuarantId() string {
	if x != nil {
		return x.RestuarantId
	}
	return ""
}

func (x *Food) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Food) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Food) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Food) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *Food) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Food) GetUpdateAt() string {
	if x != nil {
		return x.UpdateAt
	}
	return ""
}

type Foods struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foods []*Food `protobuf:"bytes,1,rep,name=foods,proto3" json:"foods,omitempty"`
}

func (x *Foods) Reset() {
	*x = Foods{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foods) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foods) ProtoMessage() {}

func (x *Foods) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foods.ProtoReflect.Descriptor instead.
func (*Foods) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{3}
}

func (x *Foods) GetFoods() []*Food {
	if x != nil {
		return x.Foods
	}
	return nil
}

type FoodId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FoodId) Reset() {
	*x = FoodId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodId) ProtoMessage() {}

func (x *FoodId) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodId.ProtoReflect.Descriptor instead.
func (*FoodId) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{4}
}

func (x *FoodId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[5]
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
	return file_menu_proto_rawDescGZIP(), []int{5}
}

func (x *Status) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type UpdateF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RestuarantId string  `protobuf:"bytes,2,opt,name=restuarantId,proto3" json:"restuarantId,omitempty"`
	Name         string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description  string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Price        float32 `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	Image        []byte  `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *UpdateF) Reset() {
	*x = UpdateF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateF) ProtoMessage() {}

func (x *UpdateF) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateF.ProtoReflect.Descriptor instead.
func (*UpdateF) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateF) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateF) GetRestuarantId() string {
	if x != nil {
		return x.RestuarantId
	}
	return ""
}

func (x *UpdateF) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateF) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateF) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateF) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_menu_proto protoreflect.FileDescriptor

var file_menu_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x65,
	0x6e, 0x75, 0x22, 0x8f, 0x01, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x12, 0x22,
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0xd8, 0x01, 0x0a,
	0x04, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x75, 0x61, 0x72,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x74, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22, 0x29, 0x0a, 0x05, 0x46, 0x6f, 0x6f, 0x64, 0x73,
	0x12, 0x20, 0x0a, 0x05, 0x66, 0x6f, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x05, 0x66, 0x6f, 0x6f,
	0x64, 0x73, 0x22, 0x18, 0x0a, 0x06, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x9f,
	0x01, 0x0a, 0x07, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65,
	0x73, 0x74, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x32, 0xdd, 0x01, 0x0a, 0x04, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x2b, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x0d, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x1a, 0x0c, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x46, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x0a, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x1a, 0x0b, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x73, 0x22, 0x00,
	0x12, 0x25, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x0c, 0x2e, 0x6d, 0x65,
	0x6e, 0x75, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x6d, 0x65, 0x6e, 0x75,
	0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x0d, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x46, 0x1a, 0x0c, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f,
	0x6f, 0x64, 0x12, 0x0c, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x64,
	0x1a, 0x0c, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00,
	0x42, 0x0f, 0x5a, 0x0d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x6e,
	0x75, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_menu_proto_rawDescOnce sync.Once
	file_menu_proto_rawDescData = file_menu_proto_rawDesc
)

func file_menu_proto_rawDescGZIP() []byte {
	file_menu_proto_rawDescOnce.Do(func() {
		file_menu_proto_rawDescData = protoimpl.X.CompressGZIP(file_menu_proto_rawDescData)
	})
	return file_menu_proto_rawDescData
}

var file_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_menu_proto_goTypes = []interface{}{
	(*CreateF)(nil), // 0: menu.CreateF
	(*Void)(nil),    // 1: menu.Void
	(*Food)(nil),    // 2: menu.Food
	(*Foods)(nil),   // 3: menu.Foods
	(*FoodId)(nil),  // 4: menu.FoodId
	(*Status)(nil),  // 5: menu.Status
	(*UpdateF)(nil), // 6: menu.UpdateF
}
var file_menu_proto_depIdxs = []int32{
	2, // 0: menu.Foods.foods:type_name -> menu.Food
	0, // 1: menu.Menu.CreateFood:input_type -> menu.CreateF
	1, // 2: menu.Menu.GetAllFoods:input_type -> menu.Void
	4, // 3: menu.Menu.GetFood:input_type -> menu.FoodId
	6, // 4: menu.Menu.UpdateFood:input_type -> menu.UpdateF
	4, // 5: menu.Menu.DeleteFood:input_type -> menu.FoodId
	5, // 6: menu.Menu.CreateFood:output_type -> menu.Status
	3, // 7: menu.Menu.GetAllFoods:output_type -> menu.Foods
	2, // 8: menu.Menu.GetFood:output_type -> menu.Food
	5, // 9: menu.Menu.UpdateFood:output_type -> menu.Status
	5, // 10: menu.Menu.DeleteFood:output_type -> menu.Status
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_menu_proto_init() }
func file_menu_proto_init() {
	if File_menu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_menu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateF); i {
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
		file_menu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_menu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Food); i {
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
		file_menu_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Foods); i {
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
		file_menu_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodId); i {
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
		file_menu_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_menu_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateF); i {
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
			RawDescriptor: file_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_menu_proto_goTypes,
		DependencyIndexes: file_menu_proto_depIdxs,
		MessageInfos:      file_menu_proto_msgTypes,
	}.Build()
	File_menu_proto = out.File
	file_menu_proto_rawDesc = nil
	file_menu_proto_goTypes = nil
	file_menu_proto_depIdxs = nil
}
