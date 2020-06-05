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

func parseProblem(r io.Reader) (int, *bufio.Scanner) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines

	return scanInt(sc), sc
}

func resolve(q int, sc *bufio.Scanner) {
	n := int(1e5)
	primes := eratosthenes(n + 2)
	acc := make([]int, n+2)
	sum := 0
	for i := 1; i <= n+1; i++ {
		acc[i] = sum
		if primes[i] && primes[(i+1)/2] {
			sum++
		}
	}
	res := []int{}
	for i := 0; i < q; i++ {
		l, r := scanInt(sc), scanInt(sc)
		res = append(res, acc[r+1]-acc[l])
	}
	for _, r := range res {
		fmt.Println(r)
	}
	return
}

func eratosthenes(n int) []bool {
	t := make([]bool, n)
	for i := 2; i < len(t); i++ {
		t[i] = true
	}
	for i := 2; i < len(t); i++ {
		if !t[i] {
			continue
		}
		for j := 2; j*i < len(t); j++ {
			t[j*i] = false
		}
	}
	return t
}

// snip-scan-funcs
func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
