package compiler

import (
	"fmt"

	"github.com/baris-inandi/brainfuck/lang"
)

const IR_BOILERPLATE string = `#include <stdio.h>
int main()
{
    %s t[%d] = {%d};
    %s *p = t;
    // ir %s
    return 0;
}
`

func SprintfIR(intermediate string, c lang.Code) string {
	return fmt.Sprintf(
		IR_BOILERPLATE,
		c.Context.String("c-cell-type"),
		c.Context.Int64("c-tape-size"),
		c.Context.Int("c-tape-init"),
		c.Context.String("c-cell-type"),
		intermediate,
	)
}
