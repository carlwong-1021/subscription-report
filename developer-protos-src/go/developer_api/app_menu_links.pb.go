// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.2
// source: app_menu_links.proto

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

type AppMenuLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationId    string            `protobuf:"bytes,1,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	NameTranslations map[string]string `protobuf:"bytes,2,rep,name=name_translations,json=nameTranslations,proto3" json:"name_translations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	UiHook           string            `protobuf:"bytes,3,opt,name=ui_hook,json=uiHook,proto3" json:"ui_hook,omitempty"`
	Position         string            `protobuf:"bytes,4,opt,name=position,proto3" json:"position,omitempty"`
	Priority         string            `protobuf:"bytes,5,opt,name=priority,proto3" json:"priority,omitempty"`
	IconUrl          string            `protobuf:"bytes,6,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url,omitempty"`
	// Deprecated: Do not use.
	AppRouteId string                 `protobuf:"bytes,7,opt,name=app_route_id,json=appRouteId,proto3" json:"app_route_id,omitempty"`
	Status     string                 `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt  *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	XId        string                 `protobuf:"bytes,12,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	AppRoute   map[string]string      `protobuf:"bytes,13,rep,name=app_route,json=appRoute,proto3" json:"app_route,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AppUrl     string                 `protobuf:"bytes,14,opt,name=app_url,json=appUrl,proto3" json:"app_url,omitempty"`
}

func (x *AppMenuLink) Reset() {
	*x = AppMenuLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_menu_links_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppMenuLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppMenuLink) ProtoMessage() {}

func (x *AppMenuLink) ProtoReflect() protoreflect.Message {
	mi := &file_app_menu_links_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppMenuLink.ProtoReflect.Descriptor instead.
func (*AppMenuLink) Descriptor() ([]byte, []int) {
	return file_app_menu_links_proto_rawDescGZIP(), []int{0}
}

func (x *AppMenuLink) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *AppMenuLink) GetNameTranslations() map[string]string {
	if x != nil {
		return x.NameTranslations
	}
	return nil
}

func (x *AppMenuLink) GetUiHook() string {
	if x != nil {
		return x.UiHook
	}
	return ""
}

func (x *AppMenuLink) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *AppMenuLink) GetPriority() string {
	if x != nil {
		return x.Priority
	}
	return ""
}

func (x *AppMenuLink) GetIconUrl() string {
	if x != nil {
		return x.IconUrl
	}
	return ""
}

// Deprecated: Do not use.
func (x *AppMenuLink) GetAppRouteId() string {
	if x != nil {
		return x.AppRouteId
	}
	return ""
}

func (x *AppMenuLink) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AppMenuLink) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *AppMenuLink) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *AppMenuLink) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *AppMenuLink) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *AppMenuLink) GetAppRoute() map[string]string {
	if x != nil {
		return x.AppRoute
	}
	return nil
}

func (x *AppMenuLink) GetAppUrl() string {
	if x != nil {
		return x.AppUrl
	}
	return ""
}

// Index
type AppMenuLinksIndexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PerformerId string `protobuf:"bytes,1,opt,name=performer_id,json=performerId,proto3" json:"performer_id,omitempty"`
	MerchantId  string `protobuf:"bytes,2,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	Page        int32  `protobuf:"varint,6,opt,name=page,proto3" json:"page,omitempty"`
	PerPage     int32  `protobuf:"varint,7,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
}

func (x *AppMenuLinksIndexRequest) Reset() {
	*x = AppMenuLinksIndexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_menu_links_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppMenuLinksIndexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppMenuLinksIndexRequest) ProtoMessage() {}

func (x *AppMenuLinksIndexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_menu_links_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppMenuLinksIndexRequest.ProtoReflect.Descriptor instead.
func (*AppMenuLinksIndexRequest) Descriptor() ([]byte, []int) {
	return file_app_menu_links_proto_rawDescGZIP(), []int{1}
}

func (x *AppMenuLinksIndexRequest) GetPerformerId() string {
	if x != nil {
		return x.PerformerId
	}
	return ""
}

func (x *AppMenuLinksIndexRequest) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *AppMenuLinksIndexRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *AppMenuLinksIndexRequest) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type AppMenuLinksIndexRepsonse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items      []*AppMenuLink   `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Pagination *PaginationReply `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *AppMenuLinksIndexRepsonse) Reset() {
	*x = AppMenuLinksIndexRepsonse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_menu_links_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppMenuLinksIndexRepsonse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppMenuLinksIndexRepsonse) ProtoMessage() {}

