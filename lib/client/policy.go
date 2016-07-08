// Copyright (c) 2016 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"fmt"

	"github.com/projectcalico/libcalico/lib/api"
	"github.com/projectcalico/libcalico/lib/backend"
	"github.com/projectcalico/libcalico/lib/common"
)

var (
	defaultTier = api.Tier{
		Metadata: api.TierMetadata{Name: common.DefaultTierName},
	}
)

// PolicyInterface has methods to work with Policy resources.
type PolicyInterface interface {
	List(api.PolicyMetadata) (*api.PolicyList, error)
	Get(api.PolicyMetadata) (*api.Policy, error)
	Create(*api.Policy) (*api.Policy, error)
	Update(*api.Policy) (*api.Policy, error)
	Apply(*api.Policy) (*api.Policy, error)
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
	if l, err := h.c.list(backend.Policy{}, metadata, h, nil); err != nil {
		return nil, err
	} else {
		hl := api.NewPolicyList()
		hl.Items = make([]api.Policy, 0, len(l))
		for _, h := range l {
			hl.Items = append(hl.Items, *h.(*api.Policy))
		}
		return hl, nil
	}
}

// Get returns information about a particular policy.
func (h *policies) Get(metadata api.PolicyMetadata) (*api.Policy, error) {
	if a, err := h.c.get(backend.Policy{}, metadata, h, nil); err != nil {
		return nil, err
	} else {
		h := a.(api.Policy)
		return &h, nil
	}
}

// Create creates a new policy.
func (h *policies) Create(a *api.Policy) (*api.Policy, error) {
	// Before creating the policy, check that the tier exists, and if this is the
	// default tier, create it if it doesn't.
	if a.Metadata.Tier == "" {
		if _, err := h.c.Tiers().Create(&defaultTier); err != nil {
			if _, ok := err.(common.ErrorResourceAlreadyExists); !ok {
				return nil, err
			}
		}
	} else if _, err := h.c.Tiers().Get(api.TierMetadata{Name: a.Metadata.Tier}); err != nil {
		return nil, common.ErrorResourceDoesNotExist{Name: fmt.Sprintf("Tier '%s'", a.Metadata.Tier)}
	}

	return a, h.c.create(*a, h, nil)
}

// Update updates a new policy.
func (h *policies) Update(a *api.Policy) (*api.Policy, error) {
	return a, h.c.update(*a, h, nil)
}

// Apply creates or updates a policy.
func (h *policies) Apply(a *api.Policy) (*api.Policy, error) {
	// Before creating the policy, check that the tier exists, and if this is the
	// default tier, create it if it doesn't.
	if a.Metadata.Tier == "" {
		if _, err := h.c.Tiers().Create(&defaultTier); err != nil {
			if _, ok := err.(common.ErrorResourceAlreadyExists); !ok {
				return nil, err
			}
		}
	} else if _, err := h.c.Tiers().Get(api.TierMetadata{Name: a.Metadata.Tier}); err != nil {
		return nil, common.ErrorResourceDoesNotExist{Name: fmt.Sprintf("Tier '%s'", a.Metadata.Tier)}
	}

	return a, h.c.apply(*a, h, nil)
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
		Tier: common.TierOrDefault(pm.Tier),
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

// Convert a Backend Policy structure to an API Policy structure.
func (h *policies) convertBackendToAPI(b interface{}) (interface{}, error) {
	bp := *b.(*backend.Policy)
	ap := api.NewPolicy()

	ap.Metadata.Name = bp.Name
	ap.Metadata.Tier = common.TierOrBlank(bp.Tier)

	ap.Spec.Order = bp.Order
	ap.Spec.IngressRules = rulesBackendToAPI(bp.InboundRules)
	ap.Spec.EgressRules = rulesBackendToAPI(bp.OutboundRules)
	ap.Spec.Selector = bp.Selector

	return ap, nil
}

func (h *policies) copyKeyValues(kvs []backend.KeyValue, b interface{}) {
	bp := b.(*backend.Policy)
	k := kvs[0].Key.(backend.PolicyKey)
	bp.PolicyKey = k
}
