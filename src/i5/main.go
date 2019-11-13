// SPDX-License-Identifier: GPL-3.0-or-later
package i5

import (
	"fmt"
	"os"

	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/i5/args_parser"
	"github.com/i5/i5/src/i5/colors"
	"github.com/i5/i5/src/interpreter"
	"github.com/i5/i5/src/printer"
)

// Parse CLI arguments
func ParseArguments() {
	var argumentsParser args_parser.ArgumentsParser

	argumentsParser.Init("i5 [options] [file] [arguments]", "Options")
	argumentsParser.Bool("help", "print help")
	argumentsParser.String("print", "string: 'tokens', 'code' or 'ast'", "string")
	argumentsParser.String("output", "string: 'no-color', 'html' or 'default'", "string")
	argumentsParser.String("eval", "evaluate code", "code")
	argumentsParser.Bool("version", "print current version")

	var err error = argumentsParser.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v%v\n", colors.Red("error: "), err.Error())
		os.Exit(1)
	}

	if argumentsParser.IsEmpty() || argumentsParser.IsTrue("help") {
		fmt.Println(argumentsParser.GetHelp())
		os.Exit(1)
	}

	if argumentsParser.IsTrue("version") {
		printVersion()
		return
	}

	if len(argumentsParser.Get("output")) > 0 {
		format := argumentsParser.Get("output")
		var err error = colors.SetColorFormat(format)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v%v\n", colors.Red("error: "), err.Error())
			os.Exit(1)
		}
	}

	var notOptions []string = argumentsParser.GetNotOptions()

	Eval := argumentsParser.Get("eval")
	if len(Eval) > 0 {
		var err error = interpreter.RunEval(Eval, notOptions)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}
	}

	if len(notOptions) > 0 {
		var toRun string = notOptions[0]
		var printOption string = argumentsParser.Get("print")
		if len(printOption) > 0 {
			var err error = printer.Print(toRun, printOption)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err.Error())
				os.Exit(1)
			}
			return
		}

		var err error = interpreter.Run(toRun, notOptions)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}
	}
}

// Print current version
func printVersion() {
	fmt.Printf("v%v\n", constants.MINOR_VERSION)
}
