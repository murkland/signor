// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: api.proto

package pb

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

type NegotiateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Which:
	//	*NegotiateRequest_Start_
	//	*NegotiateRequest_Answer_
	//	*NegotiateRequest_IceCandidate
	Which isNegotiateRequest_Which `protobuf_oneof:"which"`
}

func (x *NegotiateRequest) Reset() {
	*x = NegotiateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NegotiateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NegotiateRequest) ProtoMessage() {}

func (x *NegotiateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NegotiateRequest.ProtoReflect.Descriptor instead.
func (*NegotiateRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (m *NegotiateRequest) GetWhich() isNegotiateRequest_Which {
	if m != nil {
		return m.Which
	}
	return nil
}

func (x *NegotiateRequest) GetStart() *NegotiateRequest_Start {
	if x, ok := x.GetWhich().(*NegotiateRequest_Start_); ok {
		return x.Start
	}
	return nil
}

func (x *NegotiateRequest) GetAnswer() *NegotiateRequest_Answer {
	if x, ok := x.GetWhich().(*NegotiateRequest_Answer_); ok {
		return x.Answer
	}
	return nil
}

func (x *NegotiateRequest) GetIceCandidate() *NegotiateRequest_ICECandidate {
	if x, ok := x.GetWhich().(*NegotiateRequest_IceCandidate); ok {
		return x.IceCandidate
	}
	return nil
}

type isNegotiateRequest_Which interface {
	isNegotiateRequest_Which()
}

type NegotiateRequest_Start_ struct {
	Start *NegotiateRequest_Start `protobuf:"bytes,1,opt,name=start,proto3,oneof"`
}

type NegotiateRequest_Answer_ struct {
	Answer *NegotiateRequest_Answer `protobuf:"bytes,2,opt,name=answer,proto3,oneof"`
}

type NegotiateRequest_IceCandidate struct {
	IceCandidate *NegotiateRequest_ICECandidate `protobuf:"bytes,3,opt,name=ice_candidate,json=iceCandidate,proto3,oneof"`
}

func (*NegotiateRequest_Start_) isNegotiateRequest_Which() {}

func (*NegotiateRequest_Answer_) isNegotiateRequest_Which() {}

func (*NegotiateRequest_IceCandidate) isNegotiateRequest_Which() {}

type NegotiateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Which:
	//	*NegotiateResponse_Offer_
	//	*NegotiateResponse_Answer_
	//	*NegotiateResponse_IceCandidate
	Which isNegotiateResponse_Which `protobuf_oneof:"which"`
}

func (x *NegotiateResponse) Reset() {
	*x = NegotiateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NegotiateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NegotiateResponse) ProtoMessage() {}

func (x *NegotiateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NegotiateResponse.ProtoReflect.Descriptor instead.
func (*NegotiateResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (m *NegotiateResponse) GetWhich() isNegotiateResponse_Which {
	if m != nil {
		return m.Which
	}
	return nil
}

func (x *NegotiateResponse) GetOffer() *NegotiateResponse_Offer {
	if x, ok := x.GetWhich().(*NegotiateResponse_Offer_); ok {
		return x.Offer
	}
	return nil
}

func (x *NegotiateResponse) GetAnswer() *NegotiateResponse_Answer {
	if x, ok := x.GetWhich().(*NegotiateResponse_Answer_); ok {
		return x.Answer
	}
	return nil
}

func (x *NegotiateResponse) GetIceCandidate() *NegotiateResponse_ICECandidate {
	if x, ok := x.GetWhich().(*NegotiateResponse_IceCandidate); ok {
		return x.IceCandidate
	}
	return nil
}

type isNegotiateResponse_Which interface {
	isNegotiateResponse_Which()
}

type NegotiateResponse_Offer_ struct {
	Offer *NegotiateResponse_Offer `protobuf:"bytes,1,opt,name=offer,proto3,oneof"`
}

type NegotiateResponse_Answer_ struct {
	Answer *NegotiateResponse_Answer `protobuf:"bytes,2,opt,name=answer,proto3,oneof"`
}

type NegotiateResponse_IceCandidate struct {
	IceCandidate *NegotiateResponse_ICECandidate `protobuf:"bytes,3,opt,name=ice_candidate,json=iceCandidate,proto3,oneof"`
}

func (*NegotiateResponse_Offer_) isNegotiateResponse_Which() {}

func (*NegotiateResponse_Answer_) isNegotiateResponse_Which() {}

func (*NegotiateResponse_IceCandidate) isNegotiateResponse_Which() {}

type NegotiateRequest_Start struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	OfferSdp  string `protobuf:"bytes,2,opt,name=offer_sdp,json=offerSdp,proto3" json:"offer_sdp,omitempty"`
}

func (x *NegotiateRequest_Start) Reset() {
	*x = NegotiateRequest_Start{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NegotiateRequest_Start) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NegotiateRequest_Start) ProtoMessage() {}

func (x *NegotiateRequest_Start) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NegotiateRequest_Start.ProtoReflect.Descriptor instead.
func (*NegotiateRequest_Start) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0, 0}
}

