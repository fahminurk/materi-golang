package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
t.Run("3", func(t *testing.T) {
	result := Sum(1, 2)
	require.Equal(t, 3, result, "result must be 3")
})

t.Run("4", func(t *testing.T) {
	result := Sum(2, 2)
	require.Equal(t, 4, result, "result must be 4")
})
}