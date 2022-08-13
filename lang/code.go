package lang

import (
	"fmt"
	"time"

	"github.com/baris-inandi/brainfuck/lang/readcode"
	"github.com/urfave/cli/v2"
)

type Code struct {
	Filepath         string
	Inner            string
	Context          *cli.Context
	OLevel           int
	startTime        time.Time
	VerboseOutBuffer string
}

func (c *Code) VerboseOut(s ...interface{}) {
	if c.Context.Bool("debug") || c.Context.Bool("verbose") {
		elapsedTime := time.Since(c.startTime).Milliseconds()
		out := fmt.Sprint(s...) + fmt.Sprintf(" [%dms]", elapsedTime)
		c.VerboseOutBuffer += out
		if c.Context.Bool("verbose") {
			fmt.Println(out)
		}
	}
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
		Filepath:  filepath,
		Inner:     readcode.ReadBrainfuck(filepath),
		Context:   c,
		OLevel:    oLevel,
		startTime: time.Now(),
	}
}
