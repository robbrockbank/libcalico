package libcalico

import "regexp"

var policyRE = regexp.MustCompile(`/calico/v1/policy/tier/[^/]*/policy/[^/]*`)

/*
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

		var re = regexp.MustCompile(`/calico/v1/policy/tier/([^/]+)/policy/([^/]+)`)
		matches := re.FindStringSubmatch(node.Key)
		if matches != nil {
			policyName := matches[2]

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
*/
