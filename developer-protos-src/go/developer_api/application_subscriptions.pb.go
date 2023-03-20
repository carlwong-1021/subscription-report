// application_subscription.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.2
// source: application_subscriptions.proto

package developer_api

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

type ApplicationSubscriptionIndexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MerchantId    string `protobuf:"bytes,1,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	ApplicationId string `protobuf:"bytes,2,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
}

func (x *ApplicationSubscriptionIndexRequest) Reset() {
	*x = ApplicationSubscriptionIndexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_subscriptions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationSubscriptionIndexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationSubscriptionIndexRequest) ProtoMessage() {}

func (x *ApplicationSubscriptionIndexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_application_subscriptions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationSubscriptionIndexRequest.ProtoReflect.Descriptor instead.
func (*ApplicationSubscriptionIndexRequest) Descriptor() ([]byte, []int) {
	return file_application_subscriptions_proto_rawDescGZIP(), []int{0}
}

func (x *ApplicationSubscriptionIndexRequest) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *ApplicationSubscriptionIndexRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

type ApplicationSubscriptionShowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MerchantId    string `protobuf:"bytes,1,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	ApplicationId string `protobuf:"bytes,2,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
}

func (x *ApplicationSubscriptionShowRequest) Reset() {
	*x = ApplicationSubscriptionShowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_subscriptions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationSubscriptionShowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationSubscriptionShowRequest) ProtoMessage() {}

func (x *ApplicationSubscriptionShowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_application_subscriptions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationSubscriptionShowRequest.ProtoReflect.Descriptor instead.
func (*ApplicationSubscriptionShowRequest) Descriptor() ([]byte, []int) {
	return file_application_subscriptions_proto_rawDescGZIP(), []int{1}
}

func (x *ApplicationSubscriptionShowRequest) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *ApplicationSubscriptionShowRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

type ApplicationSubscriptionCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MerchantId         string `protobuf:"bytes,1,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	ApplicationId      string `protobuf:"bytes,2,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	SubscriptionStatus string `protobuf:"bytes,3,opt,name=subscription_status,json=subscriptionStatus,proto3" json:"subscription_status,omitempty"`
	PurchaseType       string `protobuf:"bytes,4,opt,name=purchase_type,json=purchaseType,proto3" json:"purchase_type,omitempty"`
}

func (x *ApplicationSubscriptionCreateRequest) Reset() {
	*x = ApplicationSubscriptionCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_subscriptions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationSubscriptionCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationSubscriptionCreateRequest) ProtoMessage() {}

func (x *ApplicationSubscriptionCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_application_subscriptions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationSubscriptionCreateRequest.ProtoReflect.Descriptor instead.
func (*ApplicationSubscriptionCreateRequest) Descriptor() ([]byte, []int) {
	return file_application_subscriptions_proto_rawDescGZIP(), []int{2}
}

func (x *ApplicationSubscriptionCreateRequest) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *ApplicationSubscriptionCreateRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *ApplicationSubscriptionCreateRequest) GetSubscriptionStatus() string {
	if x != nil {
		return x.SubscriptionStatus
	}
	return ""
}

func (x *ApplicationSubscriptionCreateRequest) GetPurchaseType() string {
	if x != nil {
		return x.PurchaseType
	}
	return ""
}

type TerminateTrialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MerchantId    string `protobuf:"bytes,1,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	ApplicationId string `protobuf:"bytes,2,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	PerformerId   string `protobuf:"bytes,3,opt,name=performer_id,json=performerId,proto3" json:"performer_id,omitempty"`
}

func (x *TerminateTrialRequest) Reset() {
	*x = TerminateTrialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_subscriptions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateTrialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateTrialRequest) ProtoMessage() {}

func (x *TerminateTrialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_application_subscriptions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminateTrialRequest.ProtoReflect.Descriptor instead.
func (*TerminateTrialRequest) Descriptor() ([]byte, []int) {
	return file_application_subscriptions_proto_rawDescGZIP(), []int{3}
}

func (x *TerminateTrialRequest) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *TerminateTrialRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *TerminateTrialRequest) GetPerformerId() string {
	if x != nil {
		return x.PerformerId
	}
	return ""
}

type ApplicationSubscriptionIndexReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items      []*ApplicationSubscription `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Pagination *PaginationReply           `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ApplicationSubscriptionIndexReply) Reset() {
	*x = ApplicationSubscriptionIndexReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_subscriptions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationSubscriptionIndexReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationSubscriptionIndexReply) ProtoMessage() {}

