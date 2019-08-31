package i5

import (
	"fmt"

	"github.com/i5/i5/src/errors"
	"github.com/i5/i5/src/interpreter"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/io/file"
	"github.com/i5/i5/src/lexer"
	"github.com/i5/i5/src/parser"
	"github.com/i5/i5/src/types"
)

var (
	ap      ArgsParser = InitArgsParser()
	_help              = ap.Bool("help")
	_tokens            = ap.Bool("tokens")
	_code              = ap.Bool("code")
	_ast               = ap.Bool("ast")
	_output            = ap.String("output")
	_eval              = ap.String("eval")
	_args              = ap.Default()
)

func ParseArgs() {

	ap.Parse()

	if ap.Empty() || *_help {
		PrintHelp()
		errors.Exit(0)
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
			errors.FatalError(fmt.Sprintf(errors.ARGS_UNKNOWN_CLR, *_output), 1)
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
			PrintCode(tokenList)
		}
		if *_tokens {
			PrintTokens(tokenList)
		}
		if *_ast {
			PrintAst(parser.Run(tokenList), 0, "red")
		}
	} else {
		interpreter.Run(parser.Run(tokenList))
	}
}

func PrintHelp() {
	console.Println(`
Usage:

    i5 [options] [file] [arguments]

options:

    --help                      print help
    --code                      print code
    --tokens                    print tokens
    --ast                       print ast
    --output='string'           set output format
                                ('html', 'no-color', 'default')
    --eval='string'             eval code
    `)
}
