package client

import (
	. "github.com/projectcalico/libcalico/lib/api"
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
	r  *CalicoClient
}

// newServices returns a services
func newProfiles(c *CalicoClient) *profiles {
	return &profiles{c}
}

// List takes a Metadata, and returns the list of hot endpoints that match that Metdata
// (wildcarding mising fields)
func (p *profiles) List(metadata ProfileMetadata) (*ProfileList, error) {
	return nil, nil
}

// Get returns information about a particular profile.
func (p *profiles) Get(metadata ProfileMetadata) (*Profile, error) {
	return nil, nil
}

// Create creates a new profile.
func (p *profiles) Create(hep *Profile) (*Profile, error) {
	return nil, nil
	
}

// Update updates an existing profile.
func (p *profiles) Update(hep *Profile) (*Profile, error) {
	return nil, nil
	
}

// Delete deletes an existing profile.
func (p *profiles) Delete(metadata ProfileMetadata) error {
	return nil
	
}