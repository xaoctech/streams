package streams

type Pair[T1, T2 any] struct {
	First  T1
	Second T2
}

func NewPair[T1, T2 any](v1 T1, v2 T2) Pair[T1, T2] {
	return Pair[T1, T2]{
		First:  v1,
		Second: v2,
	}
}

func (p Pair[T1, T2]) Get() (T1, T2) {
	return p.First, p.Second
}

func (p *Pair[T1, T2]) Set(v1 T1, v2 T2) {
	p.First = v1
	p.Second = v2
}

func First[T1, T2 any](v1 T1, _ T2) T1 {
	return v1
}

func Second[T1, T2 any](_ T1, v2 T2) T2 {
	return v2
}
