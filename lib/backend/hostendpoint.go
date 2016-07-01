package backend

import (
	"errors"
	"fmt"

	. "github.com/projectcalico/libcalico/lib/common"
)

type HostEndpointKey struct {
	Hostname   string `json:"-" validate:"required,hostname"`
	EndpointID string `json:"-" validate:"required,hostname"`
}

func (key HostEndpointKey) asEtcdKey() (string, error) {
	if key.Hostname == "" || key.EndpointID == "" {
		return "", errors.New("insufficient identifiers")
	}
	e := fmt.Sprintf("/calico/v1/host/%s/endpoint/%s",
		key.Hostname, key.EndpointID)
	return e, nil
}

func (key HostEndpointKey) asEtcdDeleteKey() (string, error) {
	return key.asEtcdKey()
}

type HostEndpointListOptions struct {
	Hostname   string
	EndpointID string
}

func (options HostEndpointListOptions) asEtcdKeyRegex() (string, error) {
	if options.Hostname == "" {
		return "", ErrorInsufficientIdentifiers{}
	}
	e := fmt.Sprintf("/calico/v1/host/%s/endpoint/%s",
		idOrWildcard(options.Hostname),
		idOrWildcard(options.EndpointID))
	return e, nil
}

type HostEndpoint struct {
	HostEndpointKey   `json:"-"`
	Name              string            `json:"name,omitempty" validate:"omitempty,interface"`
	ExpectedIPv4Addrs []IP              `json:"expected_ipv4_addrs,omitempty" validate:"omitempty,dive,ipv4"`
	ExpectedIPv6Addrs []IP              `json:"expected_ipv6_addrs,omitempty" validate:"omitempty,dive,ipv6"`
	Labels            map[string]string `json:"labels,omitempty" validate:"omitempty,labels"`
	ProfileIDs        []string          `json:"profile_ids,omitempty" validate:"omitempty,dive,name"`
}
