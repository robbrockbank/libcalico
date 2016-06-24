package v1

import (
	"reflect"

	. "github.com/projectcalico/libcalico/lib/common"
	"gopkg.in/go-playground/validator.v8"
)

// Register v1 structure validators to validate cross-field dependencies in any of the
// required structures.
func init() {
	RegisterStructValidator(validateRule, Rule{})
}

func validateRule(v *validator.Validate, structLevel *validator.StructLevel) {
	rule := structLevel.CurrentStruct.Interface().(Rule)
	if (rule.ICMPCode != nil && rule.ICMPType == nil) {
		structLevel.ReportError(reflect.ValueOf(rule.ICMPCode), "ICMPCode", "icmpCode", "icmpCodeWithoutType")
	}

	// TODO other cross-struct validation
}
