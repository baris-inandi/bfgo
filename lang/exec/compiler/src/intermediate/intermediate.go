package intermediate

import (
	"fmt"
	"os"

	"github.com/baris-inandi/bfgo/lang"
	"github.com/baris-inandi/bfgo/lang/exec/compiler/src/boilerplate"
	"github.com/baris-inandi/bfgo/lang/exec/compiler/src/boilerplate/ir_constants"
	"github.com/baris-inandi/bfgo/lang/exec/compiler/src/optimizer"
	"github.com/baris-inandi/bfgo/lang/exec/compiler/src/optimizer/canonicalizer"
	"github.com/baris-inandi/bfgo/lang/exec/compiler/src/optimizer/irliteral"
)

func GenerateIntermediateRepresentation(c lang.Code) string {
	// transforms BF code to intermediate representation
	if c.Inner == "" {
		return ""
	}
	intermediate := ""
	prevChar := ' '
	depth := int32(0)
	repSymbolCount := uint16(1)
	inLiteral := false
	skipChars := 0
	c.AddDebugFile("SOURCE.b", c.Inner)
	if c.OLevel == 3 {
		c = optimizer.Optimize(c)
		c.AddDebugFile("SOURCE-OPTIMIZED.b", c.Inner)
	}
	c.Inner += "\n"
	for idx, char := range c.Inner {
		if skipChars > 0 {
			skipChars--
			continue
		}
		if inLiteral {
			if prevChar != irliteral.IR_LITERAL_START {
				intermediate += string(prevChar)
			}
		}
		if prevChar == char && (prevChar == '+' ||
			prevChar == '-' ||
			char == '<' ||
			char == '>') {
			repSymbolCount += 1
			continue
		}
		switch prevChar {
		case '<':
			intermediate += fmt.Sprintf(ir_constants.ResolveCompileTargetIR(&c, "LEFT_ANGLE_REP"), repSymbolCount)
		case '>':
			intermediate += fmt.Sprintf(ir_constants.ResolveCompileTargetIR(&c, "RIGHT_ANGLE_REP"), repSymbolCount)
		case '+':
			intermediate += fmt.Sprintf(ir_constants.ResolveCompileTargetIR(&c, "PLUS_REP"), repSymbolCount)
		case '-':
			intermediate += fmt.Sprintf(ir_constants.ResolveCompileTargetIR(&c, "MINUS_REP"), repSymbolCount)
		case '.':
			intermediate += ir_constants.ResolveCompileTargetIR(&c, ".")
		case ',':
			intermediate += ir_constants.ResolveCompileTargetIR(&c, ",")
		case '[':
			depth++
			intermediate += ir_constants.ResolveCompileTargetIR(&c, "[")
		case ']':
			depth--
			intermediate += ir_constants.ResolveCompileTargetIR(&c, "]")
		case irliteral.IR_LITERAL_START:
			i := idx
			current := string(c.Inner[i])
			literal := ""
			for current != string(irliteral.IR_LITERAL_END) {
				i++
				literal += current
				current = string(c.Inner[i])
			}
			intermediate += literal
			skipChars += len(literal)
		}
		repSymbolCount = 1
		prevChar = char
	}
	if depth > 0 {
		fmt.Println("Syntax error: Unmatched [")
		os.Exit(1)
	} else if depth < 0 {
		fmt.Println("Syntax error: Unmatched ]")
		os.Exit(1)
	}

	// an experimental optimization applied post ir generation
	if c.OLevel == 3 && c.CompileTarget == "c" {
		intermediate = canonicalizer.ExperimentalCResetIncDecCanon(intermediate)
	}

	intermediate = boilerplate.GenerateIRBoilerplate(intermediate, c)
	if c.Context.Bool("d-dump-ir") {
		fmt.Println(intermediate)
	}
	c.AddDebugFile("IR."+c.CompileTarget, intermediate)
	return intermediate
}
