package main
import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
type MyFloat float64


func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	//v := Vertex{3, 4}
	//fmt.Println(v.Abs())

	//f := MyFloat(-math.Sqrt2)
	//fmt.Println(f.Abs())

	//v.Scale(10)
	//fmt.Println(v.Abs())



	//p := &Vertex{4, 3}
	//p.Scale(3)
	//ScaleFunc(p, 8)

	//fmt.Println(v, p)




	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))

	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())



}
