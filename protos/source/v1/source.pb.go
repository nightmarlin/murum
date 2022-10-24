// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: source/v1/source.proto

package v1

import (
	v1 "github.com/nightmarlin/murum/protos/shared/v1"
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

type Track struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                               // The Name of the Track.
	PlayCount  int32       `protobuf:"varint,2,opt,name=play_count,json=playCount,proto3" json:"play_count,omitempty"`   // The number of times the user has played this Track.
	AlbumName  string      `protobuf:"bytes,3,opt,name=album_name,json=albumName,proto3" json:"album_name,omitempty"`    // The name of the Album the Track appears on.
	ArtistName string      `protobuf:"bytes,4,opt,name=artist_name,json=artistName,proto3" json:"artist_name,omitempty"` // The name of the artist(s) who performed this Track.
	Images     []*v1.Image `protobuf:"bytes,5,rep,name=images,proto3" json:"images,omitempty"`                           // The artwork(s) for this Track.
}

func (x *Track) Reset() {
	*x = Track{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_v1_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Track) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Track) ProtoMessage() {}

func (x *Track) ProtoReflect() protoreflect.Message {
	mi := &file_source_v1_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Track.ProtoReflect.Descriptor instead.
func (*Track) Descriptor() ([]byte, []int) {
	return file_source_v1_source_proto_rawDescGZIP(), []int{0}
}

func (x *Track) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Track) GetPlayCount() int32 {
	if x != nil {
		return x.PlayCount
	}
	return 0
}

func (x *Track) GetAlbumName() string {
	if x != nil {
		return x.AlbumName
	}
	return ""
}

func (x *Track) GetArtistName() string {
	if x != nil {
		return x.ArtistName
	}
	return ""
}

func (x *Track) GetImages() []*v1.Image {
	if x != nil {
		return x.Images
	}
	return nil
}

type Album struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                               // The Name of the Album.
	PlayCount  int32       `protobuf:"varint,2,opt,name=play_count,json=playCount,proto3" json:"play_count,omitempty"`   // The number of times the user has played this Album.
	ArtistName string      `protobuf:"bytes,3,opt,name=artist_name,json=artistName,proto3" json:"artist_name,omitempty"` // The name of the artist(s) who performed this Album.
	Images     []*v1.Image `protobuf:"bytes,4,rep,name=images,proto3" json:"images,omitempty"`                           // The artwork(s) for this Album.
}

func (x *Album) Reset() {
	*x = Album{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_v1_source_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Album) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Album) ProtoMessage() {}

func (x *Album) ProtoReflect() protoreflect.Message {
	mi := &file_source_v1_source_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Album.ProtoReflect.Descriptor instead.
func (*Album) Descriptor() ([]byte, []int) {
	return file_source_v1_source_proto_rawDescGZIP(), []int{1}
}

func (x *Album) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Album) GetPlayCount() int32 {
	if x != nil {
		return x.PlayCount
	}
	return 0
}

func (x *Album) GetArtistName() string {
	if x != nil {
		return x.ArtistName
	}
	return ""
}

func (x *Album) GetImages() []*v1.Image {
	if x != nil {
		return x.Images
	}
	return nil
}

type Artist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                             // The Name of the Artist.
	PlayCount int32       `protobuf:"varint,2,opt,name=play_count,json=playCount,proto3" json:"play_count,omitempty"` // The number of times the user has played this Artist.
	Images    []*v1.Image `protobuf:"bytes,3,rep,name=images,proto3" json:"images,omitempty"`                         // The artwork(s) for this Artist.
}

func (x *Artist) Reset() {
	*x = Artist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_v1_source_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artist) ProtoMessage() {}

func (x *Artist) ProtoReflect() protoreflect.Message {
	mi := &file_source_v1_source_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artist.ProtoReflect.Descriptor instead.
func (*Artist) Descriptor() ([]byte, []int) {
	return file_source_v1_source_proto_rawDescGZIP(), []int{2}
}

func (x *Artist) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Artist) GetPlayCount() int32 {
	if x != nil {
		return x.PlayCount
	}
	return 0
}

