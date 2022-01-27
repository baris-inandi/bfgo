package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

var intermediate = ""
var initByteBoilerplateRequired = false
var fmtBoilerplateRequired = false
var byteInBoilerplateRequired = false
var byteOutBoilerplateRequired = false

func readBrainfuck(f string) string {
	fileBytes, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Print(err)
	}
	return toValidBf(string(fileBytes))
}

func runeInSlice(a rune, list []rune) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func toValidBf(s string) string {
	return strings.Map(
		func(r rune) rune {
			if runeInSlice(r, []rune{'<', '>', '+', '-', '.', '[', ']', ','}) {
				return r
			}
			return -1
		}, s,
	)
}

func addCode(x string, init bool) {
	if init {
		initByteBoilerplateRequired = true
		addCode("initByte(ptr);", false)
	}
	intermediate += x
}

func evalExpr(code string, ptr uint) string {
	code = code + "/"
	needsInit := true
	prevChar := ""
	repeatedCharCounter := 1
	initialRepeat := false
	for _, char := range code {
		char := string(char)
		if initialRepeat {
			if prevChar == char {
				repeatedCharCounter += 1
			} else {
				fmt.Println("taşak", prevChar, repeatedCharCounter)
				rep := strconv.Itoa(repeatedCharCounter)
				switch prevChar {
				case "<":
					needsInit = true
					addCode("ptr-="+rep+";", false)
				case ">":
					needsInit = true
					addCode("ptr+="+rep+";", false)
				case "+":
					addCode("tape[ptr]+="+rep+";", needsInit)
					needsInit = false
				case "-":
					addCode("tape[ptr]-="+rep+";", needsInit)
					needsInit = false
				case ".":
					addCode("bo();", true)
					fmtBoilerplateRequired = true
					byteOutBoilerplateRequired = true
				case ",":
					addCode("bi();", false)
					fmtBoilerplateRequired = true
					byteInBoilerplateRequired = true
				case "[":
					addCode("for {", false)
				case "]":
					addCode("if tape[ptr]==byte(0) {break}};", true)
				}
				repeatedCharCounter = 1
			}
		} else {
			initialRepeat = true
		}
		prevChar = char
	}
	return intermediate
}

func generateIntermediateCode(code string, outFile string) {
	boilerplate := "package main;"
	if fmtBoilerplateRequired {
		boilerplate += "import \"fmt\";"
	}
	if initByteBoilerplateRequired {
		boilerplate += "var tape = map[uint64]byte{}; var ptr = uint64(0);"
		boilerplate += "func initByte(x uint64)byte{if val,ok:=tape[x];ok {return val};return byte(0)};"
	}
	if byteInBoilerplateRequired {
		boilerplate += "func bi(){var bfIn byte;fmt.Printf(\"> \");fmt.Scanln(&bfIn);tape[ptr]=bfIn;};"
	}
	if byteOutBoilerplateRequired {
		boilerplate += "func bo(){fmt.Printf(string(tape[ptr]))};"
	}
	goOut := boilerplate + "func main(){" + code + "}"

	// generate temp .go file
	f, _ := os.CreateTemp("", "brainfuck-*.go")
	ioutil.WriteFile(f.Name(), []byte(goOut), 0640)
	tempDir := (path.Dir(f.Name()))

	// compile
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	arg := []string{"build", "-o", outFile, f.Name()}
	cmd := exec.Command("go", arg...)
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	cmd.Dir = tempDir
	err := cmd.Run()
	if err != nil {
		fmt.Println("Brainfuck Compilation Error:")
		fmt.Println("Error: ", stderr.String())
	}

	fmt.Println(f.Name())

	// cleanup
	// os.Remove(f.Name())
}

func generateOutFile(fileIn string, specifiedName string) string {
	path, _ := os.Getwd()
	path += "/"
	outNoWd := ""
	if specifiedName == "" {
		fileInNameSplit := strings.Split(fileIn, "/")
		fileInName := fileInNameSplit[len(fileInNameSplit)-1]
		fileInNameDotSplit := strings.Split(fileInName, ".")
		outNoWd = fileInNameDotSplit[0]
	} else {
		outNoWd = specifiedName
	}
	return path + outNoWd
}

func main() {
	args := os.Args[1:]
	outFile := ""
	if len(args) < 1 {
		fmt.Println("Brainfuck Compiler")
		fmt.Println("Usage: ./brainfuck <file> <outputFile (optional)>")
		os.Exit(0)
	} else if len(args) >= 2 {
		outFile = args[1]
	}
	fileIn := readBrainfuck(args[0])
	out := evalExpr(fileIn, 0)
	generateIntermediateCode(out, generateOutFile(args[0], outFile))
}
