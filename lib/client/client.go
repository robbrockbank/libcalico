package client

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
	"github.com/kelseyhightower/envconfig"
	. "github.com/projectcalico/libcalico/lib/api"
	backend "github.com/projectcalico/libcalico/lib/backend/client"
)

type Client struct {
	backend *backend.Client
}

// Return a new connected Client.
func New(config *ClientConfig) (c *Client, err error) {
	cc := Client{}
	cc.backend, err = backend.NewClient(config)
	return &cc, err
}

func (c *Client) Tiers() TierInterface {
	return newTiers(c)
}

func (c *Client) Policies() PolicyInterface {
	return newPolicies(c)
}

func (c *Client) Profiles() ProfileInterface {
	return newProfiles(c)
}

func (c *Client) HostEndpoints() HostEndpointInterface {
	return newHostEndpoints(c)
}

// Load the client config from the specified file (if specified) and from environment
// variables.  The values from both locations are merged together, with file values
// taking precedence).
func LoadClientConfig(f *string) (*ClientConfig, error) {
	var c ClientConfig

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
