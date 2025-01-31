// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: tensorflow/core/profiler/profiler_options.proto

package profiler_options_proto_go_proto

import (
	for_profiler_protos_go_proto "github.com/wamuir/graft/tensorflow/tsl/profiler/protobuf/for_profiler_protos_go_proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of tensorflow/tsl/profiler/protobuf/profiler_options.proto.

type ProfileOptions_DeviceType = for_profiler_protos_go_proto.ProfileOptions_DeviceType

const ProfileOptions_UNSPECIFIED = for_profiler_protos_go_proto.ProfileOptions_UNSPECIFIED
const ProfileOptions_CPU = for_profiler_protos_go_proto.ProfileOptions_CPU
const ProfileOptions_GPU = for_profiler_protos_go_proto.ProfileOptions_GPU
const ProfileOptions_TPU = for_profiler_protos_go_proto.ProfileOptions_TPU
const ProfileOptions_PLUGGABLE_DEVICE = for_profiler_protos_go_proto.ProfileOptions_PLUGGABLE_DEVICE

var ProfileOptions_DeviceType_name = for_profiler_protos_go_proto.ProfileOptions_DeviceType_name
var ProfileOptions_DeviceType_value = for_profiler_protos_go_proto.ProfileOptions_DeviceType_value

type ProfileOptions = for_profiler_protos_go_proto.ProfileOptions
type RemoteProfilerSessionManagerOptions = for_profiler_protos_go_proto.RemoteProfilerSessionManagerOptions

var File_tensorflow_core_profiler_profiler_options_proto protoreflect.FileDescriptor

var file_tensorflow_core_profiler_profiler_options_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x75,
	0x6d, 0x6d, 0x79, 0x1a, 0x37, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x74, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x5e, 0x5a, 0x5c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_tensorflow_core_profiler_profiler_options_proto_goTypes = []interface{}{}
var file_tensorflow_core_profiler_profiler_options_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_core_profiler_profiler_options_proto_init() }
func file_tensorflow_core_profiler_profiler_options_proto_init() {
	if File_tensorflow_core_profiler_profiler_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_profiler_profiler_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_profiler_profiler_options_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_profiler_profiler_options_proto_depIdxs,
	}.Build()
	File_tensorflow_core_profiler_profiler_options_proto = out.File
	file_tensorflow_core_profiler_profiler_options_proto_rawDesc = nil
	file_tensorflow_core_profiler_profiler_options_proto_goTypes = nil
	file_tensorflow_core_profiler_profiler_options_proto_depIdxs = nil
}
