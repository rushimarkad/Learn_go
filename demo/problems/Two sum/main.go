//Link: https://leetcode.com/problems/two-sum/

//Statement: Given an array of integers nums and an integer target, return indices of the
//two numbers such that they add up to target.

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	ans := twoSum(nums, target)
	fmt.Println(ans)
}
