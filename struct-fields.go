package main

import "fmt"

type Vertex struct {
	X int
	Y int
}
func main(){
	v := Vertex{1,2}
	v.X = 4

	fmt.Println("X :",v.X)
	fmt.Println("Y :",v.Y)
}
