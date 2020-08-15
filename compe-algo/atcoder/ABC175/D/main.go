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

func parseProblem() int {
	n, k := scanInt(sc), scanInt(sc)
	ps := nextIntSlice(sc, n)
	cs := nextIntSlice(sc, n)
	max := math.MinInt64
	for i := 0; i < n; i++ {
		if v := score(n, k, ps, cs, i); max < v {
			max = v
		}
	}
	return max
}

func score(n, k int, ps, cs []int, idx int) int {
	scores := make([]int, n)
	maxScores := make([]int, n)
	for i := 0; i < len(maxScores); i++ {
		maxScores[i] = math.MinInt64
	}
	steps := make([]int, n)

	// setup
	cur := idx
	visited := map[int]struct{}{}
	stepNum, score, maxScore := 0, 0, 0
	for i := 0; i < n && i < k; i++ {
		steps[cur] = stepNum
		visited[cur] = struct{}{}

		nextPos := ps[cur] - 1
		score += cs[nextPos]
		stepNum++
		if _, ok := visited[nextPos]; ok {
			break
		}
		cur = nextPos
		scores[stepNum] = score
		if maxScore < score {
			maxScore = score
		}
		maxScores[stepNum] = maxScore
	}
	if k == stepNum {
		return findMax(score, maxScores)
	}

	// k is long
	nextPos := ps[cur] - 1
	startStep := steps[nextPos]
	loopSize := stepNum - startStep
	loopGain := score - scores[startStep]
	if loopGain < 0 {
		return findMax(score, maxScores)
	}

	availableLoopNum := (k - steps[startStep]) / loopSize
	restSteps := k - startStep - availableLoopNum*loopSize
	maxRestPart := 0
	for i := 0; i < restSteps; i++ {
		if v := scores[startStep+i]; maxRestPart < v {
			maxRestPart = v
		}
	}
	baseMax := findMax(score, maxScores)
	candidates := []int{
		maxScores[startStep],
		scores[startStep] + availableLoopNum*loopGain + maxRestPart,
		baseMax + (availableLoopNum-1)*loopGain,
	}
	return max(candidates)
}
func max(a []int) int {
	m := a[0]
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
}

func findMax(score int, maxScores []int) int {
	max := score
	for i := 1; i < len(maxScores); i++ {
		if max < maxScores[i] {
			max = maxScores[i]
		}
	}
	return max
}

// package: gocom
// packed src of [/Users/masumi/dev/src/github.com/masu-mi/gamemo/lib/gocom/next.go] with goone.

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
