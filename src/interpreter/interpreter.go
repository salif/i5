// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/builtins"
	"github.com/i5/i5/src/errors"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/object"
)

var (
	TRUE  = &object.Bool{Value: true}
	FALSE = &object.Bool{Value: false}
	NIL   = &object.Nil{}
)

func Run(program ast.Node) {
	Eval(program, object.InitEnv())
}

func Eval(nodei ast.Node, env *object.Env) object.Object {
	switch node := nodei.(type) {

	case *ast.Program:
		var ret object.Object
		for _, expr := range node.Body {
			ret = Eval(expr, env)
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
		return &object.Return{Value: val}

	case *ast.Assign:
		result := Eval(node.Right, env)
		switch left := node.Left.(type) {
		case *ast.ExprList:
			env.Set(left.Body[0].String(), result)
		default:
			errors.FatalError("left assign: expected ast.ExprList", 1)
		}
		return nil
	case *ast.Call:
		function := Eval(node.Caller, env)
		args := evalExpressions(node.Arguments, env)
		return callFunction(function, args)
	case *ast.Function:
		return &object.Function{Params: node.Params, Body: node.Body, Env: env}
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.Builtin:
		return evalBuiltin(node, env)
	case *ast.Number:
		return &object.Number{Value: node.Value}
	case *ast.String:
		return &object.String{Value: node.Value}
	case *ast.Bool:
		return &object.Bool{Value: node.Value}
	case *ast.Nil:
		return &object.Nil{}
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
		// TODO change this
	}
	return nil
}

func evalIdentifier(node *ast.Identifier, env *object.Env) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	} else {
		errors.FatalError(errors.F("%v: identifier not found", node.Value), 1)
		return nil
	}
}

func evalBuiltin(node *ast.Builtin, env *object.Env) object.Object {
	if builtin, ok := builtins.Get(node.Value); ok {
		return builtin
	} else {
		errors.FatalError(errors.F("%v: builtin not found", node.Value), 1)
		return nil
	}
}

func callFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {

	case *object.Function:
		env := extendFunctionEnv(fn, args)
		return unwrapReturnValue(Eval(fn.Body, env))
	case *object.Builtin:
		if result := fn.Function(args...); result != nil {
			return result
		}
		return NIL
	default:
		console.Println("not function", fn.Type())
	}
	return NIL
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Env {
	env := fn.Env.Clone()

	for paramIdx, param := range fn.Params {
		env.Set(param.Value, args[paramIdx])
	}

	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.Return); ok {
		return returnValue.Value
	}

	return obj
}

func evalBlock(block *ast.Block, env *object.Env) object.Object {
	var result object.Object
	for _, statement := range block.Body {
		result = Eval(statement, env)
		if result != nil {
			rt := result.Type()
			if rt == object.RETURN || rt == object.ERROR {
				return result
			}
		}
	}
	return result
}

func evalExpressions(exps []ast.Expression, env *object.Env) []object.Object {
	var result []object.Object
	for _, e := range exps {
		evaluated := Eval(e, env)
		result = append(result, evaluated)
	}
	return result
}

func evalPrefix(operator string, right object.Object) object.Object {
	return &object.Error{}
}

func evalInfix(operator string, left, right object.Object) object.Object {
	return &object.Error{}
}

func evalSuffix(operator string, left object.Object) object.Object {
	return &object.Error{}
}
