// grpc 流式数据处理

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.14.0
// source: proto/pb_stream/stream_manager.proto

package pb_stream

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 包含人名的一个请求消息
type StreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uuid       string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Ip         string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Port       int32  `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Status     string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	StreamData []byte `protobuf:"bytes,6,opt,name=stream_data,json=streamData,proto3" json:"stream_data,omitempty"` //字节数组
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pb_stream_stream_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pb_stream_stream_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_proto_pb_stream_stream_manager_proto_rawDescGZIP(), []int{0}
}

func (x *StreamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StreamRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *StreamRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *StreamRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *StreamRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *StreamRequest) GetStreamData() []byte {
	if x != nil {
		return x.StreamData
	}
	return nil
}

// 包含问候语的响应消息
type StreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code         int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message      string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	StreamData   []byte `protobuf:"bytes,6,opt,name=stream_data,json=streamData,proto3" json:"stream_data,omitempty"` //字节数组
	ErrorMessage string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	Total        int32  `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *StreamResponse) Reset() {
	*x = StreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pb_stream_stream_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse) ProtoMessage() {}

func (x *StreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pb_stream_stream_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse.ProtoReflect.Descriptor instead.
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return file_proto_pb_stream_stream_manager_proto_rawDescGZIP(), []int{1}
}

func (x *StreamResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *StreamResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StreamResponse) GetStreamData() []byte {
	if x != nil {
		return x.StreamData
	}
	return nil
}

func (x *StreamResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *StreamResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_proto_pb_stream_stream_manager_proto protoreflect.FileDescriptor

var file_proto_pb_stream_stream_manager_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x94, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x22, 0x9a, 0x01,
	0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xf2, 0x03, 0x0a, 0x0e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a,
	0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x1d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x04, 0x50, 0x75, 0x6c, 0x6c, 0x12, 0x1d,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4e, 0x0a, 0x0b, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x51, 0x0a, 0x0a, 0x50, 0x75, 0x73, 0x68, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1d, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x12, 0x51, 0x0a, 0x0a, 0x50, 0x75, 0x6c, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x1d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x58, 0x0a, 0x11, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1d, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x1b, 0x5a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x5f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x3b, 0x70, 0x62, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_pb_stream_stream_manager_proto_rawDescOnce sync.Once
	file_proto_pb_stream_stream_manager_proto_rawDescData = file_proto_pb_stream_stream_manager_proto_rawDesc
)

func file_proto_pb_stream_stream_manager_proto_rawDescGZIP() []byte {
	file_proto_pb_stream_stream_manager_proto_rawDescOnce.Do(func() {
		file_proto_pb_stream_stream_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_pb_stream_stream_manager_proto_rawDescData)
	})
	return file_proto_pb_stream_stream_manager_proto_rawDescData
}

