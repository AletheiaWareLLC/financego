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
	CUSTOMER     = "Customer"
	SUBSCRIPTION = "Subscription"
	USAGE_RECORD = "UsageRecord"
)

func OpenAndLoadChargeChannel(cache bcgo.Cache, network bcgo.Network) *bcgo.PoWChannel {
	return bcgo.OpenAndLoadPoWChannel(CHARGE, bcgo.THRESHOLD_STANDARD, cache, network)
}

func OpenAndLoadCustomerChannel(cache bcgo.Cache, network bcgo.Network) *bcgo.PoWChannel {
	return bcgo.OpenAndLoadPoWChannel(CUSTOMER, bcgo.THRESHOLD_STANDARD, cache, network)
}

func OpenAndLoadSubscriptionChannel(cache bcgo.Cache, network bcgo.Network) *bcgo.PoWChannel {
	return bcgo.OpenAndLoadPoWChannel(SUBSCRIPTION, bcgo.THRESHOLD_STANDARD, cache, network)
}

func OpenAndLoadUsageRecordChannel(cache bcgo.Cache, network bcgo.Network) *bcgo.PoWChannel {
	return bcgo.OpenAndLoadPoWChannel(USAGE_RECORD, bcgo.THRESHOLD_STANDARD, cache, network)
}

func NewCharge(alias, paymentId string, amount int64, description string) (*stripe.Charge, *Charge, error) {
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

func NewCustomer(alias, email, paymentId, description string) (*stripe.Customer, *Customer, error) {
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

func NewSubscription(alias, customerId, paymentId, productId, planId string) (*stripe.Subscription, *Subscription, error) {
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

func GetChargeAsync(charges bcgo.Channel, cache bcgo.Cache, alias string, key *rsa.PrivateKey, chargeAlias string, callback func(*Charge) error) error {
	if err := bcgo.LoadHead(charges, cache, nil); err != nil {
		log.Println(err)
	}
	return bcgo.Read(charges.GetHead(), nil, cache, alias, key, nil, func(entry *bcgo.BlockEntry, key, data []byte) error {
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

func GetChargeSync(charges bcgo.Channel, cache bcgo.Cache, alias string, key *rsa.PrivateKey, chargeAlias string) (*Charge, error) {
	var charge *Charge
	if err := GetChargeAsync(charges, cache, alias, key, chargeAlias, func(c *Charge) error {
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

func GetCustomerAsync(customers bcgo.Channel, cache bcgo.Cache, alias string, key *rsa.PrivateKey, customerAlias string, callback func(*Customer) error) error {
	if err := bcgo.LoadHead(customers, cache, nil); err != nil {
		log.Println(err)
	}
	return bcgo.Read(customers.GetHead(), nil, cache, alias, key, nil, func(entry *bcgo.BlockEntry, key, data []byte) error {
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

func GetCustomerSync(customers bcgo.Channel, cache bcgo.Cache, alias string, key *rsa.PrivateKey, customerAlias string) (*Customer, error) {
	var customer *Customer
	if err := GetCustomerAsync(customers, cache, alias, key, customerAlias, func(c *Customer) error {
		customer = c
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
	return customer, nil
}

func GetSubscriptionAsync(subscriptions bcgo.Channel, cache bcgo.Cache, alias string, key *rsa.PrivateKey, subscriptionAlias, productId, planId string, callback func(*Subscription) error) error {
	if err := bcgo.LoadHead(subscriptions, cache, nil); err != nil {
		log.Println(err)
	}
	return bcgo.Read(subscriptions.GetHead(), nil, cache, alias, key, nil, func(entry *bcgo.BlockEntry, key, data []byte) error {
		// Unmarshal as Subscription
		subscription := &Subscription{}
		err := proto.Unmarshal(data, subscription)
		if err != nil {
			return err
		} else if subscription.Alias == subscriptionAlias && (productId == "" || subscription.ProductId == productId) && (planId == "" || subscription.PlanId == planId) {
			return callback(subscription)
		}
		return nil
	})
}

func GetSubscriptionSync(subscriptions bcgo.Channel, cache bcgo.Cache, alias string, key *rsa.PrivateKey, subscriptionAlias, productId, planId string) (*Subscription, error) {
	var subscription *Subscription
	if err := GetSubscriptionAsync(subscriptions, cache, alias, key, subscriptionAlias, productId, planId, func(s *Subscription) error {
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

func GetUsageRecordAsync(usages bcgo.Channel, cache bcgo.Cache, alias string, key *rsa.PrivateKey, usageAlias string, callback func(*UsageRecord) error) error {
	if err := bcgo.LoadHead(usages, cache, nil); err != nil {
		log.Println(err)
	}
	return bcgo.Read(usages.GetHead(), nil, cache, alias, key, nil, func(entry *bcgo.BlockEntry, key, data []byte) error {
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

func GetUsageRecordSync(usages bcgo.Channel, cache bcgo.Cache, alias string, key *rsa.PrivateKey, usageAlias string) (*UsageRecord, error) {
	var usage *UsageRecord
	if err := GetUsageRecordAsync(usages, cache, alias, key, usageAlias, func(u *UsageRecord) error {
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
