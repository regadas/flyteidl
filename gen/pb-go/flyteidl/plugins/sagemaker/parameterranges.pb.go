// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/sagemaker/parameterranges.proto

package plugins

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HyperparameterScalingType int32

const (
	HyperparameterScalingType_AUTO               HyperparameterScalingType = 0
	HyperparameterScalingType_LINEAR             HyperparameterScalingType = 1
	HyperparameterScalingType_LOGARITHMIC        HyperparameterScalingType = 2
	HyperparameterScalingType_REVERSELOGARITHMIC HyperparameterScalingType = 3
)

var HyperparameterScalingType_name = map[int32]string{
	0: "AUTO",
	1: "LINEAR",
	2: "LOGARITHMIC",
	3: "REVERSELOGARITHMIC",
}

var HyperparameterScalingType_value = map[string]int32{
	"AUTO":               0,
	"LINEAR":             1,
	"LOGARITHMIC":        2,
	"REVERSELOGARITHMIC": 3,
}

func (x HyperparameterScalingType) String() string {
	return proto.EnumName(HyperparameterScalingType_name, int32(x))
}

func (HyperparameterScalingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d87fa566d23b1e82, []int{0}
}

type ContinuousParameterRange struct {
	MaxValue             float64                   `protobuf:"fixed64,1,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
	MinValue             float64                   `protobuf:"fixed64,2,opt,name=min_value,json=minValue,proto3" json:"min_value,omitempty"`
	ScalingType          HyperparameterScalingType `protobuf:"varint,3,opt,name=scaling_type,json=scalingType,proto3,enum=flyteidl.plugins.sagemaker.HyperparameterScalingType" json:"scaling_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ContinuousParameterRange) Reset()         { *m = ContinuousParameterRange{} }
func (m *ContinuousParameterRange) String() string { return proto.CompactTextString(m) }
func (*ContinuousParameterRange) ProtoMessage()    {}
func (*ContinuousParameterRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87fa566d23b1e82, []int{0}
}

func (m *ContinuousParameterRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContinuousParameterRange.Unmarshal(m, b)
}
func (m *ContinuousParameterRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContinuousParameterRange.Marshal(b, m, deterministic)
}
func (m *ContinuousParameterRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContinuousParameterRange.Merge(m, src)
}
func (m *ContinuousParameterRange) XXX_Size() int {
	return xxx_messageInfo_ContinuousParameterRange.Size(m)
}
func (m *ContinuousParameterRange) XXX_DiscardUnknown() {
	xxx_messageInfo_ContinuousParameterRange.DiscardUnknown(m)
}

var xxx_messageInfo_ContinuousParameterRange proto.InternalMessageInfo

func (m *ContinuousParameterRange) GetMaxValue() float64 {
	if m != nil {
		return m.MaxValue
	}
	return 0
}

func (m *ContinuousParameterRange) GetMinValue() float64 {
	if m != nil {
		return m.MinValue
	}
	return 0
}

func (m *ContinuousParameterRange) GetScalingType() HyperparameterScalingType {
	if m != nil {
		return m.ScalingType
	}
	return HyperparameterScalingType_AUTO
}

type IntegerParameterRange struct {
	MaxValue             int64                     `protobuf:"varint,1,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
	MinValue             int64                     `protobuf:"varint,2,opt,name=min_value,json=minValue,proto3" json:"min_value,omitempty"`
	ScalingType          HyperparameterScalingType `protobuf:"varint,3,opt,name=scaling_type,json=scalingType,proto3,enum=flyteidl.plugins.sagemaker.HyperparameterScalingType" json:"scaling_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *IntegerParameterRange) Reset()         { *m = IntegerParameterRange{} }
func (m *IntegerParameterRange) String() string { return proto.CompactTextString(m) }
func (*IntegerParameterRange) ProtoMessage()    {}
func (*IntegerParameterRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87fa566d23b1e82, []int{1}
}

func (m *IntegerParameterRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegerParameterRange.Unmarshal(m, b)
}
func (m *IntegerParameterRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegerParameterRange.Marshal(b, m, deterministic)
}
func (m *IntegerParameterRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegerParameterRange.Merge(m, src)
}
func (m *IntegerParameterRange) XXX_Size() int {
	return xxx_messageInfo_IntegerParameterRange.Size(m)
}
func (m *IntegerParameterRange) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegerParameterRange.DiscardUnknown(m)
}

var xxx_messageInfo_IntegerParameterRange proto.InternalMessageInfo

func (m *IntegerParameterRange) GetMaxValue() int64 {
	if m != nil {
		return m.MaxValue
	}
	return 0
}

func (m *IntegerParameterRange) GetMinValue() int64 {
	if m != nil {
		return m.MinValue
	}
	return 0
}

func (m *IntegerParameterRange) GetScalingType() HyperparameterScalingType {
	if m != nil {
		return m.ScalingType
	}
	return HyperparameterScalingType_AUTO
}

type CategoricalParameterRange struct {
	Values               []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoricalParameterRange) Reset()         { *m = CategoricalParameterRange{} }
func (m *CategoricalParameterRange) String() string { return proto.CompactTextString(m) }
func (*CategoricalParameterRange) ProtoMessage()    {}
func (*CategoricalParameterRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87fa566d23b1e82, []int{2}
}

func (m *CategoricalParameterRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoricalParameterRange.Unmarshal(m, b)
}
func (m *CategoricalParameterRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoricalParameterRange.Marshal(b, m, deterministic)
}
func (m *CategoricalParameterRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoricalParameterRange.Merge(m, src)
}
func (m *CategoricalParameterRange) XXX_Size() int {
	return xxx_messageInfo_CategoricalParameterRange.Size(m)
}
func (m *CategoricalParameterRange) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoricalParameterRange.DiscardUnknown(m)
}

