package compile_command

import (
	"fmt"
	"os"

	"github.com/baris-inandi/bfgo/lang"
)

func GenerateCompileCommand(c *lang.Code, outFile string, tempFile *os.File) string {
	var out string
	if c.UsingJVM() {
		out = generateJavaCompileCommand(c, tempFile)
	} else {
		out = generateCCompileCommand(c, outFile, tempFile)
	}
	c.VerboseOut("generate_compile_command.go: generated compile command: ", out)
	if c.Context.Bool("d-print-compile-command") {
		fmt.Println(out)
	}
	return out
}