func (x *Artist) GetImages() []*v1.Image {
	if x != nil {
		return x.Images
	}
	return nil
}

type ArtistsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string        `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          // The user to request Artist s info for.
	Lang      string        `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`                            // The language to return results in. Default: en-GB.
	Limit     int32         `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`                         // The maximum number of Artist s to return.
	TimeRange *v1.TimeRange `protobuf:"bytes,4,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"` // The time period to filter scrobbles within.
}

func (x *ArtistsRequest) Reset() {
	*x = ArtistsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_v1_source_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtistsRequest) ProtoMessage() {}

func (x *ArtistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_source_v1_source_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtistsRequest.ProtoReflect.Descriptor instead.
func (*ArtistsRequest) Descriptor() ([]byte, []int) {
	return file_source_v1_source_proto_rawDescGZIP(), []int{3}
}

func (x *ArtistsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ArtistsRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *ArtistsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ArtistsRequest) GetTimeRange() *v1.TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

type ArtistsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Artists []*Artist `protobuf:"bytes,1,rep,name=artists,proto3" json:"artists,omitempty"`
}

func (x *ArtistsResponse) Reset() {
	*x = ArtistsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_v1_source_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtistsResponse) ProtoMessage() {}

func (x *ArtistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_source_v1_source_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtistsResponse.ProtoReflect.Descriptor instead.
func (*ArtistsResponse) Descriptor() ([]byte, []int) {
	return file_source_v1_source_proto_rawDescGZIP(), []int{4}
}

func (x *ArtistsResponse) GetArtists() []*Artist {
	if x != nil {
		return x.Artists
	}
	return nil
}

type AlbumsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string        `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          // The user to request Artist s info for.
	Lang      string        `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`                            // The language to return results in. Default: en-GB.
	Limit     int32         `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`                         // The maximum number of Album s to return.
	TimeRange *v1.TimeRange `protobuf:"bytes,4,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"` // The time period to filter scrobbles within.
	Artist    string        `protobuf:"bytes,5,opt,name=artist,proto3" json:"artist,omitempty"`                        // Which Artist s Album s to return.
}

func (x *AlbumsRequest) Reset() {
	*x = AlbumsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_v1_source_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumsRequest) ProtoMessage() {}

func (x *AlbumsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_source_v1_source_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumsRequest.ProtoReflect.Descriptor instead.
func (*AlbumsRequest) Descriptor() ([]byte, []int) {
	return file_source_v1_source_proto_rawDescGZIP(), []int{5}
}

func (x *AlbumsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AlbumsRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *AlbumsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *AlbumsRequest) GetTimeRange() *v1.TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

func (x *AlbumsRequest) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

type AlbumsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Albums []*Album `protobuf:"bytes,1,rep,name=albums,proto3" json:"albums,omitempty"`
}

func (x *AlbumsResponse) Reset() {
	*x = AlbumsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_v1_source_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumsResponse) ProtoMessage() {}

