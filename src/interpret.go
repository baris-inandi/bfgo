package src

func Interpret(filepath string) {
	/*
		func Interpret
			interprets brainfuck code from a file
			where filepath is a filepath to a brainfuck file
	*/
	EvalExpr(readBrainfuck(filepath))
}
