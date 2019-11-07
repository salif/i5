// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"bufio"
	"fmt"
	"os"

	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func _console() object.Object {
	v := _Map()
	v.Set(_String("print"), _Builtin(object.ERROR, 1, _console_print))
	v.Set(_String("println"), _Builtin(object.ERROR, 0, _console_println))
	v.Set(_String("readln"), _Builtin(object.STRING, 0, _console_readln))
	return v
}

func _console_print(args ...object.Object) object.Object {
	for _, o := range args {
		fmt.Print(o.StringValue())
	}
	return NIL
}

func _console_println(args ...object.Object) object.Object {
	for _, o := range args {
		fmt.Print(o.StringValue())
	}
	fmt.Println()
	return NIL
}

func _console_readln(args ...object.Object) object.Object {
	if len(args) > 0 {
		_console_print(args[0])
	}
	v := bufio.NewReader(os.Stdin)
	input, err := v.ReadString('\n')
	if err != nil {
		return _Error(true, constants.ERROR_INTERTAL, err.Error())
	}
	if input == "\n" {
		return _String(input)
	} else {
		return _String(input[:len(input)-1])
	}
}
