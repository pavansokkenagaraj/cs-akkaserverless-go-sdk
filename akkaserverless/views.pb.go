// Copyright 2021 Lightbend Inc.

// Extension for specifying which field in a message is to be considered an
// entity key, for the purposes of associating gRPC calls with entities and
// sharding.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: views.proto

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

// Views configuration for a gRPC method.
type View struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to UpdateOrQuery:
	//	*View_Update_
	//	*View_Query_
	UpdateOrQuery isView_UpdateOrQuery `protobuf_oneof:"update_or_query"`
}

func (x *View) Reset() {
	*x = View{}
	if protoimpl.UnsafeEnabled {
		mi := &file_views_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *View) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*View) ProtoMessage() {}

func (x *View) ProtoReflect() protoreflect.Message {
	mi := &file_views_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use View.ProtoReflect.Descriptor instead.
func (*View) Descriptor() ([]byte, []int) {
	return file_views_proto_rawDescGZIP(), []int{0}
}

func (m *View) GetUpdateOrQuery() isView_UpdateOrQuery {
	if m != nil {
		return m.UpdateOrQuery
	}
	return nil
}

func (x *View) GetUpdate() *View_Update {
	if x, ok := x.GetUpdateOrQuery().(*View_Update_); ok {
		return x.Update
	}
	return nil
}

func (x *View) GetQuery() *View_Query {
	if x, ok := x.GetUpdateOrQuery().(*View_Query_); ok {
		return x.Query
	}
	return nil
}

type isView_UpdateOrQuery interface {
	isView_UpdateOrQuery()
}

type View_Update_ struct {
	Update *View_Update `protobuf:"bytes,1,opt,name=update,proto3,oneof"`
}

type View_Query_ struct {
	Query *View_Query `protobuf:"bytes,2,opt,name=query,proto3,oneof"`
}

func (*View_Update_) isView_UpdateOrQuery() {}

func (*View_Query_) isView_UpdateOrQuery() {}

type View_Update struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates the name of the table that this processing method will persist
	// to, or that the query that is indexing will use.
	//
	// The return type of this call specifies the schema of the persisted value.
	// Any defined queries that select from this table will be selecting over this
	// schema.
	Table string `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	// Whether messages should automatically be persisted without further
	// processing.
	//
	// By default (false) all received message will be persisted at the key given
	// by the CloudEvents subject (that is, the entity key) without passing them
	// to the gRPC service call. That is, the user function does not need to
	// implement this service call.
	TransformUpdates bool `protobuf:"varint,2,opt,name=transform_updates,json=transformUpdates,proto3" json:"transform_updates,omitempty"`
}

func (x *View_Update) Reset() {
	*x = View_Update{}
	if protoimpl.UnsafeEnabled {
		mi := &file_views_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *View_Update) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*View_Update) ProtoMessage() {}

func (x *View_Update) ProtoReflect() protoreflect.Message {
	mi := &file_views_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use View_Update.ProtoReflect.Descriptor instead.
func (*View_Update) Descriptor() ([]byte, []int) {
	return file_views_proto_rawDescGZIP(), []int{0, 0}
}

func (x *View_Update) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *View_Update) GetTransformUpdates() bool {
	if x != nil {
		return x.TransformUpdates
	}
	return false
}

type View_Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A query that gets executed when this call is executed.
	//
	// This query is used to know how the view should be indexed.
	Query string `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	// Whether query results should be passed to the user function for further
	// processing before being returned to the client.
	//
	// By default (false), the user function does not process query results.
	TransformResults bool `protobuf:"varint,4,opt,name=transform_results,json=transformResults,proto3" json:"transform_results,omitempty"`
}

func (x *View_Query) Reset() {
	*x = View_Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_views_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *View_Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*View_Query) ProtoMessage() {}

func (x *View_Query) ProtoReflect() protoreflect.Message {
	mi := &file_views_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use View_Query.ProtoReflect.Descriptor instead.
func (*View_Query) Descriptor() ([]byte, []int) {
	return file_views_proto_rawDescGZIP(), []int{0, 1}
}

func (x *View_Query) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *View_Query) GetTransformResults() bool {
	if x != nil {
		return x.TransformResults
	}
	return false
}

var File_views_proto protoreflect.FileDescriptor

var file_views_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61,
	0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9d, 0x02, 0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x35, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x32, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e,
	0x56, 0x69, 0x65, 0x77, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x1a, 0x4b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73,
	0x1a, 0x4a, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x2b, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x42, 0x11, 0x0a, 0x0f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x72, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x42,
	0x6c, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x6c, 0x65, 0x73, 0x73, 0x42, 0x0a, 0x56, 0x69, 0x65, 0x77, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x62, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x3b, 0x61,
	0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_views_proto_rawDescOnce sync.Once
	file_views_proto_rawDescData = file_views_proto_rawDesc
)

func file_views_proto_rawDescGZIP() []byte {
	file_views_proto_rawDescOnce.Do(func() {
		file_views_proto_rawDescData = protoimpl.X.CompressGZIP(file_views_proto_rawDescData)
	})
	return file_views_proto_rawDescData
}

var file_views_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_views_proto_goTypes = []interface{}{
	(*View)(nil),        // 0: akkaserverless.View
	(*View_Update)(nil), // 1: akkaserverless.View.Update
	(*View_Query)(nil),  // 2: akkaserverless.View.Query
}
var file_views_proto_depIdxs = []int32{
	1, // 0: akkaserverless.View.update:type_name -> akkaserverless.View.Update
	2, // 1: akkaserverless.View.query:type_name -> akkaserverless.View.Query
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_views_proto_init() }
func file_views_proto_init() {
	if File_views_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_views_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*View); i {
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
		file_views_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*View_Update); i {
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
		file_views_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*View_Query); i {
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
	file_views_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*View_Update_)(nil),
		(*View_Query_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_views_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_views_proto_goTypes,
		DependencyIndexes: file_views_proto_depIdxs,
		MessageInfos:      file_views_proto_msgTypes,
	}.Build()
	File_views_proto = out.File
	file_views_proto_rawDesc = nil
	file_views_proto_goTypes = nil
	file_views_proto_depIdxs = nil
}