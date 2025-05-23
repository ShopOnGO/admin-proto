// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: product_variant.proto

package service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StockAction int32

const (
	StockAction_RESERVE StockAction = 0
	StockAction_RELEASE StockAction = 1
	StockAction_UPDATE  StockAction = 2
)

// Enum value maps for StockAction.
var (
	StockAction_name = map[int32]string{
		0: "RESERVE",
		1: "RELEASE",
		2: "UPDATE",
	}
	StockAction_value = map[string]int32{
		"RESERVE": 0,
		"RELEASE": 1,
		"UPDATE":  2,
	}
)

func (x StockAction) Enum() *StockAction {
	p := new(StockAction)
	*p = x
	return p
}

func (x StockAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StockAction) Descriptor() protoreflect.EnumDescriptor {
	return file_product_variant_proto_enumTypes[0].Descriptor()
}

func (StockAction) Type() protoreflect.EnumType {
	return &file_product_variant_proto_enumTypes[0]
}

func (x StockAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StockAction.Descriptor instead.
func (StockAction) EnumDescriptor() ([]byte, []int) {
	return file_product_variant_proto_rawDescGZIP(), []int{0}
}

type VariantRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Identifier:
	//
	//	*VariantRequest_Id
	//	*VariantRequest_Sku
	//	*VariantRequest_Barcode
	Identifier    isVariantRequest_Identifier `protobuf_oneof:"identifier"`
	Unscoped      bool                        `protobuf:"varint,4,opt,name=unscoped,proto3" json:"unscoped,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VariantRequest) Reset() {
	*x = VariantRequest{}
	mi := &file_product_variant_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VariantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariantRequest) ProtoMessage() {}

func (x *VariantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_variant_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariantRequest.ProtoReflect.Descriptor instead.
func (*VariantRequest) Descriptor() ([]byte, []int) {
	return file_product_variant_proto_rawDescGZIP(), []int{0}
}

func (x *VariantRequest) GetIdentifier() isVariantRequest_Identifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

func (x *VariantRequest) GetId() uint32 {
	if x != nil {
		if x, ok := x.Identifier.(*VariantRequest_Id); ok {
			return x.Id
		}
	}
	return 0
}

func (x *VariantRequest) GetSku() string {
	if x != nil {
		if x, ok := x.Identifier.(*VariantRequest_Sku); ok {
			return x.Sku
		}
	}
	return ""
}

func (x *VariantRequest) GetBarcode() string {
	if x != nil {
		if x, ok := x.Identifier.(*VariantRequest_Barcode); ok {
			return x.Barcode
		}
	}
	return ""
}

func (x *VariantRequest) GetUnscoped() bool {
	if x != nil {
		return x.Unscoped
	}
	return false
}

type isVariantRequest_Identifier interface {
	isVariantRequest_Identifier()
}

type VariantRequest_Id struct {
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3,oneof"`
}

type VariantRequest_Sku struct {
	Sku string `protobuf:"bytes,2,opt,name=sku,proto3,oneof"`
}

type VariantRequest_Barcode struct {
	Barcode string `protobuf:"bytes,3,opt,name=barcode,proto3,oneof"`
}

func (*VariantRequest_Id) isVariantRequest_Identifier() {}

func (*VariantRequest_Sku) isVariantRequest_Identifier() {}

func (*VariantRequest_Barcode) isVariantRequest_Identifier() {}

type VariantListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     uint32                 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ActiveOnly    bool                   `protobuf:"varint,2,opt,name=active_only,json=activeOnly,proto3" json:"active_only,omitempty"`
	PriceRange    *PriceRange            `protobuf:"bytes,3,opt,name=price_range,json=priceRange,proto3" json:"price_range,omitempty"`
	Limit         uint32                 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        uint32                 `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VariantListRequest) Reset() {
	*x = VariantListRequest{}
	mi := &file_product_variant_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VariantListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariantListRequest) ProtoMessage() {}

func (x *VariantListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_variant_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariantListRequest.ProtoReflect.Descriptor instead.
func (*VariantListRequest) Descriptor() ([]byte, []int) {
	return file_product_variant_proto_rawDescGZIP(), []int{1}
}

func (x *VariantListRequest) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *VariantListRequest) GetActiveOnly() bool {
	if x != nil {
		return x.ActiveOnly
	}
	return false
}

func (x *VariantListRequest) GetPriceRange() *PriceRange {
	if x != nil {
		return x.PriceRange
	}
	return nil
}

func (x *VariantListRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *VariantListRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type PriceRange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Min           uint32                 `protobuf:"varint,1,opt,name=min,proto3" json:"min,omitempty"`
	Max           uint32                 `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PriceRange) Reset() {
	*x = PriceRange{}
	mi := &file_product_variant_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PriceRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceRange) ProtoMessage() {}

