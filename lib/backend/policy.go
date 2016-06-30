package backend

import (
	"errors"
	"fmt"
)

type PolicyKey struct {
	Name string `json:"-" validate:"required,name"`
	Tier string `json:"-" validate:"required,name"`
}

func (key PolicyKey) asEtcdKey() (string, error) {
	if key.Name == "" || key.Tier == "" {
		return "", errors.New("insufficient identifiers")
	}
	e := fmt.Sprintf("/calico/v1/policy/tier/%s/policy/%s",
		key.Tier, key.Name)
	return e, nil
}

type PolicyListOptions struct {
	Name string
	Tier string
}

func (options PolicyListOptions) asEtcdKeyRegex() (string, error) {
	e := fmt.Sprintf("/calico/v1/policy/tier/%s/policy/%s",
		idOrWildcard(options.Tier),
		idOrWildcard(options.Name))
	return e, nil
}

type Policy struct {
	PolicyKey     `json:"-"`
	Order         *float32 `json:"order,omitempty"`
	InboundRules  []Rule   `json:"inbound_rules,omitempty" validate:"omitempty,dive"`
	OutboundRules []Rule   `json:"outbound_rules,omitempty" validate:"omitempty,dive"`
	Selector      string   `json:"selector" validate:"selector"`
}
