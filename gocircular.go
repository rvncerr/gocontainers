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

// At returns element from CircularBuffer by index.
func (cb *CircularBuffer) At(index int) interface{} {
	return cb.buffer[(cb.shift+index)%len(cb.buffer)]
}

// Back returns the back element in CircularBuffer.
func (cb *CircularBuffer) Back() interface{} {
	if cb.Empty() {
		panic("Calling Back() on an empty CircularBuffer.")
	}
	return cb.At(cb.Size() - 1)
}

// Capacity returns the maximum possible number elements in CircularBuffer.
func (cb *CircularBuffer) Capacity() int {
	return len(cb.buffer)
}

// Do calls function f on each element of the CircularBuffer.
func (cb *CircularBuffer) Do(f func(interface{})) {
	for i := 0; i < cb.size; i++ {
		f(cb.At(i))
	}
}

// Empty checks if CircularBuffer has no elements.
func (cb *CircularBuffer) Empty() bool {
	return cb.size == 0
}

// Front returns the front element in CircularBuffer.
func (cb *CircularBuffer) Front() interface{} {
	if cb.Empty() {
		panic("Calling Front() on an empty CircularBuffer.")
	}
	return cb.At(0)
}

// Full checks if CircularBuffer is full.
func (cb *CircularBuffer) Full() bool {
	return cb.size == len(cb.buffer)
}

// PopBack removes back element from CircularBuffer.
func (cb *CircularBuffer) PopBack() {
	if !cb.Empty() {
		cb.size = cb.size - 1
	}
}

// PopFront removes front element from CircularBuffer.
func (cb *CircularBuffer) PopFront() {
	if !cb.Empty() {
		cb.size = cb.size - 1
		cb.shift = (cb.shift + 1) % len(cb.buffer)
	}
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

// PushFront appends new element into CircularBuffer.
// If CircularBuffer is full, PopBack() will be called.
func (cb *CircularBuffer) PushFront(value interface{}) {
	if cb.Full() {
		cb.PopBack()
	}
	index := (cb.shift + len(cb.buffer) - 1) % len(cb.buffer)
	cb.buffer[index] = value
	cb.shift = index
	cb.size = cb.size + 1
}

// Size returns number of elements in CircularBuffer.
func (cb *CircularBuffer) Size() int {
	return cb.size
}

// ToArray converts CircularBuffer to Array.
func (cb *CircularBuffer) ToArray() []interface{} {
	array := make([]interface{}, cb.size)
	for i := 0; i < cb.size; i++ {
		array[i] = cb.At(i)
	}
	return array
}
