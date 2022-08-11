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

func compileIntermediateIntoFile(intermediate string, outFile string) {
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
	irccmd := exec.Command("clang", "-Ofast", "-o", outFile, f.Name())
	irccmd.Stderr = ircstderr
	irccmd.Stdout = ircstdout
	irccmd.Dir = tempDir
	err = irccmd.Run()
	if err != nil {
		fmt.Println("Brainfuck Compilation Error:\nERROR: ", ircstderr.String())
	}

	/* 	// strip binary (TODO)
	   	stripstdout := &bytes.Buffer{}
	   	stripstderr := &bytes.Buffer{}
	   	stripcmd := exec.Command("strip", "--strip-unneeded", outFile)
	   	stripcmd.Stderr = stripstderr
	   	stripcmd.Stdout = stripstdout
	   	stripcmd.Dir = filepath.Base(outFile)
	   	err = stripcmd.Run()
	   	if err != nil {
	   		fmt.Println("WARN: Cannot strip binary", err)
	   	} */

	// cleanup
	os.Remove(f.Name())
}

func generateOutFile(fileIn string, specifiedName string) string {
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
		intermediate(c.Content), // intermediate representation
		generateOutFile( // output file
			c.Filepath,            // original filepath
			c.Context.String("o"), // specified filename
		),
	)
}
