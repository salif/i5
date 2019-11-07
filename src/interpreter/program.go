// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/object"
)

func evalProgram(node ast.Program, env *object.Env) object.Object {
	for _, assign := range node.GetBody() {
		var evaluatedAssign object.Object = Eval(assign, env)
		if ErrorType(evaluatedAssign) == FATAL {
			return evaluatedAssign
		}
	}
	return Nil(node.GetLine())
}
