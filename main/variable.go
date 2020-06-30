package main

import "fmt"

func main()  {
	c1 := complex(5,7)

	c2 := 8 + 27i

	cadd := c1 + c2

	fmt.Println("sum : " , cadd)


	cmul := c1 * c2

	fmt.Println("multiply : ", cmul)


	i := 10

	var j = float64(i)

	fmt.Println(j)


	fmt.Println(calculate(2, 3))

	var area, _ = rectProps(10.8, 34.2)
	fmt.Println(area)
}

func calculate(price int, no int) int {
	var total  = price * no
	return total
}

func rectProps(length, width float64) (area, perimeter float64)  {
	area  = length * width
	perimeter = (length + width) * 2
	return

}