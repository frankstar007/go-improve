/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-03-11 15:16
 * @Project : go-improve
 * @Desc: 
*/

package main

import (
	"flag"
	"fmt"
)

func main()  {
	var name string
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
	flag.Parse()
	fmt.Printf("Hello, %v!", name)
}
