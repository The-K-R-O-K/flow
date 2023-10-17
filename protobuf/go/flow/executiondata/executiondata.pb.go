// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flow/executiondata/executiondata.proto

package access

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	entities "github.com/onflow/flow/protobuf/go/flow/entities"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// The request for GetExecutionDataByBlockID
type GetExecutionDataByBlockIDRequest struct {
	// Block ID of the block to get execution data for.
	BlockId []byte `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	// Preferred event encoding version of the block events payload.
	// Possible variants:
	// 1. CCF
	// 2. JSON-CDC
	// 3. DEFAULT, which is JSON-CDC
	EventEncodingVersion entities.EventEncodingVersion `protobuf:"varint,2,opt,name=event_encoding_version,json=eventEncodingVersion,proto3,enum=flow.entities.EventEncodingVersion" json:"event_encoding_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *GetExecutionDataByBlockIDRequest) Reset()         { *m = GetExecutionDataByBlockIDRequest{} }
func (m *GetExecutionDataByBlockIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetExecutionDataByBlockIDRequest) ProtoMessage()    {}
func (*GetExecutionDataByBlockIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da469632570513fb, []int{0}
}

func (m *GetExecutionDataByBlockIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetExecutionDataByBlockIDRequest.Unmarshal(m, b)
}
func (m *GetExecutionDataByBlockIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetExecutionDataByBlockIDRequest.Marshal(b, m, deterministic)
}
func (m *GetExecutionDataByBlockIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetExecutionDataByBlockIDRequest.Merge(m, src)
}
func (m *GetExecutionDataByBlockIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetExecutionDataByBlockIDRequest.Size(m)
}
func (m *GetExecutionDataByBlockIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetExecutionDataByBlockIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetExecutionDataByBlockIDRequest proto.InternalMessageInfo

func (m *GetExecutionDataByBlockIDRequest) GetBlockId() []byte {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *GetExecutionDataByBlockIDRequest) GetEventEncodingVersion() entities.EventEncodingVersion {
	if m != nil {
		return m.EventEncodingVersion
	}
	return entities.EventEncodingVersion_DEFAULT
}

// The response for GetExecutionDataByBlockID
type GetExecutionDataByBlockIDResponse struct {
	// BlockExecutionData for the block.
	BlockExecutionData   *entities.BlockExecutionData `protobuf:"bytes,1,opt,name=block_execution_data,json=blockExecutionData,proto3" json:"block_execution_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GetExecutionDataByBlockIDResponse) Reset()         { *m = GetExecutionDataByBlockIDResponse{} }
func (m *GetExecutionDataByBlockIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetExecutionDataByBlockIDResponse) ProtoMessage()    {}
func (*GetExecutionDataByBlockIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da469632570513fb, []int{1}
}

func (m *GetExecutionDataByBlockIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetExecutionDataByBlockIDResponse.Unmarshal(m, b)
}
func (m *GetExecutionDataByBlockIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetExecutionDataByBlockIDResponse.Marshal(b, m, deterministic)
}
func (m *GetExecutionDataByBlockIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetExecutionDataByBlockIDResponse.Merge(m, src)
}
func (m *GetExecutionDataByBlockIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetExecutionDataByBlockIDResponse.Size(m)
}
func (m *GetExecutionDataByBlockIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetExecutionDataByBlockIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetExecutionDataByBlockIDResponse proto.InternalMessageInfo

func (m *GetExecutionDataByBlockIDResponse) GetBlockExecutionData() *entities.BlockExecutionData {
	if m != nil {
		return m.BlockExecutionData
	}
	return nil
}

// The request for SubscribeExecutionData
type SubscribeExecutionDataRequest struct {
	// Block ID of the first block to get execution data for.
	// Only one of start_block_id and start_block_height may be provided,
	// otherwise an InvalidArgument error is returned. If neither are provided,
	// the latest sealed block is used.
	StartBlockId []byte `protobuf:"bytes,1,opt,name=start_block_id,json=startBlockId,proto3" json:"start_block_id,omitempty"`
	// Block height of the first block to get execution data for.
	// Only one of start_block_id and start_block_height may be provided,
	// otherwise an InvalidArgument error is returned. If neither are provided,
	// the latest sealed block is used.
	StartBlockHeight uint64 `protobuf:"varint,2,opt,name=start_block_height,json=startBlockHeight,proto3" json:"start_block_height,omitempty"`
	// Preferred event encoding version of the block events payload.
	// Possible variants:
	// 1. CCF
	// 2. JSON-CDC
	// 3. DEFAULT, which is JSON-CDC
	EventEncodingVersion entities.EventEncodingVersion `protobuf:"varint,3,opt,name=event_encoding_version,json=eventEncodingVersion,proto3,enum=flow.entities.EventEncodingVersion" json:"event_encoding_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SubscribeExecutionDataRequest) Reset()         { *m = SubscribeExecutionDataRequest{} }
func (m *SubscribeExecutionDataRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeExecutionDataRequest) ProtoMessage()    {}
func (*SubscribeExecutionDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da469632570513fb, []int{2}
}

func (m *SubscribeExecutionDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeExecutionDataRequest.Unmarshal(m, b)
}
func (m *SubscribeExecutionDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeExecutionDataRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeExecutionDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeExecutionDataRequest.Merge(m, src)
}
func (m *SubscribeExecutionDataRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeExecutionDataRequest.Size(m)
}
func (m *SubscribeExecutionDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeExecutionDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeExecutionDataRequest proto.InternalMessageInfo

func (m *SubscribeExecutionDataRequest) GetStartBlockId() []byte {
	if m != nil {
		return m.StartBlockId
	}
	return nil
}

func (m *SubscribeExecutionDataRequest) GetStartBlockHeight() uint64 {
	if m != nil {
		return m.StartBlockHeight
	}
	return 0
}

func (m *SubscribeExecutionDataRequest) GetEventEncodingVersion() entities.EventEncodingVersion {
	if m != nil {
		return m.EventEncodingVersion
	}
	return entities.EventEncodingVersion_DEFAULT
}

// The response for SubscribeExecutionData
type SubscribeExecutionDataResponse struct {
	// Block height of the block containing the execution data.
	BlockHeight uint64 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// BlockExecutionData for the block.
	// Note: The block's ID is included within the BlockExecutionData.
	BlockExecutionData *entities.BlockExecutionData `protobuf:"bytes,2,opt,name=block_execution_data,json=blockExecutionData,proto3" json:"block_execution_data,omitempty"`
	// Timestamp from the block containing the execution data.
	BlockTimestamp       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=block_timestamp,json=blockTimestamp,proto3" json:"block_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SubscribeExecutionDataResponse) Reset()         { *m = SubscribeExecutionDataResponse{} }
func (m *SubscribeExecutionDataResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeExecutionDataResponse) ProtoMessage()    {}
func (*SubscribeExecutionDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da469632570513fb, []int{3}
}

func (m *SubscribeExecutionDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeExecutionDataResponse.Unmarshal(m, b)
}
func (m *SubscribeExecutionDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeExecutionDataResponse.Marshal(b, m, deterministic)
}
func (m *SubscribeExecutionDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeExecutionDataResponse.Merge(m, src)
}
func (m *SubscribeExecutionDataResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeExecutionDataResponse.Size(m)
}
func (m *SubscribeExecutionDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeExecutionDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeExecutionDataResponse proto.InternalMessageInfo

func (m *SubscribeExecutionDataResponse) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *SubscribeExecutionDataResponse) GetBlockExecutionData() *entities.BlockExecutionData {
	if m != nil {
		return m.BlockExecutionData
	}
	return nil
}

func (m *SubscribeExecutionDataResponse) GetBlockTimestamp() *timestamppb.Timestamp {
	if m != nil {
		return m.BlockTimestamp
	}
	return nil
}

// The request for SubscribeEvents
type SubscribeEventsRequest struct {
	// Block ID of the first block to search for events.
	// Only one of start_block_id and start_block_height may be provided,
	// otherwise an InvalidArgument error is returned. If neither are provided,
	// the latest sealed block is used.
	StartBlockId []byte `protobuf:"bytes,1,opt,name=start_block_id,json=startBlockId,proto3" json:"start_block_id,omitempty"`
	// Block height of the first block to search for events.
	// Only one of start_block_id and start_block_height may be provided,
	// otherwise an InvalidArgument error is returned. If neither are provided,
	// the latest sealed block is used.
	StartBlockHeight uint64 `protobuf:"varint,2,opt,name=start_block_height,json=startBlockHeight,proto3" json:"start_block_height,omitempty"`
	// Filter to apply to events for each block searched.
	// If no filter is provided, all events are returned.
	Filter *EventFilter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	// Interval in block heights at which the server should return a heartbeat
	// message to the client. The heartbeat is a normal SubscribeEventsResponse
	// with no events, and allows clients to track which blocks were searched.
	// Clients can use this information to determine which block to start from
	// when reconnecting.
	//
	// The interval is calculated from the last response returned, which could be
	// either another heartbeat or a response containing events.
	HeartbeatInterval uint64 `protobuf:"varint,4,opt,name=heartbeat_interval,json=heartbeatInterval,proto3" json:"heartbeat_interval,omitempty"`
	// Preferred event encoding version of the block events payload.
	// Possible variants:
	// 1. CCF
	// 2. JSON-CDC
	// 3. DEFAULT, which is JSON-CDC
	EventEncodingVersion entities.EventEncodingVersion `protobuf:"varint,5,opt,name=event_encoding_version,json=eventEncodingVersion,proto3,enum=flow.entities.EventEncodingVersion" json:"event_encoding_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SubscribeEventsRequest) Reset()         { *m = SubscribeEventsRequest{} }
func (m *SubscribeEventsRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeEventsRequest) ProtoMessage()    {}
func (*SubscribeEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da469632570513fb, []int{4}
}

func (m *SubscribeEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeEventsRequest.Unmarshal(m, b)
}
func (m *SubscribeEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeEventsRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeEventsRequest.Merge(m, src)
}
func (m *SubscribeEventsRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeEventsRequest.Size(m)
}
func (m *SubscribeEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeEventsRequest proto.InternalMessageInfo

func (m *SubscribeEventsRequest) GetStartBlockId() []byte {
	if m != nil {
		return m.StartBlockId
	}
	return nil
}

func (m *SubscribeEventsRequest) GetStartBlockHeight() uint64 {
	if m != nil {
		return m.StartBlockHeight
	}
	return 0
}

func (m *SubscribeEventsRequest) GetFilter() *EventFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *SubscribeEventsRequest) GetHeartbeatInterval() uint64 {
	if m != nil {
		return m.HeartbeatInterval
	}
	return 0
}

func (m *SubscribeEventsRequest) GetEventEncodingVersion() entities.EventEncodingVersion {
	if m != nil {
		return m.EventEncodingVersion
	}
	return entities.EventEncodingVersion_DEFAULT
}

// The response for SubscribeEvents
type SubscribeEventsResponse struct {
	// Block ID of the block containing the events.
	BlockId []byte `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	// Block height of the block containing the events.
	BlockHeight uint64 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// Events matching the EventFilter in the request.
	// The API may return no events which signals a periodic heartbeat. This
	// allows clients to track which blocks were searched. Client can use this
	// information to determine which block to start from when reconnecting.
	Events []*entities.Event `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
	// Timestamp from the block containing the events.
	BlockTimestamp       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=block_timestamp,json=blockTimestamp,proto3" json:"block_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SubscribeEventsResponse) Reset()         { *m = SubscribeEventsResponse{} }
func (m *SubscribeEventsResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeEventsResponse) ProtoMessage()    {}
func (*SubscribeEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da469632570513fb, []int{5}
}

func (m *SubscribeEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeEventsResponse.Unmarshal(m, b)
}
func (m *SubscribeEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeEventsResponse.Marshal(b, m, deterministic)
}
func (m *SubscribeEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeEventsResponse.Merge(m, src)
}
func (m *SubscribeEventsResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeEventsResponse.Size(m)
}
func (m *SubscribeEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeEventsResponse proto.InternalMessageInfo

func (m *SubscribeEventsResponse) GetBlockId() []byte {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *SubscribeEventsResponse) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *SubscribeEventsResponse) GetEvents() []*entities.Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *SubscribeEventsResponse) GetBlockTimestamp() *timestamppb.Timestamp {
	if m != nil {
		return m.BlockTimestamp
	}
	return nil
}

// EventFilter defines the filter to apply to block events.
// Filters are applied as an OR operation, i.e. any event matching any of the
// filters is returned. If no filters are provided, all events are returned. If
// there are any invalid filters, the API will return an InvalidArgument error.
type EventFilter struct {
	// A list of full event types to include.
	//
	// All events exactly matching any of the provided event types will be
	// returned.
	//
	// Event types have 2 formats:
	//   - Protocol events:
	//     flow.[event name]
	//   - Smart contract events:
	//     A.[contract address].[contract name].[event name]
	EventType []string `protobuf:"bytes,1,rep,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	// A list of contracts who's events should be included.
	//
	// All events emitted by any of the provided contracts will be returned.
	//
	// Contracts have the following name formats:
	//   - Protocol events:
	//     flow
	//   - Smart contract events:
	//     A.[contract address].[contract name]
	//
	// This filter matches on the full contract including its address, not just
	// the contract's name.
	Contract []string `protobuf:"bytes,2,rep,name=contract,proto3" json:"contract,omitempty"`
	// A list of addresses who's events should be included.
	//
	// All events emitted by any contract held by any of the provided addresses
	// will be returned.
	//
	// Addresses must be Flow account addresses in hex format and valid for the
	// network the node is connected to. i.e. only a mainnet address is valid for
	// a mainnet node. Addresses may optionally include the 0x prefix.
	Address              []string `protobuf:"bytes,3,rep,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventFilter) Reset()         { *m = EventFilter{} }
func (m *EventFilter) String() string { return proto.CompactTextString(m) }
func (*EventFilter) ProtoMessage()    {}
func (*EventFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_da469632570513fb, []int{6}
}

func (m *EventFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventFilter.Unmarshal(m, b)
}
func (m *EventFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventFilter.Marshal(b, m, deterministic)
}
func (m *EventFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFilter.Merge(m, src)
}
func (m *EventFilter) XXX_Size() int {
	return xxx_messageInfo_EventFilter.Size(m)
}
func (m *EventFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFilter.DiscardUnknown(m)
}

var xxx_messageInfo_EventFilter proto.InternalMessageInfo

func (m *EventFilter) GetEventType() []string {
	if m != nil {
		return m.EventType
	}
	return nil
}

func (m *EventFilter) GetContract() []string {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *EventFilter) GetAddress() []string {
	if m != nil {
		return m.Address
	}
	return nil
}

func init() {
	proto.RegisterType((*GetExecutionDataByBlockIDRequest)(nil), "flow.access.GetExecutionDataByBlockIDRequest")
	proto.RegisterType((*GetExecutionDataByBlockIDResponse)(nil), "flow.access.GetExecutionDataByBlockIDResponse")
	proto.RegisterType((*SubscribeExecutionDataRequest)(nil), "flow.access.SubscribeExecutionDataRequest")
	proto.RegisterType((*SubscribeExecutionDataResponse)(nil), "flow.access.SubscribeExecutionDataResponse")
	proto.RegisterType((*SubscribeEventsRequest)(nil), "flow.access.SubscribeEventsRequest")
	proto.RegisterType((*SubscribeEventsResponse)(nil), "flow.access.SubscribeEventsResponse")
	proto.RegisterType((*EventFilter)(nil), "flow.access.EventFilter")
}

func init() {
	proto.RegisterFile("flow/executiondata/executiondata.proto", fileDescriptor_da469632570513fb)
}

var fileDescriptor_da469632570513fb = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x5d, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x93, 0xd2, 0x9f, 0x49, 0xd5, 0x96, 0x55, 0x55, 0x5c, 0x4b, 0x85, 0x34, 0xad, 0x50,
	0x04, 0xed, 0xa6, 0x0a, 0x27, 0x20, 0xb4, 0x40, 0x24, 0x1e, 0x90, 0x5b, 0x21, 0xc1, 0x03, 0xd1,
	0xda, 0x99, 0x3a, 0x2b, 0x52, 0x6f, 0xf0, 0x6e, 0x4a, 0xcb, 0x49, 0x38, 0x02, 0x67, 0xe1, 0x81,
	0x1b, 0xf0, 0xc4, 0x45, 0x90, 0x77, 0x6d, 0xd7, 0x4e, 0x13, 0x53, 0x14, 0xf1, 0x62, 0x69, 0x66,
	0xbe, 0xf9, 0xd9, 0x6f, 0xbe, 0x5d, 0xc3, 0xe3, 0xf3, 0xa1, 0xf8, 0xd2, 0xc2, 0x2b, 0xf4, 0xc7,
	0x8a, 0x8b, 0xb0, 0xcf, 0x14, 0x2b, 0x5a, 0x74, 0x14, 0x09, 0x25, 0x48, 0x2d, 0xc6, 0x51, 0xe6,
	0xfb, 0x28, 0xa5, 0xd3, 0x34, 0x49, 0xa1, 0xe2, 0x8a, 0xa3, 0x6c, 0x79, 0x43, 0xe1, 0x7f, 0xea,
	0x65, 0x59, 0xbd, 0x9b, 0x34, 0x67, 0xbb, 0x88, 0xc4, 0x4b, 0x0c, 0x55, 0x12, 0x7a, 0x14, 0x08,
	0x11, 0x0c, 0xb1, 0xa5, 0x2d, 0x6f, 0x7c, 0xde, 0x52, 0xfc, 0x02, 0xa5, 0x62, 0x17, 0x23, 0x03,
	0x68, 0x7c, 0xb3, 0xa0, 0xfe, 0x0a, 0xd5, 0x49, 0x5a, 0xf7, 0x98, 0x29, 0xd6, 0xb9, 0xee, 0xc4,
	0xcd, 0xba, 0xc7, 0x2e, 0x7e, 0x1e, 0xa3, 0x54, 0x64, 0x1b, 0x96, 0x4d, 0x7b, 0xde, 0xb7, 0xad,
	0xba, 0xd5, 0x5c, 0x75, 0x97, 0xb4, 0xdd, 0xed, 0x93, 0xf7, 0xb0, 0xa5, 0xfb, 0xf5, 0x30, 0xf4,
	0x45, 0x9f, 0x87, 0x41, 0xef, 0x12, 0x23, 0xc9, 0x45, 0x68, 0x57, 0xea, 0x56, 0x73, 0xad, 0xbd,
	0x47, 0xf5, 0x99, 0xd2, 0xe1, 0xe8, 0x49, 0x0c, 0x3e, 0x49, 0xb0, 0xef, 0x0c, 0xd4, 0xdd, 0xc4,
	0x29, 0xde, 0xc6, 0x15, 0xec, 0x96, 0x4c, 0x26, 0x47, 0x22, 0x94, 0x48, 0x4e, 0x61, 0x73, 0x1a,
	0x33, 0x7a, 0xcc, 0x5a, 0x7b, 0x77, 0xa2, 0xbb, 0xce, 0x2e, 0x54, 0x74, 0x89, 0x77, 0xcb, 0xd7,
	0xf8, 0x69, 0xc1, 0xce, 0xe9, 0xd8, 0x93, 0x7e, 0xc4, 0x3d, 0x2c, 0xc2, 0x13, 0x46, 0xf6, 0x61,
	0x4d, 0x2a, 0x16, 0xa9, 0xde, 0x04, 0x2f, 0xab, 0xda, 0xdb, 0x49, 0xc8, 0x39, 0x00, 0x92, 0x47,
	0x0d, 0x90, 0x07, 0x03, 0xa5, 0x89, 0x59, 0x70, 0x37, 0x6e, 0x90, 0xaf, 0xb5, 0xbf, 0x84, 0xca,
	0xea, 0xbc, 0x54, 0xfe, 0xb2, 0xe0, 0xe1, 0xac, 0x03, 0x25, 0x44, 0xee, 0xc2, 0x6a, 0x61, 0x4a,
	0x4b, 0x4f, 0x59, 0xf3, 0x72, 0x03, 0xce, 0xe2, 0xba, 0x32, 0x07, 0xd7, 0xe4, 0x05, 0xac, 0x9b,
	0xa2, 0x99, 0x32, 0xf5, 0x71, 0x6b, 0x6d, 0x87, 0x1a, 0xed, 0xd2, 0x54, 0xbb, 0xf4, 0x2c, 0x45,
	0xb8, 0x6b, 0x3a, 0x25, 0xb3, 0x1b, 0xdf, 0x2b, 0xb0, 0x75, 0x73, 0xbe, 0x98, 0x01, 0xf9, 0x3f,
	0x37, 0x75, 0x04, 0x8b, 0xe7, 0x7c, 0xa8, 0x30, 0x4a, 0x46, 0xb5, 0x69, 0xee, 0xe2, 0x9a, 0xbd,
	0xbc, 0xd4, 0x71, 0x37, 0xc1, 0x91, 0x43, 0x20, 0x03, 0x64, 0x91, 0xf2, 0x90, 0xa9, 0x1e, 0x0f,
	0x15, 0x46, 0x97, 0x6c, 0x68, 0x2f, 0xe8, 0xfa, 0xf7, 0xb3, 0x48, 0x37, 0x09, 0x94, 0x48, 0xe1,
	0xde, 0xbc, 0x52, 0xf8, 0x61, 0xc1, 0x83, 0x5b, 0x54, 0x25, 0x1a, 0x28, 0xb9, 0xe7, 0x93, 0xf2,
	0xa8, 0xdc, 0x96, 0xc7, 0x01, 0x2c, 0xea, 0x8e, 0xd2, 0xae, 0xd6, 0xab, 0xcd, 0x5a, 0x7b, 0x73,
	0xda, 0x90, 0x6e, 0x82, 0x99, 0xb6, 0xf7, 0x85, 0x7f, 0xde, 0xbb, 0x07, 0xb5, 0x1c, 0xdb, 0x64,
	0x07, 0xc0, 0xd0, 0xa6, 0xae, 0x47, 0x68, 0x5b, 0xf5, 0x6a, 0x73, 0xc5, 0x5d, 0xd1, 0x9e, 0xb3,
	0xeb, 0x11, 0x12, 0x07, 0x96, 0x7d, 0x11, 0xaa, 0x88, 0xf9, 0xf1, 0xfc, 0x71, 0x30, 0xb3, 0x89,
	0x0d, 0x4b, 0xac, 0xdf, 0x8f, 0x50, 0x9a, 0xe9, 0x57, 0xdc, 0xd4, 0x6c, 0xff, 0xae, 0xc0, 0x46,
	0x41, 0xb2, 0xcf, 0xdf, 0x76, 0xc9, 0x57, 0xd8, 0x9e, 0xf9, 0x36, 0x91, 0xc3, 0x82, 0x1c, 0xfe,
	0xf6, 0xba, 0x3a, 0xf4, 0xae, 0xf0, 0x64, 0x4b, 0x32, 0xaf, 0xf5, 0xc2, 0x5d, 0x7a, 0x52, 0xa8,
	0x54, 0xfa, 0x82, 0x39, 0x4f, 0xef, 0x84, 0x35, 0x2d, 0x8f, 0x2c, 0xf2, 0x11, 0xd6, 0x27, 0x54,
	0x43, 0xf6, 0x66, 0x54, 0xc8, 0x5f, 0x3f, 0x67, 0xbf, 0x1c, 0x94, 0xd6, 0xef, 0xbc, 0x01, 0x47,
	0x44, 0x01, 0x15, 0xa1, 0x86, 0x67, 0xab, 0x37, 0x79, 0x1f, 0x68, 0xc0, 0xd5, 0x60, 0xec, 0x51,
	0x5f, 0x5c, 0xb4, 0x0c, 0xa4, 0xa5, 0x3f, 0xd9, 0x6f, 0x2d, 0x10, 0xc6, 0x61, 0xf0, 0xde, 0xa2,
	0x0e, 0x3c, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x32, 0xa9, 0x3e, 0x79, 0x07, 0x00, 0x00,
}
