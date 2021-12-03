package streams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func genSlice(start, len int) []int {
	var slice []int
	for i := 0; i < len; i++ {
		slice = append(slice, start+i)
	}
	return slice
}

func TestCollect(t *testing.T) {
	a := assert.New(t)

	a.Equal([]int(nil), Slice([]int(nil)).Collect())
	a.Equal([]int(nil), Slice([]int{}).Collect())

	a.Equal([]int(nil), Range(0, 0).Collect())
	a.Equal([]int{0}, Range(0, 1).Collect())
	a.Equal([]int{0, 1, 2}, Range(0, 3).Collect())
	a.Equal(genSlice(0, 1000), Range(0, 1000).Collect())
}

func TestCollectFirst(t *testing.T) {
	a := assert.New(t)

	a.Equal([]int(nil), First(Range(0, 0).CollectFirst(100)))
	a.Equal([]int(nil), First(Range(0, 100).CollectFirst(0)))
	a.Equal([]int{0, 1, 2}, First(Range(0, 100).CollectFirst(3)))

	{
		first, rest := RangeFrom(0).CollectFirst(100)
		a.Equal(genSlice(0, 100), first)
		a.Equal(genSlice(100, 100), First(rest.CollectFirst(100)))
	}
}

func TestCollectLast(t *testing.T) {
	a := assert.New(t)

	a.Equal(NewPair(0, []int(nil)), NewPair(Range(0, 0).CollectLast(100, -1)))
	a.Equal([]int(nil), Second(Range(0, 100).CollectLast(0, -1)))
	a.Equal([]int{0}, Second(Range(0, 1).CollectLast(3, -1)))
	a.Equal([]int{0, 1}, Second(Range(0, 2).CollectLast(3, -1)))
	a.Equal([]int{0, 1}, Second(Range(0, 2).CollectLast(3, 10)))
	a.Equal([]int{0, 1, 2}, Second(Range(0, 10).CollectLast(3, 0)))
	a.Equal(NewPair(1, []int{1, 2, 3}), NewPair(Range(0, 10).CollectLast(3, 1)))
	a.Equal(NewPair(50, []int{50, 51, 52}), NewPair(Range(0, 100).CollectLast(3, 50)))
	a.Equal(NewPair(97, []int{97, 98, 99}), NewPair(Range(0, 100).CollectLast(3, -1)))
}
