package backend

type TierKey struct {
	Name string `json:"-" validate:"required,name"`
}

type Tier struct {
	TierKey `json:"-"`
	Order *float32 `json:"order"`
}

type TierListOptions struct {
	Name *string
}