var file_proto_pb_stream_stream_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_pb_stream_stream_manager_proto_goTypes = []interface{}{
	(*StreamRequest)(nil),  // 0: manager.stream.StreamRequest
	(*StreamResponse)(nil), // 1: manager.stream.StreamResponse
}
var file_proto_pb_stream_stream_manager_proto_depIdxs = []int32{
	0, // 0: manager.stream.Stream_Service.Push:input_type -> manager.stream.StreamRequest
	0, // 1: manager.stream.Stream_Service.Pull:input_type -> manager.stream.StreamRequest
	0, // 2: manager.stream.Stream_Service.Bidirection:input_type -> manager.stream.StreamRequest
	0, // 3: manager.stream.Stream_Service.PushStream:input_type -> manager.stream.StreamRequest
	0, // 4: manager.stream.Stream_Service.PullStream:input_type -> manager.stream.StreamRequest
	0, // 5: manager.stream.Stream_Service.BidirectionStream:input_type -> manager.stream.StreamRequest
	1, // 6: manager.stream.Stream_Service.Push:output_type -> manager.stream.StreamResponse
	1, // 7: manager.stream.Stream_Service.Pull:output_type -> manager.stream.StreamResponse
	1, // 8: manager.stream.Stream_Service.Bidirection:output_type -> manager.stream.StreamResponse
	1, // 9: manager.stream.Stream_Service.PushStream:output_type -> manager.stream.StreamResponse
	1, // 10: manager.stream.Stream_Service.PullStream:output_type -> manager.stream.StreamResponse
	1, // 11: manager.stream.Stream_Service.BidirectionStream:output_type -> manager.stream.StreamResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_pb_stream_stream_manager_proto_init() }
func file_proto_pb_stream_stream_manager_proto_init() {
	if File_proto_pb_stream_stream_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_pb_stream_stream_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamRequest); i {
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
		file_proto_pb_stream_stream_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResponse); i {
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
			RawDescriptor: file_proto_pb_stream_stream_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_pb_stream_stream_manager_proto_goTypes,
		DependencyIndexes: file_proto_pb_stream_stream_manager_proto_depIdxs,
		MessageInfos:      file_proto_pb_stream_stream_manager_proto_msgTypes,
	}.Build()
	File_proto_pb_stream_stream_manager_proto = out.File
	file_proto_pb_stream_stream_manager_proto_rawDesc = nil
	file_proto_pb_stream_stream_manager_proto_goTypes = nil
	file_proto_pb_stream_stream_manager_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// Stream_ServiceClient is the client API for Stream_Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Stream_ServiceClient interface {
	//服务端推送
	Push(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*StreamResponse, error)
	//客户端推送
	Pull(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*StreamResponse, error)
	//客户端与 服务端 互相 推送
	Bidirection(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*StreamResponse, error)
	PushStream(ctx context.Context, opts ...grpc.CallOption) (Stream_Service_PushStreamClient, error)
	//客户端推送
	PullStream(ctx context.Context, opts ...grpc.CallOption) (Stream_Service_PullStreamClient, error)
	//客户端与 服务端 互相 推送
	BidirectionStream(ctx context.Context, opts ...grpc.CallOption) (Stream_Service_BidirectionStreamClient, error)
}

type stream_ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStream_ServiceClient(cc grpc.ClientConnInterface) Stream_ServiceClient {
	return &stream_ServiceClient{cc}
}

