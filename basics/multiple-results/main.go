package main

import(
	"fmt"
)

func main(){
	a,b:= swap(12,15)
	fmt.Println("Swapped 12 and 15 : ",a,b)
}

func swap(x,y int) (int,int){
	return y,x
}