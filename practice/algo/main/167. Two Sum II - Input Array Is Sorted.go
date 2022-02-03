package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	low := 0
	high := len(numbers) - 1
	out := make([]int, 0)
	for low < high {
		if numbers[low]+numbers[high] > target {
			high--
		} else if numbers[low]+numbers[high] < target {
			low++
		} else {
			out = append(append(out, low+1), high+1)
			break
		}
	}
	fmt.Println(out)
	return out
}
