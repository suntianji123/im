// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: result.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *anypb.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Result) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Result) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type PPing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri      int32   `protobuf:"varint,1,opt,name=uri,proto3" json:"uri,omitempty"`
	AppId    int32   `protobuf:"varint,2,opt,name=appId,proto3" json:"appId,omitempty"`
	Uid      int64   `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Cts      int64   `protobuf:"varint,4,opt,name=cts,proto3" json:"cts,omitempty"`
	Channels []int32 `protobuf:"varint,5,rep,packed,name=channels,proto3" json:"channels,omitempty"`
	SyncPos  []int64 `protobuf:"varint,6,rep,packed,name=syncPos,proto3" json:"syncPos,omitempty"`
}

func (x *PPing) Reset() {
	*x = PPing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PPing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PPing) ProtoMessage() {}

func (x *PPing) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PPing.ProtoReflect.Descriptor instead.
func (*PPing) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{1}
}

func (x *PPing) GetUri() int32 {
	if x != nil {
		return x.Uri
	}
	return 0
}

func (x *PPing) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *PPing) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PPing) GetCts() int64 {
	if x != nil {
		return x.Cts
	}
	return 0
}

func (x *PPing) GetChannels() []int32 {
	if x != nil {
		return x.Channels
	}
	return nil
}

func (x *PPing) GetSyncPos() []int64 {
	if x != nil {
		return x.SyncPos
	}
	return nil
}

type PPong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri   int32 `protobuf:"varint,1,opt,name=uri,proto3" json:"uri,omitempty"`
	AppId int32 `protobuf:"varint,2,opt,name=appId,proto3" json:"appId,omitempty"`
	Uid   int64 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Cts   int64 `protobuf:"varint,4,opt,name=cts,proto3" json:"cts,omitempty"`
	Sts   int64 `protobuf:"varint,5,opt,name=sts,proto3" json:"sts,omitempty"`
}

func (x *PPong) Reset() {
	*x = PPong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PPong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PPong) ProtoMessage() {}

func (x *PPong) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PPong.ProtoReflect.Descriptor instead.
func (*PPong) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{2}
}

func (x *PPong) GetUri() int32 {
	if x != nil {
		return x.Uri
	}
	return 0
}

func (x *PPong) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *PPong) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PPong) GetCts() int64 {
	if x != nil {
		return x.Cts
	}
	return 0
}

func (x *PPong) GetSts() int64 {
	if x != nil {
		return x.Sts
	}
	return 0
}

var File_result_proto protoreflect.FileDescriptor

var file_result_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x61, 0x70, 0x69, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x28,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x89, 0x01, 0x0a, 0x05, 0x50, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x75, 0x72, 0x69, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x74, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x79,
	0x6e, 0x63, 0x50, 0x6f, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x73, 0x79, 0x6e,
	0x63, 0x50, 0x6f, 0x73, 0x22, 0x65, 0x0a, 0x05, 0x50, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x74, 0x73, 0x32, 0x3b, 0x0a, 0x11, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x72, 0x74, 0x12, 0x0a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x50, 0x50, 0x6f, 0x6e, 0x67, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_result_proto_rawDescOnce sync.Once
	file_result_proto_rawDescData = file_result_proto_rawDesc
)

func file_result_proto_rawDescGZIP() []byte {
	file_result_proto_rawDescOnce.Do(func() {
		file_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_result_proto_rawDescData)
	})
	return file_result_proto_rawDescData
}

var file_result_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_result_proto_goTypes = []interface{}{
	(*Result)(nil),    // 0: api.Result
	(*PPing)(nil),     // 1: api.PPing
	(*PPong)(nil),     // 2: api.PPong
	(*anypb.Any)(nil), // 3: google.protobuf.Any
}
var file_result_proto_depIdxs = []int32{
	3, // 0: api.Result.data:type_name -> google.protobuf.Any
	1, // 1: api.HeartBeartService.HeartBeart:input_type -> api.PPing
	2, // 2: api.HeartBeartService.HeartBeart:output_type -> api.PPong
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_result_proto_init() }
func file_result_proto_init() {
	if File_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PPing); i {
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
		file_result_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PPong); i {
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
			RawDescriptor: file_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_result_proto_goTypes,
		DependencyIndexes: file_result_proto_depIdxs,
		MessageInfos:      file_result_proto_msgTypes,
	}.Build()
	File_result_proto = out.File
	file_result_proto_rawDesc = nil
	file_result_proto_goTypes = nil
	file_result_proto_depIdxs = nil
}
