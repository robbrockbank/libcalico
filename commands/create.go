package commands

import (
	"fmt"

	"github.com/docopt/docopt-go"
	"github.com/golang/glog"
	"github.com/projectcalico/libcalico/lib/api"
	"github.com/projectcalico/libcalico/lib/api/unversioned"
	"github.com/projectcalico/libcalico/lib/client"
	"github.com/projectcalico/libcalico/lib/common"
)

func Create(args []string) error {
	doc := EtcdIntro + `Create a resource by filename or stdin.

Usage:
  calicoctl create --filename=<FILENAME> [--skip-exists] [--config=<CONFIG>]

Examples:
  # Create a policy using the data in policy.yaml.
  calicoctl create -f ./policy.yaml

  # Create a policy based on the JSON passed into stdin.
  cat policy.json | calicoctl create -f -

Options:
  -f --filename=<FILENAME>     Filename to use to create the resource.  If set to "-" loads from stdin.
  -s --skip-exists             Skip over and treat as successful any attempts to create an entry that
                               already exists.
  -c --config=<CONFIG>         Filename containing connection configuration in YAML or JSON format.
                               [default: /etc/calico/calicoctl.cfg]
`
	parsedArgs, err := docopt.Parse(doc, args, true, "calicoctl", false, false)
	if err != nil {
		return err
	}
	if len(parsedArgs) == 0 {
		return nil
	}

	cmd := create{skipIfExists: parsedArgs["--skip-exists"].(bool)}
	results := executeConfigCommand(parsedArgs, cmd)
	glog.V(2).Infof("results: %v", results)

	if results.fileInvalid {
		fmt.Printf("Error processing input file: %v\n", results.err)
	} else if results.numHandled == 0 {
		if results.numResources == 0 {
			fmt.Printf("No resources specified in file\n")
		} else if results.numResources == 1 {
			fmt.Printf("Failed to create '%s' resource: %v\n", results.singleKind, results.err)
		} else if results.singleKind != "" {
			fmt.Printf("Failed to create any '%s' resources: %v\n", results.singleKind, results.err)
		} else {
			fmt.Printf("Failed to create any resources: %v\n", results.err)
		}
	} else if results.err == nil {
		if results.singleKind != "" {
			fmt.Printf("Successfully created %d '%s' resource(s)\n", results.numHandled, results.singleKind)
		} else {
			fmt.Printf("Successfully created %d resource(s)\n", results.numHandled)
		}
	} else {
		fmt.Printf("Partial success: ")
		if results.singleKind != "" {
			fmt.Printf("created the first %d out of %d '%s' resources:\n",
				results.numHandled, results.numResources, results.singleKind)
		} else {
			fmt.Printf("created the first %d out of %d resources:\n",
				results.numHandled, results.numResources)
		}
		fmt.Printf("Hit error: %v\n", results.err)
	}

	return results.err
}

// commandInterface for create command.
// Maps the generic resource types to the typed client interface.
type create struct {
	skipIfExists bool
}

func (c create) execute(client *client.Client, resource unversioned.Resource) (unversioned.Resource, error) {
	var err error
	switch r := resource.(type) {
	case api.HostEndpoint:
		_, err = client.HostEndpoints().Create(&r)
	case api.Policy:
		_, err = client.Policies().Create(&r)
	case api.Profile:
		_, err = client.Profiles().Create(&r)
	case api.Tier:
		_, err = client.Tiers().Create(&r)
	default:
		panic(fmt.Errorf("Unhandled resource type: %v", resource))
	}

	if err == nil {
		return resource, nil
	}

	// Handle resource does not exist errors explicitly.
	switch err.(type) {
	case common.ErrorResourceAlreadyExists:
		if c.skipIfExists {
			return resource, nil
		}
	}
	return nil, err
}
