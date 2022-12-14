package canonicalizer

import (
	"strings"
)

func ExperimentalCResetIncDecCanon(ir string) string {
	// experimental reference reset optimization
	// for example, RESET and INCREMENT(5) can be interpreted as a simple SET(5)
	out := strings.ReplaceAll(ir, "*p=0;*p+=", "*p=")
	return out
}
