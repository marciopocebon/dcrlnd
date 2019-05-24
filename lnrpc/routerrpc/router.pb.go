// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routerrpc/router.proto

package routerrpc // import "github.com/decred/dcrlnd/lnrpc/routerrpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import lnrpc "github.com/decred/dcrlnd/lnrpc"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Failure_FailureCode int32

const (
	// *
	// The numbers assigned in this enumeration match the failure codes as
	// defined in BOLT #4. Because protobuf 3 requires enums to start with 0,
	// a RESERVED value is added.
	Failure_RESERVED                         Failure_FailureCode = 0
	Failure_UNKNOWN_PAYMENT_HASH             Failure_FailureCode = 1
	Failure_INCORRECT_PAYMENT_AMOUNT         Failure_FailureCode = 2
	Failure_FINAL_INCORRECT_CLTV_EXPIRY      Failure_FailureCode = 3
	Failure_FINAL_INCORRECT_HTLC_AMOUNT      Failure_FailureCode = 4
	Failure_FINAL_EXPIRY_TOO_SOON            Failure_FailureCode = 5
	Failure_INVALID_REALM                    Failure_FailureCode = 6
	Failure_EXPIRY_TOO_SOON                  Failure_FailureCode = 7
	Failure_INVALID_ONION_VERSION            Failure_FailureCode = 8
	Failure_INVALID_ONION_HMAC               Failure_FailureCode = 9
	Failure_INVALID_ONION_KEY                Failure_FailureCode = 10
	Failure_AMOUNT_BELOW_MINIMUM             Failure_FailureCode = 11
	Failure_FEE_INSUFFICIENT                 Failure_FailureCode = 12
	Failure_INCORRECT_CLTV_EXPIRY            Failure_FailureCode = 13
	Failure_CHANNEL_DISABLED                 Failure_FailureCode = 14
	Failure_TEMPORARY_CHANNEL_FAILURE        Failure_FailureCode = 15
	Failure_REQUIRED_NODE_FEATURE_MISSING    Failure_FailureCode = 16
	Failure_REQUIRED_CHANNEL_FEATURE_MISSING Failure_FailureCode = 17
	Failure_UNKNOWN_NEXT_PEER                Failure_FailureCode = 18
	Failure_TEMPORARY_NODE_FAILURE           Failure_FailureCode = 19
	Failure_PERMANENT_NODE_FAILURE           Failure_FailureCode = 20
	Failure_PERMANENT_CHANNEL_FAILURE        Failure_FailureCode = 21
)

var Failure_FailureCode_name = map[int32]string{
	0:  "RESERVED",
	1:  "UNKNOWN_PAYMENT_HASH",
	2:  "INCORRECT_PAYMENT_AMOUNT",
	3:  "FINAL_INCORRECT_CLTV_EXPIRY",
	4:  "FINAL_INCORRECT_HTLC_AMOUNT",
	5:  "FINAL_EXPIRY_TOO_SOON",
	6:  "INVALID_REALM",
	7:  "EXPIRY_TOO_SOON",
	8:  "INVALID_ONION_VERSION",
	9:  "INVALID_ONION_HMAC",
	10: "INVALID_ONION_KEY",
	11: "AMOUNT_BELOW_MINIMUM",
	12: "FEE_INSUFFICIENT",
	13: "INCORRECT_CLTV_EXPIRY",
	14: "CHANNEL_DISABLED",
	15: "TEMPORARY_CHANNEL_FAILURE",
	16: "REQUIRED_NODE_FEATURE_MISSING",
	17: "REQUIRED_CHANNEL_FEATURE_MISSING",
	18: "UNKNOWN_NEXT_PEER",
	19: "TEMPORARY_NODE_FAILURE",
	20: "PERMANENT_NODE_FAILURE",
	21: "PERMANENT_CHANNEL_FAILURE",
}
var Failure_FailureCode_value = map[string]int32{
	"RESERVED":                         0,
	"UNKNOWN_PAYMENT_HASH":             1,
	"INCORRECT_PAYMENT_AMOUNT":         2,
	"FINAL_INCORRECT_CLTV_EXPIRY":      3,
	"FINAL_INCORRECT_HTLC_AMOUNT":      4,
	"FINAL_EXPIRY_TOO_SOON":            5,
	"INVALID_REALM":                    6,
	"EXPIRY_TOO_SOON":                  7,
	"INVALID_ONION_VERSION":            8,
	"INVALID_ONION_HMAC":               9,
	"INVALID_ONION_KEY":                10,
	"AMOUNT_BELOW_MINIMUM":             11,
	"FEE_INSUFFICIENT":                 12,
	"INCORRECT_CLTV_EXPIRY":            13,
	"CHANNEL_DISABLED":                 14,
	"TEMPORARY_CHANNEL_FAILURE":        15,
	"REQUIRED_NODE_FEATURE_MISSING":    16,
	"REQUIRED_CHANNEL_FEATURE_MISSING": 17,
	"UNKNOWN_NEXT_PEER":                18,
	"TEMPORARY_NODE_FAILURE":           19,
	"PERMANENT_NODE_FAILURE":           20,
	"PERMANENT_CHANNEL_FAILURE":        21,
}

func (x Failure_FailureCode) String() string {
	return proto.EnumName(Failure_FailureCode_name, int32(x))
}
func (Failure_FailureCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_router_8d57d6a4e8581cae, []int{6, 0}
}

