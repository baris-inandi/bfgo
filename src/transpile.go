package src

import (
	"fmt"
	"strconv"
	"strings"
)

const IR string = `#include <stdio.h>
int main()
{
    char t[30000] = {0};
    int p = 0;
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
	"a": {
		pattern: "[-]",
		ir:      "t[p]=0;",
	},
	"b": {
		pattern: "[->+<]",
		ir:      "t[p+1]+=t[p];t[p]=0;",
	},
	"c": {
		pattern: "[-<->]",
		ir:      "t[p-1]-=t[p];t[p]=0;",
	},
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
					intermediate += ("t[p]+=" + rep + ";")
				case "-":
					intermediate += ("t[p]-=" + rep + ";")
				case ".":
					intermediate += ("putchar(t[p]);")
				case ",":
					intermediate += ("t[p]=getchar();")
				case "[":
					intermediate += ("while (t[p]){")
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
