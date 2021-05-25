package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	if x < math.MinInt32 || x > math.MaxInt32 {
		return false
	}
	var div int = 1
	var l, r int
	//if x % 10 == 0 {
	//	return true
	//}
	if x < 0 {
		return false
	}
	for x / div >= 10 {
		div *= 10
	}
	for x != 0 {
		l = x / div
		r = x % 10
		if l != r {
			return false
		}

		x = (x % div) / 10
		div /= 100
	}

	return true
}

func main() {
	x := 10
	res := isPalindrome(x)
	fmt.Print(res)
}
