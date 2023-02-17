// other utility functions
package utils

/*
Gets param a where a is a rune,
and list is a list of runes,
checks if a is in list
*/
func RuneInSlice(a rune, list []rune) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
