package main

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func DivZero[T Number](x, y T) T {
	if y == 0 {
		return 0
	}

	return x / y
}
