/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-10 11:01
 * @Project : go-improve
 * @Desc: 
*/

package main

import "fmt"

func main()  {
	var c = make(chan int, 3)

	//子协程写
	go func() {
		c <- 1
		close(c)
	}()


	//fmt.Println(<- c)
	//fmt.Println(<- c)


	for v := range c {
		fmt.Println(v)
	}
}
