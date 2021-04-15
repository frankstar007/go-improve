package main

import "fmt"

func setZeroes(matrix [][]int)  {
	var rows []bool = make([]bool, len(matrix))
	var cols []bool = make([]bool, len(matrix[0]))
	for m,value := range matrix {
		for n,val := range value {
			if val == 0 {
				rows[m] = true
				cols[n] = true
			}
		}
	}
	for m,value := range matrix {
		for n,_ := range value {
			if rows[m] || cols[n] {
				matrix[m][n] = 0
			}
		}
	}
	fmt.Println(matrix)

}

func main()  {
	matrix := [][]int {
		{0,1,2,0},
		{3,4,5,2},
		{1,3,1,5},
	}

	setZeroes(matrix)

}