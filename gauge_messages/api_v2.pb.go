// Code generated by protoc-gen-go.
// source: api_v2.proto
// DO NOT EDIT!

package gauge_messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Types of log level that gauge supports
type ExecutionRequest_LogLevel int32

const (
	ExecutionRequest_INFO    ExecutionRequest_LogLevel = 0
	ExecutionRequest_DEBUG   ExecutionRequest_LogLevel = 1
	ExecutionRequest_WARNING ExecutionRequest_LogLevel = 2
	ExecutionRequest_ERROR   ExecutionRequest_LogLevel = 3
)

var ExecutionRequest_LogLevel_name = map[int32]string{
	0: "INFO",
	1: "DEBUG",
	2: "WARNING",
	3: "ERROR",
}
var ExecutionRequest_LogLevel_value = map[string]int32{
	"INFO":    0,
	"DEBUG":   1,
	"WARNING": 2,
	"ERROR":   3,
}

func (x ExecutionRequest_LogLevel) Enum() *ExecutionRequest_LogLevel {
	p := new(ExecutionRequest_LogLevel)
	*p = x
	return p
}
func (x ExecutionRequest_LogLevel) String() string {
	return proto.EnumName(ExecutionRequest_LogLevel_name, int32(x))
}
func (x *ExecutionRequest_LogLevel) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ExecutionRequest_LogLevel_value, data, "ExecutionRequest_LogLevel")
	if err != nil {
		return err
	}
	*x = ExecutionRequest_LogLevel(value)
	return nil
}
func (ExecutionRequest_LogLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

// Strategies for parallelization
type ExecutionRequest_Strategy int32

const (
	ExecutionRequest_LAZY  ExecutionRequest_Strategy = 0
	ExecutionRequest_EAGER ExecutionRequest_Strategy = 1
)

var ExecutionRequest_Strategy_name = map[int32]string{
	0: "LAZY",
	1: "EAGER",
}
var ExecutionRequest_Strategy_value = map[string]int32{
	"LAZY":  0,
	"EAGER": 1,
}

func (x ExecutionRequest_Strategy) Enum() *ExecutionRequest_Strategy {
	p := new(ExecutionRequest_Strategy)
	*p = x
	return p
}
func (x ExecutionRequest_Strategy) String() string {
	return proto.EnumName(ExecutionRequest_Strategy_name, int32(x))
}
func (x *ExecutionRequest_Strategy) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ExecutionRequest_Strategy_value, data, "ExecutionRequest_Strategy")
	if err != nil {
		return err
	}
	*x = ExecutionRequest_Strategy(value)
	return nil
}
func (ExecutionRequest_Strategy) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 1} }

// Types of Execution result
type ExecutionResponse_ExecutionResponseType int32

const (
	ExecutionResponse_ScenarioResult           ExecutionResponse_ExecutionResponseType = 0
	ExecutionResponse_BeforeSuiteHookResult    ExecutionResponse_ExecutionResponseType = 1
	ExecutionResponse_AfterSuiteHookResult     ExecutionResponse_ExecutionResponseType = 2
	ExecutionResponse_BeforeSpecHookResult     ExecutionResponse_ExecutionResponseType = 3
	ExecutionResponse_AfterSpecHookResult      ExecutionResponse_ExecutionResponseType = 4
	ExecutionResponse_SuiteDataStoreInitResult ExecutionResponse_ExecutionResponseType = 5
	ExecutionResponse_SpecDataStoreInitResult  ExecutionResponse_ExecutionResponseType = 6
	ExecutionResponse_ValidationResult         ExecutionResponse_ExecutionResponseType = 7
)

var ExecutionResponse_ExecutionResponseType_name = map[int32]string{
	0: "ScenarioResult",
	1: "BeforeSuiteHookResult",
	2: "AfterSuiteHookResult",
	3: "BeforeSpecHookResult",
	4: "AfterSpecHookResult",
	5: "SuiteDataStoreInitResult",
	6: "SpecDataStoreInitResult",
	7: "ValidationResult",
}
var ExecutionResponse_ExecutionResponseType_value = map[string]int32{
	"ScenarioResult":           0,
	"BeforeSuiteHookResult":    1,
	"AfterSuiteHookResult":     2,
	"BeforeSpecHookResult":     3,
	"AfterSpecHookResult":      4,
	"SuiteDataStoreInitResult": 5,
	"SpecDataStoreInitResult":  6,
	"ValidationResult":         7,
}

