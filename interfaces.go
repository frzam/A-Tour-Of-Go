package main

import (
	"fmt"
	"math"
)

type MyFloat float64

type Abser interface{
	Abs()float64
}

type Vertex struct {
	X, Y float64
}

func (f MyFloat) Abs()float64{
	if f<0 {
		return float64(-f)
}
	return float64(f)
}

func (v *Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}

func main(){
	var a Abser
	v := Vertex{3,4}
	f := MyFloat(-math.Sqrt(2))
	
	a = f
	fmt.Println(a.Abs())
	a = &v

	fmt.Println(a.Abs())

}