func (x *PriceRange) ProtoReflect() protoreflect.Message {
	mi := &file_product_variant_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceRange.ProtoReflect.Descriptor instead.
func (*PriceRange) Descriptor() ([]byte, []int) {
	return file_product_variant_proto_rawDescGZIP(), []int{2}
}

func (x *PriceRange) GetMin() uint32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *PriceRange) GetMax() uint32 {
	if x != nil {
		return x.Max
	}
	return 0
}

type VariantListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Variants      []*ProductVariant      `protobuf:"bytes,1,rep,name=variants,proto3" json:"variants,omitempty"`
	TotalCount    uint32                 `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Error         *ErrorResponse         `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VariantListResponse) Reset() {
	*x = VariantListResponse{}
	mi := &file_product_variant_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VariantListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariantListResponse) ProtoMessage() {}

func (x *VariantListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_variant_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariantListResponse.ProtoReflect.Descriptor instead.
func (*VariantListResponse) Descriptor() ([]byte, []int) {
	return file_product_variant_proto_rawDescGZIP(), []int{3}
}

func (x *VariantListResponse) GetVariants() []*ProductVariant {
	if x != nil {
		return x.Variants
	}
	return nil
}

func (x *VariantListResponse) GetTotalCount() uint32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *VariantListResponse) GetError() *ErrorResponse {
	if x != nil {
		return x.Error
	}
	return nil
}

type StockRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	VariantId     uint32                 `protobuf:"varint,1,opt,name=variant_id,json=variantId,proto3" json:"variant_id,omitempty"`
	Quantity      uint32                 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Action        StockAction            `protobuf:"varint,3,opt,name=action,proto3,enum=proto.StockAction" json:"action,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StockRequest) Reset() {
	*x = StockRequest{}
	mi := &file_product_variant_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockRequest) ProtoMessage() {}

func (x *StockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_variant_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockRequest.ProtoReflect.Descriptor instead.
func (*StockRequest) Descriptor() ([]byte, []int) {
	return file_product_variant_proto_rawDescGZIP(), []int{4}
}

func (x *StockRequest) GetVariantId() uint32 {
	if x != nil {
		return x.VariantId
	}
	return 0
}

func (x *StockRequest) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *StockRequest) GetAction() StockAction {
	if x != nil {
		return x.Action
	}
	return StockAction_RESERVE
}

type DeleteVariantRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Unscoped      bool                   `protobuf:"varint,2,opt,name=unscoped,proto3" json:"unscoped,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteVariantRequest) Reset() {
	*x = DeleteVariantRequest{}
	mi := &file_product_variant_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteVariantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVariantRequest) ProtoMessage() {}

