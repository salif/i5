// SPDX-License-Identifier: GPL-3.0-or-later
package object

import (
	"github.com/i5/i5/src/ast"
)

type Function struct {
	Params []*ast.Identifier
	Body   *ast.Block
	Env    *Env
}

func (f *Function) Type() TYPE {
	return FUNCTION
}

func (f *Function) StringValue() string {
	return "[type: function]"
}

type Return struct {
	Value Object
}

func (r *Return) Type() TYPE {
	return RETURN
}

func (r *Return) StringValue() string {
	return r.Value.StringValue()
}
