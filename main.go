/**
 * created by
 * @project go-improve
 * @author frankstar
 * @date 2021/5/22
 * @contact frankstarye@tencent.com
 **/


package main

import (
	"fmt"
	"github.com/frankstar007/go-improve/mysql"

)

func main()  {
	fmt.Println("start go")

	//functions.Limit()

	var u = mysql.Select(1)
	fmt.Println(u)
}