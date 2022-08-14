package compiler

import (
	"fmt"
	"os"

	"github.com/baris-inandi/brainfuck/lang"
)

func FastGenerateIntermediateRepresentation(c lang.Code) string {
	// transforms brainfuck code to intermediate representation and returns a string
	if c.Inner == "" {
		return ""
	}
	depth := int32(0)
	intermediate := ""
	for _, char := range c.Inner {
		switch char {
		case '<':
			intermediate += ("--p;")
		case '>':
			intermediate += ("++p;")
		case '+':
			intermediate += ("++*p;")
		case '-':
			intermediate += ("--*p;")
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
		}
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
