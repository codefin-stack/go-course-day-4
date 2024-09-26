package mergesort

import "sync"

func ConcurrentMergeSort(arr []int, maxDepth int) []int {
	if len(arr) <= 1 {
		return arr
	}
	if maxDepth <= 0 {
		return MergeSort(arr)
	}
	mid := len(arr) / 2
	var left, right []int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		left = ConcurrentMergeSort(arr[:mid], maxDepth-1)
	}()
	go func() {
		defer wg.Done()
		right = ConcurrentMergeSort(arr[mid:], maxDepth-1)
	}()
	wg.Wait()
	return merge(left, right)
}
