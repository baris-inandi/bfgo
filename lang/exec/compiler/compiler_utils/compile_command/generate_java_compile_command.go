package compile_command

import (
	"fmt"
	"os"

	"github.com/baris-inandi/brainfuck/lang"
)

func generateJavaCompileCommand(c *lang.Code, tempFile *os.File) string {
	compileCommand := fmt.Sprintf("javac %s -d . %s", c.Context.String("c-compiler-flags"), tempFile.Name())
	return compileCommand
}
