package client

import (
	. "github.com/projectcalico/libcalico/lib/backend/objects"
)

// TierInterface has methods to work with Tier objects.
type TierInterface interface {
	List(btlo *TierListOptions) ([]Tier, error)
	Get(k *TierKey) (*Tier, error)
	Create(bp *Tier) error
	Update(bp *Tier) error
	Delete(k *TierKey) error
}

// tiers implements TierInterface
type tiers struct {
	r *Client
}

// newTiers returns a tiers
func newTiers(c *Client) *tiers {
	return &tiers{c}
}

// List takes a Metadata, and returns the list of tiers that match that Metadata
// (wildcarding missing fields)
func (t *tiers) List(btlo *TierListOptions) ([]Tier, error) {
	return nil, nil
}

// Get returns information about a particular tier.
func (t *tiers) Get(k *TierKey) (*Tier, error) {
	return nil, nil
}

// Create creates a new tier.
func (t *tiers) Create(bp *Tier) error {
	return nil

}

// Update updates an existing tier.
func (t *tiers) Update(bp *Tier) error {
	return nil

}

// Delete deletes an existing tier.
func (t *tiers) Delete(k *TierKey) error {
	return nil

}