func (x *DeleteVariantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_variant_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVariantRequest.ProtoReflect.Descriptor instead.
func (*DeleteVariantRequest) Descriptor() ([]byte, []int) {
	return file_product_variant_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteVariantRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteVariantRequest) GetUnscoped() bool {
	if x != nil {
		return x.Unscoped
	}
	return false
}

type VariantResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Variant       *ProductVariant        `protobuf:"bytes,1,opt,name=variant,proto3" json:"variant,omitempty"`
	Error         *ErrorResponse         `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VariantResponse) Reset() {
	*x = VariantResponse{}
	mi := &file_product_variant_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VariantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariantResponse) ProtoMessage() {}

func (x *VariantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_variant_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariantResponse.ProtoReflect.Descriptor instead.
func (*VariantResponse) Descriptor() ([]byte, []int) {
	return file_product_variant_proto_rawDescGZIP(), []int{6}
}

func (x *VariantResponse) GetVariant() *ProductVariant {
	if x != nil {
		return x.Variant
	}
	return nil
}

func (x *VariantResponse) GetError() *ErrorResponse {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_product_variant_proto protoreflect.FileDescriptor

var file_product_variant_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c,
	0x0a, 0x0e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x1a, 0x0a, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x62, 0x61, 0x72, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x75, 0x6e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x42, 0x0c,
	0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0xb6, 0x01, 0x0a,
	0x12, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6f, 0x6e, 0x6c,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4f,
	0x6e, 0x6c, 0x79, 0x12, 0x32, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0a, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x30, 0x0a, 0x0a, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x22, 0x95, 0x01, 0x0a, 0x13, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x75, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x42, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x6e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x75, 0x6e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x22, 0x6e, 0x0a, 0x0f, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a,
	0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x2a,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0x33, 0x0a, 0x0b, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x53,
	0x45, 0x52, 0x56, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x4c, 0x45, 0x41, 0x53,
	0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x32,
	0x89, 0x03, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x3b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0b, 0x5a, 0x09, 0x2e,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_product_variant_proto_rawDescOnce sync.Once
	file_product_variant_proto_rawDescData []byte
)

func file_product_variant_proto_rawDescGZIP() []byte {
	file_product_variant_proto_rawDescOnce.Do(func() {
		file_product_variant_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_product_variant_proto_rawDesc), len(file_product_variant_proto_rawDesc)))
	})
	return file_product_variant_proto_rawDescData
}

var file_product_variant_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_product_variant_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_product_variant_proto_goTypes = []any{
	(StockAction)(0),             // 0: proto.StockAction
	(*VariantRequest)(nil),       // 1: proto.VariantRequest
	(*VariantListRequest)(nil),   // 2: proto.VariantListRequest
	(*PriceRange)(nil),           // 3: proto.PriceRange
	(*VariantListResponse)(nil),  // 4: proto.VariantListResponse
	(*StockRequest)(nil),         // 5: proto.StockRequest
	(*DeleteVariantRequest)(nil), // 6: proto.DeleteVariantRequest
	(*VariantResponse)(nil),      // 7: proto.VariantResponse
	(*ProductVariant)(nil),       // 8: proto.ProductVariant
	(*ErrorResponse)(nil),        // 9: proto.ErrorResponse
	(*Error)(nil),                // 10: proto.Error
}
var file_product_variant_proto_depIdxs = []int32{
	3,  // 0: proto.VariantListRequest.price_range:type_name -> proto.PriceRange
	8,  // 1: proto.VariantListResponse.variants:type_name -> proto.ProductVariant
	9,  // 2: proto.VariantListResponse.error:type_name -> proto.ErrorResponse
	0,  // 3: proto.StockRequest.action:type_name -> proto.StockAction
	8,  // 4: proto.VariantResponse.variant:type_name -> proto.ProductVariant
	9,  // 5: proto.VariantResponse.error:type_name -> proto.ErrorResponse
	8,  // 6: proto.ProductVariantService.CreateVariant:input_type -> proto.ProductVariant
	8,  // 7: proto.ProductVariantService.UpdateVariant:input_type -> proto.ProductVariant
	6,  // 8: proto.ProductVariantService.DeleteVariant:input_type -> proto.DeleteVariantRequest
	1,  // 9: proto.ProductVariantService.GetVariant:input_type -> proto.VariantRequest
	2,  // 10: proto.ProductVariantService.ListVariants:input_type -> proto.VariantListRequest
	5,  // 11: proto.ProductVariantService.ManageStock:input_type -> proto.StockRequest
	7,  // 12: proto.ProductVariantService.CreateVariant:output_type -> proto.VariantResponse
	7,  // 13: proto.ProductVariantService.UpdateVariant:output_type -> proto.VariantResponse
	10, // 14: proto.ProductVariantService.DeleteVariant:output_type -> proto.Error
	7,  // 15: proto.ProductVariantService.GetVariant:output_type -> proto.VariantResponse
	4,  // 16: proto.ProductVariantService.ListVariants:output_type -> proto.VariantListResponse
	10, // 17: proto.ProductVariantService.ManageStock:output_type -> proto.Error
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_product_variant_proto_init() }
func file_product_variant_proto_init() {
	if File_product_variant_proto != nil {
		return
	}
	file_entities_proto_init()
	file_product_variant_proto_msgTypes[0].OneofWrappers = []any{
		(*VariantRequest_Id)(nil),
		(*VariantRequest_Sku)(nil),
		(*VariantRequest_Barcode)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_product_variant_proto_rawDesc), len(file_product_variant_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_variant_proto_goTypes,
		DependencyIndexes: file_product_variant_proto_depIdxs,
		EnumInfos:         file_product_variant_proto_enumTypes,
		MessageInfos:      file_product_variant_proto_msgTypes,
	}.Build()
	File_product_variant_proto = out.File
	file_product_variant_proto_goTypes = nil
	file_product_variant_proto_depIdxs = nil
}
