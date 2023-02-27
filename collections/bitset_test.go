package collections

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBitset(t *testing.T) {
	set := NewFixedBitset(256)

	input := make([]byte, 20)
	_, err := rand.Read(input)
	require.NoError(t, err)

	for _, x := range input {
		set.SetByte(x)
	}

	for _, x := range input {
		require.True(t, set.IsSetByte(x))
	}

	set.Reset()

	for _, x := range input {
		require.False(t, set.IsSetByte(x))
	}
}
