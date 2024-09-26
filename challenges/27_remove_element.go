package challenges

func RemoveElement(nums []int, val int) int {
	pos := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[pos] = nums[i]
			pos++
		}
	}
	return pos
}

