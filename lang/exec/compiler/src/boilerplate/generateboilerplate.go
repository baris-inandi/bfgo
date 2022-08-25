package boilerplate

import (
	"fmt"

	"github.com/baris-inandi/brainfuck/lang"
)

func GenerateIRBoilerplate(intermediate string, c lang.Code) string {
	intermediate = "\n\t" + intermediate
	if c.UsingJVM() {
		return fmt.Sprintf(
			JAVA_IR_BOILERPLATE,
			c.GetClassName(),
			c.Context.String("c-cell-type"),
			c.Context.String("c-cell-type"),
			c.Context.Int64("c-tape-size"),
			c.Context.String("c-cell-type"),
			intermediate,
		)
	}
	if c.UsingJS() {
		return fmt.Sprintf(
			JS_IR_BOILERPLATE,
			c.Context.Int64("c-tape-size"),
			intermediate,
		)
	}
	return fmt.Sprintf(
		C_IR_BOILERPLATE,
		c.Context.String("c-cell-type"),
		c.Context.Int64("c-tape-size"),
		c.Context.Int("c-tape-init"),
		c.Context.String("c-cell-type"),
		intermediate,
	)
}