func (x ExecutionResponse_ExecutionResponseType) Enum() *ExecutionResponse_ExecutionResponseType {
	p := new(ExecutionResponse_ExecutionResponseType)
	*p = x
	return p
}
func (x ExecutionResponse_ExecutionResponseType) String() string {
	return proto.EnumName(ExecutionResponse_ExecutionResponseType_name, int32(x))
}
func (x *ExecutionResponse_ExecutionResponseType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ExecutionResponse_ExecutionResponseType_value, data, "ExecutionResponse_ExecutionResponseType")
	if err != nil {
		return err
	}
	*x = ExecutionResponse_ExecutionResponseType(value)
	return nil
}
func (ExecutionResponse_ExecutionResponseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1, 0}
}

// Specifies the execution status
type ExecutionResponse_Status int32

const (
	ExecutionResponse_PASSED  ExecutionResponse_Status = 0
	ExecutionResponse_FAILED  ExecutionResponse_Status = 1
	ExecutionResponse_SKIPPED ExecutionResponse_Status = 2
)

var ExecutionResponse_Status_name = map[int32]string{
	0: "PASSED",
	1: "FAILED",
	2: "SKIPPED",
}
var ExecutionResponse_Status_value = map[string]int32{
	"PASSED":  0,
	"FAILED":  1,
	"SKIPPED": 2,
}

func (x ExecutionResponse_Status) Enum() *ExecutionResponse_Status {
	p := new(ExecutionResponse_Status)
	*p = x
	return p
}
func (x ExecutionResponse_Status) String() string {
	return proto.EnumName(ExecutionResponse_Status_name, int32(x))
}
func (x *ExecutionResponse_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ExecutionResponse_Status_value, data, "ExecutionResponse_Status")
	if err != nil {
		return err
	}
	*x = ExecutionResponse_Status(value)
	return nil
}
func (ExecutionResponse_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 1} }

