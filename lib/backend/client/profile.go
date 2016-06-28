package client

import (
	. "github.com/projectcalico/libcalico/lib/backend/objects"
)

// ProfileInterface has methods to work with Profile objects.
type ProfileInterface interface {
	List(bplo *ProfileListOptions) ([]Profile, error)
	Get(name string) (*Profile, error)
	Create(bp *Profile) error
	Update(bp *Profile) error
	Delete(name string) error
}

// profiles implements ProfileInterface
type profiles struct {
	r *Client
}

// newProfiles returns a profiles
func newProfiles(c *Client) *profiles {
	return &profiles{c}
}

// List takes a Metadata, and returns the list of profiles that match that Metadata
// (wildcarding missing fields)
func (p *profiles) List(bplo *ProfileListOptions) ([]Profile, error) {
	return nil, nil
}

// Get returns information about a particular profile.
func (p *profiles) Get(name string) (*Profile, error) {
	return nil, nil
}

// Create creates a new profile.
func (p *profiles) Create(bp *Profile) error {
	return nil

}

// Update updates an existing profile.
func (p *profiles) Update(bp *Profile) error {
	return nil

}

// Delete deletes an existing profile.
func (p *profiles) Delete(name string) error {
	return nil

}
