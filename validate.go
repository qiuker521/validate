package validate

import (
	"reflect"
	"strings"
)

type V struct {
	v   map[string]func(interface{}) error
	tag string
}

func AddCheck(name string, checkFunction func(interface{}) error) {
	s.v[name] = checkFunction
}

func (s *V) AddCheck(name string, checkFunction func(interface{}) error) {
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
	return s.validate(x)
}

func (s *V) Validate(x interface{}) []error {
	return s.validate(x)
}
func (s *V) validate(x interface{}) []error {
	if s.tag == "" {
		s.tag = "validate"
	}
	val := reflect.ValueOf(s)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	t := val.Type()
	if t == nil || t.Kind() != reflect.Struct {
		return nil
	}

	var errs []error

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fv := val.Field(i)
		if !fv.CanInterface() {
			continue
		}
		val := fv.Interface()
		tag := f.Tag.Get(s.tag)
		if tag == "" {
			continue
		}
		vts := strings.Split(tag, ",")

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
}

func New() *V {
	s := new(V)
	return s
}
