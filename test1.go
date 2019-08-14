package main

// This is an import block it allows multiple packages to be imported
// without repeating the "import" keyword.
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)
/*
The main function, when part of the main package,
identifies the entry point of an application.
 */

func bkr_add(x int, y int) int {
	return x + y
}

func bkr_multiply(a, b int) int {
	return a * b
}

func bkr_swap(x, y string) (string, string)  {
	return y, x
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
}