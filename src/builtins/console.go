// SPDX-License-Identifier: GPL-3.0-or-later
package builtins

import (
	"bufio"
	"fmt"
	"os"

	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func console_write(args ...object.Object) object.Object {
	for _, o := range args {
		fmt.Print(o.StringValue())
	}
	return Null
}

func console_write_line(args ...object.Object) object.Object {
	for _, o := range args {
		fmt.Print(o.StringValue())
	}
	fmt.Println()
	return Null
}

func console_read_line(args ...object.Object) object.Object {
	if len(args) > 0 {
		console_write(args[0])
	}
	v := bufio.NewReader(os.Stdin)
	input, err := v.ReadString('\n')
	if err != nil {
		return newException(constants.EXCEPTION_INTERNAL, err.Error())
	}
	if input == "\n" {
		return object.String{Value: input}
	} else {
		return object.String{Value: input[:len(input)-1]}
	}
}

func printf(args ...object.Object) object.Object {
	return Null
}
