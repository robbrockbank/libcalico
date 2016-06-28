package api

import (
	. "github.com/projectcalico/libcalico/lib/api/unversioned"
	. "github.com/projectcalico/libcalico/lib/common"
)

type HostEndpointMetadata struct {
	ObjectMetadata
	Hostname string `json:"hostname" valid:"hostname"`
	Labels        *map[string]string `json:"labels" validate:"omitempty,labels"`
}

type HostEndpointSpec struct {
	InterfaceName *string            `json:"interfaceName" validate:"omitempty,interface"`
	ExpectedIPs   *[]IP              `json:"expectedIPs" validate:"omitempty,dive,ip"`
	Profiles      *[]string          `json:"profiles" validate:"omitempty,dive,name"`
}

type HostEndpoint struct {
	TypeMetadata
	Metadata HostEndpointMetadata `json:"metadata"`
	Spec     HostEndpointSpec     `json:"spec"`
}

func NewHostEndpoint() *HostEndpoint {
	return &HostEndpoint{TypeMetadata: TypeMetadata{Kind: "hostEndpoint", APIVersion: "v1"}}
}

type HostEndpointList struct {
	TypeMetadata
	Metadata ListMetadata   `json:"metadata"`
	Items    []HostEndpoint `json:"items" validate:"dive"`
}

func NewHostEndpointList() *HostEndpointList {
	return &HostEndpointList{TypeMetadata: TypeMetadata{Kind: "hostEndpointList", APIVersion: "v1"}}
}