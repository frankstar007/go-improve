/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-03-12 14:30
 * @Project : go-improve
 * @Desc: 
*/

package main

import "fmt"

func main()  {
	aMap := map[string]int {
		"one" : 1,
		"two" : 2,
		"three" : 3,
	}

	k := "two"
	v, ok := aMap[k]
	if ok {
		fmt.Printf("the element of key %q : %d", k, v)
	} else {
		fmt.Println("Not found")
	}
}
