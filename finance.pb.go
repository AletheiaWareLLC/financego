// Code generated by protoc-gen-go. DO NOT EDIT.
// source: finance.proto

package financego

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

type PaymentProcessor int32

const (
	PaymentProcessor_UNKNOWN_PROCESSOR PaymentProcessor = 0
	PaymentProcessor_STRIPE            PaymentProcessor = 1
)

var PaymentProcessor_name = map[int32]string{
	0: "UNKNOWN_PROCESSOR",
	1: "STRIPE",
}

var PaymentProcessor_value = map[string]int32{
	"UNKNOWN_PROCESSOR": 0,
	"STRIPE":            1,
}

func (x PaymentProcessor) String() string {
	return proto.EnumName(PaymentProcessor_name, int32(x))
}

func (PaymentProcessor) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{0}
}

type Service_Interval int32

const (
	Service_UNKNOWN_INTERVAL Service_Interval = 0
	Service_DAILY            Service_Interval = 1
	Service_WEEKLY           Service_Interval = 2
	Service_MONTHLY          Service_Interval = 3
	Service_QUARTERLY        Service_Interval = 4
	Service_YEARLY           Service_Interval = 5
)

var Service_Interval_name = map[int32]string{
	0: "UNKNOWN_INTERVAL",
	1: "DAILY",
	2: "WEEKLY",
	3: "MONTHLY",
	4: "QUARTERLY",
	5: "YEARLY",
}

var Service_Interval_value = map[string]int32{
	"UNKNOWN_INTERVAL": 0,
	"DAILY":            1,
	"WEEKLY":           2,
	"MONTHLY":          3,
	"QUARTERLY":        4,
	"YEARLY":           5,
}

func (x Service_Interval) String() string {
	return proto.EnumName(Service_Interval_name, int32(x))
}

func (Service_Interval) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{7, 0}
}

type Service_Mode int32

const (
	Service_UNKNOWN_MODE            Service_Mode = 0
	Service_FIXED_AMOUNT            Service_Mode = 1
	Service_METERED_SUM_USAGE       Service_Mode = 2
	Service_METERED_MAX_USAGE       Service_Mode = 3
	Service_METERED_LAST_USAGE      Service_Mode = 4
	Service_METERED_LAST_USAGE_EVER Service_Mode = 5
)

var Service_Mode_name = map[int32]string{
	0: "UNKNOWN_MODE",
	1: "FIXED_AMOUNT",
	2: "METERED_SUM_USAGE",
	3: "METERED_MAX_USAGE",
	4: "METERED_LAST_USAGE",
	5: "METERED_LAST_USAGE_EVER",
}

var Service_Mode_value = map[string]int32{
	"UNKNOWN_MODE":            0,
	"FIXED_AMOUNT":            1,
	"METERED_SUM_USAGE":       2,
	"METERED_MAX_USAGE":       3,
	"METERED_LAST_USAGE":      4,
	"METERED_LAST_USAGE_EVER": 5,
}

func (x Service_Mode) String() string {
	return proto.EnumName(Service_Mode_name, int32(x))
}

func (Service_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{7, 1}
}

