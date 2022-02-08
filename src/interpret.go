package src

import (
	"fmt"
	"os"
)

func allocTape() {
	/*
		func allocTape
			call this function before
			using the interpreter. This function
			sets a length for the tape
			in order to avoid initializing a tape
			that is too large when the
			compiler is used where the tape
			won't be needed.
	*/
	tape = [tapeLen]byte{0}
}

func Interpret(filepath string) {
	/*
		func Interpret
			interprets brainfuck code from a file
			where filepath is a filepath to a brainfuck file
	*/
	allocTape()
	EvalExpr(readBrainfuck(filepath))
}
func Repl() {
	// welcome message
	fmt.Println("Brainfuck REPL")
	fmt.Println("Type 'exit' to exit.")
	allocTape()
	for {
		// get prompt
		fmt.Print("brainfuck> ")
		input := ""
		fmt.Scanln(&input)
		// handle exiting
		if input == "exit" {
			fmt.Println("\nGoodbye!")
			os.Exit(0)
		}
		if input == "quit" {
			fmt.Println("Type 'exit' to exit.")
		}
		EvalExpr(toValidBf(input))
		// print an empty line if a print statement is present in the input
		if runeInSlice('.', []rune(input)) {
			fmt.Println()
		}
	}
}
