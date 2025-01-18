package medium

func removeDuplicatesV2(nums []int) int {
	index := 1
	occurance := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			occurance++
		} else {
			occurance = 1
		}

		if occurance <= 2 {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}

func removeDuplicates(nums []int) int {
	length := len(nums)
	frequency := map[int]int{}
	for i := 0; i < length; i++ {
		frequency[nums[i]]++
		if frequency[nums[i]] > 2 {
			moveArrayRight(&nums, i)

			length--
			i--
		}
	}

	return length
}

func moveArrayRight(nums *[]int, index int) {
	saving := 0
	for i := index; i < len(*nums)-1; i++ {
		saving = (*nums)[i]
		(*nums)[i] = (*nums)[i+1]
		(*nums)[i+1] = saving
	}
}
