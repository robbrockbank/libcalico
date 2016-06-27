package objects

type Profile struct {
	IngressRules *[]Rule            `json:"ingress,omitempty" validate:"omitempty,dive"`
	EgressRules  *[]Rule            `json:"egress,omitempty" validate:"omitempty,dive"`
	Labels       *map[string]string `json:"labels,omitempty" validate:"omitempty,labels"`
	Tags         *[]string          `json:"tags,omitempty" validate:"omitempty,dive,tag"`
}