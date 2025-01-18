package easy

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}

func removeDuplicatesv2(nums []int) int {
	index := 0
	for _, v := range nums {
		if nums[index] != v {
			index++
			nums[index] = v
		}
	}

	return index + 1
}
