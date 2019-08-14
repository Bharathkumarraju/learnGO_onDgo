package main

// This is an import block it allows multiple packages to be imported
// without repeating the "import" keyword.
import (
	"fmt"
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

}