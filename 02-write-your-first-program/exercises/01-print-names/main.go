// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print names
//
//  Print your name and your best friend's name using
//  Println twice
//
// EXPECTED OUTPUT
//  YourName
//  YourBestFriendName
//
// BONUS
//  Use `go run` first.
//  And after that use `go build` and run your program.
// ---------------------------------------------------------

func main() {
	name := "Dũng"
	age := 18

	fmt.Println("Hello Joseph")
	fmt.Println(`Hello Joseph
	Dung`)

	fmt.Printf("My name is %s and I am %d years old.\n", name, age)
}
