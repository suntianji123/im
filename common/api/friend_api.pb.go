// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: friend_api.proto

package api

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

type FriendListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	MinFriendId int64 `protobuf:"varint,2,opt,name=minFriendId,proto3" json:"minFriendId,omitempty"`
	Size        int32 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *FriendListReq) Reset() {
	*x = FriendListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendListReq) ProtoMessage() {}

func (x *FriendListReq) ProtoReflect() protoreflect.Message {
	mi := &file_friend_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendListReq.ProtoReflect.Descriptor instead.
func (*FriendListReq) Descriptor() ([]byte, []int) {
	return file_friend_api_proto_rawDescGZIP(), []int{0}
}

func (x *FriendListReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *FriendListReq) GetMinFriendId() int64 {
	if x != nil {
		return x.MinFriendId
	}
	return 0
}

func (x *FriendListReq) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type FriendListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Friends []*UserInfo `protobuf:"bytes,1,rep,name=friends,proto3" json:"friends,omitempty"`
}

func (x *FriendListResp) Reset() {
	*x = FriendListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendListResp) ProtoMessage() {}

func (x *FriendListResp) ProtoReflect() protoreflect.Message {
	mi := &file_friend_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendListResp.ProtoReflect.Descriptor instead.
func (*FriendListResp) Descriptor() ([]byte, []int) {
	return file_friend_api_proto_rawDescGZIP(), []int{1}
}

func (x *FriendListResp) GetFriends() []*UserInfo {
	if x != nil {
		return x.Friends
	}
	return nil
}

var File_friend_api_proto protoreflect.FileDescriptor

var file_friend_api_proto_rawDesc = []byte{
	0x0a, 0x10, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x0d, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d,
	0x69, 0x6e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x39,
	0x0a, 0x0e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x27, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x32, 0x40, 0x0a, 0x0d, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e,
	0x2f, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_friend_api_proto_rawDescOnce sync.Once
	file_friend_api_proto_rawDescData = file_friend_api_proto_rawDesc
)

func file_friend_api_proto_rawDescGZIP() []byte {
	file_friend_api_proto_rawDescOnce.Do(func() {
		file_friend_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_friend_api_proto_rawDescData)
	})
	return file_friend_api_proto_rawDescData
}

var file_friend_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_friend_api_proto_goTypes = []interface{}{
	(*FriendListReq)(nil),  // 0: api.FriendListReq
	(*FriendListResp)(nil), // 1: api.FriendListResp
	(*UserInfo)(nil),       // 2: api.UserInfo
	(*Result)(nil),         // 3: api.Result
}
var file_friend_api_proto_depIdxs = []int32{
	2, // 0: api.FriendListResp.friends:type_name -> api.UserInfo
	0, // 1: api.FriendService.FriendList:input_type -> api.FriendListReq
	3, // 2: api.FriendService.FriendList:output_type -> api.Result
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_friend_api_proto_init() }
func file_friend_api_proto_init() {
	if File_friend_api_proto != nil {
		return
	}
	file_result_proto_init()
	file_user_api_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_friend_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendListReq); i {
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
		file_friend_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendListResp); i {
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
			RawDescriptor: file_friend_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_friend_api_proto_goTypes,
		DependencyIndexes: file_friend_api_proto_depIdxs,
		MessageInfos:      file_friend_api_proto_msgTypes,
	}.Build()
	File_friend_api_proto = out.File
	file_friend_api_proto_rawDesc = nil
	file_friend_api_proto_goTypes = nil
	file_friend_api_proto_depIdxs = nil
}
