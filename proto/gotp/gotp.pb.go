// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: gotp/gotp.proto

package gotp

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OTPEntry_OTPType int32

const (
	OTPEntry_TOTP OTPEntry_OTPType = 0
	OTPEntry_HOTP OTPEntry_OTPType = 1
)

// Enum value maps for OTPEntry_OTPType.
var (
	OTPEntry_OTPType_name = map[int32]string{
		0: "TOTP",
		1: "HOTP",
	}
	OTPEntry_OTPType_value = map[string]int32{
		"TOTP": 0,
		"HOTP": 1,
	}
)

func (x OTPEntry_OTPType) Enum() *OTPEntry_OTPType {
	p := new(OTPEntry_OTPType)
	*p = x
	return p
}

func (x OTPEntry_OTPType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OTPEntry_OTPType) Descriptor() protoreflect.EnumDescriptor {
	return file_gotp_gotp_proto_enumTypes[0].Descriptor()
}

func (OTPEntry_OTPType) Type() protoreflect.EnumType {
	return &file_gotp_gotp_proto_enumTypes[0]
}

func (x OTPEntry_OTPType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OTPEntry_OTPType.Descriptor instead.
func (OTPEntry_OTPType) EnumDescriptor() ([]byte, []int) {
	return file_gotp_gotp_proto_rawDescGZIP(), []int{1, 0}
}

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gotp_gotp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_gotp_gotp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_gotp_gotp_proto_rawDescGZIP(), []int{0}
}

