package client

import (
	etcd "github.com/coreos/etcd/client"
	"errors"
	"strings"
	"io/ioutil"
	"k8s.io/kubernetes/pkg/kubelet/client"
)


type CalicoClient struct {
	// Etcd connection information
	EtcdAuthority string `json:"etcdAuthority" envconfig:"ETCD_AUTHORITY" default:"127.0.0.1:2379"`
	EtcdEndpoints string `json:"etcdEndpoints" envconfig:"ETCD_ENDPOINTS"`
	EtcdUsername  string `json:"etcdUsername" envconfig:"ETCD_USERNAME"`
	EtcdPassword  string `json:"etcdPassword" envconfig:"ETCD_PASSWORD"`

	// ---- Internal package data ----
	connected bool
	etcdClient  *etcd.Client
	etcdKeysAPI *etcd.KeysAPI
}


// Connect() the client to the underlying datastore specified in the config.
func (cc *CalicoClient) Connect() (err error) {
	if cc.connected {
		panic("Client is already connected")
	}

	// Determine the location from the authority or the endpoints.  The endpoints
	// takes precedence if both are specified.
	etcdLocation := []string{}
	if cc.EtcdAuthority != "" {
		etcdLocation = []string{"http://" + cc.EtcdAuthority}
	}
	if cc.EtcdEndpoints != "" {
		etcdLocation = strings.Split(cc.EtcdEndpoints, ",")
	}

	if len(etcdLocation) == 0 {
		return errors.New("no etcd authority or endpoints specified")
	}

	// Create etcd client
	cfg := etcd.Config{
		Endpoints: etcdLocation,
		Transport: etcd.DefaultTransport}
	if cc.etcdClient, err = client.New(cfg); err != nil {
		return err
	}
	cc.etcdKeysAPI = client.NewKeysAPI(cc.etcdClient)
	cc.connected = true
	return nil
}


// Load the client config from the specified file (if specified) and from environment
// variables.  The values from both locations are merged together, with file alues
// taking precedence).
//
// Returns a connected client.
func New(f *string) (*CalicoClient, error) {
	var c CalicoClient

	// Load client config from environment variables first.
	if err := envconfig.Process("calico", &c); err != nil {
		return nil, err
	}

	// Override / merge with values loaded from the specified file.
	if f != nil {
		if b, err := ioutil.ReadFile(f); err != nil {
			return nil, err
		} else if err := yaml.Unmarshal(b, &c); err != nil {
			return nil, err
		}
	}

	return &c, c.Connect()
}
