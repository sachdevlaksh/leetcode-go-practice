package main

import "fmt"

func main1() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	search(nums, target)
}

func search(nums []int, target int) int {
	for index, num := range nums {
		if num == target {
			fmt.Println(index)
			return index
		}
	}
	fmt.Println(-1)
	return -1
}
