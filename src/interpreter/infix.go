// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
	"github.com/i5/i5/src/types"
)

func evalInfixNode(node ast.Infix, env *object.Env) object.Object {
	if node.GetOperator() == types.QM {
		return evalIsError(node, env)
	}
	var evaluatedLeft object.Object = Eval(node.GetLeft(), env)
	if ErrorType(evaluatedLeft) == FATAL {
		return evaluatedLeft
	}
	var evaluatedRight object.Object = Eval(node.GetRight(), env)
	if ErrorType(evaluatedRight) == FATAL {
		return evaluatedRight
	}
	return evalInfix(node.GetOperator(), evaluatedLeft, evaluatedRight, env, node.GetLine())
}

func evalInfix(operator string, left, right object.Object, env *object.Env, line uint32) object.Object {
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
	return newError(true, line, constants.ERROR_INTERTAL, constants.IR_INVALID_INFIX, left.Type(), operator, right.Type())
}

func evalIntegerWithIntegerInfix(operator string, left, right object.Object, line uint32) object.Object {
	var leftInteger int64 = left.(object.Integer).Value
	var rightInteger int64 = right.(object.Integer).Value

	switch operator {
	case types.PLUS:
		return object.Integer{Value: leftInteger + rightInteger}
	case types.MINUS:
		return object.Integer{Value: leftInteger - rightInteger}
	case types.MULTIPLY:
		return object.Integer{Value: leftInteger * rightInteger}
	case types.DIVIDE:
		return object.Integer{Value: leftInteger / rightInteger}
	case types.MODULO:
		return object.Integer{Value: leftInteger % rightInteger}
	case types.OR:
		return object.Integer{Value: leftInteger | rightInteger}
	case types.XOR:
		return object.Integer{Value: leftInteger ^ rightInteger}
	case types.AND:
		return object.Integer{Value: leftInteger & rightInteger}
	case types.LTLT:
		return object.Integer{Value: leftInteger << uint64(rightInteger)}
	case types.GTGT:
		return object.Integer{Value: leftInteger >> uint64(rightInteger)}
	case types.LT:
		return nativeToBool(leftInteger < rightInteger)
	case types.LTEQ:
		return nativeToBool(leftInteger <= rightInteger)
	case types.GT:
		return nativeToBool(leftInteger > rightInteger)
	case types.GTEQ:
		return nativeToBool(leftInteger >= rightInteger)
	case types.EQEQ:
		return nativeToBool(leftInteger == rightInteger)
	case types.NOTEQ:
		return nativeToBool(leftInteger != rightInteger)
	default:
		return newError(true, line, constants.ERROR_INTERTAL, constants.IR_INVALID_INFIX, left.Type(), operator, right.Type())
	}
}

func evalFloatWithFloatInfix(operator string, left, right object.Object, line uint32) object.Object {
	var leftFloat float64 = left.(object.Float).Value
	var rightFloat float64 = right.(object.Float).Value

	switch operator {
	case types.PLUS:
		return object.Float{Value: leftFloat + rightFloat}
	case types.MINUS:
		return object.Float{Value: leftFloat - rightFloat}
	case types.MULTIPLY:
		return object.Float{Value: leftFloat * rightFloat}
	case types.DIVIDE:
		return object.Float{Value: leftFloat / rightFloat}
	case types.LT:
		return nativeToBool(leftFloat < rightFloat)
	case types.LTEQ:
		return nativeToBool(leftFloat <= rightFloat)
	case types.GT:
		return nativeToBool(leftFloat > rightFloat)
	case types.GTEQ:
		return nativeToBool(leftFloat >= rightFloat)
	case types.EQEQ:
		return nativeToBool(leftFloat == rightFloat)
	case types.NOTEQ:
		return nativeToBool(leftFloat != rightFloat)
	default:
		return newError(true, line, constants.ERROR_INTERTAL, constants.IR_INVALID_INFIX, left.Type(), operator, right.Type())
	}
}

