package src

import (
	"fmt"
	"strconv"
	"strings"
)

const IR string = `#include <stdio.h>
int main()
{
    int t[30000] = {0};
    int *p = t;
    // ir %s
    return 0;
}
`

type PatternBindingPair struct {
	pattern string
	ir      string
}

type PatternBinder map[string]PatternBindingPair

var PATTERN_BINDINGS = PatternBinder{
	"1": { // reset byte
		pattern: "[-]",
		ir:      "*p=0;"},
	"2": { // add current to right
		pattern: "[->+<]",
		ir:      "*(p+1)+=*p;*p=0;"},
	"4": { // add current to right (alt notation)
		pattern: "[>+<-]",
		ir:      "*(p+1)+=*p;*p=0;"},
	"3": { // subtract left from current
		pattern: "[-<->]",
		ir:      "*(p-1)-=*p;*p=0;"},
}

func transpile(code string) string {
	// transpiles brainfuck code to intermediate representation and returns a string
	if code == "" {
		return ""
	}
	intermediate := "\n\t"
	code += "/"
	prevChar := ""
	repeatedCharCounter := uint16(1)
	initialRepeat := false
	for k, v := range PATTERN_BINDINGS {
		code = strings.Replace(code, v.pattern, k, -1)
	}
	for _, char := range code {
		char := string(char)
		if initialRepeat {
			if prevChar == char && (prevChar == "+" || prevChar == "-" || char == "<" || char == ">") {
				repeatedCharCounter += 1
			} else {
				rep := strconv.Itoa(int(repeatedCharCounter))
				switch prevChar {
				case "<":
					intermediate += ("p-=" + rep + ";")
				case ">":
					intermediate += ("p+=" + rep + ";")
				case "+":
					intermediate += ("*p+=" + rep + ";")
				case "-":
					intermediate += ("*p-=" + rep + ";")
				case ".":
					intermediate += ("putchar(*p);")
				case ",":
					intermediate += ("*p=getchar();")
				case "[":
					intermediate += ("while (*p){")
				case "]":
					intermediate += ("};")
				default:
					intermediate += PATTERN_BINDINGS[prevChar].ir
				}
				repeatedCharCounter = 1
			}
		} else {
			initialRepeat = true
		}
		prevChar = char
	}
	intermediate = fmt.Sprintf(IR, intermediate)
	return intermediate
}
