package slice

import (
	constraints2 "constraints"

	"github.com/xaoctech/streams/constraints"
)

func Sum[T constraints.Addable](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func Prod[T constraints.Number](slice []T) T {
	prod := T(1)
	for _, v := range slice {
		prod *= v
	}
	return prod
}

func Min[T constraints2.Ordered](slice ...T) (T, bool) {
	if len(slice) == 0 {
		return T{}, false
	}

	min := slice[0]
	for _, v := range slice[1:] {
		if min > v {
			min = v
		}
	}
	return min, true
}

func Max[T constraints2.Ordered](slice ...T) (T, bool) {
	if len(slice) == 0 {
		return T{}, false
	}

	max := slice[0]
	for _, v := range slice[1:] {
		if max < v {
			max = v
		}
	}
	return max, true
}

func MinMax[T constraints2.Ordered](slice ...T) (T, T, bool) {
	if len(slice) == 0 {
		return T{}, T{}, false
	}

	min := slice[0]
	max := min

	for _, v := range slice[1:] {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	return min, max, true
}
