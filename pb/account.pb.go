// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: account.proto

package API_Presidio

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type AccountResponseErrorCodes int32

const (
	AccountResponse_NO_ERROR            AccountResponseErrorCodes = 0
	AccountResponse_INTERNAL_ERROR      AccountResponseErrorCodes = 1
	AccountResponse_INVALID_PAGE_NUMBER AccountResponseErrorCodes = 2
	AccountResponse_INVALID_PAGE_SIZE   AccountResponseErrorCodes = 3
)

var AccountResponseErrorCodes_name = map[int32]string{
	0: "NO_ERROR",
	1: "INTERNAL_ERROR",
	2: "INVALID_PAGE_NUMBER",
	3: "INVALID_PAGE_SIZE",
}

var AccountResponseErrorCodes_value = map[string]int32{
	"NO_ERROR":            0,
	"INTERNAL_ERROR":      1,
	"INVALID_PAGE_NUMBER": 2,
	"INVALID_PAGE_SIZE":   3,
}

func (x AccountResponseErrorCodes) String() string {
	return proto.EnumName(AccountResponseErrorCodes_name, int32(x))
}

func (AccountResponseErrorCodes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{3, 0}
}

//*
// The request to the account API can specify a number filters, such
// as which addresses to retrieve information about or general
// health requirements. The following shows an example set of request
// parameters in JSON:
// <code>{
//   "addresses": [] // returns all accounts if empty or not included
//   "block_number": 0 // returns latest if given 0
//   "max_health": { "value": "10.0" }
//   "min_borrow_value_in_eth": { "value": "0.002" }
//   "page_number": 1
//   "page_size": 10
// }</code>
type AccountRequest struct {
	// List of account addresses to filter on, e.g.: ["0x...", ,"0x..."] [Optional]
	Addresses [][]byte `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	// Filter for accounts which total outstanding borrows exceeding given amount. [Optional]
	MinBorrowValueInEth *Precise `protobuf:"bytes,2,opt,name=min_borrow_value_in_eth,json=minBorrowValueInEth,proto3" json:"min_borrow_value_in_eth,omitempty"`
	// Filter for accounts where outstanding borrows divided by collateral value is less than the provided amount.
	// If returned value is less than 1.0, for instance, the account is subject to liquidation. If provided,
	// should be given as `{ "value": "...string formatted number..." }` [Optional]
	MaxHealth *Precise `protobuf:"bytes,3,opt,name=max_health,json=maxHealth,proto3" json:"max_health,omitempty"`
	// If provided, API returns data for given block number from our historical data. Otherwise, API defaults to returning the latest information. [Optional]
	BlockNumber uint32 `protobuf:"varint,4,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	// If provided, API returns data for given timestamp from our historical data. Otherwise, API defaults to returning the latest information. [Optional]
	BlockTimestamp uint32 `protobuf:"varint,5,opt,name=block_timestamp,json=blockTimestamp,proto3" json:"block_timestamp,omitempty"`
	// Number of accounts to include in the response, default is 10 e.g. page_size=10 [Optional]
	PageSize uint32 `protobuf:"varint,6,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Pagination number for accounts in the response, default is 1 e.g. page_number=1 [Optional]
	PageNumber           uint32   `protobuf:"varint,7,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountRequest) Reset()         { *m = AccountRequest{} }
func (m *AccountRequest) String() string { return proto.CompactTextString(m) }
func (*AccountRequest) ProtoMessage()    {}
func (*AccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}
func (m *AccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountRequest.Unmarshal(m, b)
}
func (m *AccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountRequest.Marshal(b, m, deterministic)
}
func (m *AccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRequest.Merge(m, src)
}
func (m *AccountRequest) XXX_Size() int {
	return xxx_messageInfo_AccountRequest.Size(m)
}
func (m *AccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRequest proto.InternalMessageInfo

func (m *AccountRequest) GetAddresses() [][]byte {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *AccountRequest) GetMinBorrowValueInEth() *Precise {
	if m != nil {
		return m.MinBorrowValueInEth
	}
	return nil
}

func (m *AccountRequest) GetMaxHealth() *Precise {
	if m != nil {
		return m.MaxHealth
	}
	return nil
}

func (m *AccountRequest) GetBlockNumber() uint32 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *AccountRequest) GetBlockTimestamp() uint32 {
	if m != nil {
		return m.BlockTimestamp
	}
	return 0
}

func (m *AccountRequest) GetPageSize() uint32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *AccountRequest) GetPageNumber() uint32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

