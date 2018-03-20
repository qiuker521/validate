package validate

import (
	"errors"
	"reflect"
	"strings"
)

type m map[string]func(path string, content interface{}) error

type V struct {
	v      m
	tag    string
	strict bool
}

func AddCheck(name string, checkFunction func(path string, content interface{}) error) {
	if s.v == nil {
		s.v = make(m)
	}

	s.v[name] = checkFunction
}

func (s *V) AddCheck(name string, checkFunction func(path string, content interface{}) error) {
	s.v[name] = checkFunction
}

func List() (list []string) {
	for k, _ := range s.v {
		list = append(list, k)
	}
	return
}

func (s *V) List() (list []string) {
	for k, _ := range s.v {
		list = append(list, k)
	}
	return
}
func Validate(x interface{}) []error {
	return s.validate("", x)
}

func (s *V) Validate(x interface{}) []error {
	return s.validate("", x)
}

func (s *V) validate(path string, x interface{}) []error {

	if s.tag == "" {
		s.tag = "validate"
	}

	valueOf := reflect.ValueOf(x)

	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}
	valueType := valueOf.Type()

	var errs []error

	for i := 0; i < valueType.NumField(); i++ {
		currentNode := path + valueType.Field(i).Name
		valueField := valueOf.Field(i)
		tag := valueType.Field(i).Tag.Get(s.tag)
		content := valueField.Interface()

		// log.Printf("%d: %s %s %#v = %v\n",
		// 	i, valueType.Field(i).Name, valueField.Type(), valueType.Field(i).Tag.Get("validate"), valueField.Interface())

		tags := strings.Split(tag, ",")

		for _, v := range tags {
			if checkFunc, ok := s.v[v]; ok {
				e := checkFunc(currentNode, content)
				if e != nil {
					errs = append(errs, e)
				}
			} else if s.strict {
				errs = append(errs, errors.New("Validator "+v+" does not exist!"))
			}
		}

		if valueField.Kind() == reflect.Struct {
			errs = append(errs, s.validate(currentNode+".", valueField.Interface())...)
		}
	}
	return errs
}

func ValidateTag(tag string) {

}

func (s *V) ValidateTag(tag string) {
	s.tag = tag
}

var s *V

func init() {
	s = new(V)
	s.v = make(m)
	addDefaultCheck(s)
}

func New() *V {
	s := new(V)
	s.v = make(m)
	addDefaultCheck(s)
	return s
}

func addDefaultCheck(s *V) {
	s.AddCheck("notnull", validateNotNull)
	s.AddCheck("mail", validateEmail)
	s.AddCheck("email", validateEmail)
	s.AddCheck("notempty", validateNotEmpty)
	s.AddCheck("nospace", validateNoSpace)
}

func NewEmptyValidator() *V {
	s := new(V)
	return s
}

func StrictMode(x bool) {
	s.strict = x
}

func (s *V) StrictMode(x bool) {
	s.strict = x
}