// ExecutionRequest defines the structure of requests that the API's consumers can send to request execution of specs/scenarios
type ExecutionRequest struct {
	// Each ExecutionRequest can ask to execute multiple spec files or multiple scenarios in a spec or a directory or all
	Specs []string `protobuf:"bytes,1,rep,name=specs" json:"specs,omitempty"`
	// Tag expression to filter specs and scenarios during execution. Default: ""
	Tags *string `protobuf:"bytes,2,opt,name=tags" json:"tags,omitempty"`
	// The working directory for gauge.
	WorkingDir *string `protobuf:"bytes,3,opt,name=workingDir" json:"workingDir,omitempty"`
	// Environment to choose for gauge.
	Environment *string                    `protobuf:"bytes,4,opt,name=environment" json:"environment,omitempty"`
	LogLevel    *ExecutionRequest_LogLevel `protobuf:"varint,5,opt,name=logLevel,enum=gauge.messages.ExecutionRequest_LogLevel" json:"logLevel,omitempty"`
	// Whether to run gauge in parallel mode.
	IsParallel *bool `protobuf:"varint,6,opt,name=isParallel" json:"isParallel,omitempty"`
	// If isParallel is true, this specifies how many parallel streams to run.
	ParallelStreams *int32 `protobuf:"varint,7,opt,name=parallelStreams" json:"parallelStreams,omitempty"`
	// Specify which group of specs to execute, based on number of parallel streams. Works only when strategy is EAGER.
	Group *int32 `protobuf:"varint,8,opt,name=group" json:"group,omitempty"`
	// Toggles simple console reporting
	SimpleConsole *bool `protobuf:"varint,9,opt,name=simpleConsole" json:"simpleConsole,omitempty"`
	// If true, sorts execution of specifications alphabetically
	Sort     *bool                      `protobuf:"varint,10,opt,name=sort" json:"sort,omitempty"`
	Strategy *ExecutionRequest_Strategy `protobuf:"varint,11,opt,name=strategy,enum=gauge.messages.ExecutionRequest_Strategy" json:"strategy,omitempty"`
	// Specify against which rows of datatable the scenarios should be executed
	TableRows        *string `protobuf:"bytes,12,opt,name=tableRows" json:"tableRows,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ExecutionRequest) Reset()                    { *m = ExecutionRequest{} }
func (m *ExecutionRequest) String() string            { return proto.CompactTextString(m) }
func (*ExecutionRequest) ProtoMessage()               {}
func (*ExecutionRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ExecutionRequest) GetSpecs() []string {
	if m != nil {
		return m.Specs
	}
	return nil
}

func (m *ExecutionRequest) GetTags() string {
	if m != nil && m.Tags != nil {
		return *m.Tags
	}
	return ""
}

func (m *ExecutionRequest) GetWorkingDir() string {
	if m != nil && m.WorkingDir != nil {
		return *m.WorkingDir
	}
	return ""
}

func (m *ExecutionRequest) GetEnvironment() string {
	if m != nil && m.Environment != nil {
		return *m.Environment
	}
	return ""
}

func (m *ExecutionRequest) GetLogLevel() ExecutionRequest_LogLevel {
	if m != nil && m.LogLevel != nil {
		return *m.LogLevel
	}
	return ExecutionRequest_INFO
}

func (m *ExecutionRequest) GetIsParallel() bool {
	if m != nil && m.IsParallel != nil {
		return *m.IsParallel
	}
	return false
}

func (m *ExecutionRequest) GetParallelStreams() int32 {
	if m != nil && m.ParallelStreams != nil {
		return *m.ParallelStreams
	}
	return 0
}

func (m *ExecutionRequest) GetGroup() int32 {
	if m != nil && m.Group != nil {
		return *m.Group
	}
	return 0
}

func (m *ExecutionRequest) GetSimpleConsole() bool {
	if m != nil && m.SimpleConsole != nil {
		return *m.SimpleConsole
	}
	return false
}

func (m *ExecutionRequest) GetSort() bool {
	if m != nil && m.Sort != nil {
		return *m.Sort
	}
	return false
}

func (m *ExecutionRequest) GetStrategy() ExecutionRequest_Strategy {
	if m != nil && m.Strategy != nil {
		return *m.Strategy
	}
	return ExecutionRequest_LAZY
}

func (m *ExecutionRequest) GetTableRows() string {
	if m != nil && m.TableRows != nil {
		return *m.TableRows
	}
	return ""
}

// ExecutionResponse defines the structure of response for ExecutionRequest message
type ExecutionResponse struct {
	Type *ExecutionResponse_ExecutionResponseType `protobuf:"varint,1,req,name=type,enum=gauge.messages.ExecutionResponse_ExecutionResponseType" json:"type,omitempty"`
	// An identifier for the current execution result
	ID     *string                   `protobuf:"bytes,2,opt,name=ID,json=iD" json:"ID,omitempty"`
	Status *ExecutionResponse_Status `protobuf:"varint,3,opt,name=status,enum=gauge.messages.ExecutionResponse_Status" json:"status,omitempty"`
	// Contains the Execution error and its details
	Error         []*ExecutionResponse_ExecutionError `protobuf:"bytes,4,rep,name=error" json:"error,omitempty"`
	ExecutionTime *int64                              `protobuf:"varint,5,opt,name=executionTime" json:"executionTime,omitempty"`
	// Contains the console output messages
	Stdout           *string `protobuf:"bytes,6,opt,name=stdout" json:"stdout,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ExecutionResponse) Reset()                    { *m = ExecutionResponse{} }
func (m *ExecutionResponse) String() string            { return proto.CompactTextString(m) }
func (*ExecutionResponse) ProtoMessage()               {}
func (*ExecutionResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ExecutionResponse) GetType() ExecutionResponse_ExecutionResponseType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ExecutionResponse_ScenarioResult
}

func (m *ExecutionResponse) GetID() string {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return ""
}

func (m *ExecutionResponse) GetStatus() ExecutionResponse_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ExecutionResponse_PASSED
}

func (m *ExecutionResponse) GetError() []*ExecutionResponse_ExecutionError {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ExecutionResponse) GetExecutionTime() int64 {
	if m != nil && m.ExecutionTime != nil {
		return *m.ExecutionTime
	}
	return 0
}

func (m *ExecutionResponse) GetStdout() string {
	if m != nil && m.Stdout != nil {
		return *m.Stdout
	}
	return ""
}

