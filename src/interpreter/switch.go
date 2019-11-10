// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/object"
	"github.com/i5/i5/src/types"
)

func evalSwitch(node ast.Switch, env *object.Env) object.Object {
	var evaluatedCondition object.Object = Eval(node.GetCondition(), env)
	if ErrorType(evaluatedCondition) == FATAL {
		return evaluatedCondition
	}

	for _, c := range node.GetCases() {
		result := Eval(ast.Infix{}.Set(c.GetLine(), c.GetCase(), types.EQEQ, node.GetCondition()), env)
		if ErrorType(result) == FATAL {
			return result
		}
		if isTrue(result) {
			return Eval(c.GetBody(), env)
		}
	}

	return Nil(node.GetLine())
}
