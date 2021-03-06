package main

import (
	"math"
)

func ReverseInteger(x int) int {
	rev := 0
	for x != 0{
		pop := x % 10

		// move to the next number (the next digit in base 10)
		x /= 10

		// TODO: write an example
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32 / 10 && pop > 7) {
			return 0
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32 / 10 && pop < -8) {
			return 0
		}
		rev = rev * 10 + pop
	}
	return rev
}