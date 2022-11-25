package bffmt

import (
	"os"
	"regexp"

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
		Optimizes the source, as explained in https://github.com/baris-inandi/brainfuck-go/issues/2 .
		It assumes `s` only has valid opcodes.
	*/
	var minify = func(s string) string {
		// matches 256 consecutive + or - (exclusive, so mixes don't match)
		var plus_minus_256 = regexp.MustCompile(`\+{256}|-{256}`)
		s = plus_minus_256.ReplaceAllLiteralString(s, "")

		// matches pairs of BF opcodes that cancel each other (any order)
		var mutual_cancel = regexp.MustCompile(`\+-|-\+|><|<>`)

		var size int
		// TODO: optimize later
		for do := true; do; do = (size != len(s)) {
			size = len(s)
			s = mutual_cancel.ReplaceAllLiteralString(s, "")
		}

		// matches a plus-sign between square braces an even number of times
		var plus_even = regexp.MustCompile(`\[(?:\+\+){2,128}\]`)
		s = plus_even.ReplaceAllLiteralString(s, "[++]")

		// matches a minus-sign between square braces an even number of times
		var minus_even = regexp.MustCompile(`\[(?:--){2,128}\]`)
		s = minus_even.ReplaceAllLiteralString(s, "[--]")

		var plus = regexp.MustCompile(`\[(?:\+){2,256}\]`)
		s = plus.ReplaceAllLiteralString(s, "[+]")

		var minus = regexp.MustCompile(`\[(?:-){2,256}\]`)
		s = minus.ReplaceAllLiteralString(s, "[-]")

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
