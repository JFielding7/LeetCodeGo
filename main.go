package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomArray(n, min, max int) []int {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a slice of length n
	arr := make([]int, n)

	// Fill the slice with random integers in the range [min, max]
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(max-min+1) + min
	}

	return arr
}

func main() {
	//nums := []int{41, 19, 76, 42, 21, 40, 38, 80, 65, 36, 33, 92, 64, 45, 31, 64, 74, 29, 78, 44}
	//nums := []int{10000}
	//fmt.Println(nums)
	//indexDiff := 3
	//valueDiff := 0
	//fmt.Println(containsNearbyAlmostDuplicate(nums, indexDiff, valueDiff))
	matrix := [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}
	target := 0
	fmt.Println(numSubmatrixSumTarget(matrix, target))
}
