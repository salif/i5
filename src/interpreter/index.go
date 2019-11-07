// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalIndex(node ast.Index, env *object.Env) object.Object {
	var left object.Object = Eval(node.GetLeft(), env)
	if ErrorType(left) == FATAL {
		return left
	}
	if left.Type() == object.MAP {
		switch rnode := node.GetRight().(type) {
		case ast.Identifier:
			_map := left.(object.Map)
			obj := _map.Get(object.String{Value: rnode.GetValue()})
			return obj
		case ast.Integer:
			_map := left.(object.Map)
			obj := _map.Get(object.Integer{Value: rnode.GetValue()})
			return obj

		default:
			return newError(true, node.GetLine(), constants.ERROR_INTERTAL, constants.IR_INVALID_INFIX, left.Type(), node.GetOperator(), rnode.GetType())
		}
	} else {
		return newError(true, node.GetLine(), constants.ERROR_INTERTAL, constants.IR_INVALID_POSTFIX, left.Type(), node.GetOperator())
	}
}