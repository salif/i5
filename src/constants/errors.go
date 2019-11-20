// SPDX-License-Identifier: GPL-3.0-or-later
package constants

const (
	TEST_GOT_WANT string = "error: got: %v, want: %v"

	ARGS_UNKNOWN     string = "unknown option: %v"
	ARGS_UNKNOWN_CLR string = "unknown output format: %v"
	ARGS_NO_FILE     string = "no file specified"

	LEXER_UNEXPECTED_TOKEN string = "unexpected token '%v' with ascii code '%v'"
	LEXER_OUT_OF_RANGE     string = "%v: index out of range"

	PARSER_EXPECTED_FOUND string = "expected '%v', found '%v'"
	PARSER_UNEXPECTED     string = "unexpected '%v'"
	PARSER_EXPECTED       string = "expected '%v'"
	PARSER_NOT_INT        string = "could not parse %v as integer"
	PARSER_NOT_FLOAT      string = "could not parse %v as float"
	PARSER_EXPECTED_ARG   string = "expected ',' or ')', found '%v'"

	FILE_NOT_FOUND        string = "%v: no such file or directory"
	FILE_FOUND_FILE       string = "%v: expected directory, found file"
	FILE_FOUND_DIR        string = "%v: expected file, found directory"
	FILE_CANNOT_READ_FILE string = "%v: can not read file"
	FILE_CANNOT_READ_DIR  string = "%v: can not read directory"
	FILE_FILE_EXISTS      string = "%v: file already exists"
	FILE_DIR_EXISTS       string = "%v: directory already exists"
	FILE_CANNOT_WRITE     string = "%v: can not write file"

	IR_NIL               string = "nil"
	IR_MAIN_FN_NOT_FOUND string = "main function not found"
	IR_INVALID_EVAL      string = "invalid expression: '%v'"
	IR_INVALID_INFIX     string = "invalid expression: '%v%v%v'"
	IR_INVALID_PREFIX    string = "invalid expression: '%v%v'"
	IR_INVALID_POSTFIX   string = "invalid expression: '%v%v'"
	IR_IS_NOT_A_FUNCTION string = "'%v' is not a function"
	IR_IS_NOT_A_BOOL     string = "'%v' is not a boolean"
	IR_IS_NOT_AN_ERROR   string = "'%v' is not an error"
	IR_IS_NOT_AN_ARRAY   string = "'%v' is not an array"
	IR_IS_NOT_A_MAP      string = "'%v' is not a map"
	IR_IS_NOT_A_CLASS    string = "'%v' is not a class"
	IR_BUILTIN_NOT_FOUND string = "builtin not found: '$%v'"
	IR_CANNOT_ASSIGN     string = "cannot assign to %v"
	IR_CANNOT_BE_KEY     string = "'%v' cannot be map key"
	IR_NOT_ENOUGH_ARGS   string = "not enough arguments to call function"
)
