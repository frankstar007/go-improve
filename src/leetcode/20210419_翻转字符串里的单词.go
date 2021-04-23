package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	for i := 0; i < len(strs); i++  {
		strs[i] = strings.TrimSpace(strs[i])
	}
	s = ""
	for j := len(strs)-1; j >=0; j-- {
		if strings.EqualFold(strs[j], "") {
			continue
		}
		s = s + strs[j] + " "
	}

	return strings.TrimSpace(s)
}

func main()  {
	s :=  "a good   example"
	s = reverseWords(s)
	fmt.Println(s)
}