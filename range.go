package streams

import "github.com/xaoctech/streams/constraints"

func Range[N constraints.Number](from, to N) Stream[N] {
	return Stream[N]{&rangeStreamer[N]{from, to, N(1)}}
}

func RangeBy[N constraints.Number](from, to, step N) Stream[N] {
	return Stream[N]{&rangeStreamer[N]{from, to, step}}
}

func RangeFrom[N constraints.Number](from N) Stream[N] {
	return Stream[N]{&rangeFromStreamer[N]{from, N(1)}}
}

func RangeFromBy[N constraints.Number](from, step N) Stream[N] {
	return Stream[N]{&rangeFromStreamer[N]{from, step}}
}

type rangeStreamer[N constraints.Number] struct {
	value, to, step N
}

func (r *rangeStreamer[N]) Next(v *N) bool {
	if r.value < r.to {
		*v = r.value
		r.value += r.step
		return true
	}
	return false
}

type rangeFromStreamer[N constraints.Number] struct {
	value, step N
}

func (r *rangeFromStreamer[N]) Next(v *N) bool {
	*v = r.value
	r.value += r.step
	return true
}