func (x *ApplicationSubscriptionIndexReply) ProtoReflect() protoreflect.Message {
	mi := &file_application_subscriptions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationSubscriptionIndexReply.ProtoReflect.Descriptor instead.
func (*ApplicationSubscriptionIndexReply) Descriptor() ([]byte, []int) {
	return file_application_subscriptions_proto_rawDescGZIP(), []int{4}
}

func (x *ApplicationSubscriptionIndexReply) GetItems() []*ApplicationSubscription {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ApplicationSubscriptionIndexReply) GetPagination() *PaginationReply {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ApplicationSubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XId                string                 `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	MerchantId         string                 `protobuf:"bytes,2,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	ApplicationId      string                 `protobuf:"bytes,3,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	SubscriptionStatus string                 `protobuf:"bytes,4,opt,name=subscription_status,json=subscriptionStatus,proto3" json:"subscription_status,omitempty"`
	PurchaseType       string                 `protobuf:"bytes,5,opt,name=purchase_type,json=purchaseType,proto3" json:"purchase_type,omitempty"`
	UpdatedAt          *timestamppb.Timestamp `protobuf:"bytes,99,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedAt          *timestamppb.Timestamp `protobuf:"bytes,100,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *ApplicationSubscription) Reset() {
	*x = ApplicationSubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_subscriptions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationSubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationSubscription) ProtoMessage() {}

func (x *ApplicationSubscription) ProtoReflect() protoreflect.Message {
	mi := &file_application_subscriptions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationSubscription.ProtoReflect.Descriptor instead.
func (*ApplicationSubscription) Descriptor() ([]byte, []int) {
	return file_application_subscriptions_proto_rawDescGZIP(), []int{5}
}

func (x *ApplicationSubscription) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *ApplicationSubscription) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *ApplicationSubscription) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *ApplicationSubscription) GetSubscriptionStatus() string {
	if x != nil {
		return x.SubscriptionStatus
	}
	return ""
}

func (x *ApplicationSubscription) GetPurchaseType() string {
	if x != nil {
		return x.PurchaseType
	}
	return ""
}

func (x *ApplicationSubscription) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ApplicationSubscription) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type ApplicationSubscriptionEmptyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ApplicationSubscriptionEmptyReply) Reset() {
	*x = ApplicationSubscriptionEmptyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_application_subscriptions_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationSubscriptionEmptyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationSubscriptionEmptyReply) ProtoMessage() {}

func (x *ApplicationSubscriptionEmptyReply) ProtoReflect() protoreflect.Message {
	mi := &file_application_subscriptions_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationSubscriptionEmptyReply.ProtoReflect.Descriptor instead.
func (*ApplicationSubscriptionEmptyReply) Descriptor() ([]byte, []int) {
	return file_application_subscriptions_proto_rawDescGZIP(), []int{6}
}

var File_application_subscriptions_proto protoreflect.FileDescriptor

var file_application_subscriptions_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x41, 0x70, 0x69, 0x1a,
	0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x23, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x6c, 0x0a, 0x22, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x68, 0x6f, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0xc4, 0x01, 0x0a, 0x24, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x2f, 0x0a, 0x13, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x15, 0x54, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x65, 0x72, 0x66,
	0x6f, 0x72, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x9f, 0x01, 0x0a, 0x21,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x3b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x3d,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x41, 0x70,
	0x69, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbe, 0x02,
	0x0a, 0x17, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x03, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x23,
	0x0a, 0x21, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x32, 0xa4, 0x04, 0x0a, 0x18, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x6d, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x44,
	0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x61, 0x0a, 0x04, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x30, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x68,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x44, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x12, 0x65, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x32, 0x2e, 0x44,
	0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x32, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x41,
	0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00,
	0x12, 0x68, 0x0a, 0x0e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69,
	0x61, 0x6c, 0x12, 0x23, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x41, 0x70,
	0x69, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x61, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x64, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67,
	0x6f, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_application_subscriptions_proto_rawDescOnce sync.Once
	file_application_subscriptions_proto_rawDescData = file_application_subscriptions_proto_rawDesc
)

func file_application_subscriptions_proto_rawDescGZIP() []byte {
	file_application_subscriptions_proto_rawDescOnce.Do(func() {
		file_application_subscriptions_proto_rawDescData = protoimpl.X.CompressGZIP(file_application_subscriptions_proto_rawDescData)
	})
	return file_application_subscriptions_proto_rawDescData
}

