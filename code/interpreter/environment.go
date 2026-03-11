package interpreter

import "fmt"

type scope struct {
	items map[string]*value // փոփոխականների ընթացիկ արժեքներ
	up    *scope            // ընդգրկող scope-ի ցուցիչ
}

// Կատարման միջավայրը
type environment struct {
	parent  *environment
	current *scope
}

func (e *environment) openScope() {
	e.current = &scope{
		items: make(map[string]*value),
		up:    e.current,
	}
}

func (e *environment) closeScope() {
	if e.current != nil {
		e.current = e.current.up
	}
}

// միջավայրում ավելացնում է տրված անվան տրված արժեքը
func (e *environment) set(name string, value *value) {
	if e.current != nil {
		e.current.items[name] = value
	}
}

// միջավայրում որոնում է և վերադարձնում է տրված անվանը
// համապատասխանող արժեքը
func (e *environment) get(name string) *value {
	for p := e.current; p != nil; p = p.up {
		if v, exists := p.items[name]; exists {
			return v
		}
	}

	if e.parent != nil {
		return e.parent.get(name)
	}

	return nil
}

func (e *environment) String() string {
	text := "=============\n"
	for p := e.current; p != nil; p = p.up {
		text += "-----------\n"
		for n, v := range p.items {
			text += fmt.Sprintf("| %s = %v\n", n, v)
		}
		text += "\n"
	}
	return text
}
