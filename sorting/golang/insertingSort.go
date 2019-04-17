package sorting

// Insertion Sort in Golang
//Insertion sort is good only for sorting small arrays (usually less than 100 items).
// In fact, the smaller the array, the faster insertion sort is compared to any other sorting algorithm.
// However, being an O(n2) algorithm, it becomes very slow very quick when the size of the array increases.
// It was used in the tests with arrays of size 100.

import (
	"math/rand"
	"time"
)

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func insertionSort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}



