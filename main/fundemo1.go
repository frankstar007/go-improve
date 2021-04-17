/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-07 11:55
 * @Project : go-improve
 * @Desc: 
*/


package main

import (
	"errors"
	"fmt"
)

type operate func(x,y int) int

func calculating(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

func main()  {

	x, y := 34, 32
	fmt.Println(x, y)
}
