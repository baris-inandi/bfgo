package lang

import (
	"time"

	"github.com/baris-inandi/bfgo/lang/readcode"
	"github.com/urfave/cli/v2"
)

type Code struct {
	Filepath      string
	Inner         string
	Context       *cli.Context
	OLevel        int
	startTime     time.Time
	DebugFiles    map[string]string
	IsDebugging   bool
	CompileTarget string
	// compileTarget
	// "c" 		-> binary using gcc/clang
	// "java" 	-> class file using javac
}

func NewBfCode(c *cli.Context, filepath string) Code {
	var oLevel int
	if c.Bool("o-performance") {
		oLevel = 3
	} else if c.Bool("o-compile") {
		oLevel = 1
	} else {
		oLevel = 2
	}
	return Code{
		Filepath:      filepath,
		Inner:         readcode.ReadBFCode(filepath),
		Context:       c,
		OLevel:        oLevel,
		startTime:     time.Now(),
		DebugFiles:    map[string]string{},
		IsDebugging:   c.Bool("debug"),
		CompileTarget: "c",
	}
}
