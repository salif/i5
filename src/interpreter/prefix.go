// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
	"github.com/i5/i5/src/types"
)

func evalPrefix(operator string, right object.Object, env *object.Env) object.Object {
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
			return newError(constants.IR_INVALID_PREFIX, operator, right.Type())
		}
	case types.BNOT:
		if right.Type() == object.INTEGER {
			value := right.(*object.Integer).Value
			return &object.Integer{Value: ^value}
		} else {
			return newError(constants.IR_INVALID_PREFIX, operator, right.Type())
		}
	case types.MINUS:
		if right.Type() == object.INTEGER {
			value := right.(*object.Integer).Value
			return &object.Integer{Value: -value}
		} else if right.Type() == object.FLOAT {
			value := right.(*object.Float).Value
			return &object.Float{Value: -value}
		} else {
			return newError(constants.IR_INVALID_PREFIX, operator, right.Type())
		}
	default:
		return newError(constants.IR_INVALID_PREFIX, operator, right.Type())
	}
}
