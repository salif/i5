package ast

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/i5/i5/src/types"
)

type Function struct {
	Token    types.Token
	Function string
	Params   []Identifier
	Body     Block
}

func (f Function) Value() string {
	return f.Token.Value
}

func (f Function) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Params {
		params = append(params, p.String())
	}
	out.WriteString(fmt.Sprintf("%s %s", f.Value(), f.Function))
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(f.Body.String())
	return out.String()
}

func (f Function) statement() {}
