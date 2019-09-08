package errors

import (
	"fmt"
	"os"

	"github.com/i5/i5/src/io/console"
)

func Error(text interface{}) {
	console.Println(console.Color{Value: "error:"}.Red(), text)
}

func FatalError(text interface{}, status int) {
	console.Println(console.Color{Value: "error:"}.Red(), text)
	Exit(status)
}

func Exit(status int) {
	os.Exit(status)
}

func F(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

func LN(a ...interface{}) string {
	return fmt.Sprintln(a...)
}