func evalIntegerWithFloatInfix(operator string, left, right object.Object, line uint32) object.Object {
	var leftInteger int64 = left.(object.Integer).Value
	var rightFloat float64 = right.(object.Float).Value

	switch operator {
	case types.PLUS:
		return object.Float{Value: float64(leftInteger) + rightFloat}
	case types.MINUS:
		return object.Float{Value: float64(leftInteger) - rightFloat}
	case types.MULTIPLY:
		return object.Float{Value: float64(leftInteger) * rightFloat}
	case types.DIVIDE:
		return object.Float{Value: float64(leftInteger) / rightFloat}
	case types.LT:
		return nativeToBool(float64(leftInteger) < rightFloat)
	case types.LTEQ:
		return nativeToBool(float64(leftInteger) <= rightFloat)
	case types.GT:
		return nativeToBool(float64(leftInteger) > rightFloat)
	case types.GTEQ:
		return nativeToBool(float64(leftInteger) >= rightFloat)
	case types.EQEQ:
		return nativeToBool(float64(leftInteger) == rightFloat)
	case types.NOTEQ:
		return nativeToBool(float64(leftInteger) != rightFloat)
	default:
		return newError(true, line, constants.ERROR_INTERTAL, constants.IR_INVALID_INFIX, left.Type(), operator, right.Type())
	}
}

func evalFloatWithIntegerInfix(operator string, left, right object.Object, line uint32) object.Object {
	var leftFloat float64 = left.(object.Float).Value
	var rightInteger int64 = right.(object.Integer).Value

	switch operator {
	case types.PLUS:
		return object.Float{Value: leftFloat + float64(rightInteger)}
	case types.MINUS:
		return object.Float{Value: leftFloat - float64(rightInteger)}
	case types.MULTIPLY:
		return object.Float{Value: leftFloat * float64(rightInteger)}
	case types.DIVIDE:
		return object.Float{Value: leftFloat / float64(rightInteger)}
	case types.LT:
		return nativeToBool(leftFloat < float64(rightInteger))
	case types.LTEQ:
		return nativeToBool(leftFloat <= float64(rightInteger))
	case types.GT:
		return nativeToBool(leftFloat > float64(rightInteger))
	case types.GTEQ:
		return nativeToBool(leftFloat >= float64(rightInteger))
	case types.EQEQ:
		return nativeToBool(leftFloat == float64(rightInteger))
	case types.NOTEQ:
		return nativeToBool(leftFloat != float64(rightInteger))
	default:
		return newError(true, line, constants.ERROR_INTERTAL, constants.IR_INVALID_INFIX, left.Type(), operator, right.Type())
	}
}

func evalStringInfix(operator string, left, right object.Object, line uint32) object.Object {
	var leftString string = left.(object.String).Value
	var rightString string = right.(object.String).Value

	switch operator {
	case types.EQEQ:
		return nativeToBool(leftString == rightString)
	case types.NOTEQ:
		return nativeToBool(leftString != rightString)
	default:
		return newError(true, line, constants.ERROR_INTERTAL, constants.IR_INVALID_INFIX, left.Type(), operator, right.Type())
	}
}

func evalBooleanInfix(operator string, left, right object.Object, line uint32) object.Object {
	var leftBoolean bool = left.(object.Bool).Value
	var rightBoolean bool = right.(object.Bool).Value

	switch operator {
	case types.EQEQ:
		return nativeToBool(leftBoolean == rightBoolean)
	case types.NOTEQ:
		return nativeToBool(leftBoolean != rightBoolean)
	case types.ANDAND:
		return nativeToBool(leftBoolean && rightBoolean)
	case types.OROR:
		return nativeToBool(leftBoolean || rightBoolean)
	default:
		return newError(true, line, constants.ERROR_INTERTAL, constants.IR_INVALID_INFIX, left.Type(), operator, right.Type())
	}
}

func evalIsError(node ast.Infix, env *object.Env) object.Object {
	var evaluatedRight object.Object = Eval(node.GetRight(), env)
	e := ErrorType(evaluatedRight)
	if e == FATAL {
		return evaluatedRight
	}
	if e == 0 {
		return FALSE
	} else {
		left := node.GetLeft()
		if left.GetType() == types.IDENT {
			ident := left.(ast.Identifier)
			env.Set(ident.GetValue(), evaluatedRight)
			return TRUE
		} else {
			return newError(true, node.GetLine(), constants.ERROR_INTERTAL, constants.IR_CANNOT_ASSIGN, left.GetType())
		}
	}
}
