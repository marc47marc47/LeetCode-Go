package leetcode

import (
	"fmt"
)

func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		// 統計以 i 為中心的回文總數
		fmt.Printf("i=%v\n", i)
		count += countPalindrome(s, i, i)
		fmt.Printf("countPalindrome('%v',%v, %v) => %v\n", s, i, i, count)
		// 統計以當 i, i+1 為左右邊際時的中心，其回文總數
		count += countPalindrome(s, i, i+1)
		fmt.Printf("countPalindrome('%v',%v, %v) => %v\n", s, i, i+1, count)
	}
	return count
}

// 設定左右邊界，以中點為中心，往兩側擴展，
// 判斷是否為相同，相同即代表回文
// 回傳該中心往外的回文的總數
func countPalindrome(s string, left int, right int) int {
	var count int
	for left >= 0 && right < len(s) && s[left] == s[right] {
		fmt.Printf(" s[%v:%v]:%v\n", left, right, s[left:right+1])
		count++
		left--
		right++
	}
	return count
}
