// Brainfuck Compiler written in Go.
// from @baris-inandi <https://github.com/baris-inandi/brainfuck>

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

/*

func o  -> byteOut - a utility function that prints out the byte the pointer is currently on as a string
func i  -> byteIn - a function that gets user input from stdin and writes the input to the tape as bytes
func b  -> initByte - a function that writes 0 to the current byte if not initialized, called whenever init param of addCode() is true
var  p  -> pointer (of type uint64)
var  t  -> tape (of type map[uint64]byte{})

*/

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
		addCode("b(p);", false)
	}
	intermediate += x
}

func evalExpr(code string) string {
	code = code + "/"
	needsInit := true
	prevChar := ""
	repeatedCharCounter := 1
	initialRepeat := false
	for _, char := range code {
		char := string(char)
		if initialRepeat {
			if prevChar == char && (prevChar == "+" || prevChar == "-" || char == "<" || char == ">") {
				repeatedCharCounter += 1
			} else {
				rep := strconv.Itoa(repeatedCharCounter)
				switch prevChar {
				case "<":
					needsInit = true
					addCode("p-="+rep+";", false)
				case ">":
					needsInit = true
					addCode("p+="+rep+";", false)
				case "+":
					addCode("t[p]+="+rep+";", needsInit)
					needsInit = false
				case "-":
					addCode("t[p]-="+rep+";", needsInit)
					needsInit = false
				case ".":
					addCode("o();", true)
					fmtBoilerplateRequired = true
					byteOutBoilerplateRequired = true
				case ",":
					addCode("i();", false)
					fmtBoilerplateRequired = true
					byteInBoilerplateRequired = true
				case "[":
					addCode("for {if t[p]==byte(0) {break};", false)
				case "]":
					addCode("};", true)
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
		boilerplate += "var t = map[uint64]byte{}; var p = uint64(0);"
		boilerplate += "func b(x uint64)byte{if v,ok:=t[x];ok {return v};t[x]=byte(0);return byte(0)};"
	}
	if byteInBoilerplateRequired {
		boilerplate += "func i(){var x byte;fmt.Printf(\"> \");fmt.Scanln(&x);t[p]=x;};"
	}
	if byteOutBoilerplateRequired {
		boilerplate += "func o(){fmt.Printf(string(t[p]))};"
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

	// cleanup
	os.Remove(f.Name())
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
	out := evalExpr(fileIn)
	generateIntermediateCode(out, generateOutFile(args[0], outFile))
}
