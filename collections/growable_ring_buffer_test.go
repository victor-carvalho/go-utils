package collections

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFifoFixed(t *testing.T) {
	rb := NewGrowableRingBuffer[int](4)
	require.Equal(t, rb.Cap(), 4)

	values := []int{1, 2, 3, 4}

	for _, value := range values {
		rb.PushBack(value)
	}

	require.Equal(t, 4, rb.Len())
	require.Equal(t, 4, rb.Cap())

	for _, expected := range values {
		actual, ok := rb.PopFront()
		require.Equal(t, expected, actual)
		require.True(t, ok)
	}

	require.Equal(t, 0, rb.Len())
	require.Equal(t, 4, rb.Cap())

	_, ok := rb.PopFront()
	require.False(t, ok)
}

func TestFifoGrow(t *testing.T) {
	rb := NewGrowableRingBuffer[int](4)
	require.Equal(t, rb.Cap(), 4)

	values := []int{1, 2, 3, 4, 5}

	for _, value := range values {
		rb.PushBack(value)
	}

	require.Equal(t, 5, rb.Len())
	require.Equal(t, 8, rb.Cap())

	for _, expected := range values {
		actual, ok := rb.PopFront()
		require.Equal(t, expected, actual)
		require.True(t, ok)
	}

	require.Equal(t, 0, rb.Len())
	require.Equal(t, 8, rb.Cap())

	_, ok := rb.PopFront()
	require.False(t, ok)
}
