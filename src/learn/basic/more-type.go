package main

import "fmt"

func main() {
	i, j := 42, 2701
	
	
	p := &i //i의 값의 주소를 포인터에 할당한다.
	
	*p = 21 //포인터 p를 통해 i를 설정한다.
	fmt.Println(i)
	
	p = &j //j의 값의 주소를 포인터에 할당한다.
	
	*p = *p / 37  //포인터 p를 통해 j를 설정한다.
	fmt.Println(j)
	
}

//A pointer holds the memory address of a value. 포이터는 값을 메모리 주소에 할당합니다.
//The type *T is a pointer to a T value. Its zero value is nil. 유형 *T는 T값에 대한 포인터 입니다. 0값은 0입니다.
//The & operator generates a pointer to its operand. &연산자는 포인터에 피연산자의 값을 할당합니다.
//The * operator denotes the pointer's underlying value. *연산자는 포인터의 가본값을 나타냅니다.
// This is known as "dereferencing" or "indirecting"
//> 포인터로 하여금 할당된 변수의 값을 변경하는것 을  "역 참조" "간접(참조)"라고 한다.
