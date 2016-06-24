types.gopackage v1

import (
	. "github.com/projectcalico/libcalico/lib/api/unversioned"
	. "github.com/projectcalico/libcalico/lib/common"
)

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
	TypeMetadata
	Metadata HostEndpointMetadata `json:"metadata"`
	Spec     HostEndpointSpec     `json:"spec"`
}

type HostEndpointList struct {
	TypeMetadata
	Metadata ListMetadata   `json:"metadata"`
	Items    []HostEndpoint `json:"items" validate:"dive"`
}
