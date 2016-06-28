package client

import (
	. "github.com/projectcalico/libcalico/lib/backend/objects"
)

// HostEndpointInterface has methods to work with HostEndoint objects.
type HostEndpointInterface interface {
	List(hlo *HostEndpointListOptions) ([]HostEndpoint, error)
	Get(k *HostEndpointKey) (*HostEndpoint, error)
	Create(h *HostEndpoint) error
	Update(h *HostEndpoint) error
	Delete(k *HostEndpointKey) error
}

// hostEndpoints implements HostEndpointInterface
type hostEndpoints struct {
	r *Client
}

// newHostEndpoints returns a hostEndpoints
func newHostEndpoints(c *Client) *hostEndpoints {
	return &hostEndpoints{c}
}

// List takes a Metadata, and returns the list of host endpoints that match that Metadata
// (wildcarding missing fields)
func (h *hostEndpoints) List(hlo *HostEndpointListOptions) ([]HostEndpoint, error) {

	return nil, nil
}

// Get returns information about a particular host endpoint.
func (h *hostEndpoints) Get(k *HostEndpointKey) (*HostEndpoint, error) {
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
func (h *hostEndpoints) Delete(k *HostEndpointKey) error {
	return nil
}
