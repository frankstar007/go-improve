/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-28 15:15
 * @Project : go-improve
 * @Desc: 
*/

package main


func getRow(rowIndex int) []int {
	var result = make([]int, rowIndex + 1)
	for i:=0; i< len(result); i++ {
		result[i] = 1
	}

	if rowIndex < 2 {
		return result
	}

	for i:= 1; i<rowIndex; i++ {
		for j := i; j > 0; j-- {
			result[j] = result[j] + result[j-1]
		}
	}

	return result
}
