/**
 * created by
 * @project go-improve
 * @author frankstar
 * @date 2021/5/22
 * @contact frankstarye@tencent.com
 **/

package functions

import (
	"fmt"
	"time"
)

func Limit()  {
	requests := make(chan int, 5)
	for i:=1; i<len(requests); i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<- limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i:= 0; i<3; i++ {
		burstyLimiter <- time.Now()
	}


	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequest := make(chan int, 5)
	for i:=1; i<=5; i++ {
		burstyRequest <- i
	}

	close(burstyRequest)
	for req := range burstyRequest {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}


