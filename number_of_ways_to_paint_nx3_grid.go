package main

const MOD = 1000000007

var x = 0

func count(n int, i int, cache map[int]int) int {
	x += 1
	if n == 0 {
		return 1
	}
	c, exists := cache[2*n+i]
	if exists {
		return c
	}

	res := ((2+i)*count(n-1, 1, cache) + 2*count(n-1, 0, cache)) % MOD
	cache[2*n+i] = res
	return res
}

func numOfWays(n int) int {
	cache := make(map[int]int)
	return 6 * (count(n-1, 1, cache) + count(n-1, 0, cache)) % MOD
}
