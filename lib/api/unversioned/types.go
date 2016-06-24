package unversioned

// ---- Type metadata ----
//
type TypeMetadata struct {
	Kind       string `json:"kind" validate:"required"`
	APIVersion string `json:"apiVersion" validate:"required"`
}
