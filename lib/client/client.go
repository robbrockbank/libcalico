package client

import (
	"io/ioutil"
	"encoding/json"
	"reflect"

	"github.com/ghodss/yaml"
	"github.com/kelseyhightower/envconfig"
	"github.com/projectcalico/libcalico/lib/api"
	"github.com/projectcalico/libcalico/lib/backend"
)

type Client struct {
	backend *backend.Client
}

type conversionHelper interface {
	convertAPIToBackend(interface{}) (interface{}, error)
	convertBackendToAPI(interface{}) (interface{}, error)
	convertMetadataToKeyInterface(interface{}) (backend.KeyInterface, error)
	convertMetadataToListInterface(interface{}) (backend.ListInterface, error)
}

// Return a new connected Client.
func New(config *api.ClientConfig) (c *Client, err error) {
	cc := Client{}
	cc.backend, err = backend.NewClient(config)
	return &cc, err
}

/*
func (c *Client) Tiers() TierInterface {
	return newTiers(c)
}

func (c *Client) Policies() PolicyInterface {
	return newPolicies(c)
}

func (c *Client) Profiles() ProfileInterface {
	return newProfiles(c)
}
*/

func (c *Client) HostEndpoints() HostEndpointInterface {
	return newHostEndpoints(c)
}

// Load the client config from the specified file (if specified) and from environment
// variables.  The values from both locations are merged together, with file values
// taking precedence).
func LoadClientConfig(f *string) (*api.ClientConfig, error) {
	var c api.ClientConfig

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

// Untyped interface for creating an API object.  This is called from the
// typed interface.  This assumes a 1:1 mapping between the API resource and
// the backend object.
func (c *Client) create(apiObject interface{}, helper conversionHelper) (interface{}, error) {
	// All API objects have a Metadata, so extract it.
	metadata := reflect.ValueOf(apiObject).FieldByName("Metadata").Interface()
	if k, err := helper.convertMetadataToKeyInterface(metadata); err != nil {
		return nil, err
	} else if b, err := helper.convertAPIToBackend(apiObject); err != nil {
		return nil, err
	} else if v, err := json.Marshal(b); err != nil {
		return nil, err
	} else {
		obj := backend.KeyValue{Key: k, Value: v}
		return apiObject, c.backend.Create(obj)
	}
}

// Untyped interface for updating an API object.  This is called from the
// typed interface.
func (c *Client) update(apiObject interface{}, helper conversionHelper) (interface{}, error) {
	// All API objects have a Metadata, so extract it.
	metadata := reflect.ValueOf(apiObject).FieldByName("Metadata").Interface()
	if k, err := helper.convertMetadataToKeyInterface(metadata); err != nil {
		return nil, err
	} else if b, err := helper.convertAPIToBackend(apiObject); err != nil {
		return nil, err
	} else if v, err := json.Marshal(b); err != nil {
		return nil, err
	} else {
		obj := backend.KeyValue{Key: k, Value: v}
		return apiObject, c.backend.Update(obj)
	}
}

// Untyped get interface for getting a single API object.  This is called from the typed
// interface.
func (c *Client) get(backendObject interface{}, metadata interface{}, helper conversionHelper) (interface{}, error) {
	if k, err := helper.convertMetadataToKeyInterface(metadata); err != nil {
		return nil, err
	} else if kv, err := c.backend.Get(k); err != nil {
		return nil, err
	} else if b, err := c.unmarshal(backendObject, kv); err != nil {
		return nil, err
        } else if a, err := helper.convertBackendToAPI(b); err != nil {
		return nil, err
	} else {
		return a, nil
	}
}

// Untyped get interface for deleting a single API object.  This is called from the typed
// interface.
func (c *Client) delete(metadata interface{}, helper conversionHelper) error {
	if k, err := helper.convertMetadataToKeyInterface(metadata); err != nil {
		return err
	} else if err := c.backend.Delete(k); err != nil {
		return err
	} else {
		return nil
	}
}

// Untyped get interface for getting a list of API objects.  This is called from the typed
// interface.
func (c *Client) list(backendObject interface{}, metadata interface{}, helper conversionHelper) ([]interface{}, error) {
	if l, err := helper.convertMetadataToListInterface(metadata); err != nil {
		return nil, err
	} else if kvs, err := c.backend.List(l); err != nil {
		return nil, err
	} else {
		as := make([]interface{}, len(kvs))
		for _, kv := range kvs {
			if b, err := c.unmarshal(backendObject, kv); err != nil {
				return nil, err
			} else if a, err := helper.convertBackendToAPI(b); err != nil {
				return nil, err
			} else {
				as = append(as, a)
			}
		}

		return as, nil
	}
}

func (c *Client) unmarshal(backendObject interface{}, v backend.KeyValue) (interface {}, error) {
	new := reflect.New(reflect.TypeOf(backendObject)).Interface()
	if err := json.Unmarshal(v.Value, new); err != nil {
		return nil, err
	} else {
		return new, nil
	}
}