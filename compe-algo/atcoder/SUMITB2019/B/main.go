package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	resolve(parseProblem(os.Stdin))
}

func parseProblem(r io.Reader) int {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines

	return scanInt(sc)
}

func resolve(n int) {
	// n/1.08 <= i < (n+1)/1.08: =>  ceil(n/1.08) <= i <= floor((n+1)/1.08)
	for i := int((100*float64(n) + 107) / 108); i <= int((float64(n)+1)/1.08); i++ {
		if int(float64(i)*1.08) != n {
			continue
		}
		fmt.Println(i)
		return
	}
	fmt.Println(":(")
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
