// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/object"
)

func evalInfix(operator string, left, right object.Object) object.Object {
	if operator == ":" {
		return &object.String{Value: left.StringValue() + right.StringValue()}
	} else if left.Type() == right.Type() && left.Type() == object.INTEGER {
		return evalIntegerWithIntegerInfix(operator, left, right)
	} else if left.Type() == right.Type() && left.Type() == object.FLOAT {
		return evalFloatWithFloatInfix(operator, left, right)
	} else if left.Type() == object.INTEGER && right.Type() == object.FLOAT {
		return evalIntegerWithFloatInfix(operator, left, right)
	} else if left.Type() == object.FLOAT && right.Type() == object.INTEGER {
		return evalFloatWithIntegerInfix(operator, left, right)
	} else if left.Type() == right.Type() && left.Type() == object.STRING {
		return evalStringInfix(operator, left, right)
	}
	return NIL
}

func evalIntegerWithIntegerInfix(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "%":
		return &object.Integer{Value: leftVal % rightVal}
	case "|":
		return &object.Integer{Value: leftVal | rightVal}
	case "^":
		return &object.Integer{Value: leftVal ^ rightVal}
	case "&":
		return &object.Integer{Value: leftVal & rightVal}
	case "<<":
		return &object.Integer{Value: leftVal << uint64(rightVal)}
	case ">>":
		return &object.Integer{Value: leftVal >> uint64(rightVal)}
	case "<":
		return nativeToBool(leftVal < rightVal)
	case "<=":
		return nativeToBool(leftVal <= rightVal)
	case ">":
		return nativeToBool(leftVal > rightVal)
	case ">=":
		return nativeToBool(leftVal >= rightVal)
	case "==":
		return nativeToBool(leftVal == rightVal)
	case "!=":
		return nativeToBool(leftVal != rightVal)
	default:
		return NIL
	}
}

func evalFloatWithFloatInfix(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Float).Value
	rightVal := right.(*object.Float).Value

	switch operator {
	case "+":
		return &object.Float{Value: leftVal + rightVal}
	case "-":
		return &object.Float{Value: leftVal - rightVal}
	case "*":
		return &object.Float{Value: leftVal * rightVal}
	case "/":
		return &object.Float{Value: leftVal / rightVal}
	case "<":
		return nativeToBool(leftVal < rightVal)
	case "<=":
		return nativeToBool(leftVal <= rightVal)
	case ">":
		return nativeToBool(leftVal > rightVal)
	case ">=":
		return nativeToBool(leftVal >= rightVal)
	case "==":
		return nativeToBool(leftVal == rightVal)
	case "!=":
		return nativeToBool(leftVal != rightVal)
	default:
		return NIL
	}
}

func evalIntegerWithFloatInfix(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Float).Value

	switch operator {
	case "+":
		return &object.Float{Value: float64(leftVal) + rightVal}
	case "-":
		return &object.Float{Value: float64(leftVal) - rightVal}
	case "*":
		return &object.Float{Value: float64(leftVal) * rightVal}
	case "/":
		return &object.Float{Value: float64(leftVal) / rightVal}
	case "<":
		return nativeToBool(float64(leftVal) < rightVal)
	case "<=":
		return nativeToBool(float64(leftVal) <= rightVal)
	case ">":
		return nativeToBool(float64(leftVal) > rightVal)
	case ">=":
		return nativeToBool(float64(leftVal) >= rightVal)
	case "==":
		return nativeToBool(float64(leftVal) == rightVal)
	case "!=":
		return nativeToBool(float64(leftVal) != rightVal)
	default:
		return NIL
	}
}

func evalFloatWithIntegerInfix(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Float).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Float{Value: leftVal + float64(rightVal)}
	case "-":
		return &object.Float{Value: leftVal - float64(rightVal)}
	case "*":
		return &object.Float{Value: leftVal * float64(rightVal)}
	case "/":
		return &object.Float{Value: leftVal / float64(rightVal)}
	case "<":
		return nativeToBool(leftVal < float64(rightVal))
	case "<=":
		return nativeToBool(leftVal <= float64(rightVal))
	case ">":
		return nativeToBool(leftVal > float64(rightVal))
	case ">=":
		return nativeToBool(leftVal >= float64(rightVal))
	case "==":
		return nativeToBool(leftVal == float64(rightVal))
	case "!=":
		return nativeToBool(leftVal != float64(rightVal))
	default:
		return NIL
	}
}

func evalStringInfix(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	switch operator {
	case "==":
		return nativeToBool(leftVal == rightVal)
	case "!=":
		return nativeToBool(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalBooleanInfix(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Bool).Value
	rightVal := right.(*object.Bool).Value

	switch operator {
	case "==":
		return nativeToBool(leftVal == rightVal)
	case "!=":
		return nativeToBool(leftVal != rightVal)
	case "&&":
		return nativeToBool(leftVal && rightVal)
	case "||":
		return nativeToBool(leftVal || rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}
