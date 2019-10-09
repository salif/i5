// SPDX-License-Identifier: GPL-3.0-or-later
package i5

import (
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/interpreter"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/io/file"
	"github.com/i5/i5/src/lexer"
	"github.com/i5/i5/src/parser"
	"github.com/i5/i5/src/printer"
)

var (
	ap ArgsParser = InitArgsParser("i5 [options] [file] [arguments]", "Options")

	_help    = ap.Bool("help", "print help")
	_code    = ap.Bool("code", "print code")
	_tokens  = ap.Bool("tokens", "print tokens")
	_ast     = ap.Bool("ast", "print abstract syntax tree")
	_output  = ap.String("output", "set output format", "format")
	_eval    = ap.String("eval", "evaluate code", "code")
	_init    = ap.Bool("init", "initialize new module")
	_version = ap.Bool("version", "print current version")
	_args    = ap.Default()
)

// Parse CLI arguments
func ParseArgs() {

	ap.Parse()

	if ap.Empty() || *_help {
		ap.PrintHelp()
		return
	}

	if *_init {
		InitModule()
		return
	}

	if *_version {
		PrintVersion()
		return
	}

	if len(*_output) > 0 {
		switch *_output {
		case "html":
			console.SetOutput(console.HTML)
		case "no-color":
			console.SetOutput(console.NoColor)
		case "default":
			console.SetOutput(console.Default)
		default:
			console.ThrowError(1, constants.ARGS_UNKNOWN_CLR, *_output)
		}
	}

	if len(*_eval) > 0 {
		runEval(*_eval, *_args)
	} else if len(*_args) > 0 {
		runFileOrPackage((*_args)[0], *_args)
	}
}

func runEval(content string, arguments []string) {
	if *_code || *_tokens || *_ast {
		printer.Print(content, false, *_code, *_tokens, *_ast)
	} else {
		interpreter.RunFile(parser.Run(lexer.Run([]byte(content))), arguments)
	}
}

func runFile(name string, arguments []string) {
	if *_code || *_tokens || *_ast {
		printer.Print(name, true, *_code, *_tokens, *_ast)
	} else {
		interpreter.RunFile(parser.Run(lexer.Run(file.Read(name))), arguments)
	}
}
func runPackage(name string, arguments []string) {
	if file.Exists(constants.I5_MOD_FILE_NAME) {
		interpreter.RunModule(name, arguments)
	} else {
		interpreter.RunPackage(name, arguments)
	}
}

func runFileOrPackage(name string, arguments []string) {
	var result int = file.Info(name)
	switch result {
	case 1:
		console.ThrowError(1, constants.FILE_NOT_FOUND, name)
	case 2:
		runPackage(name, arguments)
	case 3:
		runFile(name, arguments)
	}
}

// Print current minor version
func PrintVersion() {
	console.Printf("i5 version: v%v\n", constants.MINOR_VERSION)
}
