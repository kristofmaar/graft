// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.5
// source: tensorflow/core/protobuf/status.proto

package for_core_protos_go_proto

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

// If included as a payload, this message flags the Status to be a "derived"
// Status. Used by StatusGroup to ignore certain Statuses when reporting
// errors to end users.
type DerivedStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DerivedStatus) Reset() {
	*x = DerivedStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DerivedStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DerivedStatus) ProtoMessage() {}

func (x *DerivedStatus) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DerivedStatus.ProtoReflect.Descriptor instead.
func (*DerivedStatus) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_status_proto_rawDescGZIP(), []int{0}
}

// If included as a payload, this message contains associated source location
// for the error Status.
type StackTracePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StackFrames []*StackTracePayload_StackFrame `protobuf:"bytes,1,rep,name=stack_frames,json=stackFrames,proto3" json:"stack_frames,omitempty"`
}

func (x *StackTracePayload) Reset() {
	*x = StackTracePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackTracePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackTracePayload) ProtoMessage() {}

func (x *StackTracePayload) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackTracePayload.ProtoReflect.Descriptor instead.
func (*StackTracePayload) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_status_proto_rawDescGZIP(), []int{1}
}

func (x *StackTracePayload) GetStackFrames() []*StackTracePayload_StackFrame {
	if x != nil {
		return x.StackFrames
	}
	return nil
}

type StackTracePayload_StackFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName     string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	LineNumber   int32  `protobuf:"varint,2,opt,name=line_number,json=lineNumber,proto3" json:"line_number,omitempty"`
	FunctionName string `protobuf:"bytes,3,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
}

func (x *StackTracePayload_StackFrame) Reset() {
	*x = StackTracePayload_StackFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackTracePayload_StackFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackTracePayload_StackFrame) ProtoMessage() {}

func (x *StackTracePayload_StackFrame) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackTracePayload_StackFrame.ProtoReflect.Descriptor instead.
func (*StackTracePayload_StackFrame) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_status_proto_rawDescGZIP(), []int{1, 0}
}

func (x *StackTracePayload_StackFrame) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *StackTracePayload_StackFrame) GetLineNumber() int32 {
	if x != nil {
		return x.LineNumber
	}
	return 0
}

func (x *StackTracePayload_StackFrame) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

var File_tensorflow_core_protobuf_status_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_status_proto_rawDesc = []byte{
	0x0a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x22, 0x0f, 0x0a, 0x0d, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0xd1, 0x01, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x4b, 0x0a, 0x0c, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x74,
	0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x6f, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x57, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_status_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_status_proto_rawDescData = file_tensorflow_core_protobuf_status_proto_rawDesc
)

func file_tensorflow_core_protobuf_status_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_status_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_status_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_status_proto_rawDescData
}

var file_tensorflow_core_protobuf_status_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_core_protobuf_status_proto_goTypes = []interface{}{
	(*DerivedStatus)(nil),                // 0: tensorflow.DerivedStatus
	(*StackTracePayload)(nil),            // 1: tensorflow.StackTracePayload
	(*StackTracePayload_StackFrame)(nil), // 2: tensorflow.StackTracePayload.StackFrame
}
var file_tensorflow_core_protobuf_status_proto_depIdxs = []int32{
	2, // 0: tensorflow.StackTracePayload.stack_frames:type_name -> tensorflow.StackTracePayload.StackFrame
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_status_proto_init() }
func file_tensorflow_core_protobuf_status_proto_init() {
	if File_tensorflow_core_protobuf_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DerivedStatus); i {
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
		file_tensorflow_core_protobuf_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackTracePayload); i {
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
		file_tensorflow_core_protobuf_status_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackTracePayload_StackFrame); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_status_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_status_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_status_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_status_proto = out.File
	file_tensorflow_core_protobuf_status_proto_rawDesc = nil
	file_tensorflow_core_protobuf_status_proto_goTypes = nil
	file_tensorflow_core_protobuf_status_proto_depIdxs = nil
}
