package ast

import (
	"fmt"
	"strings"
)

func (p *Program) String() string {
	var text string
	for _, sb := range p.Subroutines {
		text += fmt.Sprint(sb)
		text += "\n\n"
	}
	return text
}

func (s *Subroutine) String() string {
	indent()
	defer unindent()

	text := "subroutine:\n"
	text += fmt.Sprintf("%sname: %s\n", spaces, s.Name)
	text += fmt.Sprintf("%sparameters: [%s]\n", spaces, strings.Join(s.Parameters, ","))
	text += fmt.Sprintf("%sbody:\n%s", spaces, s.Body)
	return text
}

func (s *Sequence) String() string {
	indent()
	defer unindent()

	var text strings.Builder
	fmt.Fprintf(&text, "%ssequence:\n", spaces)
	for _, e := range s.Items {
		fmt.Fprint(&text, e)
	}
	return text.String()
}

func (d *Dim) String() string {
	indent()
	defer unindent()

	text := "dim:\n"
	indent()
	text += fmt.Sprintf("%sname: %s\n", spaces, d.Name)
	text += fmt.Sprintf("%ssize: %d", spaces, d.Size)
	unindent()

	return text
}

func (l *Let) String() string {
	indent()
	defer unindent()

	text := spaces + "let:\n"
	indent()
	text += fmt.Sprintf("%splace:\n%s", spaces, l.Place)
	text += fmt.Sprintf("%svalue:\n%s", spaces, l.Value)
	unindent()

	return text
}

func (i *Input) String() string {
	indent()
	defer unindent()

	text := spaces + "input:\n"
	indent()
	text += fmt.Sprintf("%splace:\n%s", spaces, i.Place)
	unindent()

	return text
}

func (p *Print) String() string {
	indent()
	defer unindent()

	text := spaces + "print:\n"
	indent()
	text += fmt.Sprintf("%svalue:\n%s", spaces, p.Value)
	unindent()

	return text
}

func (i *If) String() string {
	indent()
	defer unindent()

	text := spaces + "if:\n"
	indent()
	text += fmt.Sprintf("%scondition:\n%s", spaces, i.Condition)
	text += fmt.Sprintf("%sdecision:\n%s", spaces, i.Decision)
	text += fmt.Sprintf("%salternative:\n%s", spaces, i.Alternative)
	unindent()

	return text
}

func (w *While) String() string {
	indent()
	defer unindent()

	text := spaces + "while:\n"
	indent()
	text += fmt.Sprintf("%scondition:\n%s", spaces, w.Condition)
	text += fmt.Sprintf("%sbody:\n%s", spaces, w.Body)
	unindent()

	return text
}

func (f *For) String() string {
	indent()
	defer unindent()

	text := spaces + "for:\n"
	text += fmt.Sprintf("%sparameter:\n%s", spaces, f.Parameter)
	text += fmt.Sprintf("%sbegin:\n%s", spaces, f.Begin)
	text += fmt.Sprintf("%send:\n%s", spaces, f.End)
	text += fmt.Sprintf("%sstep:\n%s", spaces, f.Step)
	text += fmt.Sprintf("%sbody:\n%s", spaces, f.Body)
	unindent()

	return text
}

func (c *Call) String() string {
	return applyHelper("call", c.Arguments)
}

func (a *Apply) String() string {
	return applyHelper("apply", a.Arguments)
}

func applyHelper(name string, args []Expression) string {
	indent()
	defer unindent()

	var text strings.Builder
	fmt.Fprintf(&text, "%s%s:\n", spaces, name)
	indent()
	for i, e := range args {
		indent()
		fmt.Fprintf(&text, "%sargument[%d]:\n%s", spaces, i, e)
		unindent()
	}
	unindent()

	return text.String()
}

func (b *Binary) String() string {
	indent()
	defer unindent()

	var text strings.Builder
	fmt.Fprintf(&text, "%sbinary\n", spaces)
	indent()
	fmt.Fprintf(&text, "%soperation: '%s'\n", spaces, b.Operation)
	fmt.Fprintf(&text, "%sleft:\n%s", spaces, b.Left)
	fmt.Fprintf(&text, "%sright:\n%s", spaces, b.Right)
	unindent()

	return text.String()
}

func (u *Unary) String() string {
	indent()
	defer unindent()

	var text strings.Builder
	fmt.Fprintf(&text, "%sunary\n", spaces)
	fmt.Fprintf(&text, "%soperation: '%s'\n", spaces, u.Operation)
	fmt.Fprintf(&text, "%sleft:\n%s", spaces, u.Right)
	unindent()

	return text.String()
}

func (v *Variable) String() string {
	indent()
	defer unindent()

	text := spaces + "variable:\n"
	indent()
	text += fmt.Sprintf("%sname: '%v'\n", spaces, v.Name)
	unindent()
	return text
}

func (a *Array) String() string {
	indent()
	defer unindent()

	var els strings.Builder
	fmt.Fprintf(&els, "%s[\n", spaces)
	for _, e := range a.Elements {
		fmt.Fprint(&els, e)
	}
	fmt.Fprintf(&els, "%s]\n", spaces)
	return els.String()
}

func (b *Boolean) String() string {
	indent()
	defer unindent()

	var text strings.Builder
	fmt.Fprintf(&text, "%sboolean:\n", spaces)
	indent()
	fmt.Fprintf(&text, "%svalue: %v\n", spaces, b.Value)
	unindent()
	return text.String()
}

func (n *Number) String() string {
	indent()
	defer unindent()

	var text strings.Builder
	fmt.Fprintf(&text, "%snumber:\n", spaces)
	indent()
	fmt.Fprintf(&text, "%svalue: %v\n", spaces, n.Value)
	unindent()
	return text.String()
}

func (t *Text) String() string {
	indent()
	defer unindent()

	var text strings.Builder
	fmt.Fprintf(&text, "%stext:\n", spaces)
	indent()
	fmt.Fprintf(&text, "%svalue: \"%s\"\n", spaces, t.Value)
	unindent()
	return text.String()
}

var spaces = ""

func indent() { spaces += "  " }
func unindent() {
	l := len(spaces)
	if l >= 2 {
		spaces = spaces[0 : l-2]
	}
}
