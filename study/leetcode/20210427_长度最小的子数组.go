/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-27 11:29
 * @Project : go-improve
 * @Desc: 
*/

package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {
	var slow, fast = 0,0
	var sum = 0
	var min = len(nums) + 1
	for fast < len(nums) {
   		sum += nums[fast]
   		if sum >= target {
   			for {
   				if sum - nums[slow] >= target {
   					sum = sum - nums[slow]
   					slow++
				} else {
   					break
				}
			}
   			if fast - slow + 1 < min {
   				min = fast - slow + 1
			}
		}
   		fast++

	}

	if min < len(nums) + 1 {
		return min
	} else {
		return 0
	}
}

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x

}

func main()  {

	var target = 4
	var nums = []int{1,4,4}
	var result = minSubArrayLen(target, nums)
	fmt.Println(result)
}
