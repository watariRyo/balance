// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: user_tag.proto

package proto

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

type UserTagListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserTagListRequest) Reset() {
	*x = UserTagListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTagListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTagListRequest) ProtoMessage() {}

func (x *UserTagListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTagListRequest.ProtoReflect.Descriptor instead.
func (*UserTagListRequest) Descriptor() ([]byte, []int) {
	return file_user_tag_proto_rawDescGZIP(), []int{0}
}

func (x *UserTagListRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserTagListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTagList []*UserTagResponse `protobuf:"bytes,1,rep,name=user_tag_list,json=userTagList,proto3" json:"user_tag_list,omitempty"`
}

func (x *UserTagListResponse) Reset() {
	*x = UserTagListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTagListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTagListResponse) ProtoMessage() {}

func (x *UserTagListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTagListResponse.ProtoReflect.Descriptor instead.
func (*UserTagListResponse) Descriptor() ([]byte, []int) {
	return file_user_tag_proto_rawDescGZIP(), []int{1}
}

func (x *UserTagListResponse) GetUserTagList() []*UserTagResponse {
	if x != nil {
		return x.UserTagList
	}
	return nil
}

type UserTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TagName    string `protobuf:"bytes,3,opt,name=tag_name,json=tagName,proto3" json:"tag_name,omitempty"`
	GroupId    int64  `protobuf:"varint,4,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	GrantLimit string `protobuf:"bytes,5,opt,name=grant_limit,json=grantLimit,proto3" json:"grant_limit,omitempty"`
}

func (x *UserTagRequest) Reset() {
	*x = UserTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTagRequest) ProtoMessage() {}

func (x *UserTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTagRequest.ProtoReflect.Descriptor instead.
func (*UserTagRequest) Descriptor() ([]byte, []int) {
	return file_user_tag_proto_rawDescGZIP(), []int{2}
}

func (x *UserTagRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserTagRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserTagRequest) GetTagName() string {
	if x != nil {
		return x.TagName
	}
	return ""
}

func (x *UserTagRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *UserTagRequest) GetGrantLimit() string {
	if x != nil {
		return x.GrantLimit
	}
	return ""
}

type UserTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TagName    string `protobuf:"bytes,3,opt,name=tag_name,json=tagName,proto3" json:"tag_name,omitempty"`
	HasGroup   bool   `protobuf:"varint,4,opt,name=has_group,json=hasGroup,proto3" json:"has_group,omitempty"`
	GroupId    int64  `protobuf:"varint,5,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	GrantLimit string `protobuf:"bytes,6,opt,name=grant_limit,json=grantLimit,proto3" json:"grant_limit,omitempty"`
}

func (x *UserTagResponse) Reset() {
	*x = UserTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTagResponse) ProtoMessage() {}

func (x *UserTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTagResponse.ProtoReflect.Descriptor instead.
func (*UserTagResponse) Descriptor() ([]byte, []int) {
	return file_user_tag_proto_rawDescGZIP(), []int{3}
}

func (x *UserTagResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserTagResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserTagResponse) GetTagName() string {
	if x != nil {
		return x.TagName
	}
	return ""
}

func (x *UserTagResponse) GetHasGroup() bool {
	if x != nil {
		return x.HasGroup
	}
	return false
}

func (x *UserTagResponse) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *UserTagResponse) GetGrantLimit() string {
	if x != nil {
		return x.GrantLimit
	}
	return ""
}

type UserTagID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserTagID) Reset() {
	*x = UserTagID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_tag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTagID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTagID) ProtoMessage() {}

func (x *UserTagID) ProtoReflect() protoreflect.Message {
	mi := &file_user_tag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTagID.ProtoReflect.Descriptor instead.
func (*UserTagID) Descriptor() ([]byte, []int) {
	return file_user_tag_proto_rawDescGZIP(), []int{4}
}

func (x *UserTagID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserTagID) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_user_tag_proto protoreflect.FileDescriptor

var file_user_tag_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0b,
	0x75, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x90, 0x01, 0x0a, 0x0e,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xae,
	0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x61, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74,
	0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x73, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x34, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xcf, 0x02, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x12, 0x11,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x49,
	0x44, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0f, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x12, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x12,
	0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61,
	0x67, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x61, 0x67, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x61, 0x67, 0x49, 0x44, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x74, 0x61, 0x72, 0x69, 0x52, 0x79, 0x6f, 0x2f,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_tag_proto_rawDescOnce sync.Once
	file_user_tag_proto_rawDescData = file_user_tag_proto_rawDesc
)

func file_user_tag_proto_rawDescGZIP() []byte {
	file_user_tag_proto_rawDescOnce.Do(func() {
		file_user_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_tag_proto_rawDescData)
	})
	return file_user_tag_proto_rawDescData
}

var file_user_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_tag_proto_goTypes = []interface{}{
	(*UserTagListRequest)(nil),  // 0: server.UserTagListRequest
	(*UserTagListResponse)(nil), // 1: server.UserTagListResponse
	(*UserTagRequest)(nil),      // 2: server.UserTagRequest
	(*UserTagResponse)(nil),     // 3: server.UserTagResponse
	(*UserTagID)(nil),           // 4: server.UserTagID
}
var file_user_tag_proto_depIdxs = []int32{
	3, // 0: server.UserTagListResponse.user_tag_list:type_name -> server.UserTagResponse
	0, // 1: server.UserTagService.ListUserTag:input_type -> server.UserTagListRequest
	4, // 2: server.UserTagService.GetUserTag:input_type -> server.UserTagID
	2, // 3: server.UserTagService.RegisterUserTag:input_type -> server.UserTagRequest
	2, // 4: server.UserTagService.UpdateUserTag:input_type -> server.UserTagRequest
	4, // 5: server.UserTagService.DeleteUserTag:input_type -> server.UserTagID
	1, // 6: server.UserTagService.ListUserTag:output_type -> server.UserTagListResponse
	3, // 7: server.UserTagService.GetUserTag:output_type -> server.UserTagResponse
	3, // 8: server.UserTagService.RegisterUserTag:output_type -> server.UserTagResponse
	3, // 9: server.UserTagService.UpdateUserTag:output_type -> server.UserTagResponse
	4, // 10: server.UserTagService.DeleteUserTag:output_type -> server.UserTagID
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_tag_proto_init() }
func file_user_tag_proto_init() {
	if File_user_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTagListRequest); i {
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
		file_user_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTagListResponse); i {
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
		file_user_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTagRequest); i {
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
		file_user_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTagResponse); i {
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
		file_user_tag_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTagID); i {
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
			RawDescriptor: file_user_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_tag_proto_goTypes,
		DependencyIndexes: file_user_tag_proto_depIdxs,
		MessageInfos:      file_user_tag_proto_msgTypes,
	}.Build()
	File_user_tag_proto = out.File
	file_user_tag_proto_rawDesc = nil
	file_user_tag_proto_goTypes = nil
	file_user_tag_proto_depIdxs = nil
}
