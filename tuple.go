package streams

type Pair[T1, T2 any] struct {
	First  T1
	Second T2
}

func NewPair[T1, T2 any](first T1, second T2) Pair[T1, T2] {
	return Pair[T1, T2]{first, second}
}

func (p Pair[T1, T2]) Get() (T1, T2) {
	return p.First, p.Second
}

func (p *Pair[T1, T2]) Set(v1 T1, v2 T2) {
	p.First, p.Second = v1, v2
}

type Tuple3[T1, T2, T3 any] struct {
	V1 T1
	V2 T2
	V3 T3
}

func NewTuple3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3) Tuple3[T1, T2, T3] {
	return Tuple3[T1, T2, T3]{v1, v2, v3}
}

func (t Tuple3[T1, T2, T3]) Get() (T1, T2, T3) {
	return t.V1, t.V2, t.V3
}

func (t *Tuple3[T1, T2, T3]) Set(v1 T1, v2 T2, v3 T3) {
	t.V1, t.V2, t.V3 = v1, v2, v3
}

type Tuple4[T1, T2, T3, T4 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
}

func NewTuple4[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4) Tuple4[T1, T2, T3, T4] {
	return Tuple4[T1, T2, T3, T4]{v1, v2, v3, v4}
}

func (t Tuple4[T1, T2, T3, T4]) Get() (T1, T2, T3, T4) {
	return t.V1, t.V2, t.V3, t.V4
}

func (t *Tuple4[T1, T2, T3, T4]) Set(v1 T1, v2 T2, v3 T3, v4 T4) {
	t.V1, t.V2, t.V3, t.V4 = v1, v2, v3, v4
}

func First[T any](v1 T, _ ...interface{}) T {
	return v1
}

func Second[T1, T2 any](_ T1, v T2, _ ...interface{}) T2 {
	return v
}

func Third[T1, T2, T3 any](_ T1, _ T2, v T3, _ ...interface{}) T3 {
	return v
}

func Forth[T1, T2, T3, T4 any](_ T1, _ T2, _ T3, v T4, _ ...interface{}) T4 {
	return v
}
