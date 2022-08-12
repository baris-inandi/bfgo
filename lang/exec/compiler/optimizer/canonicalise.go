package optimizer

import (
	"fmt"
	"strings"
)

const CANONICALISER_SHIFTING_PATTERN_RUNS = 16
const IR_LITERAL_START = "\\"
const IR_LITERAL_END = "/"

func Canonicalise(code string) string {
	changeShiftBf := func(loop string, amount int) string {
		// [->+<] -> [->>>+<<<] ; where amount is 3
		return strings.ReplaceAll(
			strings.ReplaceAll(loop, ">", strings.Repeat(">", amount)),
			"<", strings.Repeat("<", amount))
	}
	bindPatternToIR := func(code string, pattern string, ir string) string {
		return strings.ReplaceAll(code, pattern, IR_LITERAL_START+ir+IR_LITERAL_END)
	}

	const BF__ADD_RIGHT = "[->+<]"
	const BF__ADD_RIGHT_ALT = "[>+<-]"
	const BF__ADD_LEFT = "[-<+>]"
	const BF__ADD_LEFT_ALT = "[<+>-]"
	const BF__SUB_RIGHT = "[->-<]"
	const BF__SUB_RIGHT_ALT = "[>-<-]"
	const BF__SUB_LEFT = "[-<->]"
	const BF__SUB_LEFT_ALT = "[<->-]"

	const IR__ADD_RIGHT = "*(p+%d)+=*p;*p=0;"
	const IR__ADD_LEFT = "*(p-%d)+=*p;*p=0;"
	const IR__SUB_RIGHT = "*(p+%d)-=*p;*p=0;"
	const IR__SUB_LEFT = "*(p-%d)-=*p;*p=0;"

	// constant patterns, no shift
	code = bindPatternToIR(code, "[-]", "*p=0;")

	// a section where `runs` changes the shift of operation
	runs := CANONICALISER_SHIFTING_PATTERN_RUNS + 1
	for i := 1; i < runs; i++ {

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
	return code
}
