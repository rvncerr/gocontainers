package gocontainers

// CircularBuffer is the basic class in gocircular.
// There are no public members in this struct.
type CircularBuffer struct {
	buffer   []interface{}
	capacity int
	shift    int
	size     int
}

// NewCircularBuffer is the constructor function for CircularBuffer.
func NewCircularBuffer(size int) CircularBuffer {
	var cb CircularBuffer
	cb.buffer = make([]interface{}, size)
	cb.capacity = size
	cb.shift = 0
	cb.size = 0
	return cb
}

// At returns element from CircularBuffer by index.
func (cb *CircularBuffer) At(index int) interface{} {
	return cb.buffer[(cb.shift+index)%cb.capacity]
}

// Back returns the back element in CircularBuffer.
// In case of empty CircularBuffer nil returns.
func (cb *CircularBuffer) Back() interface{} {
	if cb.Empty() {
		return nil
	}
	return cb.At(cb.Size() - 1)
}

// Capacity returns the maximum possible number elements in CircularBuffer.
func (cb *CircularBuffer) Capacity() int {
	return cb.capacity
}

// Clear removes all the data from CircularBuffer.
func (cb *CircularBuffer) Clear() {
	cb.size = 0
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
// In case of empty CircularBuffer nil returns.
func (cb *CircularBuffer) Front() interface{} {
	if cb.Empty() {
		return nil
	}
	return cb.At(0)
}

// Full checks if CircularBuffer is full.
func (cb *CircularBuffer) Full() bool {
	return cb.size == cb.capacity
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
		cb.shift = (cb.shift + 1) % cb.capacity
	}
}

// PushBack appends new element into CircularBuffer.
// If CircularBuffer is full, PopFront() will be called.
func (cb *CircularBuffer) PushBack(value interface{}) {
	if cb.Full() {
		cb.PopFront()
	}
	cb.buffer[(cb.size+cb.shift)%cb.capacity] = value
	cb.size = cb.size + 1
}

// PushFront appends new element into CircularBuffer.
// If CircularBuffer is full, PopBack() will be called.
func (cb *CircularBuffer) PushFront(value interface{}) {
	if cb.Full() {
		cb.PopBack()
	}
	index := (cb.shift + cb.capacity - 1) % cb.capacity
	cb.buffer[index] = value
	cb.shift = index
	cb.size = cb.size + 1
}

// Resize affects capacity of CircularBuffer. TODO: Better algorithm.
func (cb *CircularBuffer) Resize(size int) {
	cb.Shift()
	if size > cb.size {
		if len(cb.buffer) < size {
			abuffer := make([]interface{}, size-len(cb.buffer))
			cb.buffer = append(cb.buffer, abuffer...)
		}
	} else {
		cb.size = size
	}
	cb.capacity = size
}

// Shift makes shift zero.
func (cb *CircularBuffer) Shift() {
	var swap = func(i, j int) {
		temp := cb.buffer[i]
		cb.buffer[i] = cb.buffer[j]
		cb.buffer[j] = temp
	}
	var revert = func(i, j int) {
		for k := i; k < (i+j)/2; k++ {
			swap(k, j+i-k-1)
		}
	}
	revert(0, cb.shift)
	revert(cb.shift, cb.capacity)
	revert(0, cb.capacity)
	cb.shift = 0
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

// ToRawArray returns CircularBuffer AS IS.
func (cb *CircularBuffer) ToRawArray() []interface{} {
	return cb.buffer
}
