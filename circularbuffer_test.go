package gocontainers

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCircularBufferAt(t *testing.T) {
	cb := NewCircularBuffer(4)

	cb.PushBack(0) // [0 _ _ _]
	cb.PushBack(1) // [0 1 _ _]
	cb.PushBack(2) // [0 1 2 _]
	cb.PushBack(3) // [0 1 2 3]
	cb.PushBack(4) // [1 2 3 4]
	cb.PushBack(5) // [2 3 4 5]

	v, e := cb.At(-1)
	assert.Nil(t, v)
	assert.NotNil(t, e)

	v, e = cb.At(0)
	assert.Equal(t, v, 2)
	assert.Nil(t, e)

	v, e = cb.At(1)
	assert.Equal(t, v, 3)
	assert.Nil(t, e)

	v, e = cb.At(2)
	assert.Equal(t, v, 4)
	assert.Nil(t, e)

	v, e = cb.At(3)
	assert.Equal(t, v, 5)
	assert.Nil(t, e)

	v, e = cb.At(4)
	assert.Nil(t, v)
	assert.NotNil(t, e)
}

func TestCircularBufferBack(t *testing.T) {
	cb := NewCircularBuffer(4)

	v, e := cb.Back()
	assert.Nil(t, v)
	assert.NotNil(t, e)

	cb.PushBack(0) // [0 _ _ _]
	cb.PushBack(1) // [0 1 _ _]
	cb.PushBack(2) // [0 1 2 _]

	v, e = cb.Back()
	assert.Equal(t, v, 2)
	assert.Nil(t, e)

	cb.PushBack(3) // [0 1 2 3]
	cb.PushBack(4) // [1 2 3 4]
	cb.PushBack(5) // [2 3 4 5]

	v, e = cb.Back()
	assert.Equal(t, v, 5)
	assert.Nil(t, e)
}

func TestCircularBufferCapacity(t *testing.T) {
	cb := NewCircularBuffer(4)
	assert.Equal(t, cb.Capacity(), 4)
}

func TestCircularBufferClear(t *testing.T) {
	cb := NewCircularBuffer(4)
	assert.Zero(t, cb.Size())

	cb.PushBack(0) // [0 _ _ _]
	cb.PushBack(1) // [0 1 _ _]
	assert.Equal(t, cb.Size(), 2)

	cb.PushBack(2) // [0 1 2 _]
	cb.PushBack(3) // [0 1 2 3]
	cb.PushBack(4) // [1 2 3 4]
	cb.PushBack(5) // [2 3 4 5]
	assert.Equal(t, cb.Size(), 4)

	cb.Clear()
	assert.Zero(t, cb.Size())
}

func TestCircularBufferDo(t *testing.T) {
	testMap := make(map[int]bool)

	var doAllGood = func(element interface{}) error {
		testMap[element.(int)] = true
		return nil
	}

	var doFailOn4 = func(element interface{}) error {
		if element.(int) == 4 {
			return errors.New("test error")
		}
		return nil
	}

	cb := NewCircularBuffer(4)
	cb.PushBack(0) // [0 _ _ _]
	cb.PushBack(1) // [0 1 _ _]
	cb.PushBack(2) // [0 1 2 _]
	cb.PushBack(3) // [0 1 2 3]
	cb.PushBack(4) // [1 2 3 4]
	cb.PushBack(5) // [2 3 4 5]

	e := cb.Do(doAllGood)
	assert.Nil(t, e)
	assert.Equal(t, len(testMap), 4)

	_, ok := testMap[2]
	assert.True(t, ok)
	_, ok = testMap[3]
	assert.True(t, ok)
	_, ok = testMap[4]
	assert.True(t, ok)
	_, ok = testMap[5]
	assert.True(t, ok)

	e = cb.Do(doFailOn4)
	assert.NotNil(t, e)
}

