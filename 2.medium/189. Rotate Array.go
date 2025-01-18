package medium

func rotate(nums []int, k int) {

	k %= len(nums)

	// revert the target block (k elems from the end)
	revert(nums, len(nums)-k, len(nums)-1)
	// revert the rest of the array
	revert(nums, 0, len(nums)-k-1)
	// revert the whole array
	revert(nums, 0, len(nums)-1)
}

func revert(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
