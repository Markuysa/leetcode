package easy

func majorityElement(nums []int) int {
	times := len(nums) / 2
	frequency := map[int]int{}

	for _, v := range nums {
		frequency[v]++
		if frequency[v] > times {
			return v
		}
	}

	return 0
}
