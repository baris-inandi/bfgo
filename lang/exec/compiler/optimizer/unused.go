package optimizer

func removeUnusedLeading(code string) string {
	// remove until last [ ] . or ,
	// leading <>+- operators have no effect on output
	removeLast := 0
	for i := len(code) - 1; i >= 0; i-- {
		char := code[i]
		if char == '[' ||
			char == ']' ||
			char == '.' ||
			char == ',' {
			break
		}
		removeLast += 1
	}
	code = string([]rune(code)[:len(code)-removeLast])
	return code
}
