package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var applied []int
	var currentValue int
	var currentArray []int
	currentArray = nums
	for i := 0; i < len(currentArray); i++ {
		currentValue = currentArray[i]
		//fmt.Print(currentValue)
		//fmt.Print(currentArray)
		if currentValue > 0 {
			for j := 0; j < len(currentArray); j++ {
				if j != i {
					fmt.Print(currentArray[j])
					if currentArray[i] + currentArray[j] == target {
						ret := []int{i, j}
						fmt.Print(ret)
						return ret
					}
				}
			}
		}
	}

	return applied
}
func main() {
	nums := []int{1, 2, 3}
	target := 5
	//nums = [1, 2, 3]
	//nums = [1, 2, 3]
	fmt.Println(nums, target)
	result := twoSum(nums, 5)

	fmt.Print(result)
}

