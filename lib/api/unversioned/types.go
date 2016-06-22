package unversioned

/*



 */

// ---- Type metadata ----
//
type TypeMetadata struct {
	Kind    string `json:"kind"`
	APIVersion string `json:"apiVersion"`
}

// ---- Generic resource type ----
//
type Resource struct {
	TypeMetadata `json:",inline"`
	Metadata     interface{} `json:"metadata,omitempty"`
	Spec         interface{} `json:"spec,omitempty"`
}

// ---- List of resources  ----
// Kind list
type ListMetadata struct {
}

type ListSpec struct {
	List []Resource `json:"list"`
}

func ResourceList(m *ListMetadata, s *ListSpec) Resource {
	return Resource{TypeMetadata{"list", ""}, m, s}
}
