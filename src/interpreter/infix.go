// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/object"
	"github.com/i5/i5/src/types"
)

func evalInfix(operator string, left, right object.Object, env *object.Env, line int) object.Object {
	if operator == types.COLON {
		return object.String{Value: left.StringValue() + right.StringValue()}
	} else if left.Type() == right.Type() && left.Type() == object.INTEGER {
		return evalIntegerWithIntegerInfix(operator, left, right, line)
	} else if left.Type() == right.Type() && left.Type() == object.FLOAT {
		return evalFloatWithFloatInfix(operator, left, right, line)
	} else if left.Type() == object.INTEGER && right.Type() == object.FLOAT {
		return evalIntegerWithFloatInfix(operator, left, right, line)
	} else if left.Type() == object.FLOAT && right.Type() == object.INTEGER {
		return evalFloatWithIntegerInfix(operator, left, right, line)
	} else if left.Type() == right.Type() && left.Type() == object.STRING {
		return evalStringInfix(operator, left, right, line)
	} else if left.Type() == right.Type() && left.Type() == object.BOOL {
		return evalBooleanInfix(operator, left, right, line)
	}
	return object.Error{Message: console.Format(constants.IR_INVALID_INFIX, left.Type(), operator, right.Type()), Line: line}
}

func evalIntegerWithIntegerInfix(operator string, left, right object.Object, line int) object.Object {
	leftVal := left.(object.Integer).Value
	rightVal := right.(object.Integer).Value

	switch operator {
	case types.PLUS:
		return object.Integer{Value: leftVal + rightVal}
	case types.MINUS:
		return object.Integer{Value: leftVal - rightVal}
	case types.MULTIPLY:
		return object.Integer{Value: leftVal * rightVal}
	case types.DIVIDE:
		return object.Integer{Value: leftVal / rightVal}
	case types.MODULO:
		return object.Integer{Value: leftVal % rightVal}
	case types.OR:
		return object.Integer{Value: leftVal | rightVal}
	case types.XOR:
		return object.Integer{Value: leftVal ^ rightVal}
	case types.AND:
		return object.Integer{Value: leftVal & rightVal}
	case types.LTLT:
		return object.Integer{Value: leftVal << uint64(rightVal)}
	case types.GTGT:
		return object.Integer{Value: leftVal >> uint64(rightVal)}
	case types.LT:
		return nativeToBool(leftVal < rightVal)
	case types.LTEQ:
		return nativeToBool(leftVal <= rightVal)
	case types.GT:
		return nativeToBool(leftVal > rightVal)
	case types.GTEQ:
		return nativeToBool(leftVal >= rightVal)
	case types.EQEQ:
		return nativeToBool(leftVal == rightVal)
	case types.NOTEQ:
		return nativeToBool(leftVal != rightVal)
	default:
		return object.Error{Message: console.Format(constants.IR_INVALID_INFIX, left.Type(), operator, right.Type()), Line: line}
	}
}

func evalFloatWithFloatInfix(operator string, left, right object.Object, line int) object.Object {
	leftVal := left.(object.Float).Value
	rightVal := right.(object.Float).Value

	switch operator {
	case types.PLUS:
		return object.Float{Value: leftVal + rightVal}
	case types.MINUS:
		return object.Float{Value: leftVal - rightVal}
	case types.MULTIPLY:
		return object.Float{Value: leftVal * rightVal}
	case types.DIVIDE:
		return object.Float{Value: leftVal / rightVal}
	case types.LT:
		return nativeToBool(leftVal < rightVal)
	case types.LTEQ:
		return nativeToBool(leftVal <= rightVal)
	case types.GT:
		return nativeToBool(leftVal > rightVal)
	case types.GTEQ:
		return nativeToBool(leftVal >= rightVal)
	case types.EQEQ:
		return nativeToBool(leftVal == rightVal)
	case types.NOTEQ:
		return nativeToBool(leftVal != rightVal)
	default:
		return object.Error{Message: console.Format(constants.IR_INVALID_INFIX, left.Type(), operator, right.Type()), Line: line}
	}
}

func evalIntegerWithFloatInfix(operator string, left, right object.Object, line int) object.Object {
	leftVal := left.(object.Integer).Value
	rightVal := right.(object.Float).Value

	switch operator {
	case types.PLUS:
		return object.Float{Value: float64(leftVal) + rightVal}
	case types.MINUS:
		return object.Float{Value: float64(leftVal) - rightVal}
	case types.MULTIPLY:
		return object.Float{Value: float64(leftVal) * rightVal}
	case types.DIVIDE:
		return object.Float{Value: float64(leftVal) / rightVal}
	case types.LT:
		return nativeToBool(float64(leftVal) < rightVal)
	case types.LTEQ:
		return nativeToBool(float64(leftVal) <= rightVal)
	case types.GT:
		return nativeToBool(float64(leftVal) > rightVal)
	case types.GTEQ:
		return nativeToBool(float64(leftVal) >= rightVal)
	case types.EQEQ:
		return nativeToBool(float64(leftVal) == rightVal)
	case types.NOTEQ:
		return nativeToBool(float64(leftVal) != rightVal)
	default:
		return object.Error{Message: console.Format(constants.IR_INVALID_INFIX, left.Type(), operator, right.Type()), Line: line}
	}
}

func evalFloatWithIntegerInfix(operator string, left, right object.Object, line int) object.Object {
	leftVal := left.(object.Float).Value
	rightVal := right.(object.Integer).Value

	switch operator {
	case types.PLUS:
		return object.Float{Value: leftVal + float64(rightVal)}
	case types.MINUS:
		return object.Float{Value: leftVal - float64(rightVal)}
	case types.MULTIPLY:
		return object.Float{Value: leftVal * float64(rightVal)}
	case types.DIVIDE:
		return object.Float{Value: leftVal / float64(rightVal)}
	case types.LT:
		return nativeToBool(leftVal < float64(rightVal))
	case types.LTEQ:
		return nativeToBool(leftVal <= float64(rightVal))
	case types.GT:
		return nativeToBool(leftVal > float64(rightVal))
	case types.GTEQ:
		return nativeToBool(leftVal >= float64(rightVal))
	case types.EQEQ:
		return nativeToBool(leftVal == float64(rightVal))
	case types.NOTEQ:
		return nativeToBool(leftVal != float64(rightVal))
	default:
		return object.Error{Message: console.Format(constants.IR_INVALID_INFIX, left.Type(), operator, right.Type()), Line: line}
	}
}

func evalStringInfix(operator string, left, right object.Object, line int) object.Object {
	leftVal := left.(object.String).Value
	rightVal := right.(object.String).Value

	switch operator {
	case types.EQEQ:
		return nativeToBool(leftVal == rightVal)
	case types.NOTEQ:
		return nativeToBool(leftVal != rightVal)
	default:
		return object.Error{Message: console.Format(constants.IR_INVALID_INFIX, left.Type(), operator, right.Type()), Line: line}
	}
}

func evalBooleanInfix(operator string, left, right object.Object, line int) object.Object {
	leftVal := left.(object.Bool).Value
	rightVal := right.(object.Bool).Value

	switch operator {
	case types.EQEQ:
		return nativeToBool(leftVal == rightVal)
	case types.NOTEQ:
		return nativeToBool(leftVal != rightVal)
	case types.ANDAND:
		return nativeToBool(leftVal && rightVal)
	case types.OROR:
		return nativeToBool(leftVal || rightVal)
	default:
		return object.Error{Message: console.Format(constants.IR_INVALID_INFIX, left.Type(), operator, right.Type()), Line: line}
	}
}
