package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/cmplx"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("%.10f\n", resolve(parseProblem(os.Stdin)))
}

func parseProblem(r io.Reader) (int, int, int, int) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	a, b, h, m := scanInt(sc), scanInt(sc), scanInt(sc), scanInt(sc)
	return a, b, h, m
}

func resolve(a, b, h, m int) float64 {
	aR := 2 * math.Pi * (float64(h)/12.0 + float64(m)/(60.0*12))
	bR := 2 * math.Pi * (float64(m) / 60.0)
	ac := cmplx.Rect(float64(a), aR)
	bc := cmplx.Rect(float64(b), bR)
	return cmplx.Abs(ac - bc)
}

// snip-scan-funcs
func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
