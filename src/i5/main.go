package i5

import (
	"fmt"
	"os"

	"github.com/i5/i5/src/errors"
	"github.com/i5/i5/src/interpreter"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/io/file"
	"github.com/i5/i5/src/lexer"
	"github.com/i5/i5/src/parser"
	"github.com/i5/i5/src/types"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	_code   = kingpin.Flag("code", "Print code").Short('c').Bool()
	_tokens = kingpin.Flag("tokens", "Print tokens").Short('t').Bool()
	_ast    = kingpin.Flag("ast", "Print AST").Short('s').Bool()
	_files  = kingpin.Arg("file", "Run code").Strings()
	_color  = kingpin.Flag("color", "Color").String()
	_evals  = kingpin.Flag("eval", "Eval code").Short('e').Strings()
)

func ParseArgs() {
	if len(os.Args) < 2 {
		PrintHelp()
		errors.Exit(0)
	}

	kingpin.Parse()

	if len(*_color) > 0 {
		switch *_color {
		case "html":
			console.SetColorizer(console.HTML)
		case "no":
			console.SetColorizer(console.NoColor)
		case "color":
			console.SetColorizer(console.Color)
		default:
			errors.NewFatalError(fmt.Sprintf(errors.ARGS_UNKNOWN_CLR, *_color), 1)
		}
	}
	if len(*_files) > 0 {
		Run(*_files, true)
	} else if len(*_evals) > 0 {
		Run(*_evals, false)
	}
}

func Run(names []string, areFiles bool) {
	for _, name := range names {
		var tokenList types.TokenList
		if areFiles {
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
}

func PrintHelp() {
	console.Println(`
Usage:

    i5 [options] [file] [arguments]

options:

    --help           print help
    --code, -c       print code
    --tokens, -t     print tokens
    --ast, -s        print ast
    --eval           eval code
    `)
}
