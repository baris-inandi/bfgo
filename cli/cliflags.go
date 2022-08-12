package cli

import c "github.com/urfave/cli/v2"

func Flags() []c.Flag {
	return []c.Flag{
		&c.BoolFlag{Name: "repl", Usage: "Start a read-eval-print loop"},
		&c.BoolFlag{Name: "interpret", Usage: "Interpret file instead of compiling"},
		&c.BoolFlag{Name: "optimize", Aliases: []string{"O"}, Usage: "Enable optimizations"},
		&c.BoolFlag{Name: "run", Aliases: []string{"r"}, Usage: "Immediately run binary after compilation"},
		&c.BoolFlag{Name: "dump-ir", Aliases: []string{"d"}, Usage: "Dump intermediate representation"},
		&c.BoolFlag{Name: "no-strip", Usage: "Do not strip binary, only affects when optimizations are enabled"},
		&c.PathFlag{Name: "o", Usage: "Specify output binary"},
		&c.BoolFlag{Name: "clang", Usage: "Use clang instead of default gcc"},
		&c.BoolFlag{Name: "compile-only", Aliases: []string{"C"}, Usage: "Only compile, do not output a binary"}, // TODO:
		&c.StringFlag{Name: "c-compiler-flags", Usage: "Pass arbitrary flags to the C compiler"},                 // TODO:
		&c.Int64Flag{Name: "c-tape-size", Value: 30000, Usage: "64-bit integer to specify length of brainfuck tape"},
		&c.IntFlag{Name: "c-tape-init", Value: 0, Usage: "Integer value used to initialize all elements in brainfuck tape"},
		&c.StringFlag{Name: "c-cell-type", Value: "int", Usage: "C type used in brainfuck tape in intermediate representation"},
	}
}
