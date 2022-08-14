package optimizer

import (
	"fmt"
	"strings"

	"github.com/baris-inandi/brainfuck/lang"
)

const CANONICALISER_SHIFTING_PATTERN_RUNS = 16
const IR_LITERAL_START = '\\'
const IR_LITERAL_END = '/'

func canonicalise(c lang.Code) lang.Code {

	code := c.Inner
	bindingDebugString := ""

	changeShiftBf := func(loop string, amount int) string {
		// [->+<] -> [->>>+<<<] ; where amount is 3
		return strings.ReplaceAll(
			strings.ReplaceAll(loop, ">", strings.Repeat(">", amount)),
			"<", strings.Repeat("<", amount))
	}
	changeShiftConstBf := func(loop string, shift int, constant int) string {
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

	const BF__ADD_RIGHT = "[->+<]"
	const BF__ADD_RIGHT_ALT = "[>+<-]"
	const BF__ADD_LEFT = "[-<+>]"
	const BF__ADD_LEFT_ALT = "[<+>-]"
	const BF__SUB_RIGHT = "[->-<]"
	const BF__SUB_RIGHT_ALT = "[>-<-]"
	const BF__SUB_LEFT = "[-<->]"
	const BF__SUB_LEFT_ALT = "[<->-]"
	// %s -> multiplier * "+"
	const BF__MUL_LEFT = "[-<%s>]"
	const BF__MUL_LEFT_ALT = "[<%s>-]"
	const BF__MUL_RIGHT = "[->%s<]"
	const BF__MUL_RIGHT_ALT = "[>%s<-]"

	// %d -> shift
	const IR__ADD_RIGHT = "*(p+%d)+=*p;*p=0;"
	const IR__ADD_LEFT = "*(p-%d)+=*p;*p=0;"
	const IR__SUB_RIGHT = "*(p+%d)-=*p;*p=0;"
	const IR__SUB_LEFT = "*(p-%d)-=*p;*p=0;"
	// %d -> shift
	// %d -> constant multiplier
	const IR__MUL_RIGHT = "*(p+%d)+=(*p)*%d;*p=0;"
	const IR__MUL_LEFT = "*(p-%d)+=(*p)*%d;*p=0;"

	// constant patterns, no shift
	code = bindPatternToIR(code, "[-]", "*p=0;")

	// a section where `runs` changes the shift of operation
	runs := CANONICALISER_SHIFTING_PATTERN_RUNS + 1

	c.VerboseOut("canonicalise.go: starting arithmetic canonicalisation with runs parameter ", runs)

	for i := 1; i < runs; i++ {

		for j := 1; j < runs; j++ {
			// i -> shift, j -> constant
			code = bindPatternToIR(code, changeShiftConstBf(BF__MUL_RIGHT, i, j), fmt.Sprintf(IR__MUL_RIGHT, i, j))
			code = bindPatternToIR(code, changeShiftConstBf(BF__MUL_RIGHT_ALT, i, j), fmt.Sprintf(IR__MUL_RIGHT, i, j))

			code = bindPatternToIR(code, changeShiftConstBf(BF__MUL_LEFT, i, j), fmt.Sprintf(IR__MUL_LEFT, i, j))
			code = bindPatternToIR(code, changeShiftConstBf(BF__MUL_LEFT_ALT, i, j), fmt.Sprintf(IR__MUL_LEFT, i, j))
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
