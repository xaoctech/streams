package streams

func Range[N Number](from, to N) Stream[N] {
	return Stream[N]{&rangeStreamer[N]{from, to, N(1)}}
}

func RangeBy[N Number](from, to, step N) Stream[N] {
	return Stream[N]{&rangeStreamer[N]{from, to, step}}
}

func RangeFrom[N Number](from N) Stream[N] {
	return Stream[N]{&rangeFromStreamer[N]{from, N(1)}}
}

func RangeFromBy[N Number](from, step N) Stream[N] {
	return Stream[N]{&rangeFromStreamer[N]{from, step}}
}

type rangeStreamer[N Number] struct {
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

type rangeFromStreamer[N Number] struct {
	value, step N
}

func (r *rangeFromStreamer[N]) Next(v *N) bool {
	*v = r.value
	r.value += r.step
	return true
}
