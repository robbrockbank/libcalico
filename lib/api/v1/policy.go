package v1

import (
	. "github.com/projectcalico/libcalico/lib/api/unversioned"
	. "github.com/projectcalico/libcalico/lib/common"
)

type PolicyMetadata struct {
	ObjectMetadata
	Tier *string `json:"tier,omitempty" validate:"omitempty,name"`
}

type PolicySpec struct {
	Order        Order `json:"order" validate:"required"`
	IngressRules *[]Rule       `json:"ingress,omitempty" validate:"omitempty,dive"`
	EgressRules  *[]Rule       `json:"egress,omitempty" validate:"omitempty,dive"`
	Selector     string        `json:"selector" validate:"selector"`
}

type Policy struct {
	TypeMetadata
	Metadata PolicyMetadata `json:"metadata"`
	Spec     PolicySpec     `json:"spec"`
}

type PolicyList struct {
	TypeMetadata
	Metadata ListMetadata `json:"metadata"`
	Items    []Policy     `json:"items" validate:"dive"`
}
