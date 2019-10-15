// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"path/filepath"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/io/file"
	"github.com/i5/i5/src/lexer"
	"github.com/i5/i5/src/object"
	"github.com/i5/i5/src/parser"
)

var (
	TRUE  = object.Bool{Value: true}
	FALSE = object.Bool{Value: false}
)

func RunPackage(dir string, arguments []string) {
	dir = filepath.Base(dir)
	filesToRun := file.GetFilesToRun(dir)
	env := object.InitEnv()

	for _, f := range filesToRun {
		fullPath := console.Format("%v/%v", dir, f)
		err := Eval(parser.Run(lexer.Run(file.Read(fullPath))), env)
		if err.Type() == object.ERROR {
			console.ThrowError(1, err.StringValue())
			return
		}
	}

	if mainFunction, ok := env.Get(constants.MAIN_FUNCTION_NAME); ok {
		result := callFunction(mainFunction, []object.Object{}, 0)
		if result.Type() == object.ERROR {
			console.ThrowError(1, result.StringValue())
			return
		}
	} else {
		console.ThrowError(1, constants.IR_MAIN_FN_NOT_FOUND)
	}
}

func RunModule(module string, arguments []string) {
	console.ThrowError(1, "not implemented yet")
	// TODO
}

func RunFile(program ast.Node, arguments []string) {
	env := object.InitEnv()
	err := Eval(program, env)
	if err.Type() == object.ERROR {
		console.ThrowError(1, err.StringValue())
		return
	}
	if mainFunction, ok := env.Get(constants.MAIN_FUNCTION_NAME); ok {
		result := callFunction(mainFunction, []object.Object{}, 0)
		if result.Type() == object.ERROR {
			console.ThrowError(1, result.StringValue())
			return
		}
	} else {
		console.ThrowError(1, constants.IR_MAIN_FN_NOT_FOUND)
	}
}

func isError(obj object.Object) bool {
	return obj.Type() == object.ERROR
}

func isVoid(obj object.Object) bool {
	return obj.Type() == object.VOID
}

func nativeToBool(input bool) object.Bool {
	if input {
		return TRUE
	}
	return FALSE
}

func isTrue(obj object.Object) bool {
	if obj == TRUE {
		return true
	} else if obj == FALSE {
		return false
	} else {
		return false
	}
}
