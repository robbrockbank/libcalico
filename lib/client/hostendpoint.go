package client

import (
	"errors"
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
	if bhlo, err := h.toListOptions(metadata); err != nil {
		return nil, err
	} else if bhs, err := h.c.backend.HostEndpoints().List(bhlo); err != nil {
		return nil, err
	} else {
		ahl := NewHostEndpointList()
		for _, bh := range bhs {
			if ah, err := h.backendToAPI(&bh); err != nil {
				ahl.Items = append(ahl.Items, *ah)
			}
		}
		return ahl, nil
	}
}

// Get returns information about a particular host endpoint.
func (h *hostEndpoints) Get(metadata *HostEndpointMetadata) (*HostEndpoint, error) {
	if bk, err := h.toBackendKey(metadata); err!= nil {
		return nil, err
	} else if bh, err := h.c.backend.HostEndpoints().Get(bk); err != nil {
		return nil, err
	} else {
		return h.backendToAPI(bh)
	}
}

// Create creates a new host endpoint.
func (h *hostEndpoints) Create(ah *HostEndpoint) (*HostEndpoint, error) {
	if bh, err := h.apiToBackend(ah); err != nil {
		return nil, err
	} else if err = h.c.backend.HostEndpoints().Create(bh); err != nil {
		return nil, err
	} else {
		return h.backendToAPI(bh)
	}
}

// Update updates an existing host endpoint.
func (h *hostEndpoints) Update(ah *HostEndpoint) (*HostEndpoint, error) {
	if bh, err := h.apiToBackend(ah); err != nil {
		return nil, err
	} else if err = h.c.backend.HostEndpoints().Update(bh); err != nil {
		return nil, err
	} else {
		return h.backendToAPI(bh)
	}
}

// Delete deletes an existing host endpoint.
func (h *hostEndpoints) Delete(metadata *HostEndpointMetadata) error {
	if bk, err := h.toBackendKey(metadata); err!= nil {
		return err
	} else {
		return h.c.backend.HostEndpoints().Delete(bk)
	}
}

func (h *hostEndpoints) toListOptions(m *HostEndpointMetadata) (*backend.HostEndpointListOptions, error) {
	bhlo := backend.HostEndpointListOptions{
		Hostname:   m.Hostname,
		EndpointID: m.Name,
	}
	return &bhlo, nil
}

func (h *hostEndpoints) toBackendKey(m *HostEndpointMetadata) (*backend.HostEndpointKey, error) {
	if m == nil || m.Name == nil || m.Hostname == nil {
		return nil, errors.New("insufficient identifiers supplied")
	}
	k := backend.HostEndpointKey{
		Hostname: *(m.Hostname),
		EndpointID: *(m.Name),
	}
	return &k, nil
}

// Convert an API HostEndpoint structure to a Backend Tier structure
func (h *hostEndpoints) apiToBackend(ah *HostEndpoint) (*backend.HostEndpoint, error) {
	k, err := h.toBackendKey(&ah.Metadata)
	if err != nil {
		return nil, err
	}

	var ipv4Addrs []IP
	var ipv6Addrs []IP
	if ah.Spec.ExpectedIPs != nil {
		for _, ip := range *ah.Spec.ExpectedIPs {
			if len(ip.IP) == net.IPv4len {
				ipv4Addrs = append(ipv4Addrs, ip)
			} else {
				ipv6Addrs = append(ipv6Addrs, ip)
			}
		}
	}

	bh := backend.HostEndpoint{
		HostEndpointKey: *k,
		Labels:     ah.Metadata.Labels,

		Name:       ah.Spec.InterfaceName,
		ProfileIDs: ah.Spec.Profiles,
	}
	if len(ipv4Addrs) > 0 {
		bh.ExpectedIPv4Addrs = &ipv4Addrs
	}
	if len(ipv6Addrs) > 0 {
		bh.ExpectedIPv6Addrs = &ipv4Addrs
	}

	return &bh, nil
}

// Convert a Backend HostEndpoint structure to an API Tier structure
func (h *hostEndpoints) backendToAPI(bh *backend.HostEndpoint) (*HostEndpoint, error) {
	ah := NewHostEndpoint()
	var ips []IP
	if bh.ExpectedIPv4Addrs != nil {
		ips = append(ips, *bh.ExpectedIPv4Addrs...)
	}
	if bh.ExpectedIPv6Addrs != nil {
		ips = append(ips, *bh.ExpectedIPv6Addrs...)
	}

	ah.Metadata.Hostname = &bh.HostEndpointKey.Hostname
	ah.Metadata.Name = &bh.HostEndpointKey.EndpointID
	ah.Metadata.Labels = bh.Labels

	ah.Spec.InterfaceName = bh.Name
	ah.Spec.Profiles = bh.ProfileIDs
	if len(ips) > 0 {
		ah.Spec.ExpectedIPs = &ips
	}

	return ah, nil
}
