package math

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func TestNextPowerofTwo(t *testing.T) {
	t.Run("returns 1 for 0", func(t *testing.T) {
		require.Equal(t, uint(1), NextPowerOfTwo(0))
	})

	t.Run("returns same number if it's a power of two", func(t *testing.T) {
		n := uint64(1)
		for i := 0; i < 64; i++ {
			value := n << i
			require.Equal(t, value, NextPowerOfTwo64(value))
		}
	})

	t.Run("always returns a power of two", func(t *testing.T) {
		quick.Check(func(n uint) bool {
			return IsPowerOfTwo(NextPowerOfTwo(n))
		}, nil)
	})
}
