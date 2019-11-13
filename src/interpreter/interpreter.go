// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"fmt"
	"path/filepath"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/file"
	"github.com/i5/i5/src/lexer"
	"github.com/i5/i5/src/object"
	"github.com/i5/i5/src/parser"
)

func Run(fileOrDirectoryName string, args []string) error {
	fileOrDirectoryName, err := filepath.Abs(fileOrDirectoryName)
	if err != nil {
		return Errf(err)
	}
	var result int = file.Info(fileOrDirectoryName)
	switch result {
	case 1:
		return Errf(fmt.Errorf(constants.FILE_NOT_FOUND, fileOrDirectoryName))
	case 2:
		return runModule(fileOrDirectoryName)
	case 3:
		codeBytes, err := file.Read(fileOrDirectoryName)
		if err != nil {
			return Errf(err)
		}
		return runFile(fileOrDirectoryName, codeBytes)
	default:
		return nil
	}
}

func RunEval(content string, args []string) error {
	return runFile("evaluated code", []byte(content))
}

// TODO remove dublication
func runModule(absoluteDirectoryName string) error {
	var env *object.Env = object.InitEnv()
	filesToRun, err := file.GetFilesToRun(absoluteDirectoryName)
	if err != nil {
		return Errf(err)
	}
	programs, err := parsePrograms(filesToRun)
	if err != nil {
		return err
	}
	var evaluatedPrograms object.Error = evalProgramNodes(programs, env)
	if evaluatedPrograms.GetIsFatal() {
		// TODO edit fileName
		return evaluatedPrograms.NativeError(absoluteDirectoryName)
	}
	var evaluatedMainFunction object.Object = evalMainFunction(env)
	var errorType int = ErrorType(evaluatedMainFunction)
	if errorType == 2 || errorType == 3 {
		return evaluatedMainFunction.(object.Error).NativeError(absoluteDirectoryName)
	} else {
		return nil
	}

}

func runFile(fileName string, code []byte) error {
	var env *object.Env = object.InitEnv()
	program, err := parseProgram(fileName, code)
	if err != nil {
		return err
	}
	var evaluatedProgram object.Error = evalProgramNode(program, env)
	if evaluatedProgram.GetIsFatal() {
		return evaluatedProgram.NativeError(fileName)
	}
	var evaluatedMainFunction object.Object = evalMainFunction(env)
	var errorType int = ErrorType(evaluatedMainFunction)
	if errorType == 2 || errorType == 3 {
		return evaluatedMainFunction.(object.Error).NativeError(fileName)
	} else {
		return nil
	}
}

func parsePrograms(fileNames []string) ([]ast.Node, error) {
	var programs []ast.Node = []ast.Node{}
	for _, fileName := range fileNames {
		codeBytes, err := file.Read(fileName)
		if err != nil {
			return nil, Errf(err)
		}
		program, err := parseProgram(fileName, codeBytes)
		if err != nil {
			return nil, err
		}
		programs = append(programs, program)
	}
	return programs, nil
}

func parseProgram(fileName string, code []byte) (ast.Node, error) {
	tokens, err := lexer.Run(fileName, code)
	if err != nil {
		return nil, err
	}
	program, err2 := parser.Run(fileName, tokens)
	if err2 != nil {
		return nil, err2
	}
	return program, nil
}
