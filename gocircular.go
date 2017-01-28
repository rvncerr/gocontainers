package gocircular

// CircularBuffer is the basic class in gocircular.
// There are no public members in this struct.
type CircularBuffer struct {
	buffer []interface{}
	shift  int
	size   int
}

// New is the constructor function for CircularBuffer.
func New(size int) CircularBuffer {
	var cb CircularBuffer
	cb.buffer = make([]interface{}, size)
	cb.shift = 0
	cb.size = 0
	return cb
}

// Full checks if CircularBuffer is full.
func (cb *CircularBuffer) Full() bool {
	return cb.size == len(cb.buffer)
}

// Empty checks if CircularBuffer has no elements.
func (cb *CircularBuffer) Empty() bool {
	return cb.size == 0
}

// Size returns number of elements in CircularBuffer.
func (cb *CircularBuffer) Size() int {
	return cb.size
}

// Capacity returns the maximum possible number elements in CircularBuffer.
func (cb *CircularBuffer) Capacity() int {
	return len(cb.buffer)
}

// Front returns the front element in CircularBuffer.
func (cb *CircularBuffer) Front() interface{} {
	if cb.Empty() {
		panic("Calling Front() on an empty CircularBuffer.")
	}
	return cb.buffer[cb.shift]
}

// Back returns the back element in CircularBuffer.
func (cb *CircularBuffer) Back() interface{} {
	if cb.Empty() {
		panic("Calling Back() on an empty CircularBuffer.")
	}
	return cb.buffer[(cb.shift+cb.size-1)%len(cb.buffer)]
}

// PopFront removes front element from CircularBuffer.
func (cb *CircularBuffer) PopFront() {
	if !cb.Empty() {
		cb.size = cb.size - 1
		cb.shift = (cb.shift + 1) % len(cb.buffer)
	}
}

// PopBack removes back element from CircularBuffer.
func (cb *CircularBuffer) PopBack() {
	if !cb.Empty() {
		cb.size = cb.size - 1
	}
}

// PushFront appends new element into CircularBuffer.
// If CircularBuffer is full, PopBack() will be called.
func (cb *CircularBuffer) PushFront(value interface{}) {
	if cb.Full() {
		cb.PopBack()
	}
	cb.buffer[(cb.shift-1)%len(cb.buffer)] = value
	cb.shift = (cb.shift - 1) % len(cb.buffer)
	cb.size = cb.size + 1
}

// PushBack appends new element into CircularBuffer.
// If CircularBuffer is full, PopFront() will be called.
func (cb *CircularBuffer) PushBack(value interface{}) {
	if cb.Full() {
		cb.PopFront()
	}
	cb.buffer[(cb.size+cb.shift)%len(cb.buffer)] = value
	cb.size = cb.size + 1
}

// ToArray converts CircularBuffer to Array.
func (cb *CircularBuffer) ToArray() []interface{} {
	array := make([]interface{}, cb.size)
	for i := 0; i < cb.size; i++ {
		array[i] = cb.buffer[(cb.shift+i)%len(cb.buffer)]
	}
	return array
}
