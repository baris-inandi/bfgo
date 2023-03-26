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

func RelativeIndex(i, size int) int {
	if i < 0 {
		i += size
	}
	if i < 0 {
		// should it return an err instead of panicking?
		panic("Index out of bounds")
	}
	return i
}

// emulation of "pipeline operator"
func Apply[T any](s T, fs ...func(T) T) T {
	for _, f := range fs {
		s = f(s)
	}
	return s
}
