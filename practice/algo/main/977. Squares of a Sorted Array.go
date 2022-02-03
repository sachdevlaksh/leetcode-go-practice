package main

func sortedSquares(nums []int) []int {
	length := len(nums)
	ans := make([]int, length)
	for left, right, write := 0, length-1, length-1; write >= 0; write-- {
		if -nums[left] > nums[right] {
			ans[write] = nums[left] * nums[left]
			left = left + 1
		} else {
			ans[write] = nums[right] * nums[right]
			right--
		}
	}
	return ans
}
