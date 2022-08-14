package utils

// other utility functions
func RuneInSlice(a rune, list []rune) bool {
	/*
		func runeInSlice
			Gets param a where a is a rune,
			and list is a list of runes,
			checks if a is in list
	*/
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
