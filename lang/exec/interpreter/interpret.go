package interpreter

import (
	"github.com/baris-inandi/bfgo/lang"
)

func Interpret(code lang.Code) {
	context := NewBfContext()
	context.EvalExprWithContext(code.Inner)
}
