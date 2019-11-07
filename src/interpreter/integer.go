// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/object"
)

func evalInteger(node ast.Integer, env *object.Env) object.Object {
	return object.Integer{Value: node.GetValue()}
}
