package deadcode

import (
	"github.com/baris-inandi/brainfuck/lang"
)

func RemoveLeadingDeadcode(c lang.Code) lang.Code {
	code := c.Inner
	// remove until last [ ] . or ,
	// leading <>+- operators have no effect on output
	c.VerboseOut("deadcode.go: removing unneeded leading operators")
	removeLast := 0
	for i := len(code) - 1; i >= 0; i-- {
		char := code[i]
		if char == '[' ||
			char == ']' ||
			char == '.' ||
			char == ',' {
			break
		}
		removeLast += 1
	}
	code = string([]rune(code)[:len(code)-removeLast])
	c.Inner = code
	return c
}
