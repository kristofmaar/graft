// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.5
// source: tensorflow/core/framework/graph_transfer_info.proto

package graph_transfer_info_go_proto

import (
	types_go_proto "github.com/wamuir/graft/tensorflow/core/framework/types_go_proto"
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

type GraphTransferInfo_Destination int32

const (
	GraphTransferInfo_NOP     GraphTransferInfo_Destination = 0
	GraphTransferInfo_HEXAGON GraphTransferInfo_Destination = 1
)

// Enum value maps for GraphTransferInfo_Destination.
var (
	GraphTransferInfo_Destination_name = map[int32]string{
		0: "NOP",
		1: "HEXAGON",
	}
	GraphTransferInfo_Destination_value = map[string]int32{
		"NOP":     0,
		"HEXAGON": 1,
	}
)

func (x GraphTransferInfo_Destination) Enum() *GraphTransferInfo_Destination {
	p := new(GraphTransferInfo_Destination)
	*p = x
	return p
}

func (x GraphTransferInfo_Destination) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GraphTransferInfo_Destination) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_core_framework_graph_transfer_info_proto_enumTypes[0].Descriptor()
}

func (GraphTransferInfo_Destination) Type() protoreflect.EnumType {
	return &file_tensorflow_core_framework_graph_transfer_info_proto_enumTypes[0]
}

func (x GraphTransferInfo_Destination) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GraphTransferInfo_Destination.Descriptor instead.
func (GraphTransferInfo_Destination) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_graph_transfer_info_proto_rawDescGZIP(), []int{7, 0}
}

type GraphTransferNodeInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId     int32 `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	OutputPort int32 `protobuf:"varint,2,opt,name=output_port,json=outputPort,proto3" json:"output_port,omitempty"`
}

func (x *GraphTransferNodeInput) Reset() {
	*x = GraphTransferNodeInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphTransferNodeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphTransferNodeInput) ProtoMessage() {}

func (x *GraphTransferNodeInput) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphTransferNodeInput.ProtoReflect.Descriptor instead.
func (*GraphTransferNodeInput) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_graph_transfer_info_proto_rawDescGZIP(), []int{0}
}

func (x *GraphTransferNodeInput) GetNodeId() int32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *GraphTransferNodeInput) GetOutputPort() int32 {
	if x != nil {
		return x.OutputPort
	}
	return 0
}

type GraphTransferNodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NodeId      int32  `protobuf:"varint,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	TypeName    string `protobuf:"bytes,3,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	SocOpId     int32  `protobuf:"varint,4,opt,name=soc_op_id,json=socOpId,proto3" json:"soc_op_id,omitempty"`
	PaddingId   int32  `protobuf:"varint,5,opt,name=padding_id,json=paddingId,proto3" json:"padding_id,omitempty"`
	InputCount  int32  `protobuf:"varint,6,opt,name=input_count,json=inputCount,proto3" json:"input_count,omitempty"`
	OutputCount int32  `protobuf:"varint,7,opt,name=output_count,json=outputCount,proto3" json:"output_count,omitempty"`
}

func (x *GraphTransferNodeInfo) Reset() {
	*x = GraphTransferNodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphTransferNodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphTransferNodeInfo) ProtoMessage() {}

func (x *GraphTransferNodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphTransferNodeInfo.ProtoReflect.Descriptor instead.
func (*GraphTransferNodeInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_graph_transfer_info_proto_rawDescGZIP(), []int{1}
}

func (x *GraphTransferNodeInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GraphTransferNodeInfo) GetNodeId() int32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *GraphTransferNodeInfo) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *GraphTransferNodeInfo) GetSocOpId() int32 {
	if x != nil {
		return x.SocOpId
	}
	return 0
}

func (x *GraphTransferNodeInfo) GetPaddingId() int32 {
	if x != nil {
		return x.PaddingId
	}
	return 0
}

func (x *GraphTransferNodeInfo) GetInputCount() int32 {
	if x != nil {
		return x.InputCount
	}
	return 0
}

func (x *GraphTransferNodeInfo) GetOutputCount() int32 {
	if x != nil {
		return x.OutputCount
	}
	return 0
}

type GraphTransferConstNodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NodeId int32                   `protobuf:"varint,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Shape  []int64                 `protobuf:"varint,3,rep,packed,name=shape,proto3" json:"shape,omitempty"`
	Data   []byte                  `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Dtype  types_go_proto.DataType `protobuf:"varint,5,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
}

func (x *GraphTransferConstNodeInfo) Reset() {
	*x = GraphTransferConstNodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphTransferConstNodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphTransferConstNodeInfo) ProtoMessage() {}

func (x *GraphTransferConstNodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphTransferConstNodeInfo.ProtoReflect.Descriptor instead.
func (*GraphTransferConstNodeInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_graph_transfer_info_proto_rawDescGZIP(), []int{2}
}

func (x *GraphTransferConstNodeInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GraphTransferConstNodeInfo) GetNodeId() int32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *GraphTransferConstNodeInfo) GetShape() []int64 {
	if x != nil {
		return x.Shape
	}
	return nil
}

func (x *GraphTransferConstNodeInfo) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GraphTransferConstNodeInfo) GetDtype() types_go_proto.DataType {
	if x != nil {
		return x.Dtype
	}
	return types_go_proto.DataType(0)
}

type GraphTransferNodeInputInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId    int32                     `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	NodeInput []*GraphTransferNodeInput `protobuf:"bytes,2,rep,name=node_input,json=nodeInput,proto3" json:"node_input,omitempty"`
}

func (x *GraphTransferNodeInputInfo) Reset() {
	*x = GraphTransferNodeInputInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphTransferNodeInputInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphTransferNodeInputInfo) ProtoMessage() {}

func (x *GraphTransferNodeInputInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphTransferNodeInputInfo.ProtoReflect.Descriptor instead.
func (*GraphTransferNodeInputInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_graph_transfer_info_proto_rawDescGZIP(), []int{3}
}

func (x *GraphTransferNodeInputInfo) GetNodeId() int32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *GraphTransferNodeInputInfo) GetNodeInput() []*GraphTransferNodeInput {
	if x != nil {
		return x.NodeInput
	}
	return nil
}

type GraphTransferNodeOutputInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId      int32   `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	MaxByteSize []int32 `protobuf:"varint,2,rep,packed,name=max_byte_size,json=maxByteSize,proto3" json:"max_byte_size,omitempty"`
}

func (x *GraphTransferNodeOutputInfo) Reset() {
	*x = GraphTransferNodeOutputInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphTransferNodeOutputInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphTransferNodeOutputInfo) ProtoMessage() {}

func (x *GraphTransferNodeOutputInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphTransferNodeOutputInfo.ProtoReflect.Descriptor instead.
func (*GraphTransferNodeOutputInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_graph_transfer_info_proto_rawDescGZIP(), []int{4}
}

func (x *GraphTransferNodeOutputInfo) GetNodeId() int32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *GraphTransferNodeOutputInfo) GetMaxByteSize() []int32 {
	if x != nil {
		return x.MaxByteSize
	}
	return nil
}

type GraphTransferGraphInputNodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Shape []int64                 `protobuf:"varint,2,rep,packed,name=shape,proto3" json:"shape,omitempty"`
	Dtype types_go_proto.DataType `protobuf:"varint,3,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
}

func (x *GraphTransferGraphInputNodeInfo) Reset() {
	*x = GraphTransferGraphInputNodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphTransferGraphInputNodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphTransferGraphInputNodeInfo) ProtoMessage() {}

func (x *GraphTransferGraphInputNodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphTransferGraphInputNodeInfo.ProtoReflect.Descriptor instead.
func (*GraphTransferGraphInputNodeInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_graph_transfer_info_proto_rawDescGZIP(), []int{5}
}

func (x *GraphTransferGraphInputNodeInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GraphTransferGraphInputNodeInfo) GetShape() []int64 {
	if x != nil {
		return x.Shape
	}
	return nil
}

func (x *GraphTransferGraphInputNodeInfo) GetDtype() types_go_proto.DataType {
	if x != nil {
		return x.Dtype
	}
	return types_go_proto.DataType(0)
}

type GraphTransferGraphOutputNodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Shape []int64                 `protobuf:"varint,2,rep,packed,name=shape,proto3" json:"shape,omitempty"`
	Dtype types_go_proto.DataType `protobuf:"varint,3,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
}

func (x *GraphTransferGraphOutputNodeInfo) Reset() {
	*x = GraphTransferGraphOutputNodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphTransferGraphOutputNodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphTransferGraphOutputNodeInfo) ProtoMessage() {}

func (x *GraphTransferGraphOutputNodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphTransferGraphOutputNodeInfo.ProtoReflect.Descriptor instead.
func (*GraphTransferGraphOutputNodeInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_graph_transfer_info_proto_rawDescGZIP(), []int{6}
}

func (x *GraphTransferGraphOutputNodeInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GraphTransferGraphOutputNodeInfo) GetShape() []int64 {
	if x != nil {
		return x.Shape
	}
	return nil
}

func (x *GraphTransferGraphOutputNodeInfo) GetDtype() types_go_proto.DataType {
	if x != nil {
		return x.Dtype
	}
	return types_go_proto.DataType(0)
}

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type GraphTransferInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeInfo       []*GraphTransferNodeInfo       `protobuf:"bytes,1,rep,name=node_info,json=nodeInfo,proto3" json:"node_info,omitempty"`
	ConstNodeInfo  []*GraphTransferConstNodeInfo  `protobuf:"bytes,2,rep,name=const_node_info,json=constNodeInfo,proto3" json:"const_node_info,omitempty"`
	NodeInputInfo  []*GraphTransferNodeInputInfo  `protobuf:"bytes,3,rep,name=node_input_info,json=nodeInputInfo,proto3" json:"node_input_info,omitempty"`
	NodeOutputInfo []*GraphTransferNodeOutputInfo `protobuf:"bytes,4,rep,name=node_output_info,json=nodeOutputInfo,proto3" json:"node_output_info,omitempty"`
	// Input Node parameters of transferred graph
	GraphInputNodeInfo  []*GraphTransferGraphInputNodeInfo  `protobuf:"bytes,5,rep,name=graph_input_node_info,json=graphInputNodeInfo,proto3" json:"graph_input_node_info,omitempty"`
	GraphOutputNodeInfo []*GraphTransferGraphOutputNodeInfo `protobuf:"bytes,6,rep,name=graph_output_node_info,json=graphOutputNodeInfo,proto3" json:"graph_output_node_info,omitempty"`
	// Destination of graph transfer
	Destination GraphTransferInfo_Destination `protobuf:"varint,7,opt,name=destination,proto3,enum=tensorflow.GraphTransferInfo_Destination" json:"destination,omitempty"`
}

func (x *GraphTransferInfo) Reset() {
	*x = GraphTransferInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphTransferInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphTransferInfo) ProtoMessage() {}

func (x *GraphTransferInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphTransferInfo.ProtoReflect.Descriptor instead.
func (*GraphTransferInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_graph_transfer_info_proto_rawDescGZIP(), []int{7}
}

func (x *GraphTransferInfo) GetNodeInfo() []*GraphTransferNodeInfo {
	if x != nil {
		return x.NodeInfo
	}
	return nil
}

func (x *GraphTransferInfo) GetConstNodeInfo() []*GraphTransferConstNodeInfo {
	if x != nil {
		return x.ConstNodeInfo
	}
	return nil
}

func (x *GraphTransferInfo) GetNodeInputInfo() []*GraphTransferNodeInputInfo {
	if x != nil {
		return x.NodeInputInfo
	}
	return nil
}

func (x *GraphTransferInfo) GetNodeOutputInfo() []*GraphTransferNodeOutputInfo {
	if x != nil {
		return x.NodeOutputInfo
	}
	return nil
}

func (x *GraphTransferInfo) GetGraphInputNodeInfo() []*GraphTransferGraphInputNodeInfo {
	if x != nil {
		return x.GraphInputNodeInfo
	}
	return nil
}

func (x *GraphTransferInfo) GetGraphOutputNodeInfo() []*GraphTransferGraphOutputNodeInfo {
	if x != nil {
		return x.GraphOutputNodeInfo
	}
	return nil
}

func (x *GraphTransferInfo) GetDestination() GraphTransferInfo_Destination {
	if x != nil {
		return x.Destination
	}
	return GraphTransferInfo_NOP
}

var File_tensorflow_core_framework_graph_transfer_info_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_graph_transfer_info_proto_rawDesc = []byte{
	0x0a, 0x33, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x1a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x16, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x22, 0xe0, 0x01, 0x0a,
	0x15, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x09, 0x73, 0x6f, 0x63, 0x5f, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x6f, 0x63, 0x4f, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x9f, 0x01, 0x0a, 0x1a, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x68, 0x61, 0x70, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x64, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x78, 0x0a, 0x1a, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x5a, 0x0a, 0x1b, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x42,
	0x79, 0x74, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x77, 0x0a, 0x1f, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x47, 0x72, 0x61, 0x70, 0x68, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x73,
	0x68, 0x61, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x78, 0x0a, 0x20, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x47, 0x72, 0x61, 0x70, 0x68, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x12, 0x2a,
	0x0a, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x22, 0xfb, 0x04, 0x0a, 0x11, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x3e, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x4e, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x4e, 0x0a, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x51, 0x0a, 0x10, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x6e, 0x6f, 0x64, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x5e, 0x0a, 0x15, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x12, 0x67, 0x72, 0x61, 0x70, 0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x61, 0x0a, 0x16, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x13, 0x67, 0x72, 0x61, 0x70, 0x68, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4b, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x23, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x4f, 0x50, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x48,
	0x45, 0x58, 0x41, 0x47, 0x4f, 0x4e, 0x10, 0x01, 0x42, 0x93, 0x01, 0x0a, 0x18, 0x6f, 0x72, 0x67,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x16, 0x47, 0x72, 0x61, 0x70, 0x68, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_graph_transfer_info_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_graph_transfer_info_proto_rawDescData = file_tensorflow_core_framework_graph_transfer_info_proto_rawDesc
)

func file_tensorflow_core_framework_graph_transfer_info_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_graph_transfer_info_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_graph_transfer_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_graph_transfer_info_proto_rawDescData)
	})
	return file_tensorflow_core_framework_graph_transfer_info_proto_rawDescData
}

var file_tensorflow_core_framework_graph_transfer_info_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tensorflow_core_framework_graph_transfer_info_proto_goTypes = []interface{}{
	(GraphTransferInfo_Destination)(0),       // 0: tensorflow.GraphTransferInfo.Destination
	(*GraphTransferNodeInput)(nil),           // 1: tensorflow.GraphTransferNodeInput
	(*GraphTransferNodeInfo)(nil),            // 2: tensorflow.GraphTransferNodeInfo
	(*GraphTransferConstNodeInfo)(nil),       // 3: tensorflow.GraphTransferConstNodeInfo
	(*GraphTransferNodeInputInfo)(nil),       // 4: tensorflow.GraphTransferNodeInputInfo
	(*GraphTransferNodeOutputInfo)(nil),      // 5: tensorflow.GraphTransferNodeOutputInfo
	(*GraphTransferGraphInputNodeInfo)(nil),  // 6: tensorflow.GraphTransferGraphInputNodeInfo
	(*GraphTransferGraphOutputNodeInfo)(nil), // 7: tensorflow.GraphTransferGraphOutputNodeInfo
	(*GraphTransferInfo)(nil),                // 8: tensorflow.GraphTransferInfo
	(types_go_proto.DataType)(0),             // 9: tensorflow.DataType
}
var file_tensorflow_core_framework_graph_transfer_info_proto_depIdxs = []int32{
	9,  // 0: tensorflow.GraphTransferConstNodeInfo.dtype:type_name -> tensorflow.DataType
	1,  // 1: tensorflow.GraphTransferNodeInputInfo.node_input:type_name -> tensorflow.GraphTransferNodeInput
	9,  // 2: tensorflow.GraphTransferGraphInputNodeInfo.dtype:type_name -> tensorflow.DataType
	9,  // 3: tensorflow.GraphTransferGraphOutputNodeInfo.dtype:type_name -> tensorflow.DataType
	2,  // 4: tensorflow.GraphTransferInfo.node_info:type_name -> tensorflow.GraphTransferNodeInfo
	3,  // 5: tensorflow.GraphTransferInfo.const_node_info:type_name -> tensorflow.GraphTransferConstNodeInfo
	4,  // 6: tensorflow.GraphTransferInfo.node_input_info:type_name -> tensorflow.GraphTransferNodeInputInfo
	5,  // 7: tensorflow.GraphTransferInfo.node_output_info:type_name -> tensorflow.GraphTransferNodeOutputInfo
	6,  // 8: tensorflow.GraphTransferInfo.graph_input_node_info:type_name -> tensorflow.GraphTransferGraphInputNodeInfo
	7,  // 9: tensorflow.GraphTransferInfo.graph_output_node_info:type_name -> tensorflow.GraphTransferGraphOutputNodeInfo
	0,  // 10: tensorflow.GraphTransferInfo.destination:type_name -> tensorflow.GraphTransferInfo.Destination
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_graph_transfer_info_proto_init() }
func file_tensorflow_core_framework_graph_transfer_info_proto_init() {
	if File_tensorflow_core_framework_graph_transfer_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphTransferNodeInput); i {
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
		file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphTransferNodeInfo); i {
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
		file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphTransferConstNodeInfo); i {
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
		file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphTransferNodeInputInfo); i {
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
		file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphTransferNodeOutputInfo); i {
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
		file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphTransferGraphInputNodeInfo); i {
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
		file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphTransferGraphOutputNodeInfo); i {
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
		file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphTransferInfo); i {
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
			RawDescriptor: file_tensorflow_core_framework_graph_transfer_info_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_graph_transfer_info_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_graph_transfer_info_proto_depIdxs,
		EnumInfos:         file_tensorflow_core_framework_graph_transfer_info_proto_enumTypes,
		MessageInfos:      file_tensorflow_core_framework_graph_transfer_info_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_graph_transfer_info_proto = out.File
	file_tensorflow_core_framework_graph_transfer_info_proto_rawDesc = nil
	file_tensorflow_core_framework_graph_transfer_info_proto_goTypes = nil
	file_tensorflow_core_framework_graph_transfer_info_proto_depIdxs = nil
}
