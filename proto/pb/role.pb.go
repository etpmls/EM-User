// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: role.proto

package pb

import (
	empb "github.com/Etpmls/Etpmls-Micro/v3/proto/empb"
	proto "github.com/golang/protobuf/proto"
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

type RoleCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Remark      string     `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	Permissions []*empb.Id `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *RoleCreate) Reset() {
	*x = RoleCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleCreate) ProtoMessage() {}

func (x *RoleCreate) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleCreate.ProtoReflect.Descriptor instead.
func (*RoleCreate) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{0}
}

func (x *RoleCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleCreate) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *RoleCreate) GetPermissions() []*empb.Id {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type RoleEdit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark      string     `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Permissions []*empb.Id `protobuf:"bytes,4,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *RoleEdit) Reset() {
	*x = RoleEdit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleEdit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleEdit) ProtoMessage() {}

func (x *RoleEdit) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleEdit.ProtoReflect.Descriptor instead.
func (*RoleEdit) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{1}
}

func (x *RoleEdit) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RoleEdit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleEdit) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *RoleEdit) GetPermissions() []*empb.Id {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type RoleDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roles []*empb.Id `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *RoleDelete) Reset() {
	*x = RoleDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleDelete) ProtoMessage() {}

func (x *RoleDelete) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleDelete.ProtoReflect.Descriptor instead.
func (*RoleDelete) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{2}
}

func (x *RoleDelete) GetRoles() []*empb.Id {
	if x != nil {
		return x.Roles
	}
	return nil
}

var File_role_proto protoreflect.FileDescriptor

var file_role_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x1a, 0x10, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x2f, 0x65, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x0a, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x28,
	0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x65, 0x6d, 0x2e, 0x49, 0x64, 0x52, 0x0b, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x70, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65,
	0x45, 0x64, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x12, 0x28, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x65, 0x6d, 0x2e, 0x49, 0x64, 0x52, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2a, 0x0a, 0x0a, 0x52, 0x6f,
	0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x65, 0x6d, 0x2e, 0x49, 0x64, 0x52,
	0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x32, 0xb0, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x28, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x0e, 0x2e, 0x65, 0x6d, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0c, 0x2e, 0x65, 0x6d, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x0c, 0x2e, 0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x04, 0x45, 0x64, 0x69, 0x74, 0x12, 0x0e, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x45, 0x64, 0x69, 0x74, 0x1a, 0x0c, 0x2e,
	0x65, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2a, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x0c, 0x2e, 0x65, 0x6d, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_proto_rawDescOnce sync.Once
	file_role_proto_rawDescData = file_role_proto_rawDesc
)

func file_role_proto_rawDescGZIP() []byte {
	file_role_proto_rawDescOnce.Do(func() {
		file_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_proto_rawDescData)
	})
	return file_role_proto_rawDescData
}

var file_role_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_role_proto_goTypes = []interface{}{
	(*RoleCreate)(nil),      // 0: role.RoleCreate
	(*RoleEdit)(nil),        // 1: role.RoleEdit
	(*RoleDelete)(nil),      // 2: role.RoleDelete
	(*empb.Id)(nil),         // 3: em.Id
	(*empb.Pagination)(nil), // 4: em.Pagination
	(*empb.Response)(nil),   // 5: em.Response
}
var file_role_proto_depIdxs = []int32{
	3, // 0: role.RoleCreate.permissions:type_name -> em.Id
	3, // 1: role.RoleEdit.permissions:type_name -> em.Id
	3, // 2: role.RoleDelete.roles:type_name -> em.Id
	4, // 3: role.Role.GetAll:input_type -> em.Pagination
	0, // 4: role.Role.Create:input_type -> role.RoleCreate
	1, // 5: role.Role.Edit:input_type -> role.RoleEdit
	2, // 6: role.Role.Delete:input_type -> role.RoleDelete
	5, // 7: role.Role.GetAll:output_type -> em.Response
	5, // 8: role.Role.Create:output_type -> em.Response
	5, // 9: role.Role.Edit:output_type -> em.Response
	5, // 10: role.Role.Delete:output_type -> em.Response
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_role_proto_init() }
func file_role_proto_init() {
	if File_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleCreate); i {
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
		file_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleEdit); i {
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
		file_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleDelete); i {
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
			RawDescriptor: file_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_role_proto_goTypes,
		DependencyIndexes: file_role_proto_depIdxs,
		MessageInfos:      file_role_proto_msgTypes,
	}.Build()
	File_role_proto = out.File
	file_role_proto_rawDesc = nil
	file_role_proto_goTypes = nil
	file_role_proto_depIdxs = nil
}
