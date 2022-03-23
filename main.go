package main

import (
	"fmt"

	"eden-farm-test/max"
)

func main() {
	arr1 := []int{1, 2, 3, 8, 9, 3, 2, 1} // 3
	arr2 := []int{1, 2, 1, 2, 2, 1}       // 2
	arr3 := []int{7, 1, 2, 9, 7, 2, 1}    // 2
	fmt.Println(max.MaxNumber(arr1))
	fmt.Println(max.MaxNumber(arr2))
	fmt.Println(max.MaxNumber(arr3))

}