//*
// This includes a list of cTokens contextualized to each account.
// <code>{
//   "address": "0xbac065be2e8ca097e9ac924e94af000dd3a5663"
//   "health": { "value": "1.07264275673050348990755599431194797431802239523113293682619605751591901" }
//   "tokens": [
//     {
//       "address": "0xf5dce57282a584d2746faf1593d3121fcac444dc"
//       "borrow_balance_underlying": {"value": "131.4682716123015"}
//       "lifetime_borrow_interest_accrued": {"value": "0.44430505829286"}
//       "lifetime_supply_interest_accrued": {"value": "0.0000021671829864899976"}
//       "supply_balance_underlying": {"value": "0.0"}
//     }
//   ],
//   "total_borrow_value_in_eth": {"value": "0.5100157047140227313856015174794473200000000000000000000000000000" }
//   "total_collateral_value_in_eth": {"value": "0.54706465148029978664135447293587201124121731200000000000000000000000000" }
// }</code>
type Account struct {
	// The public Ethereum address of the account
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// The value of all collateral supplied by the account. Calculated as <em>cTokens held • exchange rate • collateral factor</em>.
	// Note: assets can be supplied and gain interest without being counted as collateral.
	TotalCollateralValueInEth *Precise `protobuf:"bytes,2,opt,name=total_collateral_value_in_eth,json=totalCollateralValueInEth,proto3" json:"total_collateral_value_in_eth,omitempty"`
	// The value of all outstanding borrows with accumulated interest.
	TotalBorrowValueInEth *Precise `protobuf:"bytes,3,opt,name=total_borrow_value_in_eth,json=totalBorrowValueInEth,proto3" json:"total_borrow_value_in_eth,omitempty"`
	// <pre>total_collateral_value_in_eth / total_borrow_value_in_eth</pre>. If this value is less than 1.0, the account is subject to liquidation.
	Health *Precise `protobuf:"bytes,5,opt,name=health,proto3" json:"health,omitempty"`
	//doc-false
	BlockUpdated int32 `protobuf:"varint,6,opt,name=block_updated,json=blockUpdated,proto3" json:"block_updated,omitempty"`
	// A list of tokens held by this account, see <ref>AccountCToken</ref> below for details.
	Tokens               []*AccountCToken `protobuf:"bytes,7,rep,name=tokens,proto3" json:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Account) GetTotalCollateralValueInEth() *Precise {
	if m != nil {
		return m.TotalCollateralValueInEth
	}
	return nil
}

func (m *Account) GetTotalBorrowValueInEth() *Precise {
	if m != nil {
		return m.TotalBorrowValueInEth
	}
	return nil
}

func (m *Account) GetHealth() *Precise {
	if m != nil {
		return m.Health
	}
	return nil
}

func (m *Account) GetBlockUpdated() int32 {
	if m != nil {
		return m.BlockUpdated
	}
	return 0
}

func (m *Account) GetTokens() []*AccountCToken {
	if m != nil {
		return m.Tokens
	}
	return nil
}

//*
// An account's supply, borrow, and interest information for a particular cToken.
//
// <code>{
//   "address": "0xf5dce57282a584d2746faf1593d3121fcac444dc"
//   "borrow_balance_underlying": {"value": "131.4682716123015"}
//   "lifetime_borrow_interest_accrued": {"value": "0.44430505829286"}
//   "lifetime_supply_interest_accrued": {"value": "0.0000021671829864899976"}
//   "supply_balance_underlying": {"value": "0.0"}
// }</code>
type AccountCToken struct {
	// The address of the cToken
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// The symbol of the cToken
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// The cToken balance converted to underlying tokens <pre>cTokens held • exchange rate</pre>
	SupplyBalanceUnderlying *Precise `protobuf:"bytes,3,opt,name=supply_balance_underlying,json=supplyBalanceUnderlying,proto3" json:"supply_balance_underlying,omitempty"`
	// The borrow balance (this is denominated in the underlying token, not in cTokens)
	BorrowBalanceUnderlying *Precise `protobuf:"bytes,4,opt,name=borrow_balance_underlying,json=borrowBalanceUnderlying,proto3" json:"borrow_balance_underlying,omitempty"`
	// The amount of supply interest accrued for the lifetime of this account-cToken pair.
	LifetimeSupplyInterestAccrued *Precise `protobuf:"bytes,5,opt,name=lifetime_supply_interest_accrued,json=lifetimeSupplyInterestAccrued,proto3" json:"lifetime_supply_interest_accrued,omitempty"`
	// The amount of borrow interest accrued for the lifetime of this account-cToken pair.
	LifetimeBorrowInterestAccrued *Precise `protobuf:"bytes,6,opt,name=lifetime_borrow_interest_accrued,json=lifetimeBorrowInterestAccrued,proto3" json:"lifetime_borrow_interest_accrued,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *AccountCToken) Reset()         { *m = AccountCToken{} }
func (m *AccountCToken) String() string { return proto.CompactTextString(m) }
func (*AccountCToken) ProtoMessage()    {}
func (*AccountCToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{2}
}
func (m *AccountCToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountCToken.Unmarshal(m, b)
}
func (m *AccountCToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountCToken.Marshal(b, m, deterministic)
}
func (m *AccountCToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountCToken.Merge(m, src)
}
func (m *AccountCToken) XXX_Size() int {
	return xxx_messageInfo_AccountCToken.Size(m)
}
func (m *AccountCToken) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountCToken.DiscardUnknown(m)
}

