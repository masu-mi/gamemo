package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(resolve(parseProblem(os.Stdin)))
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
	p, q := scanInt(sc), scanInt(sc)
	base := 1
	d := newRational(p, q).denomi
	for k := range primeFactories(d) {
		base *= k
	}
	return base
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
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

func primes(n int) chan int {
	ch := make(chan int)
	go func() {
		t := eratosthenes(n)
		for i := 2; i < len(t); i++ {
			if !t[i] {
				continue
			}
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func primeFactories(n int) (terms map[int]int) {
	terms = map[int]int{}
	num := n
	for f := 2; f*f <= n; f++ {
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

type rational struct {
	frac, denomi int
	neg          int
}

func newRational(frac, denomi int) rational {
	r := rational{frac: frac, denomi: denomi}
	r.canonicalize()
	return r
}

func ratMul(x, y rational) rational {
	r := rational{
		frac:   x.frac * y.frac,
		denomi: x.denomi * y.denomi,
	}
	r.canonicalize()
	return r
}

func ratDiv(x, y rational) rational {
	return ratMul(x, ratInv(y))
}

func ratInv(x rational) rational {
	return rational{frac: x.denomi, denomi: x.frac}
}

func ratAdd(x, y rational) rational {
	r := rational{
		frac:   x.frac*y.denomi + y.frac*x.denomi,
		denomi: x.denomi * y.denomi,
	}
	r.canonicalize()
	return r
}

func ratSub(x, y rational) rational {
	return ratAdd(x, rational{frac: -y.frac, denomi: y.denomi})
}

func (r rational) Float64() float64 {
	b := float64(r.frac) / float64(r.denomi)
	if r.neg == 1 {
		return -b
	}
	return b
}

func (r rational) String() string {
	b := fmt.Sprintf("%d/%d", r.frac, r.denomi)
	if r.neg == 1 {
		return "-" + b
	}
	return b
}

func (r rational) equal(other rational) bool {
	r.canonicalize()
	other.canonicalize()

	return r.neg == other.neg &&
		r.frac == other.frac &&
		r.denomi == other.denomi
}

func (r *rational) canonicalize() {
	if r.frac < 0 {
		r.neg ^= 1
		r.frac = -r.frac
	}
	if r.denomi < 0 {
		r.neg ^= 1
		r.denomi = -r.denomi
	}
	d := gcd(r.frac, r.denomi)
	r.frac /= d
	r.denomi /= d
}

func nextRational(sc *bufio.Scanner) rational {
	s := nextString(sc)
	i := strings.Index(s, ".")
	if i != -1 {
		decimalDigit := len(s) - i - 1
		fr, _ := strconv.Atoi(strings.Replace(s, ".", "", 1))
		return newRational(fr, int(math.Pow10(decimalDigit)))
	}
	i = strings.Index(s, "/")
	if i != -1 {
		fr, _ := strconv.Atoi(s[0:i])
		de, _ := strconv.Atoi(s[i+1 : len(s)])
		return newRational(fr, de)
	}
	fr, _ := strconv.Atoi(s)
	return newRational(fr, 1)
}
func gcd(a, b int) int {
	for b > 0 {
		t := a / b
		a, b = b, a-t*b
	}
	return a
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

func resolve(n int) int {
	return n
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
