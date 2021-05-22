/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-26 19:29
 * @Project : go-improve
 * @Desc: 移除元素
*/

package main

func removeElement(nums []int, val int) int {
	var slow, fast = 0,0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow ++
		}
		fast ++
	}
	return slow
}