// Copyright 2021 Lightbend Inc.

// Extension for specifying which topics a gRPC endpoint should be connected
// to, in order to facilitate consuming and producing events from a message broker.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: eventing.proto

package akkaserverless

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Eventing configuration for a gRPC method.
type Eventing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The event source in configuration.
	In *EventSource `protobuf:"bytes,1,opt,name=in,proto3" json:"in,omitempty"`
	// The event destination out configuration.
	//
	// Optional, if unset, messages out will not be published anywhere.
	Out *EventDestination `protobuf:"bytes,2,opt,name=out,proto3" json:"out,omitempty"`
}

func (x *Eventing) Reset() {
	*x = Eventing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Eventing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Eventing) ProtoMessage() {}

func (x *Eventing) ProtoReflect() protoreflect.Message {
	mi := &file_eventing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Eventing.ProtoReflect.Descriptor instead.
func (*Eventing) Descriptor() ([]byte, []int) {
	return file_eventing_proto_rawDescGZIP(), []int{0}
}

func (x *Eventing) GetIn() *EventSource {
	if x != nil {
		return x.In
	}
	return nil
}

func (x *Eventing) GetOut() *EventDestination {
	if x != nil {
		return x.Out
	}
	return nil
}

// Event source configuration
type EventSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The consumer group id.
	//
	// By default, all rpc methods on a given service with the same source will be
	// part of the same virtual consumer group, messages will be routed to the
	// different methods by type. This can be used to override that. If you want
	// multiple methods to act as independent consumers of the same source (ie, if
	// you want the same event to be published to each consumer) then give each
	// consumer a unique name.
	//
	// Note that this does depend on the event source supporting multiple consumer
	// groups. Queue based event sources may not support this.
	ConsumerGroup string `protobuf:"bytes,1,opt,name=consumer_group,json=consumerGroup,proto3" json:"consumer_group,omitempty"`
	// Types that are assignable to Source:
	//	*EventSource_Topic
	//	*EventSource_EventSourcedEntity
	//	*EventSource_ValueEntity
	Source isEventSource_Source `protobuf_oneof:"source"`
}

func (x *EventSource) Reset() {
	*x = EventSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventSource) ProtoMessage() {}

func (x *EventSource) ProtoReflect() protoreflect.Message {
	mi := &file_eventing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventSource.ProtoReflect.Descriptor instead.
func (*EventSource) Descriptor() ([]byte, []int) {
	return file_eventing_proto_rawDescGZIP(), []int{1}
}

func (x *EventSource) GetConsumerGroup() string {
	if x != nil {
		return x.ConsumerGroup
	}
	return ""
}

func (m *EventSource) GetSource() isEventSource_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *EventSource) GetTopic() string {
	if x, ok := x.GetSource().(*EventSource_Topic); ok {
		return x.Topic
	}
	return ""
}

func (x *EventSource) GetEventSourcedEntity() string {
	if x, ok := x.GetSource().(*EventSource_EventSourcedEntity); ok {
		return x.EventSourcedEntity
	}
	return ""
}

func (x *EventSource) GetValueEntity() string {
	if x, ok := x.GetSource().(*EventSource_ValueEntity); ok {
		return x.ValueEntity
	}
	return ""
}

type isEventSource_Source interface {
	isEventSource_Source()
}

type EventSource_Topic struct {
	// A topic source.
	//
	// This will consume events from the given topic name.
	Topic string `protobuf:"bytes,2,opt,name=topic,proto3,oneof"`
}

type EventSource_EventSourcedEntity struct {
	// Source for events emitted by an Event Sourced Entity.
	//
	// This will consume events from the Event Sourced Entity with the given entity type name.
	EventSourcedEntity string `protobuf:"bytes,3,opt,name=event_sourced_entity,json=eventSourcedEntity,proto3,oneof"`
}

