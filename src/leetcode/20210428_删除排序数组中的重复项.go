/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-28 19:28
 * @Project : go-improve
 * @Desc: 
*/

package main

import "fmt"

func removeDuplicates(nums []int) int {

	if len(nums) == 1 {
		return 1
	}

	var slow, fast = 0, 1
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			fast++
		} else {
			for fast < len(nums) && nums[slow] == nums[fast] {

				nums = append(nums[:slow+1], nums[fast+1:]...)
			}
		}
	}
	return len(nums)
}

func main() {
	var nums = []int{0,0,1,1,1,2,2,3,3,4}
	var result = removeDuplicates(nums)
	fmt.Println(result)
}