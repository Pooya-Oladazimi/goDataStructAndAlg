package main

import "fmt"
import "unicode/utf8"

func main() {
	araryExp()
	sliceExp()
	matrixMuliplyExp()
}

func sliceExp() {
	var s1 []int
	fmt.Println(s1) // empty array. no default value since the slice is empty with length 0.

	s2 := make([]int, 0, 5) // slice with length 0 and capacity 5
	fmt.Println(s2)

	s3 := make([]int, 2, 5) // slice with length 2 and capacity 5. The first two elements get initiated with the default value 0. --> [0,0]
	fmt.Println(s3)

	s4 := []int{1, 2, 3}
	fmt.Println(s4)

	twoDs := make([][]int, 3) // two dimensional slice with 3 rows. columns count can be anything.
	fmt.Println(twoDs)

	twoDs2 := [][]int{
		{1, 2, 3},
		{-1, -1, -1, -1, -1},
	}
	fmt.Println(twoDs2)
  
}

func araryExp() {
	var arr1 [5]int
	fmt.Println(arr1[2]) // Prints 0 --> Go arrays are by default zero-valued

	arr2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(arr2[2]) // Prints 30

	arr3 := [...]int{10, 20, 30, 40, 50}
	fmt.Println(arr3[1]) // Prints 20

	arr4 := [...]string{"oo", 3: "pp"} // index 3 will be "pp"
	fmt.Println(arr4)                  // default value for string array would be ""

	var twoDArr1 [2][5]int // 2 * 5
	fmt.Println(twoDArr1)

	twoDArr2 := [2][5]int{
		{1, 2, 3, 4, 5},
		{10, 20, 30, 40, 50},
	}
	fmt.Println(twoDArr2)

	twoDArr3 := [...][5]int{ // second dimention length needs to be explicitly given.
		{-1, -2, -3, -4, -5},
		{10, 20, 30, 40, 50},
	}
	fmt.Println(twoDArr3)
}
