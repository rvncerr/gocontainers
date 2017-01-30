package gocircular

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rvncerr/goassert"
)

func ExampleCircularBuffer_PushBack() {
	cb := New(4)

	cb.PushBack(0)
	cb.PushBack(1)
	cb.PushBack(2)
	cb.PushBack(3)
	cb.PushBack(4)
	cb.PushBack(5)

	fmt.Printf("%v\n", cb.ToArray())
	// Output: [2 3 4 5]
}

func ExampleCircularBuffer_PushFront() {
	cb := New(4)

	cb.PushFront(0)
	cb.PushFront(1)
	cb.PushFront(2)
	cb.PushFront(3)
	cb.PushFront(4)
	cb.PushFront(5)

	fmt.Printf("%v\n", cb.ToArray())
	// Output: [5 4 3 2]
}

func TestIntegers(t *testing.T) {
	ga := goassert.New(t)
	cb := New(4)

	ga.Assert(cb.Capacity() == 4, "{_, _, _, _} // 4")
	cb.PushBack(0)
	cb.PushBack(1)
	ga.Assert(reflect.DeepEqual(cb.ToArray(), []interface{}{0, 1}), "{0, 1}")
	ga.Assert(cb.Front() == interface{}(0), "{FRONT:0, 1}")
	ga.Assert(cb.Back() == interface{}(1), "{0, BACK:1}")
	ga.Assert(cb.Size() == 2, "len({0, 1}) = 2")

	cb.PushBack(2)
	cb.PushBack(3)
	ga.Assert(reflect.DeepEqual(cb.ToArray(), []interface{}{0, 1, 2, 3}), "{0, 1, 2, 3}")
	ga.Assert(cb.Front() == interface{}(0), "0 | 1 2 3")
	ga.Assert(cb.Size() == 4, "len({0, 1, 2, 3}) = 4")

	cb.PushBack(4)
	cb.PushBack(5)
	ga.Assert(reflect.DeepEqual(cb.ToArray(), []interface{}{2, 3, 4, 5}), "{2, 3, 4, 5}")
	ga.Assert(cb.Front() == interface{}(2), "2 | 3 4 5")
	ga.Assert(cb.Size() == 4, "len({2, 3, 4, 5}) = 4")

	cb.PushBack(6)
	cb.PushBack(7)
	ga.Assert(reflect.DeepEqual(cb.ToArray(), []interface{}{4, 5, 6, 7}), "{4, 5, 6, 7}")
	ga.Assert(cb.Front() == interface{}(4), "4 | 5 6 7")
	ga.Assert(cb.Size() == 4, "len({4, 5, 6, 7}) = 4")

	cb.PopFront()
	cb.PopFront()
	ga.Assert(reflect.DeepEqual(cb.ToArray(), []interface{}{6, 7}), "{6, 7}")
	ga.Assert(cb.Front() == interface{}(6), "6 | 7")
	ga.Assert(cb.Size() == 2, "len({6, 7}) = 2")

	cb.PushFront(8)
	cb.PushFront(9)
	ga.Assert(reflect.DeepEqual(cb.ToArray(), []interface{}{9, 8, 6, 7}), "{9, 8, 6, 7}")
	ga.Assert(cb.Front() == interface{}(9), "9 | 8 6 7")
	ga.Assert(cb.Size() == 4, "len({9, 8, 6, 7}) = 4")

	cb.PopBack()
	cb.PopBack()
	ga.Assert(reflect.DeepEqual(cb.ToArray(), []interface{}{9, 8}), "{9, 8}")
	ga.Assert(cb.Front() == interface{}(9), "9 | 8")
	ga.Assert(cb.Size() == 2, "len({9, 8}) = 2")
}
