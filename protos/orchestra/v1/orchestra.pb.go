// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: orchestra/v1/orchestra.proto

package v1

import (
	v12 "github.com/nightmarlin/murum/protos/layout/v1"
	v13 "github.com/nightmarlin/murum/protos/render/v1"
	v1 "github.com/nightmarlin/murum/protos/shared/v1"
	v11 "github.com/nightmarlin/murum/protos/source/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Source int32

const (
	Source_LASTFM_SOURCE Source = 0
)

// Enum value maps for Source.
var (
	Source_name = map[int32]string{
		0: "LASTFM_SOURCE",
	}
	Source_value = map[string]int32{
		"LASTFM_SOURCE": 0,
	}
)

func (x Source) Enum() *Source {
	p := new(Source)
	*p = x
	return p
}

func (x Source) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Source) Descriptor() protoreflect.EnumDescriptor {
	return file_orchestra_v1_orchestra_proto_enumTypes[0].Descriptor()
}

func (Source) Type() protoreflect.EnumType {
	return &file_orchestra_v1_orchestra_proto_enumTypes[0]
}

func (x Source) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Source.Descriptor instead.
func (Source) EnumDescriptor() ([]byte, []int) {
	return file_orchestra_v1_orchestra_proto_rawDescGZIP(), []int{0}
}

type GenerateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bounds      *v1.Rect `protobuf:"bytes,1,opt,name=bounds,proto3" json:"bounds,omitempty"`
	RowCount    int32    `protobuf:"varint,2,opt,name=row_count,json=rowCount,proto3" json:"row_count,omitempty"`
	ColumnCount int32    `protobuf:"varint,3,opt,name=column_count,json=columnCount,proto3" json:"column_count,omitempty"`
	Source      Source   `protobuf:"varint,4,opt,name=source,proto3,enum=orchestra.Source" json:"source,omitempty"`
	// Types that are assignable to Media:
	//
	//	*GenerateRequest_Tracks
	//	*GenerateRequest_Albums
	//	*GenerateRequest_Artists
	Media isGenerateRequest_Media `protobuf_oneof:"media"`
	// Types that are assignable to LayoutMode:
	//
	//	*GenerateRequest_Random
	//	*GenerateRequest_Even
	//	*GenerateRequest_Varied
	//	*GenerateRequest_Triangular
	//	*GenerateRequest_Weighted
	LayoutMode isGenerateRequest_LayoutMode `protobuf_oneof:"layout_mode"`
	FileType   v1.FileType                  `protobuf:"varint,13,opt,name=file_type,json=fileType,proto3,enum=shared.FileType" json:"file_type,omitempty"`
	// Types that are assignable to RenderMode:
	//
	//	*GenerateRequest_Basic
	//	*GenerateRequest_WithText
	RenderMode isGenerateRequest_RenderMode `protobuf_oneof:"render_mode"`
}

func (x *GenerateRequest) Reset() {
	*x = GenerateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestra_v1_orchestra_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateRequest) ProtoMessage() {}

