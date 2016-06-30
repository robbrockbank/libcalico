package client

import (
	"github.com/projectcalico/libcalico/lib/api"
	"github.com/projectcalico/libcalico/lib/backend"
)

// ProfileInterface has methods to work with Profile resources.
type ProfileInterface interface {
	List(api.ProfileMetadata) (*api.ProfileList, error)
	Get(api.ProfileMetadata) (*api.Profile, error)
	Create(*api.Profile) (*api.Profile, error)
	Update(*api.Profile) (*api.Profile, error)
	Delete(api.ProfileMetadata) error
}

// profiles implements ProfileInterface
type profiles struct {
	c *Client
}

// newProfiles returns a profiles
func newProfiles(c *Client) *profiles {
	return &profiles{c}
}

// List takes a Metadata, and returns the list of profiles that match that Metadata
// (wildcarding missing fields)
func (h *profiles) List(metadata api.ProfileMetadata) (*api.ProfileList, error) {
	if l, err := h.c.list(backend.Profile{}, metadata, h); err != nil {
		return nil, err
	} else {
		hl := api.NewProfileList()
		hl.Items = make([]api.Profile, len(l))
		for _, h := range l {
			hl.Items = append(hl.Items, h.(api.Profile))
		}
		return hl, nil
	}
}

// Get returns information about a particular profile.
func (h *profiles) Get(metadata api.ProfileMetadata) (*api.Profile, error) {
	if a, err := h.c.get(backend.Profile{}, metadata, h); err != nil {
		return nil, err
	} else {
		h := a.(api.Profile)
		return &h, nil
	}
}

// Create creates a new profile.
func (h *profiles) Create(a *api.Profile) (*api.Profile, error) {
	if na, err := h.c.create(*a, h); err != nil {
		return nil, err
	} else {
		nh := na.(api.Profile)
		return &nh, nil
	}
}

// Create creates a new profile.
func (h *profiles) Update(a *api.Profile) (*api.Profile, error) {
	if na, err := h.c.update(*a, h); err != nil {
		return nil, err
	} else {
		nh := na.(api.Profile)
		return &nh, nil
	}
}

// Delete deletes an existing profile.
func (h *profiles) Delete(metadata api.ProfileMetadata) error {
	return h.c.delete(metadata, h)
}

// Convert a ProfileMetadata to a ProfileListInterface
func (h *profiles) convertMetadataToListInterface(m interface{}) (backend.ListInterface, error) {
	hm := m.(api.ProfileMetadata)
	l := backend.ProfileListOptions{
		Name: hm.Name,
	}
	return l, nil
}

// Convert a ProfileMetadata to a ProfileKeyInterface
func (h *profiles) convertMetadataToKeyInterface(m interface{}) (backend.KeyInterface, error) {
	hm := m.(api.ProfileMetadata)
	k := backend.ProfileKey{
		Name: hm.Name,
	}
	return k, nil
}

// Convert an API Profile structure to a Backend Profile structure
func (h *profiles) convertAPIToBackend(a interface{}) (interface{}, error) {
	ap := a.(api.Profile)
	k, err := h.convertMetadataToKeyInterface(ap.Metadata)
	if err != nil {
		return nil, err
	}
	pk := k.(backend.ProfileKey)

	bp := backend.Profile{
		ProfileKey: pk,
		Rules: backend.ProfileRules{
			InboundRules:  rulesAPIToBackend(ap.Spec.IngressRules),
			OutboundRules: rulesAPIToBackend(ap.Spec.EgressRules),
		},
		Tags:   ap.Spec.Tags,
		Labels: ap.Metadata.Labels,
	}

	return bp, nil
}

// Convert a Backend Profile structure to an API Profile structure
func (h *profiles) convertBackendToAPI(b interface{}) (interface{}, error) {
	bp := b.(backend.Profile)
	ap := api.NewProfile()

	ap.Metadata.Name = bp.Name
	ap.Metadata.Labels = bp.Labels

	ap.Spec.IngressRules = rulesBackendToAPI(bp.Rules.InboundRules)
	ap.Spec.EgressRules = rulesBackendToAPI(bp.Rules.OutboundRules)
	ap.Spec.Tags = bp.Tags

	return ap, nil
}
