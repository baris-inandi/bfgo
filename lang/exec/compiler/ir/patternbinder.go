package ir

import (
	"fmt"
	"strings"
)

type PatternBindingPair struct {
	pattern string
	ir      string
}

func GeneratePatternBindings() []PatternBindingPair {
	generateAddSubBindings := func(runs int) []PatternBindingPair {
		changeShift := func(loop string, amount int) string {
			// [->+<] -> [->>>+<<<] ; where amount is 3
			return strings.ReplaceAll(
				strings.ReplaceAll(loop, ">", strings.Repeat(">", amount)),
				"<", strings.Repeat("<", amount))
		}
		generateIr := func(shift int, isAdd bool) string {
			if isAdd {
				return fmt.Sprintf("*(p+%d)+=*p;*p=0;", shift)
			}
			return fmt.Sprintf("*(p-%d)-=*p;*p=0;", shift)
		}
		out := []PatternBindingPair{}
		const ADD_RIGHT_INIT = "[->+<]"     // add current to right
		const ADD_RIGHT_ALT_INIT = "[>+<-]" // add current to right (alt notation)
		const SUB_LEFT_INIT = "[-<->]"      // subtract left from current
		const SUB_LEFT_ALT_INIT = "[<->-]"  // subtract left from current (alt notation)
		runsInc := runs + 1
		for i := 1; i < runsInc; i++ {
			out = append(append(append(append(
				out, PatternBindingPair{pattern: changeShift(ADD_RIGHT_INIT, i), ir: generateIr(i, true)}),
				PatternBindingPair{pattern: changeShift(ADD_RIGHT_ALT_INIT, i), ir: generateIr(i, true)}),
				PatternBindingPair{pattern: changeShift(SUB_LEFT_INIT, i), ir: generateIr(i, false)}),
				PatternBindingPair{pattern: changeShift(SUB_LEFT_ALT_INIT, i), ir: generateIr(i, false)})
		}
		return out
	}

	patterns := []PatternBindingPair{
		{ // reset byte, constant 0
			pattern: "[-]",
			ir:      "*p=0;"},
	}
	patterns = append(patterns, generateAddSubBindings(24)...)
	return patterns
}
