package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
	resolve()
}

func resolve() {
	// load and prepare
	h, w := scanInt(sc), scanInt(sc)
	grid := make([]string, h)
	whiteRow := make([]bool, h)
	whiteNum := make([]int, w)
	for i := 0; i < h; i++ {
		grid[i] = scanString(sc)
		whiteRow[i] = true
		for j := range grid[i] {
			if grid[i][j] != '.' {
				whiteRow[i] = false
			} else {
				whiteNum[j]++
			}
		}
	}

	// build output
	builder := &strings.Builder{}
	for i := 0; i < h; i++ {
		if whiteRow[i] {
			continue
		}
		for j := 0; j < w; j++ {
			if whiteNum[j] == h {
				continue
			}
			builder.WriteByte(grid[i][j])
		}
		builder.WriteByte('\n')
	}
	// write answer
	fmt.Printf("%s", builder.String())
	return
}

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
