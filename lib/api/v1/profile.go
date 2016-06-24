package v1

import (
	. "github.com/projectcalico/libcalico/lib/api/unversioned"
	. "github.com/projectcalico/libcalico/lib/common"
)

type ProfileMetadata ObjectMetadata

type ProfileSpec struct {
	IngressRules *[]Rule            `json:"ingress,omitempty" validate:"omitempty,dive"`
	EgressRules  *[]Rule            `json:"egress,omitempty" validate:"omitempty,dive"`
	Labels       *map[string]string `json:"labels,omitempty" validate:"omitempty,labels"`
	Tags         *[]string          `json:"tags,omitempty" validate:"omitempty,dive,tag"`
}

type Profile struct {
	TypeMetadata
	Metadata ProfileMetadata `json:"metadata"`
	Spec     ProfileSpec     `json:"spec"`
}

type ProfileList struct {
	TypeMetadata
	Metadata ListMetadata `json:"metadata"`
	Items    []Profile    `json:"items" validate:"dive"`
}
