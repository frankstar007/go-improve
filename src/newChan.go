/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2020-07-13 17:19
 * @Project : go-improve
 * @Desc: 
*/

package main

import "fmt"

func main()  {

	fmt.Println("run in main coroutine")

	count := 10
	c := make(chan  bool, 10)

	for i:=0; i<count; i++ {
		go testFunc(i, c)
	}

	for i:= 0; i<count; i++ {
		<- c
	}

}

func testFunc(i int, c chan bool) {

	fmt.Printf("child coroutine %d\n",i)
	c <- true

}
