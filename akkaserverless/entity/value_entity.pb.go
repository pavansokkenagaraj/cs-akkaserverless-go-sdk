// Copyright 2019 Lightbend Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// gRPC interface for common messages and services for value-based Entity user functions.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: value_entity.proto

package entity

import (
	protocol "github.com/pavansokkenagaraj/akkaserverless-go-sdk/akkaserverless/protocol"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Input message type for the gRPC stream in.
type ValueEntityStreamIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*ValueEntityStreamIn_Init
	//	*ValueEntityStreamIn_Command
	Message isValueEntityStreamIn_Message `protobuf_oneof:"message"`
}

func (x *ValueEntityStreamIn) Reset() {
	*x = ValueEntityStreamIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_value_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueEntityStreamIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueEntityStreamIn) ProtoMessage() {}

func (x *ValueEntityStreamIn) ProtoReflect() protoreflect.Message {
	mi := &file_value_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueEntityStreamIn.ProtoReflect.Descriptor instead.
func (*ValueEntityStreamIn) Descriptor() ([]byte, []int) {
	return file_value_entity_proto_rawDescGZIP(), []int{0}
}

func (m *ValueEntityStreamIn) GetMessage() isValueEntityStreamIn_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *ValueEntityStreamIn) GetInit() *ValueEntityInit {
	if x, ok := x.GetMessage().(*ValueEntityStreamIn_Init); ok {
		return x.Init
	}
	return nil
}

func (x *ValueEntityStreamIn) GetCommand() *protocol.Command {
	if x, ok := x.GetMessage().(*ValueEntityStreamIn_Command); ok {
		return x.Command
	}
	return nil
}

type isValueEntityStreamIn_Message interface {
	isValueEntityStreamIn_Message()
}

type ValueEntityStreamIn_Init struct {
	Init *ValueEntityInit `protobuf:"bytes,1,opt,name=init,proto3,oneof"`
}

type ValueEntityStreamIn_Command struct {
	Command *protocol.Command `protobuf:"bytes,2,opt,name=command,proto3,oneof"`
}

func (*ValueEntityStreamIn_Init) isValueEntityStreamIn_Message() {}

func (*ValueEntityStreamIn_Command) isValueEntityStreamIn_Message() {}

// The init message. This will always be the first message sent to the entity when it is loaded.
type ValueEntityInit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the service that implements this entity.
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The id of the entity.
	EntityId string `protobuf:"bytes,2,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	// The initial state of the entity.
	State *ValueEntityInitState `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *ValueEntityInit) Reset() {
	*x = ValueEntityInit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_value_entity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueEntityInit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueEntityInit) ProtoMessage() {}

func (x *ValueEntityInit) ProtoReflect() protoreflect.Message {
	mi := &file_value_entity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueEntityInit.ProtoReflect.Descriptor instead.
func (*ValueEntityInit) Descriptor() ([]byte, []int) {
	return file_value_entity_proto_rawDescGZIP(), []int{1}
}

func (x *ValueEntityInit) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *ValueEntityInit) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *ValueEntityInit) GetState() *ValueEntityInitState {
	if x != nil {
		return x.State
	}
	return nil
}

