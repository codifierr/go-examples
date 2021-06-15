package main

import (
	fmt "fmt"

	arraylist "github.com/emirpasic/gods/lists/arraylist"
	utils "github.com/emirpasic/gods/utils"
)

func main() {
	list := arraylist.New()
	list.Add("a")
	list.Add("c", "b")
	list.Sort(utils.StringComparator)
	v, ok := list.Get(0)
	fmt.Printf("Value :%v and Present? %t\n", v, ok)

	ok = list.Contains("a", "b", "c")

	fmt.Printf("Present? %t\n", ok)

	list.Swap(0, 1) // ["b","a",c"]
	fmt.Print("Post swap : ", list.String())
	list.Remove(2) // ["b","a"]
	list.Remove(1) // ["b"]
	list.Remove(0) // []
	list.Remove(0) // [] (ignored)
	fmt.Println("\nPost remove operation : ", list.String())
	empty := list.Empty() // true
	size := list.Size()   // 0
	fmt.Printf("Empty ? %t", empty)
	fmt.Printf("\nSize : %d", size)
	list.Add("a")       // ["a"]
	list.Clear()        // []
	list.Insert(0, "b") // ["b"]
	list.Insert(0, "a") // ["a","b"]
	fmt.Printf("\nPost insertion : %s", list.String())
}
