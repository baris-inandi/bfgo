package bffmt

import (
	"os"

	"github.com/baris-inandi/bfgo/lang/readcode"
)

func MinifyFile(files ...string) {
	for _, f := range files {
		minified := readcode.ReadBFCode(f)
		err := os.WriteFile(f, []byte(minified), 0644)
		if err != nil {
			panic(err)
		}
	}
}
