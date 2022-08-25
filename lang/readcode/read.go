package readcode

import (
	"fmt"
	"os"
	"strings"

	"github.com/baris-inandi/brainfuck/utils"
)

func ToValidBrainfuck(s string) string {
	return strings.Map(
		func(r rune) rune {
			if utils.RuneInSlice(r, []rune{'<', '>', '+', '-', '.', '[', ']', ','}) {
				return r
			}
			return -1
		}, s,
	)
}

func ReadBrainfuck(f string) string {
	/*
		func readBrainfuck
			Gets param f where f is
			a filepath to a brainfuck file,
			reads the brainfuck file,
			removes all non-brainfuck characters,
			returns valid brainfuck code
	*/
	fileBytes, err := os.ReadFile(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return ToValidBrainfuck(string(fileBytes))
}
