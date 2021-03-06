package go_learning

import "fmt"

// 数组和slice
func arraySlice() {
	a := [3]int8{1, 2, 3}
	a[1] = 4

	fmt.Println(a)

	b := a[1:]
	b[0] = 6

	/*
		在此之前a、b底层数组一样。
		在此之后 append 扩充了slice，导致复出底层数组开辟了新的空间，a和b已经不一样了
	*/

	b = append(b, 10)
	b[1] = 11
	fmt.Println("b: ", b)

	fmt.Println("a: ", a)

	t := fmt.Sprintf("%v", a)
	fmt.Printf("xw: %v", t)

}
