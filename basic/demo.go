/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-03-11 16:29
 * @Project : go-improve
 * @Desc: 
*/

package basic

import (
	"fmt"
)

var container  = []string{"zero","one", "two"}

func main()  {

	container := map[int]string{0:"zero", 1:"one", 2:"two"}
	fmt.Printf("the element is %q.", container[1])
}


