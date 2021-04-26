/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-23 18:07
 * @Project : go-improve
 * @Desc: 
*/

package main

import "errors"

func main() {
	outerFunc()
}

func outerFunc() {
	innerFunc()
}

func innerFunc() {
	panic(errors.New("this. is a panic error"))
}
