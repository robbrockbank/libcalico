package client

import (
	. "github.com/projectcalico/libcalico/lib/api"
	backend "github.com/projectcalico/libcalico/lib/backend/objects"
	"github.com/coreos/etcd/mvcc/backend"
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
	c  *CalicoClient
}

// newServices returns a services
func newPolicies(c *CalicoClient) *policies {
	return &policies{c}
}

// List takes a Metadata, and returns the list of policies that match that Metadata
// (wildcarding missing fields)
func (p *policies) List(metadata *PolicyMetadata) (*PolicyList, error) {
	bplo := backend.PolicyListOptions{
		Name: metadata.Name,
	}
	if bps, err := p.c.backend.Policies().List(bplo); err != nil {
		return nil, err
	} else {
		apl := NewPolicyList()
		apl.Items = bps
		return &apl, nil
	}
}

// Get returns information about a particular policy.
func (p *policies) Get(metadata *PolicyMetadata) (*Policy, error) {
	if bp, err := p.c.backend.Policies().Get(metadata.Name); err != nil {
		return nil, err
	} else {
		return policyBackendToAPI(bp), nil
	}
}

// Create creates a new policy.
func (p *policies) Create(ap *Policy) (*Policy, error) {
	if bp, err := p.c.Policies().Create(policyAPIToBackend(ap)); err != nil {
		return nil, err
	} else {
		return policyBackendToAPI(bp), nil
	}
}

// Update updates an existing policy.
func (p *policies) Update(ap *Policy) (*Policy, error) {
	if bp, err := p.c.Policies().Update(policyAPIToBackend(ap)); err != nil {
		return nil, err
	} else {
		return policyBackendToAPI(bp), nil
	}
}

// Delete deletes an existing policy.
func (p *policies) Delete(metadata *PolicyMetadata) error {
	return p.c.backend.Policies().Delete(metadata.Name)
}

// Convert an API Policy structure to a Backend Tier structure
func policyAPIToBackend(ap *Policy) *backend.Policy {
	bp := backend.Policy{
		Name: ap.Metadata.Name,

		Order: ap.Spec.Order,
		InboundRules: rulesAPIToBackend(ap.Spec.IngressRules),
		OutboundRules: rulesAPIToBackend(ap.Spec.EgressRules),
		Selector: ap.Spec.Selector,
	}

	return &bp
}

// Convert a Backend Policy structure to an API Tier structure
func policyBackendToAPI(bp *backend.Policy) *Policy {
	ap := NewPolicy()
	ap.Metadata.Name = bp.Name

	ap.Spec.Order = bp.Order
	ap.Spec.IngressRules = rulesBackendToAPI(bp.InboundRules)
	ap.Spec.EgressRules = rulesBackendToAPI(bp.OutboundRules)
	ap.Spec.Selector = bp.Selector

	return &ap
}
