package lexer

import (
	"testing"
)

var (
	test_text1 string = `
	fn main() {
		$print(1)
	}
	`
	test_s1 int    = 15
	test1   []byte = []byte(test_text1)
)

func TestRun(t *testing.T) {
	res1 := Run(test1)
	if res1.Size() != test_s1 {
		t.Errorf("Error! got: %v, want: %v", res1, test_s1)
	}
}
