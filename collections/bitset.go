package collections

import mu "github.com/victor-carvalho/go-utils/math"

const bitSizeExp = 5
const bitMask = 63

type FixedBitset struct {
	buf []uint64
}

func NewFixedBitset(cap uint64) FixedBitset {
	bufSize := mu.NextPowerOfTwo64(uint64(cap)) >> bitSizeExp
	if bufSize == 0 {
		bufSize = 1
	}

	return FixedBitset{buf: make([]uint64, bufSize)}
}

func (b FixedBitset) Set(value uint) {
	i := uint64(value) >> bitSizeExp
	if i > uint64(len(b.buf)) {
		return
	}

	b.buf[i] = b.buf[i] | 1<<(value&bitMask)
}

func (b FixedBitset) SetByte(value byte) {
	i := uint64(value) >> bitSizeExp
	if i > uint64(len(b.buf)) {
		return
	}

	b.buf[i] = b.buf[i] | 1<<(value&bitMask)
}

func (b FixedBitset) SetUint16(value uint16) {
	i := uint64(value) >> bitSizeExp
	if i > uint64(len(b.buf)) {
		return
	}

	b.buf[i] = b.buf[i] | 1<<(value&bitMask)
}

func (b FixedBitset) SetUint32(value uint32) {
	i := uint64(value) >> bitSizeExp
	if i > uint64(len(b.buf)) {
		return
	}

	b.buf[i] = b.buf[i] | 1<<(value&bitMask)
}

func (b FixedBitset) SetUint64(value uint64) {
	i := uint64(value) >> bitSizeExp
	if i > uint64(len(b.buf)) {
		return
	}

	b.buf[i] = b.buf[i] | 1<<(value&bitMask)
}

func (b FixedBitset) IsSet(value uint) bool {
	i := uint64(value) >> bitSizeExp
	if i >= uint64(len(b.buf)) {
		return false
	}

	return b.buf[i]&(1<<(value&bitMask)) > 0
}

func (b FixedBitset) IsSetByte(value byte) bool {
	i := uint64(value) >> bitSizeExp
	if i >= uint64(len(b.buf)) {
		return false
	}

	return b.buf[i]&(1<<(value&bitMask)) > 0
}

func (b FixedBitset) IsSetUint16(value uint16) bool {
	i := uint64(value) >> bitSizeExp
	if i >= uint64(len(b.buf)) {
		return false
	}

	return b.buf[i]&(1<<(value&bitMask)) > 0
}

func (b FixedBitset) IsSetUint32(value uint32) bool {
	i := uint64(value) >> bitSizeExp
	if i >= uint64(len(b.buf)) {
		return false
	}

	return b.buf[i]&(1<<(value&bitMask)) > 0
}

func (b FixedBitset) IsSetUint64(value uint64) bool {
	i := uint64(value) >> bitSizeExp
	if i >= uint64(len(b.buf)) {
		return false
	}

	return b.buf[i]&(1<<(value&bitMask)) > 0
}

func (b FixedBitset) Reset() {
	for i := range b.buf {
		b.buf[i] = 0
	}
}
