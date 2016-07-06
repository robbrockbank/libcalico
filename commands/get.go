package commands

import (
	"github.com/docopt/docopt-go"
	"github.com/projectcalico/libcalico/lib/api"
	"github.com/projectcalico/libcalico/lib/client"

	"fmt"

	"github.com/ghodss/yaml"
	"github.com/golang/glog"
	"github.com/projectcalico/libcalico/lib/api/unversioned"
)

func Get(args []string) error {
	doc := EtcdIntro + `Display one or many resources identified by file, stdin or resource type and name.

Possible resource types include: policy

By specifying the output as 'template' and providing a Go template as the value
of the --template flag, you can filter the attributes of the fetched resource(s).

Usage:
  calicoctl get ([--tier=<TIER>] [--hostname=<HOSTNAME>] (<KIND> [<NAME>]) | --filename=<FILENAME>) [--output=<OUTPUT>] [--config=<CONFIG>]

Examples:
  # List all policy in default output format.
  calicoctl get policy

  # List a specific policy in YAML format
  calicoctl get -o yaml policy my-policy-1

Options:
  -f --filename=<FILENAME>     Filename to use to get the resource.  If set to "-" loads from stdin.
  -o --output=<OUTPUT FORMAT>  Output format.  One of: yaml, json.  [Default: yaml]
  -t --tier=<TIER>             The policy tier.
  -n --hostname=<HOSTNAME>     The hostname.
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

	cmd := get{}
	results := executeConfigCommand(parsedArgs, cmd)
	glog.V(2).Infof("results: %v", results)

	if results.err != nil {
		fmt.Printf("Error getting resources: %v\n", results.err)
		return err
	}

	// TODO Handle better - results should be groups as per input file
	// For simplicity convert the returned list of resources to expand any lists
	resources := convertToSliceOfResources(results.resources)

	if output, err := yaml.Marshal(resources); err != nil {
		fmt.Printf("Error outputing data: %v", err)
	} else {
		fmt.Printf("%s", string(output))
	}

	return nil
}

// commandInterface for replace command.
// Maps the generic resource types to the typed client interface.
type get struct {
}

func (g get) execute(client *client.Client, resource unversioned.Resource) (unversioned.Resource, error) {
	var err error
	switch r := resource.(type) {
	case api.HostEndpoint:
		resource, err = client.HostEndpoints().List(r.Metadata)
	case api.Policy:
		resource, err = client.Policies().List(r.Metadata)
	case api.Profile:
		resource, err = client.Profiles().List(r.Metadata)
	case api.Tier:
		resource, err = client.Tiers().List(r.Metadata)
	default:
		panic(fmt.Errorf("Unhandled resource type: %v", resource))
	}

	return resource, err
}
