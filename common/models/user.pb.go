// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: common/models/proto/user.proto

package models

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_models_proto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_common_models_proto_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_common_models_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UserLists struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lists []*User `protobuf:"bytes,1,rep,name=lists,proto3" json:"lists,omitempty"`
}

func (x *UserLists) Reset() {
	*x = UserLists{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_models_proto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLists) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLists) ProtoMessage() {}

func (x *UserLists) ProtoReflect() protoreflect.Message {
	mi := &file_common_models_proto_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLists.ProtoReflect.Descriptor instead.
func (*UserLists) Descriptor() ([]byte, []int) {
	return file_common_models_proto_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserLists) GetLists() []*User {
	if x != nil {
		return x.Lists
	}
	return nil
}

var File_common_models_proto_user_proto protoreflect.FileDescriptor

var file_common_models_proto_user_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x2f, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x22,
	0x0a, 0x05, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x6c, 0x69, 0x73,
	0x74, 0x73, 0x32, 0x67, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x36, 0x0a, 0x07, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x73, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0c, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x65, 0x78, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_models_proto_user_proto_rawDescOnce sync.Once
	file_common_models_proto_user_proto_rawDescData = file_common_models_proto_user_proto_rawDesc
)

func file_common_models_proto_user_proto_rawDescGZIP() []byte {
	file_common_models_proto_user_proto_rawDescOnce.Do(func() {
		file_common_models_proto_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_models_proto_user_proto_rawDescData)
	})
	return file_common_models_proto_user_proto_rawDescData
}

var file_common_models_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_models_proto_user_proto_goTypes = []interface{}{
	(*User)(nil),        // 0: models.User
	(*UserLists)(nil),   // 1: models.UserLists
	(*empty.Empty)(nil), // 2: google.protobuf.Empty
}
var file_common_models_proto_user_proto_depIdxs = []int32{
	0, // 0: models.UserLists.lists:type_name -> models.User
	2, // 1: models.Users.FindAll:input_type -> google.protobuf.Empty
	0, // 2: models.Users.Create:input_type -> models.User
	1, // 3: models.Users.FindAll:output_type -> models.UserLists
	0, // 4: models.Users.Create:output_type -> models.User
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_models_proto_user_proto_init() }
func file_common_models_proto_user_proto_init() {
	if File_common_models_proto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_models_proto_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_common_models_proto_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLists); i {
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
			RawDescriptor: file_common_models_proto_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_models_proto_user_proto_goTypes,
		DependencyIndexes: file_common_models_proto_user_proto_depIdxs,
		MessageInfos:      file_common_models_proto_user_proto_msgTypes,
	}.Build()
	File_common_models_proto_user_proto = out.File
	file_common_models_proto_user_proto_rawDesc = nil
	file_common_models_proto_user_proto_goTypes = nil
	file_common_models_proto_user_proto_depIdxs = nil
}
