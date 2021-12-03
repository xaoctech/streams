package streams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStream(t *testing.T) {
	assert.Panics(t, func() { NewStream[int](Streamer[int](nil)) }, "NewStream with nil Streamer should panic")
}
