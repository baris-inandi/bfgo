package optimizer

import (
	"github.com/baris-inandi/bfgo/lang"
	"github.com/baris-inandi/bfgo/lang/exec/compiler/src/optimizer/canonicalizer"
	"github.com/baris-inandi/bfgo/lang/exec/compiler/src/optimizer/deadcode"
)

func Optimize(c lang.Code) lang.Code {
	c.VerboseOut("optimize.go: starting optimizer")
	return canonicalizer.Canonicalize(
		deadcode.RemoveLeadingDeadcode(
			c,
		),
	)
}
