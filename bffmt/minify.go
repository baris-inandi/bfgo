package bffmt

import (
	"os"
	"strings"

	"github.com/baris-inandi/brainfuck/lang/readcode"
)

func MinifyFile(files ...string) {

	/*
		finds the address of the 1st `rune` that isn't in the set {"[", "]"}.
		`start` ignores all runes before that index.
	*/
	indexOfNoBrace := func(s string, start int) int {
		for start < len(s) {
			if s[start] != "["[0] && s[start] != "]"[0] {
				return start
			}
			start += 1
		}
		return -1
	}

	minify := func(s string) string {
		// exploit mod 256 wrap-around
		s = strings.ReplaceAll(s, strings.Repeat("+", 0x100), "")
		s = strings.ReplaceAll(s, strings.Repeat("-", 0x100), "")

		for i := indexOfNoBrace(s, 0); i < len(s); i += 1 {
			//todo
		}

		return s
	}

	for _, f := range files {
		minified := minify(readcode.ReadBrainfuck(f))
		err := os.WriteFile(f, []byte(minified), 0o644)
		if err != nil {
			panic(err)
		}
	}
}
