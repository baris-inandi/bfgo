package cli

import (
	"fmt"
	"os"
	"time"

	bfexec "github.com/baris-inandi/brainfuck/lang/exec"
	ucli "github.com/urfave/cli/v2"
)

func CmdHandler(c *ucli.Context) error {
	startTime := time.Now()
	f := c.Args().Get(0)
	if c.Bool("repl") {
		bfexec.Repl(c)
	} else if c.Bool("interpret") {
		c := bfexec.Interpret(c, f)
		c.VerboseOut("cmdhandler.go: interpret job exited in")
	} else {
		if f == "" {
			fmt.Println("No input files")
			fmt.Println("Use brainfuck --help for usage")
			os.Exit(0)
		}
		c := bfexec.Compile(c, f)
		c.VerboseOut("cmdhandler.go: compile job exited in")
	}
	if c.Bool("time") {
		fmt.Printf("Time: executed in [%dms]\n", time.Since(startTime).Milliseconds())
	}
	return nil
}
