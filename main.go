package main

import (
	"fmt"

	// easy "github.com/Iretoms/daily-code-challenge/challenges"
	"github.com/Iretoms/daily-code-challenge/dsa"
)

func main() {
	// fmt.Println(easy.TwoSum([]int{1, 2, 3, 5}, 8))
	// fmt.Println(easy.IsPalindrome(223))
	// fmt.Println(easy.RomanToInt("XXVI"))
	// fmt.Println(easy.LongestCommonPrefix([]string{"flower","flow","flight"}))
	// fmt.Println(easy.IsValid("([])"))

	// list3 := easy.ListNode{
	// 	Val:  4,
	// 	Next: nil,
	// }

	// list2 := easy.ListNode{
	// 	Val:  2,
	// 	Next: &list3,
	// }

	// list1 := easy.ListNode{
	// 	Val:  1,
	// 	Next: &list2,
	// }

	// list6 := easy.ListNode{
	// 	Val:  4,
	// 	Next: nil,
	// }

	// list5 := easy.ListNode{
	// 	Val:  3,
	// 	Next: &list6,
	// }

	// list4 := easy.ListNode{
	// 	Val:  1,
	// 	Next: &list5,
	// }
	// fmt.Println(easy.MergeTwoLists(&list1, &list4))

	// list := dsa.LinkedList{Head: nil, Length: 0}
	// list.InsertAtHead(2)

	fmt.Println(dsa.BinarySearch([]int{1, 2, 3, 5, 6, 7, 9}, 3))
}
