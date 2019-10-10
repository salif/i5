// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func Eval(nodei ast.Node, env *object.Env) object.Object {
	switch node := nodei.(type) {

	case *ast.Program:
		prog := evalProgram(node, env)
		if isError(prog) {
			return prog
		}
		return Eval(&ast.Call{Caller: &ast.Identifier{Value: "main"}, Arguments: []ast.Expression{}}, env)

	case *ast.Expr:
		return Eval(node.Body, env)

	case *ast.Block:
		return evalBlock(node, env)

	case *ast.Return:
		val := Eval(node.Body, env)
		if isError(val) {
			return val
		}
		return &object.Return{Value: val}

	case *ast.Assign:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		if ident, ok := node.Left.(*ast.Identifier); ok {
			env.Set(ident.Value, right)
			return right
		} else {
			return newError("can not assign to %v", node.Left.StringValue())
		}
	case *ast.Call:
		function := Eval(node.Caller, env)
		if isError(function) {
			return function
		}
		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		return callFunction(function, args)
	case *ast.Function:
		return &object.Function{Params: node.Params, Body: node.Body, Env: env}
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.Builtin:
		return evalBuiltin(node, env)
	case *ast.Integer:
		return &object.Integer{Value: node.Value}
	case *ast.Float:
		return &object.Float{Value: node.Value}
	case *ast.String:
		return &object.String{Value: node.Value}
	case *ast.Bool:
		return nativeToBool(node.Value)
	case *ast.Throw:
		val := Eval(node.Body, env)
		if isError(val) {
			return val
		}
		return &object.Throw{Value: val}
	case *ast.Prefix:
		right := Eval(node.Right, env)
		return evalPrefix(node.Operator, right, env)
	case *ast.Infix:
		left := Eval(node.Left, env)
		right := Eval(node.Right, env)
		return evalInfix(node.Operator, left, right, env)
	case *ast.Suffix:
		left := Eval(node.Left, env)
		return evalSuffix(node.Operator, left, env)
	case *ast.If:
		return evalIf(node, env)
	case *ast.Switch:
		return evalSwitch(node, env)
	case *ast.While:
		return evalWhile(node, env)
	case *ast.ImportExpr:
		return evalImportExpr(node, env)
	case *ast.ImportStatement:
		return evalImportStatement(node, env)
	case *ast.Try:
		return evalTry(node, env)
	default:
		return newError(constants.IR_INVALID_EVAL, node.StringValue())
	}
}
