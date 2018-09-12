package main

import (
	"fmt"
	"math"
)


// alternate classes, Define methods on types.
// A method is a function with a special receiver args.
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Declare a method on non-struct types, too
type SingleVertex float64

func (s SingleVertex) SingleAbs() float64 {
	if( s  < 0 ) {
		return float64(-s)
	}
	return float64(s);
}

//Declare methods with pointer receivers
//Syntax : *T.
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X *f
	v.Y = v.Y *f
}


func (v *Vertex)Scale (f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}



func main() {
	v := Vertex {3,4}
	ScaleFunc(&v, 10);
	fmt.Println(v.Abs())
	
	s := SingleVertex (-math.Sqrt2)
   	fmt.Println(s.SingleAbs())	
	
	//*T (pointer receivers) 의 위치에 따라 사용법이 다르다ㅣ.
	// ScaleFunc(v, 10)  // compile error
	ScaleFunc(&v, 10) // OK
	
	//arguments 에 포인터 리시버를 쓰는 경우 반드시 포인터가 해당 매개변수에 할당되어야 한다.
	
// 메소드에 정의된 포인터 리시버는 값 또는 포인터 둘다 할당이 가능하다.
	v.Scale(10)
	
	p := &v
	p.Scale(10)
	
	
	
	
	
	
}


