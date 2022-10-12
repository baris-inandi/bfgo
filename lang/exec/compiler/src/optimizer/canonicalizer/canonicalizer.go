package canonicalizer

import (
	"fmt"
	"strings"

	"github.com/baris-inandi/brainfuck/lang"
	"github.com/baris-inandi/brainfuck/lang/exec/compiler/src/optimizer/canonicalizer/canon_constants"
	"github.com/baris-inandi/brainfuck/lang/exec/compiler/src/optimizer/irliteral"
)

const CANONICALIZER_SHIFTING_PATTERN_RUNS = 16

func Canonicalize(c lang.Code) lang.Code {
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
		// changeShiftBf for multiplication by constants
		return strings.ReplaceAll(
			strings.ReplaceAll(
				fmt.Sprintf(loop, strings.Repeat("+", constant)),
				">", strings.Repeat(">", shift)),
			"<", strings.Repeat("<", shift))
	}
	bindPatternToIR := func(code string, pattern string, ir string) string {
		if c.IsDebugging {
			bindingDebugString += pattern + "  " + ir + "\n"
		}
		return strings.ReplaceAll(code, pattern, string(irliteral.IR_LITERAL_START)+ir+string(irliteral.IR_LITERAL_END))
	}
	resolveAndFormatIRFromKey := func(key string, format ...interface{}) string {
		return fmt.Sprintf(canon_constants.ResolveCompileTargetIRBinding(&c, key), format...)
	}

	// constant patterns, no shift
	code = bindPatternToIR(code, "[-]", canon_constants.ResolveCompileTargetIRBinding(&c, "IR__RESET_BYTE"))
	// same with wraparound
	code = bindPatternToIR(code, "[+]", canon_constants.ResolveCompileTargetIRBinding(&c, "IR__RESET_BYTE"))

	// a section where `runs` changes the shift of operation
	runs := CANONICALIZER_SHIFTING_PATTERN_RUNS + 1

	c.VerboseOut("canonicalizer.go: starting arithmetic canonicalisation with runs parameter ", CANONICALIZER_SHIFTING_PATTERN_RUNS)

	for i := 1; i < runs; i++ {

		for j := 2; j < runs; j++ {
			// start with index 2, no need for mul/div by 1 is already implemented using add/sub
			// i -> shift, j -> constant
			code = bindPatternToIR(code, changeShiftUpper(canon_constants.BF__MUL_RIGHT, i, j), resolveAndFormatIRFromKey("IR__MUL_RIGHT", i, j))
			code = bindPatternToIR(code, changeShiftUpper(canon_constants.BF__MUL_RIGHT_ALT, i, j), resolveAndFormatIRFromKey("IR__MUL_RIGHT", i, j))
			code = bindPatternToIR(code, changeShiftUpper(canon_constants.BF__MUL_LEFT, i, j), resolveAndFormatIRFromKey("IR__MUL_LEFT", i, j))
			code = bindPatternToIR(code, changeShiftUpper(canon_constants.BF__MUL_LEFT_ALT, i, j), resolveAndFormatIRFromKey("IR__MUL_LEFT", i, j))
		}

		// patterns that add right
		code = bindPatternToIR(code, changeShiftBf(canon_constants.BF__ADD_RIGHT, i), resolveAndFormatIRFromKey("IR__ADD_RIGHT", i))
		code = bindPatternToIR(code, changeShiftBf(canon_constants.BF__ADD_RIGHT_ALT, i), resolveAndFormatIRFromKey("IR__ADD_RIGHT", i))

		// patterns that add left
		code = bindPatternToIR(code, changeShiftBf(canon_constants.BF__ADD_LEFT, i), resolveAndFormatIRFromKey("IR__ADD_LEFT", i))
		code = bindPatternToIR(code, changeShiftBf(canon_constants.BF__ADD_LEFT_ALT, i), resolveAndFormatIRFromKey("IR__ADD_LEFT", i))

		// patterns that subtract left
		code = bindPatternToIR(code, changeShiftBf(canon_constants.BF__SUB_LEFT, i), resolveAndFormatIRFromKey("IR__SUB_LEFT", i))
		code = bindPatternToIR(code, changeShiftBf(canon_constants.BF__SUB_LEFT_ALT, i), resolveAndFormatIRFromKey("IR__SUB_LEFT", i))

		// patterns that subtract right
		code = bindPatternToIR(code, changeShiftBf(canon_constants.BF__SUB_RIGHT, i), resolveAndFormatIRFromKey("IR__SUB_RIGHT", i))
		code = bindPatternToIR(code, changeShiftBf(canon_constants.BF__SUB_RIGHT_ALT, i), resolveAndFormatIRFromKey("IR__SUB_RIGHT", i))

	}
	c.Inner = code
	c.AddDebugFile("ARITHMETIC_BINDINGS", bindingDebugString)
	return c
}