func (x *AppMenuLinksIndexRepsonse) ProtoReflect() protoreflect.Message {
	mi := &file_app_menu_links_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppMenuLinksIndexRepsonse.ProtoReflect.Descriptor instead.
func (*AppMenuLinksIndexRepsonse) Descriptor() ([]byte, []int) {
	return file_app_menu_links_proto_rawDescGZIP(), []int{2}
}

func (x *AppMenuLinksIndexRepsonse) GetItems() []*AppMenuLink {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *AppMenuLinksIndexRepsonse) GetPagination() *PaginationReply {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_app_menu_links_proto protoreflect.FileDescriptor

var file_app_menu_links_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x70, 0x5f, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65,
	0x72, 0x41, 0x70, 0x69, 0x1a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x05, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x4d,
	0x65, 0x6e, 0x75, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x5c,
	0x0a, 0x11, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x44, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x4d, 0x65, 0x6e, 0x75,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x6e, 0x61, 0x6d, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x69, 0x5f, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x69, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x24, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02,
	0x18, 0x01, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0f, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x44, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x4d, 0x65,
	0x6e, 0x75, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x61, 0x70, 0x70, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x61, 0x70, 0x70, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x70, 0x70, 0x55, 0x72, 0x6c, 0x1a, 0x43, 0x0a, 0x15, 0x4e, 0x61, 0x6d, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d,
	0x41, 0x70, 0x70, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8d, 0x01, 0x0a, 0x18, 0x41, 0x70,
	0x70, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65,
	0x72, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x19, 0x41, 0x70,
	0x70, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52,
	0x65, 0x70, 0x73, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x44,
	0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x6a, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x4d, 0x65,
	0x6e, 0x75, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x5a, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x26, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e,
	0x41, 0x70, 0x70, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x44, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x65, 0x72, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x4d, 0x65, 0x6e, 0x75, 0x4c,
	0x69, 0x6e, 0x6b, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x70, 0x73, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c,
	0x6f, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_menu_links_proto_rawDescOnce sync.Once
	file_app_menu_links_proto_rawDescData = file_app_menu_links_proto_rawDesc
)

func file_app_menu_links_proto_rawDescGZIP() []byte {
	file_app_menu_links_proto_rawDescOnce.Do(func() {
		file_app_menu_links_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_menu_links_proto_rawDescData)
	})
	return file_app_menu_links_proto_rawDescData
}

var file_app_menu_links_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_app_menu_links_proto_goTypes = []interface{}{
	(*AppMenuLink)(nil),               // 0: DeveloperApi.AppMenuLink
	(*AppMenuLinksIndexRequest)(nil),  // 1: DeveloperApi.AppMenuLinksIndexRequest
	(*AppMenuLinksIndexRepsonse)(nil), // 2: DeveloperApi.AppMenuLinksIndexRepsonse
	nil,                               // 3: DeveloperApi.AppMenuLink.NameTranslationsEntry
	nil,                               // 4: DeveloperApi.AppMenuLink.AppRouteEntry
	(*timestamppb.Timestamp)(nil),     // 5: google.protobuf.Timestamp
	(*PaginationReply)(nil),           // 6: DeveloperApi.PaginationReply
}
var file_app_menu_links_proto_depIdxs = []int32{
	3, // 0: DeveloperApi.AppMenuLink.name_translations:type_name -> DeveloperApi.AppMenuLink.NameTranslationsEntry
	5, // 1: DeveloperApi.AppMenuLink.created_at:type_name -> google.protobuf.Timestamp
	5, // 2: DeveloperApi.AppMenuLink.updated_at:type_name -> google.protobuf.Timestamp
	5, // 3: DeveloperApi.AppMenuLink.deleted_at:type_name -> google.protobuf.Timestamp
	4, // 4: DeveloperApi.AppMenuLink.app_route:type_name -> DeveloperApi.AppMenuLink.AppRouteEntry
	0, // 5: DeveloperApi.AppMenuLinksIndexRepsonse.items:type_name -> DeveloperApi.AppMenuLink
	6, // 6: DeveloperApi.AppMenuLinksIndexRepsonse.pagination:type_name -> DeveloperApi.PaginationReply
	1, // 7: DeveloperApi.AppMenuLinks.Index:input_type -> DeveloperApi.AppMenuLinksIndexRequest
	2, // 8: DeveloperApi.AppMenuLinks.Index:output_type -> DeveloperApi.AppMenuLinksIndexRepsonse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_app_menu_links_proto_init() }
func file_app_menu_links_proto_init() {
	if File_app_menu_links_proto != nil {
		return
	}
	file_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_app_menu_links_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppMenuLink); i {
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
		file_app_menu_links_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppMenuLinksIndexRequest); i {
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
		file_app_menu_links_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppMenuLinksIndexRepsonse); i {
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
			RawDescriptor: file_app_menu_links_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_menu_links_proto_goTypes,
		DependencyIndexes: file_app_menu_links_proto_depIdxs,
		MessageInfos:      file_app_menu_links_proto_msgTypes,
	}.Build()
	File_app_menu_links_proto = out.File
	file_app_menu_links_proto_rawDesc = nil
	file_app_menu_links_proto_goTypes = nil
	file_app_menu_links_proto_depIdxs = nil
}