// The state of the entity when it is first activated.
type ValueEntityInitState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The value of the entity state, if the entity has already been created.
	Value *any.Any `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ValueEntityInitState) Reset() {
	*x = ValueEntityInitState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_value_entity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueEntityInitState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueEntityInitState) ProtoMessage() {}

func (x *ValueEntityInitState) ProtoReflect() protoreflect.Message {
	mi := &file_value_entity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueEntityInitState.ProtoReflect.Descriptor instead.
func (*ValueEntityInitState) Descriptor() ([]byte, []int) {
	return file_value_entity_proto_rawDescGZIP(), []int{2}
}

func (x *ValueEntityInitState) GetValue() *any.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

// Output message type for the gRPC stream out.
type ValueEntityStreamOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//	*ValueEntityStreamOut_Reply
	//	*ValueEntityStreamOut_Failure
	Message isValueEntityStreamOut_Message `protobuf_oneof:"message"`
}

func (x *ValueEntityStreamOut) Reset() {
	*x = ValueEntityStreamOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_value_entity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueEntityStreamOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueEntityStreamOut) ProtoMessage() {}

func (x *ValueEntityStreamOut) ProtoReflect() protoreflect.Message {
	mi := &file_value_entity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueEntityStreamOut.ProtoReflect.Descriptor instead.
func (*ValueEntityStreamOut) Descriptor() ([]byte, []int) {
	return file_value_entity_proto_rawDescGZIP(), []int{3}
}

func (m *ValueEntityStreamOut) GetMessage() isValueEntityStreamOut_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *ValueEntityStreamOut) GetReply() *ValueEntityReply {
	if x, ok := x.GetMessage().(*ValueEntityStreamOut_Reply); ok {
		return x.Reply
	}
	return nil
}

func (x *ValueEntityStreamOut) GetFailure() *protocol.Failure {
	if x, ok := x.GetMessage().(*ValueEntityStreamOut_Failure); ok {
		return x.Failure
	}
	return nil
}

type isValueEntityStreamOut_Message interface {
	isValueEntityStreamOut_Message()
}

type ValueEntityStreamOut_Reply struct {
	Reply *ValueEntityReply `protobuf:"bytes,1,opt,name=reply,proto3,oneof"`
}

type ValueEntityStreamOut_Failure struct {
	Failure *protocol.Failure `protobuf:"bytes,2,opt,name=failure,proto3,oneof"`
}

func (*ValueEntityStreamOut_Reply) isValueEntityStreamOut_Message() {}

func (*ValueEntityStreamOut_Failure) isValueEntityStreamOut_Message() {}

// A reply to a command.
type ValueEntityReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The command being replied to
	CommandId int64 `protobuf:"varint,1,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	// The action to take for the client response
	ClientAction *protocol.ClientAction `protobuf:"bytes,2,opt,name=client_action,json=clientAction,proto3" json:"client_action,omitempty"`
	// Any side effects to perform
	SideEffects []*protocol.SideEffect `protobuf:"bytes,3,rep,name=side_effects,json=sideEffects,proto3" json:"side_effects,omitempty"`
	// The action to take on the value entity
	StateAction *ValueEntityAction `protobuf:"bytes,4,opt,name=state_action,json=stateAction,proto3" json:"state_action,omitempty"`
}

func (x *ValueEntityReply) Reset() {
	*x = ValueEntityReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_value_entity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueEntityReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueEntityReply) ProtoMessage() {}

func (x *ValueEntityReply) ProtoReflect() protoreflect.Message {
	mi := &file_value_entity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueEntityReply.ProtoReflect.Descriptor instead.
func (*ValueEntityReply) Descriptor() ([]byte, []int) {
	return file_value_entity_proto_rawDescGZIP(), []int{4}
}

func (x *ValueEntityReply) GetCommandId() int64 {
	if x != nil {
		return x.CommandId
	}
	return 0
}

func (x *ValueEntityReply) GetClientAction() *protocol.ClientAction {
	if x != nil {
		return x.ClientAction
	}
	return nil
}

func (x *ValueEntityReply) GetSideEffects() []*protocol.SideEffect {
	if x != nil {
		return x.SideEffects
	}
	return nil
}

func (x *ValueEntityReply) GetStateAction() *ValueEntityAction {
	if x != nil {
		return x.StateAction
	}
	return nil
}

// An action to take for changing the entity state.
type ValueEntityAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//	*ValueEntityAction_Update
	//	*ValueEntityAction_Delete
	Action isValueEntityAction_Action `protobuf_oneof:"action"`
}

func (x *ValueEntityAction) Reset() {
	*x = ValueEntityAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_value_entity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueEntityAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueEntityAction) ProtoMessage() {}

func (x *ValueEntityAction) ProtoReflect() protoreflect.Message {
	mi := &file_value_entity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueEntityAction.ProtoReflect.Descriptor instead.
func (*ValueEntityAction) Descriptor() ([]byte, []int) {
	return file_value_entity_proto_rawDescGZIP(), []int{5}
}

func (m *ValueEntityAction) GetAction() isValueEntityAction_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *ValueEntityAction) GetUpdate() *ValueEntityUpdate {
	if x, ok := x.GetAction().(*ValueEntityAction_Update); ok {
		return x.Update
	}
	return nil
}

