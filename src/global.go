/*
	global.go
		Global variables and functions that will be used everywhere in src/
*/

package src

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var intermediate = "" // store for the whole intermediate code
const tapeLen = 30000 // default tape length

// global variables needed for the interpreter/repl
var tape [tapeLen]byte // the tape declaration
var ptr = uint64(0)    // the pointer

// other utility functions
func runeInSlice(a rune, list []rune) bool {
	/*
		func runeInSlice
			Gets param a where a is a rune,
			and list is a list of runes,
			checks if a is in list
	*/
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func toValidBf(s string) string {
	/*
		func toValidBf
			removes all illegal characters from brainfuck code,
			where s is a string of code
	*/
	return strings.Map(
		func(r rune) rune {
			if runeInSlice(r, []rune{'<', '>', '+', '-', '.', '[', ']', ','}) {
				return r
			}
			return -1
		}, s,
	)
}

func readBrainfuck(f string) string {
	/*
		func readBrainfuck
			Gets param f where f is
			a filepath to a brainfuck file,
			reads the brainfuck file,
			removes all non-brainfuck characters,
			returns valid brainfuck code
	*/
	fileBytes, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Print(err)
	}
	return toValidBf(string(fileBytes))
}
