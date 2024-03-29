package compiler

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/baris-inandi/bfgo/lang"
	"github.com/baris-inandi/bfgo/lang/exec/compiler/compiler_utils/compile_command"
	"github.com/baris-inandi/bfgo/lang/exec/compiler/compiler_utils/generate_out_file"
	"github.com/baris-inandi/bfgo/lang/exec/compiler/compiler_utils/strip"
	"github.com/baris-inandi/bfgo/lang/exec/compiler/src/boilerplate/bfhtml"
	"github.com/baris-inandi/bfgo/lang/exec/compiler/src/intermediate"
)

func compileIntermediateIntoFile(c *lang.Code, intermediate string, outFile string) string {
	if intermediate == "" {
		return ""
	}

	// generate temp ir file
	f, _ := os.CreateTemp("", "baris-inandi__bfgo_*."+c.CompileTarget)
	err := os.WriteFile(f.Name(), []byte(intermediate), 0644)
	if err != nil {
		fmt.Print(err)
		fmt.Println("BFGO Error: Could not write temporary file.")
	}
	c.VerboseOut("compile.go: generated temporary IR file at ", f.Name())

	if c.Context.Bool("d-print-ir-filepath") {
		fmt.Println(f.Name())
	}

	// compile
	ircstdout := &bytes.Buffer{}
	ircstderr := &bytes.Buffer{}
	compileCommand := compile_command.GenerateCompileCommand(c, outFile, f)
	irccmd := exec.Command("bash", "-c", compileCommand)
	irccmd.Stderr = ircstderr
	irccmd.Stdout = ircstdout
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	irccmd.Dir = wd
	if !c.Context.Bool("compile-only") {
		err = irccmd.Run()
	}
	if err != nil {
		fmt.Println("BFGO Compilation Error:\nERROR: ", ircstderr.String())
	}

	if c.OLevel == 3 && !c.Context.Bool("compile-only") && !c.Context.Bool("jvm") {
		c.VerboseOut("compile.go: stripping binary at ", outFile)
		strip.Strip(outFile, wd)
		c.VerboseOut("compile.go: stripped binary: ", outFile)
	}

	if c.Context.Bool("compile-only") {
		c.VerboseOut("compile.go: using -C, skipping output file")
	}

	if c.Context.Bool("run") && !c.UsingJS() {
		c.VerboseOut("compile.go: running binary at ", outFile)
		var abs string
		if c.UsingJVM() {
			abs = "java " + c.GetClassName()
		} else {
			abs, err = filepath.Abs(outFile)
			if err != nil {
				fmt.Println(err)
			}
		}
		cmd := exec.Command("bash", "-c", abs)
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

	return outFile
}

func CompileCodeIntoFile(c lang.Code) string {
	/*
		compiles code, a BF string to a binary
		where fileOut is the name of the output file.
		if fileOut is an empty string, the output file
		will be named automatically according to the
		name of the input file.
	*/
	var ir string
	c.VerboseOut("compile.go: optimization level is ", c.OLevel)
	if c.OLevel == 1 {
		c.VerboseOut("compile.go: using fast IR generation")
		ir = intermediate.FastGenerateIntermediateRepresentation(c)
	} else {
		c.VerboseOut("compile.go: using optimizing IR generation")
		ir = intermediate.GenerateIntermediateRepresentation(c)
	}

	var o string
	if c.UsingJVM() {
		o = c.GetClassName() + ".class"
	} else {
		o = generate_out_file.GenerateOutFile(c)
	}
	c.VerboseOut("compile.go: output file is ", o)

	if c.UsingJS() {
		err := os.WriteFile(o, []byte(ir), 0644)
		if err != nil {
			fmt.Println(err)
		}
		c.VerboseOut("compile.go: finished js compilation")
		bfhtml.GenerateHTMLForJSFile(o)
		return o
	}
	compileIntermediateIntoFile(
		&c,
		ir,
		o, // output binary path
	)
	c.VerboseOut("compile.go: finished compilation")
	return o
}
