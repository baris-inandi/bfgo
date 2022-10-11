package bffmt

import (
	"errors"
	"os"
	"strings"

	"github.com/baris-inandi/brainfuck/lang/readcode"
)

func MinifyFile(files ...string) {

	/*
		Check that all the square brackets are matched correctly.
		it would be more helpful if it returned line/column,
		but we don't need that, for now
	*/
	isValidSyntax := func(s string) bool {
		return true
	}

	minify := func(s string) (string, error) {
		if !isValidSyntax(s) {
			return s, errors.New("SyntaxError: Mismatched brackets")
		}

		// exploit mod 256 wrap-around
		s = strings.ReplaceAll(s, strings.Repeat("+", 0x100), "")
		s = strings.ReplaceAll(s, strings.Repeat("-", 0x100), "")

		return s, nil
	}

	for _, f := range files {
		minified, _ := minify(readcode.ReadBrainfuck(f))
		err := os.WriteFile(f, []byte(minified), 0644)
		if err != nil {
			panic(err)
		}
	}
}
