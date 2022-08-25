package canon_constants

import (
	"github.com/baris-inandi/brainfuck/lang"
	"github.com/baris-inandi/brainfuck/lang/exec/compiler/optimizer/canonicalizer/canon_constants/ir_bindings"
)

func ResolveCompileTargetIR(c *lang.Code, key string) string {
	if c.UsingJVM() {
		return ir_bindings.JAVAIR[key]
	} else {
		return ir_bindings.CIR[key]
	}
}
