// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: pkg/model/order/order.proto

package order

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

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	ItemIds []string `protobuf:"bytes,2,rep,name=item_ids,json=itemIds,proto3" json:"item_ids,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_order_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_order_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_pkg_model_order_order_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderRequest) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *CreateOrderRequest) GetItemIds() []string {
	if x != nil {
		return x.ItemIds
	}
	return nil
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_order_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_order_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_pkg_model_order_order_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type GetOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_order_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_order_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_pkg_model_order_order_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type GetRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	RecordType string `protobuf:"bytes,2,opt,name=record_type,json=recordType,proto3" json:"record_type,omitempty"`
}

func (x *GetRecordRequest) Reset() {
	*x = GetRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_order_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordRequest) ProtoMessage() {}

func (x *GetRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_order_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordRequest.ProtoReflect.Descriptor instead.
func (*GetRecordRequest) Descriptor() ([]byte, []int) {
	return file_pkg_model_order_order_proto_rawDescGZIP(), []int{3}
}

func (x *GetRecordRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *GetRecordRequest) GetRecordType() string {
	if x != nil {
		return x.RecordType
	}
	return ""
}

type GetOrderReponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address     *Address               `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	ItemIds     []string               `protobuf:"bytes,2,rep,name=item_ids,json=itemIds,proto3" json:"item_ids,omitempty"`
	OrderDate   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=order_date,json=orderDate,proto3" json:"order_date,omitempty"`
	OrderStatus string                 `protobuf:"bytes,4,opt,name=order_status,json=orderStatus,proto3" json:"order_status,omitempty"`
}

func (x *GetOrderReponse) Reset() {
	*x = GetOrderReponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_order_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderReponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderReponse) ProtoMessage() {}

func (x *GetOrderReponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_order_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderReponse.ProtoReflect.Descriptor instead.
func (*GetOrderReponse) Descriptor() ([]byte, []int) {
	return file_pkg_model_order_order_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrderReponse) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *GetOrderReponse) GetItemIds() []string {
	if x != nil {
		return x.ItemIds
	}
	return nil
}

func (x *GetOrderReponse) GetOrderDate() *timestamppb.Timestamp {
	if x != nil {
		return x.OrderDate
	}
	return nil
}

func (x *GetOrderReponse) GetOrderStatus() string {
	if x != nil {
		return x.OrderStatus
	}
	return ""
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HouseNameOrNumber string `protobuf:"bytes,1,opt,name=house_name_or_number,json=houseNameOrNumber,proto3" json:"house_name_or_number,omitempty"`
	Street            string `protobuf:"bytes,2,opt,name=street,proto3" json:"street,omitempty"`
	Town              string `protobuf:"bytes,3,opt,name=town,proto3" json:"town,omitempty"`
	Postcode          string `protobuf:"bytes,4,opt,name=postcode,proto3" json:"postcode,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_model_order_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_model_order_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_pkg_model_order_order_proto_rawDescGZIP(), []int{5}
}

func (x *Address) GetHouseNameOrNumber() string {
	if x != nil {
		return x.HouseNameOrNumber
	}
	return ""
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetTown() string {
	if x != nil {
		return x.Town
	}
	return ""
}

func (x *Address) GetPostcode() string {
	if x != nil {
		return x.Postcode
	}
	return ""
}

var File_pkg_model_order_order_proto protoreflect.FileDescriptor

var file_pkg_model_order_order_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x73,
	0x22, 0x30, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x4e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x22, 0xb4, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x2f, 0x0a, 0x14, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x6f, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x72, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x6f, 0x77, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f, 0x77, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x08, 0x5a, 0x06,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_model_order_order_proto_rawDescOnce sync.Once
	file_pkg_model_order_order_proto_rawDescData = file_pkg_model_order_order_proto_rawDesc
)

func file_pkg_model_order_order_proto_rawDescGZIP() []byte {
	file_pkg_model_order_order_proto_rawDescOnce.Do(func() {
		file_pkg_model_order_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_model_order_order_proto_rawDescData)
	})
	return file_pkg_model_order_order_proto_rawDescData
}

var file_pkg_model_order_order_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_model_order_order_proto_goTypes = []interface{}{
	(*CreateOrderRequest)(nil),    // 0: order.CreateOrderRequest
	(*CreateOrderResponse)(nil),   // 1: order.CreateOrderResponse
	(*GetOrderRequest)(nil),       // 2: order.GetOrderRequest
	(*GetRecordRequest)(nil),      // 3: order.GetRecordRequest
	(*GetOrderReponse)(nil),       // 4: order.GetOrderReponse
	(*Address)(nil),               // 5: order.Address
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_pkg_model_order_order_proto_depIdxs = []int32{
	5, // 0: order.CreateOrderRequest.address:type_name -> order.Address
	5, // 1: order.GetOrderReponse.address:type_name -> order.Address
	6, // 2: order.GetOrderReponse.order_date:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_model_order_order_proto_init() }
func file_pkg_model_order_order_proto_init() {
	if File_pkg_model_order_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_model_order_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
		file_pkg_model_order_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderResponse); i {
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
		file_pkg_model_order_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderRequest); i {
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
		file_pkg_model_order_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecordRequest); i {
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
		file_pkg_model_order_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderReponse); i {
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
		file_pkg_model_order_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
			RawDescriptor: file_pkg_model_order_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_model_order_order_proto_goTypes,
		DependencyIndexes: file_pkg_model_order_order_proto_depIdxs,
		MessageInfos:      file_pkg_model_order_order_proto_msgTypes,
	}.Build()
	File_pkg_model_order_order_proto = out.File
	file_pkg_model_order_order_proto_rawDesc = nil
	file_pkg_model_order_order_proto_goTypes = nil
	file_pkg_model_order_order_proto_depIdxs = nil
}
