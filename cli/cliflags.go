package cli

import c "github.com/urfave/cli/v2"

func Flags() []c.Flag {
	return []c.Flag{

		// compiler-specific options
		&c.BoolFlag{Name: "run", Aliases: []string{"r"}, Usage: "Immediately run binary after compilation"},
		&c.PathFlag{Name: "output", Aliases: []string{"o"}, Usage: "Specify output binary"},
		&c.BoolFlag{Name: "compile-only", Aliases: []string{"C"}, Usage: "Only compile, do not output a binary"},

		// optimizations
		&c.BoolFlag{Name: "o-compile", Aliases: []string{"F"}, Usage: "Disable optimizations and use fast compiler: fast compile time, slow execution"},
		&c.BoolFlag{Name: "o-balanced", Aliases: []string{"B"}, Usage: "Minimal optimizations for balanced compile time and performance, default behavior"},
		&c.BoolFlag{Name: "o-performance", Aliases: []string{"O"}, Usage: "Enable optimizations: fast execution, slow compile time"},

		// alternate interpreted modes
		&c.BoolFlag{Name: "interpret", Usage: "Interpret file instead of compiling"},
		&c.BoolFlag{Name: "repl", Usage: "Start a read-eval-print loop"},

		// C IR related options
		&c.BoolFlag{Name: "clang", Usage: "Use clang instead of default gcc"},
		&c.StringFlag{Name: "c-compiler-flags", Usage: "Pass arbitrary flags to the C compiler"}, // TODO:
		&c.Int64Flag{Name: "c-tape-size", Value: 30000, Usage: "64-bit integer to specify length of brainfuck tape"},
		&c.IntFlag{Name: "c-tape-init", Value: 0, Usage: "Integer value used to initialize all elements in brainfuck tape"},
		&c.StringFlag{Name: "c-cell-type", Value: "int", Usage: "C type used in brainfuck tape in intermediate representation"},

		// debug options
		&c.BoolFlag{Name: "d-dump-ir", Usage: "Dump intermediate representation"},
		&c.BoolFlag{Name: "d-keep-temp", Usage: "Do not remove temporary IR files"},
		&c.BoolFlag{Name: "d-print-ir-filepath", Usage: "Dump temporary IR filepath, use -d-keep-temp to keep them from being deleted"},
		&c.BoolFlag{Name: "d-print-compile-command", Usage: "Print C IR compiler command"},
	}
}
