package main

import "fmt"

func main() {
	nums := []int{7, 12, 3, 9, 14, 6, 5, 21, 8, 1}
	sum := 140
	fmt.Println(nums)
	num1, num2 := twoSum(nums, sum)
	if num1 == 0 && num2 == 0 && sum != 0 {
		fmt.Println("None of the numbers added up to", sum)
	} else {
		fmt.Println("Two nums that add up to", sum, "are", num1, "and", num2)
	}
}

func twoSum(nums []int, sum int) (int, int) {
	complements := make(map[int]bool)
	for _, val := range nums {
		complement := sum - val
		if complements[complement] == true {
			return complement, val
		} else {
			complements[val] = true
		}
	}
	return 0, 0
}
