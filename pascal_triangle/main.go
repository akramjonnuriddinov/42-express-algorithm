package main

import "fmt"

func main() {
	fmt.Println(generate(100))
}

func generate(n int) [][]int {
	if n < 1 {
		return [][]int{}
	}
	row := []int{1}
	result := [][]int{row}

	for i := 0; i < n-1; i++ {
		row = generate_row(row)
		result = append(result, row)
	}

	return result
}

func generate_row(prev_ []int) []int {
	next_ := []int{1}
	for i := 0; i < len(prev_)-1; i++ {
		next_ = append(next_, prev_[i]+prev_[i+1])
	}
	next_ = append(next_, 1)
	return next_
}
