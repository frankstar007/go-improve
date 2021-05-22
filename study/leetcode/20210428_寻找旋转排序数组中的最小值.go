/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-28 17:31
 * @Project : go-improve
 * @Desc: 
*/

package main


func findMin(nums []int) int {

	var slow, fast = 0,1
	for fast < len(nums) {
		if nums[slow] < nums[fast] {
			fast ++
		}else {
			return nums[fast]
		}
	}
	return nums[slow]
}


