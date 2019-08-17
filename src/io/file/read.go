package file

import (
	"github.com/i5/i5/src/errors"
	"io/ioutil"
)

func Read(fileName string) []byte {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		errors.FatalError("error: "+fileName+": no such file or directory", 1)
	}
	return file
}
