package canon_constants

import (
	"github.com/baris-inandi/bfgo/lang"
	"github.com/baris-inandi/bfgo/lang/exec/compiler/src/optimizer/canonicalizer/canon_constants/ir_bindings"
)

func ResolveCompileTargetIRBinding(c *lang.Code, key string) string {
	if c.UsingJVM() {
		return ir_bindings.JAVAIR[key]
	} else if c.UsingJS() {
		return ir_bindings.JSIR[key]
	} else {
		return ir_bindings.CIR[key]
	}
}
