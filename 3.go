package main

import "fmt"

// 3. 无重复字符的最长子串
// 难度 中等
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

func main() {

	//var s string = "abcabcbb"
	//var s string = "bbbbb"
	//var s string = "pwwkew"
	var s string = ""

	i := lengthOfLongestSubstring(s)
	fmt.Println(i)
}

func lengthOfLongestSubstring(s string) int {

	var i int = 0
	for j := 0; j < len(s); j++ {
		var temp map[string]string = make(map[string]string)
		for _, v := range s[j:] {
			if _, ok := temp[string(v)]; ok {
				if len(temp) > i {
					i = len(temp)
				}
				break
			}
			temp[string(v)] = string(v)
		}
	}

	return i
}
