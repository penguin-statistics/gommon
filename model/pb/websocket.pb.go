// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.7
// source: websocket.proto

package pb

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

type Language int32

const (
	Language_ZH_CN Language = 0
	Language_EN_US Language = 1
	Language_JA_JP Language = 2
	Language_KO_KR Language = 3
	Language_OTHER Language = 4
)

// Enum value maps for Language.
var (
	Language_name = map[int32]string{
		0: "ZH_CN",
		1: "EN_US",
		2: "JA_JP",
		3: "KO_KR",
		4: "OTHER",
	}
	Language_value = map[string]int32{
		"ZH_CN": 0,
		"EN_US": 1,
		"JA_JP": 2,
		"KO_KR": 3,
		"OTHER": 4,
	}
)

func (x Language) Enum() *Language {
	p := new(Language)
	*p = x
	return p
}

func (x Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language) Descriptor() protoreflect.EnumDescriptor {
	return file_websocket_proto_enumTypes[0].Descriptor()
}

func (Language) Type() protoreflect.EnumType {
	return &file_websocket_proto_enumTypes[0]
}

func (x Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language.Descriptor instead.
func (Language) EnumDescriptor() ([]byte, []int) {
	return file_websocket_proto_rawDescGZIP(), []int{0}
}

type MessageType int32

const (
	// default value, leave for unknown
	MessageType_UNKNOWN                       MessageType = 0
	MessageType_PROBE_NAVIGATED               MessageType = 1
	MessageType_PROBE_ENTERED_SEARCH_RESULT   MessageType = 2
	MessageType_PROBE_EXECUTED_ADVANCED_QUERY MessageType = 3
	MessageType_PROBE_SERVER_ACK              MessageType = 64
	// server push messages. start from 1 << 8
	/// uses MatrixUpdateMessage
	MessageType_MATRIX_UPDATE_MESSAGE MessageType = 256
	// request-response pattern messages. start from 1 << 12
	/// uses MatrixUpdateSubscribeReq
	MessageType_MATRIX_UPDATE_SUBSCRIBE_REQ MessageType = 4096
	/// uses MatrixUpdateSubscribeResp
	MessageType_MATRIX_UPDATE_SUBSCRIBE_RESP MessageType = 4097
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0:    "UNKNOWN",
		1:    "PROBE_NAVIGATED",
		2:    "PROBE_ENTERED_SEARCH_RESULT",
		3:    "PROBE_EXECUTED_ADVANCED_QUERY",
		64:   "PROBE_SERVER_ACK",
		256:  "MATRIX_UPDATE_MESSAGE",
		4096: "MATRIX_UPDATE_SUBSCRIBE_REQ",
		4097: "MATRIX_UPDATE_SUBSCRIBE_RESP",
	}
	MessageType_value = map[string]int32{
		"UNKNOWN":                       0,
		"PROBE_NAVIGATED":               1,
		"PROBE_ENTERED_SEARCH_RESULT":   2,
		"PROBE_EXECUTED_ADVANCED_QUERY": 3,
		"PROBE_SERVER_ACK":              64,
		"MATRIX_UPDATE_MESSAGE":         256,
		"MATRIX_UPDATE_SUBSCRIBE_REQ":   4096,
		"MATRIX_UPDATE_SUBSCRIBE_RESP":  4097,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_websocket_proto_enumTypes[1].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_websocket_proto_enumTypes[1]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_websocket_proto_rawDescGZIP(), []int{1}
}

type Skeleton struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *Skeleton) Reset() {
	*x = Skeleton{}
	if protoimpl.UnsafeEnabled {
		mi := &file_websocket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Skeleton) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Skeleton) ProtoMessage() {}

func (x *Skeleton) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Skeleton.ProtoReflect.Descriptor instead.
func (*Skeleton) Descriptor() ([]byte, []int) {
	return file_websocket_proto_rawDescGZIP(), []int{0}
}

func (x *Skeleton) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=MessageType" json:"type,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_websocket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_websocket_proto_rawDescGZIP(), []int{1}
}

func (x *Header) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_UNKNOWN
}

type MatrixUpdateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header   *Header                        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Segments []*MatrixUpdateMessage_Element `protobuf:"bytes,2,rep,name=segments,proto3" json:"segments,omitempty"`
}

func (x *MatrixUpdateMessage) Reset() {
	*x = MatrixUpdateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_websocket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatrixUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatrixUpdateMessage) ProtoMessage() {}

func (x *MatrixUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatrixUpdateMessage.ProtoReflect.Descriptor instead.
func (*MatrixUpdateMessage) Descriptor() ([]byte, []int) {
	return file_websocket_proto_rawDescGZIP(), []int{2}
}

func (x *MatrixUpdateMessage) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *MatrixUpdateMessage) GetSegments() []*MatrixUpdateMessage_Element {
	if x != nil {
		return x.Segments
	}
	return nil
}

type MatrixUpdateSubscribeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Server Server  `protobuf:"varint,2,opt,name=server,proto3,enum=Server" json:"server,omitempty"`
	// Types that are assignable to Id:
	//	*MatrixUpdateSubscribeReq_StageId
	//	*MatrixUpdateSubscribeReq_ItemId
	Id isMatrixUpdateSubscribeReq_Id `protobuf_oneof:"id"`
}

func (x *MatrixUpdateSubscribeReq) Reset() {
	*x = MatrixUpdateSubscribeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_websocket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatrixUpdateSubscribeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatrixUpdateSubscribeReq) ProtoMessage() {}

func (x *MatrixUpdateSubscribeReq) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatrixUpdateSubscribeReq.ProtoReflect.Descriptor instead.
func (*MatrixUpdateSubscribeReq) Descriptor() ([]byte, []int) {
	return file_websocket_proto_rawDescGZIP(), []int{3}
}

func (x *MatrixUpdateSubscribeReq) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *MatrixUpdateSubscribeReq) GetServer() Server {
	if x != nil {
		return x.Server
	}
	return Server_CN
}

func (m *MatrixUpdateSubscribeReq) GetId() isMatrixUpdateSubscribeReq_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (x *MatrixUpdateSubscribeReq) GetStageId() uint32 {
	if x, ok := x.GetId().(*MatrixUpdateSubscribeReq_StageId); ok {
		return x.StageId
	}
	return 0
}

func (x *MatrixUpdateSubscribeReq) GetItemId() uint32 {
	if x, ok := x.GetId().(*MatrixUpdateSubscribeReq_ItemId); ok {
		return x.ItemId
	}
	return 0
}

type isMatrixUpdateSubscribeReq_Id interface {
	isMatrixUpdateSubscribeReq_Id()
}

type MatrixUpdateSubscribeReq_StageId struct {
	StageId uint32 `protobuf:"varint,3,opt,name=stage_id,json=stageId,proto3,oneof"`
}

type MatrixUpdateSubscribeReq_ItemId struct {
	ItemId uint32 `protobuf:"varint,4,opt,name=item_id,json=itemId,proto3,oneof"`
}

func (*MatrixUpdateSubscribeReq_StageId) isMatrixUpdateSubscribeReq_Id() {}

func (*MatrixUpdateSubscribeReq_ItemId) isMatrixUpdateSubscribeReq_Id() {}

type MatrixUpdateSubscribeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Error  string  `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *MatrixUpdateSubscribeResp) Reset() {
	*x = MatrixUpdateSubscribeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_websocket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatrixUpdateSubscribeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatrixUpdateSubscribeResp) ProtoMessage() {}

func (x *MatrixUpdateSubscribeResp) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatrixUpdateSubscribeResp.ProtoReflect.Descriptor instead.
func (*MatrixUpdateSubscribeResp) Descriptor() ([]byte, []int) {
	return file_websocket_proto_rawDescGZIP(), []int{4}
}

func (x *MatrixUpdateSubscribeResp) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *MatrixUpdateSubscribeResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type MatrixUpdateMessage_Element struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server   Server `protobuf:"varint,1,opt,name=server,proto3,enum=Server" json:"server,omitempty"`
	StageId  uint32 `protobuf:"varint,2,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	ItemId   uint32 `protobuf:"varint,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Quantity uint64 `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Times    uint64 `protobuf:"varint,5,opt,name=times,proto3" json:"times,omitempty"`
}

func (x *MatrixUpdateMessage_Element) Reset() {
	*x = MatrixUpdateMessage_Element{}
	if protoimpl.UnsafeEnabled {
		mi := &file_websocket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatrixUpdateMessage_Element) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatrixUpdateMessage_Element) ProtoMessage() {}

func (x *MatrixUpdateMessage_Element) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatrixUpdateMessage_Element.ProtoReflect.Descriptor instead.
func (*MatrixUpdateMessage_Element) Descriptor() ([]byte, []int) {
	return file_websocket_proto_rawDescGZIP(), []int{2, 0}
}

func (x *MatrixUpdateMessage_Element) GetServer() Server {
	if x != nil {
		return x.Server
	}
	return Server_CN
}

func (x *MatrixUpdateMessage_Element) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *MatrixUpdateMessage_Element) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *MatrixUpdateMessage_Element) GetQuantity() uint64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *MatrixUpdateMessage_Element) GetTimes() uint64 {
	if x != nil {
		return x.Times
	}
	return 0
}

var File_websocket_proto protoreflect.FileDescriptor

var file_websocket_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2b, 0x0a, 0x08, 0x53, 0x6b, 0x65, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x2a, 0x0a, 0x06,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x83, 0x02, 0x0a, 0x13, 0x4d, 0x61, 0x74,
	0x72, 0x69, 0x78, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1f, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x38, 0x0a, 0x08, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x08, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x90, 0x01, 0x0a, 0x07,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x22, 0x9a,
	0x01, 0x0a, 0x18, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x06,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1b, 0x0a,
	0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48,
	0x00, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x07, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x06, 0x69,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x42, 0x04, 0x0a, 0x02, 0x69, 0x64, 0x22, 0x52, 0x0a, 0x19, 0x4d,
	0x61, 0x74, 0x72, 0x69, 0x78, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a,
	0x41, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x5a,
	0x48, 0x5f, 0x43, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4e, 0x5f, 0x55, 0x53, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x4a, 0x41, 0x5f, 0x4a, 0x50, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05,
	0x4b, 0x4f, 0x5f, 0x4b, 0x52, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52,
	0x10, 0x04, 0x2a, 0xea, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x13, 0x0a, 0x0f, 0x50, 0x52, 0x4f, 0x42, 0x45, 0x5f, 0x4e, 0x41, 0x56, 0x49, 0x47, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x52, 0x4f, 0x42, 0x45, 0x5f, 0x45, 0x4e,
	0x54, 0x45, 0x52, 0x45, 0x44, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x52, 0x45, 0x53,
	0x55, 0x4c, 0x54, 0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x50, 0x52, 0x4f, 0x42, 0x45, 0x5f, 0x45,
	0x58, 0x45, 0x43, 0x55, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x44, 0x56, 0x41, 0x4e, 0x43, 0x45, 0x44,
	0x5f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x52, 0x4f, 0x42,
	0x45, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x4b, 0x10, 0x40, 0x12, 0x1a,
	0x0a, 0x15, 0x4d, 0x41, 0x54, 0x52, 0x49, 0x58, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x80, 0x02, 0x12, 0x20, 0x0a, 0x1b, 0x4d, 0x41,
	0x54, 0x52, 0x49, 0x58, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x55, 0x42, 0x53,
	0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x10, 0x80, 0x20, 0x12, 0x21, 0x0a, 0x1c,
	0x4d, 0x41, 0x54, 0x52, 0x49, 0x58, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x55,
	0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x10, 0x81, 0x20, 0x42,
	0x1d, 0x5a, 0x1b, 0x65, 0x78, 0x75, 0x73, 0x69, 0x61, 0x69, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x67,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_websocket_proto_rawDescOnce sync.Once
	file_websocket_proto_rawDescData = file_websocket_proto_rawDesc
)

func file_websocket_proto_rawDescGZIP() []byte {
	file_websocket_proto_rawDescOnce.Do(func() {
		file_websocket_proto_rawDescData = protoimpl.X.CompressGZIP(file_websocket_proto_rawDescData)
	})
	return file_websocket_proto_rawDescData
}

var file_websocket_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_websocket_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_websocket_proto_goTypes = []interface{}{
	(Language)(0),                       // 0: Language
	(MessageType)(0),                    // 1: MessageType
	(*Skeleton)(nil),                    // 2: Skeleton
	(*Header)(nil),                      // 3: Header
	(*MatrixUpdateMessage)(nil),         // 4: MatrixUpdateMessage
	(*MatrixUpdateSubscribeReq)(nil),    // 5: MatrixUpdateSubscribeReq
	(*MatrixUpdateSubscribeResp)(nil),   // 6: MatrixUpdateSubscribeResp
	(*MatrixUpdateMessage_Element)(nil), // 7: MatrixUpdateMessage.Element
	(Server)(0),                         // 8: Server
}
var file_websocket_proto_depIdxs = []int32{
	3, // 0: Skeleton.header:type_name -> Header
	1, // 1: Header.type:type_name -> MessageType
	3, // 2: MatrixUpdateMessage.header:type_name -> Header
	7, // 3: MatrixUpdateMessage.segments:type_name -> MatrixUpdateMessage.Element
	3, // 4: MatrixUpdateSubscribeReq.header:type_name -> Header
	8, // 5: MatrixUpdateSubscribeReq.server:type_name -> Server
	3, // 6: MatrixUpdateSubscribeResp.header:type_name -> Header
	8, // 7: MatrixUpdateMessage.Element.server:type_name -> Server
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_websocket_proto_init() }
func file_websocket_proto_init() {
	if File_websocket_proto != nil {
		return
	}
	file_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_websocket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Skeleton); i {
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
		file_websocket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_websocket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatrixUpdateMessage); i {
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
		file_websocket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatrixUpdateSubscribeReq); i {
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
		file_websocket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatrixUpdateSubscribeResp); i {
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
		file_websocket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatrixUpdateMessage_Element); i {
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
	file_websocket_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*MatrixUpdateSubscribeReq_StageId)(nil),
		(*MatrixUpdateSubscribeReq_ItemId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_websocket_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_websocket_proto_goTypes,
		DependencyIndexes: file_websocket_proto_depIdxs,
		EnumInfos:         file_websocket_proto_enumTypes,
		MessageInfos:      file_websocket_proto_msgTypes,
	}.Build()
	File_websocket_proto = out.File
	file_websocket_proto_rawDesc = nil
	file_websocket_proto_goTypes = nil
	file_websocket_proto_depIdxs = nil
}