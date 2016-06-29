package backend

type PolicyKey struct {
	Name string `json:"-" validate:"required,name"`
}

type Policy struct {
	PolicyKey  `json:"-"`
	Order         *float32 `json:"order"`
	InboundRules  *[]Rule  `json:"inbound_rules,omitempty" validate:"omitempty,dive"`
	OutboundRules *[]Rule  `json:"outbound_rules,omitempty" validate:"omitempty,dive"`
	Selector      string   `json:"selector" validate:"selector"`
}

type PolicyListOptions struct {
	Name *string
}
