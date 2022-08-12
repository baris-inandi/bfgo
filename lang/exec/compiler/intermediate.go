package compiler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/baris-inandi/brainfuck/lang"
	"github.com/baris-inandi/brainfuck/lang/exec/compiler/optimizer"
)

func GenerateIntermediateRepresentation(c lang.Code) string {
	// transforms brainfuck code to intermediate representation and returns a string
	code := c.Content
	if code == "" {
		return ""
	}
	intermediate := "\n\t"
	prevChar := ' '
	depth := int32(0)
	repSymbolCount := uint16(1)
	inLiteral := false
	skipChars := 0
	if c.Context.Bool("o-performance") {
		code = optimizer.Optimize(code)
	}
	code += "\n"
	for idx, char := range code {
		if skipChars > 0 {
			skipChars--
			continue
		}
		if inLiteral {
			if prevChar != optimizer.IR_LITERAL_START {
				intermediate += string(prevChar)
			}
		}
		if prevChar == char && (prevChar == '+' ||
			prevChar == '-' ||
			char == '<' ||
			char == '>') {
			repSymbolCount += 1
		} else {
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
			case optimizer.IR_LITERAL_START:
				i := idx
				current := string(code[i])
				literal := ""
				for current != string(optimizer.IR_LITERAL_END) {
					i++
					literal += current
					current = string(code[i])
				}
				intermediate += literal
				skipChars += len(literal)
			}
			repSymbolCount = 1
		}
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
	if c.Context.Bool("dump-ir") {
		fmt.Println(intermediate)
	}
	return intermediate
}
