// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalPostfixNode(node ast.Postfix, env *object.Env) object.Object {
	var evaluatedLeft object.Object = Eval(node.GetLeft(), env)
	if ErrorType(evaluatedLeft) == FATAL {
		return evaluatedLeft
	}
	return evalPostfix(node.GetOperator(), evaluatedLeft, env, node.GetLine())
}

func evalPostfix(operator string, right object.Object, env *object.Env, line uint32) object.Object {
	switch operator {
	default:
		return newError(true, line, constants.ERROR_INTERTAL, constants.IR_INVALID_POSTFIX, right.Type(), operator)
	}
}
