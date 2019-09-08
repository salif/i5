package errors

const (
	ARGS_UNKNOWN           string = "unknown option: %v"
	ARGS_UNKNOWN_CLR       string = "unknown output format: %v"
	ARGS_NO_FILE           string = "no file specified"
	LEXER_UNEXPECTED_TOKEN string = "unexpected token '%v' at line %v\n"
	SCANNER_OUT_OF_RANGE   string = "line %v: %v: index out of range\n"
	PARSER_EXPECTED_FOUND  string = "expected '%v', found '%v' at line %v"
	READER_NOT_FOUND       string = "%v: no such file or directory"
	TEST_GOT_WANT          string = "error: got: %v, want: %v"
)
