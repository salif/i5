package console

import "fmt"

func Print(toWrite ...interface{}) {
	fmt.Print(toWrite...)
}

func Println(toWrite ...interface{}) {
	fmt.Println(toWrite...)
}

func Printf(format string, toWrite ...interface{}) {
	fmt.Printf(format, toWrite...)
}
