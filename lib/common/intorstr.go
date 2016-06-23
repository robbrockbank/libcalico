package intorstr

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Int32OrString is a type that can hold an int32 or a string.  When used in
// JSON or YAML marshalling and unmarshalling, it produces or consumes the
// inner type.  This allows you to have, for example, a JSON field that can
// accept a name or number.
type Int32OrString struct {
	Type   Type
	IntVal int32
	StrVal string
}

// Type represents the stored type of Int32OrString.
type Type int

const (
	Int    Type = iota // The Int32OrString holds an int.
	String             // The Int32OrString holds a string.
)

// UnmarshalJSON implements the json.Unmarshaller interface.
func (intstr *Int32OrString) UnmarshalJSON(value []byte) error {
	if value[0] == '"' {
		intstr.Type = String
		return json.Unmarshal(value, &intstr.StrVal)
	}
	intstr.Type = Int
	return json.Unmarshal(value, &intstr.IntVal)
}

// String returns the string value, or the Itoa of the int value.
func (intstr *Int32OrString) String() string {
	if intstr.Type == String {
		return intstr.StrVal
	}
	return strconv.Itoa(intstr.IntValue())
}

// IntValue returns the IntVal if type Int, or if
// it is a String, will attempt a conversion to int.
func (intstr *Int32OrString) IntValue() (int, error) {
	if intstr.Type == String {
		return strconv.Atoi(intstr.StrVal)
	}
	return int(intstr.IntVal), nil
}

// MarshalJSON implements the json.Marshaller interface.
func (intstr Int32OrString) MarshalJSON() ([]byte, error) {
	switch intstr.Type {
	case Int:
		return json.Marshal(intstr.IntVal)
	case String:
		return json.Marshal(intstr.StrVal)
	default:
		return []byte{}, fmt.Errorf("impossible Int32OrString.Type")
	}
}
