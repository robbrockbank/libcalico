package objects

type ProfileRules struct {
	InboundRules  *[]Rule `json:"inbound_rules,omitempty" validate:"omitempty,dive"`
	OutboundRules *[]Rule `json:"outbound_rules,omitempty" validate:"omitempty,dive"`
}

type Profile struct {
	Name   string
	Rules  ProfileRules
	Tags   *[]string
	Labels *map[string]string
}

type ProfileListOptions struct {
	Name *string
}
