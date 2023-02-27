package concurrent

import (
	"github.com/victor-carvalho/go-utils/collections"
)

func unboundedPipe[T any](in <-chan T, out chan<- T) {
	queue := collections.NewRingBuffer[T](32)
	closed := false

mainLoop:
	for {
		if queue.Len() > 0 {
			if closed {
				for value, ok := queue.PopFront(); ok; value, ok = queue.PopFront() {
					out <- value
				}
				close(out)
				return
			}

			value, _ := queue.PeekFront()
			for {
				select {
				case item, ok := <-in:
					if !ok {
						closed = true
						continue mainLoop
					}
					queue.PushBack(item)
				case out <- value:
					queue.RemoveFront()
					continue mainLoop
				}
			}
		}

		item, ok := <-in
		if !ok {
			close(out)
			return
		}
		queue.PushBack(item)
	}
}

func Unbounded[T any]() (chan<- T, <-chan T) {
	in := make(chan T)
	out := make(chan T)
	go unboundedPipe(in, out)
	return in, out
}
