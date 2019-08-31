package errors

import (
	"os"

	"github.com/i5/i5/src/io/console"
)

func Error(text interface{}, status int) {
	console.Println(console.Color{Value: "error:"}.Red(), text)
}

func FatalError(text interface{}, status int) {
	console.Println(console.Color{Value: "error:"}.Red(), text)
	Exit(status)
}

func Exit(status int) {
	os.Exit(status)
}
