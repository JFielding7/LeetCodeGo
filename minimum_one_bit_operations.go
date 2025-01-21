package main

import "math/bits"

func minimumOneBitOperations(n int) int {
	target := 0
	operations := 0

	for i := bits.Len(uint(n)) - 1; i > 0; i-- {
		if ((n >> i) & 1) == target {
			target = 0
		} else {
			target = 1
			operations += 1 << i
		}
	}

	return operations + ((n & 1) ^ target)
}
