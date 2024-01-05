package readcode

import (
	"fmt"
	"os"
	"strings"

	"github.com/baris-inandi/bfgo/utils"
)

func ToValidBF(s string) string {
	return strings.Map(
		func(r rune) rune {
			if utils.RuneInSlice(r, []rune{'<', '>', '+', '-', '.', '[', ']', ','}) {
				return r
			}
			return -1
		}, s,
	)
}

func ReadBFCode(f string) string {
	/*
		func ReadBFCode
			Gets param f where f is
			a filepath to a BF file,
			reads the BF file,
			removes all non-BF characters,
			returns valid BF code
	*/
	fileBytes, err := os.ReadFile(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return ToValidBF(string(fileBytes))
}
