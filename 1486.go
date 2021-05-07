package main

import (
	"fmt"
)

// 1486. 数组异或操作
// 难度 简单
// https://leetcode-cn.com/problems/xor-operation-in-an-array/

func main() {

	var n, start int = 5, 0
	//var n, start int = 4, 3
	//var n, start int = 1, 7
	//var n, start int = 10, 5
	var nums []int = make([]int, n)
	var result int

	for i := 0; i < n; i++ {
		nums[i] = start + i*2;
	}
	fmt.Println(nums)
	result = xorOperation(n, start)
	fmt.Println(result)
}

func xorOperation(n int, start int) int {

	var result int = 0
	for i := 0; i < n; i++ {
		result = result ^ (start + i*2)
	}
	return result
}
