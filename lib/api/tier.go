package api

import (
	. "github.com/projectcalico/libcalico/lib/api/unversioned"
)

type TierMetadata ObjectMetadata

type TierSpec struct {
	Order *float32 `json:"order,omitempty"`
}

type Tier struct {
	TypeMetadata
	Metadata TierMetadata `json:"metadata,omitempty"`
	Spec     TierSpec     `json:"spec,omitempty"`
}

func NewTier() *Tier {
	return &Tier{TypeMetadata: TypeMetadata{Kind: "tier", APIVersion: "v1"}}
}

type TierList struct {
	TypeMetadata
	Metadata ListMetadata `json:"metadata,omitempty"`
	Items    []Tier       `json:"items,omitempty" validate:"dive"`
}

func NewTierList() *TierList {
	return &TierList{TypeMetadata: TypeMetadata{Kind: "tierList", APIVersion: "v1"}}
}
