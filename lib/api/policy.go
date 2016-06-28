package api

import (
	. "github.com/projectcalico/libcalico/lib/api/unversioned"
)

type PolicyMetadata struct {
	ObjectMetadata
	Tier *string `json:"tier,omitempty" validate:"omitempty,name"`
}

type PolicySpec struct {
	Order        *float32 `json:"order"`
	IngressRules *[]Rule  `json:"ingress,omitempty" validate:"omitempty,dive"`
	EgressRules  *[]Rule  `json:"egress,omitempty" validate:"omitempty,dive"`
	Selector     string   `json:"selector" validate:"selector"`
}

type Policy struct {
	TypeMetadata
	Metadata PolicyMetadata `json:"metadata"`
	Spec     PolicySpec     `json:"spec"`
}

func NewPolicy() *Policy {
	return &Policy{TypeMetadata: TypeMetadata{Kind: "policy", APIVersion: "v1"}}
}

type PolicyList struct {
	TypeMetadata
	Metadata ListMetadata `json:"metadata"`
	Items    []Policy     `json:"items" validate:"dive"`
}

func NewPolicyList() *PolicyList {
	return &PolicyList{TypeMetadata: TypeMetadata{Kind: "policyList", APIVersion: "v1"}}
}
