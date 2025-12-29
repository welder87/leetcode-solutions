package problem155

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	t.Run("Preset Case 1", func(t *testing.T) {
		obj := Constructor()
		obj.Push(-2)
		checkValue(t, obj.minimums[len(obj.minimums)-1], -2)
		checkValue(t, obj.stack[len(obj.stack)-1], -2)
		obj.Push(0)
		checkValue(t, obj.minimums[len(obj.minimums)-1], -2)
		checkValue(t, obj.stack[len(obj.stack)-1], 0)
		obj.Push(-3)
		checkValue(t, obj.minimums[len(obj.minimums)-1], -3)
		checkValue(t, obj.stack[len(obj.stack)-1], -3)
		checkValue(t, obj.GetMin(), -3)
		obj.Pop()
		checkValue(t, obj.minimums[len(obj.minimums)-1], -2)
		checkValue(t, obj.stack[len(obj.stack)-1], 0)
		checkValue(t, obj.Top(), 0)
		checkValue(t, obj.GetMin(), -2)
	})
	t.Run("Case 1", func(t *testing.T) {
		obj := Constructor()
		obj.Push(-2)
		checkValue(t, obj.minimums[len(obj.minimums)-1], -2)
		checkValue(t, obj.stack[len(obj.stack)-1], -2)
		checkValue(t, obj.GetMin(), -2)
		checkValue(t, obj.Top(), -2)
		obj.Push(0)
		checkValue(t, obj.minimums[len(obj.minimums)-1], -2)
		checkValue(t, obj.stack[len(obj.stack)-1], 0)
		checkValue(t, obj.GetMin(), -2)
		checkValue(t, obj.Top(), 0)
		obj.Push(1)
		checkValue(t, obj.minimums[len(obj.minimums)-1], -2)
		checkValue(t, obj.stack[len(obj.stack)-1], 1)
		checkValue(t, obj.GetMin(), -2)
		checkValue(t, obj.Top(), 1)
		obj.Pop()
		checkValue(t, obj.minimums[len(obj.minimums)-1], -2)
		checkValue(t, obj.stack[len(obj.stack)-1], 0)
		checkValue(t, obj.Top(), 0)
		checkValue(t, obj.GetMin(), -2)
	})
}

func checkValue(t *testing.T, ans, want int) {
	if ans != want {
		t.Errorf("got %v, want %v", ans, want)
	}
}
