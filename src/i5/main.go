// SPDX-License-Identifier: GPL-3.0-or-later
package i5

import (
	"fmt"
	"os"

	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/i5/args_parser"
	"github.com/i5/i5/src/interpreter"
)

// Parse CLI arguments
func ParseArguments(args []string) int {
	argumentsParser := args_parser.ArgumentsParser{}
	argumentsParser.Init(args, "i5 [options] [file] [arguments]")
	argumentsParser.Bool("help", "print help")
	argumentsParser.String("print", "string: 'code' or 'ast'", "string")
	argumentsParser.String("eval", "evaluate code", "code")
	argumentsParser.Bool("version", "print current version")

	var parseError error = argumentsParser.Parse()
	if parseError != nil {
		fmt.Fprintf(os.Stderr, "%v%v\n", "error: ", parseError.Error())
		return 1
	}

	if argumentsParser.IsTrue("help") {
		fmt.Println(argumentsParser.GetHelp())
		return 0
	}

	if argumentsParser.IsTrue("version") {
		printVersion()
		return 0
	}

	var realArguments []string = argumentsParser.GetRealArguments()
	newInterpreter := interpreter.New(realArguments, argumentsParser.Get("print"))
	var err error

	eval := argumentsParser.Get("eval")

	if len(eval) > 0 {
		err = newInterpreter.RunEval(eval)
	} else if len(realArguments) > 0 {
		err = newInterpreter.Run()
	} else {
		err = newInterpreter.RunRepl()
	}

	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return 1
	} else {
		return 0
	}
}

// Print current version
func printVersion() {
	fmt.Printf("v%v\n", constants.MINOR_VERSION)
}
