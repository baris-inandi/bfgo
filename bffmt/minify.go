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

	// # Compression Ratio Optimizer
	//
	// Uses [frequency analysis] to find the best chars to replace.
	//
	// Current implementation only replaces cell-reseters.
	//
	// [frequency analysis]: https://en.wikipedia.org/wiki/Frequency_analysis
	var optimizeCompress = func(s string) string {
		plusCount, minusCount := counterPlusMinus(s)
		// this isn't the best way to do it,
		// because the counter still counts the cell-reseters themselves,
		// which adds a bias towards "-"
		if plusCount >= minusCount {
			s = strings.ReplaceAll(s, "[-]", "[+]")
			s = strings.ReplaceAll(s, "[--]", "[++]")
		}
		return s
	}

	// Optimizes the source, as explained in https://github.com/baris-inandi/brainfuck-go/issues/2 .
	// It assumes `s` only has valid ops.
	var minify = func(s string) string {
		// matches even-count cell-reseters
		var evenPlusMinus = regexp.MustCompile(`\[(?:(?:\+\+){1,128}|(?:--){2,128})\]`)
		s = evenPlusMinus.ReplaceAllLiteralString(s, "[--]")

		// matches odd-count cell-reseters
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

		return optimizeCompress(s)
	}

	for _, f := range files {
		minified := minify(readcode.ReadBrainfuck(f))
		err := os.WriteFile(f, []byte(minified), 0o644)
		if err != nil {
			panic(err)
		}
	}
}
