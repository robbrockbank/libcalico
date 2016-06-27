package api

import (
	"errors"

	"fmt"
	"reflect"
	"strings"

	. "github.com/projectcalico/libcalico/lib/api/unversioned"
)

var helpers map[TypeMetadata]ResourceHelper

// Register all of the available resource types.
func init() {
	helpers = make(map[TypeMetadata]ResourceHelper)

	// Register all known resources.
	// TODO Would be better for each version of the API to do this?
	registerHelper(ResourceHelper{
		NewTier(),
		NewTierList(),
	})
	registerHelper(ResourceHelper{
		NewPolicy(),
		NewPolicyList(),
	})
	registerHelper(ResourceHelper{
		NewProfile(),
		NewProfileList(),
	})
	registerHelper(ResourceHelper{
		NewHostEndpoint(),
		NewHostEndpointList(),
	})
}


// Register a resource helper.
func registerHelper(r ResourceHelper) {
	tmd := reflect.ValueOf(r.ResourceType).Elem().FieldByName("TypeMetadata").Interface().(TypeMetadata)
	helpers[tmd] = r
}


// ResourceHelper encapsulates details about a specific version of a specific resource:
// -  The type of resource (Kind and Version)
// -  The concrete resource struct for this version
// -  The concrete resource list struct for this version
type ResourceHelper struct {
	ResourceType     interface{}
	ResourceListType interface{}
}


// Create a new concrete resource structure based on the type.  If the type is
// a list, this creates a concrete Resource-List of the required type.
func NewResource(tm TypeMetadata) (interface{}, error) {
	// If this is a list, farm out to NewResourceList passing in the TypeMetadata
	// for the actual resource type.
	if strings.HasSuffix(tm.Kind, "List") {
		tm = TypeMetadata{
			Kind: strings.TrimSuffix(tm.Kind, "List"),
			APIVersion: tm.APIVersion,
		}
		return NewResourceList(tm)
	}

	rh, ok := helpers[tm]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Unknown resource type (%s) and/or version (%s)", tm.Kind, tm.APIVersion))
	}
	fmt.Printf("Found resource helper: %v\n", rh)
	fmt.Printf("Type: %v\n", reflect.TypeOf(rh.ResourceType))
	new := reflect.New(reflect.TypeOf(rh.ResourceType)).Interface()
	return new, nil
}

func NewResourceList(tm TypeMetadata) (interface{}, error) {
	rh, ok := helpers[tm]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Unknown resource type (%s) and/or version (%s)", tm.Kind, tm.APIVersion))
	}
	fmt.Printf("Found resource helper: %v\n", rh)
	new := reflect.New(reflect.TypeOf(rh.ResourceListType)).Interface()
	reflect.ValueOf(new).Elem().FieldByName("Kind").SetString(tm.Kind + "List")
	reflect.ValueOf(new).Elem().FieldByName("APIVersion").SetString(tm.APIVersion)
	return new, nil
}
