package streams

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatter(t *testing.T) {
	a := assert.New(t)

	a.Equal("", fmt.Sprintf("%d", Range(0, 0).Format()))
	a.Equal("", fmt.Sprintf("%d", Range(0, 0).Format(FirstLen(0), LastLen(0))))
	a.Equal("", fmt.Sprintf("%d", RangeFrom(0).Format(FirstLen(0), LastLen(0))))
	a.Equal("0", fmt.Sprintf("%d", Range(0, 10).Format(FirstLen(1), LastLen(0))))
	a.Equal("9", fmt.Sprintf("%d", Range(0, 10).Format(FirstLen(0), LastLen(1))))
	a.Equal("0, 1", fmt.Sprintf("%d", Range(0, 10).Format(FirstLen(2), LastLen(0))))
	a.Equal("8, 9", fmt.Sprintf("%d", Range(0, 10).Format(FirstLen(0), LastLen(2))))
	a.Equal("0, 1, 2, 3, 4", fmt.Sprintf("%d", Range(0, 5).Format(FirstLen(10), LastLen(10))))
	a.Equal("0, ..., 10", fmt.Sprintf("%d", Range(0, 11).Format(FirstLen(1), LastLen(1), MaxSkip(10))))
	a.Equal("0, ..., 11", fmt.Sprintf("%d", Range(0, 12).Format(FirstLen(1), LastLen(1), MaxSkip(10))))
	a.Equal("0, ..., 11, ...", fmt.Sprintf("%d", Range(0, 13).Format(FirstLen(1), LastLen(1), MaxSkip(10))))
}
