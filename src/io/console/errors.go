// SPDX-License-Identifier: GPL-3.0-or-later
package console

const (
	ARGS_UNKNOWN     string = "unknown option: %v"
	ARGS_UNKNOWN_CLR string = "unknown output format: %v"
	ARGS_NO_FILE     string = "no file specified"

	LEXER_UNEXPECTED_TOKEN string = "unexpected token '%v' at line %v"
	LEXER_OUT_OF_RANGE     string = "line %v: %v: index out of range"

	PARSER_EXPECTED_FOUND string = "expected '%v', found '%v' at line %v"
	PARSER_UNEXPECTED     string = "unexpected '%v' at line %v"
	PARSER_NOT_NUMBER     string = "could not parse %q as number"

	FILE_READ_NOT_FOUND   string = "%v: no such file or directory"
	FILE_READ_DIR         string = "%v: expected file, found directory"
	FILE_READ_CANNOT_READ string = "%v: can not read file"

	FILE_WRITE_EXISTS       string = "%v: file already exists"
	FILE_WRITE_CANNOT_WRITE string = "%v: can not write file"

	TEST_GOT_WANT string = "error: got: %v, want: %v"
)
