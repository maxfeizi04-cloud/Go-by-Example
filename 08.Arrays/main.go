package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("a = ", a)

	a[4] = 100
	fmt.Println("a = ", a)
	fmt.Println("a[4] = ", a[4])

	fmt.Println("len = ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b = ", b)

	b = [...]int{100, 3: 400, 52}
	fmt.Println("b = ", b)

	// 二维数组两行三列
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{3, 4, 5},
	}
	fmt.Println(twoD)
}
