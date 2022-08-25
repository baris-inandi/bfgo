package optimizer

import (
	"github.com/baris-inandi/brainfuck/lang"
	"github.com/baris-inandi/brainfuck/lang/exec/compiler/optimizer/canonicalizer"
	"github.com/baris-inandi/brainfuck/lang/exec/compiler/optimizer/deadcode"
)

func Optimize(c lang.Code) lang.Code {
	c.VerboseOut("optimize.go: starting optimizer")
	return canonicalizer.Canonicalize(
		deadcode.RemoveLeadingDeadcode(
			c,
		),
	)
}
