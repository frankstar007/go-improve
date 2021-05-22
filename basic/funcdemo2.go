/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-07 19:54
 * @Project : go-improve
 * @Desc: 
*/

package basic

import "fmt"

func main()  {
	arr1 := [3]string{"a","b","c"}
	fmt.Println("arr : ", arr1)
	arr2 := modifyArr(arr1)
	fmt.Println("arr :", arr1)
	fmt.Println("arr : ", arr2)
}

func modifyArr(strings [3]string) [3]string {
	strings[1] = "x"
	return strings
}
