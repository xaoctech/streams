package streams

type Streamer[T any] interface {
	// Next returns true and stores next stream value if stream is not empty
	Next(value *T) bool
}

type Failable interface {
	// Err returns stream error
	Err() error
}

type Derived interface {
	Derived() []interface{}
}

type Reversible[T any] interface {
	Rev() Stream[T]
}

type Stream[T any] struct {
	s Streamer[T]
}

func NewStream[T any](s Streamer[T]) Stream[T] {
	if s == nil {
		panic("nil Streamer")
	}
	return Stream[T]{s}
}

func (s Stream[T]) Next(value *T) bool {
	return s.s.Next(value)
}

func Count[T any](stream Stream[T]) int {
	var count int
	var v T
	for stream.Next(&v) {
		count++
	}
	return count
}

func (s Stream[T]) Count() int {
	return Count(s)
}

func DerivedErr(d Derived) error {
	for _, i := range d.Derived() {
		if err := Err(i); err != nil {
			return err
		}
	}
	return nil
}

func Err(i interface{}) error {
	if failable, ok := i.(Failable); ok {
		return failable.Err()
	}
	if derived, ok := i.(Derived); ok {
		return DerivedErr(derived)
	}
	return nil
}

func (s Stream[T]) Err() error {
	return Err(s.s)
}

type assertStream[T any] struct {
	stream Stream[T]
}

func (a *assertStream[T]) Next(value *T) bool {
	if a.stream.Next(value) {
		return true
	}
	if err := a.stream.Err(); err != nil {
		panic(err)
	}
	return false
}

func Assert[T any](stream Stream[T]) Stream[T] {
	return Stream[T]{&assertStream[T]{stream}}
}

func (s Stream[T]) Assert() Stream[T] {
	return Assert(s)
}

func Fold[T any, B any](s Stream[T], accum B, f func(B, T) B) B {
	var v T
	for s.Next(&v) {
		accum = f(accum, v)
	}
	return accum
}
