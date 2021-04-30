/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-28 14:36
 * @Project : go-improve
 * @Desc: 
*/

package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1},{1,1}}
	}
	var result = make([][]int, numRows)
	var first = []int{1}
	result[0] = first
	var sec  = []int{1,1}
	result[1] = sec
	for i := 2; i< numRows; i++ {
		var tmp = make([]int, i+1)
		for j := 0; j < len(tmp); j++ {
			if j== 0 || j == len(tmp) - 1 {
				tmp[j] = 1
			} else {
				tmp[j] = result[i-1][j-1] + result[i-1][j]
			}
		}
		result[i] = tmp
	}

	return result
}

func main()  {
	var result = generate(5)
	fmt.Println(result)
}
