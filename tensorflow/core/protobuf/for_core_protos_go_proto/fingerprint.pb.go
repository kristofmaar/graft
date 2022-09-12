// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: tensorflow/core/protobuf/fingerprint.proto

package for_core_protos_go_proto

import (
	versions_go_proto "github.com/wamuir/graft/tensorflow/core/framework/versions_go_proto"
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

// Protocol buffer representing a SavedModel Fingerprint.
//
// If there are multiple MetaGraphDefs in the SavedModel, the FingerprintDef
// corresponds to the first one.
type FingerprintDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hash of the graph_def, referred to as a "checksum".
	GraphDefChecksum uint64 `protobuf:"varint,1,opt,name=graph_def_checksum,json=graphDefChecksum,proto3" json:"graph_def_checksum,omitempty"`
	// Hash of regularized graph_def.
	GraphDefProgramHash uint64 `protobuf:"varint,2,opt,name=graph_def_program_hash,json=graphDefProgramHash,proto3" json:"graph_def_program_hash,omitempty"`
	// Hash of the regularized (sorted) SignatureDefs.
	SignatureDefHash uint64 `protobuf:"varint,3,opt,name=signature_def_hash,json=signatureDefHash,proto3" json:"signature_def_hash,omitempty"`
	// Hash of the regularized SavedObjectGraph.
	SavedObjectGraphHash uint64 `protobuf:"varint,4,opt,name=saved_object_graph_hash,json=savedObjectGraphHash,proto3" json:"saved_object_graph_hash,omitempty"`
	// Hash of the checkpoint.
	CheckpointHash uint64 `protobuf:"varint,5,opt,name=checkpoint_hash,json=checkpointHash,proto3" json:"checkpoint_hash,omitempty"`
	// Version specification of the fingerprint.
	Version *versions_go_proto.VersionDef `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *FingerprintDef) Reset() {
	*x = FingerprintDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_fingerprint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FingerprintDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FingerprintDef) ProtoMessage() {}

func (x *FingerprintDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_fingerprint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FingerprintDef.ProtoReflect.Descriptor instead.
func (*FingerprintDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_fingerprint_proto_rawDescGZIP(), []int{0}
}

func (x *FingerprintDef) GetGraphDefChecksum() uint64 {
	if x != nil {
		return x.GraphDefChecksum
	}
	return 0
}

func (x *FingerprintDef) GetGraphDefProgramHash() uint64 {
	if x != nil {
		return x.GraphDefProgramHash
	}
	return 0
}

func (x *FingerprintDef) GetSignatureDefHash() uint64 {
	if x != nil {
		return x.SignatureDefHash
	}
	return 0
}

func (x *FingerprintDef) GetSavedObjectGraphHash() uint64 {
	if x != nil {
		return x.SavedObjectGraphHash
	}
	return 0
}

func (x *FingerprintDef) GetCheckpointHash() uint64 {
	if x != nil {
		return x.CheckpointHash
	}
	return 0
}

func (x *FingerprintDef) GetVersion() *versions_go_proto.VersionDef {
	if x != nil {
		return x.Version
	}
	return nil
}

var File_tensorflow_core_protobuf_fingerprint_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_fingerprint_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x6e, 0x67, 0x65,
	0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x28, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb3, 0x02, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x44, 0x65, 0x66, 0x12, 0x2c, 0x0a, 0x12, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x64,
	0x65, 0x66, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x10, 0x67, 0x72, 0x61, 0x70, 0x68, 0x44, 0x65, 0x66, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x12, 0x33, 0x0a, 0x16, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x64, 0x65, 0x66,
	0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x13, 0x67, 0x72, 0x61, 0x70, 0x68, 0x44, 0x65, 0x66, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x44,
	0x65, 0x66, 0x48, 0x61, 0x73, 0x68, 0x12, 0x35, 0x0a, 0x17, 0x73, 0x61, 0x76, 0x65, 0x64, 0x5f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x73, 0x61, 0x76, 0x65, 0x64, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x61, 0x70, 0x68, 0x48, 0x61, 0x73, 0x68, 0x12, 0x27, 0x0a,
	0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x30, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x89, 0x01, 0x0a, 0x18, 0x6f, 0x72, 0x67,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x11, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_fingerprint_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_fingerprint_proto_rawDescData = file_tensorflow_core_protobuf_fingerprint_proto_rawDesc
)

func file_tensorflow_core_protobuf_fingerprint_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_fingerprint_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_fingerprint_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_fingerprint_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_fingerprint_proto_rawDescData
}

var file_tensorflow_core_protobuf_fingerprint_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_core_protobuf_fingerprint_proto_goTypes = []interface{}{
	(*FingerprintDef)(nil),               // 0: tensorflow.FingerprintDef
	(*versions_go_proto.VersionDef)(nil), // 1: tensorflow.VersionDef
}
var file_tensorflow_core_protobuf_fingerprint_proto_depIdxs = []int32{
	1, // 0: tensorflow.FingerprintDef.version:type_name -> tensorflow.VersionDef
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_fingerprint_proto_init() }
func file_tensorflow_core_protobuf_fingerprint_proto_init() {
	if File_tensorflow_core_protobuf_fingerprint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_fingerprint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FingerprintDef); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_fingerprint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_fingerprint_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_fingerprint_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_fingerprint_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_fingerprint_proto = out.File
	file_tensorflow_core_protobuf_fingerprint_proto_rawDesc = nil
	file_tensorflow_core_protobuf_fingerprint_proto_goTypes = nil
	file_tensorflow_core_protobuf_fingerprint_proto_depIdxs = nil
}
