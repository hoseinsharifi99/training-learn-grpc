// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: calculator.proto

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

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNumber  int32 `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber int32 `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *SumRequest) GetFirstNumber() int32 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *SumRequest) GetSecondNumber() int32 {
	if x != nil {
		return x.SecondNumber
	}
	return 0
}

type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SumResult int32 `protobuf:"varint,1,opt,name=sum_result,json=sumResult,proto3" json:"sum_result,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *SumResponse) GetSumResult() int32 {
	if x != nil {
		return x.SumResult
	}
	return 0
}

type PrimeNumberDecompositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PrimeNumberDecompositionRequest) Reset() {
	*x = PrimeNumberDecompositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberDecompositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberDecompositionRequest) ProtoMessage() {}

func (x *PrimeNumberDecompositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberDecompositionRequest.ProtoReflect.Descriptor instead.
func (*PrimeNumberDecompositionRequest) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *PrimeNumberDecompositionRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type PrimeNumberDecompositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrimeFactor int64 `protobuf:"varint,1,opt,name=prime_factor,json=primeFactor,proto3" json:"prime_factor,omitempty"`
}

func (x *PrimeNumberDecompositionResponse) Reset() {
	*x = PrimeNumberDecompositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberDecompositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberDecompositionResponse) ProtoMessage() {}

func (x *PrimeNumberDecompositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberDecompositionResponse.ProtoReflect.Descriptor instead.
func (*PrimeNumberDecompositionResponse) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *PrimeNumberDecompositionResponse) GetPrimeFactor() int64 {
	if x != nil {
		return x.PrimeFactor
	}
	return 0
}

type ComputeAverageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *ComputeAverageRequest) Reset() {
	*x = ComputeAverageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeAverageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeAverageRequest) ProtoMessage() {}

func (x *ComputeAverageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeAverageRequest.ProtoReflect.Descriptor instead.
func (*ComputeAverageRequest) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{4}
}

func (x *ComputeAverageRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type ComputeAverageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Average float64 `protobuf:"fixed64,1,opt,name=average,proto3" json:"average,omitempty"`
}

func (x *ComputeAverageResponse) Reset() {
	*x = ComputeAverageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeAverageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeAverageResponse) ProtoMessage() {}

func (x *ComputeAverageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeAverageResponse.ProtoReflect.Descriptor instead.
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{5}
}

func (x *ComputeAverageResponse) GetAverage() float64 {
	if x != nil {
		return x.Average
	}
	return 0
}

type FindMaximumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *FindMaximumRequest) Reset() {
	*x = FindMaximumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaximumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaximumRequest) ProtoMessage() {}

func (x *FindMaximumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaximumRequest.ProtoReflect.Descriptor instead.
func (*FindMaximumRequest) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{6}
}

func (x *FindMaximumRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type FindMaximumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Maximum int32 `protobuf:"varint,1,opt,name=maximum,proto3" json:"maximum,omitempty"`
}

func (x *FindMaximumResponse) Reset() {
	*x = FindMaximumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaximumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaximumResponse) ProtoMessage() {}

func (x *FindMaximumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaximumResponse.ProtoReflect.Descriptor instead.
func (*FindMaximumResponse) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{7}
}

func (x *FindMaximumResponse) GetMaximum() int32 {
	if x != nil {
		return x.Maximum
	}
	return 0
}

type SquareRootRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *SquareRootRequest) Reset() {
	*x = SquareRootRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SquareRootRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquareRootRequest) ProtoMessage() {}

func (x *SquareRootRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquareRootRequest.ProtoReflect.Descriptor instead.
func (*SquareRootRequest) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{8}
}

func (x *SquareRootRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type SquareRootResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberRoot float64 `protobuf:"fixed64,1,opt,name=number_root,json=numberRoot,proto3" json:"number_root,omitempty"`
}

func (x *SquareRootResponse) Reset() {
	*x = SquareRootResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SquareRootResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquareRootResponse) ProtoMessage() {}

func (x *SquareRootResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquareRootResponse.ProtoReflect.Descriptor instead.
func (*SquareRootResponse) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{9}
}

func (x *SquareRootResponse) GetNumberRoot() float64 {
	if x != nil {
		return x.NumberRoot
	}
	return 0
}

var File_calculator_proto protoreflect.FileDescriptor

var file_calculator_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x54, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2c,
	0x0a, 0x0b, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x75, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x73, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x39, 0x0a, 0x1f,
	0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x45, 0x0a, 0x20, 0x50, 0x72, 0x69, 0x6d, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x2f,
	0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0x32, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x61, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x22, 0x2c, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d,
	0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x2f, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d,
	0x75, 0x6d, 0x22, 0x2b, 0x0a, 0x11, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x6f, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0x35, 0x0a, 0x12, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x72, 0x6f, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x6f, 0x6f, 0x74, 0x32, 0x8e, 0x03, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x03,
	0x53, 0x75, 0x6d, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x75, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x18, 0x50, 0x72,
	0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72,
	0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4f, 0x0a, 0x0e, 0x43, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x48, 0x0a, 0x0b, 0x46, 0x69,
	0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d,
	0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x12, 0x41, 0x0a, 0x0a, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x6f,
	0x6f, 0x74, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65,
	0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculator_proto_rawDescOnce sync.Once
	file_calculator_proto_rawDescData = file_calculator_proto_rawDesc
)

func file_calculator_proto_rawDescGZIP() []byte {
	file_calculator_proto_rawDescOnce.Do(func() {
		file_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_proto_rawDescData)
	})
	return file_calculator_proto_rawDescData
}

var file_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_calculator_proto_goTypes = []interface{}{
	(*SumRequest)(nil),                       // 0: grpc.SumRequest
	(*SumResponse)(nil),                      // 1: grpc.SumResponse
	(*PrimeNumberDecompositionRequest)(nil),  // 2: grpc.PrimeNumberDecompositionRequest
	(*PrimeNumberDecompositionResponse)(nil), // 3: grpc.PrimeNumberDecompositionResponse
	(*ComputeAverageRequest)(nil),            // 4: grpc.ComputeAverageRequest
	(*ComputeAverageResponse)(nil),           // 5: grpc.ComputeAverageResponse
	(*FindMaximumRequest)(nil),               // 6: grpc.FindMaximumRequest
	(*FindMaximumResponse)(nil),              // 7: grpc.FindMaximumResponse
	(*SquareRootRequest)(nil),                // 8: grpc.SquareRootRequest
	(*SquareRootResponse)(nil),               // 9: grpc.SquareRootResponse
}
var file_calculator_proto_depIdxs = []int32{
	0, // 0: grpc.CalculatorService.Sum:input_type -> grpc.SumRequest
	2, // 1: grpc.CalculatorService.PrimeNumberDecomposition:input_type -> grpc.PrimeNumberDecompositionRequest
	4, // 2: grpc.CalculatorService.ComputeAverage:input_type -> grpc.ComputeAverageRequest
	6, // 3: grpc.CalculatorService.FindMaximum:input_type -> grpc.FindMaximumRequest
	8, // 4: grpc.CalculatorService.SquareRoot:input_type -> grpc.SquareRootRequest
	1, // 5: grpc.CalculatorService.Sum:output_type -> grpc.SumResponse
	3, // 6: grpc.CalculatorService.PrimeNumberDecomposition:output_type -> grpc.PrimeNumberDecompositionResponse
	5, // 7: grpc.CalculatorService.ComputeAverage:output_type -> grpc.ComputeAverageResponse
	7, // 8: grpc.CalculatorService.FindMaximum:output_type -> grpc.FindMaximumResponse
	9, // 9: grpc.CalculatorService.SquareRoot:output_type -> grpc.SquareRootResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_proto_init() }
func file_calculator_proto_init() {
	if File_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
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
		file_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
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
		file_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberDecompositionRequest); i {
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
		file_calculator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberDecompositionResponse); i {
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
		file_calculator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeAverageRequest); i {
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
		file_calculator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeAverageResponse); i {
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
		file_calculator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMaximumRequest); i {
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
		file_calculator_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMaximumResponse); i {
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
		file_calculator_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SquareRootRequest); i {
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
		file_calculator_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SquareRootResponse); i {
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
			RawDescriptor: file_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_proto_depIdxs,
		MessageInfos:      file_calculator_proto_msgTypes,
	}.Build()
	File_calculator_proto = out.File
	file_calculator_proto_rawDesc = nil
	file_calculator_proto_goTypes = nil
	file_calculator_proto_depIdxs = nil
}