type Charge struct {
	MerchantAlias        string           `protobuf:"bytes,1,opt,name=merchant_alias,json=merchantAlias,proto3" json:"merchant_alias,omitempty"`
	CustomerAlias        string           `protobuf:"bytes,2,opt,name=customer_alias,json=customerAlias,proto3" json:"customer_alias,omitempty"`
	Processor            PaymentProcessor `protobuf:"varint,3,opt,name=processor,proto3,enum=finance.PaymentProcessor" json:"processor,omitempty"`
	CustomerId           string           `protobuf:"bytes,4,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	PaymentId            string           `protobuf:"bytes,5,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	ChargeId             string           `protobuf:"bytes,6,opt,name=charge_id,json=chargeId,proto3" json:"charge_id,omitempty"`
	Amount               int64            `protobuf:"varint,7,opt,name=amount,proto3" json:"amount,omitempty"`
	InvoiceId            string           `protobuf:"bytes,8,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	ReceiptUrl           string           `protobuf:"bytes,9,opt,name=receipt_url,json=receiptUrl,proto3" json:"receipt_url,omitempty"`
	ProductId            string           `protobuf:"bytes,10,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	PlanId               string           `protobuf:"bytes,11,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	Country              string           `protobuf:"bytes,12,opt,name=country,proto3" json:"country,omitempty"`
	Currency             string           `protobuf:"bytes,13,opt,name=currency,proto3" json:"currency,omitempty"`
	Description          string           `protobuf:"bytes,14,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Charge) Reset()         { *m = Charge{} }
func (m *Charge) String() string { return proto.CompactTextString(m) }
func (*Charge) ProtoMessage()    {}
func (*Charge) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{0}
}

func (m *Charge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Charge.Unmarshal(m, b)
}
func (m *Charge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Charge.Marshal(b, m, deterministic)
}
func (m *Charge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Charge.Merge(m, src)
}
func (m *Charge) XXX_Size() int {
	return xxx_messageInfo_Charge.Size(m)
}
func (m *Charge) XXX_DiscardUnknown() {
	xxx_messageInfo_Charge.DiscardUnknown(m)
}

var xxx_messageInfo_Charge proto.InternalMessageInfo

func (m *Charge) GetMerchantAlias() string {
	if m != nil {
		return m.MerchantAlias
	}
	return ""
}

func (m *Charge) GetCustomerAlias() string {
	if m != nil {
		return m.CustomerAlias
	}
	return ""
}

func (m *Charge) GetProcessor() PaymentProcessor {
	if m != nil {
		return m.Processor
	}
	return PaymentProcessor_UNKNOWN_PROCESSOR
}

func (m *Charge) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Charge) GetPaymentId() string {
	if m != nil {
		return m.PaymentId
	}
	return ""
}

func (m *Charge) GetChargeId() string {
	if m != nil {
		return m.ChargeId
	}
	return ""
}

func (m *Charge) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Charge) GetInvoiceId() string {
	if m != nil {
		return m.InvoiceId
	}
	return ""
}

func (m *Charge) GetReceiptUrl() string {
	if m != nil {
		return m.ReceiptUrl
	}
	return ""
}

func (m *Charge) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *Charge) GetPlanId() string {
	if m != nil {
		return m.PlanId
	}
	return ""
}

func (m *Charge) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Charge) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Charge) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Invoice struct {
	MerchantAlias        string           `protobuf:"bytes,1,opt,name=merchant_alias,json=merchantAlias,proto3" json:"merchant_alias,omitempty"`
	CustomerAlias        string           `protobuf:"bytes,2,opt,name=customer_alias,json=customerAlias,proto3" json:"customer_alias,omitempty"`
	Processor            PaymentProcessor `protobuf:"varint,3,opt,name=processor,proto3,enum=finance.PaymentProcessor" json:"processor,omitempty"`
	CustomerId           string           `protobuf:"bytes,4,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	PaymentId            string           `protobuf:"bytes,5,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	InvoiceId            string           `protobuf:"bytes,6,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	InvoiceUrl           string           `protobuf:"bytes,7,opt,name=invoice_url,json=invoiceUrl,proto3" json:"invoice_url,omitempty"`
	ProductId            string           `protobuf:"bytes,8,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	PlanId               string           `protobuf:"bytes,9,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	Country              string           `protobuf:"bytes,10,opt,name=country,proto3" json:"country,omitempty"`
	Currency             string           `protobuf:"bytes,11,opt,name=currency,proto3" json:"currency,omitempty"`
	Number               string           `protobuf:"bytes,12,opt,name=number,proto3" json:"number,omitempty"`
	AmountDue            int64            `protobuf:"varint,13,opt,name=amount_due,json=amountDue,proto3" json:"amount_due,omitempty"`
	AmountPaid           int64            `protobuf:"varint,14,opt,name=amount_paid,json=amountPaid,proto3" json:"amount_paid,omitempty"`
	AmountRemaining      int64            `protobuf:"varint,15,opt,name=amount_remaining,json=amountRemaining,proto3" json:"amount_remaining,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Invoice) Reset()         { *m = Invoice{} }
func (m *Invoice) String() string { return proto.CompactTextString(m) }
func (*Invoice) ProtoMessage()    {}
func (*Invoice) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{1}
}

func (m *Invoice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invoice.Unmarshal(m, b)
}
func (m *Invoice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invoice.Marshal(b, m, deterministic)
}
func (m *Invoice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invoice.Merge(m, src)
}
func (m *Invoice) XXX_Size() int {
	return xxx_messageInfo_Invoice.Size(m)
}
func (m *Invoice) XXX_DiscardUnknown() {
	xxx_messageInfo_Invoice.DiscardUnknown(m)
}

var xxx_messageInfo_Invoice proto.InternalMessageInfo

func (m *Invoice) GetMerchantAlias() string {
	if m != nil {
		return m.MerchantAlias
	}
	return ""
}

func (m *Invoice) GetCustomerAlias() string {
	if m != nil {
		return m.CustomerAlias
	}
	return ""
}

func (m *Invoice) GetProcessor() PaymentProcessor {
	if m != nil {
		return m.Processor
	}
	return PaymentProcessor_UNKNOWN_PROCESSOR
}

func (m *Invoice) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Invoice) GetPaymentId() string {
	if m != nil {
		return m.PaymentId
	}
	return ""
}

func (m *Invoice) GetInvoiceId() string {
	if m != nil {
		return m.InvoiceId
	}
	return ""
}

func (m *Invoice) GetInvoiceUrl() string {
	if m != nil {
		return m.InvoiceUrl
	}
	return ""
}

func (m *Invoice) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *Invoice) GetPlanId() string {
	if m != nil {
		return m.PlanId
	}
	return ""
}

func (m *Invoice) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Invoice) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Invoice) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Invoice) GetAmountDue() int64 {
	if m != nil {
		return m.AmountDue
	}
	return 0
}

func (m *Invoice) GetAmountPaid() int64 {
	if m != nil {
		return m.AmountPaid
	}
	return 0
}

func (m *Invoice) GetAmountRemaining() int64 {
	if m != nil {
		return m.AmountRemaining
	}
	return 0
}

type Registration struct {
	MerchantAlias        string           `protobuf:"bytes,1,opt,name=merchant_alias,json=merchantAlias,proto3" json:"merchant_alias,omitempty"`
	CustomerAlias        string           `protobuf:"bytes,2,opt,name=customer_alias,json=customerAlias,proto3" json:"customer_alias,omitempty"`
	Processor            PaymentProcessor `protobuf:"varint,3,opt,name=processor,proto3,enum=finance.PaymentProcessor" json:"processor,omitempty"`
	CustomerId           string           `protobuf:"bytes,4,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	PaymentId            string           `protobuf:"bytes,5,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Registration) Reset()         { *m = Registration{} }
