package exec

import (
	"fmt"
	"os"

	"github.com/baris-inandi/brainfuck/lang"
	"github.com/baris-inandi/brainfuck/lang/exec/compiler"
	"github.com/baris-inandi/brainfuck/lang/exec/interpreter"
	"github.com/urfave/cli/v2"
)

func Compile(ctx *cli.Context, filepath string) lang.Code {
	c := lang.NewBfCode(ctx, filepath)
	if filepath == "" {
		fmt.Println("No input files")
		fmt.Println("Use brainfuck --help for usage")
		os.Exit(0)
	}
	c.VerboseOut("exec.go: using run mode compile")
	if ctx.Bool("debug") {
		c.VerboseOut("exec.go: running compiler in debug mode")
		compiler.CompileCodeIntoFileDebug(c)
		return c
	}
	compiler.CompileCodeIntoFile(c)
	return c
}

func Interpret(ctx *cli.Context, filepath string) lang.Code {
	c := lang.NewBfCode(ctx, filepath)
	c.VerboseOut("exec.go: using run mode interpret")
	interpreter.Interpret(c)
	return c
}

func Repl(ctx *cli.Context) {
	if ctx.Bool("verbose") {
		fmt.Println("exec.go: using run mode REPL")
	}
	interpreter.Repl()
}