func (x *NegotiateRequest_Start) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *NegotiateRequest_Start) GetOfferSdp() string {
	if x != nil {
		return x.OfferSdp
	}
	return ""
}

type NegotiateRequest_Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sdp string `protobuf:"bytes,1,opt,name=sdp,proto3" json:"sdp,omitempty"`
}

func (x *NegotiateRequest_Answer) Reset() {
	*x = NegotiateRequest_Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NegotiateRequest_Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NegotiateRequest_Answer) ProtoMessage() {}

func (x *NegotiateRequest_Answer) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NegotiateRequest_Answer.ProtoReflect.Descriptor instead.
func (*NegotiateRequest_Answer) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0, 1}
}

func (x *NegotiateRequest_Answer) GetSdp() string {
	if x != nil {
		return x.Sdp
	}
	return ""
}

type NegotiateRequest_ICECandidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IceCandidate string `protobuf:"bytes,1,opt,name=ice_candidate,json=iceCandidate,proto3" json:"ice_candidate,omitempty"`
}

func (x *NegotiateRequest_ICECandidate) Reset() {
	*x = NegotiateRequest_ICECandidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NegotiateRequest_ICECandidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NegotiateRequest_ICECandidate) ProtoMessage() {}

func (x *NegotiateRequest_ICECandidate) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NegotiateRequest_ICECandidate.ProtoReflect.Descriptor instead.
func (*NegotiateRequest_ICECandidate) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0, 2}
}

func (x *NegotiateRequest_ICECandidate) GetIceCandidate() string {
	if x != nil {
		return x.IceCandidate
	}
	return ""
}

type NegotiateResponse_Offer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sdp string `protobuf:"bytes,1,opt,name=sdp,proto3" json:"sdp,omitempty"`
}

func (x *NegotiateResponse_Offer) Reset() {
	*x = NegotiateResponse_Offer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NegotiateResponse_Offer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NegotiateResponse_Offer) ProtoMessage() {}

func (x *NegotiateResponse_Offer) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NegotiateResponse_Offer.ProtoReflect.Descriptor instead.
func (*NegotiateResponse_Offer) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1, 0}
}

func (x *NegotiateResponse_Offer) GetSdp() string {
	if x != nil {
		return x.Sdp
	}
	return ""
}

type NegotiateResponse_Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sdp string `protobuf:"bytes,1,opt,name=sdp,proto3" json:"sdp,omitempty"`
}

func (x *NegotiateResponse_Answer) Reset() {
	*x = NegotiateResponse_Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NegotiateResponse_Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NegotiateResponse_Answer) ProtoMessage() {}

func (x *NegotiateResponse_Answer) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NegotiateResponse_Answer.ProtoReflect.Descriptor instead.
func (*NegotiateResponse_Answer) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1, 1}
}

func (x *NegotiateResponse_Answer) GetSdp() string {
	if x != nil {
		return x.Sdp
	}
	return ""
}

type NegotiateResponse_ICECandidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IceCandidate string `protobuf:"bytes,1,opt,name=ice_candidate,json=iceCandidate,proto3" json:"ice_candidate,omitempty"`
}

func (x *NegotiateResponse_ICECandidate) Reset() {
	*x = NegotiateResponse_ICECandidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NegotiateResponse_ICECandidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NegotiateResponse_ICECandidate) ProtoMessage() {}

func (x *NegotiateResponse_ICECandidate) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NegotiateResponse_ICECandidate.ProtoReflect.Descriptor instead.
func (*NegotiateResponse_ICECandidate) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1, 2}
}

