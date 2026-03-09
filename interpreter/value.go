package interpreter

import (
	"fmt"
	"math"
	"strings"
)

const (
	vUndefined = '?' // անորոշ
	vBoolean   = 'B' // տրամաբանական
	vNumber    = 'N' // թվային
	vText      = 'T' // տեքստային
	vArray     = 'A' // զանգված
)

// Value Ունիվերսալ արժեք
type value struct {
	kind    rune     // տեսակը
	boolean bool     // տրամաբանական արժեք
	number  float64  // թվային արժեք
	text    string   // տեքստային արժեք
	array   []*value // արժեքների զանգված
}

func (v *value) isBoolean() bool { return v.kind == vBoolean }
func (v *value) isNumber() bool  { return v.kind == vNumber }
func (v *value) isText() bool    { return v.kind == vText }
func (v *value) isArray() bool   { return v.kind == vArray }

func (v *value) String() string {
	switch v.kind {
	case vBoolean:
		return strings.ToUpper(fmt.Sprint(v.boolean))
	case vNumber:
		return fmt.Sprintf("%g", v.number)
	case vText:
		return v.text
	case vArray:
		res := ""
		for i, e := range v.array {
			if i != 0 {
				res += ", "
			}
			res += e.String()
		}
		return "[" + res + "]"
	}
	return "<undefined>"
}

func (v *value) clone() *value {
	cloned := *v

	if cloned.kind == vArray {
		cloned.array = make([]*value, len(v.array))
		for i, e := range v.array {
			cloned.array[i] = e.clone()
		}
	}

	return &cloned
}

func eq(x, y *value) bool {
	if x == nil || y == nil {
		return x == y
	}

	switch {
	case x.isBoolean() && y.isBoolean():
		return x.boolean == y.boolean
	case x.isNumber() && y.isNumber():
		return math.Abs(x.number-y.number) < 1e-9
	case x.isText() && y.isText():
		return x.text == y.text
	case x.isArray() && y.isArray():
		if len(x.array) != len(y.array) {
			return false
		}

		for i, e := range x.array {
			if !eq(e, y.array[i]) {
				return false
			}
		}
		return true
	}

	return false
}

func lt(x, y *value) bool {
	switch {
	case x.isNumber() && y.isNumber():
		return x.number < y.number
	case x.isText() && y.isText():
		return x.text < y.text
	}

	return false
}
