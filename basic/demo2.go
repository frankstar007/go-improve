/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-03-24 16:13
 * @Project : go-improve
 * @Desc: 
*/

package basic

import "fmt"

func main()  {
	var arr =  [5]int{1,3,5,7,8}
	sum(arr)



}

func sum(arr [5]int)  {
	var sum = 0
	for k:= 0; k<len(arr); k++ {
		sum += arr[k]
	}
	fmt.Println(sum)
}
