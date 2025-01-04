package main

import "fmt"

func main() {
	fmt.Println(MoveZeroes([]int{0, 1, 0, 3, 12}))
}

func MoveZeroes(nums []int) []int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
			continue
		}
		nums[i], nums[i-count] = nums[i-count], nums[i]
	}
	return nums
}
