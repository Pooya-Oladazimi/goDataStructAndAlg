package main

import "fmt"

func matrixMuliplyExp() {
	a1 := [2][5]int{
		{1, 2, 3, 4, 5},
		{10, 20, 1, 1, 1},
	}

	a2 := [5][2]int{
		{-1, 20},
		{100, 30},
		{20, 90},
		{-3, -9},
		{10, 8},
	}

	var res [len(a1)][len(a2[0])]int

	for i := 0; i < len(a1); i++ {
		for j := 0; j < len(a1[0]); j++ {
			for k := 0; k < len(a2[0]); k++ {
				res[i][k] += a1[i][j] * a2[j][k]
			}
		}
	}

	fmt.Println(res)

}