func (m *Registration) String() string { return proto.CompactTextString(m) }
func (*Registration) ProtoMessage()    {}
func (*Registration) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{2}
}

func (m *Registration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Registration.Unmarshal(m, b)
}
func (m *Registration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Registration.Marshal(b, m, deterministic)
}
func (m *Registration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Registration.Merge(m, src)
}
func (m *Registration) XXX_Size() int {
	return xxx_messageInfo_Registration.Size(m)
}
func (m *Registration) XXX_DiscardUnknown() {
	xxx_messageInfo_Registration.DiscardUnknown(m)
}

var xxx_messageInfo_Registration proto.InternalMessageInfo

func (m *Registration) GetMerchantAlias() string {
	if m != nil {
		return m.MerchantAlias
	}
	return ""
}

func (m *Registration) GetCustomerAlias() string {
	if m != nil {
		return m.CustomerAlias
	}
	return ""
}

func (m *Registration) GetProcessor() PaymentProcessor {
	if m != nil {
		return m.Processor
	}
	return PaymentProcessor_UNKNOWN_PROCESSOR
}

func (m *Registration) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Registration) GetPaymentId() string {
	if m != nil {
		return m.PaymentId
	}
	return ""
}

type Subscription struct {
	MerchantAlias        string           `protobuf:"bytes,1,opt,name=merchant_alias,json=merchantAlias,proto3" json:"merchant_alias,omitempty"`
	CustomerAlias        string           `protobuf:"bytes,2,opt,name=customer_alias,json=customerAlias,proto3" json:"customer_alias,omitempty"`
	Processor            PaymentProcessor `protobuf:"varint,3,opt,name=processor,proto3,enum=finance.PaymentProcessor" json:"processor,omitempty"`
	CustomerId           string           `protobuf:"bytes,4,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	PaymentId            string           `protobuf:"bytes,5,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	ProductId            string           `protobuf:"bytes,6,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	PlanId               string           `protobuf:"bytes,7,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	SubscriptionId       string           `protobuf:"bytes,8,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	SubscriptionItemId   string           `protobuf:"bytes,9,opt,name=subscription_item_id,json=subscriptionItemId,proto3" json:"subscription_item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Subscription) Reset()         { *m = Subscription{} }
func (m *Subscription) String() string { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()    {}
func (*Subscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{3}
}

func (m *Subscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subscription.Unmarshal(m, b)
}
func (m *Subscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subscription.Marshal(b, m, deterministic)
}
func (m *Subscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscription.Merge(m, src)
}
func (m *Subscription) XXX_Size() int {
	return xxx_messageInfo_Subscription.Size(m)
}
func (m *Subscription) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscription.DiscardUnknown(m)
}

var xxx_messageInfo_Subscription proto.InternalMessageInfo

func (m *Subscription) GetMerchantAlias() string {
	if m != nil {
		return m.MerchantAlias
	}
	return ""
}

