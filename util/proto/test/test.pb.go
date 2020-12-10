// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: util/proto/test/test.proto

package test

import (
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

type Primitives_SomeEnum int32

const (
	Primitives_UNKNOWN  Primitives_SomeEnum = 0
	Primitives_POSITIVE Primitives_SomeEnum = 1
	Primitives_NEGATIVE Primitives_SomeEnum = -1
)

// Enum value maps for Primitives_SomeEnum.
var (
	Primitives_SomeEnum_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "POSITIVE",
		-1: "NEGATIVE",
	}
	Primitives_SomeEnum_value = map[string]int32{
		"UNKNOWN":  0,
		"POSITIVE": 1,
		"NEGATIVE": -1,
	}
)

func (x Primitives_SomeEnum) Enum() *Primitives_SomeEnum {
	p := new(Primitives_SomeEnum)
	*p = x
	return p
}

func (x Primitives_SomeEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Primitives_SomeEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_util_proto_test_test_proto_enumTypes[0].Descriptor()
}

func (Primitives_SomeEnum) Type() protoreflect.EnumType {
	return &file_util_proto_test_test_proto_enumTypes[0]
}

func (x Primitives_SomeEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Primitives_SomeEnum.Descriptor instead.
func (Primitives_SomeEnum) EnumDescriptor() ([]byte, []int) {
	return file_util_proto_test_test_proto_rawDescGZIP(), []int{0, 0}
}

type Primitives struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldA []byte              `protobuf:"bytes,1,opt,name=field_a,json=fieldA,proto3" json:"field_a,omitempty"`
	FieldB string              `protobuf:"bytes,2,opt,name=field_b,json=fieldB,proto3" json:"field_b,omitempty"`
	FieldC bool                `protobuf:"varint,200,opt,name=field_c,json=fieldC,proto3" json:"field_c,omitempty"`
	FieldD int32               `protobuf:"varint,201,opt,name=field_d,json=fieldD,proto3" json:"field_d,omitempty"`
	FieldE uint32              `protobuf:"varint,202,opt,name=field_e,json=fieldE,proto3" json:"field_e,omitempty"`
	FieldF int64               `protobuf:"varint,203,opt,name=field_f,json=fieldF,proto3" json:"field_f,omitempty"`
	FieldG uint64              `protobuf:"varint,204,opt,name=field_g,json=fieldG,proto3" json:"field_g,omitempty"`
	FieldH Primitives_SomeEnum `protobuf:"varint,300,opt,name=field_h,json=fieldH,proto3,enum=test.Primitives_SomeEnum" json:"field_h,omitempty"`
}

func (x *Primitives) Reset() {
	*x = Primitives{}
	if protoimpl.UnsafeEnabled {
		mi := &file_util_proto_test_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Primitives) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Primitives) ProtoMessage() {}

func (x *Primitives) ProtoReflect() protoreflect.Message {
	mi := &file_util_proto_test_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Primitives.ProtoReflect.Descriptor instead.
func (*Primitives) Descriptor() ([]byte, []int) {
	return file_util_proto_test_test_proto_rawDescGZIP(), []int{0}
}

func (x *Primitives) GetFieldA() []byte {
	if x != nil {
		return x.FieldA
	}
	return nil
}

func (x *Primitives) GetFieldB() string {
	if x != nil {
		return x.FieldB
	}
	return ""
}

func (x *Primitives) GetFieldC() bool {
	if x != nil {
		return x.FieldC
	}
	return false
}

func (x *Primitives) GetFieldD() int32 {
	if x != nil {
		return x.FieldD
	}
	return 0
}

func (x *Primitives) GetFieldE() uint32 {
	if x != nil {
		return x.FieldE
	}
	return 0
}

func (x *Primitives) GetFieldF() int64 {
	if x != nil {
		return x.FieldF
	}
	return 0
}

func (x *Primitives) GetFieldG() uint64 {
	if x != nil {
		return x.FieldG
	}
	return 0
}

func (x *Primitives) GetFieldH() Primitives_SomeEnum {
	if x != nil {
		return x.FieldH
	}
	return Primitives_UNKNOWN
}

type RepPrimitives struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldA [][]byte `protobuf:"bytes,1,rep,name=field_a,json=fieldA,proto3" json:"field_a,omitempty"`
	FieldB []string `protobuf:"bytes,2,rep,name=field_b,json=fieldB,proto3" json:"field_b,omitempty"`
	FieldC []int32  `protobuf:"varint,3,rep,packed,name=field_c,json=fieldC,proto3" json:"field_c,omitempty"`
	FieldD []uint32 `protobuf:"varint,4,rep,packed,name=field_d,json=fieldD,proto3" json:"field_d,omitempty"`
	FieldE []int64  `protobuf:"varint,5,rep,packed,name=field_e,json=fieldE,proto3" json:"field_e,omitempty"`
	FieldF []uint64 `protobuf:"varint,6,rep,packed,name=field_f,json=fieldF,proto3" json:"field_f,omitempty"`
}