var xxx_messageInfo_AccountCToken proto.InternalMessageInfo

func (m *AccountCToken) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *AccountCToken) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *AccountCToken) GetSupplyBalanceUnderlying() *Precise {
	if m != nil {
		return m.SupplyBalanceUnderlying
	}
	return nil
}

func (m *AccountCToken) GetBorrowBalanceUnderlying() *Precise {
	if m != nil {
		return m.BorrowBalanceUnderlying
	}
	return nil
}

func (m *AccountCToken) GetLifetimeSupplyInterestAccrued() *Precise {
	if m != nil {
		return m.LifetimeSupplyInterestAccrued
	}
	return nil
}

func (m *AccountCToken) GetLifetimeBorrowInterestAccrued() *Precise {
	if m != nil {
		return m.LifetimeBorrowInterestAccrued
	}
	return nil
}

//
// The account API returns an overall picture of accounts matching the filters on Compound.
type AccountResponse struct {
	// If set and non-zero, indicates an error returning data.
	// <pre>NO_ERROR = 0
	// INTERNAL_ERROR = 1
	// INVALID_PAGE_NUMBER = 2
	// INVALID_PAGE_SIZE = 3</pre>
	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	// The request parameters are echoed in the response.
	Request *AccountRequest `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	// For example
	// <pre>{
	//   "page_number": 1,
	//   "page_size": 100,
	//   "total_entries": 83,
	//   "total_pages": 1,
	// }</pre>
	PaginationSummary *PaginationSummary `protobuf:"bytes,3,opt,name=pagination_summary,json=paginationSummary,proto3" json:"pagination_summary,omitempty"`
	// The portion of an outstanding borrow that can be closed in a liquidation,
	// which is a percentage of the total underlying borrow balance. For example if the
	// close factor is 0.1, then an account in liqudation is liable to have 10% of its borrows
	// liquidated.
	CloseFactor float32 `protobuf:"fixed32,4,opt,name=close_factor,json=closeFactor,proto3" json:"close_factor,omitempty"`
	// The amount of extra collateral that will be seized to incentivize liquidation.
	// For example, an incentive of 1.05 implies that a liquidator will receive a 5%
	// bonus on the exchange of collateral during a liquidation
	LiquidationIncentive float32 `protobuf:"fixed32,5,opt,name=liquidation_incentive,json=liquidationIncentive,proto3" json:"liquidation_incentive,omitempty"`
	// The list of accounts (see <ref>Account</ref> below) matching the requested filter,
	// with the associated account and cToken data.
	Accounts             []*Account `protobuf:"bytes,6,rep,name=accounts,proto3" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AccountResponse) Reset()         { *m = AccountResponse{} }
