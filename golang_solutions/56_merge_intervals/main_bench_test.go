package problem56

import (
	"testing"
)

func BenchmarkMerge(b *testing.B) {
	data := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
		{1, 4},
		{4, 5},
		{100, 340},
		{4, 7},
		{1, 4},
		{1, 7},
		{50, 380},
		{1, 1},
		{2, 3},
		{2, 6},
		{2, 15},
		{100, 100},
		{1, 18},
	}
	data1 := make([][]int, len(data))
	data2 := make([][]int, len(data))
	copy(data1, data)
	copy(data2, data)
	b.ResetTimer()
	b.Run("mergeV0", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			merge(data)
		}
	})
	b.Run("mergeV1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mergeV1(data1)
		}
	})
	b.Run("mergeV2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			mergeV2(data2)
		}
	})
}