var file_application_subscriptions_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_application_subscriptions_proto_goTypes = []interface{}{
	(*ApplicationSubscriptionIndexRequest)(nil),  // 0: DeveloperApi.ApplicationSubscriptionIndexRequest
	(*ApplicationSubscriptionShowRequest)(nil),   // 1: DeveloperApi.ApplicationSubscriptionShowRequest
	(*ApplicationSubscriptionCreateRequest)(nil), // 2: DeveloperApi.ApplicationSubscriptionCreateRequest
	(*TerminateTrialRequest)(nil),                // 3: DeveloperApi.TerminateTrialRequest
	(*ApplicationSubscriptionIndexReply)(nil),    // 4: DeveloperApi.ApplicationSubscriptionIndexReply
	(*ApplicationSubscription)(nil),              // 5: DeveloperApi.ApplicationSubscription
	(*ApplicationSubscriptionEmptyReply)(nil),    // 6: DeveloperApi.ApplicationSubscriptionEmptyReply
	(*PaginationReply)(nil),                      // 7: DeveloperApi.PaginationReply
	(*timestamppb.Timestamp)(nil),                // 8: google.protobuf.Timestamp
}
var file_application_subscriptions_proto_depIdxs = []int32{
	5, // 0: DeveloperApi.ApplicationSubscriptionIndexReply.items:type_name -> DeveloperApi.ApplicationSubscription
	7, // 1: DeveloperApi.ApplicationSubscriptionIndexReply.pagination:type_name -> DeveloperApi.PaginationReply
	8, // 2: DeveloperApi.ApplicationSubscription.updated_at:type_name -> google.protobuf.Timestamp
	8, // 3: DeveloperApi.ApplicationSubscription.created_at:type_name -> google.protobuf.Timestamp
	0, // 4: DeveloperApi.ApplicationSubscriptions.Index:input_type -> DeveloperApi.ApplicationSubscriptionIndexRequest
	1, // 5: DeveloperApi.ApplicationSubscriptions.Show:input_type -> DeveloperApi.ApplicationSubscriptionShowRequest
	2, // 6: DeveloperApi.ApplicationSubscriptions.Create:input_type -> DeveloperApi.ApplicationSubscriptionCreateRequest
	2, // 7: DeveloperApi.ApplicationSubscriptions.Update:input_type -> DeveloperApi.ApplicationSubscriptionCreateRequest
	3, // 8: DeveloperApi.ApplicationSubscriptions.TerminateTrial:input_type -> DeveloperApi.TerminateTrialRequest
	4, // 9: DeveloperApi.ApplicationSubscriptions.Index:output_type -> DeveloperApi.ApplicationSubscriptionIndexReply
	5, // 10: DeveloperApi.ApplicationSubscriptions.Show:output_type -> DeveloperApi.ApplicationSubscription
	5, // 11: DeveloperApi.ApplicationSubscriptions.Create:output_type -> DeveloperApi.ApplicationSubscription
	5, // 12: DeveloperApi.ApplicationSubscriptions.Update:output_type -> DeveloperApi.ApplicationSubscription
	6, // 13: DeveloperApi.ApplicationSubscriptions.TerminateTrial:output_type -> DeveloperApi.ApplicationSubscriptionEmptyReply
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_application_subscriptions_proto_init() }
func file_application_subscriptions_proto_init() {
	if File_application_subscriptions_proto != nil {
		return
	}
	file_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_application_subscriptions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationSubscriptionIndexRequest); i {
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
		file_application_subscriptions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationSubscriptionShowRequest); i {
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
		file_application_subscriptions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationSubscriptionCreateRequest); i {
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
		file_application_subscriptions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminateTrialRequest); i {
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
		file_application_subscriptions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationSubscriptionIndexReply); i {
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
		file_application_subscriptions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationSubscription); i {
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
		file_application_subscriptions_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationSubscriptionEmptyReply); i {
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
			RawDescriptor: file_application_subscriptions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_application_subscriptions_proto_goTypes,
		DependencyIndexes: file_application_subscriptions_proto_depIdxs,
		MessageInfos:      file_application_subscriptions_proto_msgTypes,
	}.Build()
	File_application_subscriptions_proto = out.File
	file_application_subscriptions_proto_rawDesc = nil
	file_application_subscriptions_proto_goTypes = nil
	file_application_subscriptions_proto_depIdxs = nil
}
