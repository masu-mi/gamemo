package gocom

import (
	"bufio"
	"strconv"
)

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return int(a)
}
