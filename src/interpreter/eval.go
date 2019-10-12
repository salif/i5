// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/object"
	"github.com/i5/i5/src/types"
)

func Eval(nodei ast.Node, env *object.Env) object.Object {
	switch node := nodei.(type) {

	case *ast.Program:
		return evalProgram(node, env, node.GetLine())

	case *ast.Expr:
		return Eval(node.Body, env)

	case *ast.Block:
		return evalBlock(node, env, node.GetLine())

	case *ast.Return:
		val := Eval(node.Body, env)
		if isError(val) {
			return val
		}
		if isVoid(val) {
			return &object.Error{Message: val.StringValue(), Line: node.GetLine()}
		}
		return &object.Return{Value: val}

	case *ast.Assign:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		if ident, ok := node.Left.(*ast.Identifier); ok {
			if isVoid(right) {
				(right.(*object.Void)).Value = ident.StringValue()
			}
			env.Set(ident.Value, right)
			return right
		} else {
			return &object.Error{Message: console.Format(constants.IR_CANNOT_ASSIGN, node.Left.StringValue()), Line: node.GetLine()}
		}
	case *ast.Call:
		function := Eval(node.Caller, env)
		if isError(function) {
			return function
		}
		if isVoid(function) {
			return &object.Error{Message: function.StringValue(), Line: node.GetLine()}
		}
		args := evalExpressions(node.Arguments, env, node.GetLine())
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		if len(args) == 1 && isVoid(args[0]) {
			return &object.Error{Message: args[0].StringValue(), Line: node.GetLine()}
		}
		return callFunction(function, args, node.GetLine())
	case *ast.Function:
		return &object.Function{Params: node.Params, Body: node.Body, Env: env}
	case *ast.Identifier:
		return evalIdentifier(node, env, node.GetLine())
	case *ast.Builtin:
		return evalBuiltin(node, env, node.GetLine())
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
		return &object.Error{Message: val.StringValue()}
	case *ast.Prefix:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		if isVoid(right) {
			return &object.Error{Message: right.StringValue(), Line: node.GetLine()}
		}
		return evalPrefix(node.Operator, right, env, node.GetLine())
	case *ast.Infix:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		if isVoid(left) {
			return &object.Error{Message: left.StringValue(), Line: node.GetLine()}
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		if isVoid(right) {
			return &object.Error{Message: right.StringValue(), Line: node.GetLine()}
		}
		return evalInfix(node.Operator, left, right, env, node.GetLine())
	case *ast.Suffix:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		if node.Operator == types.QM {
			return nativeToBool(left.Type() != object.VOID)
		}
		if isVoid(left) {
			return &object.Error{Message: left.StringValue(), Line: node.GetLine()}
		}
		return evalSuffix(node.Operator, left, env, node.GetLine())
	case *ast.If:
		return evalIf(node, env, node.GetLine())
	case *ast.Switch:
		return evalSwitch(node, env, node.GetLine())
	case *ast.While:
		return evalWhile(node, env, node.GetLine())
	case *ast.Import:
		return evalImport(node, env, node.GetLine())
	case *ast.Try:
		return evalTry(node, env, node.GetLine())
	case *ast.Break:
		return &object.Break{}
	case *ast.Continue:
		return &object.Continue{}
	default:
		return &object.Error{Message: console.Format(constants.IR_INVALID_EVAL, node.StringValue()), Line: node.GetLine()}
	}
}
