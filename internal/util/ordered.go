package util

import (
	"golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](x T, y T) T {
	if x < y {
		return x
	}

	return y
}

func Max[T constraints.Ordered](x T, y T) T {
	if x > y {
		return x
	}

	return y
}
