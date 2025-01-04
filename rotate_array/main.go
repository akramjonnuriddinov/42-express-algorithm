package main

func main() {

}

func Rotate(nums []int, k int) []int {
	l := len(nums)
	k = k % l

	reverse(nums, 0, l-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, l-1)

	return nums
}

func reverse(nums []int, start, end int) []int {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
	return nums
}
