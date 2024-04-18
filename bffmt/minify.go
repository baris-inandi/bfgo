package bffmt

import (
	"os"
	"regexp"
	"strings"

	"github.com/baris-inandi/bfgo/lang/readcode"
	"github.com/baris-inandi/bfgo/utils"
)

/*
// Minification level
type Level int

// enum emulation
// https://stackoverflow.com/a/14426447
const (

	// only removes non-BF chars,
	// therefore it has maximum portability
	// across implementations.
	BASIC Level = iota
	// assumes code isn't a main program.
	// useful for "libraries" and "modules", such as subroutines.
	LIB Level = iota
	// assume all code will be run as-is
	MAIN Level = iota

)
*/
func MinifyFile( /*l Level,*/ files ...string) {
	// canonical `-` unconditional cell reseter
	const ODD_RESET = "[-]"
	// canonical `-` conditional (break-if-even) cell reseter
	const EVEN_RESET = "[--]"

	// matches any odd (unconditional) reseters, except ODD_RESET
	var isOddReset = regexp.MustCompile(`\[(?:(?:\+\+)*\+|(?:--)+-)\]`)

	// matches any even (conditional break) reseters, except EVEN_RESET
	var isEvenReset = regexp.MustCompile(`\[(?:(?:\+\+)+|(?:--){2,})\]`)

	// matches ODD_RESET, preceded by 1 or more "+" or "-" (mixed)
	var isPrefixedReset = regexp.MustCompile(`[+-]+\[-\]`)

	// returns a pair of indices of matching braces, searched from start.
	// -1 if not found
	//
	// `start` ignores all runes before that index.
	// If `start` is negative, it becomes relative to the end.
	var getMatchingBraces = func(s string, start int) (int, int) {
		start = utils.RelativeIndex(start, len(s))

		open := -1
		for start < len(s) {
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
		// this covers the edge-case where
		// "[" is located just before EOF (start >= size)
		close := -1
		for start < len(s) {
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

	// finds index of 1st byte that isn't in the charset "[],.", or -1 if not found.
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
	var rmLoopLoop = func(s string) string {
		return s
	}

	// removes all loops before any memory write is done.
	//
	// this is safe, because memory is all-zeros, and loops are guaranteed to never run.
	var rm0Loop = func(s string) string {
		for i := 0; i < len(s); i++ {
			c := s[i]
			if c == ',' || c == '+' || c == '-' {
				// can't guarantee cell is 0
				break
			}
			if c == '[' {
				open, close := getMatchingBraces(s, i)
				// assert open == i
				s = s[0:open] + s[close+1:]
			}
		}
		return s
	}

	// removes all bytes after last char in the set ".,]".
	// this ensures `stdin` side effects still happen,
	// and infinite loops are still executed.
	//
	// a mismatched '[' doesn't matter, because it either:
	//
	// 1. continues execution
	//
	// 2. halts/crashes the program
	var rmAfterEffects = func(s string) string {
		// reverse iter
		for i := len(s) - 1; i >= 0; i-- {
			c := s[i]
			if c == '.' || c == ',' || c == ']' {
				s = s[0 : i+1]
				break
			}
		}
		return s
	}

	// # Memory Simulator
	//
	// Statically analyses code, removing some no-ops.
	//
	// It assumes `IOBrace` is a black-box with potential-side effects.
	//
	// current implementation is identity fn
	var memSim = func(s string) string {
		// simulated BF memory/tape
		var mem = map[int]uint8{}
		// relative memory pointer
		var ptr int = 0

		/* # pseudo-code
		0. split s by IOBrace (consecutives are treated as 1).
		1. sim each substr in the resulting array,
		such that each sub has its own isolated mem.
		2. replace each substr by its "canonical form"
		derived from mem.
		3. re-insert delimiters at corresponding positions.
		*/
		// we need an outer loop to cleanup mem.
		// and inner loop should break whenever it finds IOBrace
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
		// So we should iterate over the whole s, rather than a list of pointers to s.
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
		for {
			tmp := s
			s = utils.Apply(
				s,
				// calling this 1st may speed up the others
				rmAfterEffects,
				// order matters, (from this point onwards)
				memSim,
				rmLoopLoop,
				rm0Loop,
			)
			// these 3 are "amplified" by mem-sim
			s = isEvenReset.ReplaceAllLiteralString(s, EVEN_RESET)
			s = isOddReset.ReplaceAllLiteralString(s, ODD_RESET)
			s = isPrefixedReset.ReplaceAllLiteralString(s, ODD_RESET)

			// prevent potential infinite loop and OOM panic
			// by using `>=` rather than `==`
			if len(s) >= len(tmp) {
				// ensure smallest s
				return optimizeCompress(tmp)
			}
		}
	}

	for _, f := range files {
		minified := minify(readcode.ReadBFCode(f))
		err := os.WriteFile(f, []byte(minified), 0o644)
		if err != nil {
			panic(err)
		}
	}
}
