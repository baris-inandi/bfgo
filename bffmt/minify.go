package bffmt

import (
	"os"
	"regexp"
	"strings"

	"github.com/baris-inandi/brainfuck/lang/readcode"
)

func MinifyFile(files ...string) {
	// minified `-` uncondtional cell reseter
	const ODD_RESET = "[-]"
	// minified `-` condtional (halt-if-even) cell reseter
	const EVEN_RESET = "[--]"

	// matches any odd (unconditional) reseters, except ODD_RESET
	var isOddReset = regexp.MustCompile(`\[(?:(?:\+\+)*\+|(?:--)+-)\]`)

	// matches any even (conditional halt) reseters, except EVEN_RESET
	var isEvenReset = regexp.MustCompile(`\[(?:(?:\+\+)+|(?:--){2,})\]`)

	// matches ODD_RESET, preceded by 1 or more "+" or "-" (mixed)
	var isPrefixedReset = regexp.MustCompile(`[+-]+\[-\]`)

	/* // returns a pair of indices of matching braces, searched from start.
	//
	// `start` ignores all runes before that index.
	// If start is negative, it becomes relative to the end.
	var getMatchingBraces = func(s string, start int) (int, int) {
		size := len(s)
		if start < 0 {
			start += size
		}
		if start < 0 {
			panic("Index out of bounds")
		}

		var open int = -1
		for start < size {
			c := s[start]
			if c == "["[0] {
				open = start
				break
			}
			start += 1
		}

		// avoid double-counting "["
		start++
		var depth int = 0

		var close int = -1
		for start < size {
			c := s[start]
			if c == "["[0] {
				depth++
			}
			if c == "]"[0] {
				if depth == 0 {
					close = start
					break
				}
				depth--
			}
			start += 1
		}

		return open, close
	} */

	// finds index of 1st rune that isn't in the charset "[],.", or -1 if not found.
	//
	// `start` ignores all runes before that index.
	// If start is negative, it becomes relative to the end.
	var indexNoIOBrace = func(s string, start int) int {
		size := len(s)
		if start < 0 {
			start += size
		}
		if start < 0 {
			// should it return an err instead of panicking?
			panic("Index out of bounds")
		}

		// "I couldn't find a way to write it in functional-paradigm" @Rudxain
		for start < size {
			c := s[start]
			if c != "["[0] && c != "]"[0] && c != ","[0] && c != "."[0] {
				return start
			}
			start += 1
		}
		return -1
	}

	// removes all runes after last `.`,
	// [iff] there are no braces in the part to be removed
	// (this ensures infinite loops are still executed)
	//
	//[iff]: https://en.wikipedia.org/wiki/If_and_only_if
	var noOutputRemover = func(s string) string {
		for i := len(s) - 1; i >= 0; i-- {
			c := s[i]
			if c == "["[0] || c == "]"[0] {
				break
			}
			if c == "."[0] {
				s = s[0 : i+1]
				break
			}
		}
		return s
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

		for i := indexNoIOBrace(s, 0); i < len(s) && i != -1; i = indexNoIOBrace(s, i+1) {
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
			}
		}
		return s
	}

	// # Compression Ratio Optimizer
	//
	// Uses [frequency analysis] to increase compression-ratio by 3rd-party algorithms.
	//
	// Current implementation only replaces minified "-" reseters.
	// It assumes there's no "+" reseters.
	//
	// [frequency analysis]: https://en.wikipedia.org/wiki/Frequency_analysis
	var optimizeCompress = func(s string) string {
		plus, minus := strings.Count(s, "+"), strings.Count(s, "-")
		odd, even := strings.Count(s, ODD_RESET), strings.Count(s, EVEN_RESET)
		// this ensures the choice is unbiased
		isMorePlusThanMinus := plus-minus+odd+2*even > 0

		if isMorePlusThanMinus {
			if odd > 0 {
				s = strings.ReplaceAll(s, ODD_RESET, "[+]")
			}
			if even > 0 {
				s = strings.ReplaceAll(s, EVEN_RESET, "[++]")
			}
		}
		return s
	}

	// # Advanced BF minifier
	//
	// Explained in [#2]. It assumes s only has valid ops.
	//
	//[#2]: https://github.com/baris-inandi/brainfuck-go/issues/2
	var minify = func(s string) string {
		// calling this now may speed up the others
		s = noOutputRemover(s)
		// order matters, (from this point onwards)
		s = memSimulator(s)
		s = isEvenReset.ReplaceAllLiteralString(s, EVEN_RESET)
		s = isOddReset.ReplaceAllLiteralString(s, ODD_RESET)
		s = isPrefixedReset.ReplaceAllLiteralString(s, ODD_RESET)
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
