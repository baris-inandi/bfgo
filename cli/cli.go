package cli

import (
	"github.com/urfave/cli/v2"
)

var Cli = &cli.App{
	EnableBashCompletion: true,
	Authors:              []*cli.Author{{Name: "@baris-inandi"}},
	Flags:                Flags(),
	Action:               CmdHandler,
	Usage:                "A blazingly fast, optimizing brainfuck compiler and interpreter",
}
