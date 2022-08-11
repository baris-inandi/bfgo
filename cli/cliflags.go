package cli

import ucli "github.com/urfave/cli/v2"

func Flags() []ucli.Flag {
	return []ucli.Flag{
		&ucli.BoolFlag{
			Name:  "repl",
			Usage: "Start a read-eval-print loop",
		},
		&ucli.BoolFlag{
			Name:  "interpret",
			Usage: "Interpret file instead of compiling",
		},
		&ucli.BoolFlag{
			Name:    "optimize",
			Aliases: []string{"O"},
			Usage:   "Enable optimizations",
		},
		&ucli.BoolFlag{
			Name:  "run",
			Usage: "Run binary after compilation",
		},
		&ucli.BoolFlag{
			Name:    "dump-ir",
			Aliases: []string{"d"},
			Usage:   "Dump intermediate representation",
		},
		&ucli.BoolFlag{
			Name:  "no-strip",
			Usage: "Do not strip binary",
		},
		&ucli.PathFlag{
			Name:  "o",
			Usage: "Specify output binary",
		},
		&ucli.BoolFlag{
			Name:  "clang",
			Usage: "Use clang instead of default gcc",
		},
	}
}
