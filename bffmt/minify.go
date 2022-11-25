package bffmt

import (
	"os"
	"regexp"
	"strings"

	"github.com/baris-inandi/brainfuck/lang/readcode"
)

func MinifyFile(files ...string) {

	// returns a pair containing the char frequency of "+" and "-", respectively
	var counterPlusMinus = func(s string) (uint, uint) {
		var plus uint = 0
		var minus uint = 0

		for i := 0; i < len(s); i++ {
			c := s[i]
			if c == "+"[0] {
				plus += 1
			}
			if c == "-"[0] {
				minus += 1
			}
		}

		return plus, minus
	}

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

	// Optimizes the source, as explained in https://github.com/baris-inandi/brainfuck-go/issues/2 .
	// It assumes `s` only has valid opcodes.
	var minify = func(s string) string {
		// matches 256 consecutive + or - (exclusive, so mixes don't match)
		var x256PlusMinus = regexp.MustCompile(`\+{256}|-{256}`)
		s = x256PlusMinus.ReplaceAllLiteralString(s, "")

		// matches pairs of BF opcodes that cancel each other (any order)
		var mutualCancel = regexp.MustCompile(`\+-|-\+|><|<>`)

		var size int
		// TODO: optimize later
		for do := true; do; do = (size != len(s)) {
			size = len(s)
			s = mutualCancel.ReplaceAllLiteralString(s, "")
		}

		// matches + or - between square braces an even number of times (max 128)
		var evenPlusMinus = regexp.MustCompile(`\[(?:(?:\+\+){1,128}|(?:--){2,128})\]`)
		s = evenPlusMinus.ReplaceAllLiteralString(s, "[--]")

		// matches + or - between square braces an odd number of times (max 128)
		var oddPlusMinus = regexp.MustCompile(`\[(?:(?:\+\+){0,128}\+|(?:--){1,128}-)\]`)
		s = oddPlusMinus.ReplaceAllLiteralString(s, "[-]")

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

		// this is not the optimal way to do it,
		// because the counter still counts the cell-reseters themselves,
		// which adds an unwanted bias towards `-`
		plusCount, minusCount := counterPlusMinus(s)
		if plusCount >= minusCount {
			s = strings.ReplaceAll(s, "[-]", "[+]")
			s = strings.ReplaceAll(s, "[--]", "[++]")
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
