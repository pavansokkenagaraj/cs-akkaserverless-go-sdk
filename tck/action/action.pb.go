// Copyright 2021 Lightbend Inc.

//
// == Akka Serverless TCK model test for actions ==
// see tck/src/main/scala/com/akkaserverless/tck/ActionTCK.scala

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: action/action.proto

package action

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

//
// A `Request` message contains the steps that the entity should process.
// Steps are grouped for streamed responses. Steps must be processed in order.
//
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*ProcessGroup `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_action_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_action_action_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetGroups() []*ProcessGroup {
	if x != nil {
		return x.Groups
	}
	return nil
}

//
// A `ProcessGroup` contains the steps for one response.
//
type ProcessGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Steps []*ProcessStep `protobuf:"bytes,1,rep,name=steps,proto3" json:"steps,omitempty"`
}

func (x *ProcessGroup) Reset() {
	*x = ProcessGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_action_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessGroup) ProtoMessage() {}

func (x *ProcessGroup) ProtoReflect() protoreflect.Message {
	mi := &file_action_action_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessGroup.ProtoReflect.Descriptor instead.
func (*ProcessGroup) Descriptor() ([]byte, []int) {
	return file_action_action_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessGroup) GetSteps() []*ProcessStep {
	if x != nil {
		return x.Steps
	}
	return nil
}

//
// Each `ProcessStep` is one of:
//
// - Reply: reply with the given message in a `Response`.
// - Forward: forward to another service, in place of replying with a `Response`.
// - Fail: fail the current `Process` command by sending a failure.
// - SideEffect: add a side effect to the current reply, forward, or failure.
//
type ProcessStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Step:
	//	*ProcessStep_Reply
	//	*ProcessStep_Forward
	//	*ProcessStep_Fail
	//	*ProcessStep_Effect
	Step isProcessStep_Step `protobuf_oneof:"step"`
}

func (x *ProcessStep) Reset() {
	*x = ProcessStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_action_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessStep) ProtoMessage() {}

func (x *ProcessStep) ProtoReflect() protoreflect.Message {
	mi := &file_action_action_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessStep.ProtoReflect.Descriptor instead.
func (*ProcessStep) Descriptor() ([]byte, []int) {
	return file_action_action_proto_rawDescGZIP(), []int{2}
}

func (m *ProcessStep) GetStep() isProcessStep_Step {
	if m != nil {
		return m.Step
	}
	return nil
}

func (x *ProcessStep) GetReply() *Reply {
	if x, ok := x.GetStep().(*ProcessStep_Reply); ok {
		return x.Reply
	}
	return nil
}

func (x *ProcessStep) GetForward() *Forward {
	if x, ok := x.GetStep().(*ProcessStep_Forward); ok {
		return x.Forward
	}
	return nil
}

func (x *ProcessStep) GetFail() *Fail {
	if x, ok := x.GetStep().(*ProcessStep_Fail); ok {
		return x.Fail
	}
	return nil
}

func (x *ProcessStep) GetEffect() *SideEffect {
	if x, ok := x.GetStep().(*ProcessStep_Effect); ok {
		return x.Effect
	}
	return nil
}

type isProcessStep_Step interface {
	isProcessStep_Step()
}

type ProcessStep_Reply struct {
	Reply *Reply `protobuf:"bytes,1,opt,name=reply,proto3,oneof"`
}

type ProcessStep_Forward struct {
	Forward *Forward `protobuf:"bytes,2,opt,name=forward,proto3,oneof"`
}

type ProcessStep_Fail struct {
	Fail *Fail `protobuf:"bytes,3,opt,name=fail,proto3,oneof"`
}

type ProcessStep_Effect struct {
	Effect *SideEffect `protobuf:"bytes,4,opt,name=effect,proto3,oneof"`
}

func (*ProcessStep_Reply) isProcessStep_Step() {}

func (*ProcessStep_Forward) isProcessStep_Step() {}

func (*ProcessStep_Fail) isProcessStep_Step() {}

func (*ProcessStep_Effect) isProcessStep_Step() {}

//
// Reply with a message in the reponse.
//
type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_action_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_action_action_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_action_action_proto_rawDescGZIP(), []int{3}
}

func (x *Reply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

//
// Replace the response with a forward to `akkaserverless.tck.model.ActionTwo/Call`.
// The payload must be an `OtherRequest` message with the given `id`.
//
type Forward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Forward) Reset() {
	*x = Forward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_action_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Forward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Forward) ProtoMessage() {}

func (x *Forward) ProtoReflect() protoreflect.Message {
	mi := &file_action_action_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Forward.ProtoReflect.Descriptor instead.
func (*Forward) Descriptor() ([]byte, []int) {
	return file_action_action_proto_rawDescGZIP(), []int{4}
}

func (x *Forward) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

//
// Fail the current command with the given description `message`.
//
type Fail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Fail) Reset() {
	*x = Fail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_action_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fail) ProtoMessage() {}

func (x *Fail) ProtoReflect() protoreflect.Message {
	mi := &file_action_action_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fail.ProtoReflect.Descriptor instead.
func (*Fail) Descriptor() ([]byte, []int) {
	return file_action_action_proto_rawDescGZIP(), []int{5}
}

func (x *Fail) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

//
// Add a side effect to the reply, to `akkaserverless.tck.model.ActionTwo/Call`.
// The payload must be an `OtherRequest` message with the given `id`.
// The side effect should be marked synchronous based on the given `synchronous` value.
//
type SideEffect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Synchronous bool   `protobuf:"varint,2,opt,name=synchronous,proto3" json:"synchronous,omitempty"`
}

func (x *SideEffect) Reset() {
	*x = SideEffect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_action_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SideEffect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SideEffect) ProtoMessage() {}

func (x *SideEffect) ProtoReflect() protoreflect.Message {
	mi := &file_action_action_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SideEffect.ProtoReflect.Descriptor instead.
func (*SideEffect) Descriptor() ([]byte, []int) {
	return file_action_action_proto_rawDescGZIP(), []int{6}
}

func (x *SideEffect) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SideEffect) GetSynchronous() bool {
	if x != nil {
		return x.Synchronous
	}
	return false
}

//
// The `Response` message must contain the message from the corresponding reply step.
//
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_action_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_action_action_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_action_action_proto_rawDescGZIP(), []int{7}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

//
// The `OtherRequest` message must contain the id for the forward or side effect.
//
type OtherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OtherRequest) Reset() {
	*x = OtherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_action_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtherRequest) ProtoMessage() {}

func (x *OtherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_action_action_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtherRequest.ProtoReflect.Descriptor instead.
func (*OtherRequest) Descriptor() ([]byte, []int) {
	return file_action_action_proto_rawDescGZIP(), []int{8}
}

func (x *OtherRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_action_action_proto protoreflect.FileDescriptor

var file_action_action_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x45, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73,
	0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x52, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x42, 0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53,
	0x74, 0x65, 0x70, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x22, 0x9f, 0x02, 0x0a, 0x0b, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x65, 0x70, 0x12, 0x3e, 0x0a, 0x05, 0x72, 0x65,
	0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x6b, 0x6b, 0x61,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x44, 0x0a, 0x07, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x6b,
	0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x63, 0x6b,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x48, 0x00, 0x52, 0x07, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x12, 0x3b, 0x0a, 0x04, 0x66, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e,
	0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x46, 0x61, 0x69, 0x6c, 0x48, 0x00, 0x52, 0x04, 0x66, 0x61, 0x69, 0x6c, 0x12, 0x45, 0x0a,
	0x06, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x74,
	0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x69, 0x64, 0x65, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x06, 0x65, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x22, 0x21, 0x0a, 0x05,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x19, 0x0a, 0x07, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x04, 0x46, 0x61,
	0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3e, 0x0a, 0x0a,
	0x53, 0x69, 0x64, 0x65, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x79,
	0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x6f, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x6f, 0x75, 0x73, 0x22, 0x24, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x1e, 0x0a, 0x0c, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x32, 0xdf, 0x03, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x63, 0x6b,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x87, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x28, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x29, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73,
	0x73, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x74, 0x63, 0x6b, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x3a, 0x01, 0x2a, 0x12,
	0x6a, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x65, 0x64, 0x49, 0x6e, 0x12, 0x28, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e,
	0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x6b, 0x0a, 0x12, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x4f, 0x75,
	0x74, 0x12, 0x28, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65,
	0x73, 0x73, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x6b,
	0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x63, 0x6b,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x6a, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x12, 0x28, 0x2e, 0x61, 0x6b,
	0x6b, 0x61, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x63, 0x6b,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x28, 0x01, 0x30, 0x01, 0x32, 0x6d, 0x0a, 0x09, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x77,
	0x6f, 0x12, 0x60, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x2d, 0x2e, 0x61, 0x6b, 0x6b, 0x61,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x74, 0x68, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x5c, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6b, 0x6b, 0x61, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x63, 0x6b, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x62, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x6b, 0x6b, 0x61, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x73, 0x64, 0x6b, 0x2f,
	0x74, 0x63, 0x6b, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_action_action_proto_rawDescOnce sync.Once
	file_action_action_proto_rawDescData = file_action_action_proto_rawDesc
)

func file_action_action_proto_rawDescGZIP() []byte {
	file_action_action_proto_rawDescOnce.Do(func() {
		file_action_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_action_action_proto_rawDescData)
	})
	return file_action_action_proto_rawDescData
}

var file_action_action_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_action_action_proto_goTypes = []interface{}{
	(*Request)(nil),      // 0: akkaserverless.tck.model.action.Request
	(*ProcessGroup)(nil), // 1: akkaserverless.tck.model.action.ProcessGroup
	(*ProcessStep)(nil),  // 2: akkaserverless.tck.model.action.ProcessStep
	(*Reply)(nil),        // 3: akkaserverless.tck.model.action.Reply
	(*Forward)(nil),      // 4: akkaserverless.tck.model.action.Forward
	(*Fail)(nil),         // 5: akkaserverless.tck.model.action.Fail
	(*SideEffect)(nil),   // 6: akkaserverless.tck.model.action.SideEffect
	(*Response)(nil),     // 7: akkaserverless.tck.model.action.Response
	(*OtherRequest)(nil), // 8: akkaserverless.tck.model.action.OtherRequest
}
var file_action_action_proto_depIdxs = []int32{
	1,  // 0: akkaserverless.tck.model.action.Request.groups:type_name -> akkaserverless.tck.model.action.ProcessGroup
	2,  // 1: akkaserverless.tck.model.action.ProcessGroup.steps:type_name -> akkaserverless.tck.model.action.ProcessStep
	3,  // 2: akkaserverless.tck.model.action.ProcessStep.reply:type_name -> akkaserverless.tck.model.action.Reply
	4,  // 3: akkaserverless.tck.model.action.ProcessStep.forward:type_name -> akkaserverless.tck.model.action.Forward
	5,  // 4: akkaserverless.tck.model.action.ProcessStep.fail:type_name -> akkaserverless.tck.model.action.Fail
	6,  // 5: akkaserverless.tck.model.action.ProcessStep.effect:type_name -> akkaserverless.tck.model.action.SideEffect
	0,  // 6: akkaserverless.tck.model.action.ActionTckModel.ProcessUnary:input_type -> akkaserverless.tck.model.action.Request
	0,  // 7: akkaserverless.tck.model.action.ActionTckModel.ProcessStreamedIn:input_type -> akkaserverless.tck.model.action.Request
	0,  // 8: akkaserverless.tck.model.action.ActionTckModel.ProcessStreamedOut:input_type -> akkaserverless.tck.model.action.Request
	0,  // 9: akkaserverless.tck.model.action.ActionTckModel.ProcessStreamed:input_type -> akkaserverless.tck.model.action.Request
	8,  // 10: akkaserverless.tck.model.action.ActionTwo.Call:input_type -> akkaserverless.tck.model.action.OtherRequest
	7,  // 11: akkaserverless.tck.model.action.ActionTckModel.ProcessUnary:output_type -> akkaserverless.tck.model.action.Response
	7,  // 12: akkaserverless.tck.model.action.ActionTckModel.ProcessStreamedIn:output_type -> akkaserverless.tck.model.action.Response
	7,  // 13: akkaserverless.tck.model.action.ActionTckModel.ProcessStreamedOut:output_type -> akkaserverless.tck.model.action.Response
	7,  // 14: akkaserverless.tck.model.action.ActionTckModel.ProcessStreamed:output_type -> akkaserverless.tck.model.action.Response
	7,  // 15: akkaserverless.tck.model.action.ActionTwo.Call:output_type -> akkaserverless.tck.model.action.Response
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_action_action_proto_init() }
func file_action_action_proto_init() {
	if File_action_action_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_action_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_action_action_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessGroup); i {
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
		file_action_action_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessStep); i {
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
		file_action_action_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
		file_action_action_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Forward); i {
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
		file_action_action_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fail); i {
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
		file_action_action_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SideEffect); i {
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
		file_action_action_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_action_action_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtherRequest); i {
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
	file_action_action_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ProcessStep_Reply)(nil),
		(*ProcessStep_Forward)(nil),
		(*ProcessStep_Fail)(nil),
		(*ProcessStep_Effect)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_action_action_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_action_action_proto_goTypes,
		DependencyIndexes: file_action_action_proto_depIdxs,
		MessageInfos:      file_action_action_proto_msgTypes,
	}.Build()
	File_action_action_proto = out.File
	file_action_action_proto_rawDesc = nil
	file_action_action_proto_goTypes = nil
	file_action_action_proto_depIdxs = nil
}
