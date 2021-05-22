package main

func main()  {
	matrix := [][]int {
		{0,1,2,0},
		{3,4,5,2},
		{1,3,1,5},
	}

	findDiagonalOrder(matrix)


}


func findDiagonalOrder(mat [][]int) []int {

	var total = len(mat) + len(mat[0]) - 1

	nums := make(map[int][]int)
	for m,values := range mat {
		for n,_ := range values {
			sum := m + n
			nums[sum] = append(nums[sum], mat[m][n])
		}
	}
	result := make([]int, len(mat) * len(mat[0]))

	index := 0
	for i := 0; i < total; i = i+1 {

		for j:=0; j < len(nums[i]) && i % 2 != 0; j++ {
			result[index] = nums[i][j]
			index = index + 1
		}
		for j:=len(nums[i]) - 1; j >= 0 && i % 2 == 0; j-- {
			result[index] = nums[i][j]
			index = index + 1
		}
	}
	return result
}