func (x *ValueEntityAction) GetDelete() *ValueEntityDelete {
	if x, ok := x.GetAction().(*ValueEntityAction_Delete); ok {
		return x.Delete
	}
	return nil
}

type isValueEntityAction_Action interface {
	isValueEntityAction_Action()
}

type ValueEntityAction_Update struct {
	Update *ValueEntityUpdate `protobuf:"bytes,1,opt,name=update,proto3,oneof"`
}

type ValueEntityAction_Delete struct {
	Delete *ValueEntityDelete `protobuf:"bytes,2,opt,name=delete,proto3,oneof"`
}

func (*ValueEntityAction_Update) isValueEntityAction_Action() {}

func (*ValueEntityAction_Delete) isValueEntityAction_Action() {}

// An action which updates the persisted value of the Value entity. If the entity is not yet persisted, it will be created.
type ValueEntityUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The value to set.
	Value *any.Any `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ValueEntityUpdate) Reset() {
	*x = ValueEntityUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_value_entity_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueEntityUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueEntityUpdate) ProtoMessage() {}

func (x *ValueEntityUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_value_entity_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueEntityUpdate.ProtoReflect.Descriptor instead.
func (*ValueEntityUpdate) Descriptor() ([]byte, []int) {
	return file_value_entity_proto_rawDescGZIP(), []int{6}
}

func (x *ValueEntityUpdate) GetValue() *any.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

// An action which deletes the persisted value of the Value entity.
type ValueEntityDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ValueEntityDelete) Reset() {
	*x = ValueEntityDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_value_entity_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueEntityDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueEntityDelete) ProtoMessage() {}

func (x *ValueEntityDelete) ProtoReflect() protoreflect.Message {
	mi := &file_value_entity_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueEntityDelete.ProtoReflect.Descriptor instead.
func (*ValueEntityDelete) Descriptor() ([]byte, []int) {
	return file_value_entity_proto_rawDescGZIP(), []int{7}
}

var File_value_entity_proto protoreflect.FileDescriptor

var file_value_entity_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x90, 0x01, 0x0a, 0x13, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x12, 0x3d, 0x0a, 0x04, 0x69, 0x6e, 0x69, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x69, 0x74, 0x48,
	0x00, 0x52, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48, 0x00, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x69, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x42, 0x0a, 0x14, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x69, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x94, 0x01, 0x0a, 0x14, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f, 0x75, 0x74, 0x12, 0x40, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x07, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x48, 0x00, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xf9, 0x01, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0d, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0c, 0x73, 0x69, 0x64,
	0x65, 0x5f, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x69, 0x64,
	0x65, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x0b, 0x73, 0x69, 0x64, 0x65, 0x45, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x73, 0x12, 0x4c, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xa7, 0x01, 0x0a, 0x11, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x43, 0x0a,
	0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x00, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x11,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x13, 0x0a,
	0x11, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x32, 0x78, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x69, 0x0a, 0x06, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x2b, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x1a, 0x2c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x55, 0x0a, 0x16,
	0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x69, 0x6f, 0x2f,
	0x67, 0x6f, 0x2d, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x3b, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_value_entity_proto_rawDescOnce sync.Once
	file_value_entity_proto_rawDescData = file_value_entity_proto_rawDesc
)

func file_value_entity_proto_rawDescGZIP() []byte {
	file_value_entity_proto_rawDescOnce.Do(func() {
		file_value_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_value_entity_proto_rawDescData)
	})
	return file_value_entity_proto_rawDescData
}

var file_value_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_value_entity_proto_goTypes = []interface{}{
	(*ValueEntityStreamIn)(nil),   // 0: cloudstate.valueentity.ValueEntityStreamIn
	(*ValueEntityInit)(nil),       // 1: cloudstate.valueentity.ValueEntityInit
	(*ValueEntityInitState)(nil),  // 2: cloudstate.valueentity.ValueEntityInitState
	(*ValueEntityStreamOut)(nil),  // 3: cloudstate.valueentity.ValueEntityStreamOut
	(*ValueEntityReply)(nil),      // 4: cloudstate.valueentity.ValueEntityReply
	(*ValueEntityAction)(nil),     // 5: cloudstate.valueentity.ValueEntityAction
	(*ValueEntityUpdate)(nil),     // 6: cloudstate.valueentity.ValueEntityUpdate
	(*ValueEntityDelete)(nil),     // 7: cloudstate.valueentity.ValueEntityDelete
	(*protocol.Command)(nil),      // 8: cloudstate.Command
	(*any.Any)(nil),               // 9: google.protobuf.Any
	(*protocol.Failure)(nil),      // 10: cloudstate.Failure
	(*protocol.ClientAction)(nil), // 11: cloudstate.ClientAction
	(*protocol.SideEffect)(nil),   // 12: cloudstate.SideEffect
}
var file_value_entity_proto_depIdxs = []int32{
	1,  // 0: cloudstate.valueentity.ValueEntityStreamIn.init:type_name -> cloudstate.valueentity.ValueEntityInit
	8,  // 1: cloudstate.valueentity.ValueEntityStreamIn.command:type_name -> cloudstate.Command
	2,  // 2: cloudstate.valueentity.ValueEntityInit.state:type_name -> cloudstate.valueentity.ValueEntityInitState
	9,  // 3: cloudstate.valueentity.ValueEntityInitState.value:type_name -> google.protobuf.Any
	4,  // 4: cloudstate.valueentity.ValueEntityStreamOut.reply:type_name -> cloudstate.valueentity.ValueEntityReply
	10, // 5: cloudstate.valueentity.ValueEntityStreamOut.failure:type_name -> cloudstate.Failure
	11, // 6: cloudstate.valueentity.ValueEntityReply.client_action:type_name -> cloudstate.ClientAction
	12, // 7: cloudstate.valueentity.ValueEntityReply.side_effects:type_name -> cloudstate.SideEffect
	5,  // 8: cloudstate.valueentity.ValueEntityReply.state_action:type_name -> cloudstate.valueentity.ValueEntityAction
	6,  // 9: cloudstate.valueentity.ValueEntityAction.update:type_name -> cloudstate.valueentity.ValueEntityUpdate
	7,  // 10: cloudstate.valueentity.ValueEntityAction.delete:type_name -> cloudstate.valueentity.ValueEntityDelete
	9,  // 11: cloudstate.valueentity.ValueEntityUpdate.value:type_name -> google.protobuf.Any
	0,  // 12: cloudstate.valueentity.ValueEntity.handle:input_type -> cloudstate.valueentity.ValueEntityStreamIn
	3,  // 13: cloudstate.valueentity.ValueEntity.handle:output_type -> cloudstate.valueentity.ValueEntityStreamOut
	13, // [13:14] is the sub-list for method output_type
	12, // [12:13] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_value_entity_proto_init() }
func file_value_entity_proto_init() {
	if File_value_entity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_value_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueEntityStreamIn); i {
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
		file_value_entity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueEntityInit); i {
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
		file_value_entity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueEntityInitState); i {
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
		file_value_entity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueEntityStreamOut); i {
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
		file_value_entity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueEntityReply); i {
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
		file_value_entity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueEntityAction); i {
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
		file_value_entity_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueEntityUpdate); i {
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
		file_value_entity_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueEntityDelete); i {
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
	file_value_entity_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ValueEntityStreamIn_Init)(nil),
		(*ValueEntityStreamIn_Command)(nil),
	}
	file_value_entity_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ValueEntityStreamOut_Reply)(nil),
		(*ValueEntityStreamOut_Failure)(nil),
	}
	file_value_entity_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*ValueEntityAction_Update)(nil),
		(*ValueEntityAction_Delete)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_value_entity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_value_entity_proto_goTypes,
		DependencyIndexes: file_value_entity_proto_depIdxs,
		MessageInfos:      file_value_entity_proto_msgTypes,
	}.Build()
	File_value_entity_proto = out.File
	file_value_entity_proto_rawDesc = nil
	file_value_entity_proto_goTypes = nil
	file_value_entity_proto_depIdxs = nil
}