func (x *RepPrimitives) Reset() {
	*x = RepPrimitives{}
	if protoimpl.UnsafeEnabled {
		mi := &file_util_proto_test_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepPrimitives) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepPrimitives) ProtoMessage() {}

func (x *RepPrimitives) ProtoReflect() protoreflect.Message {
	mi := &file_util_proto_test_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepPrimitives.ProtoReflect.Descriptor instead.
func (*RepPrimitives) Descriptor() ([]byte, []int) {
	return file_util_proto_test_test_proto_rawDescGZIP(), []int{1}
}

func (x *RepPrimitives) GetFieldA() [][]byte {
	if x != nil {
		return x.FieldA
	}
	return nil
}

func (x *RepPrimitives) GetFieldB() []string {
	if x != nil {
		return x.FieldB
	}
	return nil
}

func (x *RepPrimitives) GetFieldC() []int32 {
	if x != nil {
		return x.FieldC
	}
	return nil
}

func (x *RepPrimitives) GetFieldD() []uint32 {
	if x != nil {
		return x.FieldD
	}
	return nil
}

func (x *RepPrimitives) GetFieldE() []int64 {
	if x != nil {
		return x.FieldE
	}
	return nil
}

func (x *RepPrimitives) GetFieldF() []uint64 {
	if x != nil {
		return x.FieldF
	}
	return nil
}

var File_util_proto_test_test_proto protoreflect.FileDescriptor

var file_util_proto_test_test_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65,
	0x73, 0x74, 0x22, 0xb3, 0x02, 0x0a, 0x0a, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x73, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x41, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x42, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x63, 0x18, 0xc8,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x12, 0x18, 0x0a,
	0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x64, 0x18, 0xc9, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x65, 0x18, 0xca, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x45, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x66, 0x18, 0xcb, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x12, 0x18, 0x0a, 0x07, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x67, 0x18, 0xcc, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x47, 0x12, 0x33, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x68,
	0x18, 0xac, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x50,
	0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x73, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x48, 0x22, 0x3c, 0x0a, 0x08, 0x53, 0x6f,
	0x6d, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x56, 0x45, 0x10,
	0x01, 0x12, 0x15, 0x0a, 0x08, 0x4e, 0x45, 0x47, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x22, 0xa5, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x70,
	0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x41, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x12, 0x17, 0x0a, 0x07,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x63, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x43, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x64,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x12, 0x17,
	0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x66, 0x18, 0x06, 0x20, 0x03, 0x28, 0x04, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x46,
	0x42, 0x11, 0x5a, 0x0f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_util_proto_test_test_proto_rawDescOnce sync.Once
	file_util_proto_test_test_proto_rawDescData = file_util_proto_test_test_proto_rawDesc
)

func file_util_proto_test_test_proto_rawDescGZIP() []byte {
	file_util_proto_test_test_proto_rawDescOnce.Do(func() {
		file_util_proto_test_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_util_proto_test_test_proto_rawDescData)
	})
	return file_util_proto_test_test_proto_rawDescData
}

var file_util_proto_test_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_util_proto_test_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_util_proto_test_test_proto_goTypes = []interface{}{
	(Primitives_SomeEnum)(0), // 0: test.Primitives.SomeEnum
	(*Primitives)(nil),       // 1: test.Primitives
	(*RepPrimitives)(nil),    // 2: test.RepPrimitives
}
var file_util_proto_test_test_proto_depIdxs = []int32{
	0, // 0: test.Primitives.field_h:type_name -> test.Primitives.SomeEnum
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_util_proto_test_test_proto_init() }
func file_util_proto_test_test_proto_init() {
	if File_util_proto_test_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_util_proto_test_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Primitives); i {
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
		file_util_proto_test_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepPrimitives); i {
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
			RawDescriptor: file_util_proto_test_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_util_proto_test_test_proto_goTypes,
		DependencyIndexes: file_util_proto_test_test_proto_depIdxs,
		EnumInfos:         file_util_proto_test_test_proto_enumTypes,
		MessageInfos:      file_util_proto_test_test_proto_msgTypes,
	}.Build()
	File_util_proto_test_test_proto = out.File
	file_util_proto_test_test_proto_rawDesc = nil
	file_util_proto_test_test_proto_goTypes = nil
	file_util_proto_test_test_proto_depIdxs = nil
}