func (x *GenerateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestra_v1_orchestra_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateRequest.ProtoReflect.Descriptor instead.
func (*GenerateRequest) Descriptor() ([]byte, []int) {
	return file_orchestra_v1_orchestra_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateRequest) GetBounds() *v1.Rect {
	if x != nil {
		return x.Bounds
	}
	return nil
}

func (x *GenerateRequest) GetRowCount() int32 {
	if x != nil {
		return x.RowCount
	}
	return 0
}

func (x *GenerateRequest) GetColumnCount() int32 {
	if x != nil {
		return x.ColumnCount
	}
	return 0
}

func (x *GenerateRequest) GetSource() Source {
	if x != nil {
		return x.Source
	}
	return Source_LASTFM_SOURCE
}

func (m *GenerateRequest) GetMedia() isGenerateRequest_Media {
	if m != nil {
		return m.Media
	}
	return nil
}

func (x *GenerateRequest) GetTracks() *v11.TracksRequest {
	if x, ok := x.GetMedia().(*GenerateRequest_Tracks); ok {
		return x.Tracks
	}
	return nil
}

func (x *GenerateRequest) GetAlbums() *v11.AlbumsRequest {
	if x, ok := x.GetMedia().(*GenerateRequest_Albums); ok {
		return x.Albums
	}
	return nil
}

func (x *GenerateRequest) GetArtists() *v11.ArtistsRequest {
	if x, ok := x.GetMedia().(*GenerateRequest_Artists); ok {
		return x.Artists
	}
	return nil
}

func (m *GenerateRequest) GetLayoutMode() isGenerateRequest_LayoutMode {
	if m != nil {
		return m.LayoutMode
	}
	return nil
}

func (x *GenerateRequest) GetRandom() *v12.RandomParams {
	if x, ok := x.GetLayoutMode().(*GenerateRequest_Random); ok {
		return x.Random
	}
	return nil
}

func (x *GenerateRequest) GetEven() *v12.EvenParams {
	if x, ok := x.GetLayoutMode().(*GenerateRequest_Even); ok {
		return x.Even
	}
	return nil
}

func (x *GenerateRequest) GetVaried() *v12.VariedParams {
	if x, ok := x.GetLayoutMode().(*GenerateRequest_Varied); ok {
		return x.Varied
	}
	return nil
}

func (x *GenerateRequest) GetTriangular() *v12.TriangularParams {
	if x, ok := x.GetLayoutMode().(*GenerateRequest_Triangular); ok {
		return x.Triangular
	}
	return nil
}

func (x *GenerateRequest) GetWeighted() *v12.WeightedParams {
	if x, ok := x.GetLayoutMode().(*GenerateRequest_Weighted); ok {
		return x.Weighted
	}
	return nil
}

func (x *GenerateRequest) GetFileType() v1.FileType {
	if x != nil {
		return x.FileType
	}
	return v1.FileType(0)
}

func (m *GenerateRequest) GetRenderMode() isGenerateRequest_RenderMode {
	if m != nil {
		return m.RenderMode
	}
	return nil
}

func (x *GenerateRequest) GetBasic() *v13.BasicParams {
	if x, ok := x.GetRenderMode().(*GenerateRequest_Basic); ok {
		return x.Basic
	}
	return nil
}

func (x *GenerateRequest) GetWithText() *v13.WithTextParams {
	if x, ok := x.GetRenderMode().(*GenerateRequest_WithText); ok {
		return x.WithText
	}
	return nil
}

type isGenerateRequest_Media interface {
	isGenerateRequest_Media()
}

type GenerateRequest_Tracks struct {
	Tracks *v11.TracksRequest `protobuf:"bytes,5,opt,name=tracks,proto3,oneof"`
}

type GenerateRequest_Albums struct {
	Albums *v11.AlbumsRequest `protobuf:"bytes,6,opt,name=albums,proto3,oneof"`
}

type GenerateRequest_Artists struct {
	Artists *v11.ArtistsRequest `protobuf:"bytes,7,opt,name=artists,proto3,oneof"`
}

func (*GenerateRequest_Tracks) isGenerateRequest_Media() {}

func (*GenerateRequest_Albums) isGenerateRequest_Media() {}

func (*GenerateRequest_Artists) isGenerateRequest_Media() {}

type isGenerateRequest_LayoutMode interface {
	isGenerateRequest_LayoutMode()
}

type GenerateRequest_Random struct {
	Random *v12.RandomParams `protobuf:"bytes,8,opt,name=random,proto3,oneof"`
}

type GenerateRequest_Even struct {
	Even *v12.EvenParams `protobuf:"bytes,9,opt,name=even,proto3,oneof"`
}

type GenerateRequest_Varied struct {
	Varied *v12.VariedParams `protobuf:"bytes,10,opt,name=varied,proto3,oneof"`
}

type GenerateRequest_Triangular struct {
	Triangular *v12.TriangularParams `protobuf:"bytes,11,opt,name=triangular,proto3,oneof"`
}

type GenerateRequest_Weighted struct {
	Weighted *v12.WeightedParams `protobuf:"bytes,12,opt,name=weighted,proto3,oneof"`
}

func (*GenerateRequest_Random) isGenerateRequest_LayoutMode() {}

func (*GenerateRequest_Even) isGenerateRequest_LayoutMode() {}

func (*GenerateRequest_Varied) isGenerateRequest_LayoutMode() {}

func (*GenerateRequest_Triangular) isGenerateRequest_LayoutMode() {}

func (*GenerateRequest_Weighted) isGenerateRequest_LayoutMode() {}

type isGenerateRequest_RenderMode interface {
	isGenerateRequest_RenderMode()
}

type GenerateRequest_Basic struct {
	Basic *v13.BasicParams `protobuf:"bytes,14,opt,name=basic,proto3,oneof"`
}

type GenerateRequest_WithText struct {
	WithText *v13.WithTextParams `protobuf:"bytes,15,opt,name=with_text,json=withText,proto3,oneof"`
}

func (*GenerateRequest_Basic) isGenerateRequest_RenderMode() {}

func (*GenerateRequest_WithText) isGenerateRequest_RenderMode() {}

type GenerateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *v1.Image `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GenerateResponse) Reset() {
	*x = GenerateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestra_v1_orchestra_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateResponse) ProtoMessage() {}

func (x *GenerateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestra_v1_orchestra_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateResponse.ProtoReflect.Descriptor instead.
func (*GenerateResponse) Descriptor() ([]byte, []int) {
	return file_orchestra_v1_orchestra_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateResponse) GetImage() *v1.Image {
	if x != nil {
		return x.Image
	}
	return nil
}

type OAuthGetTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source  Source               `protobuf:"varint,1,opt,name=source,proto3,enum=orchestra.Source" json:"source,omitempty"`
	Request *v11.GetTokenRequest `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *OAuthGetTokenRequest) Reset() {
	*x = OAuthGetTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestra_v1_orchestra_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthGetTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthGetTokenRequest) ProtoMessage() {}

