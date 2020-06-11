package Sorting

import "testing"

func TestCountingSort(t *testing.T) {
	arr := []int{1, 5, 0, 23, 4, 12, 9}
	out := CountingSort(arr)
	t.Logf("CountingSort--> %v", out)
}
