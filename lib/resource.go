package libcalico

import (
	"io/ioutil"
	"os"

	"fmt"
	"reflect"

	"github.com/ghodss/yaml"
	"github.com/projectcalico/libcalico/lib/api"
	. "github.com/projectcalico/libcalico/lib/api/unversioned"
	"github.com/projectcalico/libcalico/lib/common"
)

// Create the Resource from the specified file f.
// -  The file format may be JSON or YAML encoding of either a single resource or list of
//    resources as defined by the API objects in /api.
// -  A filename of "-" means "Read from stdin".
//
// The returned Resource will either be a single Resource or a List containing zero or more
// Resources.  If the file does not contain any valid Resources this function returns an error.
func CreateResourceFromFile(f string) (interface{}, error) {

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

// Create the resource from the specified byte array encapsulating the resource.
// -  The byte array may be JSON or YAML encoding of either a single resource or list of
//    resources as defined by the API objects in /api.
//
// The returned Resource will either be a single resource document or a List of documents.
// If the file does not contain any valid Resources this function returns an error.
func CreateResourceFromBytes(b []byte) (interface{}, error) {
	// Start by unmarshalling the bytes into a TypeMetadata structure - this will ignore
	// other fields.
	var err error
	tm := TypeMetadata{}
	tml := []TypeMetadata{}
	if err = yaml.Unmarshal(b, &tm); err == nil {
		// We processed a metadata, so create a concrete resource struct to unpack
		// into.
		return unmarshalResource(tm, b)
	} else if err = yaml.Unmarshal(b, &tml); err == nil {
		// We processed a list of metadata's, create a list of concrete resource
		// structs to unpack into.
		return unmarshalListOfResources(tml, b)
	} else {
		// Failed to parse a single resource or list of resources.
		return nil, err
	}
}

// Unmarshal a bytearray containing a single resource of the specified type into
// a concrete structure for that resource type.
func unmarshalResource(tm TypeMetadata, b []byte) (interface{}, error) {
	fmt.Printf("Processing type %s\n", tm.Kind)
	unpacked, err := api.NewResource(tm)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(b, unpacked); err != nil {
		return nil, err
	}

	fmt.Printf("Type of unpacked data: %v\n", reflect.TypeOf(unpacked))
	if err = common.Validate(unpacked); err != nil {
		return nil, err
	}

	fmt.Printf("Unpacked: %v\n", unpacked)

	return unpacked, nil
}

// Unmarshal a bytearray containing a list of resources of the specified type into
// a list of concrete structures for that resource type.
func unmarshalListOfResources(tml []TypeMetadata, b []byte) (interface{}, error) {
	fmt.Printf("Processing list of resources\n")
	unpacked := []interface{}{}
	for _, tm := range tml {
		fmt.Printf("  - processing type %s\n", tm.Kind)
		r, err := api.NewResource(tm)
		if err != nil {
			return nil, err
		}
		unpacked = append(unpacked, r)
	}

	if err := yaml.Unmarshal(b, &unpacked); err != nil {
		return nil, err
	}

	// Validate the data in the structures.  The validator does not handle slices, so
	// validate each resource separately.
	for _, r := range unpacked {
		if err := common.Validate(r); err != nil {
			return nil, err
		}
	}

	fmt.Printf("Unpacked: %v\n", unpacked)

	return unpacked, nil
}
