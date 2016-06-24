package v1

import (
	"github.com/projectcalico/libcalico/lib/api/unversioned"
	. "github.com/projectcalico/libcalico/lib/common"
)

/*
The v1 structure definitions used in the V1 interface.
*/

// ---- Metadata common to all resources ----
type ObjectMetadata struct {
	Name string `json:"name" validate:"name"`
}

// ---- Metadata common to all lists ----
type ListMetadata struct {
}

// ----  Tier  ----
// Kind:  tier

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

// ----  Policy  ----
// Kind: policy

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
	unversioned.TypeMetadata
	Metadata PolicyMetadata `json:"metadata"`
	Spec     PolicySpec     `json:"spec"`
}

type PolicyList struct {
	unversioned.TypeMetadata
	Metadata ListMetadata `json:"metadata"`
	Items    []Policy     `json:"items" validate:"dive"`
}

// ----  Profile  ----
// Kind: profile

type ProfileMetadata ObjectMetadata

type ProfileSpec struct {
	IngressRules *[]Rule            `json:"ingress,omitempty" validate:"omitempty,dive"`
	EgressRules  *[]Rule            `json:"egress,omitempty" validate:"omitempty,dive"`
	Labels       *map[string]string `json:"labels,omitempty" validate:"omitempty,labels"`
	Tags         *[]string          `json:"tags,omitempty" validate:"omitempty,dive,tag"`
}

type Profile struct {
	unversioned.TypeMetadata
	Metadata ProfileMetadata `json:"metadata"`
	Spec     ProfileSpec     `json:"spec"`
}

type ProfileList struct {
	unversioned.TypeMetadata
	Metadata ListMetadata `json:"metadata"`
	Items    []Profile    `json:"items" validate:"dive"`
}

// ----  Rule (subtype of Profile and Policy)  ----

type Rule struct {
	Action string `json:"action" validate:"action"`

	Protocol    *Protocol `json:"protocol,omitempty" validate:"omitempty"`
	SrcTag      *string        `json:"srcTag,omitempty" validate:"omitempty,tag"`
	SrcNet      *IPNet         `json:"srcNet,omitempty" validate:"omitempty"`
	SrcSelector *string        `json:"srcSelector,omitempty" validate:"omitempty,selector"`
	SrcPorts    *[]int         `json:"srcPorts,omitempty" validate:"omitempty,dive,gte=0,lte=65535"`
	DstTag      *string        `json:"dstTag,omitempty" validate:"omitempty,tag"`
	DstSelector *string        `json:"dstSelector,omitempty" validate:"omitempty,selector"`
	DstNet      *IPNet         `json:"dstNet,omitempty" validate:"omitempty"`
	DstPorts    *[]int         `json:"dstPorts,omitempty" validate:"omitempty,dive,gte=0,lte=65535"`
	ICMPType    *int           `json:"icmpType,omitempty" validate:"omitempty,gte=0,lte=255"`
	ICMPCode    *int           `json:"icmpCode,omitempty" validate:"omitempty,gte=0,lte=255"`

	NotProtocol    *Protocol `json:"!protocol,omitempty" validate:"omitempty"`
	NotSrcTag      *string        `json:"!srcTag,omitempty" validate:"omitempty,tag"`
	NotSrcNet      *IPNet         `json:"!srcNet,omitempty" validate:"omitempty"`
	NotSrcSelector *string        `json:"!srcSelector,omitempty" validate:"omitempty,selector"`
	NotSrcPorts    *[]int         `json:"!srcPorts,omitempty" validate:"omitempty,dive,gte=0,lte=65535"`
	NotDstTag      *string        `json:"!dstTag,omitempty" validate:"omitempty"`
	NotDstSelector *string        `json:"!dstSelector,omitempty" validate:"omitempty,selector"`
	NotDstNet      *IPNet         `json:"!dstNet,omitempty" validate:"omitempty"`
	NotDstPorts    *[]int         `json:"!dstPorts,omitempty" validate:"omitempty,dive,gte=0,lte=65535"`
	NotICMPType    *int           `json:"!icmpType,omitempty" validate:"omitempty,gte=0,lte=255"`
	NotICMPCode    *int           `json:"!icmpCode,omitempty" validate:"omitempty,gte=0,lte=255"`
}

// ----  Host Endpoint  ----
// Kind: host-endpoint

type HostEndpointMetadata struct {
	ObjectMetadata
	Hostname string `json:"hostname" valid:"hostname"`
}

type HostEndpointSpec struct {
	InterfaceName *string            `json:"interfaceName" validate:"omitempty,interface"`
	ExpectedIPs   *[]IP              `json:"expectedIPs" validate:"omitempty,dive,ip"`
	Labels        *map[string]string `json:"labels" validate:"omitempty,labels"`
	Profiles      *[]string          `json:"profiles" validate:"omitempty,dive,name"`
}

type HostEndpoint struct {
	unversioned.TypeMetadata
	Metadata HostEndpointMetadata `json:"metadata"`
	Spec     HostEndpointSpec     `json:"spec"`
}

type HostEndpointList struct {
	unversioned.TypeMetadata
	Metadata ListMetadata   `json:"metadata"`
	Items    []HostEndpoint `json:"items" validate:"dive"`
}
