package dsa

func BinarySearch(array []int, number int) bool {
	found := false
	low := 0
	high := len(array) - 1

	if len(array) <= 0 {
		return found
	}

	for low <= high {
		mid := (high + low) / 2
		if array[mid] == number {
			found = true
			break
		}
		if array[mid] < number {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return found
}
