package objects

import (
	. "github.com/projectcalico/libcalico/lib/common"
)

type Tier struct {
	Name string `json:"-" validate:"required,name"`

	Order Order `json:"order" validate:"required"`
}

type TierListOptions struct {
	Name     *string
}