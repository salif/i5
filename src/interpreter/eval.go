// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"fmt"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalMainFunction(env *object.Env) (object.Object, error) {
	if mainFunction, ok := env.Get(constants.MAIN_FUNCTION_NAME); ok {
		return callFunction(mainFunction, []object.Object{}, 0)
	} else {
		return Null, constants.Error{Line: 0, Type: constants.ERROR_FATAL, Message: constants.IR_MAIN_FN_NOT_FOUND}
	}
}

func Eval(node ast.Node, env *object.Env) (object.Object, error) {
	switch node := node.(type) {

	// statements
	case ast.Program:
		return Null, evalProgram(node, env)
	case ast.Block:
		return Null, evalBlock(node, env)
	case ast.Function:
		return Null, evalFunction(node, env)
	case ast.Return:
		return Null, evalReturn(node, env)
	case ast.If:
		return Null, evalIf(node, env)
	case ast.Switch:
		return Null, evalSwitch(node, env)
	case ast.Loop:
		return Null, evalLoop(node, env)
	case ast.Break:
		return Null, evalBreak(node, env)
	case ast.Throw:
		return Null, evalThrow(node, env)

	// expressions
	case ast.Integer:
		return object.Integer{Value: node.GetValue()}, nil
	case ast.Float:
		return object.Float{Value: node.GetValue()}, nil
	case ast.String:
		return object.String{Value: node.GetValue()}, nil
	case ast.Identifier:
		return evalIdentifier(node, env)
	case ast.Builtin:
		return evalBuiltin(node, env)
	case ast.Assign:
		return evalAssign(node, env)
	case ast.Call:
		return evalCall(node, env)
	case ast.FunctionExpr:
		return object.Function{Params: node.GetParams(), Body: node.GetBody(), Env: env}, nil
	case ast.Prefix:
		return evalPrefixNode(node, env)
	case ast.Infix:
		return evalInfixNode(node, env)
	case ast.Postfix:
		return evalPostfixNode(node, env)
	case ast.Index:
		return evalIndex(node, env)

	default:
		return Null, constants.Error{Line: node.GetLine(), Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_INVALID_EVAL, node.GetType())}
	}
}
