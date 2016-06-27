package api

import (
	. "github.com/projectcalico/libcalico/lib/api/unversioned"
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

func NewProfile() *Profile {
	return &Profile{TypeMetadata: TypeMetadata{Kind: "profile", APIVersion: "v1"}}
}

type ProfileList struct {
	TypeMetadata
	Metadata ListMetadata `json:"metadata"`
	Items    []Profile    `json:"items" validate:"dive"`
}

func NewProfileList() *Profile {
	return &Profile{TypeMetadata: TypeMetadata{Kind: "profile", APIVersion: "v1"}}
}
