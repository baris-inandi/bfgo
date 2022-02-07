package src

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func compile(outFile string) {
	// generate temp .go file
	f, _ := os.CreateTemp("", "brainfuck-*.go")
	err := ioutil.WriteFile(f.Name(), []byte(intermediate), 0640)
	if err != nil {
		fmt.Print(err)
		fmt.Println("Brainfuck Error: Could not write temporary file.")
	}
	tempDir := (path.Dir(f.Name()))

	// compile
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	arg := []string{"build", "-o", outFile, f.Name()}
	cmd := exec.Command("go", arg...)
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	cmd.Dir = tempDir
	err = cmd.Run()
	if err != nil {
		fmt.Println("Brainfuck Compilation Error:")
		fmt.Println("Error: ", stderr.String())
	}

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

func compileRun() {
	// compile code to binary in the temp directory
	tempBinaryPath := filepath.Join(os.TempDir(), "/brainfuck-temp")
	compile(tempBinaryPath)

	// run binary
	// TODO: make this work

	// cleanup
	os.Remove(tempBinaryPath)
}

func Compile(filename string, fileOut string, run bool) {
	/*
		func Compile
			compiles contents of filename to a binary file
			where fileOut is the name of the output file.
			if fileOut is an empty string, the output file
			will be named automatically according to the
			name of the input file. If param run is true,
			the binary will be immediately executed
			after compilation.
	*/
	brainfuckCode := readBrainfuck(filename) // get valid brainfuck code from file
	transpile(brainfuckCode)                 // variable `intermediate` will be updated
	if run {
		fmt.Println("compiled")
		compileRun() // will compile to temp directory and run binary
		return
	}
	compile(generateOutFile(filename, fileOut)) // will compile to fileOut
}
