package main

import "fmt"

func rotate(matrix [][]int)  {

	//先对角线旋转
	for m,n := range matrix {
		for i,_ := range n {
			if m < i {
				var tmp = matrix[m][i]
				matrix[m][i] = matrix[i][m]
				matrix[i][m] = tmp
			}
		}
	}
	//在对每个一维数组进行旋转

	for index,value := range matrix {
		length := len(value)
		for key,_ := range value {
			if key < length / 2 {
				tmp := matrix[index][key]
				matrix[index][key] = matrix[index][length-1-key]
				matrix[index][length-1-key] = tmp
			}

		}
	}

}

func main()  {
	matrix := [][]int {
		{5,1,9,11},
		{2,4,8,10},
		{13,3,6,7},
		{15,14,12,16},
	}

	rotate(matrix)


	//亦或交换
	a := 4
	b := 3

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println("a ,b", a, b)
}