/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-17 17:15
 * @Project : go-improve
 * @Desc: 
*/

package main

import "fmt"

func longestPalindrome(s string) string {

	if len(s) < 2 {
		return s
	}

	dp := make([][]bool, len(s))
	//初始化
	for i:= 0; i< len(s); i++ {
		dp[i] = make([]bool, len(s))

	}
	for i:=0; i<len(s); i++ {
		dp[i][i] = true
	}


	//遍历
	resIdx, redLen := 0, 1
	for i := 1; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if s[j] != s[i] {
				dp[j][i] = false
			} else {
				if i - j < 3 {
					dp[j][i] = true
				} else {
					dp[j][i] = dp[j+1][i-1]
				}
			}
			if dp[j][i] && i-j+1 > redLen {
				resIdx = j
				redLen = i-j + 1
			}
		}
	}
	return s[resIdx: resIdx + redLen]
}

func main()  {
	s := "babad"
	result := longestPalindrome(s)
	fmt.Println(result)
}
