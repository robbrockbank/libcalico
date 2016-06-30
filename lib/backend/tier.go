package backend

import (
	"fmt"
	"github.com/projectcalico/libcalico/lib/common"
)

type TierKey struct {
	Name string `json:"-" validate:"required,name"`
}

func (key TierKey) asEtcdKey() (string, error) {
	k, err := key.asEtcdDeleteKey()
	return k + "/metadata", err
}

func (key TierKey) asEtcdDeleteKey() (string, error) {
	if key.Name == "" {
		return "", common.ErrorInsufficientIdentifiers{}
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
