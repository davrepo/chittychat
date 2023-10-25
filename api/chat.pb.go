// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: api/chat.proto

package api

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

type Result int32

const (
	Result_OK                 Result = 0
	Result_GENERAL_ERROR      Result = 1
	Result_UNAUTHORIZED_ERROR Result = 2
)

// Enum value maps for Result.
var (
	Result_name = map[int32]string{
		0: "OK",
		1: "GENERAL_ERROR",
		2: "UNAUTHORIZED_ERROR",
	}
	Result_value = map[string]int32{
		"OK":                 0,
		"GENERAL_ERROR":      1,
		"UNAUTHORIZED_ERROR": 2,
	}
)

func (x Result) Enum() *Result {
	p := new(Result)
	*p = x
	return p
}

func (x Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result) Descriptor() protoreflect.EnumDescriptor {
	return file_api_chat_proto_enumTypes[0].Descriptor()
}

func (Result) Type() protoreflect.EnumType {
	return &file_api_chat_proto_enumTypes[0]
}

func (x Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result.Descriptor instead.
func (Result) EnumDescriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{0}
}

type Lamport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time   uint32 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	NodeId string `protobuf:"bytes,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
}

func (x *Lamport) Reset() {
	*x = Lamport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lamport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lamport) ProtoMessage() {}

func (x *Lamport) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lamport.ProtoReflect.Descriptor instead.
func (*Lamport) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Lamport) GetTime() uint32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Lamport) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

//Minimalisic definition of a course
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lamport *Lamport `protobuf:"bytes,1,opt,name=lamport,proto3" json:"lamport,omitempty"`
	Content string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetLamport() *Lamport {
	if x != nil {
		return x.Lamport
	}
	return nil
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type PublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageHash string   `protobuf:"bytes,1,opt,name=messageHash,proto3" json:"messageHash,omitempty"`
	Lamport     *Lamport `protobuf:"bytes,2,opt,name=lamport,proto3" json:"lamport,omitempty"`
	Result      Result   `protobuf:"varint,3,opt,name=result,proto3,enum=Result" json:"result,omitempty"`
}

func (x *PublishResponse) Reset() {
	*x = PublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResponse) ProtoMessage() {}

func (x *PublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishResponse.ProtoReflect.Descriptor instead.
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{2}
}

func (x *PublishResponse) GetMessageHash() string {
	if x != nil {
		return x.MessageHash
	}
	return ""
}

func (x *PublishResponse) GetLamport() *Lamport {
	if x != nil {
		return x.Lamport
	}
	return nil
}

func (x *PublishResponse) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_OK
}

type BroadcastSubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Receiver string `protobuf:"bytes,1,opt,name=receiver,proto3" json:"receiver,omitempty"`
}

func (x *BroadcastSubscription) Reset() {
	*x = BroadcastSubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastSubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastSubscription) ProtoMessage() {}

func (x *BroadcastSubscription) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastSubscription.ProtoReflect.Descriptor instead.
func (*BroadcastSubscription) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{3}
}

func (x *BroadcastSubscription) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string   `protobuf:"bytes,1,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	Lamport  *Lamport `protobuf:"bytes,2,opt,name=lamport,proto3" json:"lamport,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{4}
}

func (x *JoinRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *JoinRequest) GetLamport() *Lamport {
	if x != nil {
		return x.Lamport
	}
	return nil
}

type JoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId  string   `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	Result  Result   `protobuf:"varint,2,opt,name=result,proto3,enum=Result" json:"result,omitempty"`
	Lamport *Lamport `protobuf:"bytes,3,opt,name=lamport,proto3" json:"lamport,omitempty"`
}

func (x *JoinResponse) Reset() {
	*x = JoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinResponse) ProtoMessage() {}

func (x *JoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinResponse.ProtoReflect.Descriptor instead.
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{5}
}

func (x *JoinResponse) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *JoinResponse) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_OK
}

func (x *JoinResponse) GetLamport() *Lamport {
	if x != nil {
		return x.Lamport
	}
	return nil
}

type LeaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId string   `protobuf:"bytes,1,opt,name=senderId,proto3" json:"senderId,omitempty"`
	Lamport  *Lamport `protobuf:"bytes,2,opt,name=lamport,proto3" json:"lamport,omitempty"`
}

func (x *LeaveRequest) Reset() {
	*x = LeaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveRequest) ProtoMessage() {}

func (x *LeaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveRequest.ProtoReflect.Descriptor instead.
func (*LeaveRequest) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{6}
}

func (x *LeaveRequest) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *LeaveRequest) GetLamport() *Lamport {
	if x != nil {
		return x.Lamport
	}
	return nil
}

type LeaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId  string   `protobuf:"bytes,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	Result  Result   `protobuf:"varint,2,opt,name=result,proto3,enum=Result" json:"result,omitempty"`
	Lamport *Lamport `protobuf:"bytes,3,opt,name=lamport,proto3" json:"lamport,omitempty"`
}

