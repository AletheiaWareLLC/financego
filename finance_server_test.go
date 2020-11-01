package financego_test

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/AletheiaWareLLC/aliasgo"
	"github.com/AletheiaWareLLC/bcgo"
	"github.com/AletheiaWareLLC/financego"
	"github.com/AletheiaWareLLC/testinggo"
	"os"
	"testing"
)

func TestRegister(t *testing.T) {
	cache := bcgo.NewMemoryCache(10)
	keyA, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		t.Error("Could not generate key:", err)
	}
	nodeA := &bcgo.Node{
		Alias:    "Alice",
		Key:      keyA,
		Cache:    cache,
		Channels: make(map[string]*bcgo.Channel),
	}
	keyB, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		t.Error("Could not generate key:", err)
	}
	nodeB := &bcgo.Node{
		Alias:    "Bob",
		Key:      keyB,
		Cache:    cache,
		Channels: make(map[string]*bcgo.Channel),
	}
	expectedCustomerID := "cust1234"
	expectedPaymentID := "card1234"
	provider := &mockProvider{
		registration: &financego.Registration{
			MerchantAlias: nodeA.Alias,
			CustomerAlias: nodeB.Alias,
			CustomerId:    expectedCustomerID,
			PaymentId:     expectedPaymentID,
		},
	}
	aliases := aliasgo.OpenAliasChannel()
	registrations := bcgo.NewChannel("Registration")
	listener := &bcgo.PrintingMiningListener{Output: os.Stdout}
	recordA, err := aliasgo.CreateSignedAliasRecord(nodeA.Alias, nodeA.Key)
	testinggo.AssertNoError(t, err)
	recordB, err := aliasgo.CreateSignedAliasRecord(nodeB.Alias, nodeB.Key)
	testinggo.AssertNoError(t, err)
	_, err = bcgo.WriteRecord(aliasgo.ALIAS, cache, recordA)
	testinggo.AssertNoError(t, err)
	_, err = bcgo.WriteRecord(aliasgo.ALIAS, cache, recordB)
	testinggo.AssertNoError(t, err)
	_, _, err = nodeA.Mine(aliases, aliasgo.ALIAS_THRESHOLD, listener)
	testinggo.AssertNoError(t, err)
	handler := financego.Register(nodeA, provider, aliases, registrations, 0, listener)
	customerID, registrationReference, err := handler("Bob", "b@o.b", expectedPaymentID)
	testinggo.AssertNoError(t, err)
	if customerID != expectedCustomerID {
		t.Fatalf("Wrong customer ID: expected '%s' , instead got '%s'", customerID, expectedCustomerID)
	}
	if registrationReference == nil {
		t.Fatalf("Registration Reference is nil")
	}
	// Ensure Merchant can read Registration
	merchantRegistration, err := financego.GetRegistrationSync(registrations, cache, nil, nodeA.Alias, nodeA.Key, "", nil)
	testinggo.AssertNoError(t, err)
	if merchantRegistration == nil {
		t.Fatalf("Merchant Registration is nil")
	}
	// Ensure Customer can read Registration
	customerRegistration, err := financego.GetRegistrationSync(registrations, cache, nil, "", nil, nodeB.Alias, nodeB.Key)
	testinggo.AssertNoError(t, err)
	if customerRegistration == nil {
		t.Fatalf("Customer Registration is nil")
	}
}

func TestSubscribe(t *testing.T) {
	cache := bcgo.NewMemoryCache(10)
	keyA, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		t.Error("Could not generate key:", err)
	}
	nodeA := &bcgo.Node{
		Alias:    "Alice",
		Key:      keyA,
		Cache:    cache,
		Channels: make(map[string]*bcgo.Channel),
	}
	keyB, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		t.Error("Could not generate key:", err)
	}
	nodeB := &bcgo.Node{
		Alias:    "Bob",
		Key:      keyB,
		Cache:    cache,
		Channels: make(map[string]*bcgo.Channel),
	}
	expectedCustomerID := "cust1234"
	expectedPaymentID := "card1234"
	expectedProductId := "product1234"
	expectedPlanId := "plan1234"
	expectedSubscriptionId := "sub1234"
	expectedSubscriptionItemId := "subitem1234"
	provider := &mockProvider{
		subscription: &financego.Subscription{
			MerchantAlias:      nodeA.Alias,
			CustomerAlias:      nodeB.Alias,
			CustomerId:         expectedCustomerID,
			PaymentId:          expectedPaymentID,
			ProductId:          expectedProductId,
			PlanId:             expectedPlanId,
			SubscriptionId:     expectedSubscriptionId,
			SubscriptionItemId: expectedSubscriptionItemId,
		},
	}
	aliases := aliasgo.OpenAliasChannel()
	subscriptions := bcgo.NewChannel("Subscription")
	listener := &bcgo.PrintingMiningListener{Output: os.Stdout}
	recordA, err := aliasgo.CreateSignedAliasRecord(nodeA.Alias, nodeA.Key)
	testinggo.AssertNoError(t, err)
	recordB, err := aliasgo.CreateSignedAliasRecord(nodeB.Alias, nodeB.Key)
	testinggo.AssertNoError(t, err)
	_, err = bcgo.WriteRecord(aliasgo.ALIAS, cache, recordA)
	testinggo.AssertNoError(t, err)
	_, err = bcgo.WriteRecord(aliasgo.ALIAS, cache, recordB)
	testinggo.AssertNoError(t, err)
	_, _, err = nodeA.Mine(aliases, aliasgo.ALIAS_THRESHOLD, listener)
	testinggo.AssertNoError(t, err)
	handler := financego.Subscribe(nodeA, provider, aliases, subscriptions, 0, listener)
	subscriptionItemID, subscriptionReference, err := handler("Bob", expectedCustomerID, expectedProductId, expectedPlanId)
	testinggo.AssertNoError(t, err)
	if subscriptionItemID != expectedSubscriptionItemId {
		t.Fatalf("Wrong subscription item ID: expected '%s' , instead got '%s'", subscriptionItemID, expectedSubscriptionItemId)
	}
	if subscriptionReference == nil {
		t.Fatalf("Subscription Reference is nil")
	}
	// Ensure Merchant can read Subscription
	merchantSubscription, err := financego.GetSubscriptionSync(subscriptions, cache, nil, nodeA.Alias, nodeA.Key, "", nil, expectedProductId, expectedPlanId)
	testinggo.AssertNoError(t, err)
	if merchantSubscription == nil {
		t.Fatalf("Merchant Subscription is nil")
	}
	// Ensure Customer can read Subscription
	customerSubscription, err := financego.GetSubscriptionSync(subscriptions, cache, nil, "", nil, nodeB.Alias, nodeB.Key, expectedProductId, expectedPlanId)
	testinggo.AssertNoError(t, err)
	if customerSubscription == nil {
		t.Fatalf("Customer Subscription is nil")
	}
}
