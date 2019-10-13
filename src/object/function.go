// SPDX-License-Identifier: GPL-3.0-or-later
package object

import (
	"github.com/i5/i5/src/ast"
)

type Function struct {
	Params []ast.Identifier
	Body   ast.Block
	Env    *Env
}

func (this Function) Type() TYPE {
	return FUNCTION
}

func (this Function) StringValue() string {
	return "[type: function]"
}

type Return struct {
	Value Object
}

func (this Return) Type() TYPE {
	return RETURN
}

func (this Return) StringValue() string {
	return this.Value.StringValue()
}
