package main

import "fmt"

// 1. 两数之和
// 难度 简单
// https://leetcode-cn.com/problems/two-sum/

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//你可以按任意顺序返回答案。

func main() {

	//var nums []int = []int{2, 7, 11, 15}
	//var target int = 9
	// 1
	//var nums []int = []int{3,2,4}
	//var target int = 6
	// 2
	var nums []int = []int{3, 3}
	var target int = 6
	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums)-1; i++ {
		var start = nums[i]
		for j := i + 1; j < len(nums); j++ {
			temp := nums[j] + start
			if temp == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
