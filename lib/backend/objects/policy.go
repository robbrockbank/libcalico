package objects

import (
	. "github.com/projectcalico/libcalico/lib/common"
)

type Policy struct {
	Order        Order `json:"order" validate:"required"`
	IngressRules *[]Rule       `json:"ingress,omitempty" validate:"omitempty,dive"`
	EgressRules  *[]Rule       `json:"egress,omitempty" validate:"omitempty,dive"`
	Selector     string        `json:"selector" validate:"selector"`
}