var xxx_messageInfo_CategoricalParameterRange proto.InternalMessageInfo

func (m *CategoricalParameterRange) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type ParameterRangeOneOf struct {
	// Types that are valid to be assigned to ParameterRangeType:
	//	*ParameterRangeOneOf_ContinuousParameterRange
	//	*ParameterRangeOneOf_IntegerParameterRange
	//	*ParameterRangeOneOf_CategoricalParameterRange
	ParameterRangeType   isParameterRangeOneOf_ParameterRangeType `protobuf_oneof:"parameter_range_type"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *ParameterRangeOneOf) Reset()         { *m = ParameterRangeOneOf{} }
func (m *ParameterRangeOneOf) String() string { return proto.CompactTextString(m) }
func (*ParameterRangeOneOf) ProtoMessage()    {}
func (*ParameterRangeOneOf) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87fa566d23b1e82, []int{3}
}

func (m *ParameterRangeOneOf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParameterRangeOneOf.Unmarshal(m, b)
}
func (m *ParameterRangeOneOf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParameterRangeOneOf.Marshal(b, m, deterministic)
}
func (m *ParameterRangeOneOf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParameterRangeOneOf.Merge(m, src)
}
func (m *ParameterRangeOneOf) XXX_Size() int {
	return xxx_messageInfo_ParameterRangeOneOf.Size(m)
}
func (m *ParameterRangeOneOf) XXX_DiscardUnknown() {
	xxx_messageInfo_ParameterRangeOneOf.DiscardUnknown(m)
}

var xxx_messageInfo_ParameterRangeOneOf proto.InternalMessageInfo

type isParameterRangeOneOf_ParameterRangeType interface {
	isParameterRangeOneOf_ParameterRangeType()
}

type ParameterRangeOneOf_ContinuousParameterRange struct {
	ContinuousParameterRange *ContinuousParameterRange `protobuf:"bytes,1,opt,name=continuous_parameter_range,json=continuousParameterRange,proto3,oneof"`
}

type ParameterRangeOneOf_IntegerParameterRange struct {
	IntegerParameterRange *IntegerParameterRange `protobuf:"bytes,2,opt,name=integer_parameter_range,json=integerParameterRange,proto3,oneof"`
}

type ParameterRangeOneOf_CategoricalParameterRange struct {
	CategoricalParameterRange *CategoricalParameterRange `protobuf:"bytes,3,opt,name=categorical_parameter_range,json=categoricalParameterRange,proto3,oneof"`
}

func (*ParameterRangeOneOf_ContinuousParameterRange) isParameterRangeOneOf_ParameterRangeType() {}

func (*ParameterRangeOneOf_IntegerParameterRange) isParameterRangeOneOf_ParameterRangeType() {}

func (*ParameterRangeOneOf_CategoricalParameterRange) isParameterRangeOneOf_ParameterRangeType() {}

func (m *ParameterRangeOneOf) GetParameterRangeType() isParameterRangeOneOf_ParameterRangeType {
	if m != nil {
		return m.ParameterRangeType
	}
	return nil
}

func (m *ParameterRangeOneOf) GetContinuousParameterRange() *ContinuousParameterRange {
	if x, ok := m.GetParameterRangeType().(*ParameterRangeOneOf_ContinuousParameterRange); ok {
		return x.ContinuousParameterRange
	}
	return nil
}

func (m *ParameterRangeOneOf) GetIntegerParameterRange() *IntegerParameterRange {
	if x, ok := m.GetParameterRangeType().(*ParameterRangeOneOf_IntegerParameterRange); ok {
		return x.IntegerParameterRange
	}
	return nil
}

func (m *ParameterRangeOneOf) GetCategoricalParameterRange() *CategoricalParameterRange {
	if x, ok := m.GetParameterRangeType().(*ParameterRangeOneOf_CategoricalParameterRange); ok {
		return x.CategoricalParameterRange
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ParameterRangeOneOf) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ParameterRangeOneOf_ContinuousParameterRange)(nil),
		(*ParameterRangeOneOf_IntegerParameterRange)(nil),
		(*ParameterRangeOneOf_CategoricalParameterRange)(nil),
	}
}

type ParameterRanges struct {
	ParameterRangeMap    map[string]*ParameterRangeOneOf `protobuf:"bytes,1,rep,name=parameter_range_map,json=parameterRangeMap,proto3" json:"parameter_range_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ParameterRanges) Reset()         { *m = ParameterRanges{} }
func (m *ParameterRanges) String() string { return proto.CompactTextString(m) }
func (*ParameterRanges) ProtoMessage()    {}
func (*ParameterRanges) Descriptor() ([]byte, []int) {
	return fileDescriptor_d87fa566d23b1e82, []int{4}
}

func (m *ParameterRanges) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParameterRanges.Unmarshal(m, b)
}
func (m *ParameterRanges) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParameterRanges.Marshal(b, m, deterministic)
}
func (m *ParameterRanges) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParameterRanges.Merge(m, src)
}
func (m *ParameterRanges) XXX_Size() int {
	return xxx_messageInfo_ParameterRanges.Size(m)
}
func (m *ParameterRanges) XXX_DiscardUnknown() {
	xxx_messageInfo_ParameterRanges.DiscardUnknown(m)
}

