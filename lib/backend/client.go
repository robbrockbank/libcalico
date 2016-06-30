package backend

import (
	"errors"
	"strings"

	etcd "github.com/coreos/etcd/client"
	"github.com/golang/glog"
	"github.com/projectcalico/libcalico/lib/api"
	"github.com/projectcalico/libcalico/lib/common"
	"golang.org/x/net/context"
)

// Interface used to calculate a datastore key.
type KeyInterface interface {
	asEtcdKey() (string, error)
	asEtcdDeleteKey() (string, error)
}

// Interface used to perform datastore lookups.
type ListInterface interface {
	asEtcdKeyRegex() (string, error)
}

// Encapsulated datastore key interface with value.
type KeyValue struct {
	Key   KeyInterface
	Value []byte
}

// Backend client data
type Client struct {
	// Calico client config
	config *api.ClientConfig

	// ---- Internal package data ----
	connected   bool
	etcdClient  etcd.Client
	etcdKeysAPI etcd.KeysAPI
}

// NewClient creates a new backend datastore client.
func NewClient(config *api.ClientConfig) (*Client, error) {
	c := Client{config: config}
	return &c, c.connectEtcd()
}

// Connect the client to the etcd datastore specified in the config.
func (c *Client) connectEtcd() error {
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
	c.etcdClient = client
	c.etcdKeysAPI = keys
	c.connected = true
	return nil
}

// Create an entry in the datastore.  This errors if the entry already exists.
func (c *Client) Create(d KeyValue) error {
	key, _ := d.Key.asEtcdKey()
	glog.V(2).Infof("Create Key: %s\n", key)
	glog.V(2).Infof("Value: %s\n", d.Value)
	_, err := c.etcdKeysAPI.Create(context.Background(), key, string(d.Value))
	return convertError(err, key)
}

// Update an existing entry in the datastore.  This errors if the entry does
// not exist.
func (c *Client) Update(d KeyValue) error {
	key, _ := d.Key.asEtcdKey()
	glog.V(2).Infof("Update Key: %s\n", key)
	glog.V(2).Infof("Value: %s\n", d.Value)
	_, err := c.etcdKeysAPI.Update(context.Background(), key, string(d.Value))
	return convertError(err, key)
}

// Delete an entry in the datastore.  This errors if the entry does not exists.
func (c *Client) Delete(k KeyInterface) error {
	key, _ := k.asEtcdDeleteKey()
	glog.V(2).Infof("Delete Key: %s\n", key)
	_, err := c.etcdKeysAPI.Delete(context.Background(), key, &etcd.DeleteOptions{Recursive: true})
	return convertError(err, key)
}

// Get and entry from the datastore.  This errors if the entry does not exist.
func (c *Client) Get(k KeyInterface) (KeyValue, error) {
	key, _ := k.asEtcdKey()
	glog.V(2).Infof("Get Key: %s\n", key)
	return KeyValue{}, nil
}

// List entries in the datastore.  This may return an empty list of there are
// no entries matching the request in the ListInterface.
func (c *Client) List(l ListInterface) ([]KeyValue, error) {
	key, _ := l.asEtcdKeyRegex()
	glog.V(2).Infof("List Key: %s\n", key)
	return []KeyValue{}, nil
}

func convertError(err error, key string) error {
	if err == nil {
		glog.V(2).Info("Comand completed without error")
		return nil
	}

	switch err.(type) {
	case etcd.Error:
		switch err.(etcd.Error).Code {
		case etcd.ErrorCodeNodeExist:
			glog.V(2).Info("Node exists error")
			return common.ErrorResourceAlreadyExists{Err: err, Name: key}
		case etcd.ErrorCodeKeyNotFound:
			glog.V(2).Info("Key not found error")
			return common.ErrorResourceDoesNotExist{Err: err, Name: key}
		case etcd.ErrorCodeUnauthorized:
			glog.V(2).Info("Unauthorized error")
			return common.ErrorConnectionUnauthorized{Err: err}
		default:
			glog.V(2).Infof("Generic etcd error error: %v", err)
			return common.ErrorDatastoreError{Err: err}
		}
	default:
		glog.V(2).Infof("Unhandled error: %v", err)
		return common.ErrorDatastoreError{Err: err}
	}
}
