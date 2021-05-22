/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-17 16:36
 * @Project : go-improve
 * @Desc: 
*/

package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) < i + 1 || strs[j][i] != strs[0][i] {
				if i == 0 {
					return ""
				} else {
					return strs[0][0:i]
				}
			}
		}
	}
	return strs[0]
}

func main()  {
	strs := []string {
		"flower",
		"flp",
		"fl",
	}

	result := longestCommonPrefix(strs)
	fmt.Println(result)
}