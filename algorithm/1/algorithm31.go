package main

func partitionArray(nums []int, k int) int {
	// write your code here
	if len(nums) == 0 {
		return 0
	}
	var p int = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}
	return p
}
