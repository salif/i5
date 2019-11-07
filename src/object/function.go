// SPDX-License-Identifier: GPL-3.0-or-later
package object

import (
	"github.com/i5/i5/src/ast"
)

type Function struct {
	Params ast.Identifiers
	Body   ast.Node
	Env    *Env
}

func (this Function) Type() TYPE {
	return FUNCTION
}

func (this Function) StringValue() string {
	return "[type: function]"
}
