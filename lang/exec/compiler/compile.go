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
	"github.com/baris-inandi/brainfuck/lang/exec/compiler/ir"
)

func compileIntermediateIntoFile(c lang.Code, intermediate string, outFile string) {
	if intermediate == "" {
		return
	}

	// generate temp ir file
	f, _ := os.CreateTemp("", "baris-inandi__brainfuck_*.c")
	err := ioutil.WriteFile(f.Name(), []byte(intermediate), 0644)
	if err != nil {
		fmt.Print(err)
		fmt.Println("Brainfuck Error: Could not write temporary file.")
	}
	tempDir := (path.Dir(f.Name()))

	// compile
	ircstdout := &bytes.Buffer{}
	ircstderr := &bytes.Buffer{}
	optimizeFlag := ""
	if c.Context.Bool("optimize") {
		optimizeFlag = "-Ofast"
	}
	compiler := "gcc"
	if c.Context.Bool("clang") {
		compiler = "clang"
	}
	compileCommand := fmt.Sprintf("%s %s -o %s %s", compiler, optimizeFlag, outFile, f.Name())
	irccmd := exec.Command("bash", "-c", compileCommand)
	irccmd.Stderr = ircstderr
	irccmd.Stdout = ircstdout
	irccmd.Dir = tempDir
	err = irccmd.Run()
	if err != nil {
		fmt.Println("Brainfuck Compilation Error:\nERROR: ", ircstderr.String())
	}

	if !c.Context.Bool("no-strip") {
		stripstdout := &bytes.Buffer{}
		stripstderr := &bytes.Buffer{}
		stripCommand := fmt.Sprintf("strip --strip-unneeded %s", outFile)
		stripcmd := exec.Command("bash", "-c", stripCommand)
		stripcmd.Stderr = stripstderr
		stripcmd.Stdout = stripstdout
		stripcmd.Dir = filepath.Dir(outFile)
		err = stripcmd.Run()
		if err != nil {
			fmt.Println("WARN: Cannot strip binary", err)
		}
	}

	// cleanup
	os.Remove(f.Name())
}

func generateOutFile(c lang.Code) string {

	fileIn := c.Filepath
	specifiedName := c.Context.String("o")

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
	compileIntermediateIntoFile(
		c,
		ir.Intermediate(c), // intermediate representation
		generateOutFile(c), // output binary
	)
}
