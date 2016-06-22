package common

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Order is a type that can hold an int or a value indicating "default".  When used in
// JSON or YAML marshalling and unmarshalling, it produces or consumes the
// inner type.  This allows you to have, for example, a JSON field that can
// accept a number or the dtring value "default".
type Order struct {
	Kind   OrderKind
	IntVal int
}

// IntstrKind represents the stored type of IntOrString.
type OrderKind int

const (
	OrderInt     OrderKind = iota // The Order holds an int.
	OrderDefault                  // The Order holds "default".
)

// UnmarshalJSON implements the json.Unmarshaller interface.
func (order *Order) UnmarshalJSON(value []byte) error {
	if value[0] == '"' {
		var s string
		err := json.Unmarshal(value, &s)
		if err != nil {
			return err
		}
		if s != "default" {
			return fmt.Errorf("order is not an integer or default")
		}
		order.Kind = OrderDefault
		return nil
	}
	order.Kind = OrderInt
	return json.Unmarshal(value, &order.IntVal)
}

// String returns the string value, or Itoa's the int value.
func (order *Order) String() string {
	if order.Kind == OrderDefault {
		return "default"
	}
	return strconv.Itoa(order.IntVal)
}

// MarshalJSON implements the json.Marshaller interface.
func (order Order) MarshalJSON() ([]byte, error) {
	switch order.Kind {
	case OrderInt:
		return json.Marshal(order.IntVal)
	case OrderDefault:
		return []byte("default"), nil
	default:
		return []byte{}, fmt.Errorf("impossible Order.Kind")
	}
}


