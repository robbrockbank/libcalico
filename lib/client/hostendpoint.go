package client

import (
	"net"

	. "github.com/projectcalico/libcalico/lib/api"
	backend "github.com/projectcalico/libcalico/lib/backend/objects"
	. "github.com/projectcalico/libcalico/lib/common"
)

// HostEndpointInterface has methods to work with HostEndpoint resources.
type HostEndpointInterface interface {
	List(metadata *HostEndpointMetadata) (*HostEndpointList, error)
	Get(metadata *HostEndpointMetadata) (*HostEndpoint, error)
	Create(hep *HostEndpoint) (*HostEndpoint, error)
	Update(hep *HostEndpoint) (*HostEndpoint, error)
	Delete(metadata *HostEndpointMetadata) error
}

// hostEndpoints implements HostEndpointInterface
type hostEndpoints struct {
	c *Client
}

// newHostEndpoints returns a hostEndpoints
func newHostEndpoints(c *Client) *hostEndpoints {
	return &hostEndpoints{c}
}

// List takes a Metadata, and returns the list of host endpoints that match that Metadata
// (wildcarding missing fields)
func (h *hostEndpoints) List(metadata *HostEndpointMetadata) (*HostEndpointList, error) {
	bhlo := backend.HostEndpointListOptions{
		Hostname:   metadata.Hostname,
		EndpointID: metadata.Name,
	}
	if bhs, err := h.c.backend.HostEndpoints().List(bhlo); err != nil {
		return nil, err
	} else {
		ahl := NewHostEndpointList()
		ahl.Items = bhs
		return &ahl, nil
	}
}

// Get returns information about a particular host endpoint.
func (h *hostEndpoints) Get(metadata *HostEndpointMetadata) (*HostEndpoint, error) {
	if bh, err := h.c.backend.HostEndpoints().Get(metadata.Hostname, metadata.Name); err != nil {
		return nil, err
	} else {
		return hostEndpointBackendToAPI(bh), nil
	}
}

// Create creates a new host endpoint.
func (h *hostEndpoints) Create(ah *HostEndpoint) (*HostEndpoint, error) {
	if bh, err := h.c.HostEndpoints().Create(hostEndpointAPIToBackend(ah)); err != nil {
		return nil, err
	} else {
		return hostEndpointBackendToAPI(bh), nil
	}
}

// Update updates an existing host endpoint.
func (h *hostEndpoints) Update(ah *HostEndpoint) (*HostEndpoint, error) {
	if bh, err := h.c.HostEndpoints().Update(hostEndpointAPIToBackend(ah)); err != nil {
		return nil, err
	} else {
		return hostEndpointBackendToAPI(bh), nil
	}
}

// Delete deletes an existing host endpoint.
func (h *hostEndpoints) Delete(metadata *HostEndpointMetadata) error {
	return h.c.backend.HostEndpoints().Delete(metadata.Hostname, metadata.Name)
}

// Convert an API HostEndpoint structure to a Backend Tier structure
func hostEndpointAPIToBackend(ah *HostEndpoint) *backend.HostEndpoint {
	var ipv4Addrs []IP
	var ipv6Addrs []IP
	if ah.Spec.ExpectedIPs != nil {
		for _, ip := range ah.Spec.ExpectedIPs {
			if len(ip) == net.IPv4len {
				ipv4Addrs = append(ipv4Addrs, ip)
			} else {
				ipv6Addrs = append(ipv6Addrs, ip)
			}
		}
	}

	bh := backend.HostEndpoint{
		Hostname:   ah.Metadata.Hostname,
		EndpointID: ah.Metadata.Name,
		Labels:     ah.Metadata.Labels,

		Name:       ah.Spec.InterfaceName,
		ProfileIDs: ah.Spec.Profiles,
	}
	if len(ipv4Addrs) > 0 {
		bh.ExpectedIPv4Addrs = ipv4Addrs
	}
	if len(ipv6Addrs) > 0 {
		bh.ExpectedIPv6Addrs = ipv4Addrs
	}

	return &bh
}

// Convert a Backend HostEndpoint structure to an API Tier structure
func hostEndpointBackendToAPI(bh *backend.HostEndpoint) *HostEndpoint {
	ah := NewHostEndpoint()
	var ips []IP
	if bh.ExpectedIPv4Addrs != nil {
		ips = append(ips, bh.ExpectedIPv4Addrs...)
	}
	if bh.ExpectedIPv6Addrs != nil {
		ips = append(ips, bh.ExpectedIPv6Addrs...)
	}

	ah.Metadata.Hostname = bh.Hostname
	ah.Metadata.Name = bh.EndpointID
	ah.Metadata.Labels = bh.Labels

	ah.Spec.InterfaceName = bh.Name
	ah.Spec.Profiles = bh.ProfileIDs
	if len(ips) > 0 {
		ah.Spec.ExpectedIPs = &ips
	}

	return &ah
}
