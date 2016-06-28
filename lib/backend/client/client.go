package client

import (
	"errors"
	"strings"

	etcd "github.com/coreos/etcd/client"
	. "github.com/projectcalico/libcalico/lib/api"
)

type Client struct {
	// Calico client config
	config *ClientConfig

	// ---- Internal package data ----
	connected   bool
	etcdClient  *etcd.Client
	etcdKeysAPI *etcd.KeysAPI
}

func NewClient(config *ClientConfig) (*Client, error) {
	c := Client{config: config}
	return &c, c.connect()
}

// Connect() the client to the underlying datastore specified in the config.
func (c *Client) connect() error {
	if c.connected {
		panic("Client is already connected")
	}

	// Determine the location from the authority or the endpoints.  The endpoints
	// takes precedence if both are specified.
	etcdLocation := []string{}
	if c.config.EtcdAuthority != "" {
		etcdLocation = []string{"http://" + c.config.EtcdAuthority}
	}
	if c.config.EtcdEndpoints != "" {
		etcdLocation = strings.Split(c.config.EtcdEndpoints, ",")
	}

	if len(etcdLocation) == 0 {
		return errors.New("no etcd authority or endpoints specified")
	}

	// Create the etcd client
	cfg := etcd.Config{
		Endpoints: etcdLocation,
		Transport: etcd.DefaultTransport}
	client, err := etcd.New(cfg)
	if err != nil {
		return err
	}
	keys := etcd.NewKeysAPI(client)
	c.etcdClient = &client
	c.etcdKeysAPI = &keys
	c.connected = true
	return nil
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
