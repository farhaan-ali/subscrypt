package database

import (
	"fmt"

	"github.com/Catzkorn/subscrypt/internal/subscription"
	"github.com/Catzkorn/subscrypt/internal/userprofile"
)

// NewInMemorySubscriptionStore returns a instance of InMemorySubscriptionStore
func NewInMemorySubscriptionStore() *InMemorySubscriptionStore {
	return &InMemorySubscriptionStore{[]subscription.Subscription{}, &userprofile.Userprofile{}}
}

// InMemorySubscriptionStore stores information about individual subscriptions
type InMemorySubscriptionStore struct {
	subscriptions []subscription.Subscription
	userProfile   *userprofile.Userprofile
}

// GetSubscriptions is a method that returns all subscriptions
func (i *InMemorySubscriptionStore) GetSubscriptions() ([]subscription.Subscription, error) {
	return i.subscriptions, nil
}

// GetSubscription retrieves a single subscription that has the given ID from the InMemoryDataStore
// If no subscription is found with the given ID, it returns a nil pointer
func (i *InMemorySubscriptionStore) GetSubscription(ID int) (*subscription.Subscription, error) {
	index := i.findSubscriptionIndex(ID)
	if index == -1 {
		return nil, nil
	}
	return &i.subscriptions[index], nil
}

// RecordSubscription is a method that stores a subscription into the store
func (i *InMemorySubscriptionStore) RecordSubscription(subscription subscription.Subscription) (*subscription.Subscription, error) {
	i.subscriptions = append(i.subscriptions, subscription)
	return &subscription, nil
}

// DeleteSubscription deletes a subscription from the data store with the given ID
func (i *InMemorySubscriptionStore) DeleteSubscription(subscriptionID int) error {

	index := i.findSubscriptionIndex(subscriptionID)
	lastIndex := len(i.subscriptions) - 1

	if index == -1 {
		return fmt.Errorf("failed to delete subscription with ID %v", subscriptionID)
	}
	i.subscriptions[lastIndex], i.subscriptions[index] = i.subscriptions[index], i.subscriptions[lastIndex]
	i.subscriptions = i.subscriptions[:lastIndex]
	return nil
}

// RecordUserDetails stores the users name and email
func (i *InMemorySubscriptionStore) RecordUserDetails(name string, email string) (*userprofile.Userprofile, error) {
	i.userProfile = &userprofile.Userprofile{
		Name:  name,
		Email: email,
	}

	return i.userProfile, nil
}

// GetUserDetails returns the users name and email
func (i *InMemorySubscriptionStore) GetUserDetails() (*userprofile.Userprofile, error) {
	return i.userProfile, nil
}

// FindSubscriptionIndex finds the index of a given subscription ID, from the InMemoryDataStore's subscriptions
// If no corresponding subscription is found then it returns -1
func (i *InMemorySubscriptionStore) findSubscriptionIndex(ID int) (index int) {
	for index, value := range i.subscriptions {
		if value.ID == ID {
			return index
		}
	}
	return -1
}
