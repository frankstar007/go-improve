/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-10 11:06
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

	var ch = make(chan int, 8)

	//写协程
	var wg = new(sync.WaitGroup)

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go func(num int, ch chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			ch <- num
			ch <- num * 10
		}(i, ch, wg)
	}


	//读协程
	go func(ch chan int) {
		for num := range ch {
			fmt.Println(num)
		}
	}(ch)

	//wait阻塞等待所有的写通道协程结束，待计数值变成0，wait才会返回
	wg.Wait()

	close(ch)

	time.Sleep(time.Second)

	fmt.Println("finish")




}
