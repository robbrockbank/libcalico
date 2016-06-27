package client

import (
	. "github.com/projectcalico/libcalico/lib/api"
)

// TierInterface has methods to work with Tier resources.
type TierInterface interface {
	List(metadata TierMetadata) (*TierList, error)
	Get(metadata TierMetadata) (*Tier, error)
	Create(hep *Tier) (*Tier, error)
	Update(hep *Tier) (*Tier, error)
	Delete(metadata TierMetadata) error
}

// services implements ServicesNamespacer interface
type tiers struct {
	r  *CalicoClient
}

// newServices returns a services
func newTiers(c *CalicoClient) *tiers {
	return &tiers{c}
}

// List takes a Metadata, and returns the list of hot endpoints that match that Metdata
// (wildcarding mising fields)
func (t *tiers) List(metadata TierMetadata) (*TierList, error) {
	return nil, nil
}

// Get returns information about a particular tier.
func (t *tiers) Get(metadata TierMetadata) (*Tier, error) {
	return nil, nil
}

// Create creates a new tier.
func (t *tiers) Create(hep *Tier) (*Tier, error) {
	return nil, nil
	
}

// Update updates an existing tier.
func (t *tiers) Update(hep *Tier) (*Tier, error) {
	return nil, nil
	
}

// Delete deletes an existing tier.
func (t *tiers) Delete(metadata TierMetadata) error {
	return nil
	
}