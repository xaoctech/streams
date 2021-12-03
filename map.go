package streams

func Map[T, B any](stream Stream[T], f func(T) B) Stream[B] {
	return Stream[B]{&mapTo[T, B]{stream: stream, f: f}}
}

func MapWhileOk[T, B any](stream Stream[T], f func(T) (B, bool)) Stream[B] {
	return Stream[B]{&mapWhileOk[T, B]{stream, f, true}}
}

func MapUntilError[T, B any](stream Stream[T], f func(T) (B, error)) Stream[B] {
	return Stream[B]{&mapUntilError[T, B]{stream, f, nil}}
}

type mapTo[T, B any] struct {
	stream Stream[T]
	f      func(T) B
}

func (m *mapTo[T, B]) Next(value *B) bool {
	var from T
	if !m.stream.Next(&from) {
		return false
	}
	*value = m.f(from)
	return true
}

func (m *mapTo[T, B]) Derived() []interface{} {
	return []interface{}{m.stream}
}

type mapWhileOk[T, B any] struct {
	stream Stream[T]
	f      func(T) (B, bool)
	ok     bool
}

func (m *mapWhileOk[T, B]) Next(out *B) bool {
	var v T
	if !m.ok || !m.stream.Next(&v) {
		return false
	}
	*out, m.ok = m.f(v)
	return m.ok
}

func (m *mapWhileOk[T, B]) Derived() []interface{} {
	return []interface{}{m.stream}
}

type mapUntilError[T, B any] struct {
	s   Stream[T]
	f   func(T) (B, error)
	err error
}

func (m *mapUntilError[T, B]) Next(out *B) bool {
	var v T
	if m.err != nil || !m.s.Next(&v) {
		return false
	}
	*out, m.err = m.f(v)
	return m.err == nil
}

func (m mapUntilError[T, B]) Derived() []interface{} {
	return []interface{}{m.s}
}

func (m mapUntilError[T, B]) Err() error {
	if m.err != nil {
		return m.err
	}
	return DerivedErr(m)
}
