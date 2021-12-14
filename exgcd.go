package crypto

func extendedEuclideanAlgorithm(a, b int64) (int64, int64, int64) {
	// a*x + b*y == gcd(a, b)
	if b == 0 {
		return a, 1, 0
	}
	gcd, x, y := extendedEuclideanAlgorithm(b, a%b)
	return gcd, y, x - (a/b)*y
}

func invNP(n, p int64) int64 {
	if n < 0 {
		n = n%p+p
	}
	gcd, x, _ := extendedEuclideanAlgorithm(n, p)
	if gcd != 1 {
		panic("gcd not 1")
	}
	return mod(x, p)
}

func mod(x, y int64) int64 {
	// y > 0
	r := x%y
	if r < 0 {
		r += y
	}
	return r
}