func (x *AlbumsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_source_v1_source_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumsResponse.ProtoReflect.Descriptor instead.
func (*AlbumsResponse) Descriptor() ([]byte, []int) {
	return file_source_v1_source_proto_rawDescGZIP(), []int{6}
}

func (x *AlbumsResponse) GetAlbums() []*Album {
	if x != nil {
		return x.Albums
	}
	return nil
}

type TracksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string        `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          // The user to request Track s info for.
	Lang      string        `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`                            // The language to return results in. Default: en-GB.
	Limit     int32         `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`                         // The maximum number of Track s to return.
	TimeRange *v1.TimeRange `protobuf:"bytes,4,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"` // The time period to filter scrobbles within.
	Artist    string        `protobuf:"bytes,5,opt,name=artist,proto3" json:"artist,omitempty"`                        // Which Artist s Track s to return. If Album is set, Artist must appear on it.
	Album     string        `protobuf:"bytes,6,opt,name=album,proto3" json:"album,omitempty"`                          // Which Album s Track s to return. If Artist is set, they must appear on it.
}

func (x *TracksRequest) Reset() {
	*x = TracksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_v1_source_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TracksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracksRequest) ProtoMessage() {}

func (x *TracksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_source_v1_source_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TracksRequest.ProtoReflect.Descriptor instead.
func (*TracksRequest) Descriptor() ([]byte, []int) {
	return file_source_v1_source_proto_rawDescGZIP(), []int{7}
}

func (x *TracksRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TracksRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *TracksRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *TracksRequest) GetTimeRange() *v1.TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

func (x *TracksRequest) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

func (x *TracksRequest) GetAlbum() string {
	if x != nil {
		return x.Album
	}
	return ""
}

type TracksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tracks []*Track `protobuf:"bytes,1,rep,name=tracks,proto3" json:"tracks,omitempty"`
}

func (x *TracksResponse) Reset() {
	*x = TracksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_v1_source_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TracksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracksResponse) ProtoMessage() {}

func (x *TracksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_source_v1_source_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TracksResponse.ProtoReflect.Descriptor instead.
func (*TracksResponse) Descriptor() ([]byte, []int) {
	return file_source_v1_source_proto_rawDescGZIP(), []int{8}
}

func (x *TracksResponse) GetTracks() []*Track {
	if x != nil {
		return x.Tracks
	}
	return nil
}

type IsUserAuthedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // The user to check the auth state for.
}

func (x *IsUserAuthedRequest) Reset() {
	*x = IsUserAuthedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_v1_source_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsUserAuthedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsUserAuthedRequest) ProtoMessage() {}

func (x *IsUserAuthedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_source_v1_source_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsUserAuthedRequest.ProtoReflect.Descriptor instead.
func (*IsUserAuthedRequest) Descriptor() ([]byte, []int) {
	return file_source_v1_source_proto_rawDescGZIP(), []int{9}
}

func (x *IsUserAuthedRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type IsUserAuthedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether murum is authorized to use this user's info by the Source.
	// For services that don't implement OAuthService, this will always be true.
	HasAuth bool `protobuf:"varint,1,opt,name=has_auth,json=hasAuth,proto3" json:"has_auth,omitempty"`
}

func (x *IsUserAuthedResponse) Reset() {
	*x = IsUserAuthedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_v1_source_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsUserAuthedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsUserAuthedResponse) ProtoMessage() {}

func (x *IsUserAuthedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_source_v1_source_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsUserAuthedResponse.ProtoReflect.Descriptor instead.
func (*IsUserAuthedResponse) Descriptor() ([]byte, []int) {
	return file_source_v1_source_proto_rawDescGZIP(), []int{10}
}

func (x *IsUserAuthedResponse) GetHasAuth() bool {
	if x != nil {
		return x.HasAuth
	}
	return false
}

var File_source_v1_source_proto protoreflect.FileDescriptor

var file_source_v1_source_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x05, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x6c, 0x61,
	0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6c, 0x62, 0x75,
	0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x82, 0x01,
	0x0a, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x72,
	0x74, 0x69, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x22, 0x62, 0x0a, 0x06, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x25, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x06,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x0e, 0x41, 0x72, 0x74, 0x69, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x30, 0x0a, 0x0a,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x3b,
	0x0a, 0x0f, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x0d,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x30, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x0e, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06,
	0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x06, 0x61, 0x6c, 0x62,
	0x75, 0x6d, 0x73, 0x22, 0xb2, 0x01, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61,
	0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x30, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x72,
	0x74, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x22, 0x37, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x06, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x73, 0x22, 0x2e, 0x0a, 0x13, 0x49, 0x73, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x31, 0x0a, 0x14, 0x49, 0x73, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73,
	0x41, 0x75, 0x74, 0x68, 0x32, 0x88, 0x02, 0x0a, 0x0d, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x73, 0x12, 0x16, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x12, 0x15, 0x2e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x73, 0x12, 0x15, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x49, 0x73, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x64, 0x12, 0x1b, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x49, 0x73,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x49, 0x73, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x75, 0x74, 0x68, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x80, 0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x0b,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x67, 0x68, 0x74, 0x6d,
	0x61, 0x72, 0x6c, 0x69, 0x6e, 0x2f, 0x6d, 0x75, 0x72, 0x75, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53,
	0x58, 0x58, 0xaa, 0x02, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0xca, 0x02, 0x06, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0xe2, 0x02, 0x12, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_source_v1_source_proto_rawDescOnce sync.Once
	file_source_v1_source_proto_rawDescData = file_source_v1_source_proto_rawDesc
)

func file_source_v1_source_proto_rawDescGZIP() []byte {
	file_source_v1_source_proto_rawDescOnce.Do(func() {
		file_source_v1_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_source_v1_source_proto_rawDescData)
	})
	return file_source_v1_source_proto_rawDescData
}

var file_source_v1_source_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_source_v1_source_proto_goTypes = []interface{}{
	(*Track)(nil),                // 0: source.Track
	(*Album)(nil),                // 1: source.Album
	(*Artist)(nil),               // 2: source.Artist
	(*ArtistsRequest)(nil),       // 3: source.ArtistsRequest
	(*ArtistsResponse)(nil),      // 4: source.ArtistsResponse
	(*AlbumsRequest)(nil),        // 5: source.AlbumsRequest
	(*AlbumsResponse)(nil),       // 6: source.AlbumsResponse
	(*TracksRequest)(nil),        // 7: source.TracksRequest
	(*TracksResponse)(nil),       // 8: source.TracksResponse
	(*IsUserAuthedRequest)(nil),  // 9: source.IsUserAuthedRequest
	(*IsUserAuthedResponse)(nil), // 10: source.IsUserAuthedResponse
	(*v1.Image)(nil),             // 11: shared.Image
	(*v1.TimeRange)(nil),         // 12: shared.TimeRange
}
var file_source_v1_source_proto_depIdxs = []int32{
	11, // 0: source.Track.images:type_name -> shared.Image
	11, // 1: source.Album.images:type_name -> shared.Image
	11, // 2: source.Artist.images:type_name -> shared.Image
	12, // 3: source.ArtistsRequest.time_range:type_name -> shared.TimeRange
	2,  // 4: source.ArtistsResponse.artists:type_name -> source.Artist
	12, // 5: source.AlbumsRequest.time_range:type_name -> shared.TimeRange
	1,  // 6: source.AlbumsResponse.albums:type_name -> source.Album
	12, // 7: source.TracksRequest.time_range:type_name -> shared.TimeRange
	0,  // 8: source.TracksResponse.tracks:type_name -> source.Track
	3,  // 9: source.SourceService.Artists:input_type -> source.ArtistsRequest
	5,  // 10: source.SourceService.Albums:input_type -> source.AlbumsRequest
	7,  // 11: source.SourceService.Tracks:input_type -> source.TracksRequest
	9,  // 12: source.SourceService.IsUserAuthed:input_type -> source.IsUserAuthedRequest
	4,  // 13: source.SourceService.Artists:output_type -> source.ArtistsResponse
	6,  // 14: source.SourceService.Albums:output_type -> source.AlbumsResponse
	8,  // 15: source.SourceService.Tracks:output_type -> source.TracksResponse
	10, // 16: source.SourceService.IsUserAuthed:output_type -> source.IsUserAuthedResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_source_v1_source_proto_init() }
func file_source_v1_source_proto_init() {
	if File_source_v1_source_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_source_v1_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Track); i {
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
		file_source_v1_source_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Album); i {
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
		file_source_v1_source_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artist); i {
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
		file_source_v1_source_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtistsRequest); i {
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
		file_source_v1_source_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtistsResponse); i {
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
		file_source_v1_source_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumsRequest); i {
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
		file_source_v1_source_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumsResponse); i {
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
		file_source_v1_source_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TracksRequest); i {
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
		file_source_v1_source_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TracksResponse); i {
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
		file_source_v1_source_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsUserAuthedRequest); i {
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
		file_source_v1_source_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsUserAuthedResponse); i {
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
			RawDescriptor: file_source_v1_source_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_source_v1_source_proto_goTypes,
		DependencyIndexes: file_source_v1_source_proto_depIdxs,
		MessageInfos:      file_source_v1_source_proto_msgTypes,
	}.Build()
	File_source_v1_source_proto = out.File
	file_source_v1_source_proto_rawDesc = nil
	file_source_v1_source_proto_goTypes = nil
	file_source_v1_source_proto_depIdxs = nil
}
