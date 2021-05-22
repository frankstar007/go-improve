/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-28 15:44
 * @Project : go-improve
 * @Desc: 
*/

package main

import (
	"fmt"
	"strings"
)

func reverseWord(s string) string {
	var strs = strings.Split(s, " ")
	var result = ""
	for i:= 0; i< len(strs); i++ {
		result = result + reverse(strs[i]) + " "
	}

	return strings.Trim(result, " ")
}

func reverse(s string) string{
	var result = ""
	for  end := len(s) - 1; end >=0; end-- {
		result += string (s[end])
	}
	return result
}

func main()  {
	var result = reverseWord("let me see")
	fmt.Println(result)
}



