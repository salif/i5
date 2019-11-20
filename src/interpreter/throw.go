// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalThrow(node ast.Throw, env *object.Env) object.Object {
	var evaluatedRight object.Object = Eval(node.GetBody(), env)
	if ErrorType(evaluatedRight) == FATAL {
		return evaluatedRight
	}
	if evaluatedRight.Type() == object.ERROR {
		var err object.Error = evaluatedRight.(object.Error)
		err.IsFatal = true
		return err
	} else {
		return newError(true, node.GetLine(), constants.ERROR_INTERTAL, constants.IR_IS_NOT_AN_ERROR, evaluatedRight.Type())
	}
}
