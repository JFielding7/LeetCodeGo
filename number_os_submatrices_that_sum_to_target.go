package main

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	rows := len(matrix)
	cols := len(matrix[0])

	total := 0
	for i := 0; i < rows; i++ {
		sums := make([]int, cols+1)
		for j := i; j < rows; j++ {
			sum := 0
			seen := make(map[int]int)
			seen[0] = 1

			for k := 0; k < cols; k++ {
				sum += matrix[j][k]
				currSum := sums[k+1] + sum
				sums[k+1] = currSum

				if c, exists := seen[currSum-target]; exists {
					total += c
				}

				if c, exists := seen[currSum]; exists {
					seen[currSum] = c + 1
				} else {
					seen[currSum] = 1
				}
			}
		}
	}

	return total
}
