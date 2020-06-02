package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fmt.Println(resolve(parseProblem(os.Stdin)))
}

func parseProblem(r io.Reader) (n, p int) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n, p = scanInt(sc), scanInt(sc)
	return n, p
}

func resolve(n, p int) int {
	t := primeFactories(p)
	max := 1
	for k, ex := range t {
		max *= pow(k, ex/n)
	}
	return max
}

func pow(n, i int) (res int) {
	res = 1
	for i > 0 {
		if i&1 == 1 {
			res *= n
		}
		i >>= 1
		n *= n
	}
	return res
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeFactories(n int) (terms map[int]int) {
	terms = map[int]int{}
	num := n
	for f := 2; f*f < n; f++ {
		if num%f != 0 {
			continue
		}
		ex := 0
		for num%f == 0 {
			num /= f
			ex++
		}
		terms[f] = ex
	}
	if num > 1 {
		terms[num] = 1
	}
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
