// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalLoop(node ast.Loop, env *object.Env) object.Object {
	for {
		var evaluatedLoopStatement = Eval(node.GetBody(), env)
		var errorType int = ErrorType(evaluatedLoopStatement)
		if errorType == FATAL {
			return evaluatedLoopStatement
		}
		if evaluatedLoopStatement.Type() == object.RETURN {
			var forReturn object.Return = evaluatedLoopStatement.(object.Return)
			if forReturn.Value.Type() == object.ERROR {
				var errForReturn object.Error = forReturn.Value.(object.Error)
				if errForReturn.Number.Value == constants.ERROR_BREAK {
					break
				}
			}
		}
		if errorType != 1 {
			return evaluatedLoopStatement
		}
	}
	return Nil(node.GetLine())
}
