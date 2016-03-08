package main

import (
	"fmt"
	"time"

)

func main() {

	fmt.Printf("Hello, a GO Lanuage.. \n")
	fmt.Printf("The time is ", time.Now())


	//var age int = 40
	//fmt.Printf("my age is " , age)

	const pi float64 = 3.14215484215
	//var isOver40 bool = true

	fmt.Printf("\n %f.3f \n", pi)


	/*i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	*/
favNum3 := [5] float64{1,2,3,4,5}

	for i, value:=range favNum3 {
		fmt.Println(i,value)
	}

}

