// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/object"
)

func evalInfix(operator string, left, right object.Object) object.Object {
	if left.Type() == right.Type() && left.Type() == object.NUMBER {
		return evalNumberInfix(operator, left, right)
	} else if left.Type() == right.Type() && left.Type() == object.STRING {
		return evalStringInfix(operator, left, right)
	}
	return NIL
}

func evalNumberInfix(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Number).Value
	rightVal := right.(*object.Number).Value

	switch operator {
	case "+":
		return &object.Number{Value: leftVal + rightVal}
	case "-":
		return &object.Number{Value: leftVal - rightVal}
	case "*":
		return &object.Number{Value: leftVal * rightVal}
	case "/":
		return &object.Number{Value: leftVal / rightVal}
	case "%":
		return &object.Number{Value: leftVal % rightVal}
	case "|":
		return &object.Number{Value: leftVal | rightVal}
	case "^":
		return &object.Number{Value: leftVal ^ rightVal}
	case "&":
		return &object.Number{Value: leftVal & rightVal}
	case "<<":
		return &object.Number{Value: leftVal << uint64(rightVal)}
	case ">>":
		return &object.Number{Value: leftVal >> uint64(rightVal)}
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
