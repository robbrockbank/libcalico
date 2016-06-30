package backend

type ProfileRules struct {
	InboundRules  []Rule `json:"inbound_rules,omitempty" validate:"omitempty,dive"`
	OutboundRules []Rule `json:"outbound_rules,omitempty" validate:"omitempty,dive"`
}

type ProfileKey struct {
	Name string `json:"-" validate:"required,name"`
}

type Profile struct {
	Key    ProfileKey
	Rules  ProfileRules
	Tags   []string
	Labels map[string]string
}

type ProfileListOptions struct {
	Name *string
}
