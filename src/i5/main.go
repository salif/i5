// SPDX-License-Identifier: GPL-3.0-or-later
package i5

import (
	"fmt"
	"os"

	"github.com/i5/i5/src/printer"

	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/i5/args_parser"
	"github.com/i5/i5/src/i5/colors"
	"github.com/i5/i5/src/interpreter"
)

// Parse CLI arguments
func ParseArguments() {
	var argumentsParser args_parser.ArgumentsParser

	argumentsParser.Init("i5 [options] [file] [arguments]", "Options")
	argumentsParser.Bool("help", "print help")
	argumentsParser.Bool("code", "print code")
	argumentsParser.Bool("tokens", "print tokens")
	argumentsParser.Bool("ast", "print abstract syntax tree")
	argumentsParser.String("output", "set output format", "format")
	argumentsParser.String("eval", "evaluate code", "code")
	argumentsParser.Bool("init", "initialize new module")
	argumentsParser.Bool("version", "print current version")

	err := argumentsParser.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v%v\n", colors.Red("error: "), err.Error())
		return
	}

	if argumentsParser.IsEmpty() || argumentsParser.IsTrue("help") {
		fmt.Println(argumentsParser.GetHelp())
		return
	}

	if argumentsParser.IsTrue("init") {
		InitModule()
		return
	}

	if argumentsParser.IsTrue("version") {
		PrintVersion()
		return
	}

	if len(argumentsParser.Get("output")) > 0 {
		format := argumentsParser.Get("output")
		err := colors.SetColorFormat(format)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v%v\n", colors.Red("error: "), err.Error())
			return
		}
	}

	notOptions := argumentsParser.GetNotOptions()
	if len(notOptions) > 0 {
		t1 := argumentsParser.IsTrue("tokens")
		t2 := argumentsParser.IsTrue("code")
		t3 := argumentsParser.IsTrue("ast")
		if t1 || t2 || t3 {
			err := printer.Print(notOptions[0], t1, t2, t3)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err.Error())
			}
			return
		}
		cEval := argumentsParser.Get("eval")
		if len(cEval) > 0 {
			err := interpreter.RunEval(cEval, notOptions)
			if err != nil {
				fmt.Fprint(os.Stderr, err.Error())
			}
		} else {
			err := interpreter.Run(notOptions[0], notOptions)
			if err != nil {
				fmt.Fprint(os.Stderr, err.Error())
			}
		}
	}
}

// Print current version
func PrintVersion() {
	fmt.Printf("v%v\n", constants.MINOR_VERSION)
}
