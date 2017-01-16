package gocircular

// CircularBuffer is the basic class in gocircular.
type CircularBuffer struct {
	buffer []interface{}
	shift  int
	size   int
}

// NewCircularBuffer is the constructor function for CircularBuffer.
func NewCircularBuffer(size int) CircularBuffer {
	var cb CircularBuffer
	cb.buffer = make([]interface{}, size)
	cb.shift = 0
	cb.size = 0
	return cb
}

// Full returns true if CircularBuffer is full.
func (cb *CircularBuffer) Full() bool {
	return cb.size == len(cb.buffer)
}

// Empty checks if CircularBuffer has no elements.
func (cb *CircularBuffer) Empty() bool {
	return cb.size == 0
}

// Pop removes first element from CircularBuffer.
func (cb *CircularBuffer) Pop() {
	if !cb.Empty() {
		cb.size = cb.size - 1
		cb.shift = (cb.shift + 1) % len(cb.buffer)
	}
}

// Push appends new element into CircularBuffer.
// If CircularBuffer is full, Pop() will be called.
func (cb *CircularBuffer) Push(value interface{}) {
	if cb.Full() {
		cb.Pop()
	}
	cb.buffer[(cb.size+cb.shift)%len(cb.buffer)] = value
	cb.size = cb.size + 1
}

// Head returns the first element of CircularBuffer.
func (cb *CircularBuffer) Head() interface{} {
	return cb.buffer[cb.shift]
}

// ToArray converts CircularBuffer to Array.
func (cb *CircularBuffer) ToArray() []interface{} {
	array := make([]interface{}, cb.size)
	for i := 0; i < cb.size; i++ {
		array[i] = cb.buffer[(cb.shift+i)%len(cb.buffer)]
	}
	return array
}
