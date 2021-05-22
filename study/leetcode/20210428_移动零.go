/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-28 20:02
 * @Project : go-improve
 * @Desc: 
*/

package main

import "fmt"

func moveZeroes(nums []int)  {
	var start, end = 0, len(nums)

	for start < end {
		if nums[start] == 0 {
			nums = append(nums[:start], nums[start+1:]...)
			end --
			nums = append(nums, 0)
		} else {
			start ++
		}
	}
}

func main()  {
	var nums = []int{0,0,1,3,12}
	moveZeroes(nums)
	fmt.Println(nums)
}