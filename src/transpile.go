package src

import (
	"fmt"
	"strconv"
	"strings"
)

const IR string = `
#include <stdio.h>

int main()
{
    unsigned int t[30000] = {0}; // tape
    unsigned int p = 0;         // pointer
    // brainfuck ir %s
    printf("\n");
    return 0;
}
`

type bindingPair struct {
	pattern string
	ir      string
}

var PATTERN_BINDINGS = map[string]bindingPair{
	"a": {
		pattern: "[-]",
		ir:      "t[p]=0;",
	},
	"b": {
		pattern: "[->+<]",
		ir:      "t[p+1]+=t[p];t[p]=0;",
	},
}

func transpile(code string) string {
	// transpiles brainfuck code to intermediate representation and returns a string
	intermediate := "\n\t"
	code += "/"
	prevChar := ""
	repeatedCharCounter := 1
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
				rep := strconv.Itoa(repeatedCharCounter)
				switch prevChar {
				case "<":
					intermediate += ("p-=" + rep + ";")
				case ">":
					intermediate += ("p+=" + rep + ";")
				case "+":
					intermediate += ("t[p]+=" + rep + ";")
				case "-":
					intermediate += ("t[p]-=" + rep + ";")
				case ".":
					intermediate += ("printf(\"%c\",(unsigned char)t[p]);")
				case ",":
					intermediate += ("b.i();")
				case "[":
					intermediate += ("while (t[p] != 0) {")
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
