// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: comments.proto

package comments

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    int64  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	ProductID int64  `protobuf:"varint,2,opt,name=productID,proto3" json:"productID,omitempty"`
	Text      string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CreateCommentRequest) Reset() {
	*x = CreateCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentRequest) ProtoMessage() {}

func (x *CreateCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentRequest.ProtoReflect.Descriptor instead.
func (*CreateCommentRequest) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCommentRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CreateCommentRequest) GetProductID() int64 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *CreateCommentRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type CreateCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentID int64 `protobuf:"varint,1,opt,name=commentID,proto3" json:"commentID,omitempty"`
}

func (x *CreateCommentResponse) Reset() {
	*x = CreateCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentResponse) ProtoMessage() {}

func (x *CreateCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentResponse.ProtoReflect.Descriptor instead.
func (*CreateCommentResponse) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCommentResponse) GetCommentID() int64 {
	if x != nil {
		return x.CommentID
	}
	return 0
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     int64                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID int64                  `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Text   string                 `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Ts     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{2}
}

func (x *Comment) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Comment) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Comment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Comment) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

type GetCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID int64 `protobuf:"varint,1,opt,name=productID,proto3" json:"productID,omitempty"`
}

func (x *GetCommentsRequest) Reset() {
	*x = GetCommentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsRequest) ProtoMessage() {}

func (x *GetCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsRequest.ProtoReflect.Descriptor instead.
func (*GetCommentsRequest) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{3}
}

func (x *GetCommentsRequest) GetProductID() int64 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

type GetCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID int64      `protobuf:"varint,1,opt,name=productID,proto3" json:"productID,omitempty"`
	Comments  []*Comment `protobuf:"bytes,2,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *GetCommentsResponse) Reset() {
	*x = GetCommentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsResponse) ProtoMessage() {}

func (x *GetCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsResponse.ProtoReflect.Descriptor instead.
func (*GetCommentsResponse) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{4}
}

func (x *GetCommentsResponse) GetProductID() int64 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *GetCommentsResponse) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

var File_comments_proto protoreflect.FileDescriptor

var file_comments_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x24, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x25, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22,
	0x02, 0x20, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x1e,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42,
	0x07, 0x72, 0x05, 0x10, 0x05, 0x18, 0xff, 0x01, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x35,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x22, 0xce, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x17, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1f, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22,
	0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05,
	0x10, 0x05, 0x18, 0x80, 0x02, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2a, 0x0a, 0x02, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x3a, 0x3d, 0x92, 0x41, 0x3a, 0x0a, 0x38, 0x2a, 0x07,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x20, 0x69, 0x74, 0x65, 0x6d, 0xd2, 0x01, 0x02, 0x49, 0x44, 0xd2, 0x01, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0xd2, 0x01, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0xd2,
	0x01, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x44, 0x22, 0x7e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x49, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x32, 0xea, 0x02, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0xa7, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x3a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x00, 0x12, 0x9c, 0x01, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x38, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x92, 0x41, 0x00, 0x1a, 0x15, 0x92, 0x41, 0x12, 0x12, 0x10,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x42, 0x79, 0x5a, 0x24, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x92, 0x41, 0x50, 0x12, 0x26, 0x0a, 0x1d, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x20,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x20, 0x41, 0x50, 0x49, 0x32, 0x05, 0x31, 0x2e,
	0x30, 0x2e, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_comments_proto_rawDescOnce sync.Once
	file_comments_proto_rawDescData = file_comments_proto_rawDesc
)

func file_comments_proto_rawDescGZIP() []byte {
	file_comments_proto_rawDescOnce.Do(func() {
		file_comments_proto_rawDescData = protoimpl.X.CompressGZIP(file_comments_proto_rawDescData)
	})
	return file_comments_proto_rawDescData
}

var file_comments_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_comments_proto_goTypes = []interface{}{
	(*CreateCommentRequest)(nil),  // 0: example.comments.pkg.api.comments.v1.CreateCommentRequest
	(*CreateCommentResponse)(nil), // 1: example.comments.pkg.api.comments.v1.CreateCommentResponse
	(*Comment)(nil),               // 2: example.comments.pkg.api.comments.v1.Comment
	(*GetCommentsRequest)(nil),    // 3: example.comments.pkg.api.comments.v1.GetCommentsRequest
	(*GetCommentsResponse)(nil),   // 4: example.comments.pkg.api.comments.v1.GetCommentsResponse
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_comments_proto_depIdxs = []int32{
	5, // 0: example.comments.pkg.api.comments.v1.Comment.ts:type_name -> google.protobuf.Timestamp
	2, // 1: example.comments.pkg.api.comments.v1.GetCommentsResponse.comments:type_name -> example.comments.pkg.api.comments.v1.Comment
	0, // 2: example.comments.pkg.api.comments.v1.Comments.CreateComment:input_type -> example.comments.pkg.api.comments.v1.CreateCommentRequest
	3, // 3: example.comments.pkg.api.comments.v1.Comments.GetComments:input_type -> example.comments.pkg.api.comments.v1.GetCommentsRequest
	1, // 4: example.comments.pkg.api.comments.v1.Comments.CreateComment:output_type -> example.comments.pkg.api.comments.v1.CreateCommentResponse
	4, // 5: example.comments.pkg.api.comments.v1.Comments.GetComments:output_type -> example.comments.pkg.api.comments.v1.GetCommentsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_comments_proto_init() }
func file_comments_proto_init() {
	if File_comments_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentRequest); i {
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
		file_comments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentResponse); i {
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
		file_comments_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_comments_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentsRequest); i {
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
		file_comments_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentsResponse); i {
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
			RawDescriptor: file_comments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comments_proto_goTypes,
		DependencyIndexes: file_comments_proto_depIdxs,
		MessageInfos:      file_comments_proto_msgTypes,
	}.Build()
	File_comments_proto = out.File
	file_comments_proto_rawDesc = nil
	file_comments_proto_goTypes = nil
	file_comments_proto_depIdxs = nil
}
