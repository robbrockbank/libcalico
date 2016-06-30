package backend

import (
	"errors"
	"fmt"
)

type TierKey struct {
	Name string `json:"-" validate:"required,name"`
}

func (key TierKey) asEtcdKey() (string, error) {
	if key.Name == "" {
		return "", errors.New("insufficient identifiers")
	}
	e := fmt.Sprintf("/calico/v1/policy/tier/%s", key.Name)
	return e, nil
}

type TierListOptions struct {
	Name string
}

func (options TierListOptions) asEtcdKeyRegex() (string, error) {
	e := fmt.Sprintf("/calico/v1/policy/tier/%s",
		idOrWildcard(options.Name))
	return e, nil
}

type Tier struct {
	TierKey `json:"-"`
	Order   *float32 `json:"order,omitempty"`
}
