package main

import (
	"fmt"
	"math"
)

func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	}

	convertedN := float64(n)
	processedN := math.Log2(convertedN)

	buffer := processedN
	if buffer != float64(int(buffer)) {
		return false
	}

	x := int(processedN) % 2

	if (math.MinInt32 > x) || (x > math.MaxInt32) || x != 0  {
		return false
	}

	return true
}

func main() {
	checkInt := 1
	res := isPowerOfFour(checkInt)
	fmt.Println(res)
}
