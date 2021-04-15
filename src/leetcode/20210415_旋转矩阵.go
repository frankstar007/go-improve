package leetcode

import "fmt"

func rotate(matrix [][]int)  {

	for m,n := range matrix {
		for i,_ := range n {
			if (m < i) {
				tmp = matrix[m][i]
				matrix[m][i] = matrix[i][m]
				matrix[i][m] = tmp
			}
		}
	}



	for m,n := range matrix {
		fmt.Fprintln("m, n", m, n)
	}

}

func main()  {
	matrix := 
}