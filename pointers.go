package main

import (
	"fmt"
	"Strings"
)

type Vertex struct {
	X int
	Y int
}
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func printSlice(d []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(d), cap(d), d)
}
func main()  {
//pointers
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	//Struct
	fmt.Println(Vertex{1,2})
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v1, p, v2, v3)

	//Array
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	 m := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(m, b)
	fmt.Println(names)


	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	k := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(k)

	d := []int{2, 3, 5, 7, 11, 13}
	printSlice(d)
    d=d[:0]
	printSlice(d)
	d = d[1:4]
	fmt.Println(d)
	printSlice(d)
	d = d[:2]
	fmt.Println(d)
	printSlice(d)
	d = d[1:]
	fmt.Println(d)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
