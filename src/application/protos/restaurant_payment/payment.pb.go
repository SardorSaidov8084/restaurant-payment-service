// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: restaurant_payment/payment.proto

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

type Receipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId   string            `protobuf:"bytes,2,opt,name=order_id,proto3" json:"order_id,omitempty"`
	CardId    string            `protobuf:"bytes,3,opt,name=card_id,proto3" json:"card_id,omitempty"`
	Amount    int32             `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Data      map[string]string `protobuf:"bytes,5,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedAt string            `protobuf:"bytes,6,opt,name=created_at,proto3" json:"created_at,omitempty"`
	UpdatedAt string            `protobuf:"bytes,7,opt,name=updated_at,proto3" json:"updated_at,omitempty"`
}

func (x *Receipt) Reset() {
	*x = Receipt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_payment_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Receipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Receipt) ProtoMessage() {}

func (x *Receipt) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_payment_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Receipt.ProtoReflect.Descriptor instead.
func (*Receipt) Descriptor() ([]byte, []int) {
	return file_restaurant_payment_payment_proto_rawDescGZIP(), []int{0}
}

func (x *Receipt) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Receipt) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Receipt) GetCardId() string {
	if x != nil {
		return x.CardId
	}
	return ""
}

func (x *Receipt) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Receipt) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Receipt) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Receipt) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Currency   string            `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	Status     string            `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	CreateTime int32             `protobuf:"varint,4,opt,name=create_time,proto3" json:"create_time,omitempty"`
	PayTime    int32             `protobuf:"varint,5,opt,name=pay_time,proto3" json:"pay_time,omitempty"`
	CancelTime int32             `protobuf:"varint,6,opt,name=cancel_time,proto3" json:"cancel_time,omitempty"`
	Amount     int64             `protobuf:"varint,7,opt,name=amount,proto3" json:"amount,omitempty"`
	Account    map[string]string `protobuf:"bytes,8,rep,name=account,proto3" json:"account,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedAt  string            `protobuf:"bytes,9,opt,name=created_at,proto3" json:"created_at,omitempty"`
	UpdatedAt  string            `protobuf:"bytes,10,opt,name=updated_at,proto3" json:"updated_at,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_payment_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_payment_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_restaurant_payment_payment_proto_rawDescGZIP(), []int{1}
}

func (x *Transaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Transaction) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Transaction) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Transaction) GetCreateTime() int32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Transaction) GetPayTime() int32 {
	if x != nil {
		return x.PayTime
	}
	return 0
}

func (x *Transaction) GetCancelTime() int32 {
	if x != nil {
		return x.CancelTime
	}
	return 0
}

func (x *Transaction) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetAccount() map[string]string {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *Transaction) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Transaction) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GetReceiptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReceiptId string `protobuf:"bytes,1,opt,name=receipt_id,proto3" json:"receipt_id,omitempty"`
	OrderId   string `protobuf:"bytes,2,opt,name=order_id,proto3" json:"order_id,omitempty"`
}

func (x *GetReceiptRequest) Reset() {
	*x = GetReceiptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_payment_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceiptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceiptRequest) ProtoMessage() {}

