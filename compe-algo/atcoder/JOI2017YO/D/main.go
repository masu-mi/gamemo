package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

const (
	initialBufSize = 100000
	maxBufSize     = 1000000
)

var sc *bufio.Scanner

func initScanner(r io.Reader) *bufio.Scanner {
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	return sc
}

func main() {
	sc = initScanner(os.Stdin)
	fmt.Println(resolve(parseProblem()))
}

func parseProblem() (int, int, []int) {
	n, m := scanInt(sc), scanInt(sc)
	as := nextIntSlice(sc, n)
	return n, m, as
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return int(a)
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func nextIntSlice(sc *bufio.Scanner, n int) (a []int) {

	a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt(sc)
	}
	return a
}

func resolve(n, m int, as []int) int {
	accum := make([][]int, m)
	for i := 0; i < m; i++ {
		accum[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < m; j++ {
			accum[j][i] = accum[j][i-1]
		}
		cat := as[i-1] - 1
		accum[cat][i]++
	}
	// How can I debug in dp code
	dp := make([]pair, 1<<uint(m))
	for i := 1; i < 1<<uint(m); i++ {
		dp[i].min = math.MaxInt64
	}
	for i := 0; i < m; i++ {
		segLen := numOfCat(accum, i, n)
		movedNum := segLen - involvedNum(accum, i, 0, segLen)
		dp[1<<uint(i)].min = movedNum
		dp[1<<uint(i)].length = segLen
	}
	for i := 1; i < 1<<uint(m); i++ {
		// take DP
		for j := 0; j < m; j++ {
			if i&(1<<uint(j)) == 0 {
				continue
			}
			base := i &^ (1 << uint(j))
			if base == 0 {
				continue
			}
			segLen := numOfCat(accum, j, n)
			movedNum := segLen - involvedNum(accum, j, dp[base].length, segLen)
			num := dp[base].min + movedNum
			if dp[i].min > num {
				dp[i].min = num
				dp[i].length = dp[base].length + segLen
			}
		}
	}
	return dp[(1<<uint(m))-1].min
}

func numOfCat(accum [][]int, cat, size int) (num int) {
	return accum[cat][size]
}

func involvedNum(accum [][]int, cat, start, length int) int {
	return accum[cat][start+length] - accum[cat][start]
}

type pair struct{ min, length int }

// snip-scan-funcs
func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
func scanString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
