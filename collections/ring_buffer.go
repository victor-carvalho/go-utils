package collections

import (
	"github.com/victor-carvalho/go-utils/generic"
)

type RingBuffer[T any] struct {
	buffer []T
	start  int
	len    int
}

func NewRingBuffer[T any](initialSize int) *RingBuffer[T] {
	return &RingBuffer[T]{
		buffer: make([]T, initialSize),
		start:  0,
		len:    0,
	}
}

func (rb *RingBuffer[T]) PushBack(item T) {
	if rb.len == len(rb.buffer) {
		rb.grow()
	}

	idx := rb.wrapAdd(rb.len)
	rb.buffer[idx] = item
	rb.len += 1
}

func (rb *RingBuffer[T]) PushFront(item T) {
	if rb.len == len(rb.buffer) {
		rb.grow()
	}

	rb.start = rb.wrapSub(1)
	rb.buffer[rb.start] = item
	rb.len += 1
}

func (rb *RingBuffer[T]) PopFront() (T, bool) {
	if rb.len == 0 {
		return generic.Default[T](), false
	}

	result := rb.buffer[rb.start]
	rb.buffer[rb.start] = generic.Default[T]()
	rb.start = rb.wrapAdd(1)
	rb.len -= 1

	return result, true
}

func (rb *RingBuffer[T]) PopBack() (T, bool) {
	if rb.len == 0 {
		return generic.Default[T](), false
	}

	rb.len -= 1
	idx := rb.wrapAdd(rb.len)
	result := rb.buffer[idx]
	rb.buffer[idx] = generic.Default[T]()

	return result, true
}

func (rb *RingBuffer[T]) PeekFront() (T, bool) {
	if rb.len == 0 {
		return generic.Default[T](), false
	}

	return rb.buffer[rb.start], true
}

func (rb *RingBuffer[T]) PeekBack() (T, bool) {
	if rb.len == 0 {
		return generic.Default[T](), false
	}

	return rb.buffer[rb.wrapAdd(rb.len-1)], true
}

func (rb *RingBuffer[T]) RemoveFront() bool {
	if rb.len == 0 {
		return false
	}

	rb.buffer[rb.start] = generic.Default[T]()
	rb.start = rb.wrapAdd(1)
	rb.len -= 1
	return true
}

func (rb *RingBuffer[T]) RemoveBack() bool {
	if rb.len == 0 {
		return false
	}

	rb.len -= 1
	rb.buffer[rb.wrapAdd(rb.len)] = generic.Default[T]()

	return true
}

func (rb *RingBuffer[T]) Len() int {
	return rb.len
}

func (rb *RingBuffer[T]) Cap() int {
	return len(rb.buffer)
}

func (rb *RingBuffer[T]) wrapAdd(n int) int {
	return (rb.start + n) % len(rb.buffer)
}

func (rb *RingBuffer[T]) wrapSub(n int) int {
	return (rb.start + len(rb.buffer) - n) % len(rb.buffer)
}

func (rb *RingBuffer[T]) grow() {
	newBuffer := make([]T, len(rb.buffer)*2)
	copy(newBuffer, rb.buffer[rb.start:])
	copy(newBuffer[len(rb.buffer)-rb.start:], rb.buffer[:rb.len-rb.start])
	rb.buffer = newBuffer
}
