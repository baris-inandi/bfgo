package compiler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/baris-inandi/brainfuck/lang"
	"github.com/baris-inandi/brainfuck/lang/exec/compiler/optimizer"
	"github.com/baris-inandi/brainfuck/lang/exec/compiler/optimizer/irliteral"
)

func GenerateIntermediateRepresentation(c lang.Code) string {
	// transforms brainfuck code to intermediate representation
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
		rep := strconv.Itoa(int(repSymbolCount))
		switch prevChar {
		case '<':
			intermediate += ("p-=" + rep + ";")
		case '>':
			intermediate += ("p+=" + rep + ";")
		case '+':
			intermediate += ("*p+=" + rep + ";")
		case '-':
			intermediate += ("*p-=" + rep + ";")
		case '.':
			intermediate += ("putc(*p, stdout);")
		case ',':
			intermediate += ("*p=getchar();")
		case '[':
			depth++
			intermediate += ("while (*p){")
		case ']':
			depth--
			intermediate += ("};")
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
	intermediate = SprintfIR(intermediate, c)
	if c.Context.Bool("d-dump-ir") {
		fmt.Println(intermediate)
	}
	c.AddDebugFile("IR.c", intermediate)
	return intermediate
}
