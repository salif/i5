// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package interpreter

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/file"
	"github.com/i5/i5/src/object"
	"github.com/i5/i5/src/parser"
)

type Interpreter struct {
	args []string
	p    string
}

func New(args []string, p string) Interpreter {
	this := Interpreter{}
	this.args = args
	this.p = p
	return this
}

func (this Interpreter) Run() error {
	fileOrDirectoryName, err := filepath.Abs(this.args[0])
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

func (this Interpreter) RunEval(content string) error {
	return runFile("evaluated code", []byte(content))
}

func (this Interpreter) RunRepl() error {
	fmt.Printf("i5 v%v\n", constants.PATCH_VERSION)
	reader := bufio.NewReader(os.Stdin)
	var env *object.Env = object.InitEnv()
	var fileName string = "REPL"

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println()
			return Errf(err)
		}
		if len(strings.TrimSpace(input)) == 0 {
			continue
		}
		var code []byte = []byte(input)

		var _parser parser.Parser
		_parser.Init(fileName, code)
		program, err := _parser.Parse()
		if err != nil {
			fmt.Print(err.Error())
			continue
		}

		result, err := Eval(program, env)
		if err == nil {
			fmt.Println(result.StringValue())
		} else {
			fmt.Print(Errf(err))
		}
	}
}

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
	for _, program := range programs {
		_, err := Eval(program, env)
		if err != nil {
			return Errf(err)
		}
	}
	_, err = evalMainFunction(env)
	if err != nil {
		return Errf(err)
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
	_, err = Eval(program, env)
	if err != nil {
		return Errf(err)
	}
	_, err = evalMainFunction(env)
	if err != nil {
		return Errf(err)
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
	var _parser parser.Parser
	_parser.Init(fileName, code)
	program, err := _parser.ParseProgram()
	if err != nil {
		return nil, err
	}
	return program, nil
}