type PaymentRequest struct {
	// *
	// A serialized BOLT-11 payment request that contains all information
	// required to dispatch the payment. If the pay req is invalid, or expired,
	// an error will be returned.
	PayReq string `protobuf:"bytes,1,opt,name=pay_req,json=payReq,proto3" json:"pay_req,omitempty"`
	// *
	// An absolute limit on the highest fee we should pay when looking for a route
	// to the destination. Routes with fees higher than this will be ignored, if
	// there are no routes with a fee below this amount, an error will be
	// returned.
	FeeLimitAtoms int64 `protobuf:"varint,2,opt,name=fee_limit_atoms,json=feeLimitAtoms,proto3" json:"fee_limit_atoms,omitempty"`
	// *
	// An absolute limit on the cumulative CLTV value along the route for this
	// payment. Routes with total CLTV values higher than this will be ignored,
	// if there are no routes with a CLTV value below this amount, an error will
	// be returned.
	CltvLimit int32 `protobuf:"varint,3,opt,name=cltv_limit,json=cltvLimit,proto3" json:"cltv_limit,omitempty"`
	// *
	// An upper limit on the amount of time we should spend when attempting to
	// fulfill the payment. This is expressed in seconds. If we cannot make a
	// successful payment within this time frame, an error will be returned.
	TimeoutSeconds int32 `protobuf:"varint,4,opt,name=timeout_seconds,json=timeoutSeconds,proto3" json:"timeout_seconds,omitempty"`
	// *
	// The channel id of the channel that must be taken to the first hop. If zero,
	// any channel may be used.
	OutgoingChannelId    int64    `protobuf:"varint,5,opt,name=outgoing_channel_id,json=outgoingChannelId,proto3" json:"outgoing_channel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentRequest) Reset()         { *m = PaymentRequest{} }
func (m *PaymentRequest) String() string { return proto.CompactTextString(m) }
func (*PaymentRequest) ProtoMessage()    {}
func (*PaymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_router_8d57d6a4e8581cae, []int{0}
}
func (m *PaymentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentRequest.Unmarshal(m, b)
}
func (m *PaymentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentRequest.Marshal(b, m, deterministic)
}
func (dst *PaymentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentRequest.Merge(dst, src)
}
func (m *PaymentRequest) XXX_Size() int {
	return xxx_messageInfo_PaymentRequest.Size(m)
}
func (m *PaymentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentRequest proto.InternalMessageInfo

func (m *PaymentRequest) GetPayReq() string {
	if m != nil {
		return m.PayReq
	}
	return ""
}

func (m *PaymentRequest) GetFeeLimitAtoms() int64 {
	if m != nil {
		return m.FeeLimitAtoms
	}
	return 0
}

func (m *PaymentRequest) GetCltvLimit() int32 {
	if m != nil {
		return m.CltvLimit
	}
	return 0
}

func (m *PaymentRequest) GetTimeoutSeconds() int32 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

func (m *PaymentRequest) GetOutgoingChannelId() int64 {
	if m != nil {
		return m.OutgoingChannelId
	}
	return 0
}

type PaymentResponse struct {
	// *
	// The payment hash that we paid to. Provided so callers are able to map
	// responses (which may be streaming) back to their original requests.
	PayHash []byte `protobuf:"bytes,1,opt,name=pay_hash,json=payHash,proto3" json:"pay_hash,omitempty"`
	// *
	// The pre-image of the payment successfully completed.
	PreImage []byte `protobuf:"bytes,2,opt,name=pre_image,json=preImage,proto3" json:"pre_image,omitempty"`
	// *
	// If not an empty string, then a string representation of the payment error.
	PaymentErr           string   `protobuf:"bytes,3,opt,name=payment_err,json=paymentErr,proto3" json:"payment_err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentResponse) Reset()         { *m = PaymentResponse{} }
func (m *PaymentResponse) String() string { return proto.CompactTextString(m) }
func (*PaymentResponse) ProtoMessage()    {}
func (*PaymentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_router_8d57d6a4e8581cae, []int{1}
}
func (m *PaymentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentResponse.Unmarshal(m, b)
}
func (m *PaymentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentResponse.Marshal(b, m, deterministic)
}
func (dst *PaymentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentResponse.Merge(dst, src)
}
func (m *PaymentResponse) XXX_Size() int {
	return xxx_messageInfo_PaymentResponse.Size(m)
}
func (m *PaymentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentResponse proto.InternalMessageInfo

func (m *PaymentResponse) GetPayHash() []byte {
	if m != nil {
		return m.PayHash
	}
	return nil
}

func (m *PaymentResponse) GetPreImage() []byte {
	if m != nil {
		return m.PreImage
	}
	return nil
}

func (m *PaymentResponse) GetPaymentErr() string {
	if m != nil {
		return m.PaymentErr
	}
	return ""
}

type RouteFeeRequest struct {
	// *
	// The destination once wishes to obtain a routing fee quote to.
	Dest []byte `protobuf:"bytes,1,opt,name=dest,proto3" json:"dest,omitempty"`
	// *
	// The amount one wishes to send to the target destination.
	AmtAtoms             int64    `protobuf:"varint,2,opt,name=amt_atoms,json=amtAtoms,proto3" json:"amt_atoms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteFeeRequest) Reset()         { *m = RouteFeeRequest{} }
func (m *RouteFeeRequest) String() string { return proto.CompactTextString(m) }
func (*RouteFeeRequest) ProtoMessage()    {}
func (*RouteFeeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_router_8d57d6a4e8581cae, []int{2}
}
func (m *RouteFeeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteFeeRequest.Unmarshal(m, b)
}
func (m *RouteFeeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteFeeRequest.Marshal(b, m, deterministic)
}
func (dst *RouteFeeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteFeeRequest.Merge(dst, src)
}
func (m *RouteFeeRequest) XXX_Size() int {
	return xxx_messageInfo_RouteFeeRequest.Size(m)
}
func (m *RouteFeeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteFeeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RouteFeeRequest proto.InternalMessageInfo

func (m *RouteFeeRequest) GetDest() []byte {
	if m != nil {
		return m.Dest
	}
	return nil
}

func (m *RouteFeeRequest) GetAmtAtoms() int64 {
	if m != nil {
		return m.AmtAtoms
	}
	return 0
}

type RouteFeeResponse struct {
	// *
	// A lower bound of the estimated fee to the target destination within the
	// network, expressed in milli-satoshis.
	RoutingFeeMatoms int64 `protobuf:"varint,1,opt,name=routing_fee_matoms,json=routingFeeMatoms,proto3" json:"routing_fee_matoms,omitempty"`
	// *
	// An estimate of the worst case time delay that can occur. Note that callers
	// will still need to factor in the final CLTV delta of the last hop into this
	// value.
	TimeLockDelay        int64    `protobuf:"varint,2,opt,name=time_lock_delay,json=timeLockDelay,proto3" json:"time_lock_delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteFeeResponse) Reset()         { *m = RouteFeeResponse{} }
func (m *RouteFeeResponse) String() string { return proto.CompactTextString(m) }
func (*RouteFeeResponse) ProtoMessage()    {}
func (*RouteFeeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_router_8d57d6a4e8581cae, []int{3}
}
func (m *RouteFeeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteFeeResponse.Unmarshal(m, b)
}
func (m *RouteFeeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteFeeResponse.Marshal(b, m, deterministic)
}
func (dst *RouteFeeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteFeeResponse.Merge(dst, src)
}
func (m *RouteFeeResponse) XXX_Size() int {
	return xxx_messageInfo_RouteFeeResponse.Size(m)
}
func (m *RouteFeeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteFeeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RouteFeeResponse proto.InternalMessageInfo

func (m *RouteFeeResponse) GetRoutingFeeMatoms() int64 {
	if m != nil {
		return m.RoutingFeeMatoms
	}
	return 0
}

func (m *RouteFeeResponse) GetTimeLockDelay() int64 {
	if m != nil {
		return m.TimeLockDelay
	}
	return 0
}

type SendToRouteRequest struct {
	// / The payment hash to use for the HTLC.
	PaymentHash []byte `protobuf:"bytes,1,opt,name=payment_hash,json=paymentHash,proto3" json:"payment_hash,omitempty"`
	// / Route that should be used to attempt to complete the payment.
	Route                *lnrpc.Route `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SendToRouteRequest) Reset()         { *m = SendToRouteRequest{} }
func (m *SendToRouteRequest) String() string { return proto.CompactTextString(m) }
func (*SendToRouteRequest) ProtoMessage()    {}
func (*SendToRouteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_router_8d57d6a4e8581cae, []int{4}
}
func (m *SendToRouteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendToRouteRequest.Unmarshal(m, b)
}
func (m *SendToRouteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendToRouteRequest.Marshal(b, m, deterministic)
}
func (dst *SendToRouteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendToRouteRequest.Merge(dst, src)
}
func (m *SendToRouteRequest) XXX_Size() int {
	return xxx_messageInfo_SendToRouteRequest.Size(m)
}
func (m *SendToRouteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendToRouteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendToRouteRequest proto.InternalMessageInfo

func (m *SendToRouteRequest) GetPaymentHash() []byte {
	if m != nil {
		return m.PaymentHash
	}
	return nil
}

func (m *SendToRouteRequest) GetRoute() *lnrpc.Route {
	if m != nil {
		return m.Route
	}
	return nil
}

type SendToRouteResponse struct {
	// / The preimage obtained by making the payment.
	Preimage []byte `protobuf:"bytes,1,opt,name=preimage,proto3" json:"preimage,omitempty"`
	// / The failure message in case the payment failed.
	Failure              *Failure `protobuf:"bytes,2,opt,name=failure,proto3" json:"failure,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendToRouteResponse) Reset()         { *m = SendToRouteResponse{} }
func (m *SendToRouteResponse) String() string { return proto.CompactTextString(m) }
func (*SendToRouteResponse) ProtoMessage()    {}
func (*SendToRouteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_router_8d57d6a4e8581cae, []int{5}
}
func (m *SendToRouteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendToRouteResponse.Unmarshal(m, b)
}
func (m *SendToRouteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendToRouteResponse.Marshal(b, m, deterministic)
}
func (dst *SendToRouteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendToRouteResponse.Merge(dst, src)
}
func (m *SendToRouteResponse) XXX_Size() int {
	return xxx_messageInfo_SendToRouteResponse.Size(m)
}
func (m *SendToRouteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendToRouteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendToRouteResponse proto.InternalMessageInfo

func (m *SendToRouteResponse) GetPreimage() []byte {
	if m != nil {
		return m.Preimage
	}
	return nil
}

func (m *SendToRouteResponse) GetFailure() *Failure {
	if m != nil {
		return m.Failure
	}
	return nil
}

type Failure struct {
	// / Failure code as defined in the Lightning spec
	Code Failure_FailureCode `protobuf:"varint,1,opt,name=code,proto3,enum=routerrpc.Failure_FailureCode" json:"code,omitempty"`
	// *
	// The node pubkey of the intermediate or final node that generated the failure
	// message.
	FailureSourcePubkey []byte `protobuf:"bytes,2,opt,name=failure_source_pubkey,json=failureSourcePubkey,proto3" json:"failure_source_pubkey,omitempty"`
	// / An optional channel update message.
	ChannelUpdate *ChannelUpdate `protobuf:"bytes,3,opt,name=channel_update,json=channelUpdate,proto3" json:"channel_update,omitempty"`
	// / A failure type-dependent htlc value.
	HtlcMAtoms uint64 `protobuf:"varint,4,opt,name=htlc_m_atoms,json=htlcMAtoms,proto3" json:"htlc_m_atoms,omitempty"`
	// / The sha256 sum of the onion payload.
	OnionSha_256 []byte `protobuf:"bytes,5,opt,name=onion_sha_256,json=onionSha256,proto3" json:"onion_sha_256,omitempty"`
	// / A failure type-dependent cltv expiry value.
	CltvExpiry uint32 `protobuf:"varint,6,opt,name=cltv_expiry,json=cltvExpiry,proto3" json:"cltv_expiry,omitempty"`
	// / A failure type-dependent flags value.
	Flags                uint32   `protobuf:"varint,7,opt,name=flags,proto3" json:"flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Failure) Reset()         { *m = Failure{} }
func (m *Failure) String() string { return proto.CompactTextString(m) }
func (*Failure) ProtoMessage()    {}
func (*Failure) Descriptor() ([]byte, []int) {
	return fileDescriptor_router_8d57d6a4e8581cae, []int{6}
}
func (m *Failure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Failure.Unmarshal(m, b)
}
func (m *Failure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Failure.Marshal(b, m, deterministic)
}
func (dst *Failure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Failure.Merge(dst, src)
}
func (m *Failure) XXX_Size() int {
	return xxx_messageInfo_Failure.Size(m)
}
func (m *Failure) XXX_DiscardUnknown() {
	xxx_messageInfo_Failure.DiscardUnknown(m)
}

var xxx_messageInfo_Failure proto.InternalMessageInfo

func (m *Failure) GetCode() Failure_FailureCode {
	if m != nil {
		return m.Code
	}
	return Failure_RESERVED
}

func (m *Failure) GetFailureSourcePubkey() []byte {
	if m != nil {
		return m.FailureSourcePubkey
	}
	return nil
}

func (m *Failure) GetChannelUpdate() *ChannelUpdate {
	if m != nil {
		return m.ChannelUpdate
	}
	return nil
}

func (m *Failure) GetHtlcMAtoms() uint64 {
	if m != nil {
		return m.HtlcMAtoms
	}
	return 0
}

func (m *Failure) GetOnionSha_256() []byte {
	if m != nil {
		return m.OnionSha_256
	}
	return nil
}

func (m *Failure) GetCltvExpiry() uint32 {
	if m != nil {
		return m.CltvExpiry
	}
	return 0
}

func (m *Failure) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

type ChannelUpdate struct {
	// *
	// The signature that validates the announced data and proves the ownership
	// of node id.
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// *
	// The target chain that this channel was opened within. This value
	// should be the genesis hash of the target chain. Along with the short
	// channel ID, this uniquely identifies the channel globally in a
	// blockchain.
	ChainHash []byte `protobuf:"bytes,2,opt,name=chain_hash,json=chainHash,proto3" json:"chain_hash,omitempty"`
	// *
	// The unique description of the funding transaction.
	ChanId uint64 `protobuf:"varint,3,opt,name=chan_id,json=chanId,proto3" json:"chan_id,omitempty"`
	// *
	// A timestamp that allows ordering in the case of multiple announcements.
	// We should ignore the message if timestamp is not greater than the
	// last-received.
	Timestamp uint32 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// *
	// The bitfield that describes whether optional fields are present in this
	// update. Currently, the least-significant bit must be set to 1 if the
	// optional field MaxHtlc is present.
	MessageFlags uint32 `protobuf:"varint,10,opt,name=message_flags,json=messageFlags,proto3" json:"message_flags,omitempty"`
	// *
	// The bitfield that describes additional meta-data concerning how the
	// update is to be interpreted. Currently, the least-significant bit must be
	// set to 0 if the creating node corresponds to the first node in the
	// previously sent channel announcement and 1 otherwise. If the second bit
	// is set, then the channel is set to be disabled.
	ChannelFlags uint32 `protobuf:"varint,5,opt,name=channel_flags,json=channelFlags,proto3" json:"channel_flags,omitempty"`
	// *
	// The minimum number of blocks this node requires to be added to the expiry
	// of HTLCs. This is a security parameter determined by the node operator.
	// This value represents the required gap between the time locks of the
	// incoming and outgoing HTLC's set to this node.
	TimeLockDelta uint32 `protobuf:"varint,6,opt,name=time_lock_delta,json=timeLockDelta,proto3" json:"time_lock_delta,omitempty"`
	// *
	// The minimum HTLC value which will be accepted.
	HtlcMinimumMAtoms uint64 `protobuf:"varint,7,opt,name=htlc_minimum_m_atoms,json=htlcMinimumMAtoms,proto3" json:"htlc_minimum_m_atoms,omitempty"`
	// *
	// The base fee that must be used for incoming HTLC's to this particular
	// channel. This value will be tacked onto the required for a payment
	// independent of the size of the payment.
	BaseFee uint32 `protobuf:"varint,8,opt,name=base_fee,json=baseFee,proto3" json:"base_fee,omitempty"`
	// *
	// The fee rate that will be charged per millionth of a satoshi.
	FeeRate uint32 `protobuf:"varint,9,opt,name=fee_rate,json=feeRate,proto3" json:"fee_rate,omitempty"`
	// *
	// The maximum HTLC value which will be accepted.
	HtlcMaximumMAtoms uint64 `protobuf:"varint,11,opt,name=htlc_maximum_m_atoms,json=htlcMaximumMAtoms,proto3" json:"htlc_maximum_m_atoms,omitempty"`
	// *
	// The set of data that was appended to this message, some of which we may
	// not actually know how to iterate or parse. By holding onto this data, we
	// ensure that we're able to properly validate the set of signatures that
	// cover these new fields, and ensure we're able to make upgrades to the
	// network in a forwards compatible manner.
	ExtraOpaqueData      []byte   `protobuf:"bytes,12,opt,name=extra_opaque_data,json=extraOpaqueData,proto3" json:"extra_opaque_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelUpdate) Reset()         { *m = ChannelUpdate{} }
func (m *ChannelUpdate) String() string { return proto.CompactTextString(m) }
func (*ChannelUpdate) ProtoMessage()    {}
func (*ChannelUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_router_8d57d6a4e8581cae, []int{7}
}
func (m *ChannelUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelUpdate.Unmarshal(m, b)
}
func (m *ChannelUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelUpdate.Marshal(b, m, deterministic)
}
func (dst *ChannelUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelUpdate.Merge(dst, src)
}
func (m *ChannelUpdate) XXX_Size() int {
	return xxx_messageInfo_ChannelUpdate.Size(m)
}
func (m *ChannelUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelUpdate proto.InternalMessageInfo

func (m *ChannelUpdate) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *ChannelUpdate) GetChainHash() []byte {
	if m != nil {
		return m.ChainHash
	}
	return nil
}

func (m *ChannelUpdate) GetChanId() uint64 {
	if m != nil {
		return m.ChanId
	}
	return 0
}

func (m *ChannelUpdate) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ChannelUpdate) GetMessageFlags() uint32 {
	if m != nil {
		return m.MessageFlags
	}
	return 0
}

func (m *ChannelUpdate) GetChannelFlags() uint32 {
	if m != nil {
		return m.ChannelFlags
	}
	return 0
}

func (m *ChannelUpdate) GetTimeLockDelta() uint32 {
	if m != nil {
		return m.TimeLockDelta
	}
	return 0
}

func (m *ChannelUpdate) GetHtlcMinimumMAtoms() uint64 {
	if m != nil {
		return m.HtlcMinimumMAtoms
	}
	return 0
}

func (m *ChannelUpdate) GetBaseFee() uint32 {
	if m != nil {
		return m.BaseFee
	}
	return 0
}

func (m *ChannelUpdate) GetFeeRate() uint32 {
	if m != nil {
		return m.FeeRate
	}
	return 0
}

func (m *ChannelUpdate) GetHtlcMaximumMAtoms() uint64 {
	if m != nil {
		return m.HtlcMaximumMAtoms
	}
	return 0
}

func (m *ChannelUpdate) GetExtraOpaqueData() []byte {
	if m != nil {
		return m.ExtraOpaqueData
	}
	return nil
}

func init() {
	proto.RegisterType((*PaymentRequest)(nil), "routerrpc.PaymentRequest")
	proto.RegisterType((*PaymentResponse)(nil), "routerrpc.PaymentResponse")
	proto.RegisterType((*RouteFeeRequest)(nil), "routerrpc.RouteFeeRequest")
	proto.RegisterType((*RouteFeeResponse)(nil), "routerrpc.RouteFeeResponse")
	proto.RegisterType((*SendToRouteRequest)(nil), "routerrpc.SendToRouteRequest")
	proto.RegisterType((*SendToRouteResponse)(nil), "routerrpc.SendToRouteResponse")
	proto.RegisterType((*Failure)(nil), "routerrpc.Failure")
	proto.RegisterType((*ChannelUpdate)(nil), "routerrpc.ChannelUpdate")
	proto.RegisterEnum("routerrpc.Failure_FailureCode", Failure_FailureCode_name, Failure_FailureCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouterClient interface {
	// *
	// SendPayment attempts to route a payment described by the passed
	// PaymentRequest to the final destination. If we are unable to route the
	// payment, or cannot find a route that satisfies the constraints in the
	// PaymentRequest, then an error will be returned. Otherwise, the payment
	// pre-image, along with the final route will be returned.
	SendPayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error)
	// *
	// EstimateRouteFee allows callers to obtain a lower bound w.r.t how much it
	// may cost to send an HTLC to the target end destination.
	EstimateRouteFee(ctx context.Context, in *RouteFeeRequest, opts ...grpc.CallOption) (*RouteFeeResponse, error)
	// *
	// SendToRoute attempts to make a payment via the specified route. This method
	// differs from SendPayment in that it allows users to specify a full route
	// manually. This can be used for things like rebalancing, and atomic swaps.
	SendToRoute(ctx context.Context, in *SendToRouteRequest, opts ...grpc.CallOption) (*SendToRouteResponse, error)
}

type routerClient struct {
	cc *grpc.ClientConn
}

func NewRouterClient(cc *grpc.ClientConn) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) SendPayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error) {
	out := new(PaymentResponse)
	err := c.cc.Invoke(ctx, "/routerrpc.Router/SendPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) EstimateRouteFee(ctx context.Context, in *RouteFeeRequest, opts ...grpc.CallOption) (*RouteFeeResponse, error) {
	out := new(RouteFeeResponse)
	err := c.cc.Invoke(ctx, "/routerrpc.Router/EstimateRouteFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) SendToRoute(ctx context.Context, in *SendToRouteRequest, opts ...grpc.CallOption) (*SendToRouteResponse, error) {
	out := new(SendToRouteResponse)
	err := c.cc.Invoke(ctx, "/routerrpc.Router/SendToRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterServer is the server API for Router service.
type RouterServer interface {
	// *
	// SendPayment attempts to route a payment described by the passed
	// PaymentRequest to the final destination. If we are unable to route the
	// payment, or cannot find a route that satisfies the constraints in the
	// PaymentRequest, then an error will be returned. Otherwise, the payment
	// pre-image, along with the final route will be returned.
	SendPayment(context.Context, *PaymentRequest) (*PaymentResponse, error)
	// *
	// EstimateRouteFee allows callers to obtain a lower bound w.r.t how much it
	// may cost to send an HTLC to the target end destination.
	EstimateRouteFee(context.Context, *RouteFeeRequest) (*RouteFeeResponse, error)
	// *
	// SendToRoute attempts to make a payment via the specified route. This method
	// differs from SendPayment in that it allows users to specify a full route
	// manually. This can be used for things like rebalancing, and atomic swaps.
	SendToRoute(context.Context, *SendToRouteRequest) (*SendToRouteResponse, error)
}

func RegisterRouterServer(s *grpc.Server, srv RouterServer) {
	s.RegisterService(&_Router_serviceDesc, srv)
}

func _Router_SendPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).SendPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routerrpc.Router/SendPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).SendPayment(ctx, req.(*PaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_EstimateRouteFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).EstimateRouteFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routerrpc.Router/EstimateRouteFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).EstimateRouteFee(ctx, req.(*RouteFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_SendToRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendToRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).SendToRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routerrpc.Router/SendToRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).SendToRoute(ctx, req.(*SendToRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Router_serviceDesc = grpc.ServiceDesc{
	ServiceName: "routerrpc.Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPayment",
			Handler:    _Router_SendPayment_Handler,
		},
		{
			MethodName: "EstimateRouteFee",
			Handler:    _Router_EstimateRouteFee_Handler,
		},
		{
			MethodName: "SendToRoute",
			Handler:    _Router_SendToRoute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "routerrpc/router.proto",
}

func init() { proto.RegisterFile("routerrpc/router.proto", fileDescriptor_router_8d57d6a4e8581cae) }

var fileDescriptor_router_8d57d6a4e8581cae = []byte{
	// 1221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x56, 0xef, 0x72, 0xda, 0x46,
	0x10, 0x2f, 0xf1, 0x1f, 0x60, 0x01, 0x23, 0x9f, 0xed, 0x04, 0x93, 0x38, 0x71, 0x68, 0xa7, 0xf5,
	0x64, 0x32, 0xf6, 0x0c, 0x9d, 0xe4, 0x6b, 0x47, 0x86, 0xa3, 0xd6, 0x04, 0x49, 0xf4, 0x04, 0x4e,
	0xdc, 0x7e, 0xb8, 0x39, 0x4b, 0x67, 0x50, 0x83, 0xfe, 0x58, 0x12, 0x9d, 0xf8, 0x19, 0xfa, 0x44,
	0x7d, 0x84, 0x3e, 0x42, 0x1f, 0xa1, 0x6f, 0xd1, 0xb9, 0x3b, 0x09, 0x83, 0xe3, 0x7e, 0x82, 0xfb,
	0xfd, 0x7e, 0xb7, 0x7b, 0xbb, 0xb7, 0xbb, 0x27, 0x78, 0x9a, 0x44, 0x8b, 0x8c, 0x27, 0x49, 0xec,
	0x9e, 0xa9, 0x7f, 0xa7, 0x71, 0x12, 0x65, 0x11, 0xaa, 0x2e, 0xf1, 0x76, 0x35, 0x89, 0x5d, 0x85,
	0x76, 0xfe, 0x2e, 0xc1, 0xce, 0x88, 0xdd, 0x05, 0x3c, 0xcc, 0x08, 0xbf, 0x5d, 0xf0, 0x34, 0x43,
	0xcf, 0xa0, 0x1c, 0xb3, 0x3b, 0x9a, 0xf0, 0xdb, 0x56, 0xe9, 0xb8, 0x74, 0x52, 0x25, 0xdb, 0x31,
	0xbb, 0x23, 0xfc, 0x16, 0x7d, 0x0f, 0xcd, 0x1b, 0xce, 0xe9, 0xdc, 0x0f, 0xfc, 0x8c, 0xb2, 0x2c,
	0x0a, 0xd2, 0xd6, 0x93, 0xe3, 0xd2, 0xc9, 0x06, 0x69, 0xdc, 0x70, 0x3e, 0x14, 0xa8, 0x2e, 0x40,
	0x74, 0x04, 0xe0, 0xce, 0xb3, 0x3f, 0x94, 0xb0, 0xb5, 0x71, 0x5c, 0x3a, 0xd9, 0x22, 0x55, 0x81,
	0x48, 0x0d, 0xfa, 0x01, 0x9a, 0x99, 0x1f, 0xf0, 0x68, 0x91, 0xd1, 0x94, 0xbb, 0x51, 0xe8, 0xa5,
	0xad, 0x4d, 0xa9, 0xd9, 0xc9, 0x61, 0x47, 0xa1, 0xe8, 0x14, 0xf6, 0xa2, 0x45, 0x36, 0x8d, 0xfc,
	0x70, 0x4a, 0xdd, 0x19, 0x0b, 0x43, 0x3e, 0xa7, 0xbe, 0xd7, 0xda, 0x92, 0x3e, 0x77, 0x0b, 0xaa,
	0xa7, 0x18, 0xc3, 0xeb, 0xfc, 0x0e, 0xcd, 0x65, 0x28, 0x69, 0x1c, 0x85, 0x29, 0x47, 0x87, 0x50,
	0x11, 0xb1, 0xcc, 0x58, 0x3a, 0x93, 0xc1, 0xd4, 0x89, 0x88, 0xed, 0x82, 0xa5, 0x33, 0xf4, 0x1c,
	0xaa, 0x71, 0xc2, 0xa9, 0x1f, 0xb0, 0x29, 0x97, 0x71, 0xd4, 0x49, 0x25, 0x4e, 0xb8, 0x21, 0xd6,
	0xe8, 0x15, 0xd4, 0x62, 0x65, 0x8a, 0xf2, 0x24, 0x91, 0x31, 0x54, 0x09, 0xe4, 0x10, 0x4e, 0x92,
	0xce, 0x39, 0x34, 0x89, 0xc8, 0xe7, 0x80, 0xf3, 0x22, 0x6f, 0x08, 0x36, 0x3d, 0x9e, 0x66, 0xb9,
	0x1f, 0xf9, 0x5f, 0x38, 0x61, 0xc1, 0x7a, 0xb2, 0x2a, 0x2c, 0x50, 0x79, 0xea, 0xcc, 0x40, 0xbb,
	0xb7, 0x91, 0x1f, 0xf8, 0x2d, 0x20, 0x71, 0x4f, 0x22, 0x64, 0x91, 0xeb, 0x40, 0xed, 0x2c, 0xc9,
	0x9d, 0x5a, 0xce, 0x0c, 0x38, 0x37, 0x25, 0x2e, 0x6e, 0x44, 0xe4, 0x8c, 0xce, 0x23, 0xf7, 0x33,
	0xf5, 0xf8, 0x9c, 0xdd, 0x15, 0x37, 0x22, 0xe0, 0x61, 0xe4, 0x7e, 0xee, 0x0b, 0xb0, 0xf3, 0x1b,
	0x20, 0x87, 0x87, 0xde, 0x38, 0x92, 0xfe, 0x8a, 0x03, 0xbf, 0x86, 0x7a, 0x11, 0xe4, 0x4a, 0x82,
	0x8a, 0xc0, 0x65, 0x92, 0x3a, 0xb0, 0x25, 0xcb, 0x46, 0x9a, 0xad, 0x75, 0xeb, 0xa7, 0xf3, 0x50,
	0xd4, 0x8e, 0x32, 0xa3, 0xa8, 0x0e, 0x85, 0xbd, 0x35, 0xe3, 0x79, 0x24, 0x6d, 0x10, 0xe9, 0x54,
	0xe9, 0x2d, 0x2d, 0xd3, 0x2b, 0xd7, 0xe8, 0x2d, 0x94, 0x6f, 0x98, 0x3f, 0x5f, 0x24, 0x85, 0x61,
	0x74, 0xba, 0xac, 0xce, 0xd3, 0x81, 0x62, 0x48, 0x21, 0xe9, 0xfc, 0x59, 0x86, 0x72, 0x0e, 0xa2,
	0x2e, 0x6c, 0xba, 0x91, 0xa7, 0x2c, 0xee, 0x74, 0x5f, 0x7e, 0xbd, 0xad, 0xf8, 0xed, 0x45, 0x1e,
	0x27, 0x52, 0x8b, 0xba, 0x70, 0x90, 0x9b, 0xa2, 0x69, 0xb4, 0x48, 0x5c, 0x4e, 0xe3, 0xc5, 0xf5,
	0x67, 0x7e, 0x97, 0xdf, 0xfa, 0x5e, 0x4e, 0x3a, 0x92, 0x1b, 0x49, 0x0a, 0xfd, 0x04, 0x3b, 0x45,
	0xc9, 0x2d, 0x62, 0x8f, 0x65, 0x5c, 0xd6, 0x40, 0xad, 0xdb, 0x5a, 0xf1, 0x98, 0x57, 0xde, 0x44,
	0xf2, 0xa4, 0xe1, 0xae, 0x2e, 0xd1, 0x31, 0xd4, 0x67, 0xd9, 0xdc, 0xa5, 0x41, 0x7e, 0xf9, 0xa2,
	0xc4, 0x37, 0x09, 0x08, 0xcc, 0x54, 0x6d, 0xd2, 0x81, 0x46, 0x14, 0xfa, 0x51, 0x48, 0xd3, 0x19,
	0xa3, 0xdd, 0x77, 0xef, 0x65, 0x61, 0xd7, 0x49, 0x4d, 0x82, 0xce, 0x8c, 0x75, 0xdf, 0xbd, 0x17,
	0x75, 0x28, 0x5b, 0x89, 0x7f, 0x89, 0xfd, 0xe4, 0xae, 0xb5, 0x7d, 0x5c, 0x3a, 0x69, 0x10, 0xd9,
	0x5d, 0x58, 0x22, 0x68, 0x1f, 0xb6, 0x6e, 0xe6, 0x6c, 0x9a, 0xb6, 0xca, 0x92, 0x52, 0x8b, 0xce,
	0x3f, 0x9b, 0x50, 0x5b, 0xc9, 0x03, 0xaa, 0x43, 0x85, 0x60, 0x07, 0x93, 0x4b, 0xdc, 0xd7, 0xbe,
	0x41, 0x2d, 0xd8, 0x9f, 0x58, 0x1f, 0x2c, 0xfb, 0xa3, 0x45, 0x47, 0xfa, 0x95, 0x89, 0xad, 0x31,
	0xbd, 0xd0, 0x9d, 0x0b, 0xad, 0x84, 0x5e, 0x40, 0xcb, 0xb0, 0x7a, 0x36, 0x21, 0xb8, 0x37, 0x5e,
	0x72, 0xba, 0x69, 0x4f, 0xac, 0xb1, 0xf6, 0x04, 0xbd, 0x82, 0xe7, 0x03, 0xc3, 0xd2, 0x87, 0xf4,
	0x5e, 0xd3, 0x1b, 0x8e, 0x2f, 0x29, 0xfe, 0x34, 0x32, 0xc8, 0x95, 0xb6, 0xf1, 0x98, 0xe0, 0x62,
	0x3c, 0xec, 0x15, 0x16, 0x36, 0xd1, 0x21, 0x1c, 0x28, 0x81, 0xda, 0x42, 0xc7, 0xb6, 0x4d, 0x1d,
	0xdb, 0xb6, 0xb4, 0x2d, 0xb4, 0x0b, 0x0d, 0xc3, 0xba, 0xd4, 0x87, 0x46, 0x9f, 0x12, 0xac, 0x0f,
	0x4d, 0x6d, 0x1b, 0xed, 0x41, 0xf3, 0xa1, 0xae, 0x2c, 0x4c, 0x14, 0x3a, 0xdb, 0x32, 0x6c, 0x8b,
	0x5e, 0x62, 0xe2, 0x18, 0xb6, 0xa5, 0x55, 0xd0, 0x53, 0x40, 0xeb, 0xd4, 0x85, 0xa9, 0xf7, 0xb4,
	0x2a, 0x3a, 0x80, 0xdd, 0x75, 0xfc, 0x03, 0xbe, 0xd2, 0x40, 0xa4, 0x41, 0x1d, 0x8c, 0x9e, 0xe3,
	0xa1, 0xfd, 0x91, 0x9a, 0x86, 0x65, 0x98, 0x13, 0x53, 0xab, 0xa1, 0x7d, 0xd0, 0x06, 0x18, 0x53,
	0xc3, 0x72, 0x26, 0x83, 0x81, 0xd1, 0x33, 0xb0, 0x35, 0xd6, 0xea, 0xca, 0xf3, 0x63, 0x81, 0x37,
	0xc4, 0x86, 0xde, 0x85, 0x6e, 0x59, 0x78, 0x48, 0xfb, 0x86, 0xa3, 0x9f, 0x0f, 0x71, 0x5f, 0xdb,
	0x41, 0x47, 0x70, 0x38, 0xc6, 0xe6, 0xc8, 0x26, 0x3a, 0xb9, 0xa2, 0x05, 0x3f, 0xd0, 0x8d, 0xe1,
	0x84, 0x60, 0xad, 0x89, 0x5e, 0xc3, 0x11, 0xc1, 0xbf, 0x4c, 0x0c, 0x82, 0xfb, 0xd4, 0xb2, 0xfb,
	0x98, 0x0e, 0xb0, 0x3e, 0x9e, 0x10, 0x4c, 0x4d, 0xc3, 0x71, 0x0c, 0xeb, 0x67, 0x4d, 0x43, 0xdf,
	0xc1, 0xf1, 0x52, 0xb2, 0x34, 0xf0, 0x40, 0xb5, 0x2b, 0xe2, 0x2b, 0xee, 0xd3, 0xc2, 0x9f, 0xc6,
	0x74, 0x84, 0x31, 0xd1, 0x10, 0x6a, 0xc3, 0xd3, 0x7b, 0xf7, 0xca, 0x41, 0xee, 0x7b, 0x4f, 0x70,
	0x23, 0x4c, 0x4c, 0xdd, 0x12, 0x17, 0xbc, 0xc6, 0xed, 0x8b, 0x63, 0xdf, 0x73, 0x0f, 0x8f, 0x7d,
	0xd0, 0xf9, 0x6b, 0x03, 0x1a, 0x6b, 0x95, 0x8f, 0x5e, 0x40, 0x35, 0xf5, 0xa7, 0x21, 0xcb, 0x44,
	0x3f, 0xab, 0x56, 0xbf, 0x07, 0xe4, 0x6b, 0x30, 0x63, 0x7e, 0xa8, 0x66, 0x8c, 0x6a, 0xb9, 0xaa,
	0x44, 0xe4, 0x84, 0x79, 0x06, 0x65, 0xd1, 0x38, 0x62, 0xb0, 0x6f, 0xc8, 0x16, 0xd9, 0x16, 0x4b,
	0xc3, 0x13, 0x56, 0xc5, 0x10, 0x4b, 0x33, 0x16, 0xc4, 0xb2, 0x7b, 0x1a, 0xe4, 0x1e, 0x40, 0xdf,
	0x42, 0x23, 0xe0, 0x69, 0xca, 0xa6, 0x9c, 0xaa, 0xfa, 0x07, 0xa9, 0xa8, 0xe7, 0xe0, 0x40, 0x60,
	0x42, 0x54, 0x34, 0xb1, 0x12, 0x6d, 0x29, 0x51, 0x0e, 0x2a, 0xd1, 0xc3, 0x19, 0x9a, 0xb1, 0xbc,
	0xcd, 0x56, 0x67, 0x68, 0xc6, 0xd0, 0x19, 0xec, 0xab, 0x86, 0xf6, 0x43, 0x3f, 0x58, 0x04, 0xcb,
	0xc6, 0x2e, 0xcb, 0x53, 0xef, 0xca, 0xc6, 0x56, 0x54, 0xde, 0xdf, 0x87, 0x50, 0xb9, 0x66, 0x29,
	0x17, 0x73, 0xbc, 0x55, 0x91, 0x16, 0xcb, 0x62, 0x3d, 0xe0, 0xf2, 0x59, 0x12, 0xd3, 0x3d, 0x11,
	0x73, 0xa5, 0xaa, 0xa8, 0x1b, 0xce, 0x89, 0x48, 0xe6, 0xd2, 0x0d, 0xfb, 0xb2, 0xe6, 0xa6, 0xb6,
	0xe2, 0x46, 0x51, 0xb9, 0x9b, 0x37, 0xb0, 0xcb, 0xbf, 0x64, 0x09, 0xa3, 0x51, 0xcc, 0x6e, 0x17,
	0x9c, 0x7a, 0x2c, 0x63, 0xad, 0xba, 0x4c, 0x73, 0x53, 0x12, 0xb6, 0xc4, 0xfb, 0x2c, 0x63, 0xdd,
	0x7f, 0x4b, 0xb0, 0x2d, 0xa7, 0x74, 0x82, 0xfa, 0x50, 0x13, 0x53, 0x3b, 0x7f, 0x30, 0xd1, 0xe1,
	0xca, 0x5c, 0x5b, 0xff, 0x1e, 0x68, 0xb7, 0x1f, 0xa3, 0xf2, 0x21, 0xff, 0x01, 0x34, 0x9c, 0x66,
	0x7e, 0x20, 0x06, 0x60, 0xfe, 0x94, 0xa1, 0x55, 0xfd, 0x83, 0x37, 0xb2, 0xfd, 0xfc, 0x51, 0x2e,
	0x37, 0x36, 0x54, 0x47, 0xca, 0x1f, 0x12, 0x74, 0xb4, 0xa2, 0xfd, 0xfa, 0xf5, 0x6a, 0xbf, 0xfc,
	0x3f, 0x5a, 0x59, 0x3b, 0x7f, 0xf3, 0xeb, 0xc9, 0xd4, 0xcf, 0x66, 0x8b, 0xeb, 0x53, 0x37, 0x0a,
	0xce, 0x3c, 0xee, 0x26, 0xdc, 0x3b, 0xf3, 0xdc, 0x64, 0x1e, 0x7a, 0x67, 0xf2, 0x15, 0x3b, 0x5b,
	0xee, 0xbf, 0xde, 0x96, 0x1f, 0x43, 0x3f, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xa0, 0x05,
	0x34, 0x3c, 0x09, 0x00, 0x00,
}
