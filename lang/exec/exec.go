package exec

import (
	"fmt"

	"github.com/baris-inandi/brainfuck/lang"
	"github.com/baris-inandi/brainfuck/lang/exec/compiler"
	"github.com/baris-inandi/brainfuck/lang/exec/interpreter"
	"github.com/urfave/cli/v2"
)

func Compile(ctx *cli.Context, filepath string) lang.Code {
	c := lang.NewBfCode(ctx, filepath)
	if c.Context.Bool("jvm") {
		c.VerboseOut("exec.go: running compiler with JVM compile target")
		c.UseJVM()
	} else if c.Context.Bool("js") {
		c.VerboseOut("exec.go: running compiler with JS compile target")
		c.UseJS()
	}
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
