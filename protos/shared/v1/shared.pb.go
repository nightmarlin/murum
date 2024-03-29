// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: shared/v1/shared.proto

package v1

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

type FileType int32

const (
	FileType_PNG_FILE_TYPE FileType = 0
	FileType_JPG_FILE_TYPE FileType = 1
	FileType_GIF_FILE_TYPE FileType = 2
)

// Enum value maps for FileType.
var (
	FileType_name = map[int32]string{
		0: "PNG_FILE_TYPE",
		1: "JPG_FILE_TYPE",
		2: "GIF_FILE_TYPE",
	}
	FileType_value = map[string]int32{
		"PNG_FILE_TYPE": 0,
		"JPG_FILE_TYPE": 1,
		"GIF_FILE_TYPE": 2,
	}
)

func (x FileType) Enum() *FileType {
	p := new(FileType)
	*p = x
	return p
}

func (x FileType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileType) Descriptor() protoreflect.EnumDescriptor {
	return file_shared_v1_shared_proto_enumTypes[0].Descriptor()
}

func (FileType) Type() protoreflect.EnumType {
	return &file_shared_v1_shared_proto_enumTypes[0]
}

func (x FileType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileType.Descriptor instead.
func (FileType) EnumDescriptor() ([]byte, []int) {
	return file_shared_v1_shared_proto_rawDescGZIP(), []int{0}
}

// A TimeRange represents a period of time as follows: after <= T < before.
type TimeRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	After  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=after,proto3" json:"after,omitempty"`   // Inclusive
	Before *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=before,proto3" json:"before,omitempty"` // Exclusive
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_v1_shared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_shared_v1_shared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRange.ProtoReflect.Descriptor instead.
func (*TimeRange) Descriptor() ([]byte, []int) {
	return file_shared_v1_shared_proto_rawDescGZIP(), []int{0}
}

func (x *TimeRange) GetAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.After
	}
	return nil
}

func (x *TimeRange) GetBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.Before
	}
	return nil
}

// A Point represents an arbitrary point on a 2D plane.
type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_v1_shared_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_shared_v1_shared_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_shared_v1_shared_proto_rawDescGZIP(), []int{1}
}

func (x *Point) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Point) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

// A Rect describes the bounding box formed between a pair of Point s.
//
//	Rect.min.z <= X < Rect.max.x
//	Rect.min.y <= Y < Rect.max.y
type Rect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min *Point `protobuf:"bytes,1,opt,name=min,proto3" json:"min,omitempty"` // Inclusive
	Max *Point `protobuf:"bytes,2,opt,name=max,proto3" json:"max,omitempty"` // Exclusive
}

func (x *Rect) Reset() {
	*x = Rect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_v1_shared_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rect) ProtoMessage() {}

func (x *Rect) ProtoReflect() protoreflect.Message {
	mi := &file_shared_v1_shared_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rect.ProtoReflect.Descriptor instead.
func (*Rect) Descriptor() ([]byte, []int) {
	return file_shared_v1_shared_proto_rawDescGZIP(), []int{2}
}

func (x *Rect) GetMin() *Point {
	if x != nil {
		return x.Min
	}
	return nil
}

func (x *Rect) GetMax() *Point {
	if x != nil {
		return x.Max
	}
	return nil
}

// An Image contains the raw data to display a graphic.
type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	FileType FileType `protobuf:"varint,2,opt,name=file_type,json=fileType,proto3,enum=shared.FileType" json:"file_type,omitempty"`
	Image    []byte   `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_v1_shared_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_shared_v1_shared_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_shared_v1_shared_proto_rawDescGZIP(), []int{3}
}

func (x *Image) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Image) GetFileType() FileType {
	if x != nil {
		return x.FileType
	}
	return FileType_PNG_FILE_TYPE
}

func (x *Image) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_shared_v1_shared_proto protoreflect.FileDescriptor

var file_shared_v1_shared_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x71, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x30,
	0x0a, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x12, 0x32, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x62, 0x65,
	0x66, 0x6f, 0x72, 0x65, 0x22, 0x23, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22, 0x48, 0x0a, 0x04, 0x52, 0x65, 0x63,
	0x74, 0x12, 0x1f, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x03, 0x6d,
	0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x03,
	0x6d, 0x61, 0x78, 0x22, 0x60, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2d, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x2a, 0x43, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x4e, 0x47, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4a, 0x50, 0x47, 0x5f, 0x46, 0x49, 0x4c, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x49, 0x46, 0x5f, 0x46,
	0x49, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x02, 0x42, 0x80, 0x01, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x42, 0x0b, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x6d, 0x61, 0x72, 0x6c, 0x69, 0x6e,
	0x2f, 0x6d, 0x75, 0x72, 0x75, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x06,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0xca, 0x02, 0x06, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0xe2,
	0x02, 0x12, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shared_v1_shared_proto_rawDescOnce sync.Once
	file_shared_v1_shared_proto_rawDescData = file_shared_v1_shared_proto_rawDesc
)

func file_shared_v1_shared_proto_rawDescGZIP() []byte {
	file_shared_v1_shared_proto_rawDescOnce.Do(func() {
		file_shared_v1_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_v1_shared_proto_rawDescData)
	})
	return file_shared_v1_shared_proto_rawDescData
}

var file_shared_v1_shared_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_shared_v1_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_shared_v1_shared_proto_goTypes = []interface{}{
	(FileType)(0),                 // 0: shared.FileType
	(*TimeRange)(nil),             // 1: shared.TimeRange
	(*Point)(nil),                 // 2: shared.Point
	(*Rect)(nil),                  // 3: shared.Rect
	(*Image)(nil),                 // 4: shared.Image
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_shared_v1_shared_proto_depIdxs = []int32{
	5, // 0: shared.TimeRange.after:type_name -> google.protobuf.Timestamp
	5, // 1: shared.TimeRange.before:type_name -> google.protobuf.Timestamp
	2, // 2: shared.Rect.min:type_name -> shared.Point
	2, // 3: shared.Rect.max:type_name -> shared.Point
	0, // 4: shared.Image.file_type:type_name -> shared.FileType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_shared_v1_shared_proto_init() }
func file_shared_v1_shared_proto_init() {
	if File_shared_v1_shared_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shared_v1_shared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeRange); i {
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
		file_shared_v1_shared_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_shared_v1_shared_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rect); i {
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
		file_shared_v1_shared_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
			RawDescriptor: file_shared_v1_shared_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shared_v1_shared_proto_goTypes,
		DependencyIndexes: file_shared_v1_shared_proto_depIdxs,
		EnumInfos:         file_shared_v1_shared_proto_enumTypes,
		MessageInfos:      file_shared_v1_shared_proto_msgTypes,
	}.Build()
	File_shared_v1_shared_proto = out.File
	file_shared_v1_shared_proto_rawDesc = nil
	file_shared_v1_shared_proto_goTypes = nil
	file_shared_v1_shared_proto_depIdxs = nil
}
