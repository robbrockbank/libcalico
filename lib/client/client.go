package client

import (
	"io/ioutil"
	"github.com/kelseyhightower/envconfig"
	backend "github.com/projectcalico/libcalico/lib/backend/client"
	. "github.com/projectcalico/libcalico/lib/api"
	"github.com/ghodss/yaml"
)

type CalicoClient struct {
	backend *backend.Client
}

// Return a new connected CalicoClient.
func New(config *CalicoClientConfig) (c *CalicoClient, err error) {
	cc := CalicoClient{}
	cc.backend, err = backend.NewClient(config)
	return &cc, err
}

func (c *CalicoClient) Tiers(namespace string) TierInterface {
	return newTiers(c)
}

func (c *CalicoClient) Policies(namespace string) PolicyInterface {
	return newPolicies(c)
}

func (c *CalicoClient) Profiles(namespace string) ProfileInterface {
	return newProfiles(c)
}

func (c *CalicoClient) HostEndpoints(namespace string) HostEndpointInterface {
	return newHostEndpoints(c)
}

// Load the client config from the specified file (if specified) and from environment
// variables.  The values from both locations are merged together, with file values
// taking precedence).
//
// Returns a connected client.
func LoadClientConfig(f *string) (*CalicoClientConfig, error) {
	var c CalicoClientConfig

	// Load client config from environment variables first.
	if err := envconfig.Process("calico", &c); err != nil {
		return nil, err
	}

	// Override / merge with values loaded from the specified file.
	if f != nil {
		if b, err := ioutil.ReadFile(*f); err != nil {
			return nil, err
		} else if err := yaml.Unmarshal(b, &c); err != nil {
			return nil, err
		}
	}

	return &c, nil
}


