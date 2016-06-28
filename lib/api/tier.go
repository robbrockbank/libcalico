package api

import (
	. "github.com/projectcalico/libcalico/lib/api/unversioned"
	. "github.com/projectcalico/libcalico/lib/common"
)

type TierMetadata ObjectMetadata

type TierSpec struct {
	Order *float32 `json:"order"`
}

type Tier struct {
	TypeMetadata
	Metadata TierMetadata `json:"metadata"`
	Spec     TierSpec     `json:"spec"`
}

func NewTier() *Tier {
	return &Tier{TypeMetadata: TypeMetadata{Kind: "tier", APIVersion: "v1"}}
}

type TierList struct {
	TypeMetadata
	Metadata ListMetadata `json:"metadata"`
	Items    []Tier       `json:"items" validate:"dive"`
}

func NewTierList() *TierList {
	return &TierList{TypeMetadata: TypeMetadata{Kind: "tierList", APIVersion: "v1"}}
}