func (x *UUID) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type OTPEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        OTPEntry_OTPType       `protobuf:"varint,1,opt,name=type,proto3,enum=gotp.OTPEntry_OTPType" json:"type,omitempty"`
	Uuid        string                 `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name        string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	SecretToken string                 `protobuf:"bytes,4,opt,name=secret_token,json=secretToken,proto3" json:"secret_token,omitempty"`
	Counter     uint64                 `protobuf:"varint,5,opt,name=counter,proto3" json:"counter,omitempty"`
	UpdateTime  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *OTPEntry) Reset() {
	*x = OTPEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gotp_gotp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OTPEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OTPEntry) ProtoMessage() {}

func (x *OTPEntry) ProtoReflect() protoreflect.Message {
	mi := &file_gotp_gotp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OTPEntry.ProtoReflect.Descriptor instead.
func (*OTPEntry) Descriptor() ([]byte, []int) {
	return file_gotp_gotp_proto_rawDescGZIP(), []int{1}
}

func (x *OTPEntry) GetType() OTPEntry_OTPType {
	if x != nil {
		return x.Type
	}
	return OTPEntry_TOTP
}

func (x *OTPEntry) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *OTPEntry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OTPEntry) GetSecretToken() string {
	if x != nil {
		return x.SecretToken
	}
	return ""
}

func (x *OTPEntry) GetCounter() uint64 {
	if x != nil {
		return x.Counter
	}
	return 0
}

func (x *OTPEntry) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_gotp_gotp_proto protoreflect.FileDescriptor

var file_gotp_gotp_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x6f, 0x74, 0x70, 0x2f, 0x67, 0x6f, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x67, 0x6f, 0x74, 0x70, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x22, 0xf7, 0x01, 0x0a, 0x08, 0x4f, 0x54, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x2a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x67, 0x6f, 0x74, 0x70, 0x2e, 0x4f, 0x54, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4f, 0x54,
	0x50, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x1d, 0x0a,
	0x07, 0x4f, 0x54, 0x50, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x4f, 0x54, 0x50,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x54, 0x50, 0x10, 0x01, 0x32, 0xc0, 0x02, 0x0a,
	0x04, 0x67, 0x4f, 0x54, 0x50, 0x12, 0x4c, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x0a, 0x2e, 0x67, 0x6f, 0x74, 0x70, 0x2e, 0x55, 0x55, 0x49, 0x44,
	0x1a, 0x0e, 0x2e, 0x67, 0x6f, 0x74, 0x70, 0x2e, 0x4f, 0x54, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f,
	0x74, 0x70, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x75, 0x75, 0x69, 0x64,
	0x7d, 0x30, 0x01, 0x12, 0x47, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x0e, 0x2e, 0x67, 0x6f, 0x74, 0x70, 0x2e, 0x4f, 0x54, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x1a,
	0x0e, 0x2e, 0x67, 0x6f, 0x74, 0x70, 0x2e, 0x4f, 0x54, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x74,
	0x70, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x51, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x2e, 0x67, 0x6f,
	0x74, 0x70, 0x2e, 0x4f, 0x54, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x0e, 0x2e, 0x67, 0x6f,
	0x74, 0x70, 0x2e, 0x4f, 0x54, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x74, 0x70, 0x2f, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x75, 0x75, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12,
	0x4e, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e,
	0x2e, 0x67, 0x6f, 0x74, 0x70, 0x2e, 0x4f, 0x54, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x0e,
	0x2e, 0x67, 0x6f, 0x74, 0x70, 0x2e, 0x4f, 0x54, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x2a, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x74, 0x70,
	0x2f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x75, 0x75, 0x69, 0x64, 0x7d, 0x42,
	0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61,
	0x64, 0x64, 0x6f, 0x73, 0x61, 0x75, 0x72, 0x75, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x70, 0x2f, 0x67,
	0x6f, 0x74, 0x70, 0x58, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gotp_gotp_proto_rawDescOnce sync.Once
	file_gotp_gotp_proto_rawDescData = file_gotp_gotp_proto_rawDesc
)

func file_gotp_gotp_proto_rawDescGZIP() []byte {
	file_gotp_gotp_proto_rawDescOnce.Do(func() {
		file_gotp_gotp_proto_rawDescData = protoimpl.X.CompressGZIP(file_gotp_gotp_proto_rawDescData)
	})
	return file_gotp_gotp_proto_rawDescData
}

var file_gotp_gotp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gotp_gotp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gotp_gotp_proto_goTypes = []interface{}{
	(OTPEntry_OTPType)(0),         // 0: gotp.OTPEntry.OTPType
	(*UUID)(nil),                  // 1: gotp.UUID
	(*OTPEntry)(nil),              // 2: gotp.OTPEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_gotp_gotp_proto_depIdxs = []int32{
	0, // 0: gotp.OTPEntry.type:type_name -> gotp.OTPEntry.OTPType
	3, // 1: gotp.OTPEntry.update_time:type_name -> google.protobuf.Timestamp
	1, // 2: gotp.gOTP.ListEntries:input_type -> gotp.UUID
	2, // 3: gotp.gOTP.AddEntry:input_type -> gotp.OTPEntry
	2, // 4: gotp.gOTP.UpdateEntry:input_type -> gotp.OTPEntry
	2, // 5: gotp.gOTP.DeleteEntry:input_type -> gotp.OTPEntry
	2, // 6: gotp.gOTP.ListEntries:output_type -> gotp.OTPEntry
	2, // 7: gotp.gOTP.AddEntry:output_type -> gotp.OTPEntry
	2, // 8: gotp.gOTP.UpdateEntry:output_type -> gotp.OTPEntry
	2, // 9: gotp.gOTP.DeleteEntry:output_type -> gotp.OTPEntry
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gotp_gotp_proto_init() }
func file_gotp_gotp_proto_init() {
	if File_gotp_gotp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gotp_gotp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
		file_gotp_gotp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OTPEntry); i {
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
			RawDescriptor: file_gotp_gotp_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gotp_gotp_proto_goTypes,
		DependencyIndexes: file_gotp_gotp_proto_depIdxs,
		EnumInfos:         file_gotp_gotp_proto_enumTypes,
		MessageInfos:      file_gotp_gotp_proto_msgTypes,
	}.Build()
	File_gotp_gotp_proto = out.File
	file_gotp_gotp_proto_rawDesc = nil
	file_gotp_gotp_proto_goTypes = nil
	file_gotp_gotp_proto_depIdxs = nil
}
