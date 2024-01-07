package main

import "fmt"


func square(num *int) {
		*num *= *num
	}

func main() {
	fmt.Println("pointers:")
	// объявление указателя
	var intPointer *int
	fmt.Printf("%T %#v \n", intPointer, intPointer) // выведет: *int (*int)(nil) 

	var a int = 7
	fmt.Printf("%T %#v \n", a, a)

	var pointerA = &a
	fmt.Printf("%T %#v %#v \n", pointerA, pointerA, *pointerA) // *int (*int)(0xc00000a0c8) 7 
	// изменения значения в разименованном указателе:
	*pointerA = 3
	fmt.Printf("%T %#v %#v \n", pointerA, pointerA, *pointerA) //*int (*int)(0xc000096068) 3 

	num := 2

	square(&num)
}