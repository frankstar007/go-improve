/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-09 21:09
 * @Project : go-improve
 * @Desc: 
*/

package main

import (
	"fmt"
	"time"
)

func Producer(c chan <- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}

}


func Consumer1(c <- chan int) {
	for m := range c {
		fmt.Println(m)
	}
}


func Consumer2(c <- chan int) {
	for m := range c {
		fmt.Println(m)
	}
}

func main()  {
	c := make(chan int, 2)

	go Consumer1(c)
	go Consumer2(c)

	Producer(c)

	time.Sleep(time.Second)
}
