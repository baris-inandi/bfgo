package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	tapeLen = 20
)

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
			if runeInSlice(r, []rune{'<', '>', '+', '-', '.', '[', ']'}) {
				return r
			}
			return -1
		}, s,
	)
}

func skipCode(str string, index int) string {
	s := []rune(str)
	return string(append(s[0:index], s[index+1:]...))
}

func matchLoopIndices(index int, code string) (int, int, string) {
	indexEnd := index
	depth := 0
	for indexEnd < len(code) {
		if string(code[indexEnd]) == "[" {
			depth++
		}
		if string(code[indexEnd]) == "]" {
			depth--
		}
		if depth == 0 && string(code[indexEnd]) == "]" {
			break
		}
		indexEnd++
	}
	index++
	codeRunes := []rune(code)
	loopExpr := string(codeRunes[index:indexEnd])
	return index, indexEnd, loopExpr
	// loop starts at index, ends, with indexEnd
}

func evalExpr(code string, ptr uint, tape [tapeLen]byte) ([tapeLen]byte, uint) {
	skipChars := 0
	for index, char := range code {
		if skipChars > 0 {
			skipChars--
		} else {
			char := string(char)
			switch char {
			case "<":
				ptr--
			case ">":
				ptr++
			case "+":
				tape[ptr]++
			case "-":
				tape[ptr]--
			case ".":
				fmt.Printf(string(tape[ptr]))
			case "[":
				startIndex, endIndex, loopExpr := matchLoopIndices(index, code)
				skipCount := endIndex - startIndex
				if tape[ptr] != 0 {
					for tape[ptr] > 0 {
						exprTape, exprPtr := evalExpr(loopExpr, ptr, tape)
						ptr, tape = exprPtr, exprTape
					}
				}
				skipChars = skipCount
			}
		}
	}
	return tape, ptr
}

func main() {
	code := readBrainfuck("main.bf")
	out, _ := evalExpr(code, 0, [tapeLen]byte{})
	fmt.Println(out)
}
