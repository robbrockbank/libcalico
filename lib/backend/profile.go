package backend

import (
	"errors"
	"fmt"
)

type ProfileRules struct {
	InboundRules  []Rule `json:"inbound_rules,omitempty" validate:"omitempty,dive"`
	OutboundRules []Rule `json:"outbound_rules,omitempty" validate:"omitempty,dive"`
}

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

type ProfileListOptions struct {
	Name string
}

func (options ProfileListOptions) asEtcdKeyRegex() (string, error) {
	e := fmt.Sprintf("/calico/v1/policy/profile/%s", idOrWildcard(options.Name))
	return e, nil
}

type Profile struct {
	ProfileKey
	Rules  ProfileRules
	Tags   []string
	Labels map[string]string
}
