package optimizer

import (
	"github.com/baris-inandi/brainfuck/lang"
)

func Optimize(c lang.Code) lang.Code {
	c.VerboseOut("optimize.go: starting optimizer")
	return canonicalise(
		removeUnusedLeading(
			c,
		),
	)
}
