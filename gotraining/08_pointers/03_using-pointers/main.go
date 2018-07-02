package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  // prints 43
	fmt.Println(&a) // prints address of a

	var b *int = &a

	fmt.Println(b)  // prints b which is address of a
	fmt.Println(*b) // prints value that resides in where b pointing to, 43

	*b = 42        // b says, "The value at this address, change it to 42"
	fmt.Println(a) // 42
	fmt.Println(b) // b (a's memory address) is still the same

	// one thing to note is that everything is PASS BY VALUE in golang
}
