package streams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	a := assert.New(t)

	a.Equal(0, Slice([]int(nil)).Count())
	a.Equal(0, Slice([]int{}).Count())
	a.Equal(1, Slice([]int{1}).Count())
	a.Equal([]int{1}, Slice([]int{1}).Collect())
	a.Equal([]int{1, 2, 3}, Slice([]int{1, 2, 3}).Collect())
}

func TestRSlice(t *testing.T) {
	a := assert.New(t)

	a.Equal(0, RSlice([]int(nil)).Count())
	a.Equal(0, RSlice([]int{}).Count())
	a.Equal(1, RSlice([]int{1}).Count())
	a.Equal([]int{1}, RSlice([]int{1}).Collect())
	a.Equal([]int{3, 2, 1}, RSlice([]int{1, 2, 3}).Collect())
}

func TestWindow(t *testing.T) {
}

func TestRWindow(t *testing.T) {

}

func TestChunks(t *testing.T) {

}

func TestRChunks(t *testing.T) {

}
