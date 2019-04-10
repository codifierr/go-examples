package main

import "fmt"

// Vertex which holds two field x and y
type Vertex struct {
	X,Y int
}

var (
	v1 = Vertex{1,2} // has type Vertex
	v2 = Vertex{X : 1} // Y:0 is implicit
	v3 = Vertex{} // X:0 and Y:0
	p = &Vertex{1,2} // has type *Vertex
)

func main(){
	fmt.Println(v1,v2,v3,p)
}
