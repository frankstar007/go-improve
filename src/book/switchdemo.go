/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-23 17:14
 * @Project : go-improve
 * @Desc: 
*/

package main

import "fmt"

var content string

func switchFunc(content string)  {

	switch content {

	default:
		fmt.Println("Unknown language")
	case "Python":
		fmt.Println("Python language")
	case "Go":
		fmt.Println("Go language")


	}
}

func main()  {
	content = "Python"
	switchFunc(content)
}



