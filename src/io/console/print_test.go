// SPDX-License-Identifier: GPL-3.0-or-later
package console

func ExamplePrint() {
	Print("Text!")
	// Output: Text!
}

func ExamplePrintln() {
	Println("Text!")
	// Output: Text!
}

func ExamplePrintf() {
	Printf("Text%v", "!")
	// Output: Text!
}
