//Link : https://leetcode.com/problems/longest-substring-without-repeating-characters/
//Statement: Given a string s, find the length of the longest substring without repeating characters.

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		set := make(map[byte]bool)
		length := 0
		for j := i; j < len(s); j++ {
			if set[s[j]] {
				break
			}

			set[s[j]] = true
			length++
		}

		if result < length {
			result = length
		}
	}

	return result
}

func main() {
	s := "abcabcbb"
	ans := lengthOfLongestSubstring(s)
	fmt.Println(ans)
}
