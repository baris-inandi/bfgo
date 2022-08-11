package lang

import (
	"github.com/baris-inandi/brainfuck/lang/readcode"
	"github.com/urfave/cli/v2"
)

type Code struct {
	Filepath string
	Content  string
	Context  *cli.Context
}

func NewBfCode(c *cli.Context, filepath string) Code {
	return Code{
		Filepath: filepath,
		Content:  readcode.ReadBrainfuck(filepath),
		Context:  c,
	}
}
