// SPDX-License-Identifier: GPL-3.0-or-later
package lexer

func ParseModuleFile(fileName string, f []byte) []string {
	result := []string{}
	var scanner Scanner
	scanner.Init(f)

	for scanner.HasNext() {

		if scanner.PeekEquals(10) || scanner.PeekEquals(13) {
			scanner.Next()
			continue
		}

		if scanner.PeekEquals(35) {
			scanner.Next()
			for ; scanner.HasNext() && scanner.Until(10); scanner.Next() {
			}
			continue
		}

		var value string = ""
		for ; scanner.HasNext() && scanner.Until(10); scanner.Next() {
			value += string(scanner.Peek())
		}
		result = append(result, value)
		continue
	}

	return result
}