func (m *Subscription) GetCustomerAlias() string {
	if m != nil {
		return m.CustomerAlias
	}
	return ""
}

func (m *Subscription) GetProcessor() PaymentProcessor {
	if m != nil {
		return m.Processor
	}
	return PaymentProcessor_UNKNOWN_PROCESSOR
}

func (m *Subscription) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Subscription) GetPaymentId() string {
	if m != nil {
		return m.PaymentId
	}
	return ""
}

func (m *Subscription) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *Subscription) GetPlanId() string {
	if m != nil {
		return m.PlanId
	}
	return ""
}

func (m *Subscription) GetSubscriptionId() string {
	if m != nil {
		return m.SubscriptionId
	}
	return ""
}

func (m *Subscription) GetSubscriptionItemId() string {
	if m != nil {
		return m.SubscriptionItemId
	}
	return ""
}

type UsageRecord struct {
	MerchantAlias        string           `protobuf:"bytes,1,opt,name=merchant_alias,json=merchantAlias,proto3" json:"merchant_alias,omitempty"`
	CustomerAlias        string           `protobuf:"bytes,2,opt,name=customer_alias,json=customerAlias,proto3" json:"customer_alias,omitempty"`
	Processor            PaymentProcessor `protobuf:"varint,3,opt,name=processor,proto3,enum=finance.PaymentProcessor" json:"processor,omitempty"`
	SubscriptionId       string           `protobuf:"bytes,4,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	UsageRecordId        string           `protobuf:"bytes,5,opt,name=usage_record_id,json=usageRecordId,proto3" json:"usage_record_id,omitempty"`
	Quantity             int64            `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`
	SubscriptionItemId   string           `protobuf:"bytes,7,opt,name=subscription_item_id,json=subscriptionItemId,proto3" json:"subscription_item_id,omitempty"`
	ProductId            string           `protobuf:"bytes,8,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	PlanId               string           `protobuf:"bytes,9,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UsageRecord) Reset()         { *m = UsageRecord{} }
func (m *UsageRecord) String() string { return proto.CompactTextString(m) }
func (*UsageRecord) ProtoMessage()    {}
func (*UsageRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{4}
}

func (m *UsageRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsageRecord.Unmarshal(m, b)
}
func (m *UsageRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsageRecord.Marshal(b, m, deterministic)
}
func (m *UsageRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsageRecord.Merge(m, src)
}
func (m *UsageRecord) XXX_Size() int {
	return xxx_messageInfo_UsageRecord.Size(m)
}
func (m *UsageRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_UsageRecord.DiscardUnknown(m)
}

var xxx_messageInfo_UsageRecord proto.InternalMessageInfo

func (m *UsageRecord) GetMerchantAlias() string {
	if m != nil {
		return m.MerchantAlias
	}
	return ""
}

func (m *UsageRecord) GetCustomerAlias() string {
	if m != nil {
		return m.CustomerAlias
	}
	return ""
}

func (m *UsageRecord) GetProcessor() PaymentProcessor {
	if m != nil {
		return m.Processor
	}
	return PaymentProcessor_UNKNOWN_PROCESSOR
}

func (m *UsageRecord) GetSubscriptionId() string {
	if m != nil {
		return m.SubscriptionId
	}
	return ""
}

func (m *UsageRecord) GetUsageRecordId() string {
	if m != nil {
		return m.UsageRecordId
	}
	return ""
}

func (m *UsageRecord) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *UsageRecord) GetSubscriptionItemId() string {
	if m != nil {
		return m.SubscriptionItemId
	}
	return ""
}

func (m *UsageRecord) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *UsageRecord) GetPlanId() string {
	if m != nil {
		return m.PlanId
	}
	return ""
}

type Merchant struct {
	Alias string `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"`
	// The domain name of the merchant
	// Eg. space.aletheiaware.com
	Domain               string           `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Processor            PaymentProcessor `protobuf:"varint,3,opt,name=processor,proto3,enum=finance.PaymentProcessor" json:"processor,omitempty"`
	PublishableKey       string           `protobuf:"bytes,4,opt,name=publishable_key,json=publishableKey,proto3" json:"publishable_key,omitempty"`
	RegisterUrl          string           `protobuf:"bytes,5,opt,name=register_url,json=registerUrl,proto3" json:"register_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Merchant) Reset()         { *m = Merchant{} }
func (m *Merchant) String() string { return proto.CompactTextString(m) }
func (*Merchant) ProtoMessage()    {}
func (*Merchant) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{5}
}

