package libcalico

import (
	"io/ioutil"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/coreos/etcd/client"
	"github.com/ghodss/yaml"
	"github.com/projectcalico/libcalico/lib/api"
	"github.com/projectcalico/libcalico/lib/api/unversioned"
	"fmt"
)

// Save the resource in the datastore:
// -  can_create is used to indicate whether the resource can be created if it does not exist
// -  can_replace is used to indicate whether the resource can be updated if it already exists
//
// The Resource may be of Kind "list" and therefore contain zero or more resources to save.
// If a list of resources is specified, they are saved in the list order, and this function
// returns after saving all resources, or after hitting an error.  The function returns the
// number of resources successfully updated.
func SaveResource(etcd client.KeysAPI, r *unversioned.Resource, canCreate, canReplace bool) (int, error) {
	return 0, nil
}

// Load the resource(s) from the datastore:
// -  The Resource does not need to contain the Spec - if specified it is ignored.
// -  If the supplied Resource metadata contains missing identifiers, the query will wildcard
//    those identifiers.  If multiple values are returned, the returned resource will be of
//    kind "list".
//
// The Resource may be of Kind "list" and therefore contain zero or more resources to save.
// If a list of resources is specified, they are saved in the list order, and this function
// returns after saving all resources, or after hitting an error.  The function returns the
// number of resources successfully updated.
func LoadResource(etcd client.KeysAPI, r *unversioned.Resource) (*unversioned.Resource, error) {
	return nil, nil
}

// Delete the resource(s) from the datastore:
// -  The Resource does not need to contain the Spec - if specified it is ignored.
// -  The Resource metadata should contain all identifiers required to uniquely identify a
//    single resource.
// -  The ignore_not_present flag indicates whether attempts to delete a missing resource is
//    ignored, or treated as an error.
//
// The Resource may be of Kind "list" and therefore contain zero or more resources to delete.
// If a list of resources is specified, they are deleted in the list order, and this function
// returns after saving all resources, or after hitting an error.  The function returns the
// number of resources successfully deleted.
func DeleteResource(etcd client.KeysAPI, r *unversioned.Resource, ignoreNotPresent bool) (int, error) {
	return 0, nil
}

// Create the Resource from the specified file f.
// -  The file format may be JSON or YAML encoding of either a single resource or list of
//    resources as defined by the API objects in /api.
// -  A filename of "-" means "Read from stdin".
//
// The returned Resource will either be a single Resource or a List containing zero or more
// Resources.  If the file does not contain any valid Resources this function returns an error.
func CreateResourceFromFile(f string) (*unversioned.Resource, error) {

	// Load the bytes from file or from stdin.
	var b []byte
	var err error

	if f == "-" {
		b, err = ioutil.ReadAll(os.Stdin)
	} else {
		b, err = ioutil.ReadFile(f)
	}
	if err != nil {
		return nil, err
	}

	return CreateResourceFromBytes(b)
}

// Create the Resource from the specified byte array encapsulating the resource.
// -  The byte array may be JSON or YAML encoding of either a single resource or list of
//    resources as defined by the API objects in /api.
//
// The returned Resource will either be a single Resource or a List containing zero or more
// Resources.  If the file does not contain any valid Resources this function returns an error.
func CreateResourceFromBytes(b []byte) (*unversioned.Resource, error) {
	// Start by unmarshalling the bytes into a TypeMetadata structure - this will ignore
	// other fields.
	tm := unversioned.TypeMetadata{}
	err := yaml.Unmarshal(b, &tm)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Parsed type metadata: %v\n", tm)

	// Handle unversioned resources explicitly (currently just the list type).
	var rp *unversioned.Resource
	if tm.Kind == "list" {
		fmt.Printf("Processing list type")
		r := unversioned.ResourceList(&unversioned.ListMetadata{}, &unversioned.ListSpec{})
		err = yaml.Unmarshal(b, &r)
		if err != nil {
			return nil, err
		}
		ls := r.Spec.(unversioned.ListSpec)
		rl := []unversioned.Resource{}

		// The resource list spec and meta data will be parsed into generic interfaces, so
		// need to re-parse based on the type metadata for each.
		for _, ri := range ls.List {
			rib, err := yaml.Marshal(ri)
			if err != nil {
				return nil, err
			}
			ri, err := CreateResourceFromBytes(rib)
			if err != nil {
				return nil, err
			}
			rl = append(rl, *ri)
		}

		// Update the Resource List to be a list of concrete list types.  This allows
		// the list to be validated.
		ls.List = rl
		rp = &r
	} else {
		// Now that we have a concrete type unmarshal into that resource type.
		fmt.Printf("Processing type %s\n", tm.Kind)
		r, err := api.CreateResourceManager().NewResource(tm)
		err = yaml.Unmarshal(b, &r)
		if err != nil {
			return nil, err
		}
		rp = r
	}

	fmt.Printf("Parsed: %v\n", *rp)

	// Validate the data in the structures.
	_, err = govalidator.ValidateStruct(r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
