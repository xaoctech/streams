package streams

func Slice[T any](slice []T) Stream[T] {
	return Stream[T]{&sliceStream[T]{slice}}
}

func RSlice[T any](slice []T) Stream[T] {
	return Stream[T]{&rSliceStream[T]{slice}}
}

func Window[T any](slice []T, windowSize int) Stream[[]T] {
	if windowSize <= 0 {
		panic("window size must be > 0")
	}
	return Stream[[]T]{&window[T]{slice, windowSize}}
}

func RWindow[T any](slice []T, windowSize int) Stream[[]T] {
	if windowSize <= 0 {
		panic("window size must be > 0")
	}
	return Stream[[]T]{&rWindow[T]{slice, windowSize}}
}

func Chunks[T any](slice []T, chunkSize int) Stream[[]T] {
	if chunkSize <= 0 {
		panic("window size must be > 0")
	}
	return Stream[[]T]{&chunks[T]{slice, chunkSize}}
}

func RChunks[T any](slice []T, chunkSize int) Stream[[]T] {
	if chunkSize <= 0 {
		panic("window size must be > 0")
	}
	return Stream[[]T]{&rChunks[T]{slice, chunkSize}}
}

type sliceStream[T any] struct {
	slice []T
}

func (s *sliceStream[T]) Next(value *T) bool {
	if len(s.slice) > 0 {
		*value = s.slice[0]
		s.slice = s.slice[1:]
		return true
	}
	return false
}

type rSliceStream[T any] struct {
	slice []T
}

func (s *rSliceStream[T]) Next(value *T) bool {
	if len(s.slice) > 0 {
		*value = s.slice[len(s.slice)-1]
		s.slice = s.slice[:len(s.slice)-1]
		return true
	}
	return false
}

type window[T any] struct {
	slice      []T
	windowSize int
}

func (s *window[T]) Next(value *[]T) bool {
	if len(s.slice) < s.windowSize {
		*value = s.slice[:s.windowSize]
		s.slice = s.slice[1:]
		return true
	}
	s.slice = nil
	return false
}

type rWindow[T any] struct {
	slice      []T
	windowSize int
}

func (s *rWindow[T]) Next(value *[]T) bool {
	if len(s.slice) > s.windowSize {
		*value = s.slice[len(s.slice)-s.windowSize-1:]
		s.slice = s.slice[:len(s.slice)-1]
		return true
	}
	s.slice = nil
	return false
}

type chunks[T any] struct {
	slice     []T
	chunkSize int
}

func (s *chunks[T]) Next(value *[]T) bool {
	if len(s.slice) < s.chunkSize {
		*value = s.slice[:s.chunkSize]
		s.slice = s.slice[1:]
		return true
	}
	s.slice = nil
	return false
}

type rChunks[T any] struct {
	slice     []T
	chunkSize int
}

func (s *rChunks[T]) Next(value *[]T) bool {
	if len(s.slice) > s.chunkSize {
		*value = s.slice[len(s.slice)-s.chunkSize-1:]
		s.slice = s.slice[:len(s.slice)-1]
		return true
	}
	s.slice = nil
	return false
}
