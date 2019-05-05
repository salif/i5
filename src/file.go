package src

import (
	"fmt"
	"os"
	"io/ioutil"
)

func ReadFile(name string) string {
	file, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("error: " + name + ": No such file or directory")
		os.Exit(1)
	}
	return string(file)
}
