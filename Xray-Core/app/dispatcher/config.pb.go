// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: app/dispatcher/config.proto

package dispatcher

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

type SessionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SessionConfig) Reset() {
	*x = SessionConfig{}
	mi := &file_app_dispatcher_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SessionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionConfig) ProtoMessage() {}

func (x *SessionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_app_dispatcher_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionConfig.ProtoReflect.Descriptor instead.
func (*SessionConfig) Descriptor() ([]byte, []int) {
	return file_app_dispatcher_config_proto_rawDescGZIP(), []int{0}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Settings *SessionConfig `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_app_dispatcher_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_dispatcher_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_app_dispatcher_config_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetSettings() *SessionConfig {
	if x != nil {
		return x.Settings
	}
	return nil
}

var File_app_dispatcher_config_proto protoreflect.FileDescriptor

var file_app_dispatcher_config_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x78,
	0x72, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x22, 0x15, 0x0a, 0x0d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x22, 0x48, 0x0a, 0x06, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x3e, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x42, 0x5b, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x50, 0x01,
	0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x74, 0x6c,
	0x73, 0x2f, 0x78, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0xaa, 0x02, 0x13, 0x58, 0x72, 0x61,
	0x79, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_dispatcher_config_proto_rawDescOnce sync.Once
	file_app_dispatcher_config_proto_rawDescData = file_app_dispatcher_config_proto_rawDesc
)

func file_app_dispatcher_config_proto_rawDescGZIP() []byte {
	file_app_dispatcher_config_proto_rawDescOnce.Do(func() {
		file_app_dispatcher_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_dispatcher_config_proto_rawDescData)
	})
	return file_app_dispatcher_config_proto_rawDescData
}

var file_app_dispatcher_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_dispatcher_config_proto_goTypes = []any{
	(*SessionConfig)(nil), // 0: xray.app.dispatcher.SessionConfig
	(*Config)(nil),        // 1: xray.app.dispatcher.Config
}
var file_app_dispatcher_config_proto_depIdxs = []int32{
	0, // 0: xray.app.dispatcher.Config.settings:type_name -> xray.app.dispatcher.SessionConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_dispatcher_config_proto_init() }
func file_app_dispatcher_config_proto_init() {
	if File_app_dispatcher_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_dispatcher_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_dispatcher_config_proto_goTypes,
		DependencyIndexes: file_app_dispatcher_config_proto_depIdxs,
		MessageInfos:      file_app_dispatcher_config_proto_msgTypes,
	}.Build()
	File_app_dispatcher_config_proto = out.File
	file_app_dispatcher_config_proto_rawDesc = nil
	file_app_dispatcher_config_proto_goTypes = nil
	file_app_dispatcher_config_proto_depIdxs = nil
}