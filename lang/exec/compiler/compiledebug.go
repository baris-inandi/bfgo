package compiler

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/baris-inandi/brainfuck/lang"
	"github.com/baris-inandi/brainfuck/lang/exec/compiler/compiler_utils/generate_out_file"
)

func CompileCodeIntoFileDebug(c lang.Code) string {
	// generate all resources for the debug mode
	time := time.Now().Format("010220060405")
	outFile := generate_out_file.GenerateOutFile(c)
	debugPathName := fmt.Sprintf("%s/debug-%s-%s", filepath.Dir(outFile), filepath.Base(outFile), time)
	err := os.Mkdir(debugPathName, os.FileMode(0755))
	if err != nil {
		fmt.Println(err)
	}
	c.VerboseOut("compiledebug.go: debug output path is ", debugPathName)
	err = c.Context.Set("output", filepath.Join(debugPathName, "/BIN"))
	if err != nil {
		fmt.Println(err)
	}
	out := CompileCodeIntoFile(c)
	c.WriteDebugFiles(debugPathName)
	return out
}
