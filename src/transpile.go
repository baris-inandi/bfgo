package src

import (
	"fmt"
	"os"
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
	"3": { // add current to right (alt notation)
		pattern: "[>+<-]",
		ir:      "*(p+1)+=*p;*p=0;"},
	"4": { // subtract left from current
		pattern: "[-<->]",
		ir:      "*(p-1)-=*p;*p=0;"},
	"5": { // subtract left from current (alt notation)
		pattern: "[<->-]",
		ir:      "*(p-1)-=*p;*p=0;"},

	"6": { // add current to right (2)
		pattern: "[->>+<<]",
		ir:      "*(p+2)+=*p;*p=0;"},
	"7": { // add current to right (alt notation) (2)
		pattern: "[>>+<<-]",
		ir:      "*(p+2)+=*p;*p=0;"},
	"8": { // subtract left from current (2)
		pattern: "[-<<->>]",
		ir:      "*(p-2)-=*p;*p=0;"},
	"9": { // subtract left from current (alt notation) (2)
		pattern: "[<<->>-]",
		ir:      "*(p-2)-=*p;*p=0;"},

	"0": { // add current to right (3)
		pattern: "[->>>+<<<]",
		ir:      "*(p+3)+=*p;*p=0;"},
	"a": { // add current to right (alt notation) (3)
		pattern: "[>>>+<<<-]",
		ir:      "*(p+3)+=*p;*p=0;"},
	"b": { // subtract left from current (3)
		pattern: "[-<<<->>>]",
		ir:      "*(p-3)-=*p;*p=0;"},
	"c": { // subtract left from current (alt notation) (3)
		pattern: "[<<<->>>-]",
		ir:      "*(p-3)-=*p;*p=0;"},

	"d": { // add current to right (4)
		pattern: "[->>>>+<<<<]",
		ir:      "*(p+4)+=*p;*p=0;"},
	"e": { // add current to right (alt notation) (4)
		pattern: "[>>>>+<<<<-]",
		ir:      "*(p+4)+=*p;*p=0;"},
	"f": { // subtract left from current (4)
		pattern: "[-<<<<->>>>]",
		ir:      "*(p-4)-=*p;*p=0;"},
	"g": { // subtract left from current (alt notation) (4)
		pattern: "[<<<<->>>>-]",
		ir:      "*(p-4)-=*p;*p=0;"},

	"h": { // add current to right (5)
		pattern: "[->>>>>+<<<<<]",
		ir:      "*(p+5)+=*p;*p=0;"},
	"i": { // add current to right (alt notation) (5)
		pattern: "[>>>>>+<<<<<-]",
		ir:      "*(p+5)+=*p;*p=0;"},
	"j": { // subtract left from current (5)
		pattern: "[-<<<<<->>>>>]",
		ir:      "*(p-5)-=*p;*p=0;"},
	"k": { // subtract left from current (alt notation) (5)
		pattern: "[<<<<<->>>>>-]",
		ir:      "*(p-5)-=*p;*p=0;"},

	"l": { // add current to right (6)
		pattern: "[->>>>>>+<<<<<<]",
		ir:      "*(p+6)+=*p;*p=0;"},
	"m": { // add current to right (alt notation) (6)
		pattern: "[>>>>>>+<<<<<<-]",
		ir:      "*(p+6)+=*p;*p=0;"},
	"n": { // subtract left from current (6)
		pattern: "[-<<<<<<->>>>>>]",
		ir:      "*(p-6)-=*p;*p=0;"},
	"o": { // subtract left from current (alt notation) (6)
		pattern: "[<<<<<<->>>>>>-]",
		ir:      "*(p-6)-=*p;*p=0;"},

	"p": { // add current to right (7)
		pattern: "[->>>>>>>+<<<<<<<]",
		ir:      "*(p+7)+=*p;*p=0;"},
	"r": { // add current to right (alt notation) (7)
		pattern: "[>>>>>>>+<<<<<<<-]",
		ir:      "*(p+7)+=*p;*p=0;"},
	"s": { // subtract left from current (7)
		pattern: "[-<<<<<<<->>>>>>>]",
		ir:      "*(p-7)-=*p;*p=0;"},
	"t": { // subtract left from current (alt notation) (7)
		pattern: "[<<<<<<<->>>>>>>-]",
		ir:      "*(p-7)-=*p;*p=0;"},

	"u": { // add current to right (8)
		pattern: "[->>>>>>>>+<<<<<<<<]",
		ir:      "*(p+8)+=*p;*p=0;"},
	"v": { // add current to right (alt notation) (8)
		pattern: "[>>>>>>>>+<<<<<<<<-]",
		ir:      "*(p+8)+=*p;*p=0;"},
	"w": { // subtract left from current (8)
		pattern: "[-<<<<<<<<->>>>>>>>]",
		ir:      "*(p-8)-=*p;*p=0;"},
	"x": { // subtract left from current (alt notation) (8)
		pattern: "[<<<<<<<<->>>>>>>>-]",
		ir:      "*(p-8)-=*p;*p=0;"},
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
	depth := 0
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
					depth++
					intermediate += ("while (*p){")
				case "]":
					depth--
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
	if depth > 0 {
		fmt.Println("Syntax error: Unmatched [")
		os.Exit(1)
	} else if depth < 0 {
		fmt.Println("Syntax error: Unmatched ]")
		os.Exit(1)
	}
	intermediate = fmt.Sprintf(IR, intermediate)
	return intermediate
}
