/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2020-07-01 11:32
 * @Project : go-improve
 * @Desc: 
*/

package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg int
	prob string
}

func (err *argError) Error() string {
	return fmt.Sprintf("%d - %s", err.arg, err.prob)
}

func f1(arg int) (int, error)  {
	if arg == 23 {
		return -1, errors.New("wrong number 42")
	}


	return arg + 2, nil
}

func f2(arg int) (int, error)  {
	if arg == 42 {
		return -1, &argError{arg, "wrong number"}
	}

	return arg, nil
}



func main()  {


	a, b := f1(23)

	fmt.Println("a, b", a, b)


	for _,i :=  range []int{3,42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}

	_,e := f2(42)

	ae, ok := e.(*argError)
	if  ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}


