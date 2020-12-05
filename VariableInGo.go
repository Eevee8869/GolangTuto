package main

import "fmt"

func main() {

	// `var` declares 1 or more variables.
	var a = "ExampleString"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding
	// initialization are _zero-valued_. For example, the
	// zero value for an `int` is `0`.
	var e int
	fmt.Println(e)

	// The `:=` syntax is shorthand for declaring and
	// initializing a variable, e.g. for
	// `var f string = "apple"` in this case.
	f := "take my advice and say java is madarjava"
	fmt.Println(f)

/*Go also has support for constants. Constants are basically variables whose values cannot be changed later*/
 const x string = "Hello World"
  fmt.Println(x)

//and now if u try to change value like
x = "Some other string"
//compiler will throw compile-time error:
//.\VariableInGo.go:7: cannot assign to x
 

}
