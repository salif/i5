// SPDX-License-Identifier: GPL-3.0-or-later
package i5

import (
	"github.com/i5/i5/src/errors"
	"github.com/i5/i5/src/interpreter"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/io/file"
	"github.com/i5/i5/src/lexer"
	"github.com/i5/i5/src/parser"
	"github.com/i5/i5/src/printer"
	"github.com/i5/i5/src/types"
)

const (
	MajorVersion = "0"
	MinorVersion = "0.0"
	PatchVersion = "0.0.0"
)

var (
	ap       ArgsParser = InitArgsParser()
	_help               = ap.Bool("help")
	_init               = ap.Bool("init")
	_tokens             = ap.Bool("tokens")
	_code               = ap.Bool("code")
	_ast                = ap.Bool("ast")
	_output             = ap.String("output")
	_eval               = ap.String("eval")
	_version            = ap.Bool("version")
	_args               = ap.Default()
)

func ParseArgs() {

	ap.Parse()

	if ap.Empty() || *_help {
		PrintHelp()
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
			errors.FatalError(errors.F(errors.ARGS_UNKNOWN_CLR, *_output), 1)
		}
	}

	if len(*_eval) > 0 {
		Run(*_eval, *_args, false)
	} else if len(*_args) > 0 {
		Run((*_args)[0], *_args, true)
	}
}

func Run(name string, arguments []string, isFile bool) {
	var tokenList types.TokenList
	if isFile {
		tokenList = lexer.Run(file.Read(name))
	} else {
		tokenList = lexer.Run([]byte(name))
	}
	if *_code || *_tokens || *_ast {
		if *_code {
			printer.Code(tokenList)
		}
		if *_tokens {
			printer.Tokens(tokenList)
		}
		if *_ast {
			printer.Ast(parser.Run(tokenList), 0, "red")
		}
	} else {
		interpreter.Run(parser.Run(tokenList))
	}
}

func PrintVersion() {
	console.Println("i5 version: v" + MinorVersion)
}

func PrintHelp() {
	console.Println(`
Usage:

    i5 [options] [file] [arguments]

options:

    --help                      print help
    --code                      print code
    --tokens                    print tokens
    --ast                       print abstract syntax tree
    --output='format'           set output format:
                                ('html', 'no-color', 'default')
    --eval='code'               evaluate code
    --init                      initialize new module
    --version                   print current version
    `)
}
