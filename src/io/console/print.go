// SPDX-License-Identifier: GPL-3.0-or-later
package console

import (
	"fmt"
)

func defaultPrint(text ...interface{}) {
	fmt.Print(text...)
}

func defaultPrintln(text ...interface{}) {
	fmt.Println(text...)
}

func defaultPrintf(format string, text ...interface{}) {
	fmt.Printf(format, text...)
}
