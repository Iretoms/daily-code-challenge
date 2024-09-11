package challenges

// Given an integer x, return true if x is a
// palindrome
// , and false otherwise.

// Example 1:

// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
// Example 2:

// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
// Example 3:

// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

// func IsPalindrome(x int) bool {
// 	stringX := strconv.Itoa(x)
// 	var reversedstring string

// 	for i := len(stringX); i > 0; i-- {
// 		reversedstring += string(stringX[i-1])
// 	}

// 	intVersion, _ := strconv.Atoi(reversedstring)

// 	return x == intVersion
// }

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	reversed := 0

	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return original == reversed
}
