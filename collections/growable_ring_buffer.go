package collections

import (
	"github.com/victor-carvalho/go-utils/generic"
)

type GrowableRingBuffer[T any] struct {
	buffer []T
	start  int
	len    int
}

func NewGrowableRingBuffer[T any](initialSize int) *GrowableRingBuffer[T] {
	return &GrowableRingBuffer[T]{
		buffer: make([]T, initialSize),
		start:  0,
		len:    0,
	}
}

func (rb *GrowableRingBuffer[T]) PushBack(item T) {
	if rb.len == len(rb.buffer) {
		rb.grow()
	}

	idx := rb.getIdx(rb.len)
	rb.buffer[idx] = item
	rb.len += 1
}

func (rb *GrowableRingBuffer[T]) PopFront() (T, bool) {
	if rb.len == 0 {
		return generic.Default[T](), false
	}

	result := rb.buffer[rb.start]
	rb.buffer[rb.start] = generic.Default[T]()
	rb.start = rb.getIdx(1)
	rb.len -= 1

	return result, true
}

func (rb *GrowableRingBuffer[T]) PeekFront() (T, bool) {
	if rb.len == 0 {
		return generic.Default[T](), false
	}

	return rb.buffer[rb.start], true
}

func (rb *GrowableRingBuffer[T]) DropFront() bool {
	if rb.len == 0 {
		return false
	}

	rb.buffer[rb.start] = generic.Default[T]()
	rb.start = rb.getIdx(1)
	rb.len -= 1
	return true
}

func (rb *GrowableRingBuffer[T]) Len() int {
	return rb.len
}

func (rb *GrowableRingBuffer[T]) Cap() int {
	return len(rb.buffer)
}

func (rb *GrowableRingBuffer[T]) getIdx(idx int) int {
	return (rb.start + idx) % len(rb.buffer)
}

func (rb *GrowableRingBuffer[T]) grow() {
	newBuffer := make([]T, len(rb.buffer)*2)
	copy(newBuffer, rb.buffer[rb.start:])
	copy(newBuffer[len(rb.buffer)-rb.start:], rb.buffer[:rb.len-rb.start])
	rb.buffer = newBuffer
}
