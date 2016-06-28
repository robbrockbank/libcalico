package client

import (
	. "github.com/projectcalico/libcalico/lib/backend/objects"
)

// PolicyInterface has methods to work with Policy objects.
type PolicyInterface interface {
	List(bplo *PolicyListOptions) ([]Policy, error)
	Get(name string) (*Policy, error)
	Create(bp *Policy) error
	Update(bp *Policy) error
	Delete(name string) error
}

// policies implements PolicyInterface
type policies struct {
	r *Client
}

// newPolicies returns a policies
func newPolicies(c *Client) *policies {
	return &policies{c}
}

// List takes a Metadata, and returns the list of policies that match that Metadata
// (wildcarding missing fields)
func (p *policies) List(bplo *PolicyListOptions) ([]Policy, error) {
	return nil, nil
}

// Get returns information about a particular policy.
func (p *policies) Get(name string) (*Policy, error) {
	return nil, nil
}

// Create creates a new policy.
func (p *policies) Create(bp *Policy) error {
	return nil

}

// Update updates an existing policy.
func (p *policies) Update(bp *Policy) error {
	return nil

}

// Delete deletes an existing policy.
func (p *policies) Delete(name string) error {
	return nil

}
