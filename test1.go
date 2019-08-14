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

func main() {
	fmt.Println("Hello This is my first GO program")
//	print(os.Hostname())
    fmt.Println("The time is now is", time.Now())
//	fmt.Println("The hostname of your machine is", os.Hostname())
    fmt.Println("the random number is", rand.Intn(18))
	fmt.Println("The random number with rand seed is", rand.Seed)
	fmt.Printf("Now we have %g problems", math.Sqrt(81))
    fmt.Println("The Pi value is ", math.Pi)
}