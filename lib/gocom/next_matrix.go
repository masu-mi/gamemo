package gocom

import "bufio"

func nextMatrix(n, m int, sc *bufio.Scanner) (matrix [][]int) {
	matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			matrix[i][j] = nextInt(sc)
		}
	}
	return
}
