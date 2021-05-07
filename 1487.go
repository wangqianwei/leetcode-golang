package main

import (
	"fmt"
	"strings"
)

// 1487. 保证文件名唯一
// 难度 中等
// https://leetcode-cn.com/problems/making-file-names-unique/
func main() {

	//var names []string = []string{"pes", "fifa", "gta", "pes(2019)"}
	//var names []string = []string{"gta", "gta(1)", "gta", "avalon"}
	var names []string = []string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece(3)", "onepiece"}
	getFolderNames(names)
}

func getFolderNames(names []string) []string {

	fmt.Println(names)
	var temp map[string]int = make(map[string]int)
	var result []string
	for _, v := range names {
		if _, ok := temp[v]; !ok {
			index := strings.LastIndex(v, "(")
			if index != -1 {
				// 不存在
			}
			result = append(result, v)
		}
	}
	fmt.Println(result)
	return result

}
