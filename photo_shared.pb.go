// Code generated by protoc-gen-go. DO NOT EDIT.
// source: photo_shared.proto

package wcity

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ProfilePhoto struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserEmail            string   `protobuf:"bytes,2,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Timestamp            string   `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Selected             bool     `protobuf:"varint,5,opt,name=selected,proto3" json:"selected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfilePhoto) Reset()         { *m = ProfilePhoto{} }
func (m *ProfilePhoto) String() string { return proto.CompactTextString(m) }
func (*ProfilePhoto) ProtoMessage()    {}
func (*ProfilePhoto) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbe3263bbb10a08, []int{0}
}

func (m *ProfilePhoto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfilePhoto.Unmarshal(m, b)
}
func (m *ProfilePhoto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfilePhoto.Marshal(b, m, deterministic)
}
func (m *ProfilePhoto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfilePhoto.Merge(m, src)
}
func (m *ProfilePhoto) XXX_Size() int {
	return xxx_messageInfo_ProfilePhoto.Size(m)
}
func (m *ProfilePhoto) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfilePhoto.DiscardUnknown(m)
}

var xxx_messageInfo_ProfilePhoto proto.InternalMessageInfo

func (m *ProfilePhoto) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProfilePhoto) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

func (m *ProfilePhoto) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ProfilePhoto) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *ProfilePhoto) GetSelected() bool {
	if m != nil {
		return m.Selected
	}
	return false
}

type CityPhoto struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CityId               int32    `protobuf:"varint,2,opt,name=cityId,proto3" json:"cityId,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Timestamp            string   `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Selected             bool     `protobuf:"varint,5,opt,name=selected,proto3" json:"selected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CityPhoto) Reset()         { *m = CityPhoto{} }
func (m *CityPhoto) String() string { return proto.CompactTextString(m) }
func (*CityPhoto) ProtoMessage()    {}
func (*CityPhoto) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbe3263bbb10a08, []int{1}
}

func (m *CityPhoto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CityPhoto.Unmarshal(m, b)
}
func (m *CityPhoto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CityPhoto.Marshal(b, m, deterministic)
}
func (m *CityPhoto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CityPhoto.Merge(m, src)
}
func (m *CityPhoto) XXX_Size() int {
	return xxx_messageInfo_CityPhoto.Size(m)
}
func (m *CityPhoto) XXX_DiscardUnknown() {
	xxx_messageInfo_CityPhoto.DiscardUnknown(m)
}

var xxx_messageInfo_CityPhoto proto.InternalMessageInfo

func (m *CityPhoto) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CityPhoto) GetCityId() int32 {
	if m != nil {
		return m.CityId
	}
	return 0
}

func (m *CityPhoto) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CityPhoto) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *CityPhoto) GetSelected() bool {
	if m != nil {
		return m.Selected
	}
	return false
}

type PlacePhoto struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PlaceId              int32    `protobuf:"varint,2,opt,name=placeId,proto3" json:"placeId,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Timestamp            string   `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Selected             bool     `protobuf:"varint,5,opt,name=selected,proto3" json:"selected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlacePhoto) Reset()         { *m = PlacePhoto{} }
func (m *PlacePhoto) String() string { return proto.CompactTextString(m) }
func (*PlacePhoto) ProtoMessage()    {}
func (*PlacePhoto) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbe3263bbb10a08, []int{2}
}

func (m *PlacePhoto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlacePhoto.Unmarshal(m, b)
}
func (m *PlacePhoto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlacePhoto.Marshal(b, m, deterministic)
}
func (m *PlacePhoto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlacePhoto.Merge(m, src)
}
func (m *PlacePhoto) XXX_Size() int {
	return xxx_messageInfo_PlacePhoto.Size(m)
}
func (m *PlacePhoto) XXX_DiscardUnknown() {
	xxx_messageInfo_PlacePhoto.DiscardUnknown(m)
}

var xxx_messageInfo_PlacePhoto proto.InternalMessageInfo

func (m *PlacePhoto) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PlacePhoto) GetPlaceId() int32 {
	if m != nil {
		return m.PlaceId
	}
	return 0
}

func (m *PlacePhoto) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *PlacePhoto) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *PlacePhoto) GetSelected() bool {
	if m != nil {
		return m.Selected
	}
	return false
}

type PostPhoto struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PostId               string   `protobuf:"bytes,2,opt,name=postId,proto3" json:"postId,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Timestamp            string   `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Selected             bool     `protobuf:"varint,5,opt,name=selected,proto3" json:"selected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostPhoto) Reset()         { *m = PostPhoto{} }
func (m *PostPhoto) String() string { return proto.CompactTextString(m) }
func (*PostPhoto) ProtoMessage()    {}
func (*PostPhoto) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbe3263bbb10a08, []int{3}
}

