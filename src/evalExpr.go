package src

import "fmt"

func matchLoopIndices(index int, code string) (int, int, string) {
	/*
		func matchLoopIndices
			returns the start and end indices of a loop expression
			where index is the index of the opening bracket
			and code is a string of the whole brainfuck code
	*/
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
	// loop starts at index, ends with indexEnd
	return index, indexEnd, loopExpr
}

func EvalExpr(code string) {
	/*
		func EvalExpr
			interprets brainfuck code where code is a string of brainfuck code
	*/
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
				fmt.Print(string(tape[ptr]))
			case ",":
				var bfIn byte
				fmt.Printf("> ")
				fmt.Scanln(&bfIn)
				tape[ptr] = bfIn
			case "[":
				startIndex, endIndex, loopExpr := matchLoopIndices(index, code)
				skipCount := endIndex - startIndex
				if tape[ptr] != 0 {
					for tape[ptr] > 0 {
						EvalExpr(loopExpr)
					}
				}
				skipChars = skipCount
			}
		}
	}
}
