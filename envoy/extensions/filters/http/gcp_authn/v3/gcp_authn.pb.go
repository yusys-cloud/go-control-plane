// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: envoy/extensions/filters/http/gcp_authn/v3/gcp_authn.proto

package gcp_authnv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Filter configuration.
type GcpAuthnFilterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The HTTP URI to fetch tokens from GCE Metadata Server(https://cloud.google.com/compute/docs/metadata/overview).
	// The URL format is "http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/identity?audience=[AUDIENCE]"
	HttpUri *v3.HttpUri `protobuf:"bytes,1,opt,name=http_uri,json=httpUri,proto3" json:"http_uri,omitempty"`
	// Retry policy for fetching tokens.
	// This field is optional. If it is not configured, the filter will be fail-closed (i.e., reject the requests).
	RetryPolicy *v3.RetryPolicy `protobuf:"bytes,2,opt,name=retry_policy,json=retryPolicy,proto3" json:"retry_policy,omitempty"`
}

func (x *GcpAuthnFilterConfig) Reset() {
	*x = GcpAuthnFilterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpAuthnFilterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpAuthnFilterConfig) ProtoMessage() {}

func (x *GcpAuthnFilterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpAuthnFilterConfig.ProtoReflect.Descriptor instead.
func (*GcpAuthnFilterConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescGZIP(), []int{0}
}

func (x *GcpAuthnFilterConfig) GetHttpUri() *v3.HttpUri {
	if x != nil {
		return x.HttpUri
	}
	return nil
}

func (x *GcpAuthnFilterConfig) GetRetryPolicy() *v3.RetryPolicy {
	if x != nil {
		return x.RetryPolicy
	}
	return nil
}

// Audience is the URL of the receiving service that performs token authentication.
// It will be provided to the filter through cluster's typed_filter_metadata.
type Audience struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Audience) Reset() {
	*x = Audience{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Audience) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Audience) ProtoMessage() {}

func (x *Audience) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Audience.ProtoReflect.Descriptor instead.
func (*Audience) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescGZIP(), []int{1}
}

func (x *Audience) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x67, 0x63, 0x70, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x67, 0x63, 0x70,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x67, 0x63, 0x70, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x5f, 0x75, 0x72, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x14, 0x47, 0x63, 0x70, 0x41, 0x75,
	0x74, 0x68, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x42, 0x0a, 0x08, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x55, 0x72, 0x69,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x68, 0x74, 0x74, 0x70,
	0x55, 0x72, 0x69, 0x12, 0x44, 0x0a, 0x0c, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0b, 0x72, 0x65,
	0x74, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x25, 0x0a, 0x08, 0x41, 0x75, 0x64,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x42, 0xb2, 0x01, 0x0a, 0x38, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x67, 0x63, 0x70, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x33, 0x42, 0x0d, 0x47,
	0x63, 0x70, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2f, 0x67, 0x63, 0x70, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2f, 0x76,
	0x33, 0x3b, 0x67, 0x63, 0x70, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x76, 0x33, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescData = file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDesc
)

func file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDescData
}

var file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_goTypes = []interface{}{
	(*GcpAuthnFilterConfig)(nil), // 0: envoy.extensions.filters.http.gcp_authn.v3.GcpAuthnFilterConfig
	(*Audience)(nil),             // 1: envoy.extensions.filters.http.gcp_authn.v3.Audience
	(*v3.HttpUri)(nil),           // 2: envoy.config.core.v3.HttpUri
	(*v3.RetryPolicy)(nil),       // 3: envoy.config.core.v3.RetryPolicy
}
var file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.filters.http.gcp_authn.v3.GcpAuthnFilterConfig.http_uri:type_name -> envoy.config.core.v3.HttpUri
	3, // 1: envoy.extensions.filters.http.gcp_authn.v3.GcpAuthnFilterConfig.retry_policy:type_name -> envoy.config.core.v3.RetryPolicy
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_init() }
func file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_init() {
	if File_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcpAuthnFilterConfig); i {
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
		file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Audience); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto = out.File
	file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_rawDesc = nil
	file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_goTypes = nil
	file_envoy_extensions_filters_http_gcp_authn_v3_gcp_authn_proto_depIdxs = nil
}
