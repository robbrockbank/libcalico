package common

import (
	"reflect"
	"regexp"

	"github.com/golang/glog"
	"gopkg.in/go-playground/validator.v8"
	"github.com/projectcalico/libcalico/lib/selector"
)

var validate *validator.Validate

// [smc] why not define these inline to save typing on the types?
var nameRegex = regexp.MustCompile("[a-zA-Z0-9_-]+")
var actionRegex = regexp.MustCompile("next-tier|allow|deny")
var protocolRegex = regexp.MustCompile("tcp|udp|icmp|icmpv6|sctp|udplite")

func init() {
	// Initialise static data.
	config := &validator.Config{TagName: "validate", FieldNameTag: "json"}
	validate = validator.New(config)
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
// [smc] More than one init() func?
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
	// [smc] Why not use MatchString to avoid the byte[] copy/cast?
	s := field.String()
	glog.V(2).Infof("Validate action: %s\n", s)
	return actionRegex.MatchString(s)
}

func validateName(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	b := []byte(field.String())
	glog.V(2).Infof("Validate name: %s\n", b)
	return nameRegex.Match(b)
}

func validateSelector(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	s := field.String()
	glog.V(2).Infof("Validate selector: %s\n", s)
	_, err := selector.Parse(s)
	if err != nil {
		glog.Errorf("Selector %#v was invalid: %v", s, err)
		return false
	}
	return true
}

func validateTag(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	b := []byte(field.String())
	glog.V(2).Infof("Validate tag: %s\n", b)
	return nameRegex.Match(b)
}

func validateLabels(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	l := field.Interface().(map[string]string)
	glog.V(2).Infof("Validate labels: %s\n", l)
	for k, v := range l {
		if nameRegex.Match([]byte(k)) || nameRegex.Match([]byte(v)) {
			return false
		}
	}
	return true
}

func validateInterface(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	b := []byte(field.String())
	glog.V(2).Infof("Validate interface: %s\n", b)
	return nameRegex.Match(b)
}

func validateProtocol(v *validator.Validate, structLevel *validator.StructLevel) {
	glog.V(2).Infof("Validate protocol")
	p := structLevel.CurrentStruct.Interface().(Protocol)
	glog.V(2).Infof("Validate protocol: %v %s %v\n", p.Type, p.StrVal, p.NumVal)
	// [smc] I think the kernel doesn't support protocol 0:
	if p.Type == NumOrStringNum && ((p.NumVal < 1) || (p.NumVal > 255)) {
		structLevel.ReportError(reflect.ValueOf(p.NumVal), "Protocol", "protocol", "protocolNum")
	} else if p.Type == NumOrStringString && !protocolRegex.Match([]byte(p.StrVal)) {
		structLevel.ReportError(reflect.ValueOf(p.StrVal), "Protocol", "protocol", "protocolStr")
	}
}
