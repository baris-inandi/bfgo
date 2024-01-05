package compile_command

import (
	"fmt"
	"os"

	"github.com/baris-inandi/bfgo/lang"
)

func generateCCompileCommand(c *lang.Code, outFile string, tempFile *os.File) string {
	optimizeFlag := ""
	if c.OLevel == 1 {
		optimizeFlag = "-O0"
	} else if c.OLevel == 2 {
		optimizeFlag = "-O1"
	} else if c.OLevel == 3 {
		optimizeFlag = "-Ofast"
	}
	compiler := "gcc"
	if c.Context.Bool("clang") {
		compiler = "clang"
	}
	compileCommand := fmt.Sprintf("%s %s %s -o %s %s", compiler, c.Context.String("c-compiler-flags"),
		optimizeFlag,
		outFile,
		tempFile.Name(),
	)
	return compileCommand
}
