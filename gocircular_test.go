package gocircular

import (
	"reflect"
	"testing"

	"github.com/rvncerr/goassert"
)

func TestIntegers(t *testing.T) {
	ga := goassert.New(t)
	cb := NewCircularBuffer(4)
	cb.Push(0)
	cb.Push(1)
	ga.Assert(reflect.DeepEqual(cb.ToArray(), []interface{}{0, 1}), "{0, 1}")
	ga.Assert(cb.Head() == interface{}(0), "0 | 1")
	cb.Push(2)
	cb.Push(3)
	ga.Assert(reflect.DeepEqual(cb.ToArray(), []interface{}{0, 1, 2, 3}), "{0, 1, 2, 3}")
	ga.Assert(cb.Head() == interface{}(0), "0 | 1 2 3")
	cb.Push(4)
	cb.Push(5)
	ga.Assert(reflect.DeepEqual(cb.ToArray(), []interface{}{2, 3, 4, 5}), "{2, 3, 4, 5}")
	ga.Assert(cb.Head() == interface{}(2), "2 | 3 4 5")
	cb.Push(6)
	cb.Push(7)
	ga.Assert(reflect.DeepEqual(cb.ToArray(), []interface{}{4, 5, 6, 7}), "{4, 5, 6, 7}")
	ga.Assert(cb.Head() == interface{}(4), "4 | 5 6 7")
	cb.Pop()
	cb.Pop()
	ga.Assert(reflect.DeepEqual(cb.ToArray(), []interface{}{6, 7}), "{6, 7}")
	ga.Assert(cb.Head() == interface{}(6), "6 | 7")
}
