// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
	"github.com/i5/i5/src/types"
)

func evalPrefixNode(node ast.Prefix, env *object.Env) object.Object {
	var evaluatedRight object.Object = Eval(node.GetRight(), env)
	if ErrorType(evaluatedRight) == FATAL {
		return evaluatedRight
	}
	return evalPrefix(node.GetOperator(), evaluatedRight, env, node.GetLine())
}

func evalPrefix(operator string, right object.Object, env *object.Env, line uint32) object.Object {
	switch operator {
	case types.NOT:
		if right.Type() == object.BOOL {
			switch right {
			case TRUE:
				return FALSE
			case FALSE:
				return TRUE
			default:
				return FALSE
			}
		} else {
			return newError(true, line, constants.ERROR_INTERTAL, constants.IR_INVALID_PREFIX, operator, right.Type())
		}
	case types.BNOT:
		if right.Type() == object.INTEGER {
			value := right.(object.Integer).Value
			return object.Integer{Value: ^value}
		} else {
			return newError(true, line, constants.ERROR_INTERTAL, constants.IR_INVALID_PREFIX, operator, right.Type())
		}
	case types.MINUS:
		if right.Type() == object.INTEGER {
			value := right.(object.Integer).Value
			return object.Integer{Value: -value}
		} else if right.Type() == object.FLOAT {
			value := right.(object.Float).Value
			return object.Float{Value: -value}
		} else {
			return newError(true, line, constants.ERROR_INTERTAL, constants.IR_INVALID_PREFIX, operator, right.Type())
		}
	default:
		return newError(true, line, constants.ERROR_INTERTAL, constants.IR_INVALID_PREFIX, operator, right.Type())
	}
}
