// hackerrank problems practice workspace
package main

import (
	"fmt"
	"sort"
	"time"
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

	//ar := []int32{3, 2, 1, 7, 6, 8, 8, 3}
	//println(birthdayCakeCandles(ar))

	// time := "07:05:45PM"
	// println(timeConversion(time))

	// a := []int32{1, 2, 3, 4, 1}
	// b := []int32{3, 4, 1, 2, 1, 3}
	// c := longestCommonSubsequence(a, b)
	// fmt.Printf("LongestCommonSubsequence : %v\n", c)

	// h := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}
	// word := "zaba"
	// fmt.Printf("DesignerPdfViewer : %v\n", designerPdfViewer(h, word))

	// a := []int32{3, 4, 5}
	// k := int32(2)
	// queries := []int32{1, 2}

	// fmt.Printf("CircularArrayRotation : %v\n", circularArrayRotation(a, k, queries))

	grades := []int32{73, 67, 38, 33}
	fmt.Printf("GradingStudents : %v\n", gradingStudents(grades))

}

/*
 * Complete the gradingStudents function below.
 */
func gradingStudents(grades []int32) []int32 {
	var result []int32
	for _, v := range grades {
		fmt.Printf("Value at once : %v\n", v%10)
		if v > 38 {
			if diff := v % 10; diff < 3 {
				result = append(result, v+5-diff)
			} else {
				result = append(result, v)
			}
		} else {
			//fmt.Printf("Value : %v\n", v)
			result = append(result, v)
		}
	}
	return result
}

// Complete the angryProfessor function below.
func angryProfessor(k int32, a []int32) string {

	return "nil"
}

// Complete the circularArrayRotation function below.
func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	if k < 0 || len(a) == 0 {
		return a
	}
	r := int32(len(a)) - k%int32(len(a))
	a = append(a[r:], a[:r]...)
	result := make([]int32, len(queries))
	for i, v := range queries {
		result[i] = a[v]
	}
	return result
}

// Complete the designerPdfViewer function below.
func designerPdfViewer(h []int32, word string) int32 {
	var maxH int32
	for _, val := range word {
		if h[val-'a'] > maxH {
			maxH = h[val-'a']
		}
	}
	return int32(len(word)) * maxH
}

// Complete the longestCommonSubsequence function below.
func longestCommonSubsequence(a, b []int32) []int32 {
	var c []int32
	l, c := longestCommonSubsequenceHelper(a, b, int32(len(a)), int32(len(b)), c)
	fmt.Println(l)
	return c
}

func longestCommonSubsequenceHelper(a, b []int32, m, n int32, c []int32) (int32, []int32) {
	if m == 0 || n == 0 {
		//close(ch)
		return 0, c
	}
	for _, val := range c {
		if val == a[m-1] {
			return 1, c
		}
	}
	if a[m-1] == b[n-1] {
		c = append(c, a[m-1])
		k, c := longestCommonSubsequenceHelper(a, b, m-1, n-1, c)
		return int32(1) + k, c
	}
	i, c := longestCommonSubsequenceHelper(a, b, m-1, n, c)
	j, c := longestCommonSubsequenceHelper(a, b, m, n-1, c)
	if i > j {
		//c = append(c, a[m-1])
		return i, c
	}
	//c = append(c, a[n-1])
	return j, c

}

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
	/*
	 * Write your code here.
	 */
	layoutFrom := "03:04:05PM"
	layoutTo := "15:04:05"
	time, err := time.Parse(layoutFrom, s)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return time.Format(layoutTo)
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