func (x *GetReceiptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_payment_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReceiptRequest.ProtoReflect.Descriptor instead.
func (*GetReceiptRequest) Descriptor() ([]byte, []int) {
	return file_restaurant_payment_payment_proto_rawDescGZIP(), []int{2}
}

func (x *GetReceiptRequest) GetReceiptId() string {
	if x != nil {
		return x.ReceiptId
	}
	return ""
}

func (x *GetReceiptRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type GetReceiptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Receipt *Receipt `protobuf:"bytes,1,opt,name=receipt,proto3" json:"receipt,omitempty"`
}

func (x *GetReceiptResponse) Reset() {
	*x = GetReceiptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_payment_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceiptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceiptResponse) ProtoMessage() {}

func (x *GetReceiptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_payment_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReceiptResponse.ProtoReflect.Descriptor instead.
func (*GetReceiptResponse) Descriptor() ([]byte, []int) {
	return file_restaurant_payment_payment_proto_rawDescGZIP(), []int{3}
}

func (x *GetReceiptResponse) GetReceipt() *Receipt {
	if x != nil {
		return x.Receipt
	}
	return nil
}

type SaveTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *SaveTransactionRequest) Reset() {
	*x = SaveTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_payment_payment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveTransactionRequest) ProtoMessage() {}

func (x *SaveTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_payment_payment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveTransactionRequest.ProtoReflect.Descriptor instead.
func (*SaveTransactionRequest) Descriptor() ([]byte, []int) {
	return file_restaurant_payment_payment_proto_rawDescGZIP(), []int{4}
}

func (x *SaveTransactionRequest) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type SaveTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SaveTransactionResponse) Reset() {
	*x = SaveTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_payment_payment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveTransactionResponse) ProtoMessage() {}

func (x *SaveTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_payment_payment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveTransactionResponse.ProtoReflect.Descriptor instead.
func (*SaveTransactionResponse) Descriptor() ([]byte, []int) {
	return file_restaurant_payment_payment_proto_rawDescGZIP(), []int{5}
}

type UpdateTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *UpdateTransactionRequest) Reset() {
	*x = UpdateTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_payment_payment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTransactionRequest) ProtoMessage() {}

func (x *UpdateTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_payment_payment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTransactionRequest.ProtoReflect.Descriptor instead.
func (*UpdateTransactionRequest) Descriptor() ([]byte, []int) {
	return file_restaurant_payment_payment_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTransactionRequest) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

type UpdateTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTransactionResponse) Reset() {
	*x = UpdateTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_payment_payment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTransactionResponse) ProtoMessage() {}

func (x *UpdateTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_payment_payment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTransactionResponse.ProtoReflect.Descriptor instead.
func (*UpdateTransactionResponse) Descriptor() ([]byte, []int) {
	return file_restaurant_payment_payment_proto_rawDescGZIP(), []int{7}
}

type GetTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,proto3" json:"transaction_id,omitempty"`
}

func (x *GetTransactionRequest) Reset() {
	*x = GetTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_payment_payment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionRequest) ProtoMessage() {}

func (x *GetTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_payment_payment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionRequest) Descriptor() ([]byte, []int) {
	return file_restaurant_payment_payment_proto_rawDescGZIP(), []int{8}
}

func (x *GetTransactionRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type GetTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *GetTransactionResponse) Reset() {
	*x = GetTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_payment_payment_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionResponse) ProtoMessage() {}

func (x *GetTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_payment_payment_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionResponse.ProtoReflect.Descriptor instead.
func (*GetTransactionResponse) Descriptor() ([]byte, []int) {
	return file_restaurant_payment_payment_proto_rawDescGZIP(), []int{9}
}

func (x *GetTransactionResponse) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

var File_restaurant_payment_payment_proto protoreflect.FileDescriptor

var file_restaurant_payment_payment_proto_rawDesc = []byte{
	0x0a, 0x20, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x9b, 0x02, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x1a, 0x37, 0x0a, 0x09, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x8d, 0x03, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x46, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x1a, 0x3a, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x4f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x4b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x22, 0x5b, 0x0a, 0x16, 0x53, 0x61, 0x76, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x19, 0x0a, 0x17, 0x53, 0x61, 0x76, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5d, 0x0a, 0x18, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x65,
	0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1b, 0x0a, 0x19, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x22, 0x5b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x41, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x53, 0x61, 0x72, 0x64, 0x6f, 0x72, 0x53, 0x61, 0x69, 0x64, 0x6f, 0x76, 0x38,
	0x30, 0x38, 0x34, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_restaurant_payment_payment_proto_rawDescOnce sync.Once
	file_restaurant_payment_payment_proto_rawDescData = file_restaurant_payment_payment_proto_rawDesc
)

func file_restaurant_payment_payment_proto_rawDescGZIP() []byte {
	file_restaurant_payment_payment_proto_rawDescOnce.Do(func() {
		file_restaurant_payment_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_restaurant_payment_payment_proto_rawDescData)
	})
	return file_restaurant_payment_payment_proto_rawDescData
}

var file_restaurant_payment_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_restaurant_payment_payment_proto_goTypes = []interface{}{
	(*Receipt)(nil),                   // 0: restaurant_payment.Receipt
	(*Transaction)(nil),               // 1: restaurant_payment.Transaction
	(*GetReceiptRequest)(nil),         // 2: restaurant_payment.GetReceiptRequest
	(*GetReceiptResponse)(nil),        // 3: restaurant_payment.GetReceiptResponse
	(*SaveTransactionRequest)(nil),    // 4: restaurant_payment.SaveTransactionRequest
	(*SaveTransactionResponse)(nil),   // 5: restaurant_payment.SaveTransactionResponse
	(*UpdateTransactionRequest)(nil),  // 6: restaurant_payment.UpdateTransactionRequest
	(*UpdateTransactionResponse)(nil), // 7: restaurant_payment.UpdateTransactionResponse
	(*GetTransactionRequest)(nil),     // 8: restaurant_payment.GetTransactionRequest
	(*GetTransactionResponse)(nil),    // 9: restaurant_payment.GetTransactionResponse
	nil,                               // 10: restaurant_payment.Receipt.DataEntry
	nil,                               // 11: restaurant_payment.Transaction.AccountEntry
}
var file_restaurant_payment_payment_proto_depIdxs = []int32{
	10, // 0: restaurant_payment.Receipt.data:type_name -> restaurant_payment.Receipt.DataEntry
	11, // 1: restaurant_payment.Transaction.account:type_name -> restaurant_payment.Transaction.AccountEntry
	0,  // 2: restaurant_payment.GetReceiptResponse.receipt:type_name -> restaurant_payment.Receipt
	1,  // 3: restaurant_payment.SaveTransactionRequest.transaction:type_name -> restaurant_payment.Transaction
	1,  // 4: restaurant_payment.UpdateTransactionRequest.transaction:type_name -> restaurant_payment.Transaction
	1,  // 5: restaurant_payment.GetTransactionResponse.transaction:type_name -> restaurant_payment.Transaction
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_restaurant_payment_payment_proto_init() }
func file_restaurant_payment_payment_proto_init() {
	if File_restaurant_payment_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_restaurant_payment_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Receipt); i {
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
		file_restaurant_payment_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_restaurant_payment_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReceiptRequest); i {
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
		file_restaurant_payment_payment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReceiptResponse); i {
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
		file_restaurant_payment_payment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveTransactionRequest); i {
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
		file_restaurant_payment_payment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveTransactionResponse); i {
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
		file_restaurant_payment_payment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTransactionRequest); i {
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
		file_restaurant_payment_payment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTransactionResponse); i {
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
		file_restaurant_payment_payment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionRequest); i {
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
		file_restaurant_payment_payment_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionResponse); i {
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
			RawDescriptor: file_restaurant_payment_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_restaurant_payment_payment_proto_goTypes,
		DependencyIndexes: file_restaurant_payment_payment_proto_depIdxs,
		MessageInfos:      file_restaurant_payment_payment_proto_msgTypes,
	}.Build()
	File_restaurant_payment_payment_proto = out.File
	file_restaurant_payment_payment_proto_rawDesc = nil
	file_restaurant_payment_payment_proto_goTypes = nil
	file_restaurant_payment_payment_proto_depIdxs = nil
}
