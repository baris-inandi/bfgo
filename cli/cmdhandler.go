package cli

import (
	bfexec "github.com/baris-inandi/brainfuck/lang/exec"
	ucli "github.com/urfave/cli/v2"
)

func CmdHandler(c *ucli.Context) error {
	f := c.Args().Get(0)
	if c.Bool("repl") {
		bfexec.Repl(c)
	} else if c.Bool("interpret") {
		bfexec.Interpret(c, f)
	} else {
		bfexec.Compile(c, f)
	}
	return nil
}
