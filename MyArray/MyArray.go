package MyArray

import "fmt"

func Demo1() {
	arr1 := [3]int{1, 2, 3}
	println("arr1 len=", len(arr1))

	arr2 := [3]int{}
	for i, _ := range arr2 {
		arr2[i] = 100 * i
	}
	fmt.Println("arr2=", arr2, " len=", len(arr2))

	arr2[0] = 999
	fmt.Println("arr2=", arr2, " len=", len(arr2))

	arr3 := [...]int{7, 8, 9}
	fmt.Println("arr3 len=", len(arr3), arr3, cap(arr3))

	s1 := arr3[0:1]
	fmt.Println("s1=", s1, len(s1), cap(s1))

	//多维数组
	var matrix [2][4]int
	matrix = [2][4]int{
		{1, 2, 3, 4},
		{10, 20, 30, 40},
	}
	fmt.Println("matrix=", matrix)
}
