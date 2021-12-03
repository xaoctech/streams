package streams

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type failure struct{}

func (f failure) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("failure")
}

func TestScannerStream(t *testing.T) {
	a := assert.New(t)

	var stream Stream[string]
	stream = Scanner(bufio.NewScanner(bytes.NewBuffer([]byte("1\n2\n3\n4"))))
	a.Equal([]string{"1", "2", "3", "4"}, stream.Collect())
	a.Nil(stream.Err())
	stream = Scanner(bufio.NewScanner(bytes.NewBuffer([]byte(""))))
	a.Equal(0, stream.Count())
	a.Nil(stream.Err())
	stream = Scanner(bufio.NewScanner(failure{}))
	a.Equal(0, stream.Count())
	a.Equal(fmt.Errorf("failure"), stream.Err())
}
