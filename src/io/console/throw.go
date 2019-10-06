// SPDX-License-Identifier: GPL-3.0-or-later
package console

import (
	"os"
)

func PrintError(prefix string, text interface{}) {
	Println(Red(prefix), text)
}

func ThrowError(exitCode int, text string, format ...interface{}) {
	PrintError("error:", Format(text, format...))
	Exit(exitCode)
}

func ThrowSyntaxError(exitCode int, text string, format ...interface{}) {
	PrintError("syntax error:", Format(text, format...))
	Exit(exitCode)
}

func ThrowParsingError(exitCode int, text string, format ...interface{}) {
	PrintError("parsing error:", Format(text, format...))
	Exit(exitCode)
}

func Exit(exitCode int) {
	os.Exit(exitCode)
}
