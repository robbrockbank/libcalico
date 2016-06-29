package client

import (
	"errors"
	"net"

	"github.com/projectcalico/libcalico/lib/api"
	. "github.com/projectcalico/libcalico/lib/common"
	"github.com/projectcalico/libcalico/lib/backend"
)

// api.HostEndpointInterface has methods to work with api.HostEndpoint resources.
type HostEndpointInterface interface {
	List(api.HostEndpointMetadata) (*api.HostEndpointList, error)
	Get(api.HostEndpointMetadata) (*api.HostEndpoint, error)
	Create(api.HostEndpoint) (*api.HostEndpoint, error)
	Update(api.HostEndpoint) (*api.HostEndpoint, error)
	Delete(api.HostEndpointMetadata) error
}

// hostEndpoints implements api.HostEndpointInterface
type hostEndpoints struct {
	c *Client
}

// newapi.HostEndpoints returns a hostEndpoints
func newHostEndpoints(c *Client) *hostEndpoints {
	return &hostEndpoints{c}
}

// List takes a Metadata, and returns the list of host endpoints that match that Metadata
// (wildcarding missing fields)
func (h *hostEndpoints) List(metadata api.HostEndpointMetadata) (*api.HostEndpointList, error) {
	if l, err := h.c.list(backend.HostEndpoint{}, metadata, h); err != nil {
		return nil, err
	} else {
		hl := api.NewHostEndpointList()
		hl.Items = make([]api.HostEndpoint, len(l))
		for _, h := range l {
			hl.Items = append(hl.Items, h.(api.HostEndpoint))
		}
		return hl, nil
	}
}

// Get returns information about a particular host endpoint.
func (h *hostEndpoints) Get(metadata api.HostEndpointMetadata) (*api.HostEndpoint, error) {
	if a, err := h.c.get(backend.HostEndpoint{}, metadata, h); err != nil {
		return nil, err
	} else {
		h := a.(api.HostEndpoint)
		return &h, nil
	}
}

// Create creates a new host endpoint.
func (h *hostEndpoints) Create(a api.HostEndpoint) (*api.HostEndpoint, error) {
	if err := h.c.create(a, h); err != nil {
		return nil, err
	} else {
		return &a, nil
	}
}

// Create creates a new host endpoint.
func (h *hostEndpoints) Update(a api.HostEndpoint) (*api.HostEndpoint, error) {
	if err := h.c.update(a, h); err != nil {
		return nil, err
	} else {
		return &a, nil
	}
}

// Delete deletes an existing host endpoint.
func (h *hostEndpoints) Delete(metadata api.HostEndpointMetadata) error {
	return h.c.delete(metadata, h)
}

func (h *hostEndpoints) convertMetadataToListInterface(m interface{}) (backend.ListInterface, error) {
	hm := m.(api.HostEndpointMetadata)
	l := backend.HostEndpointListOptions{
		Hostname:   hm.Hostname,
		EndpointID: hm.Name,
	}
	return l, nil
}

func (h *hostEndpoints) convertMetadataToKeyInterface(m interface{}) (backend.KeyInterface, error) {
	hm := m.(api.HostEndpointMetadata)
	if hm.Name == nil || hm.Hostname == nil {
		return nil, errors.New("insufficient identifiers supplied")
	}
	k := backend.HostEndpointKey{
		Hostname: *(hm.Hostname),
		EndpointID: *(hm.Name),
	}
	return k, nil
}

// Convert an API api.HostEndpoint structure to a Backend Tier structure
func (h *hostEndpoints) convertAPIToBackend(a interface{}) (interface{}, error) {
	ah := a.(api.HostEndpoint)
	k, err := h.convertMetadataToKeyInterface(ah.Metadata)
	if err != nil {
		return nil, err
	}
	hk := k.(backend.HostEndpointKey)

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
		HostEndpointKey: hk,
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

	return bh, nil
}

// Convert a Backend api.HostEndpoint structure to an API api.HostEndpoint structure
func (h *hostEndpoints) convertBackendToAPI(b interface{}) (interface{}, error) {
	ah := api.NewHostEndpoint()
	bh := b.(backend.HostEndpoint)
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
