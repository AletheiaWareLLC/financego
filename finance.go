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
	"errors"
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
	"time"
)

const (
	CHARGE       = "Charge"
	CUSTOMER     = "Customer"
	SUBSCRIPTION = "Subscription"
	USAGE_RECORD = "UsageRecord"
)

func OpenChargeChannel() (*bcgo.Channel, error) {
	return bcgo.OpenChannel(CHARGE)
}

func OpenCustomerChannel() (*bcgo.Channel, error) {
	return bcgo.OpenChannel(CUSTOMER)
}

func OpenSubscriptionChannel() (*bcgo.Channel, error) {
	return bcgo.OpenChannel(SUBSCRIPTION)
}

func OpenUsageRecordChannel() (*bcgo.Channel, error) {
	return bcgo.OpenChannel(USAGE_RECORD)
}

func NewCharge(alias string, paymentId string, amount int64, description string) (*stripe.Charge, *Charge, error) {
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
		Alias:     alias,
		Processor: PaymentProcessor_STRIPE,
		PaymentId: paymentId,
		ChargeId:  ch.ID,
	}
	log.Println("Charge", charge)
	return ch, charge, nil
}

func NewCustomerCharge(customer *Customer, amount int64, description string) (*stripe.Charge, *Charge, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	chargeParams := &stripe.ChargeParams{
		Amount:      stripe.Int64(amount),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Customer:    stripe.String(customer.CustomerId),
		Description: stripe.String(description),
	}
	ch, err := charge.New(chargeParams)
	if err != nil {
		return nil, nil, err
	}

	charge := &Charge{
		Alias:      customer.Alias,
		Processor:  PaymentProcessor_STRIPE,
		CustomerId: customer.CustomerId,
		ChargeId:   ch.ID,
	}
	log.Println("Charge", charge)
	return ch, charge, nil
}

func GetChargeAsync(charges *bcgo.Channel, alias string, key *rsa.PrivateKey, chargeAlias string, callback func(*Charge) error) error {
	return charges.Read(alias, key, nil, func(entry *bcgo.BlockEntry, key, data []byte) error {
		// Unmarshal as Charge
		charge := &Charge{}
		err := proto.Unmarshal(data, charge)
		if err != nil {
			return err
		} else if charge.Alias == chargeAlias {
			return callback(charge)
		}
		return nil
	})
}

func GetChargeSync(charges *bcgo.Channel, alias string, key *rsa.PrivateKey, chargeAlias string) (*Charge, error) {
	// Load Charge Information
	ch := make(chan *Charge, 1)
	err := GetChargeAsync(charges, alias, key, chargeAlias, func(charge *Charge) error {
		ch <- charge
		return nil
	})
	if err != nil {
		return nil, err
	}
	select {
	case charge := <-ch:
		return charge, nil
	case <-time.After(1 * time.Minute):
		return nil, errors.New("Timeout getting charge - 1 minute")
	}
}

func NewCustomer(alias string, email string, paymentId string, description string) (*stripe.Customer, *Customer, error) {
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

	customer := &Customer{
		Alias:      alias,
		Processor:  PaymentProcessor_STRIPE,
		CustomerId: c.ID,
		PaymentId:  paymentId,
	}
	log.Println("Customer", customer)
	return c, customer, nil
}

func GetCustomerAsync(customers *bcgo.Channel, alias string, key *rsa.PrivateKey, customerAlias string, callback func(*Customer) error) error {
	return customers.Read(alias, key, nil, func(entry *bcgo.BlockEntry, key, data []byte) error {
		// Unmarshal as Customer
		customer := &Customer{}
		err := proto.Unmarshal(data, customer)
		if err != nil {
			return err
		} else if customer.Alias == customerAlias {
			return callback(customer)
		}
		return nil
	})
}

func GetCustomerSync(customers *bcgo.Channel, alias string, key *rsa.PrivateKey, customerAlias string) (*Customer, error) {
	// Load Customer Information
	ch := make(chan *Customer, 1)
	err := GetCustomerAsync(customers, alias, key, customerAlias, func(customer *Customer) error {
		ch <- customer
		return nil
	})
	if err != nil {
		return nil, err
	}
	select {
	case customer := <-ch:
		return customer, nil
	case <-time.After(1 * time.Minute):
		return nil, errors.New("Timeout getting customer - 1 minute")
	}
}

func NewSubscription(alias string, customerId string, paymentId string, productId string, planId string) (*stripe.Subscription, *Subscription, error) {
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
		Alias:              alias,
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

func GetSubscriptionAsync(subscriptions *bcgo.Channel, alias string, key *rsa.PrivateKey, subscriptionAlias string, callback func(*Subscription) error) error {
	return subscriptions.Read(alias, key, nil, func(entry *bcgo.BlockEntry, key, data []byte) error {
		// Unmarshal as Subscription
		subscription := &Subscription{}
		err := proto.Unmarshal(data, subscription)
		if err != nil {
			return err
		} else if subscription.Alias == subscriptionAlias {
			return callback(subscription)
		}
		return nil
	})
}

func GetSubscriptionSync(subscriptions *bcgo.Channel, alias string, key *rsa.PrivateKey, subscriptionAlias string) (*Subscription, error) {
	// Load Subscription Information
	ch := make(chan *Subscription, 1)
	err := GetSubscriptionAsync(subscriptions, alias, key, subscriptionAlias, func(subscription *Subscription) error {
		ch <- subscription
		return nil
	})
	if err != nil {
		return nil, err
	}
	select {
	case subscription := <-ch:
		return subscription, nil
	case <-time.After(1 * time.Minute):
		return nil, errors.New("Timeout getting subscription - 1 minute")
	}
}

func NewUsageRecord(alias string, subscription string, timestamp int64, size int64) (*stripe.UsageRecord, *UsageRecord, error) {
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
		Alias:          alias,
		Processor:      PaymentProcessor_STRIPE,
		SubscriptionId: subscription,
		UsageRecordId:  ur.ID,
	}
	log.Println("UsageRecord", usage)
	return ur, usage, nil
}

func GetUsageRecordAsync(usages *bcgo.Channel, alias string, key *rsa.PrivateKey, usageAlias string, callback func(*UsageRecord) error) error {
	return usages.Read(alias, key, nil, func(entry *bcgo.BlockEntry, key, data []byte) error {
		// Unmarshal as UsageRecord
		usage := &UsageRecord{}
		err := proto.Unmarshal(data, usage)
		if err != nil {
			return err
		} else if usage.Alias == usageAlias {
			return callback(usage)
		}
		return nil
	})
}

func GetUsageRecordSync(usages *bcgo.Channel, alias string, key *rsa.PrivateKey, usageAlias string) (*UsageRecord, error) {
	// Load UsageRecord Information
	ch := make(chan *UsageRecord, 1)
	err := GetUsageRecordAsync(usages, alias, key, usageAlias, func(usage *UsageRecord) error {
		ch <- usage
		return nil
	})
	if err != nil {
		return nil, err
	}
	select {
	case usage := <-ch:
		return usage, nil
	case <-time.After(1 * time.Minute):
		return nil, errors.New("Timeout getting usage - 1 minute")
	}
}

func ConstructEvent(data []byte, signature string) (stripe.Event, error) {
	secretKey := os.Getenv("STRIPE_WEB_HOOK_SECRET_KEY")
	return webhook.ConstructEvent(data, signature, secretKey)
}
