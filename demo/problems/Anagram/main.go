//problem link : https://leetcode.com/problems/valid-anagram/
//Statement : Given two strings s and t, return true if t is an anagram of s, and false otherwise.

package main

import "fmt"

func isAnagram(s string, t string) bool {
	sSize := len(s)
	tSize := len(t)

	if sSize != tSize {
		return false
	}

	charMap := make(map[byte]int, 0)

	for i := 0; i < sSize; i++ {
		charMap[s[i]]++
	}

	for i := 0; i < tSize; i++ {
		if _, ok := charMap[t[i]]; !ok {
			return false
		}
		charMap[t[i]]--
		if charMap[t[i]] == 0 {
			delete(charMap, t[i])
		}
	}

	// if the length of charMap still has some values, then the string is not an anagram.
	return len(charMap) == 0
}

func main() {
	s := "anagram"
	t := "nagaram"
	ans := isAnagram(s, t)
	fmt.Println(ans)
}
