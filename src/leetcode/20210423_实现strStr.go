/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-23 14:48
 * @Project : go-improve
 * @Desc: 
*/

package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	//如果needle是空串
	if strings.Compare(needle, "") == 0 {
		return 0
	}
	if strings.Compare(haystack, "") == 0 {
		return -1
	}
	var index = kmp(haystack, needle)
	return index

}

func kmp(str string, pattern string) int {
	//先获取到next数组
	var prefix = next(pattern, len(pattern))

	var i, j = 0, 0

	for i < len(str) && j < len(pattern) {
		if j < 0 || str[i] == pattern[j]  {
			i = i + 1
			j = j + 1
		} else {
			j = prefix[j]
		}
	}
	if j == len(pattern) {
		return i-j
	}

	return -1

}

func next(pattern string, n int) []int {
	var prefix = make([]int, n)
	prefix[0] = -1
	var i, j = -1, 0
	for j < n -1 {
		if i < 0 || pattern[j] == pattern[i] {
			j = j + 1
			i = i + 1
			prefix[j] = i
		} else {
			i = prefix[i]
		}
	}
	return prefix
}


func main() {
	var hayStack = "aaaaa"
	var needle = "bba"
	var result = strStr(hayStack, needle)
	fmt.Println(result)
}
