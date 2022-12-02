package bffmt

import (
	"os"
	"regexp"
	"strings"

	"github.com/baris-inandi/brainfuck/lang/readcode"
)

func MinifyFile(files ...string) {
	// matches even cell-reseters
	var isEvenReset = regexp.MustCompile(`\[(?:(?:\+\+){1,128}|(?:--){2,128})\]`)

	// matches odd cell-reseters
	var isOddReset = regexp.MustCompile(`\[(?:(?:\+\+){0,128}\+|(?:--){1,128}-)\]`)

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

	// finds index of 1st rune that isn't "[" or "]", or -1 if not found.
	//
	// `start` ignores all runes before that index.
	// If start is negative, it becomes relative to the end.
	var indexNoBrace = func(s string, start int) int {
		size := len(s)
		if start < 0 {
			start += size
		}
		if start < 0 {
			// should it return an err instead of panicking?
			panic("Index out of bounds")
		}

		// "I couldn't find a way to write it as funtional-paradigm" @Rudxain
		for start < size {
			if s[start] != "["[0] && s[start] != "]"[0] {
				return start
			}
			start += 1
		}
		return -1
	}

	// # Memory Simulator
	//
	// Statically analyses non-loop code, to remove some no-ops.
	//
	// current implementation is no-op itself,
	// so this func is equivalent to the identity fn
	var memSimulator = func(s string) string {
		// simulated BF memory/tape
		var mem = map[int]uint8{}
		// relative memory pointer
		var ptr int = 0

	label:
		for i := indexNoBrace(s, 0); i < len(s); i += 1 {
			if i == -1 {
				break
			}

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
		}
		return s
	}

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
	// It assumes s only has valid ops.
	var minify = func(s string) string {
		// order matters, A LOT
		s = memSimulator(s)
		s = isEvenReset.ReplaceAllLiteralString(s, "[--]")
		s = isOddReset.ReplaceAllLiteralString(s, "[-]")
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