type EventSource_ValueEntity struct {
	// Source for updates issued by a Value Entity.
	//
	// This will consume changes to the Value Entity with the given entity type name.
	ValueEntity string `protobuf:"bytes,4,opt,name=value_entity,json=valueEntity,proto3,oneof"`
}

func (*EventSource_Topic) isEventSource_Source() {}

func (*EventSource_EventSourcedEntity) isEventSource_Source() {}

func (*EventSource_ValueEntity) isEventSource_Source() {}

type EventDestination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Destination:
	//	*EventDestination_Topic
	Destination isEventDestination_Destination `protobuf_oneof:"destination"`
}

func (x *EventDestination) Reset() {
	*x = EventDestination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventDestination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventDestination) ProtoMessage() {}

func (x *EventDestination) ProtoReflect() protoreflect.Message {
	mi := &file_eventing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventDestination.ProtoReflect.Descriptor instead.
func (*EventDestination) Descriptor() ([]byte, []int) {
	return file_eventing_proto_rawDescGZIP(), []int{2}
}

func (m *EventDestination) GetDestination() isEventDestination_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (x *EventDestination) GetTopic() string {
	if x, ok := x.GetDestination().(*EventDestination_Topic); ok {
		return x.Topic
	}
	return ""
}

type isEventDestination_Destination interface {
	isEventDestination_Destination()
}

type EventDestination_Topic struct {
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3,oneof"`
}

func (*EventDestination_Topic) isEventDestination_Destination() {}

var File_eventing_proto protoreflect.FileDescriptor

var file_eventing_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x08, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2b,
	0x0a, 0x02, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x6b, 0x6b,
	0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x02, 0x69, 0x6e, 0x12, 0x32, 0x0a, 0x03, 0x6f,
	0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x6f, 0x75, 0x74, 0x22,
	0xaf, 0x01, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x16, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x32,
	0x0a, 0x14, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x64, 0x5f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x12,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x23, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x22, 0x39, 0x0a, 0x10, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x42, 0x0d, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6d, 0x0a, 0x12,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65,
	0x73, 0x73, 0x42, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x62, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x6b,
	0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x3b, 0x61, 0x6b, 0x6b,
	0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_eventing_proto_rawDescOnce sync.Once
	file_eventing_proto_rawDescData = file_eventing_proto_rawDesc
)

func file_eventing_proto_rawDescGZIP() []byte {
	file_eventing_proto_rawDescOnce.Do(func() {
		file_eventing_proto_rawDescData = protoimpl.X.CompressGZIP(file_eventing_proto_rawDescData)
	})
	return file_eventing_proto_rawDescData
}

var file_eventing_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_eventing_proto_goTypes = []interface{}{
	(*Eventing)(nil),         // 0: akkaserverless.Eventing
	(*EventSource)(nil),      // 1: akkaserverless.EventSource
	(*EventDestination)(nil), // 2: akkaserverless.EventDestination
}
var file_eventing_proto_depIdxs = []int32{
	1, // 0: akkaserverless.Eventing.in:type_name -> akkaserverless.EventSource
	2, // 1: akkaserverless.Eventing.out:type_name -> akkaserverless.EventDestination
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eventing_proto_init() }
func file_eventing_proto_init() {
	if File_eventing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eventing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Eventing); i {
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
		file_eventing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventSource); i {
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
		file_eventing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventDestination); i {
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
	file_eventing_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*EventSource_Topic)(nil),
		(*EventSource_EventSourcedEntity)(nil),
		(*EventSource_ValueEntity)(nil),
	}
	file_eventing_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*EventDestination_Topic)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eventing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eventing_proto_goTypes,
		DependencyIndexes: file_eventing_proto_depIdxs,
		MessageInfos:      file_eventing_proto_msgTypes,
	}.Build()
	File_eventing_proto = out.File
	file_eventing_proto_rawDesc = nil
	file_eventing_proto_goTypes = nil
	file_eventing_proto_depIdxs = nil
}