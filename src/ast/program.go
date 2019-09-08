package ast

import "bytes"

type Program struct {
	Body []Statement
}

func (p Program) Value() string {
	if len(p.Body) > 0 {
		return p.Body[0].Value()
	}
	return ""
}

func (p Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Body {
		out.WriteString(s.String())
	}
	return out.String()
}
