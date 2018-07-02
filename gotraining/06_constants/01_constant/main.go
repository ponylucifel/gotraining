package main

import "fmt"

const p = "death & taxes"

func main() {
	const q = 42
	// p = "hello" // This will raise an error
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

}

// a CONSTANT is a simple unchanging value
