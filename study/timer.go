/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2020-07-13 16:00
 * @Project : go-improve
 * @Desc: 
*/

package main

import (
	"fmt"
	"time"
)

func main()  {

	timer1 := time.NewTimer(time.Second * 2)


	<-timer1.C


	fmt.Println("timer1 expired")
}


