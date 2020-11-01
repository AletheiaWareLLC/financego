package financego_test

import (
	"github.com/AletheiaWareLLC/financego"
)

var _ financego.Processor = (*mockProvider)(nil)

type mockProvider struct {
	charge              *financego.Charge
	registration        *financego.Registration
	customerCharge      *financego.Charge
	subscription        *financego.Subscription
	usageRecord         *financego.UsageRecord
	chargeError         error
	registrationError   error
	customerChargeError error
	subscriptionError   error
	usageRecordError    error
}

func (p *mockProvider) NewCharge(merchantAlias, customerAlias, paymentId, productId, planId, country, currency string, amount int64, description string) (*financego.Charge, error) {
	return p.charge, p.chargeError
}

func (p *mockProvider) NewRegistration(merchantAlias, customerAlias, email, paymentId, description string) (*financego.Registration, error) {
	return p.registration, p.registrationError
}

func (p *mockProvider) NewCustomerCharge(registration *financego.Registration, productId, planId, country, currency string, amount int64, description string) (*financego.Charge, error) {
	return p.customerCharge, p.customerChargeError
}

func (p *mockProvider) NewSubscription(merchantAlias, customerAlias, customerId, paymentId, productId, planId string) (*financego.Subscription, error) {
	return p.subscription, p.subscriptionError
}

func (p *mockProvider) NewUsageRecord(merchantAlias, customerAlias, subscriptionId, subscriptionItemId, productId, planId string, timestamp int64, size int64) (*financego.UsageRecord, error) {
	return p.usageRecord, p.usageRecordError
}
