/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2020-07-01 14:06
 * @Project : go-improve
 * @Desc: 
*/

package main

import "fmt"

func f(from string)  {

	for i:=0; i<3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main()  {

	f("direct")


	go f("goroutine")



	go func(msg string) {
		fmt.Println(msg)
	}("going")


	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
