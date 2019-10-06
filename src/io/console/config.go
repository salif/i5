// SPDX-License-Identifier: GPL-3.0-or-later
package console

type output struct {
	colorize func(string, string) string
	print    func(...interface{})
	println  func(...interface{})
	printf   func(string, ...interface{})
}

var localPrint func(...interface{}) = defaultPrint
var localPrintln func(...interface{}) = defaultPrintln
var localPrintf func(string, ...interface{}) = defaultPrintf

var (
	HTML    = output{htmlColor, defaultPrint, defaultPrintln, defaultPrintf}
	Default = output{defaultColor, defaultPrint, defaultPrintln, defaultPrintf}
	NoColor = output{noColor, defaultPrint, defaultPrintln, defaultPrintf}
)

func SetOutput(p output) {
	colorize = p.colorize
	localPrint = p.print
	localPrintln = p.println
	localPrintf = p.printf
}