var xxx_messageInfo_ParameterRanges proto.InternalMessageInfo

func (m *ParameterRanges) GetParameterRangeMap() map[string]*ParameterRangeOneOf {
	if m != nil {
		return m.ParameterRangeMap
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl.plugins.sagemaker.HyperparameterScalingType", HyperparameterScalingType_name, HyperparameterScalingType_value)
	proto.RegisterType((*ContinuousParameterRange)(nil), "flyteidl.plugins.sagemaker.ContinuousParameterRange")
	proto.RegisterType((*IntegerParameterRange)(nil), "flyteidl.plugins.sagemaker.IntegerParameterRange")
	proto.RegisterType((*CategoricalParameterRange)(nil), "flyteidl.plugins.sagemaker.CategoricalParameterRange")
	proto.RegisterType((*ParameterRangeOneOf)(nil), "flyteidl.plugins.sagemaker.ParameterRangeOneOf")
	proto.RegisterType((*ParameterRanges)(nil), "flyteidl.plugins.sagemaker.ParameterRanges")
	proto.RegisterMapType((map[string]*ParameterRangeOneOf)(nil), "flyteidl.plugins.sagemaker.ParameterRanges.ParameterRangeMapEntry")
}

func init() {
	proto.RegisterFile("flyteidl/plugins/sagemaker/parameterranges.proto", fileDescriptor_d87fa566d23b1e82)
}

var fileDescriptor_d87fa566d23b1e82 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xd1, 0x6e, 0x12, 0x41,
	0x14, 0x65, 0x58, 0x6d, 0xca, 0xc5, 0x58, 0x9c, 0x5a, 0x04, 0xfa, 0x42, 0x78, 0x22, 0x26, 0xee,
	0x2a, 0xd8, 0xc4, 0xf8, 0x06, 0x64, 0x23, 0x24, 0xad, 0x98, 0x29, 0x36, 0xa6, 0x2f, 0x64, 0x58,
	0x87, 0x75, 0xc2, 0xee, 0xec, 0x64, 0x76, 0x56, 0xbb, 0x9f, 0xe0, 0xb7, 0x98, 0xf8, 0x85, 0x3e,
	0x18, 0x06, 0x0a, 0x76, 0xcb, 0xa2, 0x4f, 0xbe, 0xed, 0x9d, 0x7b, 0xcf, 0x9e, 0x33, 0x73, 0x4e,
	0x2e, 0xbc, 0x9c, 0x07, 0xa9, 0x66, 0xfc, 0x73, 0xe0, 0xc8, 0x20, 0xf1, 0xb9, 0x88, 0x9d, 0x98,
	0xfa, 0x2c, 0xa4, 0x0b, 0xa6, 0x1c, 0x49, 0x15, 0x0d, 0x99, 0x66, 0x4a, 0x51, 0xe1, 0xb3, 0xd8,
	0x96, 0x2a, 0xd2, 0x11, 0x6e, 0xdc, 0x22, 0xec, 0x35, 0xc2, 0xde, 0x20, 0x5a, 0x3f, 0x11, 0xd4,
	0x06, 0x91, 0xd0, 0x5c, 0x24, 0x51, 0x12, 0x7f, 0xb8, 0xc5, 0x93, 0x25, 0x1e, 0x9f, 0x42, 0x29,
	0xa4, 0x37, 0xd3, 0xaf, 0x34, 0x48, 0x58, 0x0d, 0x35, 0x51, 0x1b, 0x91, 0xc3, 0x90, 0xde, 0x5c,
	0x2d, 0x6b, 0xd3, 0xe4, 0x62, 0xdd, 0x2c, 0xae, 0x9b, 0x5c, 0xac, 0x9a, 0x9f, 0xe0, 0x51, 0xec,
	0xd1, 0x80, 0x0b, 0x7f, 0xaa, 0x53, 0xc9, 0x6a, 0x56, 0x13, 0xb5, 0x1f, 0x77, 0xce, 0xec, 0x7c,
	0x25, 0xf6, 0x30, 0x95, 0x4c, 0x6d, 0x2e, 0x70, 0xb9, 0x42, 0x4f, 0x52, 0xc9, 0x48, 0x39, 0xde,
	0x16, 0xad, 0x1f, 0x08, 0x4e, 0x46, 0x42, 0x33, 0x9f, 0xa9, 0xbf, 0xa9, 0xb5, 0xf6, 0xa9, 0xb5,
	0xfe, 0x8b, 0xda, 0x2e, 0xd4, 0x07, 0x54, 0x33, 0x3f, 0x52, 0xdc, 0xa3, 0x41, 0x46, 0x70, 0x15,
	0x0e, 0x8c, 0x9e, 0xb8, 0x86, 0x9a, 0x56, 0xbb, 0x44, 0xd6, 0x55, 0xeb, 0xbb, 0x05, 0xc7, 0x77,
	0x47, 0xc7, 0x82, 0x8d, 0xe7, 0x58, 0x43, 0xc3, 0xdb, 0x58, 0x35, 0xdd, 0x90, 0x4f, 0x8d, 0xd9,
	0xe6, 0xc6, 0xe5, 0xce, 0xeb, 0x7d, 0xa2, 0xf3, 0x8c, 0x1e, 0x16, 0x48, 0xcd, 0xcb, 0x0b, 0xc1,
	0x02, 0x9e, 0xf1, 0xd5, 0x7b, 0xdf, 0xa3, 0x2c, 0x1a, 0xca, 0x57, 0xfb, 0x28, 0x77, 0x5a, 0x35,
	0x2c, 0x90, 0x13, 0xbe, 0xd3, 0xc3, 0x6f, 0x70, 0xea, 0x6d, 0xdf, 0xeb, 0x1e, 0xa1, 0x65, 0x08,
	0xf7, 0x1a, 0x93, 0xfb, 0xdc, 0xc3, 0x02, 0xa9, 0x7b, 0x79, 0xcd, 0x7e, 0x15, 0x9e, 0x66, 0xc8,
	0x4c, 0x14, 0x5a, 0xbf, 0x10, 0x1c, 0xdd, 0x1d, 0x8d, 0xb1, 0x82, 0xe3, 0xec, 0x6c, 0x48, 0xa5,
	0x31, 0xb1, 0xdc, 0xe9, 0xef, 0x13, 0x97, 0xf9, 0x53, 0xa6, 0xbe, 0xa0, 0xd2, 0x15, 0x5a, 0xa5,
	0xe4, 0x89, 0xcc, 0x9e, 0x37, 0x12, 0xa8, 0xee, 0x1e, 0xc6, 0x15, 0xb0, 0x16, 0x2c, 0x35, 0xf6,
	0x97, 0xc8, 0xf2, 0x13, 0xbb, 0xf0, 0x70, 0x9b, 0xf3, 0x72, 0xc7, 0xf9, 0x77, 0x45, 0x26, 0x67,
	0x64, 0x85, 0x7e, 0x5b, 0x7c, 0x83, 0x9e, 0x5f, 0x43, 0x3d, 0x37, 0xe9, 0xf8, 0x10, 0x1e, 0xf4,
	0x3e, 0x4e, 0xc6, 0x95, 0x02, 0x06, 0x38, 0x38, 0x1f, 0xbd, 0x77, 0x7b, 0xa4, 0x82, 0xf0, 0x11,
	0x94, 0xcf, 0xc7, 0xef, 0x7a, 0x64, 0x34, 0x19, 0x5e, 0x8c, 0x06, 0x95, 0x22, 0xae, 0x02, 0x26,
	0xee, 0x95, 0x4b, 0x2e, 0xdd, 0x3f, 0xcf, 0xad, 0xfe, 0xd9, 0x75, 0xd7, 0xe7, 0xfa, 0x4b, 0x32,
	0xb3, 0xbd, 0x28, 0x74, 0x82, 0x74, 0xae, 0x9d, 0xcd, 0x6a, 0xf3, 0x99, 0x70, 0xe4, 0xec, 0x85,
	0x1f, 0x39, 0xd9, 0x6d, 0x37, 0x3b, 0x30, 0x4b, 0xad, 0xfb, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xaf,
	0x48, 0xf7, 0xbe, 0x08, 0x05, 0x00, 0x00,
}
