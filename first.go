package main

import (
	"fmt"
	"math"
	"math/rand"
	"math/cmplx"
	"time"
)
const Pi = 3.14
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i))


func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}










func main() {
	fmt.Println("My favorite number is", rand.Intn(2))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(16))
	fmt.Println(math.Pi)
		fmt.Println(add(20, 30))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(20))


		var i, j int = 1, 2
		k := 3
		c, python, java := true, false, "no!"

		fmt.Println(i, j, k, c, python, java)


	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var p int
	var f float64
	var d bool
	var s string
	fmt.Printf("%v %v %v %q\n", p, f, d, s)


	var x, y int = 3, 4
	var h float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(h)
	fmt.Println(x, y, z)

	v := 42 // int
	fmt.Printf("v is of type %T\n", v)

	w:= 42.5 // float64
	fmt.Printf("w is of type %T\n", w)

	o := 0.78+0.5i // complex128 ,i should
	fmt.Printf("o is of type %T\n", o)





//Constants are declared like variables, but with the const keyword.
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	//Numeric Constants
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

//For
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
//he init and post statements are optional.
	tot := 1
	for ; tot< 1000; {
		tot += tot
	}
	fmt.Println(tot)
	// IF
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	//switch case
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	//switch with no condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
//defer
	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")



}


