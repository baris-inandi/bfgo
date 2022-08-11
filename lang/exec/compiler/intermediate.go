package compiler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/baris-inandi/brainfuck/lang"
	"github.com/baris-inandi/brainfuck/lang/exec/compiler/optimizer"
)

func Intermediate(c lang.Code) string {
	// transforms brainfuck code to intermediate representation and returns a string
	code := c.Content
	ctx := c.Context

	if code == "" {
		return ""
	}
	intermediate := "\n\t"
	code += "/"
	prevChar := ""
	repeatedCharCounter := uint16(1)
	initialRepeat := false
	depth := 0
	optimized := ctx.Bool("optimize")
	inLiteral := false
	skipChars := 0
	if optimized {
		code = optimizer.Canonicalise(code)
	}
	for idx, char := range code {
		char := string(char)
		if skipChars > 0 {
			skipChars--
			continue
		}
		if initialRepeat {
			if inLiteral {
				if prevChar != optimizer.IR_LITERAL_START {
					intermediate += prevChar
				}
			}
			if prevChar == char && (prevChar == "+" || prevChar == "-" || char == "<" || char == ">") {
				repeatedCharCounter += 1
			} else {
				rep := strconv.Itoa(int(repeatedCharCounter))
				switch prevChar {
				case "<":
					intermediate += ("p-=" + rep + ";")
				case ">":
					intermediate += ("p+=" + rep + ";")
				case "+":
					intermediate += ("*p+=" + rep + ";")
				case "-":
					intermediate += ("*p-=" + rep + ";")
				case ".":
					intermediate += ("putchar(*p);")
				case ",":
					intermediate += ("*p=getchar();")
				case "[":
					depth++
					intermediate += ("while (*p){")
				case "]":
					depth--
					intermediate += ("};")
				case optimizer.IR_LITERAL_START:
					i := idx
					current := string(code[i])
					literal := ""
					for current != optimizer.IR_LITERAL_END {
						i++
						literal += current
						current = string(code[i])
					}
					intermediate += literal
					skipChars += len(literal)
				}
				repeatedCharCounter = 1
			}
		} else {
			initialRepeat = true
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
	intermediate = fmt.Sprintf(IR, intermediate)
	if c.Context.Bool("dump-ir") {
		fmt.Println(intermediate)
	}
	return intermediate
}
