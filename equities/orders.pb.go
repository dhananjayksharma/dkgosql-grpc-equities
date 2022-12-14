// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: orders.proto

package equities

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

type OrderType int32

const (
	OrderType_ORDER_UNSPECIFIED OrderType = 0
	OrderType_ORDER_SELL        OrderType = 1
	OrderType_ORDER_BUY         OrderType = 2
)

// Enum value maps for OrderType.
var (
	OrderType_name = map[int32]string{
		0: "ORDER_UNSPECIFIED",
		1: "ORDER_SELL",
		2: "ORDER_BUY",
	}
	OrderType_value = map[string]int32{
		"ORDER_UNSPECIFIED": 0,
		"ORDER_SELL":        1,
		"ORDER_BUY":         2,
	}
)

func (x OrderType) Enum() *OrderType {
	p := new(OrderType)
	*p = x
	return p
}

func (x OrderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderType) Descriptor() protoreflect.EnumDescriptor {
	return file_orders_proto_enumTypes[0].Descriptor()
}

func (OrderType) Type() protoreflect.EnumType {
	return &file_orders_proto_enumTypes[0]
}

func (x OrderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderType.Descriptor instead.
func (OrderType) EnumDescriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{0}
}

type OrderStatus int32

const (
	OrderStatus_ORDER_PENDING   OrderStatus = 0
	OrderStatus_ORDER_PROCESSED OrderStatus = 1
	OrderStatus_ORDER_CANCELED  OrderStatus = 2
	OrderStatus_ORDER_EXPIRED   OrderStatus = 3
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "ORDER_PENDING",
		1: "ORDER_PROCESSED",
		2: "ORDER_CANCELED",
		3: "ORDER_EXPIRED",
	}
	OrderStatus_value = map[string]int32{
		"ORDER_PENDING":   0,
		"ORDER_PROCESSED": 1,
		"ORDER_CANCELED":  2,
		"ORDER_EXPIRED":   3,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_orders_proto_enumTypes[1].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_orders_proto_enumTypes[1]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_orders_proto_rawDescGZIP(), []int{1}
}

var File_orders_proto protoreflect.FileDescriptor

var file_orders_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x65, 0x71, 0x75, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2a, 0x41, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x55, 0x59, 0x10, 0x02, 0x2a, 0x5c, 0x0a, 0x0b, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x52,
	0x44, 0x45, 0x52, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x13, 0x0a,
	0x0f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x43, 0x41, 0x4e, 0x43,
	0x45, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f,
	0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x03, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x68, 0x61, 0x6e, 0x61, 0x6e, 0x6a, 0x61,
	0x79, 0x6b, 0x73, 0x68, 0x61, 0x72, 0x6d, 0x61, 0x2f, 0x64, 0x6b, 0x67, 0x6f, 0x73, 0x71, 0x6c,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x71, 0x75, 0x69, 0x74, 0x69, 0x65, 0x73, 0x3b, 0x65,
	0x71, 0x75, 0x69, 0x74, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orders_proto_rawDescOnce sync.Once
	file_orders_proto_rawDescData = file_orders_proto_rawDesc
)

func file_orders_proto_rawDescGZIP() []byte {
	file_orders_proto_rawDescOnce.Do(func() {
		file_orders_proto_rawDescData = protoimpl.X.CompressGZIP(file_orders_proto_rawDescData)
	})
	return file_orders_proto_rawDescData
}

var file_orders_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_orders_proto_goTypes = []interface{}{
	(OrderType)(0),   // 0: equities.OrderType
	(OrderStatus)(0), // 1: equities.OrderStatus
}
var file_orders_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orders_proto_init() }
func file_orders_proto_init() {
	if File_orders_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orders_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orders_proto_goTypes,
		DependencyIndexes: file_orders_proto_depIdxs,
		EnumInfos:         file_orders_proto_enumTypes,
	}.Build()
	File_orders_proto = out.File
	file_orders_proto_rawDesc = nil
	file_orders_proto_goTypes = nil
	file_orders_proto_depIdxs = nil
}
