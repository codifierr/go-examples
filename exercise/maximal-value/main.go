package main

import "fmt"

func main() {
	nums := []int{16, 8, 35, 87, 3, 75, 90}
	max := nums[0]
	for _, v := range nums[1:] {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}
