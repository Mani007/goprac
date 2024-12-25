package main

import "fmt"

func main() {
	fmt.Println("Variables here")
	// String variables
	var username string = "ABC"
	fmt.Println(username)
	fmt.Printf(" Variable type is %T\n", username)
	var isLog bool = true
	fmt.Println(isLog)
	fmt.Printf(" Variable type is %T\n", isLog)
	var someint int32 = 25
	fmt.Println(someint)
	fmt.Printf(" Variable type is %T\n", someint)
	// Default value for un initialized int
	var defaultval int32
	fmt.Println(defaultval) // it will be zero
	fmt.Printf(" Variable type is %T\n", defaultval)
	// inplicit type
	var webadd = "abc" // you cannot use '' insted use ""
	println(webadd)    // its string
	webadd = "new"     // you can chage the value but not the type
	//webadd = 25 // error type cannot be chaged
}
