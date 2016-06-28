package client

import (
	. "github.com/projectcalico/libcalico/lib/api"
	backend "github.com/projectcalico/libcalico/lib/backend/objects"
)

// TierInterface has methods to work with Tier resources.
type TierInterface interface {
	List(metadata *TierMetadata) (*TierList, error)
	Get(metadata *TierMetadata) (*Tier, error)
	Create(hep *Tier) (*Tier, error)
	Update(hep *Tier) (*Tier, error)
	Delete(metadata *TierMetadata) error
}

// tiers implements TierInterface
type tiers struct {
	c *Client
}

// newTiers returns a tiers
func newTiers(c *Client) *tiers {
	return &tiers{c}
}

// List takes a Metadata, and returns the list of tiers that match that Metadata
// (wildcarding missing fields)
func (t *tiers) List(metadata *TierMetadata) (*TierList, error) {
	btlo := backend.TierListOptions{
		Name: metadata.Name,
	}
	if bts, err := t.c.backend.Tiers().List(btlo); err != nil {
		return nil, err
	} else {
		atl := NewTierList()
		atl.Items = bts
		return &atl, nil
	}
}

// Get returns information about a particular tier.
func (t *tiers) Get(metadata *TierMetadata) (*Tier, error) {
	if bt, err := t.c.backend.Tiers().Get(metadata.Name); err != nil {
		return nil, err
	} else {
		return tierBackendToAPI(bt), nil
	}
}

// Create creates a new tier.
func (t *tiers) Create(ap *Tier) (*Tier, error) {
	if bt, err := t.c.Tiers().Create(tierAPIToBackend(ap)); err != nil {
		return nil, err
	} else {
		return tierBackendToAPI(bt), nil
	}
}

// Update updates an existing tier.
func (t *tiers) Update(at *Tier) (*Tier, error) {
	if bt, err := t.c.Tiers().Update(tierAPIToBackend(at)); err != nil {
		return nil, err
	} else {
		return tierBackendToAPI(bt), nil
	}
}

// Delete deletes an existing tier.
func (t *tiers) Delete(metadata *TierMetadata) error {
	return t.c.backend.Tiers().Delete(metadata.Name)
}

// Convert an API Tier structure to a Backend Tier structure
func tierAPIToBackend(at *Tier) *backend.Tier {
	bt := backend.Tier{
		Name: at.Metadata.Name,

		Order: at.Spec.Order,
	}

	return &bt
}

// Convert a Backend Tier structure to an API Tier structure
func tierBackendToAPI(bt *backend.Tier) *Tier {
	at := NewTier()
	at.Metadata.Name = bt.Name

	at.Spec.Order = bt.Order

	return &at
}
