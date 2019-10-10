// SPDX-License-Identifier: GPL-3.0-or-later
package constants

const (
	TEST_GOT_WANT string = "error: got: %v, want: %v"

	ARGS_UNKNOWN     string = "unknown option: %v"
	ARGS_UNKNOWN_CLR string = "unknown output format: %v"
	ARGS_NO_FILE     string = "no file specified"

	LEXER_UNEXPECTED_TOKEN string = "unexpected token '%v' at line %v"
	LEXER_OUT_OF_RANGE     string = "line %v: %v: index out of range"

	PARSER_EXPECTED_FOUND string = "expected '%v', found '%v' at line %v"
	PARSER_UNEXPECTED     string = "unexpected '%v' at line %v"
	PARSER_NOT_INT        string = "could not parse %v as integer"
	PARSER_NOT_FLOAT      string = "could not parse %v as float"
	PARSER_EXPECTED_ARG   string = "expected ',' or ')', found '%v' at line %v"

	FILE_NOT_FOUND    string = "%v: no such file or directory"
	FILE_FOUND_FILE   string = "%v: expected directory, found file"
	FILE_FOUND_DIR    string = "%v: expected file, found directory"
	FILE_CANNOT_READ  string = "%v: can not read file"
	FILE_FILE_EXISTS  string = "%v: file already exists"
	FILE_DIR_EXISTS   string = "%v: directory already exists"
	FILE_CANNOT_WRITE string = "%v: can not write file"

	IR_INVALID_EVAL      string = "invalid expression: %v"
	IR_INVALID_INFIX     string = "invalid expression: %v %v %v"
	IR_INVALID_PREFIX    string = "invalid expression: %v %v"
	IR_INVALID_SUFFIX    string = "invalid expression: %v %v"
	IR_INVALID_CALL      string = "invalid function caller: %v"
	IR_BUILTIN_NOT_FOUND string = "builtin not found: %v"
)
