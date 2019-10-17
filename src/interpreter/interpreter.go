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

func RunDirectory(directoryName string) {
	absoluteDirectoryName, _ := filepath.Abs(directoryName)
	modFile := console.Format("%s/%s", absoluteDirectoryName, constants.I5_MOD_FILE_NAME)
	if file.Exists(modFile) {
		RunModule(modFile, absoluteDirectoryName)
	} else {
		RunPackage(absoluteDirectoryName)
	}
}

func RunPackage(absoluteDirectoryName string) {
	env := object.InitEnv()
	EvalPackage(absoluteDirectoryName, env)
	EvalMainFunction(env)
}

func RunModule(moduleFileName, module string) {
	f := file.Read(moduleFileName)
	moduleInfoArray := lexer.ParseModuleFile(f)
	if len(moduleInfoArray) < 2 {
		console.ThrowError(1, constants.IR_INVALID_MOD_FILE)
	}
}

func RunFile(code []byte) {
	program := parser.Run(lexer.Run(code))
	env := object.InitEnv()
	EvalFile(program, env)
	EvalMainFunction(env)
}

func EvalPackage(absoluteDirectoryName string, env *object.Env) {
	filesToRun := file.GetFilesToRun(absoluteDirectoryName)
	for _, fileToRun := range filesToRun {
		fileToRunPath := console.Format("%v/%v", absoluteDirectoryName, fileToRun)
		EvalFile(parser.Run(lexer.Run(file.Read(fileToRunPath))), env)
	}
}

func EvalFile(program ast.Node, env *object.Env) {
	err := Eval(program, env)
	if err.Type() == object.ERROR {
		console.ThrowError(1, err.StringValue())
		return
	}
}

func EvalMainFunction(env *object.Env) {
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
