package compile_command

import (
	"fmt"
	"os"

	"github.com/baris-inandi/bfgo/lang"
)

func generateJavaCompileCommand(c *lang.Code, tempFile *os.File) string {
	compileCommand := fmt.Sprintf("javac %s -O -d . %s", c.Context.String("c-compiler-flags"), tempFile.Name())
	return compileCommand
}