func (x *LeaveResponse) Reset() {
	*x = LeaveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveResponse) ProtoMessage() {}

func (x *LeaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveResponse.ProtoReflect.Descriptor instead.
func (*LeaveResponse) Descriptor() ([]byte, []int) {
	return file_api_chat_proto_rawDescGZIP(), []int{7}
}

func (x *LeaveResponse) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *LeaveResponse) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_OK
}

func (x *LeaveResponse) GetLamport() *Lamport {
	if x != nil {
		return x.Lamport
	}
	return nil
}

var File_api_chat_proto protoreflect.FileDescriptor

var file_api_chat_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x35, 0x0a, 0x07, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x07, 0x6c,
	0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x78, 0x0a, 0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x61,
	0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x22, 0x0a, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x33, 0x0a, 0x15, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x22,
	0x4d, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x6c, 0x61,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4c, 0x61,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x6b,
	0x0a, 0x0c, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x22, 0x0a, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4c, 0x61, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x4e, 0x0a, 0x0c, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4c, 0x61, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x6c, 0x0a, 0x0d, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x22, 0x0a, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2a, 0x3b, 0x0a, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x47,
	0x45, 0x4e, 0x45, 0x52, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x16,
	0x0a, 0x12, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x32, 0xb2, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x0c,
	0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x4a,
	0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x12, 0x0d, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x08,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x42, 0x26, 0x5a, 0x24, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x6e, 0x61, 0x6e, 0x69,
	0x6e, 0x61, 0x73, 0x2f, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x2d, 0x63, 0x68, 0x61, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_chat_proto_rawDescOnce sync.Once
	file_api_chat_proto_rawDescData = file_api_chat_proto_rawDesc
)

func file_api_chat_proto_rawDescGZIP() []byte {
	file_api_chat_proto_rawDescOnce.Do(func() {
		file_api_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_chat_proto_rawDescData)
	})
	return file_api_chat_proto_rawDescData
}

var file_api_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_chat_proto_goTypes = []interface{}{
	(Result)(0),                   // 0: Result
	(*Lamport)(nil),               // 1: Lamport
	(*Message)(nil),               // 2: Message
	(*PublishResponse)(nil),       // 3: PublishResponse
	(*BroadcastSubscription)(nil), // 4: BroadcastSubscription
	(*JoinRequest)(nil),           // 5: JoinRequest
	(*JoinResponse)(nil),          // 6: JoinResponse
	(*LeaveRequest)(nil),          // 7: LeaveRequest
	(*LeaveResponse)(nil),         // 8: LeaveResponse
}
var file_api_chat_proto_depIdxs = []int32{
	1,  // 0: Message.lamport:type_name -> Lamport
	1,  // 1: PublishResponse.lamport:type_name -> Lamport
	0,  // 2: PublishResponse.result:type_name -> Result
	1,  // 3: JoinRequest.lamport:type_name -> Lamport
	0,  // 4: JoinResponse.result:type_name -> Result
	1,  // 5: JoinResponse.lamport:type_name -> Lamport
	1,  // 6: LeaveRequest.lamport:type_name -> Lamport
	0,  // 7: LeaveResponse.result:type_name -> Result
	1,  // 8: LeaveResponse.lamport:type_name -> Lamport
	5,  // 9: ChatService.Join:input_type -> JoinRequest
	7,  // 10: ChatService.Leave:input_type -> LeaveRequest
	2,  // 11: ChatService.Publish:input_type -> Message
	4,  // 12: ChatService.Broadcast:input_type -> BroadcastSubscription
	6,  // 13: ChatService.Join:output_type -> JoinResponse
	8,  // 14: ChatService.Leave:output_type -> LeaveResponse
	3,  // 15: ChatService.Publish:output_type -> PublishResponse
	2,  // 16: ChatService.Broadcast:output_type -> Message
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_chat_proto_init() }
func file_api_chat_proto_init() {
	if File_api_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lamport); i {
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
		file_api_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_api_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishResponse); i {
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
		file_api_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastSubscription); i {
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
		file_api_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
		file_api_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinResponse); i {
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
		file_api_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveRequest); i {
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
		file_api_chat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveResponse); i {
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
			RawDescriptor: file_api_chat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_chat_proto_goTypes,
		DependencyIndexes: file_api_chat_proto_depIdxs,
		EnumInfos:         file_api_chat_proto_enumTypes,
		MessageInfos:      file_api_chat_proto_msgTypes,
	}.Build()
	File_api_chat_proto = out.File
	file_api_chat_proto_rawDesc = nil
	file_api_chat_proto_goTypes = nil
	file_api_chat_proto_depIdxs = nil
}
