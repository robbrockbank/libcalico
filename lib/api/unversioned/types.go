package unversioned

type Resource interface {
	GetTypeMetadata() TypeMetadata
}

// ---- Type metadata ----
//
type TypeMetadata struct {
	Kind       string `json:"kind" validate:"required"`
	APIVersion string `json:"apiVersion" validate:"required"`
}

func (md TypeMetadata) GetTypeMetadata() TypeMetadata {
	return md
}

// ---- Metadata common to all resources ----
type ObjectMetadata struct {
	Name string `json:"name,omitempty" validate:"omitempty,name"`
}

// ---- Metadata common to all lists ----
type ListMetadata struct {
}