func (m *Merchant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Merchant.Unmarshal(m, b)
}
func (m *Merchant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Merchant.Marshal(b, m, deterministic)
}
func (m *Merchant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Merchant.Merge(m, src)
}
func (m *Merchant) XXX_Size() int {
	return xxx_messageInfo_Merchant.Size(m)
}
func (m *Merchant) XXX_DiscardUnknown() {
	xxx_messageInfo_Merchant.DiscardUnknown(m)
}

var xxx_messageInfo_Merchant proto.InternalMessageInfo

func (m *Merchant) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *Merchant) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Merchant) GetProcessor() PaymentProcessor {
	if m != nil {
		return m.Processor
	}
	return PaymentProcessor_UNKNOWN_PROCESSOR
}

func (m *Merchant) GetPublishableKey() string {
	if m != nil {
		return m.PublishableKey
	}
	return ""
}

func (m *Merchant) GetRegisterUrl() string {
	if m != nil {
		return m.RegisterUrl
	}
	return ""
}

type Product struct {
	ProductId            string   `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Country              string   `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Currency             string   `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	UnitPrice            int64    `protobuf:"varint,4,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	PurchaseUrl          string   `protobuf:"bytes,5,opt,name=purchase_url,json=purchaseUrl,proto3" json:"purchase_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{6}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *Product) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Product) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Product) GetUnitPrice() int64 {
	if m != nil {
		return m.UnitPrice
	}
	return 0
}

func (m *Product) GetPurchaseUrl() string {
	if m != nil {
		return m.PurchaseUrl
	}
	return ""
}

