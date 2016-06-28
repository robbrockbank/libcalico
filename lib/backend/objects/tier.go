package objects

type Tier struct {
	Name string `json:"-" validate:"required,name"`

	Order *float32 `json:"order"`
}

type TierListOptions struct {
	Name *string
}
