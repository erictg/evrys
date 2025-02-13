// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: evrys.proto

package proto

import (
	pb "github.com/cloudevents/sdk-go/binding/format/protobuf/v2/pb"
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

type Order int32

const (
	Order_None          Order = 0
	Order_TimestampDesc Order = 1
	Order_TimestampAsc  Order = 2
)

// Enum value maps for Order.
var (
	Order_name = map[int32]string{
		0: "None",
		1: "TimestampDesc",
		2: "TimestampAsc",
	}
	Order_value = map[string]int32{
		"None":          0,
		"TimestampDesc": 1,
		"TimestampAsc":  2,
	}
)

func (x Order) Enum() *Order {
	p := new(Order)
	*p = x
	return p
}

func (x Order) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Order) Descriptor() protoreflect.EnumDescriptor {
	return file_evrys_proto_enumTypes[0].Descriptor()
}

func (Order) Type() protoreflect.EnumType {
	return &file_evrys_proto_enumTypes[0]
}

func (x Order) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Order.Descriptor instead.
func (Order) EnumDescriptor() ([]byte, []int) {
	return file_evrys_proto_rawDescGZIP(), []int{0}
}

type RecordEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RecordEventResponse) Reset() {
	*x = RecordEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evrys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordEventResponse) ProtoMessage() {}

func (x *RecordEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_evrys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordEventResponse.ProtoReflect.Descriptor instead.
func (*RecordEventResponse) Descriptor() ([]byte, []int) {
	return file_evrys_proto_rawDescGZIP(), []int{0}
}

type GetEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *GetEventRequest) Reset() {
	*x = GetEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evrys_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventRequest) ProtoMessage() {}

func (x *GetEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_evrys_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventRequest.ProtoReflect.Descriptor instead.
func (*GetEventRequest) Descriptor() ([]byte, []int) {
	return file_evrys_proto_rawDescGZIP(), []int{1}
}

func (x *GetEventRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type AllWithTraceId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId string `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
}

func (x *AllWithTraceId) Reset() {
	*x = AllWithTraceId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evrys_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllWithTraceId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllWithTraceId) ProtoMessage() {}

func (x *AllWithTraceId) ProtoReflect() protoreflect.Message {
	mi := &file_evrys_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllWithTraceId.ProtoReflect.Descriptor instead.
func (*AllWithTraceId) Descriptor() ([]byte, []int) {
	return file_evrys_proto_rawDescGZIP(), []int{2}
}

func (x *AllWithTraceId) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

type AllWithTraceIdAndAfterTimestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId   string                 `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *AllWithTraceIdAndAfterTimestamp) Reset() {
	*x = AllWithTraceIdAndAfterTimestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evrys_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllWithTraceIdAndAfterTimestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllWithTraceIdAndAfterTimestamp) ProtoMessage() {}

func (x *AllWithTraceIdAndAfterTimestamp) ProtoReflect() protoreflect.Message {
	mi := &file_evrys_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllWithTraceIdAndAfterTimestamp.ProtoReflect.Descriptor instead.
func (*AllWithTraceIdAndAfterTimestamp) Descriptor() ([]byte, []int) {
	return file_evrys_proto_rawDescGZIP(), []int{3}
}

func (x *AllWithTraceIdAndAfterTimestamp) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *AllWithTraceIdAndAfterTimestamp) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*Filter_AllWithTraceId
	//	*Filter_AllWithTraceIdAndAfterTimestamp
	Value isFilter_Value `protobuf_oneof:"value"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evrys_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_evrys_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_evrys_proto_rawDescGZIP(), []int{4}
}

func (m *Filter) GetValue() isFilter_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Filter) GetAllWithTraceId() *AllWithTraceId {
	if x, ok := x.GetValue().(*Filter_AllWithTraceId); ok {
		return x.AllWithTraceId
	}
	return nil
}

func (x *Filter) GetAllWithTraceIdAndAfterTimestamp() *AllWithTraceIdAndAfterTimestamp {
	if x, ok := x.GetValue().(*Filter_AllWithTraceIdAndAfterTimestamp); ok {
		return x.AllWithTraceIdAndAfterTimestamp
	}
	return nil
}

type isFilter_Value interface {
	isFilter_Value()
}

type Filter_AllWithTraceId struct {
	AllWithTraceId *AllWithTraceId `protobuf:"bytes,1,opt,name=all_with_trace_id,json=allWithTraceId,proto3,oneof"`
}

type Filter_AllWithTraceIdAndAfterTimestamp struct {
	AllWithTraceIdAndAfterTimestamp *AllWithTraceIdAndAfterTimestamp `protobuf:"bytes,2,opt,name=all_with_trace_id_and_after_timestamp,json=allWithTraceIdAndAfterTimestamp,proto3,oneof"`
}

func (*Filter_AllWithTraceId) isFilter_Value() {}

func (*Filter_AllWithTraceIdAndAfterTimestamp) isFilter_Value() {}

type SliceEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filters []*Filter `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty"`
	Order   Order     `protobuf:"varint,2,opt,name=order,proto3,enum=evrys.Order" json:"order,omitempty"`
}

func (x *SliceEventsRequest) Reset() {
	*x = SliceEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evrys_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SliceEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SliceEventsRequest) ProtoMessage() {}

func (x *SliceEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_evrys_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SliceEventsRequest.ProtoReflect.Descriptor instead.
func (*SliceEventsRequest) Descriptor() ([]byte, []int) {
	return file_evrys_proto_rawDescGZIP(), []int{5}
}

func (x *SliceEventsRequest) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *SliceEventsRequest) GetOrder() Order {
	if x != nil {
		return x.Order
	}
	return Order_None
}