func (x *OAuthGetTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestra_v1_orchestra_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthGetTokenRequest.ProtoReflect.Descriptor instead.
func (*OAuthGetTokenRequest) Descriptor() ([]byte, []int) {
	return file_orchestra_v1_orchestra_proto_rawDescGZIP(), []int{2}
}

func (x *OAuthGetTokenRequest) GetSource() Source {
	if x != nil {
		return x.Source
	}
	return Source_LASTFM_SOURCE
}

func (x *OAuthGetTokenRequest) GetRequest() *v11.GetTokenRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type OAuthGetTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *v11.GetTokenResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *OAuthGetTokenResponse) Reset() {
	*x = OAuthGetTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestra_v1_orchestra_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthGetTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthGetTokenResponse) ProtoMessage() {}

func (x *OAuthGetTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestra_v1_orchestra_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthGetTokenResponse.ProtoReflect.Descriptor instead.
func (*OAuthGetTokenResponse) Descriptor() ([]byte, []int) {
	return file_orchestra_v1_orchestra_proto_rawDescGZIP(), []int{3}
}

func (x *OAuthGetTokenResponse) GetResponse() *v11.GetTokenResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type OAuthStartSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source  Source                        `protobuf:"varint,1,opt,name=source,proto3,enum=orchestra.Source" json:"source,omitempty"`
	Request *v11.StartOAuthSessionRequest `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *OAuthStartSessionRequest) Reset() {
	*x = OAuthStartSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestra_v1_orchestra_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthStartSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthStartSessionRequest) ProtoMessage() {}

func (x *OAuthStartSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestra_v1_orchestra_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthStartSessionRequest.ProtoReflect.Descriptor instead.
func (*OAuthStartSessionRequest) Descriptor() ([]byte, []int) {
	return file_orchestra_v1_orchestra_proto_rawDescGZIP(), []int{4}
}

func (x *OAuthStartSessionRequest) GetSource() Source {
	if x != nil {
		return x.Source
	}
	return Source_LASTFM_SOURCE
}

func (x *OAuthStartSessionRequest) GetRequest() *v11.StartOAuthSessionRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type OAuthStartSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OAuthStartSessionResponse) Reset() {
	*x = OAuthStartSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestra_v1_orchestra_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthStartSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthStartSessionResponse) ProtoMessage() {}

func (x *OAuthStartSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestra_v1_orchestra_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthStartSessionResponse.ProtoReflect.Descriptor instead.
func (*OAuthStartSessionResponse) Descriptor() ([]byte, []int) {
	return file_orchestra_v1_orchestra_proto_rawDescGZIP(), []int{5}
}

var File_orchestra_v1_orchestra_proto protoreflect.FileDescriptor

var file_orchestra_v1_orchestra_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x72, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xee, 0x05, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x52,
	0x65, 0x63, 0x74, 0x52, 0x06, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x72, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6f, 0x72,
	0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x06, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00,
	0x52, 0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x12, 0x2e, 0x0a, 0x06,
	0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c,
	0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x48, 0x01, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x12, 0x28, 0x0a, 0x04,
	0x65, 0x76, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x48, 0x01,
	0x52, 0x04, 0x65, 0x76, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x06, 0x76, 0x61, 0x72, 0x69, 0x65, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e,
	0x56, 0x61, 0x72, 0x69, 0x65, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x48, 0x01, 0x52, 0x06,
	0x76, 0x61, 0x72, 0x69, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x74, 0x72, 0x69, 0x61, 0x6e, 0x67,
	0x75, 0x6c, 0x61, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x2e, 0x54, 0x72, 0x69, 0x61, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x48, 0x01, 0x52, 0x0a, 0x74, 0x72, 0x69, 0x61, 0x6e, 0x67, 0x75, 0x6c,
	0x61, 0x72, 0x12, 0x34, 0x0a, 0x08, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x48, 0x01, 0x52, 0x08,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x48, 0x02, 0x52, 0x05, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x12, 0x35, 0x0a, 0x09, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x2e, 0x57, 0x69, 0x74, 0x68, 0x54, 0x65, 0x78, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x48,
	0x02, 0x52, 0x08, 0x77, 0x69, 0x74, 0x68, 0x54, 0x65, 0x78, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x42, 0x0d, 0x0a, 0x0b, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x22, 0x37, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x74, 0x0a, 0x14, 0x4f,
	0x41, 0x75, 0x74, 0x68, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x2e,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x31,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x4d, 0x0a, 0x15, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x81, 0x01, 0x0a, 0x18, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x1b, 0x0a, 0x19, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2a, 0x1b, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x4c,
	0x41, 0x53, 0x54, 0x46, 0x4d, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x00, 0x32, 0x8b,
	0x02, 0x0a, 0x10, 0x4f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6f, 0x72,
	0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x4f, 0x41, 0x75, 0x74,
	0x68, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x2e, 0x6f, 0x72, 0x63, 0x68,
	0x65, 0x73, 0x74, 0x72, 0x61, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6f, 0x72, 0x63,
	0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x11,
	0x4f, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x23, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x2e, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74,
	0x72, 0x61, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x95, 0x01, 0x0a,
	0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x42, 0x0e,
	0x4f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x67,
	0x68, 0x74, 0x6d, 0x61, 0x72, 0x6c, 0x69, 0x6e, 0x2f, 0x6d, 0x75, 0x72, 0x75, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x2f,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x4f, 0x72, 0x63, 0x68, 0x65,
	0x73, 0x74, 0x72, 0x61, 0xca, 0x02, 0x09, 0x4f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61,
	0xe2, 0x02, 0x15, 0x4f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x4f, 0x72, 0x63, 0x68, 0x65,
	0x73, 0x74, 0x72, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orchestra_v1_orchestra_proto_rawDescOnce sync.Once
	file_orchestra_v1_orchestra_proto_rawDescData = file_orchestra_v1_orchestra_proto_rawDesc
)

func file_orchestra_v1_orchestra_proto_rawDescGZIP() []byte {
	file_orchestra_v1_orchestra_proto_rawDescOnce.Do(func() {
		file_orchestra_v1_orchestra_proto_rawDescData = protoimpl.X.CompressGZIP(file_orchestra_v1_orchestra_proto_rawDescData)
	})
	return file_orchestra_v1_orchestra_proto_rawDescData
}

var file_orchestra_v1_orchestra_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_orchestra_v1_orchestra_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_orchestra_v1_orchestra_proto_goTypes = []interface{}{
	(Source)(0),                          // 0: orchestra.Source
	(*GenerateRequest)(nil),              // 1: orchestra.GenerateRequest
	(*GenerateResponse)(nil),             // 2: orchestra.GenerateResponse
	(*OAuthGetTokenRequest)(nil),         // 3: orchestra.OAuthGetTokenRequest
	(*OAuthGetTokenResponse)(nil),        // 4: orchestra.OAuthGetTokenResponse
	(*OAuthStartSessionRequest)(nil),     // 5: orchestra.OAuthStartSessionRequest
	(*OAuthStartSessionResponse)(nil),    // 6: orchestra.OAuthStartSessionResponse
	(*v1.Rect)(nil),                      // 7: shared.Rect
	(*v11.TracksRequest)(nil),            // 8: source.TracksRequest
	(*v11.AlbumsRequest)(nil),            // 9: source.AlbumsRequest
	(*v11.ArtistsRequest)(nil),           // 10: source.ArtistsRequest
	(*v12.RandomParams)(nil),             // 11: layout.RandomParams
	(*v12.EvenParams)(nil),               // 12: layout.EvenParams
	(*v12.VariedParams)(nil),             // 13: layout.VariedParams
	(*v12.TriangularParams)(nil),         // 14: layout.TriangularParams
	(*v12.WeightedParams)(nil),           // 15: layout.WeightedParams
	(v1.FileType)(0),                     // 16: shared.FileType
	(*v13.BasicParams)(nil),              // 17: render.BasicParams
	(*v13.WithTextParams)(nil),           // 18: render.WithTextParams
	(*v1.Image)(nil),                     // 19: shared.Image
	(*v11.GetTokenRequest)(nil),          // 20: source.GetTokenRequest
	(*v11.GetTokenResponse)(nil),         // 21: source.GetTokenResponse
	(*v11.StartOAuthSessionRequest)(nil), // 22: source.StartOAuthSessionRequest
}
var file_orchestra_v1_orchestra_proto_depIdxs = []int32{
	7,  // 0: orchestra.GenerateRequest.bounds:type_name -> shared.Rect
	0,  // 1: orchestra.GenerateRequest.source:type_name -> orchestra.Source
	8,  // 2: orchestra.GenerateRequest.tracks:type_name -> source.TracksRequest
	9,  // 3: orchestra.GenerateRequest.albums:type_name -> source.AlbumsRequest
	10, // 4: orchestra.GenerateRequest.artists:type_name -> source.ArtistsRequest
	11, // 5: orchestra.GenerateRequest.random:type_name -> layout.RandomParams
	12, // 6: orchestra.GenerateRequest.even:type_name -> layout.EvenParams
	13, // 7: orchestra.GenerateRequest.varied:type_name -> layout.VariedParams
	14, // 8: orchestra.GenerateRequest.triangular:type_name -> layout.TriangularParams
	15, // 9: orchestra.GenerateRequest.weighted:type_name -> layout.WeightedParams
	16, // 10: orchestra.GenerateRequest.file_type:type_name -> shared.FileType
	17, // 11: orchestra.GenerateRequest.basic:type_name -> render.BasicParams
	18, // 12: orchestra.GenerateRequest.with_text:type_name -> render.WithTextParams
	19, // 13: orchestra.GenerateResponse.image:type_name -> shared.Image
	0,  // 14: orchestra.OAuthGetTokenRequest.source:type_name -> orchestra.Source
	20, // 15: orchestra.OAuthGetTokenRequest.request:type_name -> source.GetTokenRequest
	21, // 16: orchestra.OAuthGetTokenResponse.response:type_name -> source.GetTokenResponse
	0,  // 17: orchestra.OAuthStartSessionRequest.source:type_name -> orchestra.Source
	22, // 18: orchestra.OAuthStartSessionRequest.request:type_name -> source.StartOAuthSessionRequest
	1,  // 19: orchestra.OrchestraService.Generate:input_type -> orchestra.GenerateRequest
	3,  // 20: orchestra.OrchestraService.OAuthGetToken:input_type -> orchestra.OAuthGetTokenRequest
	5,  // 21: orchestra.OrchestraService.OAuthStartSession:input_type -> orchestra.OAuthStartSessionRequest
	2,  // 22: orchestra.OrchestraService.Generate:output_type -> orchestra.GenerateResponse
	4,  // 23: orchestra.OrchestraService.OAuthGetToken:output_type -> orchestra.OAuthGetTokenResponse
	6,  // 24: orchestra.OrchestraService.OAuthStartSession:output_type -> orchestra.OAuthStartSessionResponse
	22, // [22:25] is the sub-list for method output_type
	19, // [19:22] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_orchestra_v1_orchestra_proto_init() }
func file_orchestra_v1_orchestra_proto_init() {
	if File_orchestra_v1_orchestra_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orchestra_v1_orchestra_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateRequest); i {
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
		file_orchestra_v1_orchestra_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateResponse); i {
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
		file_orchestra_v1_orchestra_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthGetTokenRequest); i {
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
		file_orchestra_v1_orchestra_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthGetTokenResponse); i {
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
		file_orchestra_v1_orchestra_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthStartSessionRequest); i {
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
		file_orchestra_v1_orchestra_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthStartSessionResponse); i {
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
	file_orchestra_v1_orchestra_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GenerateRequest_Tracks)(nil),
		(*GenerateRequest_Albums)(nil),
		(*GenerateRequest_Artists)(nil),
		(*GenerateRequest_Random)(nil),
		(*GenerateRequest_Even)(nil),
		(*GenerateRequest_Varied)(nil),
		(*GenerateRequest_Triangular)(nil),
		(*GenerateRequest_Weighted)(nil),
		(*GenerateRequest_Basic)(nil),
		(*GenerateRequest_WithText)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orchestra_v1_orchestra_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orchestra_v1_orchestra_proto_goTypes,
		DependencyIndexes: file_orchestra_v1_orchestra_proto_depIdxs,
		EnumInfos:         file_orchestra_v1_orchestra_proto_enumTypes,
		MessageInfos:      file_orchestra_v1_orchestra_proto_msgTypes,
	}.Build()
	File_orchestra_v1_orchestra_proto = out.File
	file_orchestra_v1_orchestra_proto_rawDesc = nil
	file_orchestra_v1_orchestra_proto_goTypes = nil
	file_orchestra_v1_orchestra_proto_depIdxs = nil
}