func (m *AccountResponse) String() string { return proto.CompactTextString(m) }
func (*AccountResponse) ProtoMessage()    {}
func (*AccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{3}
}
func (m *AccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountResponse.Unmarshal(m, b)
}
func (m *AccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountResponse.Marshal(b, m, deterministic)
}
func (m *AccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountResponse.Merge(m, src)
}
func (m *AccountResponse) XXX_Size() int {
	return xxx_messageInfo_AccountResponse.Size(m)
}
func (m *AccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountResponse proto.InternalMessageInfo

func (m *AccountResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *AccountResponse) GetRequest() *AccountRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *AccountResponse) GetPaginationSummary() *PaginationSummary {
	if m != nil {
		return m.PaginationSummary
	}
	return nil
}

func (m *AccountResponse) GetCloseFactor() float32 {
	if m != nil {
		return m.CloseFactor
	}
	return 0
}

func (m *AccountResponse) GetLiquidationIncentive() float32 {
	if m != nil {
		return m.LiquidationIncentive
	}
	return 0
}

func (m *AccountResponse) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func init() {
	proto.RegisterEnum("API.Presidio.AccountResponseErrorCodes", AccountResponseErrorCodes_name, AccountResponseErrorCodes_value)
	proto.RegisterType((*AccountRequest)(nil), "API.Presidio.AccountRequest")
	proto.RegisterType((*Account)(nil), "API.Presidio.Account")
	proto.RegisterType((*AccountCToken)(nil), "API.Presidio.AccountCToken")
	proto.RegisterType((*AccountResponse)(nil), "API.Presidio.AccountResponse")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 781 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xc6, 0x76, 0x63, 0x27, 0xc7, 0x9b, 0x34, 0x99, 0x92, 0x66, 0x9b, 0x36, 0xc2, 0x18, 0x24,
	0x2c, 0x21, 0x19, 0x91, 0x48, 0xdc, 0xdb, 0x61, 0x4b, 0x57, 0x2a, 0x6e, 0x34, 0x8e, 0x2b, 0x84,
	0x04, 0xa3, 0xf1, 0xee, 0x34, 0x1e, 0x75, 0x77, 0x66, 0x3b, 0x33, 0x1b, 0xea, 0x5e, 0xc2, 0x23,
	0x70, 0xcd, 0x03, 0xf0, 0x0c, 0x3c, 0x01, 0xd7, 0xbc, 0x02, 0x0f, 0x82, 0x76, 0x66, 0x6c, 0xf2,
	0xb3, 0x84, 0x9b, 0x5e, 0xce, 0xf7, 0x7d, 0xe7, 0x9b, 0xf1, 0x77, 0x8e, 0xcf, 0xc2, 0x36, 0x4d,
	0x12, 0x59, 0x0a, 0x33, 0x2c, 0x94, 0x34, 0x12, 0x05, 0xa3, 0xb3, 0x78, 0x78, 0xa6, 0x98, 0xe6,
	0x29, 0x97, 0x87, 0x7b, 0x54, 0x08, 0x69, 0xa8, 0xe1, 0x52, 0x68, 0x27, 0x38, 0x0c, 0x12, 0x99,
	0xe7, 0x52, 0xb8, 0x53, 0xff, 0xf7, 0x26, 0xec, 0x8c, 0x9c, 0x01, 0x66, 0x6f, 0x4a, 0xa6, 0x0d,
	0x7a, 0x02, 0x5b, 0x34, 0x4d, 0x15, 0xd3, 0x9a, 0xe9, 0xb0, 0xd1, 0x6b, 0x0d, 0x02, 0xfc, 0x2f,
	0x80, 0xc6, 0x70, 0x90, 0x73, 0x41, 0xe6, 0x52, 0x29, 0xf9, 0x13, 0xb9, 0xa4, 0x59, 0xc9, 0x08,
	0x17, 0x84, 0x99, 0x45, 0xd8, 0xec, 0x35, 0x06, 0xdd, 0xe3, 0x60, 0xe8, 0x5f, 0x90, 0x70, 0xcd,
	0xf0, 0x83, 0x9c, 0x8b, 0xb1, 0xd5, 0xbe, 0xac, 0xa4, 0xb1, 0x88, 0xcc, 0x02, 0x7d, 0x0e, 0x90,
	0xd3, 0xb7, 0x64, 0xc1, 0x68, 0x66, 0x16, 0x61, 0xab, 0xa6, 0x6c, 0x2b, 0xa7, 0x6f, 0x9f, 0x59,
	0x1a, 0x7d, 0x0c, 0xc1, 0x3c, 0x93, 0xc9, 0x6b, 0x22, 0xca, 0x7c, 0xce, 0x54, 0x78, 0xaf, 0xd7,
	0x18, 0x6c, 0xe3, 0xae, 0xc5, 0x26, 0x16, 0x42, 0x9f, 0xc1, 0x7d, 0x27, 0x31, 0x3c, 0x67, 0xda,
	0xd0, 0xbc, 0x08, 0x37, 0xac, 0x6a, 0xc7, 0xc2, 0xe7, 0x2b, 0x14, 0x3d, 0x86, 0xad, 0x82, 0x5e,
	0x30, 0xa2, 0xf9, 0x3b, 0x16, 0xb6, 0xad, 0x64, 0xb3, 0x02, 0xa6, 0xfc, 0x1d, 0x43, 0x1f, 0x41,
	0xd7, 0x92, 0xfe, 0x9e, 0x8e, 0xa5, 0xa1, 0x82, 0xdc, 0x35, 0xfd, 0x3f, 0x9a, 0xd0, 0xf1, 0x59,
	0xa1, 0x10, 0x3a, 0x3e, 0x93, 0xb0, 0xd1, 0x6b, 0x0c, 0x02, 0xbc, 0x3a, 0xa2, 0x09, 0x1c, 0x19,
	0x69, 0x68, 0x46, 0x12, 0x99, 0x65, 0xd4, 0x30, 0x45, 0xb3, 0xff, 0x8f, 0xe9, 0x91, 0x2d, 0x39,
	0x5d, 0x57, 0x5c, 0x09, 0xeb, 0x29, 0x38, 0xb2, 0x36, 0xf2, 0xba, 0xec, 0xf6, 0xad, 0xfc, 0x56,
	0xe8, 0x9f, 0x42, 0xdb, 0x07, 0xbe, 0x51, 0x53, 0xe4, 0x39, 0xf4, 0x09, 0x6c, 0xbb, 0x28, 0xcb,
	0x22, 0xa5, 0x86, 0xa5, 0x36, 0xa5, 0x0d, 0xec, 0x5a, 0x30, 0x73, 0x18, 0x3a, 0x81, 0xb6, 0x91,
	0xaf, 0x99, 0xd0, 0x61, 0xa7, 0xd7, 0x1a, 0x74, 0x8f, 0x1f, 0x0f, 0xaf, 0x0e, 0xdd, 0xd0, 0x67,
	0x74, 0x7a, 0x5e, 0x69, 0xb0, 0x97, 0xf6, 0x7f, 0x69, 0xc1, 0xf6, 0x35, 0xe6, 0x8e, 0x0c, 0x1f,
	0x42, 0x5b, 0x2f, 0xf3, 0xb9, 0xcc, 0x6c, 0x58, 0x5b, 0xd8, 0x9f, 0xd0, 0x33, 0x78, 0xa4, 0xcb,
	0xa2, 0xc8, 0x96, 0x64, 0x4e, 0x33, 0x2a, 0x12, 0x46, 0x4a, 0x91, 0x32, 0x95, 0x2d, 0xb9, 0xb8,
	0xa8, 0xcd, 0xe2, 0xc0, 0xc9, 0xc7, 0x4e, 0x3d, 0x5b, 0x8b, 0x2b, 0x27, 0x9f, 0x67, 0x8d, 0xd3,
	0xbd, 0x3a, 0x27, 0x27, 0xbf, 0xed, 0x34, 0x83, 0x5e, 0xc6, 0x5f, 0xb1, 0x6a, 0xf4, 0x88, 0x7f,
	0x1c, 0x17, 0x86, 0x29, 0xa6, 0x0d, 0xa1, 0x49, 0xa2, 0x4a, 0x96, 0xd6, 0x26, 0x7e, 0xb4, 0xaa,
	0x9a, 0xda, 0xa2, 0xd8, 0xd7, 0x8c, 0x5c, 0xc9, 0x35, 0x5b, 0xff, 0xd2, 0x5b, 0xb6, 0xed, 0xbb,
	0x6c, 0xdd, 0x00, 0xdc, 0xb0, 0xed, 0xff, 0xd6, 0x82, 0xfb, 0xeb, 0xff, 0xbb, 0x2e, 0xa4, 0xd0,
	0x0c, 0xf5, 0x60, 0x83, 0x29, 0x25, 0x95, 0xed, 0x42, 0xf7, 0x18, 0xac, 0x5f, 0x54, 0x21, 0xd8,
	0x11, 0xe8, 0x2b, 0xe8, 0x28, 0xb7, 0x1d, 0xfc, 0xf4, 0x3e, 0xa9, 0xed, 0xb8, 0xdf, 0x20, 0x78,
	0x25, 0x46, 0x11, 0xa0, 0x82, 0x5e, 0x70, 0x61, 0x17, 0x10, 0xd1, 0x65, 0x9e, 0x53, 0xb5, 0xf4,
	0x8d, 0x7a, 0xe8, 0x2c, 0xd6, 0xf4, 0xd4, 0xb1, 0x78, 0xaf, 0xb8, 0x09, 0x55, 0x2b, 0x20, 0xc9,
	0xa4, 0x66, 0xe4, 0x15, 0x4d, 0x8c, 0x74, 0x2b, 0xa0, 0x89, 0xbb, 0x16, 0x7b, 0x6a, 0x21, 0x74,
	0x02, 0xfb, 0x19, 0x7f, 0x53, 0xf2, 0xd4, 0x5d, 0xc5, 0x45, 0xc2, 0x84, 0xe1, 0x97, 0xcc, 0x46,
	0xdf, 0xc4, 0x1f, 0x5e, 0x21, 0xe3, 0x15, 0x87, 0xbe, 0x84, 0x4d, 0xbf, 0x3c, 0x75, 0xd8, 0xb6,
	0x93, 0xbc, 0x5f, 0xff, 0xbb, 0xd6, 0xb2, 0xfe, 0x8f, 0x00, 0x36, 0x92, 0x53, 0x99, 0x32, 0x8d,
	0x02, 0xd8, 0x9c, 0xbc, 0x20, 0x11, 0xc6, 0x2f, 0xf0, 0xee, 0x07, 0x08, 0xc1, 0x4e, 0x3c, 0x39,
	0x8f, 0xf0, 0x64, 0xf4, 0xdc, 0x63, 0x0d, 0x74, 0x00, 0x0f, 0xe2, 0xc9, 0xcb, 0xd1, 0xf3, 0xf8,
	0x6b, 0x72, 0x36, 0xfa, 0x26, 0x22, 0x93, 0xd9, 0xb7, 0xe3, 0x08, 0xef, 0x36, 0xd1, 0x3e, 0xec,
	0x5d, 0x23, 0xa6, 0xf1, 0xf7, 0xd1, 0x6e, 0xeb, 0xf8, 0xcf, 0xc6, 0x7a, 0x1f, 0x4f, 0x99, 0xba,
	0xe4, 0x09, 0x43, 0xdf, 0x41, 0xc7, 0x5f, 0x8f, 0xee, 0x8c, 0xfd, 0xf0, 0xe8, 0x3f, 0x58, 0xd7,
	0xe6, 0xfe, 0xee, 0xcf, 0x7f, 0xfd, 0xfd, 0x6b, 0x13, 0xd0, 0xe6, 0x17, 0x2b, 0xbb, 0x1f, 0x20,
	0x28, 0xa4, 0x9b, 0xa7, 0xf7, 0x66, 0xdf, 0x5f, 0xdb, 0xcf, 0xdb, 0xf6, 0x13, 0x73, 0xf2, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x77, 0x32, 0xa6, 0xa2, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServiceClient interface {
	Account(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	PostAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) Account(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/API.Presidio.AccountService/account", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) PostAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/API.Presidio.AccountService/post_account", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	Account(context.Context, *AccountRequest) (*AccountResponse, error)
	PostAccount(context.Context, *AccountRequest) (*AccountResponse, error)
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_Account_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Account(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API.Presidio.AccountService/Account",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Account(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_PostAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).PostAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API.Presidio.AccountService/PostAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).PostAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "API.Presidio.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "account",
			Handler:    _AccountService_Account_Handler,
		},
		{
			MethodName: "post_account",
			Handler:    _AccountService_PostAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}