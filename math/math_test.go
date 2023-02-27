package math

import (
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func TestNextPowerofTwo_PowersOfTwo(t *testing.T) {
	n := uint64(1)
	for i := 0; i < 64; i++ {
		value := n << i
		require.Equal(t, value, NextPowerOfTwo64(value))
	}
}

func TestNextPowerofTwo_General(t *testing.T) {
	quick.Check(func(n uint) bool {
		return IsPowerOfTwo(NextPowerOfTwo(n))
	}, nil)
}
