// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/object"
)

func Eval(nodei ast.Node, env *object.Env) object.Object {
	switch node := nodei.(type) {

	case *ast.Program:
		var ret object.Object
		for _, expr := range node.Body {
			ret = Eval(expr, env)
			switch ret := ret.(type) {
			case *object.Error:
				return ret
			}
		}

		Eval(&ast.Call{Caller: &ast.Identifier{Value: "main"}, Arguments: []ast.Expression{}}, env)
		// TODO errors.FatalError("main function not found", 1)
		return ret

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
		result := Eval(node.Right, env)
		switch left := node.Left.(type) {
		case *ast.Identifier:
			env.Set(left.Value, result)
		default:
			console.ThrowError(1, "left assign error")
		}
		return nil
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
	case *ast.If:
		return evalIf(node, env)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.Builtin:
		return evalBuiltin(node, env)
	case *ast.Number:
		return &object.Number{Value: node.Value}
	case *ast.String:
		return &object.String{Value: node.Value}
	case *ast.Bool:
		return nativeToBool(node.Value)
	case *ast.Nil:
		return NIL
	case *ast.Prefix:
		right := Eval(node.Right, env)
		return evalPrefix(node.Operator, right)
	case *ast.Infix:
		left := Eval(node.Left, env)
		right := Eval(node.Right, env)
		return evalInfix(node.Operator, left, right)
	case *ast.Suffix:
		left := Eval(node.Left, env)
		return evalSuffix(node.Operator, left)
	default:
		console.Println("eval error", node)
	}
	return nil
}
