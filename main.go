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
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.print()
	lfu.Put(2, 2)
	lfu.print()
	lfu.Get(1)
	lfu.Put(3, 3)
	lfu.print()
	fmt.Println(lfu.Get(2))
	lfu.Get(3)
	lfu.print()
	lfu.Put(4, 4)
	lfu.print()
	fmt.Println("Count", lfu.count)
	fmt.Println(lfu.Get(1))
	lfu.Get(3)
	lfu.Get(4)
}
