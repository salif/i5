// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/object"
)

func evalReturn(node ast.Return, env *object.Env) object.Object {
	var evaluatedRight object.Object = Eval(node.GetBody(), env)
	if ErrorType(evaluatedRight) == FATAL {
		return evaluatedRight
	}
	return object.Return{Value: evaluatedRight}
}
