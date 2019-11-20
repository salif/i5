// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalIf(node ast.If, env *object.Env) object.Object {
	var evaluatedCondition object.Object = Eval(node.GetCondition(), env)

	if ErrorType(evaluatedCondition) == FATAL {
		return evaluatedCondition
	}

	if evaluatedCondition.Type() != object.BOOL {
		return newError(true, node.GetLine(), constants.ERROR_INTERTAL, constants.IR_IS_NOT_A_BOOL, evaluatedCondition.Type())
	}
	if isTrue(evaluatedCondition) {
		return Eval(node.GetConsequence(), env)
	} else if node.HaveAlternative() {
		return Eval(node.GetAlternative(), env)
	} else {
		return Nil(node.GetLine())
	}
}
