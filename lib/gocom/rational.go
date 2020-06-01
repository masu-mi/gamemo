package gocom

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

type rational struct {
	frac, denomi int
}

func ratMul(x, y rational) rational {
	r := rational{
		frac:   x.frac * y.frac,
		denomi: x.denomi * y.denomi,
	}
	d := gcd(r.frac, r.denomi)
	r.frac /= d
	r.denomi /= d
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
	d := gcd(r.frac, r.denomi)
	r.frac /= d
	r.denomi /= d
	return r
}

func ratSub(x, y rational) rational {
	return ratAdd(x, rational{frac: -y.frac, denomi: y.denomi})
}

func (r rational) Float64() float64 {
	return float64(r.frac) / float64(r.denomi)
}

func nextRational(sc *bufio.Scanner) rational {
	s := nextString(sc)
	i := strings.Index(s, ".")
	if i != -1 {
		decimalDigit := len(s) - i - 1
		fr, _ := strconv.Atoi(strings.Replace(s, ".", "", 1))
		return rational{frac: fr, denomi: int(math.Pow10(decimalDigit))}
	}
	i = strings.Index(s, "/")
	if i != -1 {
		fr, _ := strconv.Atoi(s[0:i])
		de, _ := strconv.Atoi(s[i+1 : len(s)])
		return rational{frac: fr, denomi: de}
	}
	fr, _ := strconv.Atoi(s)
	return rational{frac: fr, denomi: 1}
}