func (m *PostPhoto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostPhoto.Unmarshal(m, b)
}
func (m *PostPhoto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostPhoto.Marshal(b, m, deterministic)
}
func (m *PostPhoto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostPhoto.Merge(m, src)
}
func (m *PostPhoto) XXX_Size() int {
	return xxx_messageInfo_PostPhoto.Size(m)
}
func (m *PostPhoto) XXX_DiscardUnknown() {
	xxx_messageInfo_PostPhoto.DiscardUnknown(m)
}

var xxx_messageInfo_PostPhoto proto.InternalMessageInfo

func (m *PostPhoto) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PostPhoto) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

func (m *PostPhoto) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *PostPhoto) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *PostPhoto) GetSelected() bool {
	if m != nil {
		return m.Selected
	}
	return false
}

func init() {
	proto.RegisterType((*ProfilePhoto)(nil), "wcity.ProfilePhoto")
	proto.RegisterType((*CityPhoto)(nil), "wcity.CityPhoto")
	proto.RegisterType((*PlacePhoto)(nil), "wcity.PlacePhoto")
	proto.RegisterType((*PostPhoto)(nil), "wcity.PostPhoto")
}

func init() { proto.RegisterFile("photo_shared.proto", fileDescriptor_dcbe3263bbb10a08) }

var fileDescriptor_dcbe3263bbb10a08 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x86, 0xc9, 0xae, 0x7b, 0xde, 0x0e, 0x22, 0x47, 0x0a, 0x89, 0x62, 0x21, 0x57, 0x59, 0x6d,
	0x23, 0xbe, 0xc0, 0x1d, 0x16, 0x76, 0x43, 0x7c, 0x00, 0x59, 0x37, 0xd1, 0x0b, 0x64, 0x49, 0x48,
	0x72, 0x88, 0x95, 0x60, 0xe5, 0x73, 0xf8, 0xa4, 0x26, 0x71, 0xdd, 0x6b, 0x6e, 0xbb, 0xed, 0xe6,
	0xfb, 0x67, 0x8a, 0x8f, 0xe1, 0x07, 0x6a, 0x77, 0x26, 0x98, 0x67, 0xbf, 0x6b, 0x9d, 0x14, 0x8d,
	0x75, 0x11, 0x68, 0xf5, 0xde, 0xa9, 0xf0, 0xb1, 0xfe, 0x26, 0x70, 0x86, 0xce, 0xbc, 0x2a, 0x2d,
	0x31, 0x1d, 0xd1, 0x73, 0x28, 0x94, 0x60, 0xe4, 0x86, 0xdc, 0x56, 0x3c, 0x4e, 0xf4, 0x1a, 0xea,
	0xbd, 0x97, 0xee, 0xa1, 0x6f, 0x95, 0x66, 0x45, 0x8c, 0x6b, 0x7e, 0x08, 0xe8, 0x0a, 0xca, 0xbd,
	0xd3, 0xac, 0xcc, 0x79, 0x1a, 0xd3, 0x7d, 0x50, 0xbd, 0xf4, 0xa1, 0xed, 0x2d, 0x3b, 0xf9, 0xbb,
	0x1f, 0x03, 0x7a, 0x05, 0x4b, 0x2f, 0xb5, 0xec, 0x82, 0x14, 0xac, 0x8a, 0xcb, 0x25, 0x1f, 0x79,
	0xfd, 0x09, 0xf5, 0x36, 0x2a, 0x1d, 0xd7, 0xb8, 0x80, 0x45, 0xf2, 0x7d, 0x14, 0xd9, 0xa1, 0xe2,
	0x03, 0xcd, 0x2a, 0xf0, 0x45, 0x00, 0x50, 0xb7, 0xdd, 0xc4, 0x27, 0x18, 0x9c, 0xda, 0xb4, 0x1d,
	0x1d, 0xfe, 0x71, 0xee, 0x2f, 0xa0, 0xf1, 0x61, 0xf2, 0x0b, 0x36, 0x2e, 0x07, 0x83, 0x9a, 0x0f,
	0x34, 0xa7, 0xc0, 0xe6, 0x1e, 0x2e, 0x95, 0x69, 0xde, 0x9c, 0xed, 0x9a, 0x5c, 0x91, 0x26, 0x97,
	0xe7, 0x29, 0x77, 0x67, 0xb3, 0xc2, 0x03, 0x60, 0xea, 0x11, 0x92, 0x9f, 0xa2, 0xe4, 0xb8, 0x7d,
	0x59, 0xe4, 0x5a, 0xdd, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xbb, 0x09, 0xa4, 0x6c, 0x02,
	0x00, 0x00,
}
