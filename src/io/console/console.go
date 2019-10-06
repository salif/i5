// SPDX-License-Identifier: GPL-3.0-or-later
package console

import (
	"fmt"
)

func Format(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

func Print(toWrite ...interface{}) {
	localPrint(toWrite...)
}

func Println(toWrite ...interface{}) {
	localPrintln(toWrite...)
}

func Printf(format string, toWrite ...interface{}) {
	localPrintf(format, toWrite...)
}
