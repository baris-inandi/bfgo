package bffmt

import (
	"os"
	"regexp"
	"strings"

	"github.com/baris-inandi/brainfuck/lang/readcode"
	"github.com/baris-inandi/brainfuck/utils"
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

	// returns a pair of indices of matching braces, searched from start.
	// -1 if not found
	//
	// `start` ignores all runes before that index.
	// If start is negative, it becomes relative to the end.
	var getMatchingBraces = func(s string, start int) (int, int) {
		size := len(s)
		start = utils.RelativeIndex(start, size)

		open := -1
		for start < size {
			c := s[start]
			if c == '[' {
				open = start
				break
			}
			start += 1
		}

		// avoid double-counting "["
		start++
		depth := 0

		close := -1
		for start < size {
			c := s[start]
			if c == '[' {
				depth++
			}
			if c == ']' {
				if depth == 0 {
					close = start
					break
				}
				depth--
			}
			start += 1
		}

		return open, close
	}

	// finds index of 1st rune that isn't in the charset "[],.", or -1 if not found.
	//
	// `start` ignores all runes before that index.
	// If start is negative, it becomes relative to the end.
	var indexNoIOBrace = func(s string, start int) int {
		size := len(s)
		start = utils.RelativeIndex(start, size)

		// "I couldn't find a way to write it in functional-paradigm" @Rudxain
		for start < size {
			c := s[start]
			if c != '[' && c != ']' && c != ',' && c != '.' {
				return start
			}
			start += 1
		}
		return -1
	}

	// removes consecutive loops, keeping the 1st.
	//
	// current impl is identity fn
	var removeConsecutiveLoop = func(s string) string {
		return s
	}

	// removes all loops before any memory write is done.
	//
	// this is safe, because memory is all-zeros, and loops are guaranteed to never run.
	//
	// current impl is identity fn
	var zeroLoopRemover = func(s string) string {
		return s
	}

	// removes all runes after last ".",
	// [iff] there's no `,` or `]` in the part to be removed,
	// this ensures `stdin` side effects still happen,
	// and infinite loops are still executed.
	//
	// a mismatched `[` doesn't matter, because it either:
	//
	// 1. continues execution
	//
	// 2. halts/crashes the program
	//
	// [iff]: https://en.wikipedia.org/wiki/If_and_only_if
	var removeAfterLastDot = func(s string) string {
		for i := len(s) - 1; i >= 0; i-- {
			c := s[i]
			if c == ',' || c == ']' {
				break
			}
			if c == '.' {
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
	// current implementation is identity fn
	var memSimulator = func(s string) string {
		// simulated BF memory/tape
		var mem = map[int]uint8{}
		// relative memory pointer
		var ptr int = 0

		for i := indexNoIOBrace(s, 0); i < len(s) && i > -1; i = indexNoIOBrace(s, i+1) {
			switch s[i] {
			case '+':
				{
					mem[ptr] += 1
					continue
				}
			case '-':
				{
					mem[ptr] -= 1
					continue
				}
			case '>':
				{
					ptr += 1
					continue
				}
			case '<':
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
		// "I hope the compiler optimizes this from 4n iterations to n iters"
		// @Rudxain
		plus, minus := strings.Count(s, "+"), strings.Count(s, "-")
		odd, even := strings.Count(s, ODD_RESET), strings.Count(s, EVEN_RESET)
		// this ensures the choice is unbiased
		isMorePlusThanMinus := plus-minus+odd+2*even > 0

		// A space-time tradeoff isn't worth it,
		// because time is O(n) and space is O(1) (ignoring s).
		// If (while counting) we were to allocate a list of indices to all ocurrences
		// of ODD_RESET and EVEN_RESET, space would become O(n),
		// but time would still be O(n) (despite being practically faster).
		// So we should iterate over the whole s, instead of iter over a list of pointers to s.
		//
		// CPU cache already helps a bit.
		// allocating more memory just reduces the available cache space,
		// therefore reducing iteration speed
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
	// [#2]: https://github.com/baris-inandi/brainfuck-go/issues/2
	var minify = func(s string) string {
		s = utils.Apply(
			s,
			// calling this 1st may speed up the others
			removeAfterLastDot,
			// order matters, (from this point onwards)
			memSimulator,
			removeConsecutiveLoop,
			zeroLoopRemover,
		)
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
