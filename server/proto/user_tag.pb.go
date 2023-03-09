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

type ListUserTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ListUserTagRequest) Reset() {
	*x = ListUserTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserTagRequest) ProtoMessage() {}

func (x *ListUserTagRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListUserTagRequest.ProtoReflect.Descriptor instead.
func (*ListUserTagRequest) Descriptor() ([]byte, []int) {
	return file_user_tag_proto_rawDescGZIP(), []int{0}
}

func (x *ListUserTagRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ListUserTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTagList []*UserTagResponse `protobuf:"bytes,1,rep,name=user_tag_list,json=userTagList,proto3" json:"user_tag_list,omitempty"`
}

func (x *ListUserTagResponse) Reset() {
	*x = ListUserTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserTagResponse) ProtoMessage() {}

func (x *ListUserTagResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListUserTagResponse.ProtoReflect.Descriptor instead.
func (*ListUserTagResponse) Descriptor() ([]byte, []int) {
	return file_user_tag_proto_rawDescGZIP(), []int{1}
}

func (x *ListUserTagResponse) GetUserTagList() []*UserTagResponse {
	if x != nil {
		return x.UserTagList
	}
	return nil
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
		mi := &file_user_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTagResponse) ProtoMessage() {}

func (x *UserTagResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserTagResponse.ProtoReflect.Descriptor instead.
func (*UserTagResponse) Descriptor() ([]byte, []int) {
	return file_user_tag_proto_rawDescGZIP(), []int{2}
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

type GetUserTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserTagRequest) Reset() {
	*x = GetUserTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTagRequest) ProtoMessage() {}

func (x *GetUserTagRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetUserTagRequest.ProtoReflect.Descriptor instead.
func (*GetUserTagRequest) Descriptor() ([]byte, []int) {
	return file_user_tag_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserTagRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserTagResponse struct {
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

func (x *GetUserTagResponse) Reset() {
	*x = GetUserTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_tag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTagResponse) ProtoMessage() {}

func (x *GetUserTagResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetUserTagResponse.ProtoReflect.Descriptor instead.
func (*GetUserTagResponse) Descriptor() ([]byte, []int) {
	return file_user_tag_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserTagResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetUserTagResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetUserTagResponse) GetTagName() string {
	if x != nil {
		return x.TagName
	}
	return ""
}

func (x *GetUserTagResponse) GetHasGroup() bool {
	if x != nil {
		return x.HasGroup
	}
	return false
}

func (x *GetUserTagResponse) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GetUserTagResponse) GetGrantLimit() string {
	if x != nil {
		return x.GrantLimit
	}
	return ""
}

type RegisterUserTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TagName    string `protobuf:"bytes,2,opt,name=tag_name,json=tagName,proto3" json:"tag_name,omitempty"`
	HasGroup   bool   `protobuf:"varint,3,opt,name=has_group,json=hasGroup,proto3" json:"has_group,omitempty"`
	GroupId    int64  `protobuf:"varint,4,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	GrantLimit string `protobuf:"bytes,5,opt,name=grant_limit,json=grantLimit,proto3" json:"grant_limit,omitempty"`
}

func (x *RegisterUserTagRequest) Reset() {
	*x = RegisterUserTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_tag_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserTagRequest) ProtoMessage() {}

func (x *RegisterUserTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_tag_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserTagRequest.ProtoReflect.Descriptor instead.
func (*RegisterUserTagRequest) Descriptor() ([]byte, []int) {
	return file_user_tag_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterUserTagRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RegisterUserTagRequest) GetTagName() string {
	if x != nil {
		return x.TagName
	}
	return ""
}

func (x *RegisterUserTagRequest) GetHasGroup() bool {
	if x != nil {
		return x.HasGroup
	}
	return false
}

func (x *RegisterUserTagRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *RegisterUserTagRequest) GetGrantLimit() string {
	if x != nil {
		return x.GrantLimit
	}
	return ""
}

type RegisterUserTagResponse struct {
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

func (x *RegisterUserTagResponse) Reset() {
	*x = RegisterUserTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_tag_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserTagResponse) ProtoMessage() {}

func (x *RegisterUserTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_tag_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserTagResponse.ProtoReflect.Descriptor instead.
func (*RegisterUserTagResponse) Descriptor() ([]byte, []int) {
	return file_user_tag_proto_rawDescGZIP(), []int{6}
}

func (x *RegisterUserTagResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RegisterUserTagResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RegisterUserTagResponse) GetTagName() string {
	if x != nil {
		return x.TagName
	}
	return ""
}

func (x *RegisterUserTagResponse) GetHasGroup() bool {
	if x != nil {
		return x.HasGroup
	}
	return false
}

func (x *RegisterUserTagResponse) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *RegisterUserTagResponse) GetGrantLimit() string {
	if x != nil {
		return x.GrantLimit
	}
	return ""
}

type UpdateUserTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UpdateUserTagRequest) Reset() {
	*x = UpdateUserTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_tag_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserTagRequest) ProtoMessage() {}

func (x *UpdateUserTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_tag_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserTagRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserTagRequest) Descriptor() ([]byte, []int) {
	return file_user_tag_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateUserTagRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UpdateUserTagResponse struct {
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

func (x *UpdateUserTagResponse) Reset() {
	*x = UpdateUserTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_tag_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserTagResponse) ProtoMessage() {}

func (x *UpdateUserTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_tag_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserTagResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserTagResponse) Descriptor() ([]byte, []int) {
	return file_user_tag_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateUserTagResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateUserTagResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateUserTagResponse) GetTagName() string {
	if x != nil {
		return x.TagName
	}
	return ""
}

func (x *UpdateUserTagResponse) GetHasGroup() bool {
	if x != nil {
		return x.HasGroup
	}
	return false
}

func (x *UpdateUserTagResponse) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *UpdateUserTagResponse) GetGrantLimit() string {
	if x != nil {
		return x.GrantLimit
	}
	return ""
}

type DeleteUserTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteUserTagRequest) Reset() {
	*x = DeleteUserTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_tag_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserTagRequest) ProtoMessage() {}

func (x *DeleteUserTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_tag_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserTagRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserTagRequest) Descriptor() ([]byte, []int) {
	return file_user_tag_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteUserTagRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DeleteUserTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteUserTagResponse) Reset() {
	*x = DeleteUserTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_tag_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserTagResponse) ProtoMessage() {}

func (x *DeleteUserTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_tag_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserTagResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserTagResponse) Descriptor() ([]byte, []int) {
	return file_user_tag_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteUserTagResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteUserTagResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_user_tag_proto protoreflect.FileDescriptor

var file_user_tag_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a,
	0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0b, 0x75, 0x73,
	0x65, 0x72, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xae, 0x01, 0x0a, 0x0f, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x72, 0x61,
	0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x67, 0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x2c, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xb1, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x67,
	0x72, 0x61, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xa5, 0x01, 0x0a,
	0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68,
	0x61, 0x73, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x68, 0x61, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0xb6, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x67,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x73, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x2f, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xb4,
	0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x68, 0x61, 0x73, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x68, 0x61, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x2f, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0x83, 0x03, 0x0a, 0x0e, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x61, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x12,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x61, 0x67, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x74,
	0x61, 0x72, 0x69, 0x52, 0x79, 0x6f, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_user_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_user_tag_proto_goTypes = []interface{}{
	(*ListUserTagRequest)(nil),      // 0: proto.ListUserTagRequest
	(*ListUserTagResponse)(nil),     // 1: proto.ListUserTagResponse
	(*UserTagResponse)(nil),         // 2: proto.UserTagResponse
	(*GetUserTagRequest)(nil),       // 3: proto.GetUserTagRequest
	(*GetUserTagResponse)(nil),      // 4: proto.GetUserTagResponse
	(*RegisterUserTagRequest)(nil),  // 5: proto.RegisterUserTagRequest
	(*RegisterUserTagResponse)(nil), // 6: proto.RegisterUserTagResponse
	(*UpdateUserTagRequest)(nil),    // 7: proto.UpdateUserTagRequest
	(*UpdateUserTagResponse)(nil),   // 8: proto.UpdateUserTagResponse
	(*DeleteUserTagRequest)(nil),    // 9: proto.DeleteUserTagRequest
	(*DeleteUserTagResponse)(nil),   // 10: proto.DeleteUserTagResponse
}
var file_user_tag_proto_depIdxs = []int32{
	2,  // 0: proto.ListUserTagResponse.user_tag_list:type_name -> proto.UserTagResponse
	0,  // 1: proto.UserTagService.ListUserTag:input_type -> proto.ListUserTagRequest
	3,  // 2: proto.UserTagService.GetUserTag:input_type -> proto.GetUserTagRequest
	5,  // 3: proto.UserTagService.RegisterUserTag:input_type -> proto.RegisterUserTagRequest
	7,  // 4: proto.UserTagService.UpdateUserTag:input_type -> proto.UpdateUserTagRequest
	9,  // 5: proto.UserTagService.DeleteUserTag:input_type -> proto.DeleteUserTagRequest
	1,  // 6: proto.UserTagService.ListUserTag:output_type -> proto.ListUserTagResponse
	4,  // 7: proto.UserTagService.GetUserTag:output_type -> proto.GetUserTagResponse
	6,  // 8: proto.UserTagService.RegisterUserTag:output_type -> proto.RegisterUserTagResponse
	8,  // 9: proto.UserTagService.UpdateUserTag:output_type -> proto.UpdateUserTagResponse
	10, // 10: proto.UserTagService.DeleteUserTag:output_type -> proto.DeleteUserTagResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_user_tag_proto_init() }
func file_user_tag_proto_init() {
	if File_user_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserTagRequest); i {
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
			switch v := v.(*ListUserTagResponse); i {
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
		file_user_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTagRequest); i {
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
			switch v := v.(*GetUserTagResponse); i {
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
		file_user_tag_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterUserTagRequest); i {
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
		file_user_tag_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterUserTagResponse); i {
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
		file_user_tag_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserTagRequest); i {
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
		file_user_tag_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserTagResponse); i {
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
		file_user_tag_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserTagRequest); i {
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
		file_user_tag_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserTagResponse); i {
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
			NumMessages:   11,
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
