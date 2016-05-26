package libcalico

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/coreos/etcd/client"
	"github.com/ghodss/yaml"
	"golang.org/x/net/context"
)

var policyRE = regexp.MustCompile(`/calico/v1/policy/tier/[^/]*/policy/[^/]*`)

type TierResource struct {
	Kind     string     `json:"kind"`
	Version  string     `json:"version"`
	Metadata PolicyMetadata `json:"metadata"`
	Spec     PolicySpec `json:"spec"`
}

type TierMetadata struct {
	Name string `json:"name"`
	Tier string `json:"tier,omitempty"`
}

type TierSpec struct {
	Order int `json:"order"`
}

type PolicyResource struct {
	Kind     string     `json:"kind"`
	Version  string     `json:"version"`
	Metadata PolicyMetadata `json:"metadata"`
	Spec     PolicySpec `json:"spec"`
}

type PolicyMetadata struct {
	Name string `json:"name"`
	Tier string `json:"tier,omitempty"`
}

type PolicySpec struct {
	Order         int    `json:"order"`
	InboundRules  []Rule `json:"inbound_rules"`
	OutboundRules []Rule `json:"outbound_rules"`
	Selector string `json:"selector"`
}

type Rule struct {
	Action string `json:"action"`

	Protocol    string `json:"protocol,omitempty"`
	SrcTag      string `json:"src_tag,omitempty"`
	SrcNet      string `json:"src_net,omitempty"`
	SrcSelector string `json:"src_selector,omitempty"`
	SrcPorts    []int  `json:"src_ports,omitempty"`
	DstTag      string `json:"dst_tag,omitempty"`
	DstSelector string `json:"dst_selector,omitempty"`
	DstNet      string `json:"dst_net,omitempty"`
	DstPorts    []int  `json:"dst_ports,omitempty"`
	IcmpType    int    `json:"icmp_type,omitempty"`
	IcmpCode    int    `json:"icmp_code,omitempty"`

	NotProtocol    string `json:"!protocol,omitempty"`
	NotSrcTag      string `json:"!src_tag,omitempty"`
	NotSrcNet      string `json:"!src_net,omitempty"`
	NotSrcSelector string `json:"!src_selector,omitempty"`
	NotSrcPorts    []int  `json:"!src_ports,omitempty"`
	NotDstTag      string `json:"!dst_tag,omitempty"`
	NotDstSelector string `json:"!dst_selector,omitempty"`
	NotDstNet      string `json:"!dst_net,omitempty"`
	NotDstPorts    []int  `json:"!dst_ports,omitempty"`
	NotIcmpType    int    `json:"!icmp_type,omitempty"`
	NotIcmpCode    int    `json:"!icmp_code,omitempty"`
}

func LoadPolicy(policyBytes []byte) (*PolicyResource, error) {
	var pq PolicyResource
	var err error

	// Load the policy string.  This should be a fully qualified set of policy.
	err = yaml.Unmarshal(policyBytes, &pq)
	if err != nil {
		return nil, err
	}

	if pq.Kind != "policy" {
		return nil, errors.New(fmt.Sprintf("Expecting kind 'policy', but got '%s'.", pq.Kind))
	}
	if pq.Version != "v1" {
		return nil, errors.New(fmt.Sprintf("Expecting version 'v1', but got '%s'.", pq.Version))
	}

	return &pq, nil
}

func CreateOrReplacePolicy(etcd client.KeysAPI, pq *PolicyResource, replace bool) error {

	var err error

	// If the default tier is specified, or no tier is specified we may need to create the default.
	tierName := pq.Metadata.Tier
	if tierName == "default" || tierName == "" {
		err = createDefaultTier(etcd)
		tierName = "default"
	}

	// Construct the policy key, and marshal the policy spec into JSON format
	// required by Felix.
	pk := fmt.Sprintf("/calico/v1/policy/tier/%s/policy/%s", tierName, pq.Metadata.Name)
	pb, err := json.Marshal(pq.Spec)
	if err != nil {
		return err
	}

	// Write the policy object to etcd.  If replacing policy, the we expect the policy to
	// already exist, otherwise we expect it to not exist.
	if replace {
		_, err = etcd.Update(context.Background(), pk, string(pb))
	} else {
		_, err = etcd.Create(context.Background(), pk, string(pb))
	}
	return err
}

func GetPolicies(etcd client.KeysAPI, tierName string) ([]PolicyResource, error) {
	var pqs []PolicyResource

	actualTierName := tierName
	if actualTierName == "" {
		actualTierName = "default"
	}

	resp, err := etcd.Get(context.Background(), fmt.Sprintf("/calico/v1/policy/tier/%s/policy", actualTierName), &client.GetOptions{Recursive: true})
	if err != nil {
		if !client.IsKeyNotFound(err) {
			return nil, err
		}
		return pqs, nil
	}

	for _, node := range resp.Node.Nodes {
		var ps PolicySpec

		var re = regexp.MustCompile(`/calico/v1/policy/tier/([^/]+?)/policy/([^/]+?)`)
		matches := re.FindStringSubmatch(node.Key)
		if matches != nil {
			policyName := matches[1]

			err = json.Unmarshal([]byte(node.Value), &ps)
			if err != nil {
				log.Fatal(err)
			}
			pm := PolicyMetadata{Name: policyName, Tier: tierName}
			pq := PolicyResource{
				Kind:     "policy",
				Version:  "v1",
				Metadata: pm,
				Spec:     ps,
			}
			pqs = append(pqs, pq)
		}
	}
	return pqs, nil
}

func GetPolicy(etcd client.KeysAPI, pm PolicyMetadata) (*PolicyResource, error) {
	var pq PolicyResource
	var ps PolicySpec

	tierName := pm.Tier
	if tierName == "" {
		tierName = "default"
	}
	pk := fmt.Sprintf("/calico/v1/policy/tier/%s/policy/%s", tierName, pm.Name)

	resp, err := etcd.Get(context.Background(), pk, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(resp.Node.Value), &ps)
	if err != nil {
		return nil, err
	}
	pq = PolicyResource{
		Kind:     "policy",
		Version:  "v1",
		Metadata: pm,
		Spec:     ps,
	}

	return &pq, nil
}

func DeletePolicy(etcd client.KeysAPI, pm PolicyMetadata) error {
	tierName := pm.Tier
	if tierName == "" {
		tierName = "default"
	}
	pk := fmt.Sprintf("/calico/v1/policy/tier/%s/policy/%s", tierName, pm.Name)

	_, err := etcd.Delete(context.Background(), pk, nil)
	return err
}

func createDefaultTier(etcd client.KeysAPI) error {
	ts := TierSpec{Order: 1000}
	tb, _ := json.Marshal(ts)

	//TODO: Handle already exists, for now just overwrite the default each time
	_, err := etcd.Set(context.Background(), "/calico/v1/policy/tier/default/metadata", string(tb), &client.SetOptions{})
	return err
}
