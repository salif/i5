// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/object"
)

func evalString(node ast.String, env *object.Env) object.Object {
	return object.String{Value: node.GetValue()}
}
