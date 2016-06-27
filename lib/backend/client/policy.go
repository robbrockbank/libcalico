package client

import (
	. "github.com/projectcalico/libcalico/lib/backend/objects"
)

// PolicyInterface has methods to work with Service resources.
type PolicyInterface interface {
	List(hostname, name *string) ([]Policy, error)
	Get(hostname, name string) (*Policy, error)
	Create(hostname, name string, data *Policy) error
	Update(hostname, name string, data *Policy) error
	Delete(hostname, name string) error
}

// services implements ServicesNamespacer interface
type policies struct {
	r  *Client
}

// newServices returns a services
func newPolicies(c *Client) *policies {
	return &policies{c}
}

// List takes a Metadata, and returns the list of hot endpoints that match that Metadata
// (wildcarding mising fields)
func (h *policies) List(hostname, name *string) ([]Policy, error) {
	return nil, nil
}

// Get returns information about a particular policy.
func (h *policies) Get(hostname, name string) (*Policy, error) {
	return nil, nil
}

// Create creates a new policy.
func (h *policies) Create(hostname, name string, data *Policy) error {
	return nil

}

// Update updates an existing policy.
func (h *policies) Update(hostname, name string, data *Policy) error {
	return nil

}

// Delete deletes an existing policy.
func (h *policies) Delete(hostname, name string) error {
	return nil

}