package Conditionals

import "fmt"

/*
SCOPE in GoLang
- import "fmt" or import f "fmt"
*/

// Package scopes
const done = false

var b int = 10

func main() {
	var task = "Running" // This is local to main() scope
	fmt.Println(task, done)

	const done = true //Local scope to main()
	fmt.Printf("<done> in main() is %v\n", done)
	f1()

	fmt.Println("Bye bye")
}

func f1() {
	fmt.Printf("done in f1() is %v\n", done) // This is <done> from package scope
	a := 10
	_ = a
}
