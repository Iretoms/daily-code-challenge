package challenges

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]

		value, ok := seen[diff]
		if ok {
			return []int{value, i}
		}
		seen[nums[i]] = i
	}
	return nil
}
