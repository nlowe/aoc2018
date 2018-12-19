package util

import "strconv"

func AtoiOrPanic(a string) int {
	if i, err := strconv.Atoi(a); err != nil {
		panic(err)
	} else {
		return i
	}
}
