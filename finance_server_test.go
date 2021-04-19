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
	"aletheiaware.com/aliasgo"
	"aletheiaware.com/bcgo"
	"aletheiaware.com/bcgo/account"
	"aletheiaware.com/bcgo/cache"
	"aletheiaware.com/bcgo/channel"
	"aletheiaware.com/bcgo/node"
	"aletheiaware.com/financego"
	"aletheiaware.com/testinggo"
	"os"
	"testing"
)

func TestRegister(t *testing.T) {
	cache := cache.NewMemory(10)
	aliceAlias := "Alice"
	aliceAccount, err := account.GenerateRSA(aliceAlias)
	testinggo.AssertNoError(t, err)
	aliceNode := node.New(aliceAccount, cache, nil)

	bobAlias := "Bob"
	bobAccount, err := account.GenerateRSA(bobAlias)
	testinggo.AssertNoError(t, err)

	expectedCustomerID := "cust1234"
	expectedPaymentID := "card1234"
	provider := &mockProvider{
		registration: &financego.Registration{
			MerchantAlias: aliceAlias,
			CustomerAlias: bobAlias,
			CustomerId:    expectedCustomerID,
			PaymentId:     expectedPaymentID,
		},
	}
	aliases := aliasgo.OpenAliasChannel()
	registrations := channel.New("Registration")
	listener := &bcgo.PrintingMiningListener{Output: os.Stdout}
	recordA, err := aliasgo.CreateSignedAliasRecord(aliceAccount)
	testinggo.AssertNoError(t, err)
	recordB, err := aliasgo.CreateSignedAliasRecord(bobAccount)
	testinggo.AssertNoError(t, err)
	_, err = bcgo.WriteRecord(aliasgo.ALIAS, cache, recordA)
	testinggo.AssertNoError(t, err)
	_, err = bcgo.WriteRecord(aliasgo.ALIAS, cache, recordB)
	testinggo.AssertNoError(t, err)
	_, _, err = bcgo.Mine(aliceNode, aliases, aliasgo.ALIAS_THRESHOLD, listener)
	testinggo.AssertNoError(t, err)
	handler := financego.Register(aliceNode, provider, aliases, registrations, 0, listener)
	customerID, registrationReference, err := handler("Bob", "b@o.b", expectedPaymentID)
	testinggo.AssertNoError(t, err)
	if customerID != expectedCustomerID {
		t.Fatalf("Wrong customer ID: expected '%s' , instead got '%s'", customerID, expectedCustomerID)
	}
	if registrationReference == nil {
		t.Fatalf("Registration Reference is nil")
	}
	// Ensure Merchant can read Registration
	merchantRegistration, err := financego.RegistrationSync(registrations, cache, nil, aliceAccount, aliceAlias, bobAlias)
	testinggo.AssertNoError(t, err)
	if merchantRegistration == nil {
		t.Fatalf("Merchant Registration is nil")
	}
	// Ensure Customer can read Registration
	customerRegistration, err := financego.RegistrationSync(registrations, cache, nil, bobAccount, aliceAlias, bobAlias)
	testinggo.AssertNoError(t, err)
	if customerRegistration == nil {
		t.Fatalf("Customer Registration is nil")
	}
}

func TestSubscribe(t *testing.T) {
	cache := cache.NewMemory(10)
	aliceAlias := "Alice"
	aliceAccount, err := account.GenerateRSA(aliceAlias)
	testinggo.AssertNoError(t, err)
	aliceNode := node.New(aliceAccount, cache, nil)

	bobAlias := "Bob"
	bobAccount, err := account.GenerateRSA(bobAlias)
	testinggo.AssertNoError(t, err)

	expectedCustomerID := "cust1234"
	expectedPaymentID := "card1234"
	expectedProductId := "product1234"
	expectedPlanId := "plan1234"
	expectedSubscriptionId := "sub1234"
	expectedSubscriptionItemId := "subitem1234"
	provider := &mockProvider{
		subscription: &financego.Subscription{
			MerchantAlias:      aliceAlias,
			CustomerAlias:      bobAlias,
			CustomerId:         expectedCustomerID,
			PaymentId:          expectedPaymentID,
			ProductId:          expectedProductId,
			PlanId:             expectedPlanId,
			SubscriptionId:     expectedSubscriptionId,
			SubscriptionItemId: expectedSubscriptionItemId,
		},
	}
	aliases := aliasgo.OpenAliasChannel()
	subscriptions := channel.New("Subscription")
	listener := &bcgo.PrintingMiningListener{Output: os.Stdout}
	recordA, err := aliasgo.CreateSignedAliasRecord(aliceAccount)
	testinggo.AssertNoError(t, err)
	recordB, err := aliasgo.CreateSignedAliasRecord(bobAccount)
	testinggo.AssertNoError(t, err)
	_, err = bcgo.WriteRecord(aliasgo.ALIAS, cache, recordA)
	testinggo.AssertNoError(t, err)
	_, err = bcgo.WriteRecord(aliasgo.ALIAS, cache, recordB)
	testinggo.AssertNoError(t, err)
	_, _, err = bcgo.Mine(aliceNode, aliases, aliasgo.ALIAS_THRESHOLD, listener)
	testinggo.AssertNoError(t, err)
	handler := financego.Subscribe(aliceNode, provider, aliases, subscriptions, 0, listener, expectedProductId, expectedPlanId)
	subscriptionItemID, subscriptionReference, err := handler("Bob", expectedCustomerID)
	testinggo.AssertNoError(t, err)
	if subscriptionItemID != expectedSubscriptionItemId {
		t.Fatalf("Wrong subscription item ID: expected '%s' , instead got '%s'", subscriptionItemID, expectedSubscriptionItemId)
	}
	if subscriptionReference == nil {
		t.Fatalf("Subscription Reference is nil")
	}
	// Ensure Merchant can read Subscription
	merchantSubscription, err := financego.SubscriptionSync(subscriptions, cache, nil, aliceAccount, aliceAlias, bobAlias, expectedProductId, expectedPlanId)
	testinggo.AssertNoError(t, err)
	if merchantSubscription == nil {
		t.Fatalf("Merchant Subscription is nil")
	}
	// Ensure Customer can read Subscription
	customerSubscription, err := financego.SubscriptionSync(subscriptions, cache, nil, bobAccount, aliceAlias, bobAlias, expectedProductId, expectedPlanId)
	testinggo.AssertNoError(t, err)
	if customerSubscription == nil {
		t.Fatalf("Customer Subscription is nil")
	}
}
