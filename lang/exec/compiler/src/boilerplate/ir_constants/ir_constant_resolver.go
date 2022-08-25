package ir_constants

import "github.com/baris-inandi/brainfuck/lang"

func ResolveCompileTargetIR(c *lang.Code, key string) string {
	if c.UsingJVM() {
		return JAVAIR[key]
	} else {
		return CIR[key]
	}
}
