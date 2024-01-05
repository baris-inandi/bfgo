package repl

import (
	"fmt"
	"os"

	"github.com/baris-inandi/bfgo/lang/exec/interpreter"
	"github.com/baris-inandi/bfgo/lang/readcode"
	"github.com/baris-inandi/bfgo/utils"
)

func Repl() {
	context := interpreter.NewBfContext()
	// welcome message
	fmt.Println("bfgo REPL\nType 'exit' to exit.")
	for {
		// get prompt
		fmt.Print("bfgo> ")
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
		context.EvalExprContextually(readcode.ToValidBF(input))
		// print an empty line if a print statement is present in the input
		if utils.RuneInSlice('.', []rune(input)) {
			fmt.Println()
		}
	}
}
