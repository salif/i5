package file

import (
	"io/ioutil"
	"os"

	"github.com/i5/i5/src/errors"
)

func Write(fileName string, content string, perm os.FileMode) (result string) {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		err := ioutil.WriteFile(fileName, []byte(content), perm)
		if err != nil {
			return errors.F(errors.WRITER_CANNOT_WRITE, fileName)
		} else {
			return
		}
	} else {
		return errors.F(errors.WRITER_EXISTS, fileName)
	}
}
