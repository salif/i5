// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"fmt"
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalInfixNode(node ast.Infix, env *object.Env) (object.Object, error) {
	if node.GetOperator() == constants.TOKEN_QM {
		return evalIsError(node, env)
	}

	evLeft, err := Eval(node.GetLeft(), env)
	if err != nil {
		return nil, err
	}

	evRight, err := Eval(node.GetRight(), env)
	if err != nil {
		return nil, err
	}

	return evalInfix(node.GetOperator(), evLeft, evRight, env, node.GetLine())
}

func evalInfix(operator string, left, right object.Object, env *object.Env, line uint32) (object.Object, error) {
	if operator == constants.TOKEN_COLON {
		return object.String{Value: left.StringValue() + right.StringValue()}, nil
	} else if left.Type() == right.Type() && left.Type() == constants.TYPE_INTEGER {
		return evalIntegerWithIntegerInfix(operator, left, right, line)
	} else if left.Type() == right.Type() && left.Type() == constants.TYPE_FLOAT {
		return evalFloatWithFloatInfix(operator, left, right, line)
	} else if left.Type() == constants.TYPE_INTEGER && right.Type() == constants.TYPE_FLOAT {
		return evalIntegerWithFloatInfix(operator, left, right, line)
	} else if left.Type() == constants.TYPE_FLOAT && right.Type() == constants.TYPE_INTEGER {
		return evalFloatWithIntegerInfix(operator, left, right, line)
	} else if left.Type() == right.Type() && left.Type() == constants.TYPE_STRING {
		return evalStringInfix(operator, left, right, line)
	} else if left.Type() == right.Type() && left.Type() == constants.TYPE_BOOLEAN {
		return evalBooleanInfix(operator, left, right, line)
	}
	return nil, constants.Error{Line: line, Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_INVALID_INFIX, left.Type(), operator, right.Type())}
}

func evalIntegerWithIntegerInfix(operator string, left, right object.Object, line uint32) (object.Object, error) {
	var leftInteger int64 = left.(object.Integer).Value
	var rightInteger int64 = right.(object.Integer).Value

	switch operator {
	case constants.TOKEN_PLUS:
		return object.Integer{Value: leftInteger + rightInteger}, nil
	case constants.TOKEN_MINUS:
		return object.Integer{Value: leftInteger - rightInteger}, nil
	case constants.TOKEN_MULTIPLY:
		return object.Integer{Value: leftInteger * rightInteger}, nil
	case constants.TOKEN_DIVIDE:
		return object.Integer{Value: leftInteger / rightInteger}, nil
	case constants.TOKEN_MODULO:
		return object.Integer{Value: leftInteger % rightInteger}, nil
	case constants.TOKEN_OR:
		return object.Integer{Value: leftInteger | rightInteger}, nil
	case constants.TOKEN_XOR:
		return object.Integer{Value: leftInteger ^ rightInteger}, nil
	case constants.TOKEN_AND:
		return object.Integer{Value: leftInteger & rightInteger}, nil
	case constants.TOKEN_LTLT:
		return object.Integer{Value: leftInteger << uint64(rightInteger)}, nil
	case constants.TOKEN_GTGT:
		return object.Integer{Value: leftInteger >> uint64(rightInteger)}, nil
	case constants.TOKEN_LT:
		return nativeToBool(leftInteger < rightInteger), nil
	case constants.TOKEN_LTEQ:
		return nativeToBool(leftInteger <= rightInteger), nil
	case constants.TOKEN_GT:
		return nativeToBool(leftInteger > rightInteger), nil
	case constants.TOKEN_GTEQ:
		return nativeToBool(leftInteger >= rightInteger), nil
	case constants.TOKEN_EQEQ:
		return nativeToBool(leftInteger == rightInteger), nil
	case constants.TOKEN_NOTEQ:
		return nativeToBool(leftInteger != rightInteger), nil
	default:
		return nil, constants.Error{Line: line, Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_INVALID_INFIX, left.Type(), operator, right.Type())}
	}
}

func evalFloatWithFloatInfix(operator string, left, right object.Object, line uint32) (object.Object, error) {
	var leftFloat float64 = left.(object.Float).Value
	var rightFloat float64 = right.(object.Float).Value

	switch operator {
	case constants.TOKEN_PLUS:
		return object.Float{Value: leftFloat + rightFloat}, nil
	case constants.TOKEN_MINUS:
		return object.Float{Value: leftFloat - rightFloat}, nil
	case constants.TOKEN_MULTIPLY:
		return object.Float{Value: leftFloat * rightFloat}, nil
	case constants.TOKEN_DIVIDE:
		return object.Float{Value: leftFloat / rightFloat}, nil
	case constants.TOKEN_LT:
		return nativeToBool(leftFloat < rightFloat), nil
	case constants.TOKEN_LTEQ:
		return nativeToBool(leftFloat <= rightFloat), nil
	case constants.TOKEN_GT:
		return nativeToBool(leftFloat > rightFloat), nil
	case constants.TOKEN_GTEQ:
		return nativeToBool(leftFloat >= rightFloat), nil
	case constants.TOKEN_EQEQ:
		return nativeToBool(leftFloat == rightFloat), nil
	case constants.TOKEN_NOTEQ:
		return nativeToBool(leftFloat != rightFloat), nil
	default:
		return nil, constants.Error{Line: line, Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_INVALID_INFIX, left.Type(), operator, right.Type())}
	}
}

