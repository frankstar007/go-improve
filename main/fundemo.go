/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-07 11:34
 * @Project : go-improve
 * @Desc: 
*/

package main

import "fmt"

type Printer func(content string) (n int, err error)

func printToStd(content string) (bytesNum int, err error)  {
	return fmt.Println(content)
}

func main()  {
	var p Printer

	p = printToStd
	p("I am here")
}