package ast

import (
	"testing"
)

func TestLiterals(t *testing.T) {
	n0 := &Number{Value: 42}
	e0 := "  number:\n    value: 42\n"
	if n0.String() != e0 {
		t.Fail()
	}

	t0 := &Text{Value: "hello"}
	e1 := "  text:\n    value: \"hello\"\n"
	if t0.String() != e1 {
		t.Fail()
	}

	b0 := &Boolean{Value: true}
	e2 := "  boolean:\n    value: true\n"
	if b0.String() != e2 {
		t.Fail()
	}
}

func TestVariable(t *testing.T) {
	v0 := &Variable{Name: "x"}
	e0 := "  variable:\n    name: 'x'\n"
	if v0.String() != e0 {
		t.Fail()
	}
}
