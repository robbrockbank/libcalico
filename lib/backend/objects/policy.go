package objects

import (
	. "github.com/projectcalico/libcalico/lib/common"
)

type Policy struct {
	Name string `json:"-" validate:"required,name"`

	Order        Order `json:"order" validate:"required"`
	InboundRules *[]Rule       `json:"inbound_rules,omitempty" validate:"omitempty,dive"`
	OutboundRules  *[]Rule       `json:"outbound_rules,omitempty" validate:"omitempty,dive"`
	Selector     string        `json:"selector" validate:"selector"`
}

type PolicyListOptions struct {
	Name     *string
}