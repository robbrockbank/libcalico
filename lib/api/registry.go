package api

import (
	"errors"

	. "github.com/projectcalico/libcalico/lib/api/unversioned"
	"github.com/projectcalico/libcalico/lib/api/v1"
	"fmt"
	"reflect"
)

type ResourceManager struct {
	ResourceHelper map[TypeMetadata]ResourceHelper
}

func NewResourceManager() *ResourceManager {
	return &ResourceManager{make(map[TypeMetadata]ResourceHelper)}
}

type ResourceHelper struct {
	Type     TypeMetadata
	ResourceType     interface{}
	ResourceListType     interface{}
}

func (rm *ResourceManager) registerHelper(r ResourceHelper) {
	rm.ResourceHelper[r.Type] = r
}

func CreateResourceManager() *ResourceManager {
	rm := NewResourceManager()
	rm.registerHelper(ResourceHelper{
		TypeMetadata{Kind: "tier", APIVersion: "v1"},
		v1.Tier{},
		v1.TierList{},
	})
	rm.registerHelper(ResourceHelper{
		TypeMetadata{Kind: "policy", APIVersion: "v1"},
		v1.Policy{},
		v1.PolicyList{},
	})
	rm.registerHelper(ResourceHelper{
		TypeMetadata{Kind: "profile", APIVersion: "v1"},
		v1.Profile{},
		v1.ProfileList{},
	})
	rm.registerHelper(ResourceHelper{
		TypeMetadata{Kind: "hostEndpoint", APIVersion: "v1"},
		v1.HostEndpoint{},
		v1.HostEndpointList{},
	})

	return rm
}

func (rm *ResourceManager) NewResource(tm TypeMetadata) (interface{}, error) {
	rh, ok := rm.ResourceHelper[tm]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Unknown resource type (%s) and/or version (%s)", tm.Kind, tm.APIVersion))
	}
	fmt.Printf("Found resource helper: %v\n", rh)
	new := reflect.New(rh.ResourceType)
	return &new, nil
}

func (rm *ResourceManager) NewResourceList(tm unversioned.TypeMetadata) (*unversioned.Resource, error) {
	rh, ok := rm.ResourceHelper[tm]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Unknown resource type (%s) and/or version (%s)", tm.Kind, tm.APIVersion))
	}
	fmt.Printf("Found resource helper: %v\n", rh)
	fmt.Printf("Returning empty resource: %v\n", new)
	new := reflect.New(rh.ResourceListType)
	return &new, nil
}
