package objects

type ProfileRules struct {
	Name string `json:"-" validate:"required,name"`

	InboundRules  *[]Rule            `json:"inbound_rules,omitempty" validate:"omitempty,dive"`
	OutboundRules *[]Rule            `json:"outbound_rules,omitempty" validate:"omitempty,dive"`
}

type ProfileLabels struct {
	Name   string

	Labels *map[string]string
}

type ProfileTags struct {
	Name string

	Tags *[]string
}

type ProfileListOptions struct {
	Name     *string
}