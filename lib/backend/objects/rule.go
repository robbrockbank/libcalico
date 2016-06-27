package objects

import (
	. "github.com/projectcalico/libcalico/lib/common"
	"reflect"
	"gopkg.in/go-playground/validator.v8"
)

type Rule struct {
	Action string `json:"action" validate:"action"`

	Protocol    *Protocol `json:"protocol,omitempty" validate:"omitempty"`
	SrcTag      *string        `json:"srcTag,omitempty" validate:"omitempty,tag"`
	SrcNet      *IPNet         `json:"srcNet,omitempty" validate:"omitempty"`
	SrcSelector *string        `json:"srcSelector,omitempty" validate:"omitempty,selector"`
	SrcPorts    *[]int         `json:"srcPorts,omitempty" validate:"omitempty,dive,gte=0,lte=65535"`
	DstTag      *string        `json:"dstTag,omitempty" validate:"omitempty,tag"`
	DstSelector *string        `json:"dstSelector,omitempty" validate:"omitempty,selector"`
	DstNet      *IPNet         `json:"dstNet,omitempty" validate:"omitempty"`
	DstPorts    *[]int         `json:"dstPorts,omitempty" validate:"omitempty,dive,gte=0,lte=65535"`
	ICMPType    *int           `json:"icmpType,omitempty" validate:"omitempty,gte=0,lte=255"`
	ICMPCode    *int           `json:"icmpCode,omitempty" validate:"omitempty,gte=0,lte=255"`

	NotProtocol    *Protocol `json:"!protocol,omitempty" validate:"omitempty"`
	NotSrcTag      *string        `json:"!srcTag,omitempty" validate:"omitempty,tag"`
	NotSrcNet      *IPNet         `json:"!srcNet,omitempty" validate:"omitempty"`
	NotSrcSelector *string        `json:"!srcSelector,omitempty" validate:"omitempty,selector"`
	NotSrcPorts    *[]int         `json:"!srcPorts,omitempty" validate:"omitempty,dive,gte=0,lte=65535"`
	NotDstTag      *string        `json:"!dstTag,omitempty" validate:"omitempty"`
	NotDstSelector *string        `json:"!dstSelector,omitempty" validate:"omitempty,selector"`
	NotDstNet      *IPNet         `json:"!dstNet,omitempty" validate:"omitempty"`
	NotDstPorts    *[]int         `json:"!dstPorts,omitempty" validate:"omitempty,dive,gte=0,lte=65535"`
	NotICMPType    *int           `json:"!icmpType,omitempty" validate:"omitempty,gte=0,lte=255"`
	NotICMPCode    *int           `json:"!icmpCode,omitempty" validate:"omitempty,gte=0,lte=255"`
}

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
