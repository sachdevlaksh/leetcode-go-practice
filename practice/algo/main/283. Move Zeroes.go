package main

func moveZeroes(nums []int) {
	// arrNonZero := make([]int, 0)
	// arrZero := make([]int, 0)
	// arrApp := make([]int, len(nums))

	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] != 0 {
	// 		arrNonZero = append(arrNonZero, nums[i])
	// 	} else {
	// 		arrZero = append(arrZero, nums[i])
	// 	}
	// }
	// arrApp = append(arrNonZero, arrZero...)
	// copy(nums, arrApp)
	// fmt.Println(nums)

	temp := 0
	for _, val := range nums {
		if val != 0 {
			nums[temp] = val
			temp++
		}
	}
	for i := temp; i < len(nums); i++ {
		nums[i] = 0
	}

}
