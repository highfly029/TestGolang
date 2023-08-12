package MySlice

import "fmt"

func Demo() {
	slice1 := []int{1, 2, 3}
	fmt.Println("slice1=", slice1, " len=", len(slice1), cap(slice1))

	slice2 := make([]int, 3)
	fmt.Println("slice2=", slice2, " len=", len(slice2), cap(slice2))

	slice3 := make([]int, 3, 5)
	fmt.Println("slice3=", slice3, " len=", len(slice3), cap(slice3))

	slice4 := slice1[1:2]
	fmt.Println("slice4=", slice4, " len=", len(slice4), cap(slice4))

	slice5 := slice1[1:]
	fmt.Println("slice5=", slice5, " len=", len(slice5), cap(slice5))

	slice6 := slice1[:2]
	fmt.Println("slice6=", slice6, " len=", len(slice6), cap(slice6))

	slice7 := append(slice1, 4)
	fmt.Println("slice7=", slice7, " len=", len(slice7), cap(slice7))

	slice8 := append(slice1, 4, 5, 6)
	fmt.Println("slice8=", slice8, " len=", len(slice8), cap(slice8))

	slice9 := append(slice1, slice2...)
	fmt.Println("slice9=", slice9, " len=", len(slice9), cap(slice9))

	slice10 := append(slice1[:1], slice2...)
	fmt.Println("slice10=", slice10, " len=", len(slice10), cap(slice10))

	slice11 := append(slice1[:1], slice2[1:]...)
	fmt.Println("slice11=", slice11, " len=", len(slice11), cap(slice11))

	slice12 := append(slice1[:1], slice2[1:2]...)
	fmt.Println("slice12=", slice12, " len=", len(slice12), cap(slice12))

}
