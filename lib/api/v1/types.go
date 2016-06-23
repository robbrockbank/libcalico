package v1

import (
	. "net"

	"github.com/projectcalico/libcalico/lib/api/unversioned"
	"github.com/projectcalico/libcalico/lib/common"
)

/*
The v1 structure definitions used in the V1 interface.
*/

// ---- Metadata common to all resources ----
type ObjectMetadata struct {
	Name string `json:"name" valid:"name"`
}

// ---- Metadata common to all lists ----
type ListMetadata struct {
}

// ----  Tier  ----
// Kind:  tier

type TierMetadata ObjectMetadata

type TierSpec struct {
	Order intorstr.Int32OrString `json:"order" valid:"matches(default|\d+)"`
}

type Tier struct {
	unversioned.TypeMetadata
	Metadata TierMetadata `json:"metadata"`
	Spec     TierSpec     `json:"spec"`
}

type TierList struct {
	unversioned.TypeMetadata
	Metadata ListMetadata `json:"metadata"`
	Items    []Tier       `json:"items"`
}

// ----  Policy  ----
// Kind: policy

type PolicyMetadata struct {
	ObjectMetadata
	Tier *string `json:"tier,omitempty" valid:"matches([a-zA-Z0-9-_]+)"`
}

type PolicySpec struct {
	Order        intorstr.Int32OrString `json:"order"`
	IngressRules []Rule                 `json:"ingress"`
	EgressRules  []Rule                 `json:"egress"`
	Selector     string                 `json:"selector"`
}

type Policy struct {
	unversioned.TypeMetadata
	Metadata PolicyMetadata `json:"metadata"`
	Spec     PolicySpec     `json:"spec"`
}

type PolicyList struct {
	unversioned.TypeMetadata
	Metadata ListMetadata `json:"metadata"`
	Items    []Policy     `json:"items"`
}

// ----  Profile  ----
// Kind: profile

type ProfileMetadata ObjectMetadata

type ProfileSpec struct {
	IngressRules *[]Rule            `json:"ingress,omitempty"`
	EgressRules  *[]Rule            `json:"egress,omitempty"`
	Labels       *map[string]string `json:"labels,omitempty"`
	Tags         *[]string          `json:"tags,omitempty"`
}

type Profile struct {
	unversioned.TypeMetadata
	Metadata ProfileMetadata `json:"metadata"`
	Spec     ProfileSpec     `json:"spec"`
}

type ProfileList struct {
	unversioned.TypeMetadata
	Metadata ListMetadata `json:"metadata"`
	Items    []Profile    `json:"items"`
}

// ----  Rule (subtype of Profile and Policy)  ----

type Rule struct {
	Action string `json:"action" valid:"matches(deny|allow|next-tier)"`

	Protocol    *string `json:"protocol,omitempty"`
	SrcTag      *string `json:"srcTag,omitempty"`
	SrcNet      *IPNet  `json:"srcNet,omitempty"`
	SrcSelector *string `json:"srcSelector,omitempty"`
	SrcPorts    *[]int  `json:"srcPorts,omitempty"`
	DstTag      *string `json:"dstTag,omitempty"`
	DstSelector *string `json:"dstSelector,omitempty"`
	DstNet      *IPNet  `json:"dstNet,omitempty"`
	DstPorts    *[]int  `json:"dstPorts,omitempty"`
	IcmpType    *int    `json:"icmpType,omitempty"`
	IcmpCode    *int    `json:"icmpCode,omitempty"`

	NotProtocol    *string `json:"!protocol,omitempty"`
	NotSrcTag      *string `json:"!srcTag,omitempty"`
	NotSrcNet      *IPNet  `json:"!srcNet,omitempty"`
	NotSrcSelector *string `json:"!srcSelector,omitempty"`
	NotSrcPorts    *[]int  `json:"!srcPorts,omitempty"`
	NotDstTag      *string `json:"!dstTag,omitempty"`
	NotDstSelector *string `json:"!dstSelector,omitempty"`
	NotDstNet      *IPNet  `json:"!dstNet,omitempty"`
	NotDstPorts    *[]int  `json:"!dstPorts,omitempty"`
	NotIcmpType    *int    `json:"!icmpType,omitempty"`
	NotIcmpCode    *int    `json:"!icmpCode,omitempty"`
}

// ----  Host Endpoint  ----
// Kind: host-endpoint

type HostEndpointMetadata struct {
	ObjectMetadata
	Hostname string `json:"hostname" valid:"hostname"`
}

type HostEndpointSpec struct {
	InterfaceName *string            `json:"interface_name"`
	ExpectedAddrs *[]IP              `json:"expectedIPv4Addrs"`
	Labels        *map[string]string `json:"labels"`
	Profiles      *[]string          `json:"profiles"` // Perhaps profiles or profile_names
}

type HostEndpoint struct {
	unversioned.TypeMetadata
	Metadata HostEndpointMetadata `json:"metadata"`
	Spec     HostEndpointSpec     `json:"spec"`
}

type HostEndpointList struct {
	unversioned.TypeMetadata
	Metadata ListMetadata   `json:"metadata"`
	Items    []HostEndpoint `json:"items"`
}
