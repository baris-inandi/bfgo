package exec

import (
	"fmt"
	"os"

	"github.com/baris-inandi/brainfuck/lang"
	"github.com/baris-inandi/brainfuck/lang/exec/compiler"
	"github.com/baris-inandi/brainfuck/lang/exec/interpreter"
	"github.com/urfave/cli/v2"
)

func Compile(ctx *cli.Context, filepath string) {
	if filepath == "" {
		fmt.Println("No input files")
		fmt.Println("Use brainfuck --help for usage")
		os.Exit(1)
	}
	compiler.CompileCodeIntoFile(lang.NewBfCode(ctx, filepath))
}

func Interpret(ctx *cli.Context, filepath string) {
	interpreter.Interpret(lang.NewBfCode(ctx, filepath))
}

func Repl(ctx *cli.Context) { interpreter.Repl() }
