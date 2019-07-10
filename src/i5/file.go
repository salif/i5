package i5

import (
	"github.com/i5-lang/i5/src/error"
	"io/ioutil"
)

func ReadFile(fileName string) []byte {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		error.FatalError("error: "+fileName+": No such file or directory", 1)
	}
	return file
}
