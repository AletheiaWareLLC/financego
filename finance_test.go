/*
 * Copyright 2020 Aletheia Ware LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package financego_test

import (
	"aletheiaware.com/financego"
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
