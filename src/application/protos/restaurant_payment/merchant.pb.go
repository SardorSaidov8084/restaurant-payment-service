// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: merchant.proto

package restaurant_payment

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

type MerchantSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId  string `protobuf:"bytes,1,opt,name=entity_id,proto3" json:"entity_id,omitempty"`
	CashboxId string `protobuf:"bytes,2,opt,name=cashbox_id,proto3" json:"cashbox_id,omitempty"`
	Enabled   bool   `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	CreatedAt string `protobuf:"bytes,8,opt,name=created_at,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,9,opt,name=updated_at,proto3" json:"updated_at,omitempty"`
}

func (x *MerchantSetting) Reset() {
	*x = MerchantSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerchantSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerchantSetting) ProtoMessage() {}

func (x *MerchantSetting) ProtoReflect() protoreflect.Message {
	mi := &file_merchant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerchantSetting.ProtoReflect.Descriptor instead.
func (*MerchantSetting) Descriptor() ([]byte, []int) {
	return file_merchant_proto_rawDescGZIP(), []int{0}
}

func (x *MerchantSetting) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *MerchantSetting) GetCashboxId() string {
	if x != nil {
		return x.CashboxId
	}
	return ""
}

func (x *MerchantSetting) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *MerchantSetting) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *MerchantSetting) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UpdateMerchantSettingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId  string `protobuf:"bytes,1,opt,name=entity_id,proto3" json:"entity_id,omitempty"`
	CashboxId string `protobuf:"bytes,2,opt,name=cashbox_id,proto3" json:"cashbox_id,omitempty"`
	Enabled   bool   `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *UpdateMerchantSettingRequest) Reset() {
	*x = UpdateMerchantSettingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMerchantSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMerchantSettingRequest) ProtoMessage() {}

func (x *UpdateMerchantSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_merchant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMerchantSettingRequest.ProtoReflect.Descriptor instead.
func (*UpdateMerchantSettingRequest) Descriptor() ([]byte, []int) {
	return file_merchant_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateMerchantSettingRequest) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *UpdateMerchantSettingRequest) GetCashboxId() string {
	if x != nil {
		return x.CashboxId
	}
	return ""
}

func (x *UpdateMerchantSettingRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type UpdateMerchantSettingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateMerchantSettingResponse) Reset() {
	*x = UpdateMerchantSettingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMerchantSettingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMerchantSettingResponse) ProtoMessage() {}

func (x *UpdateMerchantSettingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_merchant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMerchantSettingResponse.ProtoReflect.Descriptor instead.
func (*UpdateMerchantSettingResponse) Descriptor() ([]byte, []int) {
	return file_merchant_proto_rawDescGZIP(), []int{2}
}

type GetMerchantSettingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId string `protobuf:"bytes,1,opt,name=entity_id,proto3" json:"entity_id,omitempty"`
}

func (x *GetMerchantSettingRequest) Reset() {
	*x = GetMerchantSettingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMerchantSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMerchantSettingRequest) ProtoMessage() {}

func (x *GetMerchantSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_merchant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMerchantSettingRequest.ProtoReflect.Descriptor instead.
func (*GetMerchantSettingRequest) Descriptor() ([]byte, []int) {
	return file_merchant_proto_rawDescGZIP(), []int{3}
}

func (x *GetMerchantSettingRequest) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

type GetMerchantSettingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Setting *MerchantSetting `protobuf:"bytes,1,opt,name=setting,proto3" json:"setting,omitempty"`
}

func (x *GetMerchantSettingResponse) Reset() {
	*x = GetMerchantSettingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merchant_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMerchantSettingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMerchantSettingResponse) ProtoMessage() {}

func (x *GetMerchantSettingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_merchant_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMerchantSettingResponse.ProtoReflect.Descriptor instead.
func (*GetMerchantSettingResponse) Descriptor() ([]byte, []int) {
	return file_merchant_proto_rawDescGZIP(), []int{4}
}

func (x *GetMerchantSettingResponse) GetSetting() *MerchantSetting {
	if x != nil {
		return x.Setting
	}
	return nil
}

var File_merchant_proto protoreflect.FileDescriptor

var file_merchant_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa9, 0x01, 0x0a, 0x0f, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x78, 0x5f,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x76, 0x0a, 0x1c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x22, 0x1f, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x22, 0x48, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x07, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x07, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x16, 0x5a, 0x14, 0x2e, 0x2f,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_merchant_proto_rawDescOnce sync.Once
	file_merchant_proto_rawDescData = file_merchant_proto_rawDesc
)

func file_merchant_proto_rawDescGZIP() []byte {
	file_merchant_proto_rawDescOnce.Do(func() {
		file_merchant_proto_rawDescData = protoimpl.X.CompressGZIP(file_merchant_proto_rawDescData)
	})
	return file_merchant_proto_rawDescData
}

var file_merchant_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_merchant_proto_goTypes = []interface{}{
	(*MerchantSetting)(nil),               // 0: MerchantSetting
	(*UpdateMerchantSettingRequest)(nil),  // 1: UpdateMerchantSettingRequest
	(*UpdateMerchantSettingResponse)(nil), // 2: UpdateMerchantSettingResponse
	(*GetMerchantSettingRequest)(nil),     // 3: GetMerchantSettingRequest
	(*GetMerchantSettingResponse)(nil),    // 4: GetMerchantSettingResponse
}
var file_merchant_proto_depIdxs = []int32{
	0, // 0: GetMerchantSettingResponse.setting:type_name -> MerchantSetting
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_merchant_proto_init() }
func file_merchant_proto_init() {
	if File_merchant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_merchant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerchantSetting); i {
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
		file_merchant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMerchantSettingRequest); i {
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
		file_merchant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMerchantSettingResponse); i {
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
		file_merchant_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMerchantSettingRequest); i {
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
		file_merchant_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMerchantSettingResponse); i {
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
			RawDescriptor: file_merchant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_merchant_proto_goTypes,
		DependencyIndexes: file_merchant_proto_depIdxs,
		MessageInfos:      file_merchant_proto_msgTypes,
	}.Build()
	File_merchant_proto = out.File
	file_merchant_proto_rawDesc = nil
	file_merchant_proto_goTypes = nil
	file_merchant_proto_depIdxs = nil
}
