package math

import (
	"math"
	"math/bits"
)

func IsPowerOfTwo(number uint) bool {
	return bits.OnesCount(number) == 1
}

func IsPowerOfTwo8(number uint8) bool {
	return bits.OnesCount8(number) == 1
}

func IsPowerOfTwo16(number uint16) bool {
	return bits.OnesCount16(number) == 1
}

func IsPowerOfTwo32(number uint32) bool {
	return bits.OnesCount32(number) == 1
}

func IsPowerOfTwo64(number uint64) bool {
	return bits.OnesCount64(number) == 1
}

func NextPowerOfTwo(number uint) uint {
	if number <= 1 {
		return 1
	}

	return (math.MaxUint >> bits.LeadingZeros(number-1)) + 1
}

func NextPowerOfTwo8(number uint8) uint8 {
	if number <= 1 {
		return 1
	}

	return (math.MaxUint8 >> bits.LeadingZeros8(number-1)) + 1
}

func NextPowerOfTwo16(number uint16) uint16 {
	if number <= 1 {
		return 1
	}

	return (math.MaxUint16 >> bits.LeadingZeros16(number-1)) + 1
}

func NextPowerOfTwo32(number uint32) uint32 {
	if number <= 1 {
		return 1
	}

	return (math.MaxUint32 >> bits.LeadingZeros32(number-1)) + 1
}

func NextPowerOfTwo64(number uint64) uint64 {
	if number <= 1 {
		return 1
	}

	return (math.MaxUint64 >> bits.LeadingZeros64(number-1)) + 1
}
