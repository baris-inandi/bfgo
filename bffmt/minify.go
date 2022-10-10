package bffmt

import (
	"os"

	"github.com/baris-inandi/brainfuck/lang/readcode"
)

func MinifyFile(files ...string) {
	minify := func(s string) {
		return s
	}

	for _, f := range files {
		minified := minify(readcode.ReadBrainfuck(f))
		err := os.WriteFile(f, []byte(minified), 0644)
		if err != nil {
			panic(err)
		}
	}
}
