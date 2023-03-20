// app_setting.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.2
// source: app_settings.proto

package developer_api

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

type AppSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppScriptsActivated bool `protobuf:"varint,1,opt,name=app_scripts_activated,json=appScriptsActivated,proto3" json:"app_scripts_activated,omitempty"`
}

func (x *AppSettings) Reset() {
	*x = AppSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppSettings) ProtoMessage() {}

func (x *AppSettings) ProtoReflect() protoreflect.Message {
	mi := &file_app_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppSettings.ProtoReflect.Descriptor instead.
func (*AppSettings) Descriptor() ([]byte, []int) {
	return file_app_settings_proto_rawDescGZIP(), []int{0}
}

func (x *AppSettings) GetAppScriptsActivated() bool {
	if x != nil {
		return x.AppScriptsActivated
	}
	return false
}

var File_app_settings_proto protoreflect.FileDescriptor

var file_app_settings_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x41,
	0x70, 0x69, 0x22, 0x41, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x32, 0x0a, 0x15, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x13, 0x61, 0x70, 0x70, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x64, 0x42, 0x23, 0x5a, 0x21, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x65, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_app_settings_proto_rawDescOnce sync.Once
	file_app_settings_proto_rawDescData = file_app_settings_proto_rawDesc
)

func file_app_settings_proto_rawDescGZIP() []byte {
	file_app_settings_proto_rawDescOnce.Do(func() {
		file_app_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_settings_proto_rawDescData)
	})
	return file_app_settings_proto_rawDescData
}

var file_app_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_app_settings_proto_goTypes = []interface{}{
	(*AppSettings)(nil), // 0: DeveloperApi.AppSettings
}
var file_app_settings_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_settings_proto_init() }
func file_app_settings_proto_init() {
	if File_app_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppSettings); i {
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
			RawDescriptor: file_app_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_settings_proto_goTypes,
		DependencyIndexes: file_app_settings_proto_depIdxs,
		MessageInfos:      file_app_settings_proto_msgTypes,
	}.Build()
	File_app_settings_proto = out.File
	file_app_settings_proto_rawDesc = nil
	file_app_settings_proto_goTypes = nil
	file_app_settings_proto_depIdxs = nil
}
