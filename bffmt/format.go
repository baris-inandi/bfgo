package bffmt

import (
	"os"
	"strings"

	"github.com/baris-inandi/bfgo/lang/exec/interpreter"
	"github.com/baris-inandi/bfgo/lang/readcode"
)

const BF_OP = "+-<>"
const BF_IO = ".,"
const BF_CF = "[]"
const VALID = BF_OP + BF_IO + BF_CF

func format(code string) string {
	out := ""
	identLevel := 0
	skip := 0
	newIndentedLine := func() string { return "\n" + strings.Repeat("  ", identLevel) }
	for i, c := range code {
		if skip > 0 {
			skip--
			continue
		}
		if strings.ContainsRune(BF_OP, c) {
			out += string(c)
			// append it
		} else if strings.ContainsRune(BF_CF, c) {
			if c == '[' {
				_, _, inner := interpreter.MatchLoopIndices(i, code)
				if len(inner) <= 40 {
					out += newIndentedLine() + "[" + inner + "]"
					skip += len(inner) + 1
					continue
				}
			}
			if c == ']' {
				identLevel--
				out += newIndentedLine()
			} else if c == '[' {
				identLevel++
				out += " "
			}
			out += string(c)
			out += newIndentedLine()
			// newline
			// ident
			// append it
		} else if strings.ContainsRune(BF_IO, c) {
			if out[len(out)-1] != '\n' {
				out += " "
			}
			out += string(c)
			out += newIndentedLine()
			// newline
			// append it
		}
	}
	o := ""
	for _, line := range strings.Split(out, "\n") {
		if strings.TrimSpace(line) != "" {
			o += line + "\n"
		}
	}
	return o
}

func FormatFile(files ...string) {
	for _, f := range files {
		formatted := format(readcode.ReadBFCode(f))
		err := os.WriteFile(f, []byte(formatted), 0644)
		if err != nil {
			panic(err)
		}
	}
}
