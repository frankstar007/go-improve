/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2020-07-13 16:17
 * @Project : go-improve
 * @Desc: 
*/

package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("main coroutine")

	for i:=0; i<10; i++ {
		go testFun(i)
	}

	time.Sleep(time.Second * 1)

	fmt.Println("end")

	//c := make(chan int)
	//c <- 2
	//<- c
	//fmt.Println(c)

}

func testFun(i int) {

	fmt.Printf("child coroutine %d\n",i)

}
