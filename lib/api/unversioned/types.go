package unversioned

// ---- Type metadata ----
//
type TypeMetadata struct {
	Kind       string `json:"kind" validate:"required"`
	APIVersion string `json:"apiVersion" validate:"required"`
}

// ---- Metadata common to all resources ----
type ObjectMetadata struct {
	Name string `json:"name,omitempty" validate:"omitempty,name"`
}

// ---- Metadata common to all lists ----
type ListMetadata struct {
}
