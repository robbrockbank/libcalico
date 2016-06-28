package client

import (
	. "github.com/projectcalico/libcalico/lib/api"
	backend "github.com/projectcalico/libcalico/lib/backend/objects"
)

// ProfileInterface has methods to work with Profile resources.
type ProfileInterface interface {
	List(metadata *ProfileMetadata) (*ProfileList, error)
	Get(metadata *ProfileMetadata) (*Profile, error)
	Create(hep *Profile) (*Profile, error)
	Update(hep *Profile) (*Profile, error)
	Delete(metadata *ProfileMetadata) error
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
func (p *profiles) List(metadata *ProfileMetadata) (*ProfileList, error) {
	bplo := backend.ProfileListOptions{
		Name: metadata.Name,
	}
	if bps, err := p.c.backend.Profiles().List(bplo); err != nil {
		return nil, err
	} else {
		apl := NewProfileList()
		apl.Items = bps
		return &apl, nil
	}
}

// Get returns information about a particular profile.
func (p *profiles) Get(metadata *ProfileMetadata) (*Profile, error) {
	if bp, err := p.c.backend.Profiles().Get(metadata.Name); err != nil {
		return nil, err
	} else {
		return profileBackendToAPI(bp), nil
	}
}

// Create creates a new profile.
func (p *profiles) Create(ap *Profile) (*Profile, error) {
	if bp, err := p.c.Profiles().Create(profileAPIToBackend(ap)); err != nil {
		return nil, err
	} else {
		return profileBackendToAPI(bp), nil
	}
}

// Update updates an existing profile.
func (p *profiles) Update(ap *Profile) (*Profile, error) {
	if bp, err := p.c.Profiles().Update(profileAPIToBackend(ap)); err != nil {
		return nil, err
	} else {
		return profileBackendToAPI(bp), nil
	}
}

// Delete deletes an existing profile.
func (p *profiles) Delete(metadata *ProfileMetadata) error {
	return p.c.backend.Profiles().Delete(metadata.Name)
}

// Convert a API Profile structure to the Backend Profile structures.
// For profiles there is a one-to-many mapping.
func profileAPIToBackend(ap *Profile) *backend.Profile {
	bp := backend.Profile{
		Name: ap.Metadata.Name,
		Rules: backend.ProfileRules{
			InboundRules:  rulesAPIToBackend(ap.Spec.IngressRules),
			OutboundRules: rulesAPIToBackend(ap.Spec.EgressRules),
		},
		Tags:   ap.Spec.Tags,
		Labels: ap.Metadata.Labels,
	}

	return &bp
}

// Convert the Backend Profile structures to an API Profile structure
// For profiles there is a many-to-one mapping.
func profileBackendToAPI(bp *backend.Profile) *Profile {
	ap := NewProfile()
	ap.Metadata.Name = bp.Name
	ap.Metadata.Labels = bp.Labels

	ap.Spec.IngressRules = rulesBackendToAPI(bp.Rules.InboundRules)
	ap.Spec.EgressRules = rulesBackendToAPI(bp.Rules.OutboundRules)
	ap.Spec.Tags = bp.Tags

	return &ap
}
