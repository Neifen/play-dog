package util

import (
	"math/rand/v2"
)

func RandItem[S ~[]E, E any](s S) (E, int) {
	length := len(s)

	if length == 0 {
		var e E
		return e, -1 // nil value
	}
	if length == 1 {
		return s[0], 0
	}

	randN := rand.IntN(length - 1)
	return s[randN], randN
}

func RandItemSkipFirst[S ~[]E, E any](s S) (E, int) {
	var e E
	length := len(s)

	if length == 0 || length == 1 {
		return e, -1 // nil value
	}
	if length == 2 {
		return s[1], 1
	}

	randN := rand.IntN(length - 2)
	return s[randN+1], randN + 1
}
