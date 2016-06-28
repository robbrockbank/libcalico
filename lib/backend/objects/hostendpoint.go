package objects

import (
	. "github.com/projectcalico/libcalico/lib/common"
	"fmt"
)

type HostEndpointKey struct {
	Hostname       string `json:"-" validate:"required,hostname"`
	EndpointID     string `json:"-" validate:"required,hostname"`
}

func (key HostEndpointKey) asEtcdKey() string {
	return fmt.Sprintf("/calico/v1/host/%s/endpoint/%s",
		key.Hostname, key.EndpointID)
}

type HostEndpoint struct {
	HostEndpointKey `json:"-"`
	Name              *string            `json:"name" validate:"omitempty,interface"`
	ExpectedIPv4Addrs *[]IP              `json:"expected_ipv4_addrs" validate:"omitempty,dive,ipv4"`
	ExpectedIPv6Addrs *[]IP              `json:"expected_ipv6_addrs" validate:"omitempty,dive,ipv6"`
	Labels            *map[string]string `json:"labels" validate:"omitempty,labels"`
	ProfileIDs        *[]string          `json:"profile_ids" validate:"omitempty,dive,name"`
}

type HostEndpointListOptions struct {
	Hostname   *string
	EndpointID *string
}
