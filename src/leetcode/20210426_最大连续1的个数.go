/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-26 19:44
 * @Project : go-improve
 * @Desc: 
*/

package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {

	var max, cnt = 0, 0
	for index := 0; index < len(nums); index++ {
		if nums[index] == 1 {
			cnt ++
		} else {
			max = maxInt(max, cnt)
			cnt = 0
		}
	}
	return maxInt(max, cnt)
}

func maxInt(x, y int) int{
	if x > y {
		return x
	}
	return y
}

func main()  {

	var nums = []int{1,1,0,1,1,1}
	var result = findMaxConsecutiveOnes(nums)
	fmt.Println(result)
}

