package main

import "fmt"

// Vertex which holds two field x and y
type Vertex struct {
	// A struct is a collection of fields.
	X int
	Y int
}

func main(){
	fmt.Println(Vertex{1,2})
	v := Vertex{1,2}

	//Struct fields are accessed using a dot.
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v.Y)

	fmt.Printf("%+v\n", v)

	//Struct fields can be accessed through a struct pointer.
	p := &v
	p.X = 1e9

	// below two statements does same
	fmt.Println((*p).Y)
	fmt.Println(p.Y)

	//print vertex
	fmt.Println(v)
}