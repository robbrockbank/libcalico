package client

import (
	. "github.com/projectcalico/libcalico/lib/backend/objects"
)

// HostEndpointInterface has methods to work with Service resources.
type HostEndpointInterface interface {
	List(hostname, name *string) ([]HostEndpoint, error)
	Get(hostname, name string) (*HostEndpoint, error)
	Create(h *HostEndpoint) error
	Update(h *HostEndpoint) error
	Delete(hostname, name string) error
}

// services implements ServicesNamespacer interface
type hostEndpoints struct {
	r *Client
}

// newServices returns a services
func newHostEndpoints(c *Client) *hostEndpoints {
	return &hostEndpoints{c}
}

// List takes a Metadata, and returns the list of host endpoints that match that Metadata
// (wildcarding missing fields)
func (h *hostEndpoints) List(hlo *HostEndpointListOptions) ([]HostEndpoint, error) {
	return nil, nil
}

// Get returns information about a particular host endpoint.
func (h *hostEndpoints) Get(hostname, name string) (*HostEndpoint, error) {
	return nil, nil
}

// Create creates a new host endpoint.
func (h *hostEndpoints) Create(hep *HostEndpoint) error {
	return nil
}

// Update updates an existing host endpoint.
func (h *hostEndpoints) Update(hep *HostEndpoint) error {
	return nil

}

// Delete deletes an existing host endpoint.
func (h *hostEndpoints) Delete(hostname, name string) error {
	return nil

}
