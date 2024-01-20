package main

import "fmt"

// https://www.youtube.com/watch?v=ryHn-zHj7E8&list=PLc2Vkg57qmuRNHp6NNvYRVgg3OP-b5E_v&index=17&ab_channel=ThisIsIT

func main() {
	fmt.Println("Slices 3")
	getSlice()
	copySlice()
	deleteSliceElem()
}

func getSlice() {
	fmt.Println("getSlice:")
	intArr := [...]int{13, 14, 15, 16}
	fmt.Printf("Type: %T, Value: %#v\n", intArr, intArr)

	intSlice := intArr[1:3]
	fmt.Printf("Type: %T, Value: %#v\n", intSlice, intSlice)
	fmt.Printf("Lenght: %d, Capacity: %d\n\n", len(intSlice), cap(intSlice))

	fullSlice := intArr[:]
	fmt.Printf("Type: %T, Value: %#v\n", fullSlice, fullSlice)
	fmt.Printf("Lenght: %d, Capacity: %d\n\n", len(fullSlice), cap(fullSlice))
}

func copySlice() {
	fmt.Println("copySlice:")
	destination := make([]string, 0, 2)
	source := []string{"Vasya", "Petya", "Katya"}

	fmt.Println("Copied", copy(destination, source))
	fmt.Printf("Type: %T, Value: %#v\n", destination, destination)

	destination = make([]string, 2, 3)
	fmt.Println("Copied", copy(destination, source))
	fmt.Printf("Type: %T, Value: %#v\n", destination, destination)
	fmt.Printf("Lenght: %d, Capacity: %d\n\n", len(destination), cap(destination))

	destination = make([]string, len(source))
	fmt.Println("Copied", copy(destination, source))
	fmt.Printf("Type: %T, Value: %#v\n", destination, destination)
	fmt.Printf("Lenght: %d, Capacity: %d\n\n", len(destination), cap(destination))

	var defaultSlice []string // ТАК НЕ ДЕЛАТЬ! ПОЛУЧАЕМ nil МАССИВ!
	fmt.Println("Copied", copy(defaultSlice, source))
	fmt.Printf("Type: %T, Value: %#v\n", defaultSlice, defaultSlice)
	fmt.Printf("Lenght: %d, Capacity: %d\n\n", len(defaultSlice), cap(defaultSlice))

	rightCopy := append(make([]string, 0, len(source)), source...)
	fmt.Println("Copied", copy(rightCopy, source))
	fmt.Printf("Type: %T, Value: %#v\n", rightCopy, rightCopy)
	fmt.Printf("Lenght: %d, Capacity: %d\n\n", len(rightCopy), cap(rightCopy))
}

func deleteSliceElem() {
	fmt.Println("deleteSliceElem:")
	slice := []int{1, 2, 3, 4, 5}
	i := 2
	fmt.Printf("Type: %T, Value: %#v\n", slice, slice)
	fmt.Printf("Lenght: %d, Capacity: %d\n\n", len(slice), cap(slice))

	withAppend := append(slice[:i], slice[i+1:]...) // ломает исходный слайс
	fmt.Printf("Type: %T, Value: %#v\n", withAppend, withAppend)
	fmt.Println(slice)

	slice = []int{1, 2, 3, 4, 5}

	withCopy := slice[:i+copy(slice[:i], slice[i+1:])] // ломает исходный слайс
	fmt.Printf("Type: %T, Value: %#v\n", withCopy, withCopy)
	fmt.Println(slice)
}
