package objects

import (
	. "github.com/projectcalico/libcalico/lib/common"
)

type Tier struct {
	Name string `json:"-" validate:"required,name"`

	Order *float32 `json:"order"`
}

type TierListOptions struct {
	Name *string
}
