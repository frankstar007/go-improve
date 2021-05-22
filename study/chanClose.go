/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2020-07-13 19:52
 * @Project : go-improve
 * @Desc: 
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {

	var c = make(chan int, 3)

	go func() {
		c<-1
		close(c)
	}()


	for v:= range c {
		fmt.Println(v)
	}


	var ch = make(chan int, 8)
	var wg = new(sync.WaitGroup)

	for i:=1; i<=4;i++ {
		wg.Add(1)
		go func(num int, ch chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			ch <- num
			ch <- num * 10
		}(i, ch, wg)
	}


	//读协程
	go func(ch chan int) {
		for num:= range ch{
			fmt.Println(num)
		}
	}(ch)

	//wg
	wg.Wait()

	close(ch)

	time.Sleep(time.Second)

	fmt.Println("finish")
}
