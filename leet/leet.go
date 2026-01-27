package leet

import "slices"

// p128 Longest Consecutive Sequence
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

// p15 3Sum
func ThreeSum(nums []int) [][]int {
	res := [][]int{}

	slices.Sort(nums)

	for i, v := range nums {
		// remaining values are positive, no more solutions
		if v > 0 {
			break
		}

		// duplicate
		if i > 0 && v == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1
		for l < r {
			threesum := v + nums[l] + nums[r]
			if threesum > 0 {
				r--
			} else if threesum < 0 {
				l++
			} else {
				res = append(res, []int{v, nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return res
}
