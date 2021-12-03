package streams

func Collect[T any](stream Stream[T]) []T {
	var v T
	var result []T
	for stream.Next(&v) {
		result = append(result, v)
	}
	return result
}

func (s Stream[T]) Collect() []T {
	return Collect(s)
}

func CollectFirst[T any](stream Stream[T], count int) (first []T, rest Stream[T]) {
	if count < 0 {
		panic("CollectFirst: count < 0")
	}

	var v T
	for len(first) < count && stream.Next(&v) {
		first = append(first, v)
	}
	return first, stream
}

func (s Stream[T]) CollectFirst(count int) (first []T, rest Stream[T]) {
	return CollectFirst(s, count)
}

func CollectLast[T any](stream Stream[T], count, maxSkip int) (skipped int, last []T, hasRest bool) {
	processed := 0
	if count < 0 {
		panic("CollectLast: count < 0")
	}
	if count == 0 {
		var v T
		return 0, nil, stream.Next(&v)
	}
	var buffer []T
	lastIdx := 0
	var v T
	for (maxSkip < 0 || processed < maxSkip+count) && stream.Next(&v) {
		processed += 1
		if len(buffer) < count {
			buffer = append(buffer, v)
		} else {
			buffer[lastIdx] = v
			lastIdx = (lastIdx + 1) % count
		}
	}
	return processed - len(buffer), append(buffer[lastIdx:], buffer[:lastIdx]...), stream.Next(&v)
}

func (s Stream[T]) CollectLast(count, maxSkip int) (skip int, last []T, hasRest bool) {
	return CollectLast(s, count, maxSkip)
}

func CollectFirstLast[T any](stream Stream[T], firstCount, lastCount, maxSkip int) (
	first []T, skipped int, last []T, hasRest bool,
) {
	if firstCount < 0 {
		panic("CollectFirstLast: firstCount < 0")
	}
	if lastCount < 0 {
		panic("CollectFirstLast: lastCount < 0")
	}

	first, stream = CollectFirst(stream, firstCount)
	if len(first) < firstCount {
		if len(first) < lastCount {
			lastCount = len(first)
		}
		if lastCount == 0 {
			var v T
			return nil, 0, nil, stream.Next(&v)
		}
		last = make([]T, lastCount)
		copy(last, first[len(first)-lastCount:])
		return first, -lastCount, last, false
	}
	skipped, last, hasRest = CollectLast(stream, lastCount, maxSkip)
	if len(last) < lastCount {
		if len(last)+len(first) < lastCount {
			lastCount = len(last) + len(first)
		}
		if lastCount == 0 {
			return nil, skipped, nil, hasRest
		}
		copyFirst := lastCount - len(last)
		mergedLast := make([]T, lastCount)

		copy(mergedLast, first[len(first)-copyFirst:])
		copy(mergedLast[copyFirst:], last)
		return first, -copyFirst, mergedLast, hasRest
	}
	return first, skipped, last, hasRest
}

func (s Stream[T]) CollectFirstLast(firstCount, lastCount, maxSkip int) (
	first []T, skipped int, last []T, hasRest bool,
) {
	return CollectFirstLast(s, firstCount, lastCount, maxSkip)
}
