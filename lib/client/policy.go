package client

import (
	"github.com/projectcalico/libcalico/lib/api"
	"github.com/projectcalico/libcalico/lib/backend"
)

// PolicyInterface has methods to work with Policy resources.
type PolicyInterface interface {
	List(api.PolicyMetadata) (*api.PolicyList, error)
	Get(api.PolicyMetadata) (*api.Policy, error)
	Create(*api.Policy) (*api.Policy, error)
	Update(*api.Policy) (*api.Policy, error)
	Delete(api.PolicyMetadata) error
}

// policies implements PolicyInterface
type policies struct {
	c *Client
}

// newPolicies returns a policies
func newPolicies(c *Client) *policies {
	return &policies{c}
}

// List takes a Metadata, and returns the list of policies that match that Metadata
// (wildcarding missing fields)
func (h *policies) List(metadata api.PolicyMetadata) (*api.PolicyList, error) {
	if l, err := h.c.list(backend.Policy{}, metadata, h); err != nil {
		return nil, err
	} else {
		hl := api.NewPolicyList()
		hl.Items = make([]api.Policy, len(l))
		for _, h := range l {
			hl.Items = append(hl.Items, h.(api.Policy))
		}
		return hl, nil
	}
}

// Get returns information about a particular policy.
func (h *policies) Get(metadata api.PolicyMetadata) (*api.Policy, error) {
	if a, err := h.c.get(backend.Policy{}, metadata, h); err != nil {
		return nil, err
	} else {
		h := a.(api.Policy)
		return &h, nil
	}
}

// Create creates a new policy.
func (h *policies) Create(a *api.Policy) (*api.Policy, error) {
	if na, err := h.c.create(*a, h); err != nil {
		return nil, err
	} else {
		nh := na.(api.Policy)
		return &nh, nil
	}
}

// Create creates a new policy.
func (h *policies) Update(a *api.Policy) (*api.Policy, error) {
	if na, err := h.c.update(*a, h); err != nil {
		return nil, err
	} else {
		nh := na.(api.Policy)
		return &nh, nil
	}
}

// Delete deletes an existing policy.
func (h *policies) Delete(metadata api.PolicyMetadata) error {
	return h.c.delete(metadata, h)
}

// Convert a PolicyMetadata to a PolicyListInterface
func (h *policies) convertMetadataToListInterface(m interface{}) (backend.ListInterface, error) {
	pm := m.(api.PolicyMetadata)
	l := backend.PolicyListOptions{
		Name: pm.Name,
		Tier: pm.Tier,
	}
	return l, nil
}

// Convert a PolicyMetadata to a PolicyKeyInterface
func (h *policies) convertMetadataToKeyInterface(m interface{}) (backend.KeyInterface, error) {
	pm := m.(api.PolicyMetadata)
	k := backend.PolicyKey{
		Name: pm.Name,
		Tier: pm.Tier,
	}
	return k, nil
}

// Convert an API Policy structure to a Backend Policy structure
func (h *policies) convertAPIToBackend(a interface{}) (interface{}, error) {
	ap := a.(api.Policy)
	k, err := h.convertMetadataToKeyInterface(ap.Metadata)
	if err != nil {
		return nil, err
	}
	pk := k.(backend.PolicyKey)

	bp := backend.Policy{
		PolicyKey: pk,

		Order:         ap.Spec.Order,
		InboundRules:  rulesAPIToBackend(ap.Spec.IngressRules),
		OutboundRules: rulesAPIToBackend(ap.Spec.EgressRules),
		Selector:      ap.Spec.Selector,
	}

	return bp, nil
}

// Convert a Backend Policy structure to an API Policy structure
func (h *policies) convertBackendToAPI(b interface{}) (interface{}, error) {
	bp := b.(backend.Policy)
	ap := api.NewPolicy()

	ap.Metadata.Name = bp.Name

	ap.Spec.Order = bp.Order
	ap.Spec.IngressRules = rulesBackendToAPI(bp.InboundRules)
	ap.Spec.EgressRules = rulesBackendToAPI(bp.OutboundRules)
	ap.Spec.Selector = bp.Selector

	return ap, nil
}
