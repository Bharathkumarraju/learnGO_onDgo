package main

// This is an import block it allows multiple packages to be imported
// without repeating the "import" keyword.
import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)
/*
The main function, when part of the main package,
identifies the entry point of an application.
 */

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 -1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

//var c,python,java bool
var i, j int = 1, 2

func bkr_add(x int, y int) int {
	return x + y
}

func bkr_multiply(a, b int) int {
	return a * b
}

func bkr_swap(x, y string) (string, string)  {
	return y, x
}

func bkr_spilt(sum int) (x, y int)  {
	x = sum * 4/9
	y = sum - x
	return
}



func main() {
	fmt.Println("Hello This is my first GO program")
//	print(os.Hostname())
    fmt.Println("The time is now is", time.Now())
//	fmt.Println("The hostname of your machine is", os.Hostname())
    fmt.Println("the random number is", rand.Intn(18))
	fmt.Println("The random number with rand seed is", rand.Seed)
	fmt.Printf("Now we have %g problems", math.Sqrt(81))
    fmt.Println("The Pi value is ", math.Pi)
	fmt.Println(bkr_add(287485, 39539354))
	fmt.Println(bkr_multiply(3565636, 3453453))
	a, b := bkr_swap("SRI ANJANEYAM", "PRASANNA ANAJANEYAM")
	fmt.Println(a, b)
	fmt.Println(bkr_spilt(18))
	//var z int
//	fmt.Println(i, c, python, java)
	var c, python, java = true, "yes!", "no!"
	fmt.Println(i, j, c, python, java)
	var i, j int = 18, 19
	k := 46675
	fmt.Println(i, j, k, c, python, java)
	fmt.Printf("Type: %T Value: %v \n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v \n", MaxInt, MaxInt)
	fmt.Printf("Type: %T value: %v \n", z, z)
	var f float64
	var y bool
	var s int
	var n string
	fmt.Printf("%v %v %v %q", f, y, s, n)
}