func evalIntegerWithFloatInfix(operator string, left, right object.Object, line uint32) (object.Object, error) {
	var leftInteger int64 = left.(object.Integer).Value
	var rightFloat float64 = right.(object.Float).Value

	switch operator {
	case constants.TOKEN_PLUS:
		return object.Float{Value: float64(leftInteger) + rightFloat}, nil
	case constants.TOKEN_MINUS:
		return object.Float{Value: float64(leftInteger) - rightFloat}, nil
	case constants.TOKEN_MULTIPLY:
		return object.Float{Value: float64(leftInteger) * rightFloat}, nil
	case constants.TOKEN_DIVIDE:
		return object.Float{Value: float64(leftInteger) / rightFloat}, nil
	case constants.TOKEN_LT:
		return nativeToBool(float64(leftInteger) < rightFloat), nil
	case constants.TOKEN_LTEQ:
		return nativeToBool(float64(leftInteger) <= rightFloat), nil
	case constants.TOKEN_GT:
		return nativeToBool(float64(leftInteger) > rightFloat), nil
	case constants.TOKEN_GTEQ:
		return nativeToBool(float64(leftInteger) >= rightFloat), nil
	case constants.TOKEN_EQEQ:
		return nativeToBool(float64(leftInteger) == rightFloat), nil
	case constants.TOKEN_NOTEQ:
		return nativeToBool(float64(leftInteger) != rightFloat), nil
	default:
		return nil, constants.Error{Line: line, Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_INVALID_INFIX, left.Type(), operator, right.Type())}
	}
}

func evalFloatWithIntegerInfix(operator string, left, right object.Object, line uint32) (object.Object, error) {
	var leftFloat float64 = left.(object.Float).Value
	var rightInteger int64 = right.(object.Integer).Value

	switch operator {
	case constants.TOKEN_PLUS:
		return object.Float{Value: leftFloat + float64(rightInteger)}, nil
	case constants.TOKEN_MINUS:
		return object.Float{Value: leftFloat - float64(rightInteger)}, nil
	case constants.TOKEN_MULTIPLY:
		return object.Float{Value: leftFloat * float64(rightInteger)}, nil
	case constants.TOKEN_DIVIDE:
		return object.Float{Value: leftFloat / float64(rightInteger)}, nil
	case constants.TOKEN_LT:
		return nativeToBool(leftFloat < float64(rightInteger)), nil
	case constants.TOKEN_LTEQ:
		return nativeToBool(leftFloat <= float64(rightInteger)), nil
	case constants.TOKEN_GT:
		return nativeToBool(leftFloat > float64(rightInteger)), nil
	case constants.TOKEN_GTEQ:
		return nativeToBool(leftFloat >= float64(rightInteger)), nil
	case constants.TOKEN_EQEQ:
		return nativeToBool(leftFloat == float64(rightInteger)), nil
	case constants.TOKEN_NOTEQ:
		return nativeToBool(leftFloat != float64(rightInteger)), nil
	default:
		return nil, constants.Error{Line: line, Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_INVALID_INFIX, left.Type(), operator, right.Type())}
	}
}

func evalStringInfix(operator string, left, right object.Object, line uint32) (object.Object, error) {
	var leftString string = left.(object.String).Value
	var rightString string = right.(object.String).Value

	switch operator {
	case constants.TOKEN_EQEQ:
		return nativeToBool(leftString == rightString), nil
	case constants.TOKEN_NOTEQ:
		return nativeToBool(leftString != rightString), nil
	default:
		return nil, constants.Error{Line: line, Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_INVALID_INFIX, left.Type(), operator, right.Type())}
	}
}

func evalBooleanInfix(operator string, left, right object.Object, line uint32) (object.Object, error) {
	var leftBoolean bool = left.(object.Boolean).Value
	var rightBoolean bool = right.(object.Boolean).Value

	switch operator {
	case constants.TOKEN_EQEQ:
		return nativeToBool(leftBoolean == rightBoolean), nil
	case constants.TOKEN_NOTEQ:
		return nativeToBool(leftBoolean != rightBoolean), nil
	case constants.TOKEN_ANDAND:
		return nativeToBool(leftBoolean && rightBoolean), nil
	case constants.TOKEN_OROR:
		return nativeToBool(leftBoolean || rightBoolean), nil
	default:
		return nil, constants.Error{Line: line, Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_INVALID_INFIX, left.Type(), operator, right.Type())}
	}
}

func evalIsError(node ast.Infix, env *object.Env) (object.Object, error) {
	evRight, err := Eval(node.GetRight(), env)
	if err != nil {
		return nil, err
	}

	if evRight.Type() != constants.TYPE_EXCEPTION {
		return FALSE, nil
	} else {
		left := node.GetLeft()
		if left.GetType() == constants.TOKEN_IDENTIFIER {
			ident := left.(ast.Identifier)
			env.Set(ident.GetValue(), evRight)
			return TRUE, nil
		} else {
			return nil, constants.Error{Line: node.GetLine(), Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_CANNOT_ASSIGN, left.GetType())}
		}
	}
}
