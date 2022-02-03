package main

import (
	"fmt"
)

func rotate(nums []int, k int) {
	var outArr = make([]int, len(nums))
	j := 0
	for len(nums) < k {
		k = k - len(nums)
	}
	if len(nums) > k {
		for i := len(nums) - k; i < len(nums); i++ {
			outArr[j] = nums[i]
			j++
		}
		j = 0
		for i := k; i < len(nums); i++ {
			outArr[i] = nums[j]
			j++
		}
	} else if len(nums) < k {
		for i := k - len(nums); i < len(nums); i++ {
			outArr[j] = nums[i]
			j++
		}
		j = 0
		for i := k; i < len(nums); i++ {
			outArr[i] = nums[j]
			j++
		}
	}
	nums = outArr
	fmt.Println(nums)
}
