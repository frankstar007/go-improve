/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-03-11 15:34
 * @Project : go-improve
 * @Desc: 
*/

package main

import "fmt"

var block  = "package"

func main()  {

	block	:= "function"
	{
		block := "inner"
		fmt.Printf("the block is %s.\n", block)
	}
	fmt.Printf("this block is %s", block)
}
