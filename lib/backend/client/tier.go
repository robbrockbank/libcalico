package client

import (
	. "github.com/projectcalico/libcalico/lib/backend/objects"
)


// TierInterface has methods to work with Service resources.
type TierInterface interface {
	List(hostname, name *string) ([]Tier, error)
	Get(hostname, name string) (*Tier, error)
	Create(hostname, name string, data *Tier) error
	Update(hostname, name string, data *Tier) error
	Delete(hostname, name string) error
}

// services implements ServicesNamespacer interface
type tiers struct {
	r  *Client
}

// newServices returns a services
func newTiers(c *Client) *tiers {
	return &tiers{c}
}

// List takes a Metadata, and returns the list of hot endpoints that match that Metadata
// (wildcarding mising fields)
func (h *tiers) List(hostname, name *string) ([]Tier, error) {
	return nil, nil
}

// Get returns information about a particular tier.
func (h *tiers) Get(hostname, name string) (*Tier, error) {
	return nil, nil
}

// Create creates a new tier.
func (h *tiers) Create(hostname, name string, data *Tier) error {
	return nil

}

// Update updates an existing tier.
func (h *tiers) Update(hostname, name string, data *Tier) error {
	return nil

}

// Delete deletes an existing tier.
func (h *tiers) Delete(hostname, name string) error {
	return nil

}