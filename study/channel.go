/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2020-07-01 14:58
 * @Project : go-improve
 * @Desc: 
*/

package main

import "fmt"

func main()  {

	messages := make(chan string)


	go func() {messages <- "frankstar"}()

	msg := <-messages
	fmt.Println(msg)



	}
