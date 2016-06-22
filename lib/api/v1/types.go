package v1

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/projectcalico/libcalico/lib/api/unversioned"
)

/*
The v1 structure definitions used in the V1 interface.
*/

// ----  Tier  ----
// Kind:  tier

type TierMetadata struct {
	Name string `json:"name" valid:"name"`
}

type TierSpec struct {
	Order Order `json:"order" valid:"matches(default|\d*)"`
}

func ResourceTier(m *TierMetadata, s *TierSpec) unversioned.Resource {
	return unversioned.Resource{unversioned.TypeMetadata{"tier", "v1"}, m, s}
}

// ----  Policy  ----
// Kind: policy

type PolicyMetadata struct {
	Name string `json:"name" valid:"matches([a-zA-Z0-9-_]+)"`
	Tier string `json:"tier,omitempty" valid:"matches([a-zA-Z0-9-_]+),optional"`
}

type PolicySpec struct {
	Order         Order  `json:"order" valid:"matches(default|\d*)"`
	InboundRules  []Rule `json:"ingress"`
	OutboundRules []Rule `json:"egress"`
	Selector      string `json:"selector"`
}

func ResourcePolicy(m *PolicyMetadata, s *PolicySpec) unversioned.Resource {
	return unversioned.Resource{unversioned.TypeMetadata{"policy", "v1"}, m, s}
}

// ----  Profile  ----
// Kind: profile

type ProfileMetadata struct {
	Name string `json:"name" valid:"name"`
}

type ProfileSpec struct {
	InboundRules  *[]Rule            `json:"ingress,omitempty" valid:"optional"`
	OutboundRules *[]Rule            `json:"egress,omitempty" valid:"optional"`
	Labels        *map[string]string `json:"labels,omitempty" valid:"matches([a-zA-Z0-9-_/]+),optional"`
	Tags          *[]string          `json:"tags,omitempty" valid:"optional"`
}

func ResourceProfile(m *ProfileMetadata, s *ProfileSpec) unversioned.Resource {
	return unversioned.Resource{unversioned.TypeMetadata{"profile", "v1"}, m, s}
}

// ----  Rule (subtype of Profile and Policy)  ----

type Rule struct {
	Action string `json:"action" valid:"matches(deny|allow|next-tier)"`

	Protocol    *string `json:"protocol,omitempty" valid:"protocol,optional"`
	SrcTag      *string `json:"src_tag,omitempty" valid:"optional"`
	SrcNet      *string `json:"src_net,omitempty" valid:"cidr,optional"`
	SrcSelector *string `json:"src_selector,omitempty" valid:"selector,optional"`
	SrcPorts    *[]int  `json:"src_ports,omitempty" valid:"port,optional"`
	DstTag      *string `json:"dst_tag,omitempty" valid:"optional"`
	DstSelector *string `json:"dst_selector,omitempty" valid:"selector,optional"`
	DstNet      *string `json:"dst_net,omitempty" valid:"cidr,optional"`
	DstPorts    *[]int  `json:"dst_ports,omitempty" valid:"port,optional"`
	IcmpType    *int    `json:"icmp_type,omitempty" valid:"icmp_type,optional"`
	IcmpCode    *int    `json:"icmp_code,omitempty" valid:"icmp_code,optional"`

	NotProtocol    *string `json:"!protocol,omitempty" valid:"protocol,optional"`
	NotSrcTag      *string `json:"!src_tag,omitempty" valid:"optional"`
	NotSrcNet      *string `json:"!src_net,omitempty" valid:"cidr,optional"`
	NotSrcSelector *string `json:"!src_selector,omitempty" valid:"selector,optional"`
	NotSrcPorts    *[]int  `json:"!src_ports,omitempty" valid:"port,optional"`
	NotDstTag      *string `json:"!dst_tag,omitempty" valid:"optional"`
	NotDstSelector *string `json:"!dst_selector,omitempty" valid:"selector,optional"`
	NotDstNet      *string `json:"!dst_net,omitempty" valid:"cidr,optional"`
	NotDstPorts    *[]int  `json:"!dst_ports,omitempty" valid:"port,optional"`
	NotIcmpType    *int    `json:"!icmp_type,omitempty" valid:"icmp_type,optional"`
	NotIcmpCode    *int    `json:"!icmp_code,omitempty" valid:"icmp_code,optional"`
}

// ----  Host Endpoint  ----
// Kind: host-endpoint

type HostEndpointMetadata struct {
	Hostname string `json:"hostname" valid:"hostname"`
	Name     string `json:"name" valid:"name"`
}

type HostEndpointSpec struct {
	InterfaceName     *string            `json:"interface_name" valid:"interface,optional"`
	ExpectedIPv4Addrs *[]string          `json:"egress" valid:"ipv4,optional"`
	ExpectedIPv6Addrs *[]string          `json:"egress" valid:"ipv6,optional"` // Perhaps contract into a single field in the Northbound API
	Labels            *map[string]string `json:"labels" valid:"matches([a-zA-Z0-9-_/]*),optional"`
	ProfileIDs        *[]string          `json:"profile_ids" valid:"profile,optional"` // Perhaps profiles or profile_names
}

func ResourceHostEndpoint(m *HostEndpointMetadata, s *HostEndpointSpec) unversioned.Resource {
	return unversioned.Resource{unversioned.TypeMetadata{"host-endpoint", "v1"}, m, s}
}

// Order is a type that can hold an int or a value indicating "default".  When used in
// JSON or YAML marshalling and unmarshalling, it produces or consumes the
// inner type.  This allows you to have, for example, a JSON field that can
// accept a number or the dtring value "default".
type Order struct {
	Kind   OrderKind
	IntVal int
}

// IntstrKind represents the stored type of IntOrString.
type OrderKind int

const (
	OrderInt     OrderKind = iota // The Order holds an int.
	OrderDefault                  // The Order holds "default".
)

// UnmarshalJSON implements the json.Unmarshaller interface.
func (order *Order) UnmarshalJSON(value []byte) error {
	if value[0] == '"' {
		var s string
		err := json.Unmarshal(value, &s)
		if err != nil {
			return err
		}
		if s != "default" {
			return fmt.Errorf("order is not an integer or default")
		}
		order.Kind = OrderDefault
		return nil
	}
	order.Kind = OrderInt
	return json.Unmarshal(value, &order.IntVal)
}

// String returns the string value, or Itoa's the int value.
func (order *Order) String() string {
	if order.Kind == OrderDefault {
		return "default"
	}
	return strconv.Itoa(order.IntVal)
}

// MarshalJSON implements the json.Marshaller interface.
func (order Order) MarshalJSON() ([]byte, error) {
	switch order.Kind {
	case OrderInt:
		return json.Marshal(order.IntVal)
	case OrderDefault:
		return []byte("default"), nil
	default:
		return []byte{}, fmt.Errorf("impossible Order.Kind")
	}
}
