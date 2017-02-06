package gocircular

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleCircularBuffer_At() {
	cb := New(4)
	cb.PushBack(0)
	cb.PushBack(1)
	cb.PushBack(2)
	cb.PushBack(3)
	cb.PushBack(4)
	cb.PushBack(5)
	fmt.Printf("%v_%v_%v_%v\n", cb.At(0), cb.At(1), cb.At(2), cb.At(3))

	// Output: 2_3_4_5
}

func ExampleCircularBuffer_Full() {
	cb := New(4)

	cb.PushBack(0)
	cb.PushBack(1)
	fmt.Printf("%v\n", cb.Full())
	cb.PushBack(2)
	cb.PushBack(3)
	fmt.Printf("%v\n", cb.Full())
	cb.PushBack(4)
	cb.PushBack(5)
	fmt.Printf("%v\n", cb.Full())

	// Output:
	// false
	// true
	// true
}

func ExampleCircularBuffer_Empty() {
	cb := New(4)

	fmt.Printf("%v\n", cb.Empty())
	cb.PushBack(0)
	cb.PushBack(1)
	fmt.Printf("%v\n", cb.Empty())
	cb.PopFront()
	cb.PopFront()
	fmt.Printf("%v\n", cb.Empty())

	// Output:
	// true
	// false
	// true
}

func ExampleCircularBuffer_PushBack() {
	cb := New(4)

	cb.PushBack(0) // [0 _ _ _]
	cb.PushBack(1) // [0 1 _ _]
	cb.PushBack(2) // [0 1 2 _]
	cb.PushBack(3) // [0 1 2 3]
	cb.PushBack(4) // [1 2 3 4]
	cb.PushBack(5) // [2 3 4 5]

	fmt.Printf("%v\n", cb.ToArray())
	// Output: [2 3 4 5]
}

func ExampleCircularBuffer_PushFront() {
	cb := New(4)

	cb.PushFront(0) // [0 _ _ _]
	cb.PushFront(1) // [1 0 _ _]
	cb.PushFront(2) // [2 1 0 _]
	cb.PushFront(3) // [3 2 1 0]
	cb.PushFront(4) // [4 3 2 1]
	cb.PushFront(5) // [5 4 3 2]

	fmt.Printf("%v\n", cb.ToArray())
	// Output: [5 4 3 2]
}

/* func do(element interface{}) {
	fmt.Printf("---> %v <---\n", element)
} */

func ExampleCircularBuffer_Do() {
	var do = func(element interface{}) {
		fmt.Printf("---> %v <---\n", element)
	}

	cb := New(4)

	cb.PushBack(0) // [0 _ _ _]
	cb.PushBack(1) // [0 1 _ _]
	cb.PushBack(2) // [0 1 2 _]
	cb.PushBack(3) // [0 1 2 3]

	cb.Do(do)

	// Output:
	// ---> 0 <---
	// ---> 1 <---
	// ---> 2 <---
	// ---> 3 <---
}

func TestIntegers(t *testing.T) {
	assert := assert.New(t)
	cb := New(4)

	assert.Equal(cb.Capacity(), 4, "{_, _, _, _} // 4")
	cb.PushBack(0)
	cb.PushBack(1)
	assert.Equal(cb.ToArray(), []interface{}{0, 1}, "{0, 1}")
	assert.Equal(cb.Front(), interface{}(0), "{FRONT:0, 1}")
	assert.Equal(cb.Back(), interface{}(1), "{0, BACK:1}")
	assert.Equal(cb.Size(), 2, "len({0, 1}) = 2")

	cb.PushBack(2)
	cb.PushBack(3)
	assert.Equal(cb.ToArray(), []interface{}{0, 1, 2, 3}, "{0, 1, 2, 3}")
	assert.Equal(cb.Front(), interface{}(0), "0 | 1 2 3")
	assert.Equal(cb.Size(), 4, "len({0, 1, 2, 3}) = 4")

	cb.PushBack(4)
	cb.PushBack(5)
	assert.Equal(cb.ToArray(), []interface{}{2, 3, 4, 5}, "{2, 3, 4, 5}")
	assert.Equal(cb.Front(), interface{}(2), "2 | 3 4 5")
	assert.Equal(cb.Size(), 4, "len({2, 3, 4, 5}) = 4")

	cb.PushBack(6)
	cb.PushBack(7)
	assert.Equal(cb.ToArray(), []interface{}{4, 5, 6, 7}, "{4, 5, 6, 7}")
	assert.Equal(cb.Front(), interface{}(4), "4 | 5 6 7")
	assert.Equal(cb.Size(), 4, "len({4, 5, 6, 7}) = 4")

	cb.PopFront()
	cb.PopFront()
	assert.Equal(cb.ToArray(), []interface{}{6, 7}, "{6, 7}")
	assert.Equal(cb.Front(), interface{}(6), "6 | 7")
	assert.Equal(cb.Size(), 2, "len({6, 7}) = 2")

	cb.PushFront(8)
	cb.PushFront(9)
	assert.Equal(cb.ToArray(), []interface{}{9, 8, 6, 7}, "{9, 8, 6, 7}")
	assert.Equal(cb.Front(), interface{}(9), "9 | 8 6 7")
	assert.Equal(cb.Size(), 4, "len({9, 8, 6, 7}) = 4")

	cb.PopBack()
	cb.PopBack()
	assert.Equal(cb.ToArray(), []interface{}{9, 8}, "{9, 8}")
	assert.Equal(cb.Front(), interface{}(9), "9 | 8")
	assert.Equal(cb.Size(), 2, "len({9, 8}) = 2")
}
