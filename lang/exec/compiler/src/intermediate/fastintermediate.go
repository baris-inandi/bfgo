package intermediate

import (
	"fmt"

	"github.com/baris-inandi/bfgo/lang"
	"github.com/baris-inandi/bfgo/lang/exec/compiler/src/boilerplate"
	"github.com/baris-inandi/bfgo/lang/exec/compiler/src/boilerplate/ir_constants"
)

func FastGenerateIntermediateRepresentation(c lang.Code) string {
	// transforms BF code to intermediate representation and returns a string
	if c.Inner == "" {
		return ""
	}
	intermediate := ""
	for _, char := range c.Inner {
		intermediate += ir_constants.ResolveCompileTargetIR(&c, string(char))
	}
	intermediate = boilerplate.GenerateIRBoilerplate(intermediate, c)
	if c.Context.Bool("d-dump-ir") {
		fmt.Println(intermediate)
	}
	c.AddDebugFile("IR."+c.CompileTarget, intermediate)
	return intermediate
}