// Specifies the execution time
// ExecutionError represents the failure during execution along with details of failure
type ExecutionResponse_ExecutionError struct {
	// Stacktrace from the failure
	StackTrace *string `protobuf:"bytes,1,opt,name=stackTrace" json:"stackTrace,omitempty"`
	// Error message from the failure
	ErrorMessage *string `protobuf:"bytes,2,opt,name=errorMessage" json:"errorMessage,omitempty"`
	// Byte array holding the screenshot taken at the time of failure.
	ScreenShot       []byte `protobuf:"bytes,3,opt,name=screenShot" json:"screenShot,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ExecutionResponse_ExecutionError) Reset()         { *m = ExecutionResponse_ExecutionError{} }
func (m *ExecutionResponse_ExecutionError) String() string { return proto.CompactTextString(m) }
func (*ExecutionResponse_ExecutionError) ProtoMessage()    {}
func (*ExecutionResponse_ExecutionError) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1, 0}
}

func (m *ExecutionResponse_ExecutionError) GetStackTrace() string {
	if m != nil && m.StackTrace != nil {
		return *m.StackTrace
	}
	return ""
}

func (m *ExecutionResponse_ExecutionError) GetErrorMessage() string {
	if m != nil && m.ErrorMessage != nil {
		return *m.ErrorMessage
	}
	return ""
}

func (m *ExecutionResponse_ExecutionError) GetScreenShot() []byte {
	if m != nil {
		return m.ScreenShot
	}
	return nil
}

func init() {
	proto.RegisterType((*ExecutionRequest)(nil), "gauge.messages.ExecutionRequest")
	proto.RegisterType((*ExecutionResponse)(nil), "gauge.messages.ExecutionResponse")
	proto.RegisterType((*ExecutionResponse_ExecutionError)(nil), "gauge.messages.ExecutionResponse.ExecutionError")
	proto.RegisterEnum("gauge.messages.ExecutionRequest_LogLevel", ExecutionRequest_LogLevel_name, ExecutionRequest_LogLevel_value)
	proto.RegisterEnum("gauge.messages.ExecutionRequest_Strategy", ExecutionRequest_Strategy_name, ExecutionRequest_Strategy_value)
	proto.RegisterEnum("gauge.messages.ExecutionResponse_ExecutionResponseType", ExecutionResponse_ExecutionResponseType_name, ExecutionResponse_ExecutionResponseType_value)
	proto.RegisterEnum("gauge.messages.ExecutionResponse_Status", ExecutionResponse_Status_name, ExecutionResponse_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Execution service

type ExecutionClient interface {
	// Bind RPC method
	Execute(ctx context.Context, in *ExecutionRequest, opts ...grpc.CallOption) (Execution_ExecuteClient, error)
}

type executionClient struct {
	cc *grpc.ClientConn
}

func NewExecutionClient(cc *grpc.ClientConn) ExecutionClient {
	return &executionClient{cc}
}

func (c *executionClient) Execute(ctx context.Context, in *ExecutionRequest, opts ...grpc.CallOption) (Execution_ExecuteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Execution_serviceDesc.Streams[0], c.cc, "/gauge.messages.Execution/execute", opts...)
	if err != nil {
		return nil, err
	}
	x := &executionExecuteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Execution_ExecuteClient interface {
	Recv() (*ExecutionResponse, error)
	grpc.ClientStream
}

type executionExecuteClient struct {
	grpc.ClientStream
}

func (x *executionExecuteClient) Recv() (*ExecutionResponse, error) {
	m := new(ExecutionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Execution service

type ExecutionServer interface {
	// Bind RPC method
	Execute(*ExecutionRequest, Execution_ExecuteServer) error
}

func RegisterExecutionServer(s *grpc.Server, srv ExecutionServer) {
	s.RegisterService(&_Execution_serviceDesc, srv)
}

func _Execution_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExecutionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExecutionServer).Execute(m, &executionExecuteServer{stream})
}

type Execution_ExecuteServer interface {
	Send(*ExecutionResponse) error
	grpc.ServerStream
}

type executionExecuteServer struct {
	grpc.ServerStream
}

func (x *executionExecuteServer) Send(m *ExecutionResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Execution_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gauge.messages.Execution",
	HandlerType: (*ExecutionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "execute",
			Handler:       _Execution_Execute_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("api_v2.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdf, 0x52, 0xd3, 0x4e,
	0x14, 0xa6, 0xff, 0xdb, 0xd3, 0xfe, 0xfa, 0xab, 0x2b, 0x48, 0x44, 0x47, 0x6b, 0xc6, 0x8b, 0x7a,
	0x61, 0x87, 0xe9, 0x0d, 0xb7, 0x16, 0x1b, 0xb0, 0x43, 0x85, 0xce, 0x06, 0x75, 0xf4, 0x86, 0x59,
	0xcb, 0x21, 0x66, 0x48, 0xb3, 0x71, 0x77, 0x03, 0xf2, 0x32, 0x3e, 0x8d, 0x8f, 0xe2, 0x83, 0xb8,
	0xbb, 0x09, 0xd0, 0x22, 0x23, 0x7a, 0x77, 0xf6, 0xfb, 0xf2, 0x9d, 0x3d, 0x7f, 0xbe, 0x0d, 0xb4,
	0x58, 0x12, 0x1e, 0x9d, 0x0d, 0xfa, 0x89, 0xe0, 0x8a, 0x93, 0x76, 0xc0, 0xd2, 0x00, 0xfb, 0x73,
	0x94, 0x92, 0x05, 0x28, 0xdd, 0xef, 0x65, 0xe8, 0x78, 0xdf, 0x70, 0x96, 0xaa, 0x90, 0xc7, 0x14,
	0xbf, 0xa6, 0x28, 0x15, 0x59, 0x85, 0x8a, 0x4c, 0x70, 0x26, 0x9d, 0x42, 0xb7, 0xd4, 0x6b, 0xd0,
	0xec, 0x40, 0x08, 0x94, 0x15, 0x0b, 0xa4, 0x53, 0xec, 0x16, 0x34, 0x68, 0x63, 0xf2, 0x04, 0xe0,
	0x9c, 0x8b, 0xd3, 0x30, 0x0e, 0x46, 0xa1, 0x70, 0x4a, 0x96, 0x59, 0x40, 0x48, 0x17, 0x9a, 0x18,
	0x9f, 0x85, 0x82, 0xc7, 0x73, 0x8c, 0x95, 0x53, 0xb6, 0x1f, 0x2c, 0x42, 0xc4, 0x83, 0x7a, 0xc4,
	0x83, 0x09, 0x9e, 0x61, 0xe4, 0x54, 0x34, 0xdd, 0x1e, 0xbc, 0xe8, 0x2f, 0xd7, 0xd8, 0xbf, 0x59,
	0x5f, 0x7f, 0x92, 0x0b, 0xe8, 0x95, 0xd4, 0x14, 0x12, 0xca, 0x29, 0x13, 0x2c, 0x8a, 0x74, 0xa2,
	0xaa, 0x4e, 0x54, 0xa7, 0x0b, 0x08, 0xe9, 0xc1, 0xff, 0x49, 0x1e, 0xfb, 0x4a, 0x20, 0x9b, 0x4b,
	0xa7, 0xa6, 0x3f, 0xaa, 0xd0, 0x9b, 0xb0, 0x69, 0x3e, 0x10, 0x3c, 0x4d, 0x9c, 0xba, 0xe5, 0xb3,
	0x03, 0x79, 0x0e, 0xff, 0xc9, 0x70, 0x9e, 0x44, 0xf8, 0x9a, 0xc7, 0x92, 0x47, 0xe8, 0x34, 0xec,
	0x15, 0xcb, 0xa0, 0x19, 0x91, 0xe4, 0x42, 0x39, 0x60, 0x49, 0x1b, 0x9b, 0x06, 0xa5, 0x12, 0x4c,
	0x61, 0x70, 0xe1, 0x34, 0xff, 0xb2, 0x41, 0x3f, 0x17, 0xd0, 0x2b, 0x29, 0x79, 0x0c, 0x0d, 0xc5,
	0x3e, 0x47, 0x48, 0xf9, 0xb9, 0x74, 0x5a, 0x76, 0x8e, 0xd7, 0x80, 0xbb, 0x05, 0xf5, 0xcb, 0xa1,
	0x90, 0x3a, 0x94, 0xc7, 0xfb, 0x3b, 0x07, 0x9d, 0x15, 0xd2, 0x80, 0xca, 0xc8, 0xdb, 0x7e, 0xb7,
	0xdb, 0x29, 0x90, 0x26, 0xd4, 0x3e, 0x0c, 0xe9, 0xfe, 0x78, 0x7f, 0xb7, 0x53, 0x34, 0xb8, 0x47,
	0xe9, 0x01, 0xed, 0x94, 0xdc, 0xa7, 0x50, 0xbf, 0xbc, 0xcc, 0x08, 0x27, 0xc3, 0x4f, 0x1f, 0x33,
	0xa1, 0x37, 0xdc, 0xf5, 0x68, 0xa7, 0xe0, 0xfe, 0xa8, 0xc0, 0xbd, 0x85, 0xfa, 0x64, 0xa2, 0x5b,
	0x45, 0xb2, 0xa7, 0xbd, 0x70, 0x91, 0xa0, 0x36, 0x48, 0x51, 0x37, 0xb4, 0xf5, 0x87, 0x86, 0x32,
	0xc1, 0xef, 0xc8, 0xa1, 0x96, 0x53, 0x9b, 0x84, 0xb4, 0xa1, 0x38, 0x1e, 0xe5, 0xb6, 0x2a, 0x86,
	0x23, 0xf2, 0x0a, 0xaa, 0x52, 0x31, 0x95, 0x4a, 0x6b, 0xa8, 0xf6, 0xa0, 0x77, 0x77, 0x7a, 0xdf,
	0x7e, 0x4f, 0x73, 0x1d, 0xd9, 0x81, 0x0a, 0x0a, 0xc1, 0x85, 0x36, 0x5c, 0xa9, 0xd7, 0x1c, 0x6c,
	0xfe, 0x43, 0x7d, 0x9e, 0xd1, 0xd1, 0x4c, 0x6e, 0xb6, 0x8e, 0x97, 0xc4, 0x61, 0x38, 0x47, 0xeb,
	0xd0, 0x12, 0x5d, 0x06, 0xc9, 0x03, 0x53, 0xef, 0x31, 0x4f, 0x95, 0xf5, 0x5d, 0x83, 0xe6, 0xa7,
	0x0d, 0x05, 0xed, 0xe5, 0xb4, 0xc6, 0xa5, 0xba, 0xc2, 0xd9, 0xe9, 0xa1, 0x60, 0x33, 0x33, 0x3c,
	0xfb, 0x5c, 0xae, 0x11, 0xe2, 0x42, 0xcb, 0x5e, 0xfc, 0x36, 0x2b, 0x34, 0x9f, 0xc9, 0x12, 0x66,
	0x73, 0xcc, 0x04, 0x62, 0xec, 0x7f, 0xe1, 0xca, 0x4e, 0xa8, 0x45, 0x17, 0x10, 0xf7, 0x67, 0x01,
	0xd6, 0x6e, 0x9d, 0xb6, 0x76, 0x67, 0xdb, 0x9f, 0x61, 0xcc, 0x44, 0xc8, 0x35, 0x9e, 0x46, 0x4a,
	0x6f, 0xfa, 0x21, 0xac, 0x6d, 0xe3, 0x09, 0x17, 0xe8, 0xa7, 0xa1, 0xc2, 0x37, 0x9c, 0x9f, 0xe6,
	0x54, 0x81, 0x38, 0xb0, 0x3a, 0x3c, 0x51, 0x28, 0x6e, 0x32, 0x45, 0xc3, 0xe4, 0x22, 0xfd, 0x63,
	0x58, 0x60, 0x4a, 0x64, 0x1d, 0xee, 0x67, 0x9a, 0x65, 0xa2, 0xac, 0xed, 0xeb, 0xd8, 0x3c, 0x23,
	0xa6, 0x98, 0xaf, 0xb4, 0x74, 0x1c, 0x87, 0x2a, 0x67, 0x2b, 0xe4, 0x11, 0xac, 0x1b, 0xc5, 0x6d,
	0x64, 0x55, 0x3f, 0xc8, 0xce, 0x7b, 0x16, 0x85, 0xc7, 0x2c, 0x6f, 0xc8, 0xa0, 0x35, 0xf7, 0x25,
	0x54, 0xb3, 0xa5, 0x13, 0x80, 0xea, 0x74, 0xe8, 0xfb, 0xde, 0x48, 0xb7, 0xa3, 0xe3, 0x9d, 0xe1,
	0x78, 0xa2, 0x63, 0x6b, 0x79, 0x7f, 0x6f, 0x3c, 0x9d, 0xea, 0x43, 0x71, 0x70, 0x04, 0x8d, 0xab,
	0xa1, 0x10, 0x0a, 0xb5, 0x6c, 0x83, 0x48, 0xba, 0x77, 0xbd, 0xc5, 0x8d, 0x67, 0x77, 0x9a, 0xc7,
	0x5d, 0xd9, 0x2c, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x64, 0x94, 0x25, 0x4c, 0x67, 0x05, 0x00,
	0x00,
}