package src

import (
	"fmt"
	"os"
)

func Repl() {
	// welcome message
	fmt.Println("Brainfuck REPL")
	fmt.Println("Type 'exit' to exit.")
	for {
		// get prompt
		fmt.Print("brainfuck> ")
		input := ""
		fmt.Scanln(&input)
		// handle exiting
		if input == "exit" || input == "quit" {
			fmt.Println("\nGoodbye!")
			os.Exit(0)
		}
		EvalExpr(toValidBf(input))
		// print an empty line if a print statement is present in the input
		if runeInSlice('.', []rune(input)) {
			fmt.Println()
		}
	}
}
