package main

import "fmt"
import "runtime"
import "time"

func main() {
	sum :=0
	for i := 0; i < 10 ; i++ {
		sum += i
	}
	
	fmt.Println(sum)
	
	//The init and post statement are optional.
	sum2 :=1
	for ; sum2 < 1000; {
		sum2++
	}
	fmt.Println(sum2)
	
	//While : can drop the semicolons:C's while is spelled for in Go
	sum3 := 1
	for sum3 < 1000 {
		sum3++
	}
	
	//it loops forever..
	// for {
	// 	// if sum3 == 1000 { continue }
	// }
	
	if sum3 == 1000 {
		fmt.Println("sum3", sum3)
	}
	
	switch os := runtime.GOOS; os {
		// case "darwin": fmt.Println("OS X.")
		// case "linux": fmt.Println("OS X.")
		default: fmt.Println("os:",os)
	}
	//switch with condition
	switch today := time.Now().Weekday(); time.Saturday {
		case today : fmt.Println("It's today..")
		default : fmt.Println("In's not today..")
	}
	//switch with no condition
	t:= time.Now()
	switch {
		case t.Hour() < 12: fmt.Println("good morning")
		case t.Hour() < 17: fmt.Println("Good afternoon.")
		default:fmt.Println("Good evening")
	}
	
	//A defer statement defers the execution of a function until the surrounding function returns.
	defer fmt.Println("world")
	
	fmt.Println("hello")
	//Stacking defers
	//Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in lase-in-first-out order.
	fmt.Println("starting.. counting")
	for ii:=0;ii<10; {
		defer fmt.Print(ii)
		ii++
	}
	fmt.Println()
	fmt.Println("done.. counting")
	
	
}









