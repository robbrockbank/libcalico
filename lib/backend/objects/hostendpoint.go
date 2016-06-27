package objects

import (
	. "github.com/projectcalico/libcalico/lib/common"
)

type HostEndpoint struct {
	InterfaceName *string            `json:"interfaceName" validate:"omitempty,interface"`
	ExpectedIPs   *[]IP              `json:"expectedIPs" validate:"omitempty,dive,ip"`
	Labels        *map[string]string `json:"labels" validate:"omitempty,labels"`
	Profiles      *[]string          `json:"profiles" validate:"omitempty,dive,name"`
}