var File_evrys_proto protoreflect.FileDescriptor

var file_evrys_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x76, 0x72, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65,
	0x76, 0x72, 0x79, 0x73, 0x1a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x64, 0x6b,
	0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x70,
	0x62, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x0e, 0x41, 0x6c, 0x6c, 0x57,
	0x69, 0x74, 0x68, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x76, 0x0a, 0x1f, 0x41, 0x6c, 0x6c, 0x57, 0x69, 0x74, 0x68,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x41, 0x66, 0x74, 0x65, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xcf, 0x01,
	0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x11, 0x61, 0x6c, 0x6c, 0x5f,
	0x77, 0x69, 0x74, 0x68, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x76, 0x72, 0x79, 0x73, 0x2e, 0x41, 0x6c, 0x6c, 0x57,
	0x69, 0x74, 0x68, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x48, 0x00, 0x52, 0x0e, 0x61, 0x6c,
	0x6c, 0x57, 0x69, 0x74, 0x68, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x78, 0x0a, 0x25,
	0x61, 0x6c, 0x6c, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x76,
	0x72, 0x79, 0x73, 0x2e, 0x41, 0x6c, 0x6c, 0x57, 0x69, 0x74, 0x68, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x41, 0x6e, 0x64, 0x41, 0x66, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x1f, 0x61, 0x6c, 0x6c, 0x57, 0x69, 0x74, 0x68, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x41, 0x66, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x61, 0x0a, 0x12, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65, 0x76, 0x72, 0x79, 0x73, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x22,
	0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e,
	0x65, 0x76, 0x72, 0x79, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2a, 0x36, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x44, 0x65, 0x73, 0x63, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x41, 0x73, 0x63, 0x10, 0x02, 0x32, 0xb2, 0x01, 0x0a, 0x05, 0x45,
	0x76, 0x72, 0x79, 0x73, 0x12, 0x39, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x76, 0x72, 0x79, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x65, 0x76,
	0x72, 0x79, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x19, 0x2e, 0x65, 0x76, 0x72, 0x79, 0x73, 0x2e, 0x53, 0x6c, 0x69, 0x63, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x30, 0x01, 0x42,
	0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x35,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x76, 0x72, 0x79, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_evrys_proto_rawDescOnce sync.Once
	file_evrys_proto_rawDescData = file_evrys_proto_rawDesc
)

func file_evrys_proto_rawDescGZIP() []byte {
	file_evrys_proto_rawDescOnce.Do(func() {
		file_evrys_proto_rawDescData = protoimpl.X.CompressGZIP(file_evrys_proto_rawDescData)
	})
	return file_evrys_proto_rawDescData
}

var file_evrys_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_evrys_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_evrys_proto_goTypes = []interface{}{
	(Order)(0),                              // 0: evrys.Order
	(*RecordEventResponse)(nil),             // 1: evrys.RecordEventResponse
	(*GetEventRequest)(nil),                 // 2: evrys.GetEventRequest
	(*AllWithTraceId)(nil),                  // 3: evrys.AllWithTraceId
	(*AllWithTraceIdAndAfterTimestamp)(nil), // 4: evrys.AllWithTraceIdAndAfterTimestamp
	(*Filter)(nil),                          // 5: evrys.Filter
	(*SliceEventsRequest)(nil),              // 6: evrys.SliceEventsRequest
	(*timestamppb.Timestamp)(nil),           // 7: google.protobuf.Timestamp
	(*pb.CloudEvent)(nil),                   // 8: pb.CloudEvent
}
var file_evrys_proto_depIdxs = []int32{
	7, // 0: evrys.AllWithTraceIdAndAfterTimestamp.timestamp:type_name -> google.protobuf.Timestamp
	3, // 1: evrys.Filter.all_with_trace_id:type_name -> evrys.AllWithTraceId
	4, // 2: evrys.Filter.all_with_trace_id_and_after_timestamp:type_name -> evrys.AllWithTraceIdAndAfterTimestamp
	5, // 3: evrys.SliceEventsRequest.filters:type_name -> evrys.Filter
	0, // 4: evrys.SliceEventsRequest.order:type_name -> evrys.Order
	8, // 5: evrys.Evrys.RecordEvent:input_type -> pb.CloudEvent
	2, // 6: evrys.Evrys.GetEvent:input_type -> evrys.GetEventRequest
	6, // 7: evrys.Evrys.SliceEvents:input_type -> evrys.SliceEventsRequest
	1, // 8: evrys.Evrys.RecordEvent:output_type -> evrys.RecordEventResponse
	8, // 9: evrys.Evrys.GetEvent:output_type -> pb.CloudEvent
	8, // 10: evrys.Evrys.SliceEvents:output_type -> pb.CloudEvent
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_evrys_proto_init() }
func file_evrys_proto_init() {
	if File_evrys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_evrys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordEventResponse); i {
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
		file_evrys_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventRequest); i {
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
		file_evrys_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllWithTraceId); i {
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
		file_evrys_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllWithTraceIdAndAfterTimestamp); i {
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
		file_evrys_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_evrys_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SliceEventsRequest); i {
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
	file_evrys_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Filter_AllWithTraceId)(nil),
		(*Filter_AllWithTraceIdAndAfterTimestamp)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_evrys_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_evrys_proto_goTypes,
		DependencyIndexes: file_evrys_proto_depIdxs,
		EnumInfos:         file_evrys_proto_enumTypes,
		MessageInfos:      file_evrys_proto_msgTypes,
	}.Build()
	File_evrys_proto = out.File
	file_evrys_proto_rawDesc = nil
	file_evrys_proto_goTypes = nil
	file_evrys_proto_depIdxs = nil
}
