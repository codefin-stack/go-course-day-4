package mergesort_test

import (
	"benchmark-test/mergesort"
	"math/rand"
	"testing"
)

//	func TestMergeSort(t *testing.T) {
//		arr := []int{5, 3, 8, 6, 2, 7, 1, 4}
//		eSorted := []int{1, 2, 3, 4, 5, 6, 7, 8}
//		sorted := mergesort.MergeSort(arr)
//		assert.Equal(t, sorted, eSorted)
//		t.Log(sorted)
//	}
func GenerateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(1_000_000)
	}
	return arr
}
func BenchmarkMergeSort(b *testing.B) {
	arr := GenerateRandomArray(1_000_000)
	for n := 0; n < b.N; n++ {
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		mergesort.MergeSort(arrCopy)
	}
}

func BenchmarkConcurrentMergeSort(b *testing.B) {
	arr := GenerateRandomArray(1000000)
	maxDepth := 3
	for n := 0; n < b.N; n++ {
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		mergesort.ConcurrentMergeSort(arrCopy, maxDepth)
	}
}
