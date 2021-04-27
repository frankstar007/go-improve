/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-26 17:40
 * @Project : go-improve
 * @Desc:  输入有序数组
*/

package main

func twoSum(numbers []int, target int) []int {

	for i:=0 ; i< len(numbers); i++ {
		var find = target - numbers[i]
		for j := i+1; j< len(numbers); j++ {
			if find == numbers[j] {
				return []int{i+1,j+1}
			}
		}
	}
	return []int{}
}

func twoSum2(numbers []int, target int) []int {

	var start, end = 0, len(numbers)-1
	for start < end {
		var sum = numbers[start] + numbers[end]
		if sum == target {
			return []int{start + 1, end + 1}
		} else if sum < target {
			start++
		} else {
			end--
		}
	}
	return []int{}
}





