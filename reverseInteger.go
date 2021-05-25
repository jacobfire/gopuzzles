package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	sign := 1
	newX := 0
	//var tempArray []int
	var tempReversedString string
	var tempString string
	var err error

	if x < 0 {
		sign = -1
		newX = x * sign
	} else {
		newX = x
	}

	tempString = strconv.Itoa(newX)
	elems := strings.Split(tempString, "")
		for i := len(tempString); i > 0; i-- {
			index := i - 1
			tempReversedString += elems[index]
		}
	res, err := strconv.Atoi(tempReversedString)
	if err != nil {
		//panic("huy")
	}

	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}
	return res * sign
}

func main() {
	x := 1534236469
	alreadyReversed := reverse(x)
	fmt.Print(alreadyReversed)
}
