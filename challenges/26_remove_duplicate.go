package challenges

//not right
// func RemoveDuplicates(nums []int) int {
// 	hashTable := make(map[int]int)
// 	var newArray []int

// 	for idx, val := range nums {
// 		_, exist := hashTable[val]
// 		if !exist {
// 			hashTable[val] = idx
// 		}
// 	}

// 	for key, _ := range hashTable {
// 		newArray = append(newArray, key)
// 	}

// 	return len(newArray)
// }

//right
// func RemoveDuplicates(nums []int) int {
//     if len(nums) == 0 {
//         return 0
//     }

//     hashTable := make(map[int]bool)
//     pos := 0

//     for _, val := range nums {
//         if !hashTable[val] {
//             nums[pos] = val
//             pos++
//             hashTable[val] = true
//         }
//     }

//     return pos
// }

//right
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	replace := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[replace] = nums[i]
			replace++
		}
	}

	return replace
}
