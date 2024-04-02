package ir_constants

import "github.com/baris-inandi/bfgo/lang"

func ResolveCompileTargetIR(c *lang.Code, key string) string {
	if c.UsingJVM() {
		return JAVAIR[key]
	} else if c.UsingJS() {
		return JSIR[key]
	} else {
		return CIR[key]
	}
}
