/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2020-07-13 17:06
 * @Project : go-improve
 * @Desc: 
*/

package main

import "fmt"

func main()  {

	var c = make(chan int)
	var a string

	go func() {
		a = "hello world"
		<- c
	}()

	c <- 0
	fmt.Println(a)
}
