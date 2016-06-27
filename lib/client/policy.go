package client

import (
	. "github.com/projectcalico/libcalico/lib/api"
)

// PolicyInterface has methods to work with Policy resources.
type PolicyInterface interface {
	List(metadata PolicyMetadata) (*PolicyList, error)
	Get(metadata PolicyMetadata) (*Policy, error)
	Create(hep *Policy) (*Policy, error)
	Update(hep *Policy) (*Policy, error)
	Delete(metadata PolicyMetadata) error
}

// services implements ServicesNamespacer interface
type policies struct {
	r  *CalicoClient
}

// newServices returns a services
func newPolicies(c *CalicoClient) *policies {
	return &policies{c}
}

// List takes a Metadata, and returns the list of hot endpoints that match that Metdata
// (wildcarding mising fields)
func (p *policies) List(metadata PolicyMetadata) (*PolicyList, error) {
	return nil, nil
}

// Get returns information about a particular policy.
func (p *policies) Get(metadata PolicyMetadata) (*Policy, error) {
	return nil, nil
}

// Create creates a new policy.
func (p *policies) Create(hep *Policy) (*Policy, error) {
	return nil, nil
	
}

// Update updates an existing policy.
func (p *policies) Update(hep *Policy) (*Policy, error) {
	return nil, nil
	
}

// Delete deletes an existing policy.
func (p *policies) Delete(metadata PolicyMetadata) error {
	return nil
	
}