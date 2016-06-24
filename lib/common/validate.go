package common

import (
	"reflect"
	"regexp"

	"gopkg.in/go-playground/validator.v8"
	"fmt"
)

var validate *validator.Validate
var nameRegex *regexp.Regexp
var actionRegex *regexp.Regexp
var protocolRegex *regexp.Regexp

func init() {
	// Initialise static data.
	config := &validator.Config{TagName: "validate"}

	config = &validator.Config{TagName: "validate", FieldNameTag: "json"}
	validate = validator.New(config)

	var err error
	if nameRegex, err = regexp.Compile("[a-zA-Z0-9_-]+"); err != nil { panic(err) }
	if actionRegex, err = regexp.Compile("next-tier|allow|deny"); err != nil { panic(err) }
	if protocolRegex, err = regexp.Compile("tcp|udp|icmp|icmpv6|sctp|udplite"); err != nil { panic(err) }
}

func RegisterFieldValidator(key string, fn validator.Func) {
	validate.RegisterValidation(key, fn)
}

func RegisterStructValidator(fn validator.StructLevelFunc, t ...interface{}) {
	validate.RegisterStructValidation(fn, t...)
}

func Validate(current interface{}) error {
	return validate.Struct(current)
}

// Common validation functions (accessed using the validate field tag)

func init() {
	// Register some common validators.
	RegisterFieldValidator("action", validateAction)
	RegisterFieldValidator("name", validateName)
	RegisterFieldValidator("selector", validateSelector)
	RegisterFieldValidator("tag", validateTag)
	RegisterFieldValidator("labels", validateLabels)
	RegisterFieldValidator("interface", validateInterface)

	RegisterStructValidator(validateOrder, Order{})
	RegisterStructValidator(validateProtocol, Protocol{})
}

func validateAction(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	b := []byte(field.String())
	fmt.Printf("Validate action: %s\n", b)
	return actionRegex.Match(b)
}

func validateName(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	b := []byte(field.String())
	fmt.Printf("Validate name: %s\n", b)
	return nameRegex.Match(b)
}

func validateSelector(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	b := []byte(field.String())
	fmt.Printf("Validate selector: %s\n", b)
	return nameRegex.Match(b)
}

func validateTag(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	b := []byte(field.String())
	fmt.Printf("Validate tag: %s\n", b)
	return nameRegex.Match(b)
}

func validateLabels(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	l := field.Interface().(map[string]string)
	fmt.Printf("Validate labels: %s\n", l)
	for k, v := range l {
		if nameRegex.Match([]byte(k)) || nameRegex.Match([]byte(v)) {
			return false
		}
	}
	return true
}

func validateInterface(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	b := []byte(field.String())
	fmt.Printf("Validate interface: %s\n", b)
	return nameRegex.Match(b)
}

func validateOrder(v *validator.Validate, structLevel *validator.StructLevel) {
	fmt.Printf("Validate order")
	o := structLevel.CurrentStruct.Interface().(Order)
	fmt.Printf("Validate protocol: %v %s %v\n", o.Type, o.StrVal, o.NumVal)
	if o.Type == NumOrStringString && o.StrVal != "default" {
		structLevel.ReportError(reflect.ValueOf(o.StrVal), "Order", "order", "orderStr")
	}
}

func validateProtocol(v *validator.Validate, structLevel *validator.StructLevel) {
	fmt.Printf("Validate protocol")
	p := structLevel.CurrentStruct.Interface().(Protocol)
	fmt.Printf("Validate protocol: %v %s %v\n", p.Type, p.StrVal, p.NumVal)
	if p.Type == NumOrStringNum && ((p.NumVal < 0) || (p.NumVal > 255)) {
		structLevel.ReportError(reflect.ValueOf(p.NumVal), "Protocol", "protocol", "protocolNum")
	} else if p.Type == NumOrStringString && !protocolRegex.Match([]byte(p.StrVal)) {
		structLevel.ReportError(reflect.ValueOf(p.StrVal), "Protocol", "protocol", "protocolStr")
	}
}
