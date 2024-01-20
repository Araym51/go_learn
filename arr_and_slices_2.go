package main

import "fmt"

/* https://www.youtube.com/watch?v=ecetBuEZ8oc&list=PLc2Vkg57qmuRNHp6NNvYRVgg3OP-b5E_v&index=16&ab_channel=ThisIsIT */

func main() {
	fmt.Println("Slices 2")
	variadicFuncs()
	convertToArrayPointer()
}

func variadicFuncs() {
	showAllElems(1, 2, 3)
	showAllElems(4, 5, 6, 7, 8)

	firstSlice := []int{9, 10, 11, 12}
	secondSlice := []int{13, 14, 15, 16}
	showAllElems(firstSlice...)
	showAllElems(secondSlice...)

	newSlice := append(firstSlice, secondSlice...)
	fmt.Printf("Type: %T, Value: %#v\n", newSlice, newSlice)
}

func showAllElems(values ...int) {
	for _, val := range values {
		fmt.Println("Value", val)
	}
	fmt.Println()
}


/*
Структура слайса:
type _slice struct{
	elements unsafe.Pointer // указатель на массив с определенным типом данных
	len int
	cap int
}
*/

func convertToArrayPointer(){
	initialSlice := []int{1,2}
	fmt.Printf("Type: %T, Value: %#v\n", initialSlice, initialSlice)
	fmt.Printf("Lenght: %d, Capacity: %d\n\n", len(initialSlice), cap(initialSlice))

	initArray :=(*[2]int)(initialSlice) // длинна массива должна совпадать с длинной слайса! Иначе паника!
	fmt.Printf("Type: %T, Value: %#v\n", initArray, initArray)
	fmt.Printf("Lenght: %d, Capacity: %d\n\n", len(initArray), cap(initArray))
}

