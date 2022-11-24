package bffmt

import (
	"os"
	"strings"

	"github.com/baris-inandi/brainfuck/lang/readcode"
)

func MinifyFile(files ...string) {

	/*
		Finds the address of the 1st `rune` that isn't in the set {"[", "]"}.

		`start` ignores all runes before that index.
		If `start` is negative, the index becomes relative to the end.
	*/
	/* var indexNoBrace = func(s string, start int) int {
		size := len(s)
		if start < 0 {
			start += size
		}
		if start < 0 {
			// should it return an err instead of panicking?
			panic("Index out of bounds")
		}

		for start < size {
			if s[start] != "["[0] && s[start] != "]"[0] {
				return start
			}
			start += 1
		}
		return -1
	} */

	/*
		Optimizes the source, as explained in https://github.com/baris-inandi/brainfuck-go/issues/2
	*/
	var minify = func(s string) string {
		// exploit mod 256 wrap-around
		s = strings.ReplaceAll(s, strings.Repeat("+", 0x100), "")
		s = strings.ReplaceAll(s, strings.Repeat("-", 0x100), "")

		size := len(s)
		for do := true; do; do = (size != len(s)) {
			size = len(s)
			s = strings.ReplaceAll(s, "+-", "")
			s = strings.ReplaceAll(s, "-+", "")
			s = strings.ReplaceAll(s, "><", "")
			s = strings.ReplaceAll(s, "<>", "")
		}

		/* // simulated BF memory/tape
			var mem = map[int]uint8{}
			// relative memory pointer
			var ptr int = 0

		label:
			for i := indexNoBrace(s, 0); i < len(s); i += 1 {
				switch s[i] {
				case "+"[0]:
					{
						mem[ptr] += 1
						continue
					}
				case "-"[0]:
					{
						mem[ptr] -= 1
						continue
					}
				case ">"[0]:
					{
						ptr += 1
						continue
					}
				case "<"[0]:
					{
						ptr -= 1
						continue
					}
				case ","[0]:
					break label
				case "."[0]:
					break label
				default:
					i = indexNoBrace(s, i)
				}
			} */

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
