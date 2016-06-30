package backend

import (
	"errors"
	"fmt"
)

// The profile structure is defined to allow the client to define a conversion interface
// to map between the API and backend profiles.  However, in the actual underlying
// implementation the profile is written as three separate entries - rules, tags and labels.
type Profile struct {
	ProfileKey
	Rules  ProfileRules
	Tags   []string
	Labels map[string]string
}

type ProfileRules struct {
	InboundRules  []Rule `json:"inbound_rules,omitempty" validate:"omitempty,dive"`
	OutboundRules []Rule `json:"outbound_rules,omitempty" validate:"omitempty,dive"`
}

// The profile key actually returns the common parent of the three separate entries.
// It is useful to define this to re-use some of the common machinery, and can be used
// for delete processing since delete needs to remove the common parent.
type ProfileKey struct {
	Name string `json:"-" validate:"required,name"`
}

func (key ProfileKey) asEtcdKey() (string, error) {
	if key.Name == "" {
		return "", errors.New("insufficient identifiers")
	}
	e := fmt.Sprintf("/calico/v1/policy/profile/%s", key.Name)
	return e, nil
}

func (key ProfileKey) asEtcdDeleteKey() (string, error) {
	return key.asEtcdKey()
}

// ProfileRulesKey implements the KeyInterface for the profile rules
type ProfileRulesKey struct {
	ProfileKey
}

func (key ProfileRulesKey) asEtcdKey() (string, error) {
	e, err := key.ProfileKey.asEtcdKey()
	return e + "/rules", err
}

// ProfileTagsKey implements the KeyInterface for the profile tags
type ProfileTagsKey struct {
	ProfileKey
}

func (key ProfileTagsKey) asEtcdKey() (string, error) {
	e, err := key.ProfileKey.asEtcdKey()
	return e + "/tags", err
}

// ProfileLabelsKey implements the KeyInterface for the profile labels
type ProfileLabelsKey struct {
	ProfileKey
}

func (key ProfileLabelsKey) asEtcdKey() (string, error) {
	e, err := key.ProfileKey.asEtcdKey()
	return e + "/labels", err
}
