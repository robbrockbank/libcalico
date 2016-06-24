package v1

import (
	"github.com/projectcalico/libcalico/lib/api/unversioned"
	. "github.com/projectcalico/libcalico/lib/common"
)

type TierMetadata ObjectMetadata

type TierSpec struct {
	Order Order `json:"order" validate:"required"`
}

type Tier struct {
	unversioned.TypeMetadata
	Metadata TierMetadata `json:"metadata"`
	Spec     TierSpec     `json:"spec"`
}

type TierList struct {
	unversioned.TypeMetadata
	Metadata ListMetadata `json:"metadata"`
	Items    []Tier       `json:"items" validate:"dive"`
}
