package client

import (
	. "github.com/projectcalico/libcalico/lib/backend/objects"
)

// ProfileInterface has methods to work with Service resources.
type ProfileInterface interface {
	List(hostname, name *string) ([]Profile, error)
	Get(hostname, name string) (*Profile, error)
	Create(hostname, name string, data *Profile) error
	Update(hostname, name string, data *Profile) error
	Delete(hostname, name string) error
}

// services implements ServicesNamespacer interface
type profiles struct {
	r  *Client
}

// newServices returns a services
func newProfiles(c *Client) *profiles {
	return &profiles{c}
}

// List takes a Metadata, and returns the list of hot endpoints that match that Metadata
// (wildcarding mising fields)
func (p *profiles) List(hostname, name *string) ([]Profile, error) {
	return nil, nil
}

// Get returns information about a particular profile.
func (p *profiles) Get(hostname, name string) (*Profile, error) {
	return nil, nil
}

// Create creates a new profile.
func (p *profiles) Create(hostname, name string, data *Profile) error {
	return nil

}

// Update updates an existing profile.
func (p *profiles) Update(hostname, name string, data *Profile) error {
	return nil

}

// Delete deletes an existing profile.
func (p *profiles) Delete(hostname, name string) error {
	return nil

}