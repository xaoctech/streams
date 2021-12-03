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

	a.Equal(NewTuple3(0, []int(nil), false), NewTuple3(Range(0, 0).CollectLast(100, -1)))
	a.Equal([]int(nil), Second(Range(0, 100).CollectLast(0, -1)))
	a.Equal([]int{0}, Second(Range(0, 1).CollectLast(3, -1)))
	a.Equal([]int{0, 1}, Second(Range(0, 2).CollectLast(3, -1)))
	a.Equal([]int{0, 1}, Second(Range(0, 2).CollectLast(3, 10)))
	a.Equal([]int{0, 1, 2}, Second(Range(0, 10).CollectLast(3, 0)))
	a.Equal(NewTuple3(1, []int{1, 2, 3}, true), NewTuple3(Range(0, 10).CollectLast(3, 1)))
	a.Equal(NewTuple3(50, []int{50, 51, 52}, true), NewTuple3(Range(0, 100).CollectLast(3, 50)))
	a.Equal(NewTuple3(97, []int{97, 98, 99}, false), NewTuple3(Range(0, 100).CollectLast(3, -1)))
}

func TestCollectFirstLast(t *testing.T) {
	a := assert.New(t)

	a.Equal(NewTuple4([]int(nil), 0, []int(nil), false), NewTuple4(Range(0, 0).CollectFirstLast(3, 3, -1)))
	a.Equal(NewTuple4([]int(nil), 0, []int(nil), true), NewTuple4(Range(0, 100).CollectFirstLast(0, 0, -1)))
	a.Equal(NewTuple4([]int{0, 1, 2}, 0, []int{3, 4, 5}, false), NewTuple4(Range(0, 6).CollectFirstLast(3, 3, -1)))
	a.Equal(NewTuple4([]int{0, 1, 2}, 4, []int{7, 8, 9}, false), NewTuple4(Range(0, 10).CollectFirstLast(3, 3, -1)))
	a.Equal(
		NewTuple4([]int{0, 1, 2}, -3, []int{0, 1, 2, 3, 4, 5}, false),
		NewTuple4(Range(0, 6).CollectFirstLast(3, 6, -1)),
	)
	a.Equal(NewTuple4([]int{0, 1, 2}, -3, []int{0, 1, 2}, false), NewTuple4(Range(0, 3).CollectFirstLast(3, 3, -1)))
	a.Equal(NewTuple4([]int{0, 1}, -2, []int{0, 1}, false), NewTuple4(Range(0, 2).CollectFirstLast(3, 3, -1)))
	a.Equal(NewTuple4([]int{0}, -1, []int{0}, false), NewTuple4(Range(0, 1).CollectFirstLast(3, 3, -1)))

	a.Equal(NewTuple4([]int{0, 1, 2}, 0, []int{3, 4, 5}, true), NewTuple4(RangeFrom(0).CollectFirstLast(3, 3, 0)))
	a.Equal(NewTuple4([]int{0, 1, 2}, 3, []int{6, 7, 8}, true), NewTuple4(RangeFrom(0).CollectFirstLast(3, 3, 3)))
}
