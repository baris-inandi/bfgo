package interpreter

type BfContext struct {
	tape [30000]byte
	ptr  uint16
}

func NewBfContext() BfContext {
	return BfContext{
		tape: [30000]byte{},
		ptr:  0,
	}
}
