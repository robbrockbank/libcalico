package client

import (
	. "github.com/projectcalico/libcalico/lib/api"
	backend "github.com/projectcalico/libcalico/lib/backend/objects"
	"github.com/coreos/etcd/mvcc/backend"
)

// ProfileInterface has methods to work with Profile resources.
type ProfileInterface interface {
	List(metadata ProfileMetadata) (*ProfileList, error)
	Get(metadata ProfileMetadata) (*Profile, error)
	Create(hep *Profile) (*Profile, error)
	Update(hep *Profile) (*Profile, error)
	Delete(metadata ProfileMetadata) error
}

// services implements ServicesNamespacer interface
type profiles struct {
	c  *CalicoClient
}

// newServices returns a services
func newProfiles(c *CalicoClient) *profiles {
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
func profileAPIToBackend(ap *Profile) (*backend.ProfileRules, *backend.ProfileTags, *backend.ProfileLabels) {
	bpr := backend.ProfileRules{
		Name: ap.Metadata.Name,

		InboundRules: rulesAPIToBackend(ap.Spec.IngressRules),
		OutboundRules: rulesAPIToBackend(ap.Spec.EgressRules),
	}
	bpt := backend.ProfileTags{
		Name: ap.Metadata.Name,

		Tags: ap.Spec.Tags,
	}
	bpl := backend.ProfileLabels{
		Name: ap.Metadata.Name,

		Labels: ap.Spec.Labels,
	}

	return &bpr, &bpt, &bpl
}

// Convert the Backend Profile structures to an API Profile structure
// For profiles there is a many-to-one mapping.
func profileBackendToAPI(bpr *backend.ProfileRules, bpt *backend.ProfileTags, bpl *backend.ProfileLabels) *Profile {
	ap := NewProfile()
	ap.Metadata.Name = bpr.Name

	ap.Spec.IngressRules = rulesBackendToAPI(bpr.InboundRules)
	ap.Spec.EgressRules = rulesBackendToAPI(bpr.OutboundRules)
	ap.Spec.Tags = bpt.Tags
	ap.Spec.Labels = bpl.Labels

	return &ap
}