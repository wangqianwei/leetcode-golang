package main

import "fmt"

// 4. 寻找两个正序数组的中位数
// 难度 困难
// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

func main() {

	//var nums1, nums2 []int = []int{1, 3}, []int{2}
	//var nums1, nums2 []int = []int{1, 2}, []int{3, 4}
	//var nums1, nums2 []int = []int{0, 0}, []int{0, 0}
	//var nums1, nums2 []int = []int{}, []int{1}
	var nums1, nums2 []int = []int{2}, []int{}

	r := findMedianSortedArrays(nums1, nums2)
	fmt.Println(r)

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	var i, j, sum []int = nums1, nums2, []int{}
	var ic, jc int = 0, 0
	for c := 0; c < len(i)+len(j); c++ {
		if ic != len(i) && jc != len(j) {
			if i[ic] < j[jc] {
				sum = append(sum, i[ic])
				ic++
			} else {
				sum = append(sum, j[jc])
				jc++
			}
		} else {
			if ic != len(i) {
				for {
					if ic == len(i) {
						break
					}
					sum = append(sum, i[ic])
					ic++
				}
			}
			if jc != len(j) {
				for {
					if jc == len(j) {
						break
					}
					sum = append(sum, j[jc])
					jc++
				}
			}
		}
	}

	if len(sum)%2 == 0 {
		mid := len(sum) / 2
		return float64((sum[mid-1] + sum[mid])) / 2
	} else {
		mid := len(sum) / 2
		return float64(sum[mid])
	}
	return 0
}