func TestCircularBufferEmpty(t *testing.T) {
	cb := NewCircularBuffer(4)
	assert.True(t, cb.Empty())

	cb.PushBack(0) // [0 _ _ _]
	cb.PushBack(1) // [0 1 _ _]

	assert.False(t, cb.Empty())

	cb.PopFront() // [1 _ _ _]
	cb.PopFront() // [_ _ _ _]

	assert.True(t, cb.Empty())
}

func TestCircularBufferFront(t *testing.T) {
	cb := NewCircularBuffer(4)

	v, e := cb.Front()
	assert.Nil(t, v)
	assert.NotNil(t, e)

	cb.PushBack(0) // [0 _ _ _]
	cb.PushBack(1) // [0 1 _ _]

	v, e = cb.Front()
	assert.Equal(t, v, 0)
	assert.Nil(t, e)

	cb.PushBack(2) // [0 1 2 _]
	cb.PushBack(3) // [0 1 2 3]
	cb.PushBack(4) // [1 2 3 4]
	cb.PushBack(5) // [2 3 4 5]

	v, e = cb.Front()
	assert.Equal(t, v, 2)
	assert.Nil(t, e)
}

func TestCircularBufferFull(t *testing.T) {
	cb := NewCircularBuffer(4)

	cb.PushBack(0) // [0, _, _, _]
	cb.PushBack(1) // [0, 1, _, _]
	assert.False(t, cb.Full())

	cb.PushBack(2) // [0, 1, 2, _]
	cb.PushBack(3) // [0, 1, 2, 3]
	assert.True(t, cb.Full())

	cb.PushBack(4) // [1, 2, 3, 4]
	cb.PushBack(5) // [2, 3, 4, 5]
	assert.True(t, cb.Full())
}

func TestCircularBufferPopBack(t *testing.T) {
	cb := NewCircularBuffer(4)

	cb.PushBack(0) // [0 _ _ _]
	cb.PushBack(1) // [0 1 _ _]
	cb.PushBack(2) // [0 1 2 _]
	cb.PushBack(3) // [0 1 2 3]
	cb.PushBack(4) // [1 2 3 4]
	cb.PushBack(5) // [2 3 4 5]

	cb.PopBack() // [2 3 4 _]
	cb.PopBack() // [2 3 _ _]

	a := cb.ToArray()
	assert.Equal(t, a, []interface{}{2, 3})
}

func TestCircularBufferPopFront(t *testing.T) {
	cb := NewCircularBuffer(4)

	cb.PushBack(0) // [0 _ _ _]
	cb.PushBack(1) // [0 1 _ _]
	cb.PushBack(2) // [0 1 2 _]
	cb.PushBack(3) // [0 1 2 3]
	cb.PushBack(4) // [1 2 3 4]
	cb.PushBack(5) // [2 3 4 5]

	cb.PopFront() // [3 4 5 _]
	cb.PopFront() // [4 5 _ _]

	a := cb.ToArray()
	assert.Equal(t, a, []interface{}{4, 5})
}

func TestCircularBufferPushBack(t *testing.T) {
	cb := NewCircularBuffer(4)

	cb.PushBack(0) // [0 _ _ _]
	cb.PushBack(1) // [0 1 _ _]
	cb.PushBack(2) // [0 1 2 _]
	cb.PushBack(3) // [0 1 2 3]
	cb.PushBack(4) // [1 2 3 4]
	cb.PushBack(5) // [2 3 4 5]

	assert.Equal(t, cb.ToArray(), []interface{}{2, 3, 4, 5})
}

func TestCircularBufferPushFront(t *testing.T) {
	cb := NewCircularBuffer(4)

	cb.PushFront(0) // [0 _ _ _]
	cb.PushFront(1) // [1 0 _ _]
	cb.PushFront(2) // [2 1 0 _]
	cb.PushFront(3) // [3 2 1 0]
	cb.PushFront(4) // [4 3 2 1]
	cb.PushFront(5) // [5 4 3 2]

	assert.Equal(t, cb.ToArray(), []interface{}{5, 4, 3, 2})
}

