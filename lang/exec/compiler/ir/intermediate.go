package ir

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/baris-inandi/brainfuck/lang"
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
	patternBindings := []PatternBindingPair{}
	if optimized {
		patternBindings = GeneratePatternBindings()
		for idx, binding := range patternBindings {
			code = strings.ReplaceAll(code, binding.pattern, "("+strconv.Itoa(idx)+")")
		}
	}
	for idx, char := range code {
		char := string(char)
		if initialRepeat {
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
				case "(":
					i := idx
					current := string(code[i])
					bindingIdentifierStr := ""
					for current != ")" {
						i++
						bindingIdentifierStr += current
						current = string(code[i])
					}
					bindingIdentifier, _ := strconv.Atoi(bindingIdentifierStr)
					intermediate += (patternBindings[bindingIdentifier].ir)
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
