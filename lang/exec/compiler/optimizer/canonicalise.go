package optimizer

import (
	"fmt"
	"strings"

	"github.com/baris-inandi/brainfuck/lang"
)

const CANONICALISER_SHIFTING_PATTERN_RUNS = 16
const IR_LITERAL_START = '~'
const IR_LITERAL_END = '`'

func canonicalise(c lang.Code) lang.Code {
	code := c.Inner
	bindingDebugString := ""

	code = strings.ReplaceAll(code, "+-", "")
	code = strings.ReplaceAll(code, "-+", "")

	changeShiftBf := func(loop string, amount int) string {
		// [->+<] -> [->>>+<<<] ; where amount is 3
		return strings.ReplaceAll(
			strings.ReplaceAll(loop, ">", strings.Repeat(">", amount)),
			"<", strings.Repeat("<", amount))
	}
	changeShiftUpper := func(loop string, shift int, constant int) string {
		return strings.ReplaceAll(
			strings.ReplaceAll(
				fmt.Sprintf(loop, strings.Repeat("+", constant)),
				">", strings.Repeat(">", shift)),
			"<", strings.Repeat("<", shift))
	}
	bindPatternToIR := func(code string, pattern string, ir string) string {
		bindingDebugString += pattern + "  " + ir + "\n"
		return strings.ReplaceAll(code, pattern, string(IR_LITERAL_START)+ir+string(IR_LITERAL_END))
	}

	// constant patterns, no shift
	code = bindPatternToIR(code, "[-]", "*p=0;")

	// a section where `runs` changes the shift of operation
	runs := CANONICALISER_SHIFTING_PATTERN_RUNS + 1

	c.VerboseOut("canonicalise.go: starting arithmetic canonicalisation with runs parameter ", CANONICALISER_SHIFTING_PATTERN_RUNS)

	for i := 1; i < runs; i++ {

		for j := 2; j < runs; j++ {
			// start with index 2, no need for mul/div by 1 is already implemented using add/sub
			// i -> shift, j -> constant
			code = bindPatternToIR(code, changeShiftUpper(BF__MUL_RIGHT, i, j), fmt.Sprintf(IR__MUL_RIGHT, i, j))
			code = bindPatternToIR(code, changeShiftUpper(BF__MUL_RIGHT_ALT, i, j), fmt.Sprintf(IR__MUL_RIGHT, i, j))
			code = bindPatternToIR(code, changeShiftUpper(BF__MUL_LEFT, i, j), fmt.Sprintf(IR__MUL_LEFT, i, j))
			code = bindPatternToIR(code, changeShiftUpper(BF__MUL_LEFT_ALT, i, j), fmt.Sprintf(IR__MUL_LEFT, i, j))
		}

		// patterns that add right
		code = bindPatternToIR(code, changeShiftBf(BF__ADD_RIGHT, i), fmt.Sprintf(IR__ADD_RIGHT, i))
		code = bindPatternToIR(code, changeShiftBf(BF__ADD_RIGHT_ALT, i), fmt.Sprintf(IR__ADD_RIGHT, i))

		// patterns that add left
		code = bindPatternToIR(code, changeShiftBf(BF__ADD_LEFT, i), fmt.Sprintf(IR__ADD_LEFT, i))
		code = bindPatternToIR(code, changeShiftBf(BF__ADD_LEFT_ALT, i), fmt.Sprintf(IR__ADD_LEFT, i))

		// patterns that subtract left
		code = bindPatternToIR(code, changeShiftBf(BF__SUB_LEFT, i), fmt.Sprintf(IR__SUB_LEFT, i))
		code = bindPatternToIR(code, changeShiftBf(BF__SUB_LEFT_ALT, i), fmt.Sprintf(IR__SUB_LEFT, i))

		// patterns that subtract right
		code = bindPatternToIR(code, changeShiftBf(BF__SUB_RIGHT, i), fmt.Sprintf(IR__SUB_RIGHT, i))
		code = bindPatternToIR(code, changeShiftBf(BF__SUB_RIGHT_ALT, i), fmt.Sprintf(IR__SUB_RIGHT, i))

	}
	c.Inner = code
	c.AddDebugFile("ARITHMETIC_BINDINGS", bindingDebugString)
	return c
}
