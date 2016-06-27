package client

import (
	. "github.com/projectcalico/libcalico/lib/api"
)

// HostEndpointInterface has methods to work with HostEndpoint resources.
type HostEndpointInterface interface {
	List(metadata HostEndpointMetadata) (*HostEndpointList, error)
	Get(metadata HostEndpointMetadata) (*HostEndpoint, error)
	Create(hep *HostEndpoint) (*HostEndpoint, error)
	Update(hep *HostEndpoint) (*HostEndpoint, error)
	Delete(metadata HostEndpointMetadata) error
}

// services implements ServicesNamespacer interface
type hostEndpoints struct {
	r  *CalicoClient
}

// newServices returns a services
func newHostEndpoints(c *CalicoClient) *hostEndpoints {
	return &hostEndpoints{c}
}

// List takes a Metadata, and returns the list of hot endpoints that match that Metdata
// (wildcarding mising fields)
func (h *hostEndpoints) List(metadata HostEndpointMetadata) (*HostEndpointList, error) {
	return nil, nil
}

// Get returns information about a particular host endpoint.
func (h *hostEndpoints) Get(metadata HostEndpointMetadata) (*HostEndpoint, error) {
	return nil, nil
}

// Create creates a new host endpoint.
func (h *hostEndpoints) Create(hep *HostEndpoint) (*HostEndpoint, error) {
	return nil, nil

}

// Update updates an existing host endpoint.
func (h *hostEndpoints) Update(hep *HostEndpoint) (*HostEndpoint, error) {
	return nil, nil

}

// Delete deletes an existing host endpoint.
func (h *hostEndpoints) Delete(metadata HostEndpointMetadata) error {
	return nil

}