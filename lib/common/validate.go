package common

import (
	"reflect"
	"regexp"

	"fmt"

	"gopkg.in/go-playground/validator.v8"
)

var validate *validator.Validate
var nameRegex *regexp.Regexp
var actionRegex *regexp.Regexp
var protocolRegex *regexp.Regexp

func init() {
	// Initialise static data.
	config := &validator.Config{TagName: "validate", FieldNameTag: "json"}
	validate = validator.New(config)

	nameRegex = regexp.MustCompile("[a-zA-Z0-9_-]+")
	actionRegex = regexp.MustCompile("next-tier|allow|deny")
	protocolRegex = regexp.MustCompile("tcp|udp|icmp|icmpv6|sctp|udplite")
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
