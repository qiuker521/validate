package validate

import (
	"testing"
)

type C struct {
	A string `validate:"notnull"`
	B string `validate:"notnull,notempty"`
	C string `validate:"nospace"`
	D string `validate:"nospace,email"`
}
type M struct {
	A string `validate:"notnull"`
	B string `validate:"notnull,notempty"`
	C C      `validate:"xx"`
}

func TestValidate(t *testing.T) {
	c := M{A: "N", B: "X", C: C{A: "", B: "", C: " ", D: "a@a.cx"}}
	errs := Validate(c)
	if len(errs) == 0 {
		t.Error("none strict mode not equals 0")
	}

	StrictMode(true)
	c = M{A: "N", B: "X", C: C{A: "X", B: "B", C: "C", D: "a@a.cx"}}
	errs = Validate(c)
	if len(errs) != 1 {
		t.Error("strict mode error not equals 1")
	}

}

func TestValidators(t *testing.T) {
	var c = C{}
	StrictMode(false)
	c = C{A: "", B: "x", C: "", D: "a@a.cx"}
	errs := Validate(c)
	if len(errs) != 0 {
		t.Error("Errors must be 0")
	}

	c = C{A: "", B: "", C: " ", D: "a@a.cx"}
	errs = Validate(c)
	if len(errs) != 2 {
		t.Error("Errors must be 2")
	}

}
