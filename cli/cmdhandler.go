package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/baris-inandi/brainfuck/bffmt"
	bfexec "github.com/baris-inandi/brainfuck/lang/exec"
	ucli "github.com/urfave/cli/v2"
)

func CmdHandler(c *ucli.Context) error {
	startTime := time.Now()
	if c.Bool("repl") {
		bfexec.Repl(c)
		os.Exit(0)
	}
	if len(c.Args().Slice()) == 0 {
		fmt.Println("No input files")
		fmt.Println("Use brainfuck --help for usage")
		os.Exit(0)
	}
	if c.Bool("minify") {
		bffmt.MinifyFile(c.Args().Slice()...)
		os.Exit(0)
	} else if c.Bool("fmt") {
		bffmt.FormatFile(c.Args().Slice()...)
		os.Exit(0)
	}
	for _, f := range c.Args().Slice() {
		if c.Bool("interpret") {
			c := bfexec.Interpret(c, f)
			c.VerboseOut("cmdhandler.go: interpret job exited in")
		} else {
			c := bfexec.Compile(c, f)
			c.VerboseOut("cmdhandler.go: compile job exited in")
		}
	}
	if c.Bool("time") {
		fmt.Printf("Time: executed in [%dms]\n", time.Since(startTime).Milliseconds())
	}
	return nil
}
