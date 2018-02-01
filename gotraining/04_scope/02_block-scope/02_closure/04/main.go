package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps is limit the scope of variables used by multiple functions
without closure, for two ormore funcs to have access to the same variable,
that variable would need to be package scope

anonymous function
a function without a name.

func expression
assigning a function to a variable
*/
