package v1

import (
	"encoding/json"
	"fmt"
	"strconv"

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
	Order common.Order `json:"order" valid:"matches(default|\d*)"`
}

type Tier struct {
	unversioned.TypeMetadata
	Metadata TierMetadata  `json:"metadata"`
	Spec TierSpec  `json:"spec"`
}

type TierList struct {
	unversioned.TypeMetadata
	Metadata ListMetadata `json:"metadata"`
	Items []Tier  `json:"items"`
}

// ----  Policy  ----
// Kind: policy

type PolicyMetadata struct {
	ObjectMetadata
	Tier *string `json:"tier,omitempty" valid:"matches([a-zA-Z0-9-_]+)"`
}

type PolicySpec struct {
	Order         common.IntOrStr  `json:"order" valid:"matches(default|\d*)"`
	IngressRules  []Rule `json:"ingress"`
	EgressRules []Rule `json:"egress"`
	Selector      string `json:"selector" valid:"selector"`
}

type Policy struct {
	unversioned.TypeMetadata
	Metadata PolicyMetadata  `json:"metadata"`
	Spec PolicySpec  `json:"spec"`
}

type PolicyList struct {
	unversioned.TypeMetadata
	Metadata ListMetadata `json:"metadata"`
	Items []Policy  `json:"items"`
}

// ----  Profile  ----
// Kind: profile

type ProfileMetadata ObjectMetadata

type ProfileSpec struct {
	IngressRules  *[]Rule            `json:"ingress,omitempty"`
	EgressRules   *[]Rule            `json:"egress,omitempty"`
	Labels        *map[string]string `json:"labels,omitempty" valid:"matches([a-zA-Z0-9-_/]+)"`
	Tags          *[]string          `json:"tags,omitempty"`
}

type Profile struct {
	unversioned.TypeMetadata
	Metadata ProfileMetadata  `json:"metadata"`
	Spec ProfileSpec          `json:"spec"`
}

type ProfileList struct {
	unversioned.TypeMetadata
	Metadata ListMetadata  `json:"metadata"`
	Items []Profile        `json:"items"`
}


// ----  Rule (subtype of Profile and Policy)  ----

type Rule struct {
	Action string `json:"action" valid:"matches(deny|allow|next-tier)"`

	Protocol    *string `json:"protocol,omitempty" valid:"protocol"`
	SrcTag      *string `json:"srcTag,omitempty" valid:"optional"`
	SrcNet      *string `json:"srcNet,omitempty" valid:"cidr"`
	SrcSelector *string `json:"srcSelector,omitempty" valid:"selector"`
	SrcPorts    *[]int  `json:"srcPorts,omitempty" valid:"port"`
	DstTag      *string `json:"dstTag,omitempty" valid:"optional"`
	DstSelector *string `json:"dstSelector,omitempty" valid:"selector"`
	DstNet      *string `json:"dstNet,omitempty" valid:"cidr"`
	DstPorts    *[]int  `json:"dstPorts,omitempty" valid:"port"`
	IcmpType    *int    `json:"icmpType,omitempty" valid:"icmp_type"`
	IcmpCode    *int    `json:"icmpCode,omitempty" valid:"icmp_code"`

	NotProtocol    *string `json:"!protocol,omitempty" valid:"protocol"`
	NotSrcTag      *string `json:"!srcTag,omitempty" valid:"optional"`
	NotSrcNet      *string `json:"!srcNet,omitempty" valid:"cidr"`
	NotSrcSelector *string `json:"!srcSelector,omitempty" valid:"selector"`
	NotSrcPorts    *[]int  `json:"!srcPorts,omitempty" valid:"port"`
	NotDstTag      *string `json:"!dstTag,omitempty" valid:"optional"`
	NotDstSelector *string `json:"!dstSelector,omitempty" valid:"selector"`
	NotDstNet      *string `json:"!dstNet,omitempty" valid:"cidr"`
	NotDstPorts    *[]int  `json:"!dstPorts,omitempty" valid:"port"`
	NotIcmpType    *int    `json:"!icmpType,omitempty" valid:"icmp_type"`
	NotIcmpCode    *int    `json:"!icmpCode,omitempty" valid:"icmp_code"`
}

// ----  Host Endpoint  ----
// Kind: host-endpoint

type HostEndpointMetadata struct {
	ObjectMetadata
	Hostname string `json:"hostname" valid:"hostname"`
}

type HostEndpointSpec struct {
	InterfaceName     *string            `json:"interface_name" valid:"interface"`
	ExpectedIPv4Addrs *[]string          `json:"expectedIPv4Addrs" valid:"ipv4"`
	ExpectedIPv6Addrs *[]string          `json:"expectedIPv6Addrs" valid:"ipv6"` // Perhaps contract into a single field in the Northbound API
	Labels            *map[string]string `json:"labels" valid:"matches([a-zA-Z0-9-_/]*)"`
	Profiles        *[]string          `json:"profiles" valid:"profile"` // Perhaps profiles or profile_names
}

type HostEndpoint struct {
	unversioned.TypeMetadata
	Metadata HostEndpointMetadata  `json:"metadata"`
	Spec HostEndpointSpec  `json:"spec"`
}

type HostEndpointList struct {
	unversioned.TypeMetadata
	Metadata ListMetadata  `json:"metadata"`
	Items []HostEndpoint  `json:"items"`
}

