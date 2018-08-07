package main

// import (
// 	"fmt"
// 	"math/rand"
// )
import "fmt"
import "math"

//declares a list of variables; the type is last.
// var c, python, java bool
//variables with initializers
var i, j int = 1, 2


func main() {
	// fmt.Println("My favorite number is ", rand.Intn(10))
	fmt.Printf("now you have %g problems.", math.Sqrt(7))
	// fmt.Printf("add : " , add(10,1))
	// fmt.Printf("add2 : " , add2(10,11))
	
	a, b := swap("multiple", "returns.")
	fmt.Println(a,b)
	fmt.Println(split(17))
	k := 3
	c, python, java := true, false, "no!"
	
	fmt.Println(k, c, python, java)
}

//the type of function cames after the variable name.
func add(x int, y int) int {
	return x + y
}

//function parameters share a type, but the last
func add2(x, y int) int {
	return x + y
}

//multiple results
func swap(x, y string)(string, string) {
	return x, y
}

//Named return values
func split(sum int)(x, y int) {
	x = sum * 4/9
	y = sum - x
	return
}