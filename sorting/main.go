package main

import (
	"fmt"
	"sort"
	strconv "strconv"
)

func main() {
	arr := []string{"6",
		"31415926535897932384626433832795",
		"1",
		"3",
		"10",
		"3",
		"5"}
	fmt.Println("arr: ", arr)
	//sort.Sort(ByLen(arr))
	fmt.Println("sorted arr :", bigSorting(arr))
	//fmt.Println("sorted arr :", arr)
}

func bigSorting(unsorted []string) []string {
	sort.Sort(ByIntValue(unsorted))
	result := unsorted
	return result
}

// ByIntValue custom type
type ByIntValue []string

func (a ByIntValue) Len() int {
	return len(a)
}

func (a ByIntValue) Less(i, j int) bool {
	z, _ := strconv.Atoi(a[i])
	x, _ := strconv.Atoi(a[j])
	return z < x
}

func (a ByIntValue) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
