// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: ping/ping.proto

package ping

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

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ping_ping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_ping_ping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_ping_ping_proto_rawDescGZIP(), []int{0}
}

func (x *Ping) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping *Ping `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ping_ping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ping_ping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_ping_ping_proto_rawDescGZIP(), []int{1}
}

func (x *PingRequest) GetPing() *Ping {
	if x != nil {
		return x.Ping
	}
	return nil
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping *Ping `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ping_ping_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ping_ping_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_ping_ping_proto_rawDescGZIP(), []int{2}
}

func (x *PingResponse) GetPing() *Ping {
	if x != nil {
		return x.Ping
	}
	return nil
}

var File_ping_ping_proto protoreflect.FileDescriptor

var file_ping_ping_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x1a, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x22, 0x2d, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x70, 0x69,
	0x6e, 0x67, 0x22, 0x2e, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x70, 0x69,
	0x6e, 0x67, 0x32, 0x39, 0x0a, 0x08, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x65, 0x73, 0x74, 0x12, 0x2d,
	0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x11, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ping_ping_proto_rawDescOnce sync.Once
	file_ping_ping_proto_rawDescData = file_ping_ping_proto_rawDesc
)

func file_ping_ping_proto_rawDescGZIP() []byte {
	file_ping_ping_proto_rawDescOnce.Do(func() {
		file_ping_ping_proto_rawDescData = protoimpl.X.CompressGZIP(file_ping_ping_proto_rawDescData)
	})
	return file_ping_ping_proto_rawDescData
}

var file_ping_ping_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ping_ping_proto_goTypes = []interface{}{
	(*Ping)(nil),         // 0: ping.Ping
	(*PingRequest)(nil),  // 1: ping.PingRequest
	(*PingResponse)(nil), // 2: ping.PingResponse
}
var file_ping_ping_proto_depIdxs = []int32{
	0, // 0: ping.PingRequest.ping:type_name -> ping.Ping
	0, // 1: ping.PingResponse.ping:type_name -> ping.Ping
	1, // 2: ping.PingTest.Ping:input_type -> ping.PingRequest
	2, // 3: ping.PingTest.Ping:output_type -> ping.PingResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ping_ping_proto_init() }
func file_ping_ping_proto_init() {
	if File_ping_ping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ping_ping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_ping_ping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_ping_ping_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
			RawDescriptor: file_ping_ping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ping_ping_proto_goTypes,
		DependencyIndexes: file_ping_ping_proto_depIdxs,
		MessageInfos:      file_ping_ping_proto_msgTypes,
	}.Build()
	File_ping_ping_proto = out.File
	file_ping_ping_proto_rawDesc = nil
	file_ping_ping_proto_goTypes = nil
	file_ping_ping_proto_depIdxs = nil
}