func (c *stream_ServiceClient) Push(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*StreamResponse, error) {
	out := new(StreamResponse)
	err := c.cc.Invoke(ctx, "/manager.stream.Stream_Service/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stream_ServiceClient) Pull(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*StreamResponse, error) {
	out := new(StreamResponse)
	err := c.cc.Invoke(ctx, "/manager.stream.Stream_Service/Pull", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stream_ServiceClient) Bidirection(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*StreamResponse, error) {
	out := new(StreamResponse)
	err := c.cc.Invoke(ctx, "/manager.stream.Stream_Service/Bidirection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stream_ServiceClient) PushStream(ctx context.Context, opts ...grpc.CallOption) (Stream_Service_PushStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Stream_Service_serviceDesc.Streams[0], "/manager.stream.Stream_Service/PushStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &stream_ServicePushStreamClient{stream}
	return x, nil
}

type Stream_Service_PushStreamClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type stream_ServicePushStreamClient struct {
	grpc.ClientStream
}

func (x *stream_ServicePushStreamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *stream_ServicePushStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *stream_ServiceClient) PullStream(ctx context.Context, opts ...grpc.CallOption) (Stream_Service_PullStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Stream_Service_serviceDesc.Streams[1], "/manager.stream.Stream_Service/PullStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &stream_ServicePullStreamClient{stream}
	return x, nil
}

type Stream_Service_PullStreamClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type stream_ServicePullStreamClient struct {
	grpc.ClientStream
}

func (x *stream_ServicePullStreamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *stream_ServicePullStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *stream_ServiceClient) BidirectionStream(ctx context.Context, opts ...grpc.CallOption) (Stream_Service_BidirectionStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Stream_Service_serviceDesc.Streams[2], "/manager.stream.Stream_Service/BidirectionStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &stream_ServiceBidirectionStreamClient{stream}
	return x, nil
}

type Stream_Service_BidirectionStreamClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type stream_ServiceBidirectionStreamClient struct {
	grpc.ClientStream
}

func (x *stream_ServiceBidirectionStreamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *stream_ServiceBidirectionStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Stream_ServiceServer is the server API for Stream_Service service.
type Stream_ServiceServer interface {
	//服务端推送
	Push(context.Context, *StreamRequest) (*StreamResponse, error)
	//客户端推送
	Pull(context.Context, *StreamRequest) (*StreamResponse, error)
	//客户端与 服务端 互相 推送
	Bidirection(context.Context, *StreamRequest) (*StreamResponse, error)
	PushStream(Stream_Service_PushStreamServer) error
	//客户端推送
	PullStream(Stream_Service_PullStreamServer) error
	//客户端与 服务端 互相 推送
	BidirectionStream(Stream_Service_BidirectionStreamServer) error
}

// UnimplementedStream_ServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStream_ServiceServer struct {
}

func (*UnimplementedStream_ServiceServer) Push(context.Context, *StreamRequest) (*StreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (*UnimplementedStream_ServiceServer) Pull(context.Context, *StreamRequest) (*StreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pull not implemented")
}
func (*UnimplementedStream_ServiceServer) Bidirection(context.Context, *StreamRequest) (*StreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bidirection not implemented")
}
func (*UnimplementedStream_ServiceServer) PushStream(Stream_Service_PushStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method PushStream not implemented")
}
func (*UnimplementedStream_ServiceServer) PullStream(Stream_Service_PullStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method PullStream not implemented")
}
func (*UnimplementedStream_ServiceServer) BidirectionStream(Stream_Service_BidirectionStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionStream not implemented")
}

func RegisterStream_ServiceServer(s *grpc.Server, srv Stream_ServiceServer) {
	s.RegisterService(&_Stream_Service_serviceDesc, srv)
}

func _Stream_Service_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Stream_ServiceServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.stream.Stream_Service/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Stream_ServiceServer).Push(ctx, req.(*StreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stream_Service_Pull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Stream_ServiceServer).Pull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.stream.Stream_Service/Pull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Stream_ServiceServer).Pull(ctx, req.(*StreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stream_Service_Bidirection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Stream_ServiceServer).Bidirection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.stream.Stream_Service/Bidirection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Stream_ServiceServer).Bidirection(ctx, req.(*StreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stream_Service_PushStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(Stream_ServiceServer).PushStream(&stream_ServicePushStreamServer{stream})
}

type Stream_Service_PushStreamServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type stream_ServicePushStreamServer struct {
	grpc.ServerStream
}

func (x *stream_ServicePushStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *stream_ServicePushStreamServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Stream_Service_PullStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(Stream_ServiceServer).PullStream(&stream_ServicePullStreamServer{stream})
}

type Stream_Service_PullStreamServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type stream_ServicePullStreamServer struct {
	grpc.ServerStream
}

func (x *stream_ServicePullStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *stream_ServicePullStreamServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Stream_Service_BidirectionStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(Stream_ServiceServer).BidirectionStream(&stream_ServiceBidirectionStreamServer{stream})
}

type Stream_Service_BidirectionStreamServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type stream_ServiceBidirectionStreamServer struct {
	grpc.ServerStream
}

func (x *stream_ServiceBidirectionStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *stream_ServiceBidirectionStreamServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Stream_Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "manager.stream.Stream_Service",
	HandlerType: (*Stream_ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _Stream_Service_Push_Handler,
		},
		{
			MethodName: "Pull",
			Handler:    _Stream_Service_Pull_Handler,
		},
		{
			MethodName: "Bidirection",
			Handler:    _Stream_Service_Bidirection_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PushStream",
			Handler:       _Stream_Service_PushStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "PullStream",
			Handler:       _Stream_Service_PullStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "BidirectionStream",
			Handler:       _Stream_Service_BidirectionStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/pb_stream/stream_manager.proto",
}
