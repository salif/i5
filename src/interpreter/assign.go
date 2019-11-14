// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalAssign(node ast.Assign, env *object.Env) object.Object {
	var left ast.Node = node.GetLeft()
	var evaluatedRight object.Object = Eval(node.GetRight(), env)

	if ErrorType(evaluatedRight) == FATAL {
		return evaluatedRight
	}

	switch left := left.(type) {

	case ast.Identifier:
		env.Set(left.GetValue(), evaluatedRight)
		return evaluatedRight

	case ast.Index:
		var leftIndex object.Object = Eval(left.GetLeft(), env)

		if ErrorType(leftIndex) == FATAL {
			return leftIndex
		}

		if leftIndex.Type() == object.CLASSOBJECT {
			switch rightIndex := left.GetRight().(type) {

			case ast.Identifier:
				_class := leftIndex.(object.ClassObject)
				_class.Set(rightIndex.GetValue(), evaluatedRight)
				return _class

			default:
				return newError(true, node.GetLine(), constants.ERROR_INTERTAL, constants.IR_INVALID_INFIX, leftIndex.Type(), left.GetOperator(), rightIndex.GetType())
			}
		} else {
			return newError(true, node.GetLine(), constants.ERROR_INTERTAL, constants.IR_INVALID_POSTFIX, leftIndex.Type(), left.GetOperator())
		}

	default:
		return newError(true, node.GetLine(), constants.ERROR_INTERTAL, constants.IR_CANNOT_ASSIGN, left.GetType())
	}
}