func TestCircularBufferResize(t *testing.T) {
	cb := NewCircularBuffer(4)

	cb.PushBack(0) // [0 _ _ _]
	cb.PushBack(1) // [0 1 _ _]
	cb.PushBack(2) // [0 1 2 _]
	cb.PushBack(3) // [0 1 2 3]
	cb.PushBack(4) // [1 2 3 4]
	cb.PushBack(5) // [2 3 4 5]
	assert.Equal(t, cb.ToArray(), []interface{}{2, 3, 4, 5})

	cb.Resize(6)   // [2 3 4 5 _ _]
	cb.PushBack(6) // [2 3 4 5 6 _]
	cb.PushBack(7) // [2 3 4 5 6 7]
	cb.PushBack(8) // [3 4 5 6 7 8]
	cb.PushBack(9) // [4 5 6 7 8 9]
	assert.Equal(t, cb.ToArray(), []interface{}{4, 5, 6, 7, 8, 9})

	cb.Resize(2) // [4 5]
	assert.Equal(t, cb.ToArray(), []interface{}{4, 5})

	cb.PushBack(10)
	assert.Equal(t, cb.ToArray(), []interface{}{5, 10})
}

func TestCircularBufferShift(t *testing.T) {
	cb := NewCircularBuffer(4)

	cb.PushBack(0) // [0 _ _ _]
	cb.PushBack(1) // [0 1 _ _]
	cb.PushBack(2) // [0 1 2 _]
	cb.PushBack(3) // [0 1 2 3]
	cb.PushBack(4) // [1 2 3 4]
	cb.PushBack(5) // [2 3 4 5]
	assert.Equal(t, cb.buffer, []interface{}{4, 5, 2, 3})

	cb.shiftToZero() // [2 3 4 5]
	assert.Equal(t, cb.buffer, []interface{}{2, 3, 4, 5})
}

func TestCircularBufferSize(t *testing.T) {
	cb := NewCircularBuffer(4)

	assert.Zero(t, cb.Size())
	cb.PushBack(0)                // [0 _ _ _]
	cb.PushBack(1)                // [0 1 _ _]
	assert.Equal(t, cb.Size(), 2)
	cb.PushBack(2)                // [0 1 2 _]
	cb.PushBack(3)                // [0 1 2 3]
	assert.Equal(t, cb.Size(), 4)
	cb.PushBack(4)                // [1 2 3 4]
	cb.PushBack(5)                // [2 3 4 5]
	assert.Equal(t, cb.Size(), 4)
}

func TestCircularBufferToArray(t *testing.T) {
	cb := NewCircularBuffer(4)

	cb.PushBack(0) // [0 _ _ _]
	cb.PushBack(1) // [0 1 _ _]
	cb.PushBack(2) // [0 1 2 _]
	cb.PushBack(3) // [0 1 2 3]
	cb.PushBack(4) // [1 2 3 4]
	cb.PushBack(5) // [2 3 4 5]

	a := cb.ToArray()
	assert.Equal(t, a, []interface{}{2, 3, 4, 5})

	a = cb.buffer
	assert.Equal(t, a, []interface{}{4, 5, 2, 3})
}

func BenchmarkCircularBuffer_PushBackUnderfill(b *testing.B) {
	cb := NewCircularBuffer(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cb.PushBack(i)
	}
}

func BenchmarkCircularBuffer_PushFrontUnderfill(b *testing.B) {
	cb := NewCircularBuffer(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cb.PushFront(i)
	}
}

func BenchmarkCircularBuffer_PushBackOverfill(b *testing.B) {
	cb := NewCircularBuffer(4)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cb.PushBack(i)
	}
}

func BenchmarkCircularBuffer_PushFrontOverfill(b *testing.B) {
	cb := NewCircularBuffer(4)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cb.PushFront(i)
	}
}