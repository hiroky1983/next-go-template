// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: app/sample_service.proto

package appv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetSampleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sample        *Sample                `protobuf:"bytes,1,opt,name=sample,proto3" json:"sample,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSampleRequest) Reset() {
	*x = GetSampleRequest{}
	mi := &file_app_sample_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSampleRequest) ProtoMessage() {}

func (x *GetSampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_sample_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSampleRequest.ProtoReflect.Descriptor instead.
func (*GetSampleRequest) Descriptor() ([]byte, []int) {
	return file_app_sample_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetSampleRequest) GetSample() *Sample {
	if x != nil {
		return x.Sample
	}
	return nil
}

type GetSampleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sample        *Sample                `protobuf:"bytes,1,opt,name=sample,proto3" json:"sample,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSampleResponse) Reset() {
	*x = GetSampleResponse{}
	mi := &file_app_sample_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSampleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSampleResponse) ProtoMessage() {}

func (x *GetSampleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_sample_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSampleResponse.ProtoReflect.Descriptor instead.
func (*GetSampleResponse) Descriptor() ([]byte, []int) {
	return file_app_sample_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetSampleResponse) GetSample() *Sample {
	if x != nil {
		return x.Sample
	}
	return nil
}

var File_app_sample_service_proto protoreflect.FileDescriptor

const file_app_sample_service_proto_rawDesc = "" +
	"\n" +
	"\x18app/sample_service.proto\x12\x06app.v1\x1a\x10app/sample.proto\":\n" +
	"\x10GetSampleRequest\x12&\n" +
	"\x06sample\x18\x01 \x01(\v2\x0e.app.v1.SampleR\x06sample\";\n" +
	"\x11GetSampleResponse\x12&\n" +
	"\x06sample\x18\x01 \x01(\v2\x0e.app.v1.SampleR\x06sample2S\n" +
	"\rSampleService\x12B\n" +
	"\tGetSample\x12\x18.app.v1.GetSampleRequest\x1a\x19.app.v1.GetSampleResponse\"\x00B\x87\x01\n" +
	"\n" +
	"com.app.v1B\x12SampleServiceProtoP\x01Z,github.com/myuser/myproject/gen/app/v1;appv1\xa2\x02\x03AXX\xaa\x02\x06App.V1\xca\x02\x06App\\V1\xe2\x02\x12App\\V1\\GPBMetadata\xea\x02\aApp::V1b\x06proto3"

var (
	file_app_sample_service_proto_rawDescOnce sync.Once
	file_app_sample_service_proto_rawDescData []byte
)

func file_app_sample_service_proto_rawDescGZIP() []byte {
	file_app_sample_service_proto_rawDescOnce.Do(func() {
		file_app_sample_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_app_sample_service_proto_rawDesc), len(file_app_sample_service_proto_rawDesc)))
	})
	return file_app_sample_service_proto_rawDescData
}

var file_app_sample_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_sample_service_proto_goTypes = []any{
	(*GetSampleRequest)(nil),  // 0: app.v1.GetSampleRequest
	(*GetSampleResponse)(nil), // 1: app.v1.GetSampleResponse
	(*Sample)(nil),            // 2: app.v1.Sample
}
var file_app_sample_service_proto_depIdxs = []int32{
	2, // 0: app.v1.GetSampleRequest.sample:type_name -> app.v1.Sample
	2, // 1: app.v1.GetSampleResponse.sample:type_name -> app.v1.Sample
	0, // 2: app.v1.SampleService.GetSample:input_type -> app.v1.GetSampleRequest
	1, // 3: app.v1.SampleService.GetSample:output_type -> app.v1.GetSampleResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_app_sample_service_proto_init() }
func file_app_sample_service_proto_init() {
	if File_app_sample_service_proto != nil {
		return
	}
	file_app_sample_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_app_sample_service_proto_rawDesc), len(file_app_sample_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_sample_service_proto_goTypes,
		DependencyIndexes: file_app_sample_service_proto_depIdxs,
		MessageInfos:      file_app_sample_service_proto_msgTypes,
	}.Build()
	File_app_sample_service_proto = out.File
	file_app_sample_service_proto_goTypes = nil
	file_app_sample_service_proto_depIdxs = nil
}
