package compiler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/baris-inandi/brainfuck/lang"
)

func compileIntermediateIntoFile(c lang.Code, intermediate string, outFile string) {
	if intermediate == "" {
		return
	}

	// generate temp ir file
	f, _ := os.CreateTemp("", "baris-inandi__brainfuck_go_*.c")
	err := ioutil.WriteFile(f.Name(), []byte(intermediate), 0644)
	if err != nil {
		fmt.Print(err)
		fmt.Println("Brainfuck Error: Could not write temporary file.")
	}
	c.VerboseOut("compile.go: generated temporary IR file at ", f.Name())

	if c.Context.Bool("d-print-ir-filepath") {
		fmt.Println(f.Name())
	}

	tempDir := (path.Dir(f.Name()))

	// compile
	ircstdout := &bytes.Buffer{}
	ircstderr := &bytes.Buffer{}
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
	compileCommand := fmt.Sprintf("%s %s %s -o %s %s", compiler, c.Context.String("c-compiler-flags"), optimizeFlag, outFile, f.Name())
	c.VerboseOut("compile.go: generated compile command for IR: ", compileCommand)
	if c.Context.Bool("d-print-compile-command") {
		fmt.Println(compileCommand)
	}
	irccmd := exec.Command("bash", "-c", compileCommand)
	irccmd.Stderr = ircstderr
	irccmd.Stdout = ircstdout
	irccmd.Dir = tempDir
	c.VerboseOut("compile.go: stdout for compile command is: ", ircstdout.String())
	if !c.Context.Bool("compile-only") {
		err = irccmd.Run()
	}
	if err != nil {
		fmt.Println("Brainfuck Compilation Error:\nERROR: ", ircstderr.String())
	}

	if c.OLevel == 3 && !c.Context.Bool("compile-only") {
		stripstdout := &bytes.Buffer{}
		stripstderr := &bytes.Buffer{}
		stripCommand := fmt.Sprintf("strip --strip-unneeded %s", outFile)
		stripcmd := exec.Command("bash", "-c", stripCommand)
		stripcmd.Stderr = stripstderr
		stripcmd.Stdout = stripstdout
		stripcmd.Dir = filepath.Dir(outFile)
		err = stripcmd.Run()
		if err != nil {
			fmt.Println("WARN: Cannot strip binary\n", err)
		}
		c.VerboseOut("compile.go: stripped binary: ", outFile)
	}

	if c.Context.Bool("compile-only") {
		c.VerboseOut("compile.go: using -C, skipping output file")
	}

	if c.Context.Bool("run") {
		c.VerboseOut("compile.go: running binary at ", outFile)
		cmd := exec.Command("bash", "-c", outFile)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Println("WARN: Cannot run binary\n", err)
		}
	}

	// cleanup
	if !c.Context.Bool("d-keep-temp") {
		c.VerboseOut("compile.go: removing temporary IR file at ", f.Name())
		os.Remove(f.Name())
	}
}

func generateOutFile(c lang.Code) string {

	fileIn := c.Filepath
	specifiedName := c.Context.Path("output")

	path, _ := os.Getwd()
	outNoWd := ""
	if specifiedName == "" {
		fileInNameSplit := strings.Split(fileIn, "/")
		fileInName := fileInNameSplit[len(fileInNameSplit)-1]
		fileInNameDotSplit := strings.Split(fileInName, ".")
		outNoWd = fileInNameDotSplit[0]
	} else {
		outNoWd = specifiedName
	}
	return filepath.Join(path, outNoWd)
}

func CompileCodeIntoFile(c lang.Code) {
	/*
		compiles code, a brainfuck string to a binary
		where fileOut is the name of the output file.
		if fileOut is an empty string, the output file
		will be named automatically according to the
		name of the input file.
	*/
	var ir string
	c.VerboseOut("compile.go: optimization level is ", c.OLevel)
	if c.OLevel == 1 {
		c.VerboseOut("compile.go: using fast IR generation")
		ir = FastGenerateIntermediateRepresentation(c)
	} else {
		c.VerboseOut("compile.go: using optimizing IR generation")
		ir = GenerateIntermediateRepresentation(c)
	}

	o := generateOutFile(c)
	c.VerboseOut("compile.go: output file is ", o)

	compileIntermediateIntoFile(
		c,
		ir,
		o, // output binary path
	)
	c.VerboseOut("compile.go: finished compilation")
}
