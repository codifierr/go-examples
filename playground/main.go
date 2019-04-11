// hackerrank problems practice workspace
package main

import (
	"fmt"
	"sort"
)

func main() {
	// string comparision
	// s1 := "test string 1 test"
	// s2 := "test string 2 test"

	//fmt.Printf("String is equal ? %t",compareString(s1,s2))

	// traverse array of array
	// arr := [][]int32{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{9, 8, 9},
	// }

	// arr1 := [][]int32{
	// 	{-1, 1, -7, -8},
	// 	{-10, -8, -5, -2},
	// 	{0, 9, 7, -1},
	// 	{4, 4, -2, 1},
	// }
	// fmt.Println(diagonalDifference(arr))
	// fmt.Println(diagonalDifference(arr1))

	/*numarr := []int32{6, -4, 3, -9, 0, 4, 1}
	plusMinus(numarr)*/
	//staircase(5)

	/*arr := []int{1, 2, 3, 4, 5}
	miniMaxSum(arr)*/

	ar := []int32{3, 2, 1, 7, 6, 8, 8, 3}
	println(birthdayCakeCandles(ar))

}

// Complete the birthdayCakeCandles function below.
func birthdayCakeCandles(ar []int32) int32 {
	arrlen := len(ar)
	var maxval int32
	var count int32
	for i := arrlen - 1; i > 0; i-- {
		if ar[i] > maxval {
			maxval = ar[i]
		}
	}

	for _, v := range ar {
		if maxval == v {
			count++
		}
	}
	return count
}

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int) {
	sort.Ints(arr)
	arrlen := len(arr)
	minslice := arr[:arrlen-1]
	maxslice := arr[1:arrlen]

	var v1, v2 int

	for i := 0; i < arrlen-1; i++ {
		v1 += minslice[i]
		v2 += maxslice[i]
	}
	fmt.Printf("%d %d", v1, v2)

}

// Complete the staircase function below.
func staircase(n int32) {
	var i, j, k int32
	for i = 1; i <= n; i++ {
		for k = n - i; k > 0; k-- {
			fmt.Printf(" ")
		}
		for j = 0; j < i; j++ {
			fmt.Printf("#")
		}
		fmt.Printf("\n")
	}
}

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
	posCounter := 0
	negCounter := 0
	zeroCounter := 0
	arrlen := len(arr)

	for _, v := range arr {
		switch {
		case v > 0:
			{
				posCounter++
			}
		case v < 0:
			{
				negCounter++
			}
		case v == 0:
			{
				zeroCounter++
			}
		}
	}
	fmt.Printf("%f\n", float64(posCounter)/float64(arrlen))
	fmt.Printf("%f\n", float64(negCounter)/float64(arrlen))
	fmt.Printf("%f\n", float64(zeroCounter)/float64(arrlen))
}

func diagonalDifference(arr [][]int32) int32 {
	arrlen := len(arr)
	var diff int32
	for i, j := 0, arrlen-1; i < arrlen && j >= 0; i, j = i+1, j-1 {
		diff += arr[i][i] - arr[i][j]
	}
	if diff < 0 {
		return -diff
	}
	return diff
}

func compareString(s1, s2 string) bool {
	min := len(s1)
	if len(s2) < min {
		min = len(s2)
	}
	diff := 0 // make([]int,0,min)
	for i := 0; i < min; i++ {
		diff += int(s1[i]) - int(s2[i])
	}

	switch {
	case diff == 0:
		{
			return true
		}
	}
	return false
}
