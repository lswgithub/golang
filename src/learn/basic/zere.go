package main

import "fmt"


//Constants are declared like variables, but with the const keyword.
//Constants can be character, string, boolean, or numeric values
//Constants can't be declared using the := syntax.
const Pi = 3.14


func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i , f, b, s)
	
	ii := 42
	ff := float64(ii)
	uu := uint(ff);
	
	fmt.Printf("%v %v %v %q\n", ii, ff, uu)
}