// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalProgramNodes(programs []ast.Node, env *object.Env) object.Error {
	for _, program := range programs {
		var result object.Error = evalProgramNode(program, env)
		if result.GetIsFatal() {
			return result
		}
	}
	return Nil(0)
}

func evalProgramNode(program ast.Node, env *object.Env) object.Error {
	var result object.Object = Eval(program, env)
	if err, ok := result.(object.Error); ok && err.GetIsFatal() {
		return err
	} else {
		return Nil(program.GetLine())
	}
}

func evalMainFunction(env *object.Env) object.Object {
	if mainFunction, ok := env.Get(constants.MAIN_FUNCTION_NAME); ok {
		return callFunction(mainFunction, []object.Object{}, 0)
	} else {
		return newError(false, 0, constants.ERROR_REFERENCE, constants.IR_MAIN_FN_NOT_FOUND)
	}
}

func Eval(node ast.Node, env *object.Env) object.Object {
	switch node := node.(type) {
	case ast.Program:
		return evalProgram(node, env)
	case ast.Block:
		return evalBlock(node, env)
	case ast.Return:
		return evalReturn(node, env)
	case ast.Assign:
		return evalAssign(node, env)
	case ast.Call:
		return evalCall(node, env)
	case ast.Function:
		return evalFunction(node, env)
	case ast.Identifier:
		return evalIdentifier(node, env)
	case ast.Builtin:
		return evalBuiltin(node, env)
	case ast.Integer:
		return evalInteger(node, env)
	case ast.Float:
		return evalFloat(node, env)
	case ast.String:
		return evalString(node, env)
	case ast.Throw:
		return evalThrow(node, env)
	case ast.Prefix:
		return evalPrefixNode(node, env)
	case ast.Infix:
		return evalInfixNode(node, env)
	case ast.Postfix:
		return evalPostfixNode(node, env)
	case ast.Index:
		return evalIndex(node, env)
	case ast.If:
		return evalIf(node, env)
	case ast.Switch:
		return evalSwitch(node, env)
	case ast.Loop:
		return evalLoop(node, env)
	default:
		return newError(true, node.GetLine(), constants.ERROR_INTERTAL, constants.IR_INVALID_EVAL, node.GetType())
	}
}
