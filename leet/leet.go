package leet

import "slices"

// p128
func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slices.Sort(nums)

	res := 0
	curr := nums[0]
	streak := 0
	i := 0

	for i < len(nums) {
		if curr != nums[i] {
			curr = nums[i]
			streak = 0
		}
		// iterate through dupes
		for i < len(nums) && nums[i] == curr {
			i++
		}
		streak++
		curr++
		res = max(res, streak)
	}
	return res
}
