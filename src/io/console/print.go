package console

import "fmt"

type output struct {
	clr     func(string, string) string
	print   func(...interface{})
	println func(...interface{})
	printf  func(string, ...interface{})
}

var print func(...interface{}) = defaultPrint
var println func(...interface{}) = defaultPrintln
var printf func(string, ...interface{}) = defaultPrintf

var (
	HTML    = output{htmlColor, defaultPrint, defaultPrintln, defaultPrintf}
	Default = output{defaultColor, defaultPrint, defaultPrintln, defaultPrintf}
	NoColor = output{noColor, defaultPrint, defaultPrintln, defaultPrintf}
)

func SetOutput(p output) {
	clr = p.clr
	print = p.print
	println = p.println
	printf = p.printf
}

func Print(toWrite ...interface{}) {
	print(toWrite...)
}

func Println(toWrite ...interface{}) {
	println(toWrite...)
}

func Printf(format string, toWrite ...interface{}) {
	printf(format, toWrite...)
}

func defaultPrint(toWrite ...interface{}) {
	fmt.Print(toWrite...)
}

func defaultPrintln(toWrite ...interface{}) {
	fmt.Println(toWrite...)
}

func defaultPrintf(format string, toWrite ...interface{}) {
	fmt.Printf(format, toWrite...)
}
