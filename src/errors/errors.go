package errors

import (
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/types"
	"os"
)

func NewFatalError(text interface{}, status int) types.Node {
	console.Println(console.Color("error:", "red"), text)
	Exit(status)
	return types.Node{}
}

func NewFatalLexerError(text string, line int, char string, status int) {
	console.Printf(console.Color("error: ", "red")+text, line, char)
	Exit(status)
}

func Exit(status int) {
	os.Exit(status)
}