func (x *NegotiateResponse_ICECandidate) GetIceCandidate() string {
	if x != nil {
		return x.IceCandidate
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x69, 0x67,
	0x6e, 0x6f, 0x72, 0x22, 0xf2, 0x02, 0x0a, 0x10, 0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x6f, 0x72,
	0x2e, 0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x39, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x2e, 0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x48, 0x00, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x0d, 0x69,
	0x63, 0x65, 0x5f, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x2e, 0x4e, 0x65, 0x67, 0x6f,
	0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x43, 0x45,
	0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x63, 0x65,
	0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x43, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x73, 0x64, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x53, 0x64, 0x70, 0x1a, 0x1a,
	0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x64, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x64, 0x70, 0x1a, 0x33, 0x0a, 0x0c, 0x49, 0x43,
	0x45, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x63,
	0x65, 0x5f, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x07, 0x0a, 0x05, 0x77, 0x68, 0x69, 0x63, 0x68, 0x22, 0xcc, 0x02, 0x0a, 0x11, 0x4e, 0x65, 0x67,
	0x6f, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x73, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x2e, 0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x48, 0x00,
	0x52, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x6f, 0x72,
	0x2e, 0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x0d, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x69, 0x67,
	0x6e, 0x6f, 0x72, 0x2e, 0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x43, 0x45, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x1a, 0x19, 0x0a, 0x05, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x64, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x64, 0x70, 0x1a, 0x1a, 0x0a,
	0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x64, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x64, 0x70, 0x1a, 0x33, 0x0a, 0x0c, 0x49, 0x43, 0x45,
	0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x63, 0x65,
	0x5f, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x42, 0x07,
	0x0a, 0x05, 0x77, 0x68, 0x69, 0x63, 0x68, 0x32, 0x56, 0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x4e, 0x65, 0x67,
	0x6f, 0x74, 0x69, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x2e,
	0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x2e, 0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x62,
	0x61, 0x72, 0x65, 0x6e, 0x61, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_proto_goTypes = []interface{}{
	(*NegotiateRequest)(nil),               // 0: signor.NegotiateRequest
	(*NegotiateResponse)(nil),              // 1: signor.NegotiateResponse
	(*NegotiateRequest_Start)(nil),         // 2: signor.NegotiateRequest.Start
	(*NegotiateRequest_Answer)(nil),        // 3: signor.NegotiateRequest.Answer
	(*NegotiateRequest_ICECandidate)(nil),  // 4: signor.NegotiateRequest.ICECandidate
	(*NegotiateResponse_Offer)(nil),        // 5: signor.NegotiateResponse.Offer
	(*NegotiateResponse_Answer)(nil),       // 6: signor.NegotiateResponse.Answer
	(*NegotiateResponse_ICECandidate)(nil), // 7: signor.NegotiateResponse.ICECandidate
}
var file_api_proto_depIdxs = []int32{
	2, // 0: signor.NegotiateRequest.start:type_name -> signor.NegotiateRequest.Start
	3, // 1: signor.NegotiateRequest.answer:type_name -> signor.NegotiateRequest.Answer
	4, // 2: signor.NegotiateRequest.ice_candidate:type_name -> signor.NegotiateRequest.ICECandidate
	5, // 3: signor.NegotiateResponse.offer:type_name -> signor.NegotiateResponse.Offer
	6, // 4: signor.NegotiateResponse.answer:type_name -> signor.NegotiateResponse.Answer
	7, // 5: signor.NegotiateResponse.ice_candidate:type_name -> signor.NegotiateResponse.ICECandidate
	0, // 6: signor.SessionService.Negotiate:input_type -> signor.NegotiateRequest
	1, // 7: signor.SessionService.Negotiate:output_type -> signor.NegotiateResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NegotiateRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NegotiateResponse); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NegotiateRequest_Start); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NegotiateRequest_Answer); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NegotiateRequest_ICECandidate); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NegotiateResponse_Offer); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NegotiateResponse_Answer); i {
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
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NegotiateResponse_ICECandidate); i {
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
	file_api_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*NegotiateRequest_Start_)(nil),
		(*NegotiateRequest_Answer_)(nil),
		(*NegotiateRequest_IceCandidate)(nil),
	}
	file_api_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*NegotiateResponse_Offer_)(nil),
		(*NegotiateResponse_Answer_)(nil),
		(*NegotiateResponse_IceCandidate)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
