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
			Name:  "no-optimize",
			Usage: "Disable optimizations",
		},
		&ucli.BoolFlag{
			Name:  "run",
			Usage: "Run binary after compilation",
		},
		&ucli.BoolFlag{
			Name:  "dump-ir",
			Usage: "Dump c intermediate representation",
		},
		&ucli.BoolFlag{
			Name:  "dump-ll",
			Usage: "Dump llvm intermediate representation",
		},
		&ucli.BoolFlag{
			Name:  "no-strip",
			Usage: "Do not strip binary",
		},
		&ucli.PathFlag{
			Name:  "o",
			Usage: "Write output to specified file",
		},
	}
}
