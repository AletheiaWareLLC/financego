/*
 * Copyright 2019 Aletheia Ware LLC
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

package financego

import (
	"crypto/rsa"
	"github.com/AletheiaWareLLC/bcgo"
	"github.com/golang/protobuf/proto"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/sub"
	"github.com/stripe/stripe-go/usagerecord"
	"github.com/stripe/stripe-go/webhook"
	"log"
	"os"
)

const (
	CHARGE       = "Charge"
	REGISTRATION = "Registration"
	SUBSCRIPTION = "Subscription"
	USAGE_RECORD = "UsageRecord"
)

func OpenChargeChannel() *bcgo.PoWChannel {
	return bcgo.OpenPoWChannel(CHARGE, bcgo.THRESHOLD_STANDARD)
}

func OpenRegistrationChannel() *bcgo.PoWChannel {
	return bcgo.OpenPoWChannel(REGISTRATION, bcgo.THRESHOLD_STANDARD)
}

func OpenSubscriptionChannel() *bcgo.PoWChannel {
	return bcgo.OpenPoWChannel(SUBSCRIPTION, bcgo.THRESHOLD_STANDARD)
}

func OpenUsageRecordChannel() *bcgo.PoWChannel {
	return bcgo.OpenPoWChannel(USAGE_RECORD, bcgo.THRESHOLD_STANDARD)
}

func NewCharge(merchantAlias, customerAlias, paymentId string, amount int64, description string) (*stripe.Charge, *Charge, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	chargeParams := &stripe.ChargeParams{
		Amount:      stripe.Int64(amount),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Description: stripe.String(description),
	}
	chargeParams.SetSource(paymentId)
	ch, err := charge.New(chargeParams)
	if err != nil {
		return nil, nil, err
	}

	charge := &Charge{
		MerchantAlias: merchantAlias,
		CustomerAlias: customerAlias,
		Processor:     PaymentProcessor_STRIPE,
		PaymentId:     paymentId,
		ChargeId:      ch.ID,
	}
	log.Println("Charge", charge)
	return ch, charge, nil
}

func NewRegistration(merchantAlias, customerAlias, email, paymentId, description string) (*stripe.Customer, *Registration, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	// Create new Stripe customer
	customerParams := &stripe.CustomerParams{
		Description: stripe.String(description),
		Email:       stripe.String(email),
	}
	if err := customerParams.SetSource(paymentId); err != nil {
		return nil, nil, err
	}
	c, err := customer.New(customerParams)
	if err != nil {
		return nil, nil, err
	}

	registration := &Registration{
		MerchantAlias: merchantAlias,
		CustomerAlias: customerAlias,
		Processor:     PaymentProcessor_STRIPE,
		CustomerId:    c.ID,
		PaymentId:     paymentId,
	}
	log.Println("Registration", registration)
	return c, registration, nil
}

func NewCustomerCharge(registration *Registration, amount int64, description string) (*stripe.Charge, *Charge, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	chargeParams := &stripe.ChargeParams{
		Amount:      stripe.Int64(amount),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Customer:    stripe.String(registration.CustomerId),
		Description: stripe.String(description),
	}
	ch, err := charge.New(chargeParams)
	if err != nil {
		return nil, nil, err
	}

	charge := &Charge{
		MerchantAlias: registration.MerchantAlias,
		CustomerAlias: registration.CustomerAlias,
		Processor:     PaymentProcessor_STRIPE,
		CustomerId:    registration.CustomerId,
		ChargeId:      ch.ID,
	}
	log.Println("Charge", charge)
	return ch, charge, nil
}

func NewSubscription(merchantAlias, customerAlias, customerId, paymentId, productId, planId string) (*stripe.Subscription, *Subscription, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	// Create new Stripe subscription
	subscriptionParams := &stripe.SubscriptionParams{
		Customer: stripe.String(customerId),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Plan: stripe.String(planId),
			},
		},
	}
	s, err := sub.New(subscriptionParams)
	if err != nil {
		return nil, nil, err
	}

	// Create subscription
	subscription := &Subscription{
		MerchantAlias:      merchantAlias,
		CustomerAlias:      customerAlias,
		Processor:          PaymentProcessor_STRIPE,
		CustomerId:         customerId,
		ProductId:          productId,
		PlanId:             planId,
		SubscriptionId:     s.ID,
		SubscriptionItemId: s.Items.Data[0].ID,
	}
	if paymentId != "" {
		subscription.PaymentId = paymentId
	}
	log.Println("Subscription", subscription)
	return s, subscription, nil
}

func NewUsageRecord(merchantAlias, customerAlias, subscription string, timestamp int64, size int64) (*stripe.UsageRecord, *UsageRecord, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	params := &stripe.UsageRecordParams{
		SubscriptionItem: stripe.String(subscription),
		Timestamp:        stripe.Int64(timestamp),
		Quantity:         stripe.Int64(size),
	}
	ur, err := usagerecord.New(params)
	if err != nil {
		return nil, nil, err
	}

	// Create usage record
	usage := &UsageRecord{
		MerchantAlias:  merchantAlias,
		CustomerAlias:  customerAlias,
		Processor:      PaymentProcessor_STRIPE,
		SubscriptionId: subscription,
		UsageRecordId:  ur.ID,
	}
	log.Println("UsageRecord", usage)
	return ur, usage, nil
}

func GetChargeAsync(charges bcgo.Channel, cache bcgo.Cache, merchantAlias string, merchantKey *rsa.PrivateKey, customerAlias string, customerKey *rsa.PrivateKey, callback func(*Charge) error) error {
	if err := bcgo.LoadHead(charges, cache, nil); err != nil {
		log.Println(err)
	}
	head := charges.GetHead()
	cb := func(entry *bcgo.BlockEntry, key, data []byte) error {
		// Unmarshal as Charge
		charge := &Charge{}
		err := proto.Unmarshal(data, charge)
		if err != nil {
			return err
		} else if (merchantAlias == "" || charge.MerchantAlias == merchantAlias) && (customerAlias == "" || charge.CustomerAlias == customerAlias) {
			return callback(charge)
		}
		return nil
	}
	// Read as merchant
	if merchantAlias != "" && merchantKey != nil {
		return bcgo.Read(head, nil, cache, merchantAlias, merchantKey, nil, cb)
	}
	// Read as customer
	return bcgo.Read(head, nil, cache, customerAlias, customerKey, nil, cb)
}

func GetChargeSync(charges bcgo.Channel, cache bcgo.Cache, merchantAlias string, merchantKey *rsa.PrivateKey, customerAlias string, customerKey *rsa.PrivateKey) (*Charge, error) {
	var charge *Charge
	if err := GetChargeAsync(charges, cache, merchantAlias, merchantKey, customerAlias, customerKey, func(c *Charge) error {
		charge = c
		return bcgo.StopIterationError{}
	}); err != nil {
		switch err.(type) {
		case bcgo.StopIterationError:
			// Do nothing
			break
		default:
			return nil, err
		}
	}
	return charge, nil
}

func GetRegistrationAsync(registrations bcgo.Channel, cache bcgo.Cache, merchantAlias string, merchantKey *rsa.PrivateKey, customerAlias string, customerKey *rsa.PrivateKey, callback func(*Registration) error) error {
	if err := bcgo.LoadHead(registrations, cache, nil); err != nil {
		log.Println(err)
	}
	head := registrations.GetHead()
	cb := func(entry *bcgo.BlockEntry, key, data []byte) error {
		// Unmarshal as Registration
		registration := &Registration{}
		err := proto.Unmarshal(data, registration)
		if err != nil {
			return err
		} else if (merchantAlias == "" || registration.MerchantAlias == merchantAlias) && (customerAlias == "" || registration.CustomerAlias == customerAlias) {
			return callback(registration)
		}
		return nil
	}
	// Read as merchant
	if merchantAlias != "" && merchantKey != nil {
		return bcgo.Read(head, nil, cache, merchantAlias, merchantKey, nil, cb)
	}
	// Read as customer
	return bcgo.Read(head, nil, cache, customerAlias, customerKey, nil, cb)
}

func GetRegistrationSync(registrations bcgo.Channel, cache bcgo.Cache, merchantAlias string, merchantKey *rsa.PrivateKey, customerAlias string, customerKey *rsa.PrivateKey) (*Registration, error) {
	var registration *Registration
	if err := GetRegistrationAsync(registrations, cache, merchantAlias, merchantKey, customerAlias, customerKey, func(r *Registration) error {
		registration = r
		return bcgo.StopIterationError{}
	}); err != nil {
		switch err.(type) {
		case bcgo.StopIterationError:
			// Do nothing
			break
		default:
			return nil, err
		}
	}
	return registration, nil
}

func GetSubscriptionAsync(subscriptions bcgo.Channel, cache bcgo.Cache, merchantAlias string, merchantKey *rsa.PrivateKey, customerAlias string, customerKey *rsa.PrivateKey, productId, planId string, callback func(*Subscription) error) error {
	if err := bcgo.LoadHead(subscriptions, cache, nil); err != nil {
		log.Println(err)
	}
	head := subscriptions.GetHead()
	cb := func(entry *bcgo.BlockEntry, key, data []byte) error {
		// Unmarshal as Subscription
		subscription := &Subscription{}
		err := proto.Unmarshal(data, subscription)
		if err != nil {
			return err
		} else if (merchantAlias == "" || subscription.MerchantAlias == merchantAlias) && (customerAlias == "" || subscription.CustomerAlias == customerAlias) && (productId == "" || subscription.ProductId == productId) && (planId == "" || subscription.PlanId == planId) {
			return callback(subscription)
		}
		return nil
	}
	// Read as merchant
	if merchantAlias != "" && merchantKey != nil {
		return bcgo.Read(head, nil, cache, merchantAlias, merchantKey, nil, cb)
	}
	// Read as customer
	return bcgo.Read(head, nil, cache, customerAlias, customerKey, nil, cb)
}

func GetSubscriptionSync(subscriptions bcgo.Channel, cache bcgo.Cache, merchantAlias string, merchantKey *rsa.PrivateKey, customerAlias string, customerKey *rsa.PrivateKey, productId, planId string) (*Subscription, error) {
	var subscription *Subscription
	if err := GetSubscriptionAsync(subscriptions, cache, merchantAlias, merchantKey, customerAlias, customerKey, productId, planId, func(s *Subscription) error {
		subscription = s
		return bcgo.StopIterationError{}
	}); err != nil {
		switch err.(type) {
		case bcgo.StopIterationError:
			// Do nothing
			break
		default:
			return nil, err
		}
	}
	return subscription, nil
}

func GetUsageRecordAsync(usages bcgo.Channel, cache bcgo.Cache, merchantAlias string, merchantKey *rsa.PrivateKey, customerAlias string, customerKey *rsa.PrivateKey, callback func(*UsageRecord) error) error {
	if err := bcgo.LoadHead(usages, cache, nil); err != nil {
		log.Println(err)
	}
	head := usages.GetHead()
	cb := func(entry *bcgo.BlockEntry, key, data []byte) error {
		// Unmarshal as UsageRecord
		usage := &UsageRecord{}
		err := proto.Unmarshal(data, usage)
		if err != nil {
			return err
		} else if (merchantAlias == "" || usage.MerchantAlias == merchantAlias) && (customerAlias == "" || usage.CustomerAlias == customerAlias) {
			return callback(usage)
		}
		return nil
	}
	// Read as merchant
	if merchantAlias != "" && merchantKey != nil {
		return bcgo.Read(head, nil, cache, merchantAlias, merchantKey, nil, cb)
	}
	// Read as customer
	return bcgo.Read(head, nil, cache, customerAlias, customerKey, nil, cb)
}

func GetUsageRecordSync(usages bcgo.Channel, cache bcgo.Cache, merchantAlias string, merchantKey *rsa.PrivateKey, customerAlias string, customerKey *rsa.PrivateKey) (*UsageRecord, error) {
	var usage *UsageRecord
	if err := GetUsageRecordAsync(usages, cache, merchantAlias, merchantKey, customerAlias, customerKey, func(u *UsageRecord) error {
		usage = u
		return bcgo.StopIterationError{}
	}); err != nil {
		switch err.(type) {
		case bcgo.StopIterationError:
			// Do nothing
			break
		default:
			return nil, err
		}
	}
	return usage, nil
}

func ConstructEvent(data []byte, signature string) (stripe.Event, error) {
	secretKey := os.Getenv("STRIPE_WEB_HOOK_SECRET_KEY")
	return webhook.ConstructEvent(data, signature, secretKey)
}
