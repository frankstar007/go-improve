/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2020-07-13 17:55
 * @Project : go-improve
 * @Desc: 
*/

package main

import (
	"fmt"
	"time"
)

func main()  {

	c := make(chan int, 2)

	go consumer1(c)
	go consumer2(c)

	producer(c)

	time.Sleep(time.Second)


}

func producer(c chan<- int)  {

	for i:=0; i<10; i++ {
		c <- i
	}
}

func consumer1(c <-chan int)  {
	for m:=range c {
		fmt.Printf("1 get num : %v\n", m)
	}
}


func consumer2(c <-chan int)  {
	for m:=range c {
		fmt.Printf("2 get num : %v\n", m)
	}
}
