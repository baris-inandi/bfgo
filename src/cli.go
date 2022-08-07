package src

import "github.com/urfave/cli/v2"

var Cli = &cli.App{
	EnableBashCompletion: true,
	Authors: []*cli.Author{
		{
			Name: "@baris-inandi",
		},
	},
	Action: func(c *cli.Context) error {
		Compile(c.Args().Get(0), c.Args().Get(1))
		return nil
	},
	Usage: "A brainfuck compiler and interpreter written in Go",
	Commands: []*cli.Command{
		{
			Name:    "build",
			Aliases: []string{"compile"},
			Usage:   "Compiles given brainfuck code to binary",
			Action: func(c *cli.Context) error {
				Compile(c.Args().Get(0), c.Args().Get(1))
				return nil
			},
		},
		{
			Name:    "interpret",
			Aliases: []string{"inter"},
			Usage:   "Interprets given brainfuck code",
			Action: func(c *cli.Context) error {
				Interpret(c.Args().Get(0))
				return nil
			},
		},
		{
			Name:  "repl",
			Usage: "Starts brainfuck interactive shell",
			Action: func(c *cli.Context) error {
				Repl()
				return nil
			},
		},
	},
}
