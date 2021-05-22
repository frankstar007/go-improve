/**
 * @Author: frankstar
 * @Email: frankstar007@163.com
 * @Date: 2021-04-26 14:11
 * @Project : go-improve
 * @Desc: 
*/

package main

import "fmt"

func arrayPairSum(nums []int) int {
	//快排
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	var sum = 0
	for i:=0; i<len(nums);i+=2 {
		sum += nums[i]
	}
	return sum
}

func quickSort(nums []int, start, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		quickSort(nums, start, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

func partition(arr []int, left int, right int) int {

	pivot := left
	index := pivot + 1

	for i:= index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index++
		}
	}
	swap(arr, pivot, index-1)
	return index - 1
}

func swap(arr []int, i,j int) {
	arr[i], arr[j] = arr[j], arr[i]
}


func main() {
	var nums = []int{1,4,3,2}
	arrayPairSum(nums)
}