type Service struct {
	ProductId            string           `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	PlanId               string           `protobuf:"bytes,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	Country              string           `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Currency             string           `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	GroupPrice           int64            `protobuf:"varint,5,opt,name=group_price,json=groupPrice,proto3" json:"group_price,omitempty"`
	GroupSize            int64            `protobuf:"varint,6,opt,name=group_size,json=groupSize,proto3" json:"group_size,omitempty"`
	Interval             Service_Interval `protobuf:"varint,7,opt,name=interval,proto3,enum=finance.Service_Interval" json:"interval,omitempty"`
	Mode                 Service_Mode     `protobuf:"varint,8,opt,name=mode,proto3,enum=finance.Service_Mode" json:"mode,omitempty"`
	SubscribeUrl         string           `protobuf:"bytes,9,opt,name=subscribe_url,json=subscribeUrl,proto3" json:"subscribe_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04e2e1c1ba79a81, []int{7}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *Service) GetPlanId() string {
	if m != nil {
		return m.PlanId
	}
	return ""
}

func (m *Service) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Service) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Service) GetGroupPrice() int64 {
	if m != nil {
		return m.GroupPrice
	}
	return 0
}

func (m *Service) GetGroupSize() int64 {
	if m != nil {
		return m.GroupSize
	}
	return 0
}

func (m *Service) GetInterval() Service_Interval {
	if m != nil {
		return m.Interval
	}
	return Service_UNKNOWN_INTERVAL
}

func (m *Service) GetMode() Service_Mode {
	if m != nil {
		return m.Mode
	}
	return Service_UNKNOWN_MODE
}

func (m *Service) GetSubscribeUrl() string {
	if m != nil {
		return m.SubscribeUrl
	}
	return ""
}

func init() {
	proto.RegisterEnum("finance.PaymentProcessor", PaymentProcessor_name, PaymentProcessor_value)
	proto.RegisterEnum("finance.Service_Interval", Service_Interval_name, Service_Interval_value)
	proto.RegisterEnum("finance.Service_Mode", Service_Mode_name, Service_Mode_value)
	proto.RegisterType((*Charge)(nil), "finance.Charge")
	proto.RegisterType((*Invoice)(nil), "finance.Invoice")
	proto.RegisterType((*Registration)(nil), "finance.Registration")
	proto.RegisterType((*Subscription)(nil), "finance.Subscription")
	proto.RegisterType((*UsageRecord)(nil), "finance.UsageRecord")
	proto.RegisterType((*Merchant)(nil), "finance.Merchant")
	proto.RegisterType((*Product)(nil), "finance.Product")
	proto.RegisterType((*Service)(nil), "finance.Service")
}

func init() { proto.RegisterFile("finance.proto", fileDescriptor_c04e2e1c1ba79a81) }

var fileDescriptor_c04e2e1c1ba79a81 = []byte{
	// 1021 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0xdb, 0x6e, 0xe3, 0x54,
	0x14, 0xad, 0xe3, 0x24, 0x8e, 0x77, 0x92, 0xd6, 0x1c, 0x75, 0x3a, 0x86, 0x51, 0x35, 0x25, 0xdc,
	0x3a, 0x3c, 0x74, 0xd0, 0xa0, 0x11, 0xcf, 0x9e, 0xd6, 0x03, 0x56, 0x73, 0xc3, 0x49, 0xa6, 0x53,
	0x5e, 0x2c, 0xc7, 0x3e, 0xa4, 0x47, 0xc4, 0x17, 0x8e, 0xed, 0xa2, 0xcc, 0x1b, 0x1f, 0xc0, 0x3f,
	0xf0, 0x07, 0x7c, 0x01, 0x7f, 0x01, 0xff, 0x80, 0xc4, 0x4f, 0xa0, 0x73, 0x71, 0xeb, 0xa6, 0xcd,
	0x20, 0xc1, 0xcb, 0xa8, 0x6f, 0xd9, 0x6b, 0xaf, 0x73, 0xb2, 0xf6, 0xda, 0xdb, 0xdb, 0x86, 0xee,
	0xf7, 0x24, 0xf6, 0xe3, 0x00, 0x1f, 0xa5, 0x34, 0xc9, 0x13, 0xa4, 0xc9, 0xb0, 0xf7, 0x87, 0x0a,
	0xcd, 0xe3, 0x0b, 0x9f, 0x2e, 0x30, 0xfa, 0x04, 0xb6, 0x23, 0x4c, 0x83, 0x0b, 0x3f, 0xce, 0x3d,
	0x7f, 0x49, 0xfc, 0xcc, 0x54, 0x0e, 0x94, 0x43, 0xdd, 0xed, 0x96, 0xa8, 0xc5, 0x40, 0x46, 0x0b,
	0x8a, 0x2c, 0x4f, 0x22, 0x4c, 0x25, 0xad, 0x26, 0x68, 0x25, 0x2a, 0x68, 0x5f, 0x81, 0x9e, 0xd2,
	0x24, 0xc0, 0x59, 0x96, 0x50, 0x53, 0x3d, 0x50, 0x0e, 0xb7, 0x9f, 0xbd, 0x7f, 0x54, 0x8a, 0x18,
	0xfb, 0xab, 0x08, 0xc7, 0xf9, 0xb8, 0x24, 0xb8, 0xd7, 0x5c, 0xf4, 0x18, 0xda, 0x57, 0xf7, 0x93,
	0xd0, 0xac, 0xf3, 0xcb, 0xa1, 0x84, 0x9c, 0x10, 0xed, 0x03, 0xa4, 0xe2, 0x3c, 0xcb, 0x37, 0x78,
	0x5e, 0x97, 0x88, 0x13, 0xa2, 0x47, 0xa0, 0x07, 0xbc, 0x20, 0x96, 0x6d, 0xf2, 0x6c, 0x4b, 0x00,
	0x4e, 0x88, 0xf6, 0xa0, 0xe9, 0x47, 0x49, 0x11, 0xe7, 0xa6, 0x76, 0xa0, 0x1c, 0xaa, 0xae, 0x8c,
	0xd8, 0x9d, 0x24, 0xbe, 0x4c, 0x48, 0xc0, 0x4f, 0xb5, 0xc4, 0x9d, 0x12, 0x71, 0x42, 0xa6, 0x89,
	0xe2, 0x00, 0x93, 0x34, 0xf7, 0x0a, 0xba, 0x34, 0x75, 0xa1, 0x49, 0x42, 0x33, 0xba, 0xe4, 0x9a,
	0x68, 0x12, 0x16, 0x01, 0xd7, 0x04, 0x52, 0x93, 0x40, 0x9c, 0x10, 0x3d, 0x04, 0x2d, 0x5d, 0xfa,
	0x31, 0xcb, 0xb5, 0x79, 0xae, 0xc9, 0x42, 0x27, 0x44, 0x26, 0x68, 0x01, 0x13, 0x40, 0x57, 0x66,
	0x87, 0x27, 0xca, 0x10, 0x7d, 0x00, 0xad, 0xa0, 0xa0, 0x14, 0xc7, 0xc1, 0xca, 0xec, 0xca, 0x2a,
	0x64, 0x8c, 0x0e, 0xa0, 0x1d, 0xe2, 0x2c, 0xa0, 0x24, 0xcd, 0x49, 0x12, 0x9b, 0xdb, 0x3c, 0x5d,
	0x85, 0x7a, 0x3f, 0xd7, 0x41, 0x73, 0x84, 0xfc, 0xfb, 0xd2, 0xd7, 0x9b, 0x2d, 0x6a, 0xde, 0xd1,
	0xa2, 0x32, 0xcd, 0x5a, 0xa4, 0x89, 0xeb, 0x25, 0x74, 0xbb, 0x45, 0xad, 0xb7, 0xb4, 0x48, 0xdf,
	0xd4, 0x22, 0xd8, 0xdc, 0xa2, 0xf6, 0x5a, 0x8b, 0xf6, 0xa0, 0x19, 0x17, 0xd1, 0x1c, 0x53, 0xd9,
	0x57, 0x19, 0x31, 0x15, 0x62, 0xe4, 0xbc, 0xb0, 0xc0, 0xbc, 0xb1, 0xaa, 0xab, 0x0b, 0xe4, 0xa4,
	0xc0, 0xac, 0x0a, 0x99, 0x4e, 0x7d, 0x12, 0xf2, 0xce, 0xaa, 0xae, 0x3c, 0x31, 0xf6, 0x49, 0x88,
	0x9e, 0x80, 0x21, 0x09, 0x14, 0x47, 0x3e, 0x89, 0x49, 0xbc, 0x30, 0x77, 0x38, 0x6b, 0x47, 0xe0,
	0x6e, 0x09, 0xf7, 0xfe, 0x54, 0xa0, 0xe3, 0xe2, 0x05, 0xc9, 0x72, 0xea, 0xb3, 0xa1, 0xb8, 0x27,
	0x83, 0xd0, 0xfb, 0xab, 0x06, 0x9d, 0x49, 0x31, 0xbf, 0x1a, 0xf6, 0x7b, 0x34, 0xe0, 0x95, 0x01,
	0x6d, 0xbe, 0x65, 0x40, 0xb5, 0x1b, 0x03, 0xfa, 0x19, 0xec, 0x64, 0x15, 0x3b, 0xae, 0xa7, 0x7b,
	0xbb, 0x0a, 0x3b, 0x21, 0xfa, 0x02, 0x76, 0x6f, 0x12, 0x73, 0x1c, 0x5d, 0xcf, 0x3b, 0xba, 0xc1,
	0xce, 0x71, 0xe4, 0x84, 0xbd, 0xbf, 0x6b, 0xd0, 0x9e, 0x65, 0xfe, 0x02, 0xbb, 0x38, 0x48, 0x68,
	0xf8, 0xae, 0x38, 0x7d, 0x47, 0xc5, 0xf5, 0x3b, 0x2b, 0xfe, 0x14, 0x76, 0x0a, 0x26, 0xdf, 0xa3,
	0x5c, 0xff, 0xb5, 0xed, 0xdd, 0xe2, 0xba, 0x2a, 0x27, 0x64, 0x4f, 0xf2, 0x8f, 0x85, 0x1f, 0xe7,
	0x24, 0x5f, 0x71, 0xe3, 0x55, 0xf7, 0x2a, 0xde, 0xe8, 0x9a, 0xb6, 0xc9, 0xb5, 0xff, 0xba, 0x69,
	0x7a, 0xbf, 0x2b, 0xd0, 0x1a, 0x48, 0x23, 0xd1, 0x2e, 0x34, 0xaa, 0x0e, 0x8b, 0x80, 0xad, 0x95,
	0x30, 0x61, 0x0f, 0xb8, 0x74, 0x54, 0x46, 0xff, 0xcb, 0xca, 0xb4, 0x98, 0x2f, 0x49, 0x76, 0xe1,
	0xcf, 0x97, 0xd8, 0xfb, 0x01, 0xaf, 0x4a, 0x2b, 0x2b, 0xf0, 0x29, 0x5e, 0xa1, 0x0f, 0xa1, 0x43,
	0xf9, 0x32, 0xc1, 0x94, 0x2f, 0x58, 0xe1, 0x63, 0xbb, 0xc4, 0x66, 0x74, 0xd9, 0xfb, 0x55, 0x01,
	0x6d, 0x2c, 0xca, 0x5c, 0xf3, 0x40, 0x59, 0xf7, 0xa0, 0xb2, 0x54, 0x6b, 0x9b, 0x97, 0xaa, 0xba,
	0xb6, 0x54, 0xf7, 0x01, 0x8a, 0x98, 0xe4, 0x5e, 0x4a, 0x49, 0x80, 0xb9, 0x4e, 0xd5, 0xd5, 0x19,
	0x32, 0x66, 0x00, 0x93, 0x98, 0x16, 0xcc, 0xbe, 0x0c, 0x57, 0x25, 0x96, 0x18, 0x93, 0xf8, 0x5b,
	0x1d, 0xb4, 0x09, 0xa6, 0x97, 0x8c, 0xfe, 0x2f, 0x12, 0x2b, 0x6d, 0xaa, 0x6d, 0x7a, 0x21, 0xa8,
	0x9b, 0xb5, 0xd7, 0xd7, 0xb4, 0x3f, 0x86, 0xf6, 0x82, 0x26, 0x45, 0x2a, 0xc5, 0x37, 0xc4, 0x66,
	0xe7, 0x90, 0x50, 0xbf, 0x0f, 0x22, 0xf2, 0x32, 0xf2, 0x06, 0xcb, 0x29, 0xd4, 0x39, 0x32, 0x21,
	0x6f, 0x30, 0x7a, 0x0e, 0x2d, 0x12, 0xe7, 0x98, 0x5e, 0xfa, 0xe2, 0xe5, 0x56, 0x6d, 0xb0, 0xac,
	0xe8, 0xc8, 0x91, 0x04, 0xf7, 0x8a, 0x8a, 0x9e, 0x40, 0x3d, 0x4a, 0x42, 0xcc, 0xa7, 0x70, 0xfb,
	0xd9, 0x83, 0x5b, 0x47, 0x06, 0x49, 0x88, 0x5d, 0x4e, 0x41, 0x1f, 0x41, 0x57, 0x0e, 0xf3, 0x1c,
	0x57, 0x3e, 0x73, 0x3a, 0x57, 0x20, 0x33, 0xd0, 0x83, 0x56, 0xf9, 0x2f, 0x68, 0x17, 0x8c, 0xd9,
	0xf0, 0x74, 0x38, 0x3a, 0x1b, 0x7a, 0xce, 0x70, 0x6a, 0xbb, 0xaf, 0xac, 0xbe, 0xb1, 0x85, 0x74,
	0x68, 0x9c, 0x58, 0x4e, 0xff, 0xdc, 0x50, 0x10, 0x40, 0xf3, 0xcc, 0xb6, 0x4f, 0xfb, 0xe7, 0x46,
	0x0d, 0xb5, 0x41, 0x1b, 0x8c, 0x86, 0xd3, 0x6f, 0xfa, 0xe7, 0x86, 0x8a, 0xba, 0xa0, 0x7f, 0x3b,
	0xb3, 0xdc, 0xa9, 0xed, 0xf6, 0xcf, 0x8d, 0x3a, 0xe3, 0x9d, 0xdb, 0x16, 0xfb, 0xdd, 0xe8, 0xfd,
	0xa2, 0x40, 0x9d, 0x89, 0x42, 0x06, 0x74, 0xca, 0xdb, 0x07, 0xa3, 0x13, 0xdb, 0xd8, 0x62, 0xc8,
	0x4b, 0xe7, 0xb5, 0x7d, 0xe2, 0x59, 0x83, 0xd1, 0x6c, 0x38, 0x35, 0x14, 0xf4, 0x00, 0xde, 0x1b,
	0xd8, 0x53, 0xdb, 0xb5, 0x4f, 0xbc, 0xc9, 0x6c, 0xe0, 0xcd, 0x26, 0xd6, 0xd7, 0xb6, 0x51, 0xab,
	0xc2, 0x03, 0xeb, 0xb5, 0x84, 0x55, 0xb4, 0x07, 0xa8, 0x84, 0xfb, 0xd6, 0x64, 0x2a, 0xf1, 0x3a,
	0x7a, 0x04, 0x0f, 0x6f, 0xe3, 0x9e, 0xfd, 0xca, 0x76, 0x8d, 0xc6, 0xe7, 0xcf, 0xc1, 0x58, 0x7f,
	0x7e, 0xd8, 0xfd, 0xa5, 0xb4, 0xb1, 0x3b, 0x3a, 0xb6, 0x27, 0x93, 0x91, 0x6b, 0x6c, 0xb1, 0x32,
	0x26, 0x53, 0xd7, 0x19, 0xdb, 0x86, 0xf2, 0x62, 0x08, 0x66, 0x90, 0x44, 0x47, 0xfe, 0x12, 0xe7,
	0x17, 0x98, 0xf8, 0x3f, 0xf9, 0x14, 0x97, 0xde, 0xbf, 0xe8, 0xbc, 0x14, 0x3f, 0xc6, 0xec, 0x53,
	0xfc, 0xbb, 0x8f, 0x17, 0x24, 0xbf, 0x28, 0xe6, 0x47, 0x41, 0x12, 0x3d, 0xb5, 0x24, 0xfd, 0xcc,
	0xa7, 0xb8, 0xdf, 0x3f, 0x7e, 0x2a, 0x4f, 0x2c, 0x92, 0x79, 0x93, 0x7f, 0xb7, 0x7f, 0xf9, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x35, 0x2b, 0x4b, 0xc8, 0x0b, 0x00, 0x00,
}