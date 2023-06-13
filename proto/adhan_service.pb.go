// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: adhan_service.proto

package proto

import (
	_ "github.com/mnadev/limestone/proto/google/api"
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

type AdhanFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique id associated with this file.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The masjid id associated with this file.
	MasjidId string `protobuf:"bytes,2,opt,name=masjid_id,json=masjidId,proto3" json:"masjid_id,omitempty"`
	// The bytes of the audio file. This should be in LINEAR16 format at
	// a sample rate of 16 kHz.
	File []byte `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
	// The create time of the file. This field is output only.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The update time of the file. This field is output only.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *AdhanFile) Reset() {
	*x = AdhanFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adhan_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdhanFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdhanFile) ProtoMessage() {}

func (x *AdhanFile) ProtoReflect() protoreflect.Message {
	mi := &file_adhan_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdhanFile.ProtoReflect.Descriptor instead.
func (*AdhanFile) Descriptor() ([]byte, []int) {
	return file_adhan_service_proto_rawDescGZIP(), []int{0}
}

func (x *AdhanFile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AdhanFile) GetMasjidId() string {
	if x != nil {
		return x.MasjidId
	}
	return ""
}

func (x *AdhanFile) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *AdhanFile) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *AdhanFile) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type CreateAdhanFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The adhan file to create.
	AdhanFile *AdhanFile `protobuf:"bytes,1,opt,name=adhan_file,json=adhanFile,proto3" json:"adhan_file,omitempty"`
}

func (x *CreateAdhanFileRequest) Reset() {
	*x = CreateAdhanFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adhan_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAdhanFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdhanFileRequest) ProtoMessage() {}

func (x *CreateAdhanFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_adhan_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAdhanFileRequest.ProtoReflect.Descriptor instead.
func (*CreateAdhanFileRequest) Descriptor() ([]byte, []int) {
	return file_adhan_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAdhanFileRequest) GetAdhanFile() *AdhanFile {
	if x != nil {
		return x.AdhanFile
	}
	return nil
}

type UpdateAdhanFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The adhan file to update to. The file associated with the id
	// will be updated to this given adhan file.
	AdhanFile *AdhanFile `protobuf:"bytes,1,opt,name=adhan_file,json=adhanFile,proto3" json:"adhan_file,omitempty"`
}

func (x *UpdateAdhanFileRequest) Reset() {
	*x = UpdateAdhanFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adhan_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAdhanFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdhanFileRequest) ProtoMessage() {}

func (x *UpdateAdhanFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_adhan_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAdhanFileRequest.ProtoReflect.Descriptor instead.
func (*UpdateAdhanFileRequest) Descriptor() ([]byte, []int) {
	return file_adhan_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAdhanFileRequest) GetAdhanFile() *AdhanFile {
	if x != nil {
		return x.AdhanFile
	}
	return nil
}

type GetAdhanFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The masjid id to get the adhan file for.
	MasjidId string `protobuf:"bytes,1,opt,name=masjid_id,json=masjidId,proto3" json:"masjid_id,omitempty"`
}

func (x *GetAdhanFileRequest) Reset() {
	*x = GetAdhanFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adhan_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdhanFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdhanFileRequest) ProtoMessage() {}

func (x *GetAdhanFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_adhan_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdhanFileRequest.ProtoReflect.Descriptor instead.
func (*GetAdhanFileRequest) Descriptor() ([]byte, []int) {
	return file_adhan_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetAdhanFileRequest) GetMasjidId() string {
	if x != nil {
		return x.MasjidId
	}
	return ""
}

type DeleteAdhanFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The adhan file id of the file to delete.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAdhanFileRequest) Reset() {
	*x = DeleteAdhanFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adhan_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAdhanFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAdhanFileRequest) ProtoMessage() {}

func (x *DeleteAdhanFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_adhan_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAdhanFileRequest.ProtoReflect.Descriptor instead.
func (*DeleteAdhanFileRequest) Descriptor() ([]byte, []int) {
	return file_adhan_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteAdhanFileRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteAdhanFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAdhanFileResponse) Reset() {
	*x = DeleteAdhanFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adhan_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAdhanFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAdhanFileResponse) ProtoMessage() {}

func (x *DeleteAdhanFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_adhan_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAdhanFileResponse.ProtoReflect.Descriptor instead.
func (*DeleteAdhanFileResponse) Descriptor() ([]byte, []int) {
	return file_adhan_service_proto_rawDescGZIP(), []int{5}
}

var File_adhan_service_proto protoreflect.FileDescriptor

var file_adhan_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x64, 0x68, 0x61, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6c, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc6, 0x01, 0x0a, 0x09, 0x41, 0x64, 0x68, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x61, 0x73, 0x6a, 0x69, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x61, 0x73, 0x6a, 0x69, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x3b,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x64, 0x68, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x33, 0x0a, 0x0a, 0x61, 0x64, 0x68, 0x61, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x6f,
	0x6e, 0x65, 0x2e, 0x41, 0x64, 0x68, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x09, 0x61, 0x64,
	0x68, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x4d, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x64, 0x68, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x33, 0x0a, 0x0a, 0x61, 0x64, 0x68, 0x61, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x6f, 0x6e,
	0x65, 0x2e, 0x41, 0x64, 0x68, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x09, 0x61, 0x64, 0x68,
	0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x32, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x64, 0x68,
	0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x61, 0x73, 0x6a, 0x69, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x61, 0x73, 0x6a, 0x69, 0x64, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x68, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64,
	0x68, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xbf, 0x03, 0x0a, 0x0c, 0x41, 0x64, 0x68, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x69, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x68, 0x61, 0x6e, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x6c, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x68, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6c, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x6f,
	0x6e, 0x65, 0x2e, 0x41, 0x64, 0x68, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x1d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x61, 0x64, 0x68, 0x61, 0x6e, 0x2f,
	0x66, 0x69, 0x6c, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x69, 0x0a, 0x0f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x68, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x21,
	0x2e, 0x6c, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x64, 0x68, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x6c, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2e, 0x41, 0x64,
	0x68, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a,
	0x01, 0x2a, 0x22, 0x12, 0x2f, 0x61, 0x64, 0x68, 0x61, 0x6e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x65, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x64, 0x68,
	0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1e, 0x2e, 0x6c, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x6f,
	0x6e, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x68, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6c, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x6f,
	0x6e, 0x65, 0x2e, 0x41, 0x64, 0x68, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x1f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61, 0x64, 0x68, 0x61, 0x6e, 0x2f, 0x66, 0x69, 0x6c,
	0x65, 0x2f, 0x7b, 0x6d, 0x61, 0x73, 0x6a, 0x69, 0x64, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x72, 0x0a,
	0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x68, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x21, 0x2e, 0x6c, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x64, 0x68, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6c, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x64, 0x68, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x2a,
	0x10, 0x2f, 0x61, 0x64, 0x68, 0x61, 0x6e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x42, 0x89, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x6c, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x6f, 0x6e, 0x65, 0x42, 0x11, 0x41, 0x64, 0x68, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6e, 0x61, 0x64, 0x65, 0x76, 0x2f, 0x6c, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x4c, 0x58,
	0x58, 0xaa, 0x02, 0x09, 0x4c, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0xca, 0x02, 0x09,
	0x4c, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0xe2, 0x02, 0x15, 0x4c, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x6f, 0x6e, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x09, 0x4c, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_adhan_service_proto_rawDescOnce sync.Once
	file_adhan_service_proto_rawDescData = file_adhan_service_proto_rawDesc
)

func file_adhan_service_proto_rawDescGZIP() []byte {
	file_adhan_service_proto_rawDescOnce.Do(func() {
		file_adhan_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_adhan_service_proto_rawDescData)
	})
	return file_adhan_service_proto_rawDescData
}

var file_adhan_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_adhan_service_proto_goTypes = []interface{}{
	(*AdhanFile)(nil),               // 0: limestone.AdhanFile
	(*CreateAdhanFileRequest)(nil),  // 1: limestone.CreateAdhanFileRequest
	(*UpdateAdhanFileRequest)(nil),  // 2: limestone.UpdateAdhanFileRequest
	(*GetAdhanFileRequest)(nil),     // 3: limestone.GetAdhanFileRequest
	(*DeleteAdhanFileRequest)(nil),  // 4: limestone.DeleteAdhanFileRequest
	(*DeleteAdhanFileResponse)(nil), // 5: limestone.DeleteAdhanFileResponse
	(*timestamppb.Timestamp)(nil),   // 6: google.protobuf.Timestamp
}
var file_adhan_service_proto_depIdxs = []int32{
	6, // 0: limestone.AdhanFile.create_time:type_name -> google.protobuf.Timestamp
	6, // 1: limestone.AdhanFile.update_time:type_name -> google.protobuf.Timestamp
	0, // 2: limestone.CreateAdhanFileRequest.adhan_file:type_name -> limestone.AdhanFile
	0, // 3: limestone.UpdateAdhanFileRequest.adhan_file:type_name -> limestone.AdhanFile
	1, // 4: limestone.AdhanService.CreateAdhanFile:input_type -> limestone.CreateAdhanFileRequest
	2, // 5: limestone.AdhanService.UpdateAdhanFile:input_type -> limestone.UpdateAdhanFileRequest
	3, // 6: limestone.AdhanService.GetAdhanFile:input_type -> limestone.GetAdhanFileRequest
	4, // 7: limestone.AdhanService.DeleteAdhanFile:input_type -> limestone.DeleteAdhanFileRequest
	0, // 8: limestone.AdhanService.CreateAdhanFile:output_type -> limestone.AdhanFile
	0, // 9: limestone.AdhanService.UpdateAdhanFile:output_type -> limestone.AdhanFile
	0, // 10: limestone.AdhanService.GetAdhanFile:output_type -> limestone.AdhanFile
	5, // 11: limestone.AdhanService.DeleteAdhanFile:output_type -> limestone.DeleteAdhanFileResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_adhan_service_proto_init() }
func file_adhan_service_proto_init() {
	if File_adhan_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_adhan_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdhanFile); i {
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
		file_adhan_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAdhanFileRequest); i {
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
		file_adhan_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAdhanFileRequest); i {
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
		file_adhan_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdhanFileRequest); i {
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
		file_adhan_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAdhanFileRequest); i {
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
		file_adhan_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAdhanFileResponse); i {
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
			RawDescriptor: file_adhan_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_adhan_service_proto_goTypes,
		DependencyIndexes: file_adhan_service_proto_depIdxs,
		MessageInfos:      file_adhan_service_proto_msgTypes,
	}.Build()
	File_adhan_service_proto = out.File
	file_adhan_service_proto_rawDesc = nil
	file_adhan_service_proto_goTypes = nil
	file_adhan_service_proto_depIdxs = nil
}
