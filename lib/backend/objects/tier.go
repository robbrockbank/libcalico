package objects

import (
	. "github.com/projectcalico/libcalico/lib/common"
)

type Tier struct {
	Order Order `json:"order" validate:"required"`
}
